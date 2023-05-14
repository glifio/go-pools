package sdk

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) IFILTransfer(ctx context.Context, receiver common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iFILTransactor, err := abigen.NewPoolTokenTransactor(a.queries.IFIL(), client)
	if err != nil {
		return nil, err
	}

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{receiver, amount}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, iFILTransactor.Transfer, "Transfer iFIL")
}

func (a *fevmActions) IFILApprove(ctx context.Context, spender common.Address, allowance *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	iFILTransactor, err := abigen.NewPoolTokenTransactor(a.queries.IFIL(), client)
	if err != nil {
		return nil, err
	}

	fromAddr, _, err := util.DeriveAddrFromPk(pk)
	if err != nil {
		return nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return nil, err
	}

	args := []interface{}{spender, allowance}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), common.Big0, nonce, args, iFILTransactor.Approve, "Approve iFIL")
}
