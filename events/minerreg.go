package events

import (
	"context"
	"log"
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

	var addMinerEvents []*abigen.MinerRegistryAddMiner
	var hashmap = make(map[string]bool)

	chunkSize := big.NewInt(50000)

	len := big.NewInt(0).Sub(endEpoch, startEpoch)
	log.Println("len", len)
	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, chunkSize) {
		end := big.NewInt(0).Add(i, chunkSize)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}
		log.Println("chunk", i, "->", end)

		iter, err := minerRegFilterer.FilterAddMiner(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter, nil)
		if err != nil {
			return []*abigen.MinerRegistryAddMiner{}, err
		}

		for iter.Next() {
			if _, ok := hashmap[iter.Event.Raw.TxHash.Hex()]; !ok {
				hashmap[iter.Event.Raw.TxHash.Hex()] = true
				addMinerEvents = append(addMinerEvents, iter.Event)
			}
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

	chunkSize := big.NewInt(50000)

	len := big.NewInt(0).Sub(endEpoch, startEpoch)
	log.Println("len", len)
	for i := startEpoch; i.Cmp(endEpoch) == -1; i.Add(i, chunkSize) {
		end := big.NewInt(0).Add(i, chunkSize)
		if end.Cmp(endEpoch) == 1 {
			end = endEpoch
		}
		log.Println("chunk", i, "->", end)

		iter, err := minerRegFilterer.FilterRemoveMiner(getFilterOpts(ctx, i, end, sdk.Query().ChainID()), agentsFilter, nil)
		if err != nil {
			return []*abigen.MinerRegistryRemoveMiner{}, err
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
