package sdk

import (
	"context"
	"math/big"

	"github.com/glifio/go-pools/abigen"
	"github.com/jimpick/go-ethereum/accounts/abi/bind"
)

func (q *fevmQueries) AgentFactoryAgentCount(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agFactoryCaller, err := abigen.NewAgentFactoryCaller(q.agentFactory, client)
	if err != nil {
		return nil, err
	}

	return agFactoryCaller.AgentCount(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
}
