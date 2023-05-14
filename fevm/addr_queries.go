package fevm

import "github.com/ethereum/go-ethereum/common"

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

func (q *fevmQueries) InfinityPool() common.Address {
	return q.infinityPool
}
