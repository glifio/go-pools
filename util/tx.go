package util

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	lotusapi "github.com/filecoin-project/lotus/api"
	filtypes "github.com/filecoin-project/lotus/chain/types"
	"golang.org/x/xerrors"
)

func TxPostProcess(tx *types.Transaction, err error) (*types.Transaction, error) {
	if err != nil {
		return nil, HumanReadableRevert(err)
	}
	if tx == nil {
		return nil, fmt.Errorf("Transaction is nil")
	}

	fmt.Println("Transaction:", tx.Hash())

	return tx, nil
}

func propagateErr(returnTrace filtypes.ReturnTrace) error {
	// chop off the first two bytes to get the FEVM error selector
	matched, err := matchSelector(ToByte4Selector(returnTrace.Return[1:]))
	if matched {
		return xerrors.Errorf("Simulation reverted with error, tx not submitted: %s", err.Error())
	}

	return xerrors.Errorf("subcall error code: %s - return (could not decode) %x", returnTrace.ExitCode, returnTrace.Return[1:])
}

// recurses through the stack trace and pulls out any non 0 exit codes, returns the error
func recursiveSubcallErrorThrown(trace *filtypes.ExecutionTrace, lClient *lotusapi.FullNodeStruct) error {

	// Recursively check the exit code of each subcall
	for _, subcall := range trace.Subcalls {
		err := recursiveSubcallErrorThrown(&subcall, lClient)
		if err != nil {
			return err
		}
	}

	// Check the exit code of the current trace last to make sure we propagate the error
	if trace.MsgRct.ExitCode != 0 {
		return propagateErr(trace.MsgRct)
	}

	// If we've made it this far, none of the exit codes were non-zero
	return nil
}

func SimulateFVMMsg(ctx context.Context, msg *filtypes.Message, lClient *lotusapi.FullNodeStruct) error {
	call, err := lClient.StateCall(ctx, msg, filtypes.EmptyTSK)
	if err != nil {
		return err
	}

	return recursiveSubcallErrorThrown(&call.ExecutionTrace, lClient)
}
