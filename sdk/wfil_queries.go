package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
)

func (q *fevmQueries) WFILBalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	wFILCaller, err := abigen.NewWFILCaller(q.wFIL, client)
	if err != nil {
		return nil, err
	}

	bal, err := wFILCaller.BalanceOf(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}
	return util.ToFIL(bal), nil
}

func (q *fevmQueries) WFILAllowance(ctx context.Context, owner common.Address, spender common.Address) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	wFILCaller, err := abigen.NewWFILCaller(q.wFIL, client)
	if err != nil {
		return nil, err
	}

	allowance, err := wFILCaller.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
	if err != nil {
		return nil, err
	}

	// return the price of 1 iFIL in FIL
	return util.ToFIL(allowance), nil
}
