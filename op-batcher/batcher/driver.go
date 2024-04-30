package batcher

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"io"
	"io/ioutil"
	"math/big"
	_ "net/http/pprof"
	_ "strings"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-batcher/metrics"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	opclient "github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/txmgr"
	_ "github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	// SYSCOIN
	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
)

func ConvertToBytes32(slice []byte) eth.Bytes32 {
	var b32 eth.Bytes32
	copy(b32[:], slice)
	return b32
}

// BatchSubmitter encapsulates a service responsible for submitting L2 tx
// batches to L1 for availability.
type BatchSubmitter struct {
	Config // directly embed the config + sources

	txMgr txmgr.TxManager
	wg    sync.WaitGroup

	shutdownCtx       context.Context
	cancelShutdownCtx context.CancelFunc
	killCtx           context.Context
	cancelKillCtx     context.CancelFunc

	mutex   sync.Mutex
	running bool

	// lastStoredBlock is the last block loaded into `state`. If it is empty it should be set to the l2 safe head.
	lastStoredBlock eth.BlockID
	lastL1Tip       eth.L1BlockRef

	state *channelManager
}

// NewBatchSubmitterFromCLIConfig initializes the BatchSubmitter, gathering any resources
// that will be needed during operation.
func NewBatchSubmitterFromCLIConfig(cfg CLIConfig, l log.Logger, m metrics.Metricer) (*BatchSubmitter, error) {
	ctx := context.Background()

	// Connect to L1 and L2 providers. Perform these last since they are the
	// most expensive.
	l1Client, err := opclient.DialEthClientWithTimeout(ctx, cfg.L1EthRpc, opclient.DefaultDialTimeout)
	if err != nil {
		l.Warn("l1 dialEthClientWithTimeout", "err", err)
		return nil, err
	}

	l2Client, err := opclient.DialEthClientWithTimeout(ctx, cfg.L2EthRpc, opclient.DefaultDialTimeout)
	if err != nil {
		l.Warn("l2 dialEthClientWithTimeout", "err", err)
		return nil, err
	}

	//rollupClient, err := opclient.DialRollupClientWithTimeout(ctx, cfg.RollupRpc, opclient.DefaultDialTimeout)
	//if err != nil {
	//	l.Warn("dialRollupClientWithTimeout", "err", err)
	//	return nil, err
	//}

	// SYSCOIN
	syscoinClient, err := opclient.DialSyscoinClientWithTimeout(ctx)
	if err != nil {
		l.Warn("dialSyscoinClientWithTimeout", "err", err)
		return nil, err
	}

	// SYSCOIN
	txManager, err := txmgr.NewSimpleTxManager("batcher", l, m, cfg.TxMgrConfig, syscoinClient)
	if err != nil {
		l.Warn("txmgr.NewConfig", "err", err)
		return nil, err
	}
	//RollupNode:             rollupClient,
	batcherCfg := Config{
		L1Client:               l1Client,
		L2Client:               l2Client,
		PollInterval:           cfg.PollInterval,
		MaxPendingTransactions: cfg.MaxPendingTransactions,
		NetworkTimeout:         cfg.TxMgrConfig.NetworkTimeout,
		TxManager:              txManager,
		Channel: ChannelConfig{
			SeqWindowSize:      288,
			ChannelTimeout:     24,
			MaxChannelDuration: cfg.MaxChannelDuration,
			SubSafetyMargin:    cfg.SubSafetyMargin,
			MaxFrameSize:       cfg.MaxL1TxSize - 1, // subtract 1 byte for version
			CompressorConfig:   cfg.CompressorConfig.Config(),
		},
	}

	// Validate the batcher config
	if err := batcherCfg.Check(); err != nil {
		return nil, err
	}

	return NewBatchSubmitter(ctx, batcherCfg, l, m)
}

// NewBatchSubmitter initializes the BatchSubmitter, gathering any resources
// that will be needed during operation.
func NewBatchSubmitter(ctx context.Context, cfg Config, l log.Logger, m metrics.Metricer) (*BatchSubmitter, error) {
	balance, err := cfg.L1Client.BalanceAt(ctx, cfg.TxManager.From(), nil)
	if err != nil {
		return nil, err
	}

	cfg.log = l
	cfg.log.Info("creating batch submitter", "submitter_addr", cfg.TxManager.From(), "submitter_bal", balance)

	cfg.metr = m

	return &BatchSubmitter{
		Config: cfg,
		txMgr:  cfg.TxManager,
		state:  NewChannelManager(l, m, cfg.Channel),
	}, nil

}

