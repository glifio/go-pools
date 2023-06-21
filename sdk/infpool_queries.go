package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/econ"
	"github.com/glifio/go-pools/types"
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

// InfPoolIsApprovedWithReason returns whether a request has been approved or not, if
// it has been rejected, the reason is supplied. In the case of an error, the reason
// is set to types.RejectionReasonNone.
func (q *fevmQueries) InfPoolIsApprovedWithReason(ctx context.Context, agentAddr common.Address, agentData *vc.AgentData) (bool, types.RejectionReason, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return false, types.RejectionReasonNone, err
	}
	defer client.Close()

	chainHeight, err := q.ChainHeight(ctx)
	if err != nil {
		return false, types.RejectionReasonNone, err
	}

	agentID, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return false, types.RejectionReasonNone, err
	}

	rateModule, err := q.RateModule()
	if err != nil {
		return false, "", err
	}

	rateModuleCaller, err := abigen.NewRateModuleCaller(rateModule, client)
	if err != nil {
		return false, "", err
	}

	account, err := q.InfPoolGetAccount(ctx, agentAddr)
	if err != nil {
		return false, types.RejectionReasonNone, err
	}

	if new(big.Int).Add(account.EpochsPaid, big.NewInt(int64(constants.RepeatBorrowEpochTolerance))).Cmp(chainHeight) == -1 {
		return false, types.RejectionReasonPayUP, nil
	}

	agentLvl, err := rateModuleCaller.AccountLevel(nil, agentID)
	if err != nil {
		return false, "", err
	}

	agentCap, err := rateModuleCaller.Levels(nil, agentLvl)
	if err != nil {
		return false, "", err
	}

	maxAmtFromLvl := new(big.Int).Sub(agentCap, agentData.Principal)

	if account.Principal.Cmp(maxAmtFromLvl) == 1 {
		return false, types.RejectionReasonCap, nil
	}

	if agentData.Principal.Cmp(big.NewInt(0)) == 0 {
		return false, types.RejectionReasonZeroCollateral, nil
	}

	collateralValue := econ.CollateralValue(agentData.AgentValue)
	equityValue := new(big.Int).Sub(agentData.AgentValue, agentData.Principal)

	maxLTV, err := rateModuleCaller.MaxLTV(nil)
	if err != nil {
		return false, "", err
	}

	maxDTE, err := rateModuleCaller.MaxDTE(nil)
	if err != nil {
		return false, "", err
	}

	maxDTI, err := rateModuleCaller.MaxDTI(nil)
	if err != nil {
		return false, "", err
	}

	ltv := new(big.Int).Div(new(big.Int).Mul(agentData.Principal, constants.WAD), collateralValue)
	dte := new(big.Int).Div(new(big.Int).Mul(agentData.Principal, constants.WAD), equityValue)

	if ltv.Cmp(maxLTV) == 1 {
		return false, types.RejectionReasonLTV, nil
	}

	if dte.Cmp(maxDTE) == 1 {
		return false, types.RejectionReasonDTE, nil
	}

	nullCred, err := vc.NullishVerifiableCredential(*agentData)
	if err != nil {
		return false, types.RejectionReasonNone, err
	}

	rate, err := q.InfPoolGetRate(ctx, *nullCred)
	if err != nil {
		return false, types.RejectionReasonNone, err
	}

	dailyFees := rate.Mul(rate, big.NewInt(constants.EpochsInDay))
	dailyFees.Mul(dailyFees, agentData.Principal)
	dailyFees.Div(dailyFees, constants.WAD)

	dti := new(big.Int).Div(dailyFees, agentData.ExpectedDailyRewards)

	if dti.Cmp(maxDTI) == 1 {
		return false, types.RejectionReasonDTI, nil
	}

	return true, types.RejectionReasonNone, nil
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
	return new(big.Int).Sub(agentValue, principal)
}

// P(new) = AgentValue(old) - 2*P(old)
// where the value 2 is based on mstat.TerminationPenaltyRatio being .5
func computeMaxLTVCap(agentValue *big.Int, principal *big.Int) *big.Int {
	oldPrincipalCap := new(big.Int).Mul(principal, big.NewInt(2))
	return new(big.Int).Sub(agentValue, oldPrincipalCap)
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

func MaxBorrowFromAgentData(agentData *vc.AgentData, rate *big.Int) *big.Int {
	caps := []*big.Int{
		// TODO: maxDTI in computeMaxDTI is hardcoded to 1e18, this could be derived from the contracts
		computeMaxDTICap(rate, agentData.ExpectedDailyRewards, agentData.Principal, big.NewInt(25e16)),
		computeMaxDTECap(agentData.AgentValue, agentData.Principal),
		computeMaxLTVCap(agentData.AgentValue, agentData.Principal),
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

	maxAmtFromLvl := new(big.Int).Sub(agentCap, agentData.Principal)

	caps := []*big.Int{
		MaxBorrowFromAgentData(agentData, rate),
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
