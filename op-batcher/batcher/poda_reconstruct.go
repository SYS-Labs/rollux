package batcher

import (
	"context"
	"errors"
	"github.com/ethereum-optimism/optimism/op-node/eth"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/core/types"
)

func reconstructBlobs() {

}

// loadBlocksIntoState loads all blocks since the previous stored block
// It does the following:
// 1. Fetch the sync status of the sequencer
// 2. Check if the sync status is valid or if we are all the way up to date
// 3. Check if it needs to initialize state OR it is lagging (todo: lagging just means race condition?)
// 4. Load all new blocks into the local state.
// If there is a reorg, it will reset the last stored block but not clear the internal state so
// the state can be flushed to L1.
func (l *BatchSubmitter) loadPodaBlocksIntoState(start eth.BlockID, end eth.BlockID, ctx context.Context) error {
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

	l2ref, err := derive.L2BlockToBlockRef(latestBlock, &l.Rollup.Genesis)
	if err != nil {
		l.log.Warn("Invalid L2 block loaded into state", "err", err)
		return err
	}

	l.metr.RecordL2BlocksLoaded(l2ref)
	return nil
}

//func NewPodaBatchSubmitter(cfg CLIConfig, l log.Logger, m metrics.Metricer) (*BatchSubmitter, error) {
//	ctx := context.Background()
//
//	// Connect to L1 and L2 providers. Perform these last since they are the
//	// most expensive.
//	l1Client, err := opclient.DialEthClientWithTimeout(ctx, "https://rpc.syscoin.org", opclient.DefaultDialTimeout)
//	if err != nil {
//		l.Warn("l1 dialEthClientWithTimeout", "err", err)
//		return nil, err
//	}
//
//	l2Client, err := opclient.DialEthClientWithTimeout(ctx, "https://rpc.rollux.com", opclient.DefaultDialTimeout)
//	if err != nil {
//		l.Warn("l2 dialEthClientWithTimeout", "err", err)
//		return nil, err
//	}
//
//	rollupClient, err := opclient.DialRollupClientWithTimeout(ctx, "temp", opclient.DefaultDialTimeout)
//	if err != nil {
//		l.Warn("dialRollupClientWithTimeout", "err", err)
//		return nil, err
//	}
//
//	// SYSCOIN
//	syscoinClient, err := opclient.DialSyscoinClientWithTimeout(ctx)
//	if err != nil {
//		l.Warn("dialSyscoinClientWithTimeout", "err", err)
//		return nil, err
//	}
//
//	rcfg, err := rollupClient.RollupConfig(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("querying rollup config: %w", err)
//	}
//	txManager, err := txmgr.NewSimpleTxManager("batcher", l, m, cfg.TxMgrConfig, syscoinClient)
//	// SYSCOIN
//	batcherCfg := Config{
//		L1Client:               l1Client,
//		L2Client:               l2Client,
//		RollupNode:             rollupClient,
//		PollInterval:           cfg.PollInterval,
//		MaxPendingTransactions: cfg.MaxPendingTransactions,
//		NetworkTimeout:         cfg.TxMgrConfig.NetworkTimeout,
//		TxManager:              txManager,
//		Rollup:                 rcfg,
//		Channel: ChannelConfig{
//			SeqWindowSize:      rcfg.SeqWindowSize,
//			ChannelTimeout:     rcfg.ChannelTimeout,
//			MaxChannelDuration: cfg.MaxChannelDuration,
//			SubSafetyMargin:    cfg.SubSafetyMargin,
//			MaxFrameSize:       cfg.MaxL1TxSize - 1, // subtract 1 byte for version
//			CompressorConfig:   cfg.CompressorConfig.Config(),
//		},
//	}
//
//	if err != nil {
//		l.Warn("txmgr.NewConfig", "err", err)
//		return nil, err
//	}
//
//	// Validate the batcher config
//	if err = batcherCfg.Check(); err != nil {
//		return nil, err
//	}
//
//	return NewBatchSubmitter(ctx, batcherCfg, l, m)
//}
