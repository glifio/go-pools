package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
	"github.com/jimpick/go-ethereum/common"
)

func PrestakeEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.PreStakeDeposit, error) {

	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.PreStakeDeposit{}, err
	}

	filterer, err := abigen.NewPreStakeFilterer(common.HexToAddress("0x0EC46aD7aA8600118DA4bD64239c3DC364fD0274"), ethclient)
	if err != nil {
		return []*abigen.PreStakeDeposit{}, err
	}

	var events []*abigen.PreStakeDeposit

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterDeposit(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil)
		if err != nil {
			return []*abigen.PreStakeDeposit{}, err
		}

		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}
