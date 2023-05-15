package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/glifio/go-pools/abigen"
)

func (q *fevmQueries) DefaultEpoch(ctx context.Context) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	policeCaller, err := abigen.NewAgentPoliceCaller(q.agentPolice, client)
	if err != nil {
		return nil, err
	}

	defaultWindow, err := policeCaller.DefaultWindow(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	chainHeadHeight, err := q.ChainHeight(ctx)
	if err != nil {
		return nil, err
	}

	return chainHeadHeight.Add(chainHeadHeight, defaultWindow), nil
}
