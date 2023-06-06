package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/util"
	"github.com/glifio/go-pools/vc"
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

	infCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, ethClient)
	if err != nil {
		return nil, err
	}

	rate, err := infCaller.GetRate(&bind.CallOpts{Context: ctx}, cred)
	if err != nil {
		return nil, err
	}

	// div out the precision wad from the rate
	rate.Div(rate, big.NewInt(1e18))

	return rate, nil
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

func (q *fevmQueries) InfPoolTotalBorrowed(ctx context.Context) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalBorrowed(nil)
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolRateFromGCRED(ctx context.Context, gcred *big.Int) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	rateModule, err := q.RateModule()
	if err != nil {
		return nil, err
	}

	rateModuleCaller, err := abigen.NewRateModuleCaller(rateModule, client)
	if err != nil {
		return nil, err
	}

	agentDataWithGCRED := vc.AgentData{
		AgentValue:                  common.Big0,
		CollateralValue:             common.Big0,
		ExpectedDailyFaultPenalties: common.Big0,
		ExpectedDailyRewards:        common.Big0,
		Gcred:                       gcred,
		QaPower:                     common.Big0,
		Principal:                   common.Big0,
		FaultySectors:               common.Big0,
		LiveSectors:                 common.Big0,
		GreenScore:                  common.Big0,
	}

	nullishVC, err := vc.NullishVerifiableCredential(agentDataWithGCRED)
	if err != nil {
		return nil, err
	}

	perEpochRate, err := rateModuleCaller.GetRate(&bind.CallOpts{Context: ctx}, *nullishVC)
	if err != nil {
		return nil, err
	}

	// div out the precision wad from the annualized rate
	perEpochRate.Div(perEpochRate, big.NewInt(1e18))

	return util.ToFIL(perEpochRate), nil
}
