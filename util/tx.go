package util

import (
	"context"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/go-address"
	lotusapi "github.com/filecoin-project/lotus/api"
	filtypes "github.com/filecoin-project/lotus/chain/types"
	walletutils "github.com/glifio/go-wallet-utils"
	"github.com/glifio/go-wallet-utils/accounts"
	"golang.org/x/xerrors"
)

func WriteTx(
	ctx context.Context,
	lapi *lotusapi.FullNodeStruct,
	client *ethclient.Client,
	wallet accounts.Wallet,
	account accounts.Account,
	passphrase string,
	proposer address.Address,
	approver address.Address,
	chainID *big.Int,
	value *big.Int,
	nonce *big.Int,
	args []interface{},
	abigenTransactor interface{},
	contractAddress common.Address,
	methodName string,
	label string,
) (common.Hash, *types.Transaction, error) {
	wrappedClient, auth, err := walletutils.NewWalletTransactor(ctx, lapi, client, wallet, &account, passphrase, proposer, approver, chainID)
	if err != nil {
		return common.Hash{}, &types.Transaction{}, err
	}

	auth.Nonce = nonce
	auth.Value = value

	// Get instance using abigenTransactor and wrappedClient
	abigenTransactorValue := reflect.ValueOf(abigenTransactor)
	abigenTransactorArgs := []reflect.Value{
		reflect.ValueOf(contractAddress),
		reflect.ValueOf(wrappedClient),
	}
	abigenTransactorResults := abigenTransactorValue.Call(abigenTransactorArgs)
	instance := abigenTransactorResults[0]
	if !abigenTransactorResults[1].IsNil() {
		err = abigenTransactorResults[1].Interface().(error)
		return common.Hash{}, nil, err
	}

	// Use reflection to call the writeTx function with the required arguments
	writeTxValue := instance.MethodByName(methodName)
	writeTxArgs := []reflect.Value{reflect.ValueOf(auth)}

	argStrings := make([]string, len(args))
	for i, arg := range args {
		writeTxArgs = append(writeTxArgs, reflect.ValueOf(arg))
		argStrings[i] = StringifyArg(arg)
	}
	result := writeTxValue.Call(writeTxArgs)

	if !result[1].IsNil() {
		return common.Hash{}, nil, HumanReadableRevert(result[1].Interface().(error))
	}

	// Get the transaction and error from the result
	tx := result[0].Interface().(*types.Transaction)

	if tx == nil {
		return common.Hash{}, nil, fmt.Errorf("Transaction is nil")
	}

	// fmt.Println("Original Eth Transaction Hash:", tx.Hash())

	txHash := wrappedClient.FilecoinEthHash(tx.Hash())
	fmt.Println("Transaction:", txHash)

	return txHash, tx, err
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
