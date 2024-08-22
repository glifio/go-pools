package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
)

func (a *fevmActions) InfPoolDepositFIL(ctx context.Context, auth *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	infpool, err := abigen.NewInfinityPoolV2Transactor(a.queries.InfinityPool(), client)
	if err != nil {
		return nil, err
	}

	auth.Value = amount
	tx, err := infpool.Deposit0(auth, receiver)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) InfPoolWithdraw(ctx context.Context, auth *bind.TransactOpts, assets *big.Int, sender common.Address, receiver common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	pool, err := abigen.NewInfinityPoolV2Transactor(a.queries.InfinityPool(), client)
	if err != nil {
		return nil, err
	}

	return util.TxPostProcess(pool.WithdrawF(auth, assets, receiver, sender, common.Big0))
}

func (a *fevmActions) InfPoolRedeem(ctx context.Context, auth *bind.TransactOpts, shares *big.Int, sender common.Address, receiver common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	pool, err := abigen.NewInfinityPoolV2Transactor(a.queries.InfinityPool(), client)
	if err != nil {
		return nil, err
	}

	return util.TxPostProcess(pool.RedeemF(auth, shares, receiver, sender, common.Big0))
}
