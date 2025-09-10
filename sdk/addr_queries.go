package sdk

import (
	"github.com/ethereum/go-ethereum/common"
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

func (q *fevmQueries) GLF() common.Address {
	return q.glf
}

func (q *fevmQueries) SPPlus() common.Address {
	return q.spPlus
}

func (q *fevmQueries) Governor() common.Address {
	return q.governor
}

func (q *fevmQueries) TokenNFTWrapper() common.Address {
	return q.tokenNFTWrapper
}

func (q *fevmQueries) DelegatedClaimsCampaigns() common.Address {
	return q.delegatedClaimsCampaigns
}
