package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
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

	iter, err := filterer.FilterCreateAgent(getFilterOpts(ctx, startEpoch, endEpoch, sdk.Query().ChainID()), nil, nil, nil)
	if err != nil {
		return []*abigen.AgentFactoryCreateAgent{}, err
	}

	var events []*abigen.AgentFactoryCreateAgent

	for iter.Next() {
		events = append(events, iter.Event)
	}

	return events, nil
}
