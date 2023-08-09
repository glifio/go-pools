package events

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
)

func SimpleRampWithdrawEvents(ctx context.Context, sdk types.PoolsSDK, caller []common.Address, receiver []common.Address, owner []common.Address, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.SimpleRampWithdraw, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.SimpleRampWithdraw{}, err
	}

	filterer, err := abigen.NewSimpleRampFilterer(sdk.Query().SimpleRamp(), ethclient)
	if err != nil {
		return []*abigen.SimpleRampWithdraw{}, err
	}

	var events []*abigen.SimpleRampWithdraw
	var hashmap = make(map[string]bool)

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterWithdraw(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), caller, receiver, owner)
		if err != nil {
			return []*abigen.SimpleRampWithdraw{}, err
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
