package batcher

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/ethereum-optimism/optimism/op-batcher/flags"
	opservice "github.com/ethereum-optimism/optimism/op-service"
	"github.com/ethereum-optimism/optimism/op-service/cliapp"
	oplog "github.com/ethereum-optimism/optimism/op-service/log"
)

<<<<<<< HEAD
const (
	// SYSCOIN
	appendSequencerBatchMethodName = "appendSequencerBatch"
)

// Main is the entrypoint into the Batch Submitter. This method returns a
// closure that executes the service and blocks until the service exits. The use
// of a closure allows the parameters bound to the top-level main package, e.g.
// GitVersion, to be captured and used once the function is executed.
func Main(version string, cliCtx *cli.Context) error {
	if err := flags.CheckRequired(cliCtx); err != nil {
		return err
	}
	cfg := NewConfig(cliCtx)
	if err := cfg.Check(); err != nil {
		return fmt.Errorf("invalid CLI flags: %w", err)
	}

	l := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
	oplog.SetGlobalLogHandler(l.GetHandler())
	opservice.ValidateEnvVars(flags.EnvVarPrefix, flags.Flags, l)
	procName := "default"
	m := metrics.NewMetrics(procName)
	l.Info("Initializing Batch Submitter")

	batchSubmitter, err := NewBatchSubmitterFromCLIConfig(cfg, l, m)
	if err != nil {
		l.Error("Unable to create Batch Submitter", "error", err)
		return err
	}

	if !cfg.Stopped {
		if err := batchSubmitter.Start(); err != nil {
			l.Error("Unable to start Batch Submitter", "error", err)
			return err
=======
// Main is the entrypoint into the Batch Submitter.
// This method returns a cliapp.LifecycleAction, to create an op-service CLI-lifecycle-managed batch-submitter with.
func Main(version string) cliapp.LifecycleAction {
	return func(cliCtx *cli.Context, closeApp context.CancelCauseFunc) (cliapp.Lifecycle, error) {
		if err := flags.CheckRequired(cliCtx); err != nil {
			return nil, err
>>>>>>> upstream/develop
		}
		cfg := NewConfig(cliCtx)
		if err := cfg.Check(); err != nil {
			return nil, fmt.Errorf("invalid CLI flags: %w", err)
		}

		l := oplog.NewLogger(oplog.AppOut(cliCtx), cfg.LogConfig)
		oplog.SetGlobalLogHandler(l.GetHandler())
		opservice.ValidateEnvVars(flags.EnvVarPrefix, flags.Flags, l)

		l.Info("Initializing Batch Submitter")
		return BatcherServiceFromCLIConfig(cliCtx.Context, version, cfg, l)
	}
}