func (l *BatchSubmitter) Start() error {
	l.log.Info("Starting Batch Submitter")

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if l.running {
		return errors.New("batcher is already running")
	}
	l.running = true

	l.shutdownCtx, l.cancelShutdownCtx = context.WithCancel(context.Background())
	l.killCtx, l.cancelKillCtx = context.WithCancel(context.Background())
	l.state.Clear()
	l.lastStoredBlock = eth.BlockID{}

	l.wg.Add(1)
	go l.loop()

	l.log.Info("Batch Submitter started")

	return nil
}

func (l *BatchSubmitter) StopIfRunning(ctx context.Context) {
	_ = l.Stop(ctx)
}

func (l *BatchSubmitter) Stop(ctx context.Context) error {
	l.log.Info("Stopping Batch Submitter")

	l.mutex.Lock()
	defer l.mutex.Unlock()

	if !l.running {
		return errors.New("batcher is not running")
	}
	l.running = false

	// go routine will call cancelKill() if the passed in ctx is ever Done
	cancelKill := l.cancelKillCtx
	wrapped, cancel := context.WithCancel(ctx)
	defer cancel()
	go func() {
		<-wrapped.Done()
		cancelKill()
	}()

	l.cancelShutdownCtx()
	l.wg.Wait()
	l.cancelKillCtx()

	l.log.Info("Batch Submitter stopped")

	return nil
}

// loadBlocksIntoState loads all blocks since the previous stored block
// It does the following:
// 1. Fetch the sync status of the sequencer
// 2. Check if the sync status is valid or if we are all the way up to date
// 3. Check if it needs to initialize state OR it is lagging (todo: lagging just means race condition?)
// 4. Load all new blocks into the local state.
// If there is a reorg, it will reset the last stored block but not clear the internal state so
// the state can be flushed to L1.
func (l *BatchSubmitter) loadBlocksIntoState(ctx context.Context) error {
	//start, end, err := l.calculateL2BlockRangeToStore(ctx)
	//if err != nil {
	//	l.log.Warn("Error calculating L2 block range", "err", err)
	//	return err
	//} else if start.Number >= end.Number {
	//	return errors.New("start number is >= end number")
	//}
	//startHash := common.HexToHash("0xfa3fd601e9ada8afeb4f3379753e6f8cf65c58a2defdb2e3f5cf75482dea095e")
	//endHash := common.HexToHash("0x11ef94bf21f3004b560943f9a2e2a4dbaf46c2ccc6747e77ee6e452ca71ea907")
	//start := eth.BlockID{
	//	Hash:   startHash,
	//	Number: 7944531,
	//}
	//end := eth.BlockID{
	//	Hash:   endHash,
	//	Number: 8111443,
	//}
	// start from genesis
	startHash := common.HexToHash("0xc6b3cc4a21e79da5a3ff4981e73403be68ad93dce5c52401c7172cd8e33af6a0")

	endHash := common.HexToHash("0xa4b6c7b17c1f2098bd547227ca2a53c710431fae61127afc397da7016a7d6935")
	start := eth.BlockID{
		Hash:   startHash,
		Number: 0,
	}
	end := eth.BlockID{
		Hash:   endHash,
		Number: 50000,
	}

	var latestBlock *types.Block
	// Add all blocks to "state"
	for i := start.Number + 1; i < end.Number+1; i++ {
		block, err := l.loadBlockIntoState(ctx, i)
		if errors.Is(err, ErrReorg) {
			l.log.Warn("Found L2 reorg", "block_number", i)
			l.lastStoredBlock = eth.BlockID{}
			return err
		} else if err != nil {
			l.log.Warn("failed to load block into state", "err", err)
			return err
		}
		l.lastStoredBlock = eth.ToBlockID(block)
		latestBlock = block
	}
	l1GenHash := common.HexToHash("0x47ac31005a8c86e4fc6bdbbf4b74ebe578bd8660b3ce44ea2e04f362a5918ca8")
	l2GenHash := common.HexToHash("0xc6b3cc4a21e79da5a3ff4981e73403be68ad93dce5c52401c7172cd8e33af6a0")
	l1Gen := eth.BlockID{
		Hash:   l1GenHash,
		Number: 319398,
	}
	l2Gen := eth.BlockID{
		Hash:   l2GenHash,
		Number: 0,
	}
	overheadBytes := common.Hex2Bytes("0x00000000000000000000000000000000000000000000000000000000000000bc")
	convertedOverheadBytes := ConvertToBytes32(overheadBytes)
	sysConfig := eth.SystemConfig{
		BatcherAddr: common.HexToAddress("0x00d97b2a26cb85252998fe7b4bd4ec2118bf6b6e"),
		Overhead:    convertedOverheadBytes,
	}
	genesis := rollup.Genesis{
		L1:           l1Gen,
		L2:           l2Gen,
		L2Time:       1687365241,
		SystemConfig: sysConfig,
	}
	l2ref, err := derive.L2BlockToBlockRef(latestBlock, &genesis)
	if err != nil {
		l.log.Warn("Invalid L2 block loaded into state", "err", err)
		return err
	}

	l.metr.RecordL2BlocksLoaded(l2ref)
	return nil
}

