package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/types"
)

func InfPoolBorrowEvents(ctx context.Context, sdk types.PoolsSDK, agentsFilter []*big.Int, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolBorrow, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolBorrow{}, err
	}

	filterer, err := abigen.NewInfinityPoolFilterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolBorrow{}, err
	}

	iter, err := filterer.FilterBorrow(getFilterOpts(ctx, startEpoch, endEpoch, sdk.Query().ChainID()), agentsFilter)
	if err != nil {
		return []*abigen.InfinityPoolBorrow{}, err
	}

	var events []*abigen.InfinityPoolBorrow

	for iter.Next() {
		events = append(events, iter.Event)
	}

	return events, nil
}

func InfPoolPayEvents(ctx context.Context, sdk types.PoolsSDK, agentsFilter []*big.Int, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolPay, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolPay{}, err
	}

	filterer, err := abigen.NewInfinityPoolFilterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolPay{}, err
	}

	iter, err := filterer.FilterPay(getFilterOpts(ctx, startEpoch, endEpoch, sdk.Query().ChainID()), agentsFilter)
	if err != nil {
		return []*abigen.InfinityPoolPay{}, err
	}

	var events []*abigen.InfinityPoolPay

	for iter.Next() {
		events = append(events, iter.Event)
	}

	return events, nil
}

func InfPoolDepositEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolDeposit, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolDeposit{}, err
	}

	filterer, err := abigen.NewInfinityPoolFilterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolDeposit{}, err
	}

	iter, err := filterer.FilterDeposit(getFilterOpts(ctx, startEpoch, endEpoch, sdk.Query().ChainID()), nil, nil)
	if err != nil {
		return []*abigen.InfinityPoolDeposit{}, err
	}

	var events []*abigen.InfinityPoolDeposit

	for iter.Next() {
		events = append(events, iter.Event)
	}

	return events, nil
}
