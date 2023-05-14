package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/util"
)

func (q *fevmQueries) IFILBalanceOf(ctx context.Context, address common.Address) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolTokenCaller, err := abigen.NewPoolTokenCaller(q.iFIL, client)
	if err != nil {
		return nil, err
	}

	bal, err := poolTokenCaller.BalanceOf(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return nil, err
	}
	return util.ToFIL(bal), nil
}

func (q *fevmQueries) IFILPrice(ctx context.Context) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	infPoolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	price, err := infPoolCaller.ConvertToAssets(nil, constants.WAD)
	if err != nil {
		return nil, err
	}

	// return the price of 1 iFIL in FIL
	return util.ToFIL(price), nil
}
