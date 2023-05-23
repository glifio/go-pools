package util

import (
	"bytes"
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	filbig "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/lotus/chain/actors"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/types/ethtypes"
	multisig0 "github.com/filecoin-project/specs-actors/actors/builtin/multisig"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
)

func MsigProposeFEVMTx(
	ctx context.Context,
	senderAddr address.Address,
	msigAddr address.Address,
	contractAddr common.Address,
	nonce uint64,
	abi *abi.ABI,
	methodName string,
	args []interface{},
) (*ltypes.Message, error) {
	contractAddrFilForm, err := ethtypes.ParseEthAddress(contractAddr.String())
	if err != nil {
		return nil, err
	}

	contractAddrDel, err := contractAddrFilForm.ToFilecoinAddress()
	if err != nil {
		return nil, err
	}

	params, err := abi.Pack(methodName, args...)
	if err != nil {
		return nil, err
	}

	var cborParams bytes.Buffer
	if err := cbg.WriteByteArray(&cborParams, params); err != nil {
		return nil, err
	}

	if msigAddr == address.Undef {
		return nil, errors.New("msig address is undefined")
	}

	if senderAddr == address.Undef {
		return nil, errors.New("Invalid private key")
	}

	enc, actErr := actors.SerializeParams(&multisig0.ProposeParams{
		To:     contractAddrDel,
		Value:  filbig.Zero(),
		Method: builtin.MethodsEVM.InvokeContract,
		Params: cborParams.Bytes(),
	})

	if actErr != nil {
		return nil, xerrors.Errorf("failed to serialize parameters: %w\n", actErr)
	}

	return &ltypes.Message{
		To:     msigAddr,
		From:   senderAddr,
		Value:  filbig.Zero(),
		Nonce:  nonce,
		Method: builtin.MethodsMultisig.Propose,
		Params: enc,
	}, nil
}
