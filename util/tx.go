package util

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/glifio/go-pools/abigen"
)

func WriteTx(
	ctx context.Context,
	pk *ecdsa.PrivateKey,
	chainID *big.Int,
	value *big.Int,
	nonce *big.Int,
	args []interface{},
	writeTx interface{},
	label string,
) (*types.Transaction, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(pk, chainID)
	if err != nil {
		return &types.Transaction{}, err
	}

	auth.Nonce = nonce
	auth.Value = value

	// Use reflection to call the writeTx function with the required arguments
	writeTxValue := reflect.ValueOf(writeTx)
	writeTxArgs := []reflect.Value{reflect.ValueOf(auth)}

	argStrings := make([]string, len(args))
	for i, arg := range args {
		writeTxArgs = append(writeTxArgs, reflect.ValueOf(arg))
		argStrings[i] = StringifyArg(arg)
	}
	result := writeTxValue.Call(writeTxArgs)

	if !result[1].IsNil() {
		return nil, humanReadableRevert(result[1].Interface().(error))
	}

	// Get the transaction and error from the result
	tx := result[0].Interface().(*types.Transaction)

	if tx == nil {
		return nil, fmt.Errorf("Transaction is nil")
	}

	fmt.Println("Transaction:", tx.Hash())

	return tx, err
}

func humanReadableRevert(errMsg error) error {
	errStr := errMsg.Error()

	// Check if "revert reason: " exists in the errMsg
	if !strings.Contains(errStr, "revert reason: ") {
		return fmt.Errorf("error message does not contain a revert reason %s", errStr)
	}

	// Split the errMsg by "revert reason: "
	parts := strings.Split(errStr, "revert reason: ")

	// Check if we have at least two parts
	if len(parts) < 2 {
		return fmt.Errorf("error message is not in the expected format %s", errStr)
	}

	// The error selector should be the next 10 characters after "revert reason: "
	errorSelectorStr := parts[1][:10]

	// Check if the error selector starts with "0x"
	if !strings.HasPrefix(errorSelectorStr, "0x") {
		return fmt.Errorf("error selector does not start with '0x' %s", errStr)
	}

	// Remove the "0x" prefix
	errorSelectorStr = errorSelectorStr[2:]

	// Decode the hex string to a byte slice
	errorSelector, err := hex.DecodeString(errorSelectorStr)
	if err != nil {
		return fmt.Errorf("failed to decode error selector: %v", err)
	}

	matched, err := matchSelector([4]byte(errorSelector))

	if matched {
		return err
	}

	return fmt.Errorf("Error calling contract (could not decode): %s", errStr)
}

func matchSelector(thrownSelector [4]byte) (bool, error) {
	abis, err := gatherABIS()
	// not sure if we should throw the actual error here...
	if err != nil {
		return false, fmt.Errorf("Transaction reverted with error: %s", thrownSelector)
	}

	var humanErr error

	// Loop through all the ABIs
	for _, abi := range abis {
		for name, error := range abi.Errors {
			// Compute the selector for the error
			selector := crypto.Keccak256Hash([]byte(error.Sig))

			if [4]byte(selector[:4]) == thrownSelector {
				humanErr = fmt.Errorf("Transaction reverted with error: %s", name)
			}
		}
	}

	if humanErr == nil {
		return false, fmt.Errorf("Transaction reverted with error: %s", thrownSelector)
	}

	return true, humanErr
}

func gatherABIS() ([]*abi.ABI, error) {
	abis := []*abi.ABI{}

	agentAbi, err := abigen.AgentMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, agentAbi)

	agentFactoryAbi, err := abigen.AgentFactoryMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, agentFactoryAbi)

	agentPoliceAbi, err := abigen.AgentPoliceMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, agentPoliceAbi)

	credParserAbi, err := abigen.CredParserMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, credParserAbi)

	infPoolAbi, err := abigen.InfinityPoolMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, infPoolAbi)

	minerRegistryAbi, err := abigen.MinerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, minerRegistryAbi)

	poolRegistryAbi, err := abigen.PoolRegistryMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, poolRegistryAbi)

	poolTokenAbi, err := abigen.PoolTokenMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, poolTokenAbi)

	rateModuleAbi, err := abigen.RateModuleMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, rateModuleAbi)

	routerAbi, err := abigen.RouterMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, routerAbi)

	wrappedFILAbi, err := abigen.WFILMetaData.GetAbi()
	if err != nil {
		return nil, nil
	}
	abis = append(abis, wrappedFILAbi)

	return abis, nil
}
