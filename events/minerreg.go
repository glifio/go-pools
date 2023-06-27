package events

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/types"
)

func MinerRegAddMinerEvents(ctx context.Context, sdk types.PoolsSDK, agentsFilter []common.Address, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.MinerRegistryAddMiner, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.MinerRegistryAddMiner{}, err
	}

	minerRegFilterer, err := abigen.NewMinerRegistryFilterer(sdk.Query().MinerRegistry(), ethclient)
	if err != nil {
		return []*abigen.MinerRegistryAddMiner{}, err
	}

	iter, err := minerRegFilterer.FilterAddMiner(getFilterOpts(ctx, startEpoch, endEpoch, sdk.Query().ChainID()), agentsFilter, nil)
	if err != nil {
		return []*abigen.MinerRegistryAddMiner{}, err
	}

	var addMinerEvents []*abigen.MinerRegistryAddMiner

	for iter.Next() {
		addMinerEvents = append(addMinerEvents, iter.Event)
	}

	return addMinerEvents, nil
}

func MinerRegRmMinerEvents(ctx context.Context, sdk types.PoolsSDK, agentsFilter []common.Address, startEpoch *big.Int, endEpoch *big.Int) ([]*abigen.MinerRegistryRemoveMiner, error) {
	ethclient, err := sdk.Extern().ConnectEthClient()
	if err != nil {
		return []*abigen.MinerRegistryRemoveMiner{}, err
	}

	minerRegFilterer, err := abigen.NewMinerRegistryFilterer(sdk.Query().MinerRegistry(), ethclient)
	if err != nil {
		return []*abigen.MinerRegistryRemoveMiner{}, err
	}

	iter, err := minerRegFilterer.FilterRemoveMiner(getFilterOpts(ctx, startEpoch, endEpoch, sdk.Query().ChainID()), agentsFilter, nil)
	if err != nil {
		return []*abigen.MinerRegistryRemoveMiner{}, err
	}

	var events []*abigen.MinerRegistryRemoveMiner

	for iter.Next() {
		events = append(events, iter.Event)
	}

	return events, nil
}
