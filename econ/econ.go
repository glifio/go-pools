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

var ErrInsufficientAgentBalance = errors.New("insufficient agent balance to process transaction")

// TODO https://github.com/glif-confidential/ado/issues/9
func ComputeAgentData(
	ctx context.Context,
	sdk poolstypes.PoolsSDK,
	// the change (+/-) in the agent's balance
	agentBalDelta *big.Int,
	principal *big.Int,
	// if this is an add miner transaction, this will be the address of the miner being added
	addMiner address.Address,
	rmMiner address.Address,
	agentAddr common.Address,
	tsk *types.TipSet,
) (*vc.AgentData, error) {
	data := &vc.AgentData{}

	/* ~~~~~ CollateralValue ~~~~~ */

	agentFi, err := EstimateTerminationFeeAgent(ctx, agentAddr, sdk, tsk)
	if err != nil {
		return nil, err
	}

	// negative number means agent balance is being removed (withdraw, pay)
	// check to ensure that subtracting the agentBalDelta from the agent's available balance doesn't result in a negative number
	if (agentBalDelta.Sign() < 0) && (agentFi.AvailableBalance.CmpAbs(agentBalDelta) < 0) {
		return nil, ErrInsufficientAgentBalance
	}

	data.CollateralValue = agentFi.LiquidationValue()

	/* ~~~~~ Principal ~~~~~ */

	data.Principal = principal

	/* ~~~~~ SectorInfo ~~~~~ */

	data.LiveSectors = agentFi.LiveSectors

	data.FaultySectors = agentFi.FaultySectors

	return data, nil
}
