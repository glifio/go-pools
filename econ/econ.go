package econ

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	poolstypes "github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/vc"
)

// TODO https://github.com/glif-confidential/ado/issues/9
// note that this function is supposed to implement a post action credential
// in some cases, it's incomplete
// example 1 - pay - if the amount should reduce principal, principal in the credential is not updated - V2 pools does not use principal in the cred anyways
// example 2 - add miner - there are no cred checks on adding miner
func ComputeAgentData(
	ctx context.Context,
	agentAddr common.Address,
	// the change (+/-) in the agent's balance
	agentBalDelta *big.Int,
	principal *big.Int,
	// if this is a remove miner transaction, this will be the address of the miner to remove
	rmMiner address.Address,
	sdk poolstypes.PoolsSDK,
	tsk *types.TipSet,
) (*vc.AgentData, error) {
	// make sure we populate all fields so we dont get encoding errors
	data := &vc.AgentData{
		AgentValue:                  big.NewInt(0),
		CollateralValue:             big.NewInt(0),
		ExpectedDailyFaultPenalties: big.NewInt(0),
		ExpectedDailyRewards:        big.NewInt(0),
		Gcred:                       big.NewInt(0),
		QaPower:                     big.NewInt(0),
		Principal:                   big.NewInt(0),
		FaultySectors:               big.NewInt(0),
		LiveSectors:                 big.NewInt(0),
		GreenScore:                  big.NewInt(0),
	}

	/* ~~~~~ CollateralValue ~~~~~ */

	agentFi, err := EstimateTerminationFeeAgent(ctx, agentAddr, rmMiner, sdk, tsk)
	if err != nil {
		return nil, err
	}

	// negative number means agent balance is being removed (withdraw, pay)
	// check to ensure that subtracting the agentBalDelta from the agent's available balance doesn't result in a negative number
	if (agentBalDelta.Sign() < 0) && (agentFi.AvailableBalance.CmpAbs(agentBalDelta) < 0) {
		return nil, errors.New("insufficient agent balance to process transaction")
	}
	// add any delta to the agent's available balance (namely withdraw, borrow)
	agentFi.AvailableBalance.Add(agentFi.AvailableBalance, agentBalDelta)

	data.CollateralValue = agentFi.LiquidationValue()

	/* ~~~~~ Principal ~~~~~ */

	data.Principal = principal

	/* ~~~~~ SectorInfo ~~~~~ */

	data.LiveSectors = agentFi.LiveSectors

	data.FaultySectors = agentFi.FaultySectors

	return data, nil
}
