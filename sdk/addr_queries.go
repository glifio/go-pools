package sdk

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
)

func (q *fevmQueries) AgentPolice() common.Address {
	return q.agentPolice
}

func (q *fevmQueries) MinerRegistry() common.Address {
	return q.minerRegistry
}

func (q *fevmQueries) Router() common.Address {
	return q.router
}

func (q *fevmQueries) PoolRegistry() common.Address {
	return q.poolRegistry
}

func (q *fevmQueries) AgentFactory() common.Address {
	return q.agentFactory
}

func (q *fevmQueries) IFIL() common.Address {
	return q.iFIL
}

func (q *fevmQueries) WFIL() common.Address {
	return q.wFIL
}

func (q *fevmQueries) InfinityPool() common.Address {
	return q.infinityPool
}

func (q *fevmQueries) SimpleRamp() common.Address {
	return q.simpleRamp
}

func (q *fevmQueries) RateModule() (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	infpool, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return common.Address{}, err
	}

	return infpool.RateModule(nil)
}
