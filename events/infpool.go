package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
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

	var events []*abigen.InfinityPoolBorrow

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterBorrow(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter)
		if err != nil {
			return []*abigen.InfinityPoolBorrow{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
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

	var events []*abigen.InfinityPoolPay

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterPay(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter)
		if err != nil {
			return []*abigen.InfinityPoolPay{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
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

	var events []*abigen.InfinityPoolDeposit

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterDeposit(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil, nil)
		if err != nil {
			return []*abigen.InfinityPoolDeposit{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}
