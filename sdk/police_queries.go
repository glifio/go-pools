package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
)

func (q *fevmQueries) CredentialValidityPeriod(ctx context.Context) (*big.Int, *big.Int, error) {
	epochIssued, err := q.ChainHeight(ctx)
	if err != nil {
		return nil, nil, err
	}
	// credential stays valid for the CredentialEpochsValid duration (30 minutes worth of epochs)
	epochValidUntil := big.NewInt(0).Add(epochIssued, big.NewInt(constants.CredentialEpochsValid))
	return epochIssued, epochValidUntil, nil
}

func (q *fevmQueries) CredentialUsedEpoch(ctx context.Context, vc abigen.VerifiableCredential, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	policeCaller, err := abigen.NewAgentPoliceV2Caller(q.agentPolice, client)
	if err != nil {
		return nil, err
	}

	return policeCaller.CredentialUsed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, vc)
}

func (q *fevmQueries) SectorFaultyTolerance(ctx context.Context) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	policeCaller, err := abigen.NewAgentPoliceV2Caller(q.agentPolice, client)
	if err != nil {
		return nil, err
	}

	return policeCaller.SectorFaultyTolerancePercent(&bind.CallOpts{Context: ctx})
}
