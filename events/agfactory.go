package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
)

func AgFactoryCreateAgentEvents(ctx context.Context, sdk types.PoolsSDK, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.AgentFactoryCreateAgent, error) {

	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.AgentFactoryCreateAgent{}, err
	}

	filterer, err := abigen.NewAgentFactoryFilterer(sdk.Query().AgentFactory(), ethclient)
	if err != nil {
		return []*abigen.AgentFactoryCreateAgent{}, err
	}

	var events []*abigen.AgentFactoryCreateAgent
	// to do - can remove hashmap logic when https://github.com/filecoin-project/lotus/issues/10964 gets merged
	var hashmap = make(map[string]bool)

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := filterer.FilterCreateAgent(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), nil, nil, nil)
		if err != nil {
			return []*abigen.AgentFactoryCreateAgent{}, err
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
