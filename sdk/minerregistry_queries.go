package sdk

import (
	"context"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
)

func (q *fevmQueries) MinerRegistryAgentMinersCount(ctx context.Context, agentID *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	mregistryCaller, err := abigen.NewMinerRegistryCaller(q.minerRegistry, client)
	if err != nil {
		return nil, err
	}

	return mregistryCaller.MinersCount(nil, agentID)
}

func (q *fevmQueries) MinerRegistryAgentMinersList(ctx context.Context, agentID *big.Int) ([]address.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	mregistryCaller, err := abigen.NewMinerRegistryCaller(q.minerRegistry, client)
	if err != nil {
		return nil, err
	}

	count, err := mregistryCaller.MinersCount(nil, agentID)
	if err != nil {
		return nil, err
	}

	tasks := make([]util.TaskFunc, count.Int64())
	for i := int64(0); i < count.Int64(); i++ {
		index := big.NewInt(i) // convert loop index to *big.Int
		tasks[i] = func() (interface{}, error) {
			return mregistryCaller.GetMiner(nil, agentID, index)
		}
	}

	results, err := util.Multiread(tasks)
	if err != nil {
		return nil, err
	}

	// Convert the results to []address.Address
	minerAddresses := make([]address.Address, len(results))
	for i, result := range results {
		minerAddr, err := address.NewIDAddress(result.(uint64))
		if err != nil {
			return nil, err
		}
		minerAddresses[i] = minerAddr
	}

	return minerAddresses, nil
}
