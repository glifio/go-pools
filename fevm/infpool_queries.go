package fevm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
)

func (q *fevmQueries) InfPoolGetAccount(ctx context.Context, agentAddr common.Address) (abigen.Account, error) {
	return q.AgentAccount(ctx, agentAddr, constants.INFINITY_POOL_ID)
}

func (q *fevmQueries) InfPoolGetRate(ctx context.Context, cred abigen.VerifiableCredential) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	poolRegCaller, err := abigen.NewPoolRegistryCaller(q.poolRegistry, ethClient)
	if err != nil {
		return nil, err
	}

	poolAddr, err := poolRegCaller.AllPools(&bind.CallOpts{Context: ctx}, constants.INFINITY_POOL_ID)
	if err != nil {
		return nil, err
	}

	infCaller, err := abigen.NewInfinityPoolCaller(poolAddr, ethClient)
	if err != nil {
		return nil, err
	}

	return infCaller.GetRate(&bind.CallOpts{Context: ctx}, cred)
}
