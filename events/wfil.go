package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/types"
)

func WFilDepositEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.WFILDeposit, error) {
	chunkSize := big.NewInt(50000)

	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.WFILDeposit{}, err
	}

	filterer, err := abigen.NewWFILFilterer(sdk.Query().WFIL(), ethclient)
	if err != nil {
		return []*abigen.WFILDeposit{}, err
	}

	var events []*abigen.WFILDeposit
	// to do - can remove hashmap logic when https://github.com/filecoin-project/lotus/issues/10964 gets merged
	var hashmap = make(map[string]bool)

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, chunkSize) {
		end := big.NewInt(0).Add(i, chunkSize)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterDeposit(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil)
		if err != nil {
			return []*abigen.WFILDeposit{}, err
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
