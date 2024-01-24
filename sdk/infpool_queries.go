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

func (q *fevmQueries) InfPoolGetAccount(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (abigen.Account, error) {
	return q.AgentAccount(ctx, agentAddr, constants.INFINITY_POOL_ID, blockNumber)
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

	return rate, nil
}

func (q *fevmQueries) InfPoolTotalAssets(ctx context.Context, blockNumber *big.Int) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolBorrowableLiquidity(ctx context.Context, blockNumber *big.Int) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalBorrowableAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolTotalBorrowed(ctx context.Context, blockNumber *big.Int) (*big.Float, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	assets, err := poolCaller.TotalBorrowed(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return util.ToFIL(assets), nil
}

func (q *fevmQueries) InfPoolExitReserve(ctx context.Context, blockNumber *big.Int) (*big.Int, *big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, nil, err
	}

	minLiquidity, err := poolCaller.GetAbsMinLiquidity(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, nil, err
	}

	liquidAssets, err := poolCaller.GetLiquidAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, nil, err
	}

	if liquidAssets.Cmp(minLiquidity) == -1 {
		return liquidAssets, minLiquidity, nil
	}

	return minLiquidity, minLiquidity, nil
}

func computeMaxDTICap(epochRate *big.Int, edr *big.Int, agentExistingPrincipal *big.Int, maxDTI *big.Int) *big.Int {
	dailyRate := new(big.Int).Mul(epochRate, big.NewInt(constants.EpochsInDay))
	maxBorrowDTICap := new(big.Int).Mul(edr, maxDTI)
	// here we add a WAD in for precision, since dailyRate has an extra wad precision
	maxBorrowDTICap.Mul(maxBorrowDTICap, constants.WAD)
	maxBorrowDTICap.Div(maxBorrowDTICap, dailyRate)
	maxBorrowDTICap.Sub(maxBorrowDTICap, agentExistingPrincipal)
	return maxBorrowDTICap
}

func computeMaxDTECap(agentValue *big.Int, principal *big.Int) *big.Int {
	maxValue := new(big.Int).Mul(agentValue, constants.MAX_DTE)
	// div out wad math precision
	maxValue.Div(maxValue, constants.WAD)

	if maxValue.Cmp(principal) == -1 {
		return big.NewInt(0)
	}

	return new(big.Int).Sub(maxValue, principal)
}

// computeMaxLTVCap returns the max borrowable amount from the agent's liquidation value
// maxBorrow = (-1*((LTV*LVstart)/(LTV * RR) - 1)) - P
// recovery rate % is expressed as wad math whole number - 1e18 is 100%, 9e17 is 90%, etc
func computeMaxLTVCap(liquidationValue *big.Int, principal *big.Int, recoveryRate *big.Int) *big.Int {
	// numerator computation
	num := new(big.Int).Mul(liquidationValue, constants.MAX_LTV)
	num.Mul(num, big.NewInt(-1))

	// denom computation
	denom := new(big.Int).Mul(constants.MAX_LTV, recoveryRate)
	// div out the wad math precision from the recovery rate
	denom.Div(denom, constants.WAD)
	// this is the subtract 1 operation from (LTV * RR)
	denom.Sub(denom, big.NewInt(1e18))

	maxBorrow := new(big.Int).Div(num, denom)

	if maxBorrow.Cmp(principal) < 1 {
		return big.NewInt(0)
	}

	// subtract out the existing principal to arrive at maxBorrow
	maxBorrow.Sub(maxBorrow, principal)
	return maxBorrow
}

func findMinCap(values []*big.Int) *big.Int {
	min := new(big.Int).Set(values[0]) // Copy the first value

	for _, value := range values {
		// If value is smaller than min, replace min
		if value.Cmp(min) == -1 {
			min = value
		}
	}

	return min
}

func MaxBorrowFromAgentData(agentData *vc.AgentData, rate *big.Int, liquidationValue *big.Int, recoveryRate *big.Int) *big.Int {
	caps := []*big.Int{
		computeMaxDTICap(rate, agentData.ExpectedDailyRewards, agentData.Principal, constants.MAX_DTI),
		computeMaxDTECap(agentData.AgentValue, agentData.Principal),
		computeMaxLTVCap(liquidationValue, agentData.Principal, recoveryRate),
	}

	return findMinCap(caps)
}

func (q *fevmQueries) InfPoolAgentMaxBorrow(ctx context.Context, agentAddr common.Address, agentData *vc.AgentData) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentID, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	rateModule, err := q.RateModule()
	if err != nil {
		return nil, err
	}

	rateModuleCaller, err := abigen.NewRateModuleCaller(rateModule, client)
	if err != nil {
		return nil, err
	}

	nullCred, err := vc.NullishVerifiableCredential(*agentData)
	if err != nil {
		return nil, err
	}

	rate, err := q.InfPoolGetRate(ctx, *nullCred)
	if err != nil {
		return nil, err
	}

	agentLvl, err := rateModuleCaller.AccountLevel(nil, agentID)
	if err != nil {
		return nil, err
	}

	agentCap, err := rateModuleCaller.Levels(nil, agentLvl)
	if err != nil {
		return nil, err
	}

	ats, err := q.AgentPreviewTerminationQuick(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	maxAmtFromLvl := new(big.Int).Sub(agentCap, agentData.Principal)

	caps := []*big.Int{
		MaxBorrowFromAgentData(agentData, rate, ats.LiquidationValue(), ats.RecoveryRate()),
		maxAmtFromLvl,
	}

	return findMinCap(caps), nil
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

func (q *fevmQueries) InfPoolMaxEpochsOwedTolerance(ctx context.Context, agentAddr common.Address) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	return poolCaller.MaxEpochsOwedTolerance(nil)
}

func (q *fevmQueries) InfPoolFeesAccrued(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	poolCaller, err := abigen.NewInfinityPoolCaller(q.infinityPool, client)
	if err != nil {
		return nil, err
	}

	return poolCaller.FeesCollected(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
}

func (q *fevmQueries) InfPoolApy(ctx context.Context, blockNumber *big.Int) (*big.Int, error) {
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
		Gcred:                       big.NewInt(100),
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

	baseApr := new(big.Int).Mul(perEpochRate, big.NewInt(constants.EpochsInYear))
	// take out the treasury fee from the APR
	tFeeRate, err := q.TreasuryFeeRate(ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	// apr * (1 - tFeeRate)
	baseApr.Mul(baseApr, new(big.Int).Sub(big.NewInt(1e18), tFeeRate))
	// div out the WAD precision
	baseApr.Div(baseApr, big.NewInt(1e18))

	// use the utilization rate to get the current APY
	borrowed, err := q.InfPoolTotalBorrowed(ctx, nil)
	if err != nil {
		return nil, err
	}
	assets, err := q.InfPoolTotalAssets(ctx, nil)
	if err != nil {
		return nil, err
	}

	borrowedAtto := util.ToAtto(borrowed)
	assetsAtto := util.ToAtto(assets)

	borrowedAtto.Mul(borrowedAtto, big.NewInt(1e18))
	utilizationRate := new(big.Int).Div(borrowedAtto, assetsAtto)

	curApy := new(big.Int).Mul(baseApr, utilizationRate)
	// div out the WAD precision twice to get back to 1e18 = 100% rate
	curApy.Div(curApy, big.NewInt(1e18))
	curApy.Div(curApy, big.NewInt(1e18))

	return curApy, nil
}