// loadBlockIntoState fetches & stores a single block into `state`. It returns the block it loaded.
func (l *BatchSubmitter) loadBlockIntoState(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	ctx, cancel := context.WithTimeout(ctx, l.NetworkTimeout)
	defer cancel()
	block, err := l.L2Client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
	if err != nil {
		return nil, fmt.Errorf("getting L2 block: %w", err)
	}

	if err := l.state.AddL2Block(block); err != nil {
		return nil, fmt.Errorf("adding L2 block to state: %w", err)
	}

	l.log.Info("added L2 block to local state", "block", eth.ToBlockID(block), "tx_count", len(block.Transactions()), "time", block.Time())
	return block, nil
}

func (l *BatchSubmitter) loop() {
	defer l.wg.Done()

	ticker := time.NewTicker(l.PollInterval)
	defer ticker.Stop()

	receiptsCh := make(chan txmgr.TxReceipt[txData])
	queue := txmgr.NewQueue[txData](l.killCtx, l.txMgr, l.MaxPendingTransactions)

	for {
		select {
		case <-ticker.C:
			if err := l.loadBlocksIntoState(l.shutdownCtx); errors.Is(err, ErrReorg) {
				err := l.state.Close()
				if err != nil {
					l.log.Error("error closing the channel manager to handle a L2 reorg", "err", err)
				}
				l.publishStateToL1(queue, receiptsCh, true)
				l.state.Clear()
				continue
			}
			l.publishStateToL1(queue, receiptsCh, false)
		case r := <-receiptsCh:
			l.handleReceipt(r)
		case <-l.shutdownCtx.Done():
			err := l.state.Close()
			if err != nil {
				l.log.Error("error closing the channel manager", "err", err)
			}
			l.publishStateToL1(queue, receiptsCh, true)
			return
		}
	}
}

// publishStateToL1 loops through the block data loaded into `state` and
// submits the associated data to the L1 in the form of channel frames.
func (l *BatchSubmitter) publishStateToL1(queue *txmgr.Queue[txData], receiptsCh chan txmgr.TxReceipt[txData], drain bool) {
	txDone := make(chan struct{})
	// send/wait and receipt reading must be on a separate goroutines to avoid deadlocks
	go func() {
		defer func() {
			if drain {
				// if draining, we wait for all transactions to complete
				queue.Wait()
			}
			close(txDone)
		}()
		for {
			err := l.publishTxToL1(l.killCtx, queue, receiptsCh)
			if err != nil {
				if drain && err != io.EOF {
					l.log.Error("error sending tx while draining state", "err", err)
				}
				return
			}
		}
	}()

	for {
		select {
		case r := <-receiptsCh:
			l.handleReceipt(r)
		case <-txDone:
			return
		}
	}
}

