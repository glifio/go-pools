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

func (a *fevmActions) InfPoolDepositFIL(ctx context.Context, receiver common.Address, amount *big.Int, pk *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolTransactor, err := abigen.NewInfinityPoolTransactor(a.queries.InfinityPool(), client)
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

	args := []interface{}{receiver}

	return util.WriteTx(ctx, pk, a.queries.ChainID(), amount, nonce, args, poolTransactor.Deposit0, "Deposit FIL")
}
