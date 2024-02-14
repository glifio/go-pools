package events

import (
	"context"
	"log"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
)

type TxUniqueKey struct {
	Hash  string
	Index uint
}

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
	// to do - can remove hashmap logic when https://github.com/filecoin-project/lotus/issues/10964 gets merged
	var hashmap = make(map[TxUniqueKey]bool)

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
			if _, found := hashmap[TxUniqueKey{iter.Event.Raw.TxHash.Hex(), iter.Event.Raw.Index}]; !found {
				hashmap[TxUniqueKey{iter.Event.Raw.TxHash.Hex(), iter.Event.Raw.Index}] = true
				events = append(events, iter.Event)
			} else {
				log.Println("Duplicate event found", iter.Event.Raw.TxHash.Hex(), iter.Event.Raw.Index)
			}
		}
	}

	return events, nil
}