// publishTxToL1 submits a single state tx to the L1
func (l *BatchSubmitter) publishTxToL1(ctx context.Context, queue *txmgr.Queue[txData], receiptsCh chan txmgr.TxReceipt[txData]) error {
	// send all available transactions
	l1tip, err := l.l1Tip(ctx)
	if err != nil {
		l.log.Error("Failed to query L1 tip", "error", err)
		return err
	}
	l.recordL1Tip(l1tip)

	// Collect next transaction data
	txdata, err := l.state.TxData(l1tip.ID())
	if err == io.EOF {
		l.log.Trace("no transaction data available")
		return err
	} else if err != nil {
		l.log.Error("unable to get tx data", "err", err)
		return err
	}
	// Function name and parameter types
	_, err = bindings.BatchInboxMetaData.GetAbi()
	if err != nil {
		l.log.Error("Failed to parse contract ABI: %v", err)
		return err
	}

	// SYSCOIN Record TX Status
	//l.log.Info("BLOB", "SEE BLOB", txdata.Bytes())
	vhData := crypto.Keccak256Hash(txdata.Bytes())
	hexString := hex.EncodeToString(txdata.Bytes())
	fmt.Println("Hex string:", hexString)
	filename := fmt.Sprintf("%s.txt", vhData.String())
	// Write the hex string to a file
	err = ioutil.WriteFile(filename, []byte(hexString), 0644)
	if err != nil {
		fmt.Println(hexString)
		fmt.Println("Error writing to file:", err)
	}

	fmt.Println("Data written to file successfully.")
	vhDataHex := fmt.Sprintf("%x", vhData)
	//targetHash := "bf111e9fe9f9749492c156fc87e183916e7e17b589eef8e1e23db75621d8a16b"

	// Compare the computed hash to the target hash
	fmt.Println("Hash is:", vhDataHex)
	//if vhDataHex == targetHash {
	//	fmt.Println("Hash match found. Terminating the program.", "versionhash", vhDataHex)
	//	fileName := "transaction_data.txt"
	//	err := os.WriteFile(fileName, txdata.Bytes(), 0644)
	//	if err != nil {
	//		l.log.Crit("Failed to write to file: %v", err)
	//	}
	//	os.Exit(0) // Terminate the program
	//} else {
	//	fmt.Println("No match found. Hash is:", vhDataHex)
	//}
	l.log.Info("Blob Testing", "versionhash", vhData)
	//if receipt, err := l.sendBlobTransaction(ctx, txdata.Bytes()); err != nil || receipt.Status == types.ReceiptStatusFailed {
	//	l.recordFailedTx(txdata.ID(), err)
	//} else {
	//	l.log.Info("Blob confirmed", "versionhash", receipt.TxHash)
	//	// Create the transaction
	//	// we avoid changing Receipt object and just reuse TxHash for VH
	//	var arrayOfVHs [][32]byte
	//	var array [32]byte
	//	copy(array[:], receipt.TxHash.Bytes())
	//	arrayOfVHs = append(arrayOfVHs, array)
	//
	//	//
	//}
	return nil
}

// sendTransaction creates & submits a transaction to the batch inbox address with the given `data`.
// It currently uses the underlying `txmgr` to handle transaction sending & price management.
// This is a blocking method. It should not be called concurrently.
func (l *BatchSubmitter) sendTransaction(txdata txData, queue *txmgr.Queue[txData], receiptsCh chan txmgr.TxReceipt[txData]) {
	// Do the gas estimation offline. A value of 0 will cause the [txmgr] to estimate the gas limit.
	data := txdata.frame.data
	/*intrinsicGas, err := core.IntrinsicGas(data, nil, false, true, true, false)
	if err != nil {
		l.log.Error("Failed to calculate intrinsic gas", "error", err)
		return
	}*/
	batcherAddr := common.HexToAddress("0x678255ae6b5c4ba0e6206a8e70b59b874f20bc9c")
	candidate := txmgr.TxCandidate{
		To:     &batcherAddr,
		TxData: data,
		// SYSCOIN let L1 estimate gas due to precompile
		GasLimit: 0,
	}
	queue.Send(txdata, candidate, receiptsCh)
}

func (l *BatchSubmitter) handleReceipt(r txmgr.TxReceipt[txData]) {
	// Record TX Status
	if r.Err != nil {
		l.log.Warn("unable to publish tx", "err", r.Err, "data_size", r.ID.Len())
		l.recordFailedTx(r.ID.ID(), r.Err)
	} else {
		l.log.Info("tx successfully published", "tx_hash", r.Receipt.TxHash, "data_size", r.ID.Len())
		l.recordConfirmedTx(r.ID.ID(), r.Receipt)
	}
}

// SYSCOIN
// SendTransaction creates & submits a transaction to the batch inbox address with the given `data`.
// It currently uses the underlying `txmgr` to handle transaction sending & price management.
// This is a blocking method. It should not be called concurrently.
// TODO: where to put concurrent transaction handling logic.
func (l *BatchSubmitter) sendBlobTransaction(ctx context.Context, data []byte) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(ctx, 25*time.Minute)
	defer cancel()
	if receipt, err := l.txMgr.SendBlob(ctx, data); err != nil {
		l.log.Warn("unable to publish blob", "err", err)
		return nil, err
	} else {
		l.log.Info("blob successfully published", "version_hash", receipt.TxHash)
		l.log.Info("blob", "blobdata", data)
		return receipt, nil
	}
}

