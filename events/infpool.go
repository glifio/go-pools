package events

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
)

func InfPoolBorrowEvents(ctx context.Context, sdk types.PoolsSDK, agentsFilter []*big.Int, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolV2Borrow, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolV2Borrow{}, err
	}

	filterer, err := abigen.NewInfinityPoolV2Filterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolV2Borrow{}, err
	}

	var events []*abigen.InfinityPoolV2Borrow

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterBorrow(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter)
		if err != nil {
			return []*abigen.InfinityPoolV2Borrow{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}
	return events, nil
}

func InfPoolPayEvents(ctx context.Context, sdk types.PoolsSDK, agentsFilter []*big.Int, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolV2Pay, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolV2Pay{}, err
	}

	filterer, err := abigen.NewInfinityPoolV2Filterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolV2Pay{}, err
	}

	var events []*abigen.InfinityPoolV2Pay

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterPay(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter)
		if err != nil {
			return []*abigen.InfinityPoolV2Pay{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}

func InfPoolDepositEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolV2Deposit, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolV2Deposit{}, err
	}

	filterer, err := abigen.NewInfinityPoolV2Filterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolV2Deposit{}, err
	}

	var events []*abigen.InfinityPoolV2Deposit

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterDeposit(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil, nil)
		if err != nil {
			return []*abigen.InfinityPoolV2Deposit{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}

func InfPoolWithdrawEvents(ctx context.Context, sdk types.PoolsSDK, caller []common.Address, receiver []common.Address, owner []common.Address, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolV2Withdraw, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolV2Withdraw{}, err
	}

	filterer, err := abigen.NewInfinityPoolV2Filterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolV2Withdraw{}, err
	}

	var events []*abigen.InfinityPoolV2Withdraw

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterWithdraw(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), caller, receiver, owner)
		if err != nil {
			return []*abigen.InfinityPoolV2Withdraw{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}
	return events, nil
}
