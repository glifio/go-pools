package events

import (
	"context"
	"log"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/types"
)

func InfPoolBorrowEvents(ctx context.Context, sdk types.PoolsSDK, agentsFilter []*big.Int, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.InfinityPoolBorrow, error) {
	chunkSize := big.NewInt(50000)

	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.InfinityPoolBorrow{}, err
	}

	filterer, err := abigen.NewInfinityPoolFilterer(sdk.Query().InfinityPool(), ethclient)
	if err != nil {
		return []*abigen.InfinityPoolBorrow{}, err
	}

	var events []*abigen.InfinityPoolBorrow

	len := big.NewInt(0).Sub(endEpoch, startEpoch)
	log.Println("len", len)
	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, chunkSize) {
		end := big.NewInt(0).Add(i, chunkSize)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}
		log.Println("chunk", i, "->", end)
		iter, err := filterer.FilterBorrow(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter)
		if err != nil {
			return []*abigen.InfinityPoolBorrow{}, err
		}

		for iter.Next() {
			if _, ok := hashmap[iter.Event.Raw.TxHash.Hex()]; !ok {
				hashmap[iter.Event.Raw.TxHash.Hex()] = true
				events = append(events, iter.Event)
			}
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

	chunkSize := big.NewInt(50000)
	var events []*abigen.InfinityPoolPay

	len := big.NewInt(0).Sub(endEpoch, startEpoch)
	log.Println("len", len)
	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, chunkSize) {
		end := big.NewInt(0).Add(i, chunkSize)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}
		log.Println("chunk", i, "->", end)

		iter, err := filterer.FilterPay(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter)
		if err != nil {
			return []*abigen.InfinityPoolPay{}, err
		}

		for iter.Next() {
			if _, ok := hashmap[iter.Event.Raw.TxHash.Hex()]; !ok {
				hashmap[iter.Event.Raw.TxHash.Hex()] = true
				events = append(events, iter.Event)
			}
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

	chunkSize := big.NewInt(50000)
	var events []*abigen.InfinityPoolDeposit

	len := big.NewInt(0).Sub(endEpoch, startEpoch)
	log.Println("len", len)
	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, chunkSize) {
		end := big.NewInt(0).Add(i, chunkSize)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}
		log.Println("chunk", i, "->", end)

		iter, err := filterer.FilterDeposit(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil, nil)
		if err != nil {
			return []*abigen.InfinityPoolDeposit{}, err
		}

		for iter.Next() {
			if _, ok := hashmap[iter.Event.Raw.TxHash.Hex()]; !ok {
				hashmap[iter.Event.Raw.TxHash.Hex()] = true
				events = append(events, iter.Event)
			}
		}
	}

	return events, nil
}
