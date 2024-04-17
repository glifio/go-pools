package events

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
	"github.com/jimpick/go-ethereum/common"
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

	var addMinerEvents []*abigen.MinerRegistryAddMiner

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := minerRegFilterer.FilterAddMiner(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter, nil)
		if err != nil {
			return []*abigen.MinerRegistryAddMiner{}, err
		}

		for iter.Next() {
			addMinerEvents = append(addMinerEvents, iter.Event)
		}
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

	var events []*abigen.MinerRegistryRemoveMiner

	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, constants.CHUNKSIZE) {
		end := big.NewInt(0).Add(i, constants.CHUNKSIZE)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}

		iter, err := minerRegFilterer.FilterRemoveMiner(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter, nil)
		if err != nil {
			return []*abigen.MinerRegistryRemoveMiner{}, err
		}
		for iter.Next() {
			events = append(events, iter.Event)
		}
	}

	return events, nil
}
