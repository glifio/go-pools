package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
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

	return chainHeadHeight.Sub(chainHeadHeight, defaultWindow), nil
}

func (q *fevmQueries) MaxConsecutiveFaultEpochs(ctx context.Context) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	policeCaller, err := abigen.NewAgentPoliceCaller(q.agentPolice, client)
	if err != nil {
		return nil, err
	}

	return policeCaller.MaxConsecutiveFaultEpochs(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) CredentialValidityPeriod(ctx context.Context) (*big.Int, *big.Int, error) {
	epochIssued, err := q.ChainHeight(ctx)
	if err != nil {
		return nil, nil, err
	}
	// credential stays valid for the CredentialEpochsValid duration (30 minutes worth of epochs)
	epochValidUntil := big.NewInt(0).Add(epochIssued, big.NewInt(constants.CredentialEpochsValid))
	return epochIssued, epochValidUntil, nil
}

/*
func (q *fevmQueries) CredentialUsed(ctx context.Context, v uint8, r [32]byte, s [32]byte) (bool, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return false, err
	}
	defer client.Close()

	policeCaller, err := abigen.NewAgentPoliceCaller(q.agentPolice, client)
	if err != nil {
		return false, err
	}

	return policeCaller.CredentialUsed(&bind.CallOpts{Context: ctx}, v, r, s)
}
*/

func (q *fevmQueries) SectorFaultyTolerance(ctx context.Context) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	policeCaller, err := abigen.NewAgentPoliceCaller(q.agentPolice, client)
	if err != nil {
		return nil, err
	}

	return policeCaller.SectorFaultyTolerancePercent(&bind.CallOpts{Context: ctx})
}
