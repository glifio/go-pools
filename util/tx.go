package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
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
		log.Fatal(err)
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
		return nil, result[1].Interface().(error)
	}

	// Get the transaction and error from the result
	tx := result[0].Interface().(*types.Transaction)

	if tx == nil {
		return nil, fmt.Errorf("Transaction is nil")
	}

	fmt.Println("Transaction:", tx.Hash())

	return tx, err
}
