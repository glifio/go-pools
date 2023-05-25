package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/util"
)

func (q *fevmQueries) InfPoolGetAccount(ctx context.Context, agentAddr common.Address) (abigen.Account, error) {
	return q.AgentAccount(ctx, agentAddr, constants.INFINITY_POOL_ID)
}

func (q *fevmQueries) InfPoolGetAgentLvl(ctx context.Context, agentID *big.Int) (*big.Int, float64, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, 0, err
	}
	defer client.Close()

	infpool, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, 0, err
	}

	rateModuleAddr, err := infpool.RateModule(nil)
	if err != nil {
		return nil, 0, err
	}

	rateModule, err := abigen.NewRateModuleCaller(rateModuleAddr, client)
	if err != nil {
		return nil, 0, err
	}

	lvl, err := rateModule.AccountLevel(nil, agentID)
	if err != nil {
		return nil, 0, err
	}

	cap, err := rateModule.Levels(nil, lvl)
	if err != nil {
		return nil, 0, err
	}

	capInFIL, _ := util.ToFIL(cap).Float64()

	return lvl, capInFIL, nil
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

func (q *fevmQueries) InfPoolTotalAssets(ctx context.Context) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalAssets(nil)
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolBorrowableLiquidity(ctx context.Context) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalBorrowableAssets(nil)
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}
