package derive

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"io"

	"github.com/ethereum-optimism/optimism/op-service/eth"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const (
	// SYSCOIN
	appendSequencerBatchMethodFunction = "appendSequencerBatch(bytes32[])"
	appendSequencerBatchMethodName     = "appendSequencerBatch"
)

type DataIter interface {
	Next(ctx context.Context) (eth.Data, error)
}

type L1TransactionFetcher interface {
	InfoAndTxsByHash(ctx context.Context, hash common.Hash) (eth.BlockInfo, types.Transactions, error)
	FetchReceipts(ctx context.Context, blockHash common.Hash) (eth.BlockInfo, types.Receipts, types.Transactions, error)
	GetBlobFromCloud(vh common.Hash) ([]byte, error)
	GetBlobFromRPC(vh common.Hash) ([]byte, error)
}

// DataSource is a fault tolerant approach to fetching data.
// The constructor will never fail & it will instead re-attempt the fetcher
// at a later point.
type CalldataSource struct {
	// Internal state + data
	open bool
	data []eth.Data
	// Required to re-attempt fetching
	ref     eth.L1BlockRef
	dsCfg   DataSourceConfig
	fetcher L1TransactionFetcher
	log     log.Logger

	batcherAddr common.Address
	// SYSCOIN
	batchInboxABI              *abi.ABI
	appendSequencerFunctionSig []byte
}

// NewCalldataSource creates a new calldata source. It suppresses errors in fetching the L1 block if they occur.
// If there is an error, it will attempt to fetch the result on the next call to `Next`.
func NewCalldataSource(ctx context.Context, log log.Logger, dsCfg DataSourceConfig, fetcher L1TransactionFetcher, ref eth.L1BlockRef, batcherAddr common.Address) DataIter {
	// SYSCOIN info
	_, receipts, txs, err := fetcher.FetchReceipts(ctx, ref.Hash)
	if err != nil {
		return &CalldataSource{
			open:                       false,
			ref:                        ref,
			dsCfg:                      dsCfg,
			fetcher:                    fetcher,
			log:                        log,
			batcherAddr:                batcherAddr,
			batchInboxABI:              dsCfg.batchInboxABI,
			appendSequencerFunctionSig: dsCfg.appendSequencerFunctionSig,
		}
	} else {
		// SYSCOIN
		dataSrc := DataFromEVMTransactions(ctx, fetcher, dsCfg, batcherAddr, receipts, txs, log.New("origin", ref))
		if dataSrc == nil {
			return &CalldataSource{
				open:                       false,
				ref:                        ref,
				dsCfg:                      dsCfg,
				fetcher:                    fetcher,
				log:                        log,
				batcherAddr:                batcherAddr,
				batchInboxABI:              dsCfg.batchInboxABI,
				appendSequencerFunctionSig: dsCfg.appendSequencerFunctionSig,
			}
		}
		return &CalldataSource{
			open: true,
			data: dataSrc,
		}
	}
}

// Next returns the next piece of data if it has it. If the constructor failed, this
// will attempt to reinitialize itself. If it cannot find the block it returns a ResetError
// otherwise it returns a temporary error if fetching the block returns an error.
func (ds *CalldataSource) Next(ctx context.Context) (eth.Data, error) {
	if !ds.open {
		// SYSCOIN
		if _, receipts, txs, err := ds.fetcher.FetchReceipts(ctx, ds.ref.Hash); err == nil {
			ds.data = DataFromEVMTransactions(ctx, ds.fetcher, ds.dsCfg, ds.batcherAddr, receipts, txs, log.New("origin", ds.ref))
			// SYSCOIN
			if ds.data == nil {
				return nil, NewTemporaryError(fmt.Errorf("failed to open calldata cloud source"))
			}
			ds.open = true
		} else if errors.Is(err, ethereum.NotFound) {
			return nil, NewResetError(fmt.Errorf("failed to open calldata source: %w", err))
		} else {
			return nil, NewTemporaryError(fmt.Errorf("failed to open calldata source: %w", err))
		}
	}
	if len(ds.data) == 0 {
		return nil, io.EOF
	} else {
		data := ds.data[0]
		ds.data = ds.data[1:]
		return data, nil
	}
}

// SYSCOIN DataFromEVMTransactions filters all of the transactions and returns the calldata from transactions
// that are sent to the batch inbox address from the batch sender address.
// This will return an empty array if no valid transactions are found.
func DataFromEVMTransactions(dsCfg DataSourceConfig, batcherAddr common.Address, txs types.Transactions, log log.Logger) []eth.Data {
	out := []eth.Data{}
	for _, tx := range txs {
		if isValidBatchTx(tx, dsCfg.l1Signer, dsCfg.batchInboxAddress, batcherAddr, log) {
			out = append(out, tx.Data())
		}
	}
	return out
}
