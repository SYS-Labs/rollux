package fetch

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"path"
	"reflect"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type TransactionWithMetadata struct {
	TxIndex     uint64             `json:"tx_index"`
	InboxAddr   common.Address     `json:"inbox_address"`
	BlockNumber uint64             `json:"block_number"`
	BlockHash   common.Hash        `json:"block_hash"`
	BlockTime   uint64             `json:"block_time"`
	ChainId     uint64             `json:"chain_id"`
	Sender      common.Address     `json:"sender"`
	ValidSender bool               `json:"valid_sender"`
	Frames      []derive.Frame     `json:"frames"`
	FrameErr    string             `json:"frame_parse_error"`
	ValidFrames bool               `json:"valid_data"`
	Tx          *types.Transaction `json:"tx"`
}

type Config struct {
	Start, End   uint64
	ChainID      *big.Int
	BatchInbox   common.Address
	BatchSenders map[common.Address]struct{}
	OutDirectory string
}

const (
	// SYSCOIN
	appendSequencerBatchMethodFunction = "appendSequencerBatch(bytes32[])"
	appendSequencerBatchMethodName     = "appendSequencerBatch"
)

// Batches fetches & stores all transactions sent to the batch inbox address in
// the given block range (inclusive to exclusive).
// The transactions & metadata are written to the out directory.
func Batches(client *ethclient.Client, config Config) (totalValid, totalInvalid int) {
	if err := os.MkdirAll(config.OutDirectory, 0750); err != nil {
		log.Fatal(err)
	}
	number := new(big.Int).SetUint64(config.Start)
	signer := types.LatestSignerForChainID(config.ChainID)
	for i := config.Start; i < config.End; i++ {
		valid, invalid := fetchBatchesPerBlock(client, number, signer, config)
		totalValid += valid
		totalInvalid += invalid
		number = number.Add(number, common.Big1)
	}
	return
}

// fetchBatchesPerBlock gets a block & the parses all of the transactions in the block.
func fetchBatchesPerBlock(client *ethclient.Client, number *big.Int, signer types.Signer, config Config) (validBatchCount, invalidBatchCount int) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	block, err := client.BlockByNumber(ctx, number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Fetched block: ", number)
	out := make([]byte, 0)
	for i, tx := range block.Transactions() {
		if tx.To() != nil && *tx.To() == config.BatchInbox {
			fmt.Printf("Found a transaction (%s) sent to the batch inbox\n", tx.Hash().String())
			sender, err := signer.Sender(tx)
			if err != nil {
				log.Fatal(err)
			}
			validSender := true
			if _, ok := config.BatchSenders[sender]; !ok {
				fmt.Printf("Found a transaction (%s) from an invalid sender (%s)\n", tx.Hash().String(), sender.String())
				invalidBatchCount += 1
				validSender = false
			}
			// Assuming you have a transaction object tx
			dataBytes := tx.Data()
			appendSequencerFunctionSig := crypto.Keccak256([]byte(appendSequencerBatchMethodFunction))[:4]
			if !reflect.DeepEqual(appendSequencerFunctionSig, dataBytes[:4]) {
				fmt.Println("DataFromEVMTransactions: append function not found as method signature", "index", i)
				invalidBatchCount += 1
				validSender = false
			}
			batchInboxABI, err := bindings.BatchInboxMetaData.GetAbi()
			batchData, err := batchInboxABI.Methods[appendSequencerBatchMethodName].Inputs.Unpack(dataBytes[4:])
			if err != nil {
				fmt.Println("DataFromEVMTransactions: Failed to unpack data for function call", "index", i, "err", err)
			}
			batchDataParam, ok := batchData[0].([][32]byte)
			if !ok {
				fmt.Println("DataFromEVMTransactions: Invalid item, expected [][32]byte", "batchData[0]", batchData[0], "len", len(batchDataParam), "receipt index", i)
			}
			numVHs := len(batchDataParam)
			for j := 0; j < numVHs; j++ {
				// get version hash from calldata and lookup data via syscoinclient
				// 1. get data from syscoin rpc
				vh := common.BytesToHash(batchDataParam[j][:])
				data, err := GetBlobFromCloud(vh)
				if err != nil {
					fmt.Println("DataFromEVMTransactions", "failed to fetch L1 block info and receipts", err)
					// instead of continuing this is a hard reset which means the entire set of blobs for this block/tx should be refetched
				}
				// check data is valid locally
				vhData := crypto.Keccak256Hash(data)
				if vh != vhData {
					fmt.Println("DataFromEVMTransactions", "blob data hash mismatch want", vh, "have", vhData)
				}
				fmt.Println("GetBlobFromCloud", "len", len(data), "vh", vh)
				out = append(out, data...)
			}
			validFrames := true
			frameError := ""

			// Convert the data to a hex string
			//dataHex := hexutil.Encode(dataBytes)

			//newdatahex := hexutil.Encode(data)
			//fmt.Println("NEW Transaction data in hex:", newdatahex)

			frames, err := derive.ParseFrames(out)
			fmt.Println(tx.Hash())
			if err != nil {
				fmt.Printf("Found a transaction (%s) with invalid data: %v\n", tx.Hash().String(), err)
				validFrames = false
				frameError = err.Error()
			}

			if validSender && validFrames {
				validBatchCount += 1
			} else {
				invalidBatchCount += 1
			}

			txm := &TransactionWithMetadata{
				Tx:          tx,
				Sender:      sender,
				ValidSender: validSender,
				TxIndex:     uint64(i),
				BlockNumber: block.NumberU64(),
				BlockHash:   block.Hash(),
				BlockTime:   block.Time(),
				ChainId:     config.ChainID.Uint64(),
				InboxAddr:   config.BatchInbox,
				Frames:      frames,
				FrameErr:    frameError,
				ValidFrames: validFrames,
			}
			filename := path.Join(config.OutDirectory, fmt.Sprintf("%s.json", tx.Hash().String()))
			file, err := os.Create(filename)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			enc := json.NewEncoder(file)
			if err := enc.Encode(txm); err != nil {
				log.Fatal(err)
			}
		}
	}
	return
}

func GetBlobFromCloud(vh common.Hash) ([]byte, error) {
	url := "https://poda.syscoin.org/vh/" + vh.String()[2:]
	var res *http.Response
	var err error
	// try 4 times incase of timeout or reset/hanging socket with 5+i second expiry each attempt
	for i := 0; i < 4; i++ {
		client := http.Client{
			Timeout: (5 + time.Duration(i)) * time.Second,
		}
		res, err = client.Get(url)
		if err != nil {
			continue
		} else {
			err = nil
			break
		}
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close() // we need to close the connection
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	txBytes, err := hex.DecodeString(string(body))
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}
