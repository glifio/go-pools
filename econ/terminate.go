package econ

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/builtin"
	lotusapi "github.com/filecoin-project/lotus/api"
	minertypes "github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"

	"github.com/glifio/go-pools/constants"
	poolstypes "github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/util"
	"golang.org/x/xerrors"
)

func EstimateTerminationFeeAgent(
	ctx context.Context,
	agentAddr common.Address,
	withoutMiner address.Address,
	psdk poolstypes.PoolsSDK,
	tsk *types.TipSet,
) (*AgentFi, error) {
	lapi, closer, err := psdk.Extern().ConnectLotusClient()
	if err != nil {
		return nil, err
	}
	defer closer()

	height := big.NewInt(int64(tsk.Height()) - 1)
	agentAvail, err := psdk.Query().AgentLiquidAssets(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	miners, err := psdk.Query().AgentMiners(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	// remove any miner that is supplied as the withoutMiner
	if withoutMiner != address.Undef {
		minerCount := len(miners)
		for i, miner := range miners {
			if miner == withoutMiner {
				miners = append(miners[:i], miners[i+1:]...)
				break
			}
		}
		// check that we actually removed a miner
		if minerCount == len(miners) {
			return nil, xerrors.Errorf("miner %s not found in agent's miners", withoutMiner)
		}
	}

	minerCount := int64(len(miners))

	tasks := make([]util.TaskFunc, minerCount)
	for i := int64(0); i < minerCount; i++ {
		minerAddr := miners[i]
		tasks[i] = func() (interface{}, error) {
			return EstimateTerminationFeeMiner(ctx, lapi, minerAddr, tsk)
		}
	}

	results, err := util.Multiread(tasks)
	if err != nil {
		return nil, err
	}

	baseFis := make([]*BaseFi, minerCount)

	for i, res := range results {
		baseFis[i] = res.(*TerminateSectorResult).ToBaseFi()
	}

	principal, err := psdk.Query().AgentPrincipal(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	interest, err := psdk.Query().AgentInterestOwed(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	l := Liability{
		Principal: principal,
		Interest:  interest,
	}

	tier, err := psdk.Query().PlusTierFromAgentAddress(ctx, agentAddr, height)
	if err != nil {
		return nil, err
	}

	maxDTL, exists := constants.TierDTL[tier]
	if !exists {
		// Default to MAX_BORROW_DTL if tier is not recognized
		maxDTL = constants.MAX_BORROW_DTL
	}

	return NewAgentFi(agentAvail, l, baseFis, maxDTL), nil
}

func EstimateTerminationFeeMiner(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*TerminateSectorResult, error) {
	res, err := ComputeMaxTerminationFee(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	if err := failOnBrokenAssumptions(minerAddr, res); err != nil {
		return nil, err
	}

	return res, nil
}

func ComputeMaxTerminationFee(ctx context.Context, api *lotusapi.FullNodeStruct, minerAddr address.Address, ts *types.TipSet) (*TerminateSectorResult, error) {
	result := TerminateSectorResult{
		TotalBalance:     big.NewInt(0),
		AvailableBalance: big.NewInt(0),
		VestingFunds:     big.NewInt(0),
		InitialPledge:    big.NewInt(0),
		FeeDebt:          big.NewInt(0),

		EstimatedTerminationFee: big.NewInt(0),

		LiveSectors:   0,
		FaultySectors: 0,
	}

	actor, mstate, err := util.LoadMinerActor(ctx, api, minerAddr, ts)
	if err != nil {
		return nil, err
	}

	lf, err := mstate.LockedFunds()
	if err != nil {
		return nil, err
	}
	result.TotalBalance = actor.Balance.Int
	result.VestingFunds = lf.VestingFunds.Int
	result.InitialPledge = lf.InitialPledgeRequirement.Int

	avail, err := mstate.AvailableBalance(actor.Balance)
	if err != nil {
		return nil, err
	}

	feeDebt, err := mstate.FeeDebt()
	if err != nil {
		return nil, err
	}
	result.FeeDebt = feeDebt.Int

	// sanity check that the available is a positive number
	if feeDebt.Int.Sign() > 0 && avail.Int.Sign() > 0 {
		// this should never happen
		return nil, xerrors.Errorf("miner: %s: fee debt and available balance are both positive", minerAddr)
	} else if feeDebt.Int.Sign() > 0 {
		result.AvailableBalance = big.NewInt(0)
	} else {
		// otherwise no fee debt, go on normally
		result.AvailableBalance = avail.Int
	}

	epochRewardSmoothed, err := util.ThisEpochRewardsSmoothed(ctx, api, ts)
	if err != nil {
		return nil, err
	}

	epochQaPowerSmoothed, err := util.TotalPowerSmoothed(ctx, api, ts)
	if err != nil {
		return nil, err
	}

	nv, err := api.StateNetworkVersion(ctx, ts.Key())
	if err != nil {
		return nil, err
	}

	p, err := api.StateMinerPower(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}

	sectorCount, err := api.StateMinerSectorCount(context.Background(), minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}
	result.LiveSectors = sectorCount.Live
	result.FaultySectors = sectorCount.Faulty

	faultFee, err := minertypes.PledgePenaltyForContinuedFault(
		nv,
		epochRewardSmoothed,
		epochQaPowerSmoothed,
		p.MinerPower.QualityAdjPower,
	)
	if err != nil {
		return nil, err
	}

	penalty, err := minertypes.PledgePenaltyForTermination(
		nv,
		lf.InitialPledgeRequirement,
		// something bigger than 140 days of epochs since we aren't computing on a per sector basis
		miner.TerminationLifetimeCap*builtin.EpochsInDay,
		faultFee,
	)
	if err != nil {
		return nil, err
	}

	result.EstimatedTerminationFee = penalty.Int

	return &result, nil
}

// Assumptions:
// 1. there cannot be negative fee debt
// 2. there cannot be positive fee debt with a positive balance - BROKEN ASSUMPTION
// 3. there cannot be positive fee debt with a positive available balance
// 4. there cannot be a positive fee debt with a positive initial pledge - BROKEN ASSUMPTION
// 5. there cannot be a positive fee debt with a positive vesting funds - BROKEN ASSUMPTION
// 6. there cannot be a negative termination fee
// 7. there cannot be a negative initial pledge
// 8. with no fee debt, available balance should always be positive (or 0)
// 9. with no fee debt, initial pledge should always be positive (or 0)
// 10. with no fee debt, balance should always be positive (or 0)
func failOnBrokenAssumptions(minerAddr address.Address, result *TerminateSectorResult) error {
	// 6. there cannot be a negative termination fee
	if result.EstimatedTerminationFee.Cmp(big.NewInt(0)) < 0 {
		return xerrors.Errorf("miner %s: negative termination fee: %v", minerAddr, result.EstimatedTerminationFee)
	}
	// 7. there cannot be a negative initial pledge
	if result.InitialPledge.Cmp(big.NewInt(0)) < 0 {
		return xerrors.Errorf("miner %s: negative initial pledge: %v", minerAddr, result.InitialPledge)
	}

	// fee debt exists
	if result.FeeDebt.Cmp(big.NewInt(0)) > 0 {
		// 3. there cannot be positive fee debt with a positive available balance
		if result.AvailableBalance.Cmp(big.NewInt(0)) > 0 {
			return xerrors.Errorf("miner %s: positive fee debt with positive available balance: %v", minerAddr, result.FeeDebt)
		}
	} else if result.FeeDebt.Cmp(big.NewInt(0)) == 0 {
		// no fee debt
		// 8. with no fee debt, available balance should always be positive (or 0)
		if result.AvailableBalance.Cmp(big.NewInt(0)) < 0 {
			return xerrors.Errorf("miner %s: no fee debt, but negative available balance: %v", minerAddr, result.AvailableBalance)
		}
		// 9. with no fee debt, initial pledge should always be positive (or 0)
		if result.InitialPledge.Cmp(big.NewInt(0)) < 0 {
			return xerrors.Errorf("miner %s: no fee debt, but negative initial pledge: %v", minerAddr, result.InitialPledge)
		}
		// 10. with no fee debt, balance should always be positive (or 0)
		if result.TotalBalance.Cmp(big.NewInt(0)) < 0 {
			return xerrors.Errorf("miner %s: no fee debt, but negative balance: %v", minerAddr, result.TotalBalance)
		}
	} else {
		// 1. fee debt is negative - error
		return xerrors.Errorf("miner %s: negative fee debt: %v", minerAddr, result.FeeDebt)
	}

	return nil
}
