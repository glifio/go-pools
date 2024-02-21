package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
)

func IFilTransferEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.PoolTokenTransfer, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.PoolTokenTransfer{}, err
	}

	filterer, err := abigen.NewPoolTokenFilterer(sdk.Query().IFIL(), ethclient)
	if err != nil {
		return []*abigen.PoolTokenTransfer{}, err
	}

	var events []*abigen.PoolTokenTransfer

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterTransfer(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil, nil)
		if err != nil {
			return []*abigen.PoolTokenTransfer{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}
