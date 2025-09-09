package sdk

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/util"
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

func (q *fevmQueries) CredentialUsed(ctx context.Context, vc abigen.VerifiableCredential, blockNumber *big.Int) (bool, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return false, err
	}
	defer client.Close()

	downgradeMethodID, err := util.MethodStrToBytes(constants.MethodPlusDowngrade)
	if err != nil {
		return false, fmt.Errorf("error converting method string to bytes: %v", err)
	}

	if vc.Action == downgradeMethodID {
		plusCaller, err := abigen.NewPlusCaller(q.agentPolice, client)
		if err != nil {
			return false, err
		}

		usedEpoch, err := plusCaller.CredentialUsed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, vc)
		if err != nil {
			return false, err
		}
		return usedEpoch.Sign() > 0, nil
	} else {
		policeCaller, err := abigen.NewAgentPoliceV2Caller(q.agentPolice, client)
		if err != nil {
			return false, err
		}

		usedEpoch, err := policeCaller.CredentialUsed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, vc)
		if err != nil {
			return false, err
		}
		return usedEpoch.Sign() > 0, nil
	}
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
