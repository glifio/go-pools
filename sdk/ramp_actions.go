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

func (a *fevmActions) RampWithdraw(ctx context.Context, auth *bind.TransactOpts, assets *big.Int, sender common.Address, receiver common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ramp, err := abigen.NewSimpleRampTransactor(a.queries.SimpleRamp(), client)
	if err != nil {
		return nil, err
	}

	tx, err := ramp.WithdrawF(auth, assets, receiver, sender, common.Big0)

	return util.TxPostProcess(tx, err)
}

func (a *fevmActions) RampRedeem(ctx context.Context, auth *bind.TransactOpts, shares *big.Int, sender common.Address, receiver common.Address) (*types.Transaction, error) {
	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	ramp, err := abigen.NewSimpleRampTransactor(a.queries.SimpleRamp(), client)
	if err != nil {
		return nil, err
	}

	tx, err := ramp.RedeemF(auth, shares, receiver, sender, common.Big0)

	return util.TxPostProcess(tx, err)
}
