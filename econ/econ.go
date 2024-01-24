package econ

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/mstat"
	poolstypes "github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/vc"
)

// TODO https://github.com/glif-confidential/ado/issues/9
func ComputeAgentData(
	ctx context.Context,
	sdk poolstypes.PoolsSDK,
	agentAvailableBalance *big.Int,
	principal *big.Int,
	aggMinerStats *mstat.MinerStats,
	agentAddr common.Address,
	tsk *types.TipSet,
) (*vc.AgentData, error) {
	// here we just work our way through the AgentData, computing each key
	// TODO: could probably parellize this elegently to speed things up
	data := &vc.AgentData{}

	data.QaPower = aggMinerStats.QualityAdjPower

	data.GreenScore = aggMinerStats.GreenScore

	data.AgentValue = big.NewInt(0).Add(agentAvailableBalance, aggMinerStats.Balance)

	/* ~~~~~ CollateralValue ~~~~~ */

	ats, err := sdk.Query().AgentPreviewTerminationPrecise(ctx, agentAddr, tsk)
	if err != nil {
		return nil, err
	}
	// here we replace the ats.AgentAvailableBal with the agentAvailableBalance passed in this call to compute the post-action liquidation value
	ats.AgentAvailableBal = agentAvailableBalance

	data.CollateralValue = ats.LiquidationValue()

	/* ~~~~~ Principal ~~~~~ */
	data.Principal = principal

	/* ~~~~~ SectorInfo ~~~~~ */
	data.LiveSectors = aggMinerStats.LiveSectors

	// if the LTV (loan to liquidationm value) is greater than the max LTV, we report faulty sectors in order to trigger a liquidation
	// this is a bit of a workaround until the liquidation value buffer is built-in to the contracts directly
	// using wad math
	numerator := new(big.Int).Mul(principal, constants.WAD)

	if (new(big.Int).Div(numerator, data.CollateralValue)).Cmp(constants.MAX_LTV) > 0 {
		// 100% faulty sectors if loan to liquidation value buffer is breached
		data.FaultySectors = aggMinerStats.LiveSectors
	} else {
		data.FaultySectors = aggMinerStats.FaultySectors
	}

	/* ~~~~~ ExpectedDailyFaultPenalties ~~~~~ */

	// COULD REMOVE
	data.ExpectedDailyFaultPenalties = aggMinerStats.PenaltyFaultPerDay

	/* ~~~~~ ExpectedDailyRewards ~~~~~ */

	data.ExpectedDailyRewards = aggMinerStats.ExpectedDailyReward

	/* ~~~~~ GCRED (NOT IN USE) ~~~~~ */
	data.Gcred = big.NewInt(100)

	return data, nil
}