func (l *BatchSubmitter) recordL1Tip(l1tip eth.L1BlockRef) {
	if l.lastL1Tip == l1tip {
		return
	}
	l.lastL1Tip = l1tip
	l.metr.RecordLatestL1Block(l1tip)
}

func (l *BatchSubmitter) recordFailedTx(id txID, err error) {
	l.log.Warn("Failed to send transaction", "err", err)
	l.state.TxFailed(id)
}

func (l *BatchSubmitter) recordConfirmedTx(id txID, receipt *types.Receipt) {
	l.log.Info("Transaction confirmed", "tx_hash", receipt.TxHash, "status", receipt.Status, "block_hash", receipt.BlockHash, "block_number", receipt.BlockNumber)
	l1block := eth.BlockID{Number: receipt.BlockNumber.Uint64(), Hash: receipt.BlockHash}
	l.state.TxConfirmed(id, l1block)
}

// l1Tip gets the current L1 tip as a L1BlockRef. The passed context is assumed
// to be a lifetime context, so it is internally wrapped with a network timeout.
func (l *BatchSubmitter) l1Tip(ctx context.Context) (eth.L1BlockRef, error) {
	tctx, cancel := context.WithTimeout(ctx, l.NetworkTimeout)
	defer cancel()
	head, err := l.L1Client.HeaderByNumber(tctx, nil)
	if err != nil {
		return eth.L1BlockRef{}, fmt.Errorf("getting latest L1 block: %w", err)
	}
	return eth.InfoToL1BlockRef(eth.HeaderBlockInfo(head)), nil
}

// calculateL2BlockRangeToStore determines the range (start,end] that should be loaded into the local state.
// It also takes care of initializing some local state (i.e. will modify l.lastStoredBlock in certain conditions)
//func (l *BatchSubmitter) calculateL2BlockRangeToStore(ctx context.Context) (eth.BlockID, eth.BlockID, error) {
//	ctx, cancel := context.WithTimeout(ctx, l.NetworkTimeout)
//	defer cancel()
//	syncStatus, err := l.RollupNode.SyncStatus(ctx)
//	// Ensure that we have the sync status
//	if err != nil {
//		return eth.BlockID{}, eth.BlockID{}, fmt.Errorf("failed to get sync status: %w", err)
//	}
//	if syncStatus.HeadL1 == (eth.L1BlockRef{}) {
//		return eth.BlockID{}, eth.BlockID{}, errors.New("empty sync status")
//	}
//
//	// Check last stored to see if it needs to be set on startup OR set if is lagged behind.
//	// It lagging implies that the op-node processed some batches that were submitted prior to the current instance of the batcher being alive.
//	if l.lastStoredBlock == (eth.BlockID{}) {
//		l.log.Info("Starting batch-submitter work at safe-head", "safe", syncStatus.SafeL2)
//		l.lastStoredBlock = syncStatus.SafeL2.ID()
//	} else if l.lastStoredBlock.Number < syncStatus.SafeL2.Number {
//		l.log.Warn("last submitted block lagged behind L2 safe head: batch submission will continue from the safe head now", "last", l.lastStoredBlock, "safe", syncStatus.SafeL2)
//		l.lastStoredBlock = syncStatus.SafeL2.ID()
//	}
//
//	// Check if we should even attempt to load any blocks. TODO: May not need this check
//	if syncStatus.SafeL2.Number >= syncStatus.UnsafeL2.Number {
//		return eth.BlockID{}, eth.BlockID{}, errors.New("L2 safe head ahead of L2 unsafe head")
//	}
//
//	return l.lastStoredBlock, syncStatus.UnsafeL2.ID(), nil
//}

// The following things occur:
// New L2 block (reorg or not)
// L1 transaction is confirmed
//
// What the batcher does:
// Ensure that channels are created & submitted as frames for an L2 range
//
// Error conditions:
// Submitted batch, but it is not valid
// Missed L2 block somehow.
// 73746fa55ae73fb9806269ebd886476415e26aeb75bc5db3874ae41e4c88de05
