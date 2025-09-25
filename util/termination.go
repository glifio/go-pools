package util

import (
	"context"

	"github.com/filecoin-project/go-state-types/abi"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	lotusminer "github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/actors/builtin/power"
	"github.com/filecoin-project/lotus/chain/actors/builtin/reward"
	"github.com/filecoin-project/lotus/chain/types"
	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"
	cbor "github.com/ipfs/go-ipld-cbor"
)

// the SimpleQAPower bool is missing here on the conversion
func ConvertSectorType(sector *lotusminer.SectorOnChainInfo) *miner8.SectorOnChainInfo {
	var expectedDayReward, expectedStoragePledge, replacedDayReward abi.TokenAmount

	if sector.ExpectedDayReward != nil {
		expectedDayReward = *sector.ExpectedDayReward
	}
	if sector.ExpectedStoragePledge != nil {
		expectedStoragePledge = *sector.ExpectedStoragePledge
	}
	if sector.ReplacedDayReward != nil {
		replacedDayReward = *sector.ReplacedDayReward
	}

	return &miner8.SectorOnChainInfo{
		SectorNumber:          sector.SectorNumber,
		SealProof:             sector.SealProof,
		SealedCID:             sector.SealedCID,
		DealIDs:               []abi.DealID{},
		Activation:            sector.Activation,
		Expiration:            sector.Expiration,
		DealWeight:            sector.DealWeight,
		VerifiedDealWeight:    sector.VerifiedDealWeight,
		InitialPledge:         sector.InitialPledge,
		ExpectedDayReward:     expectedDayReward,
		ExpectedStoragePledge: expectedStoragePledge,
		ReplacedDayReward:     replacedDayReward,
		SectorKeyCID:          sector.SectorKeyCID,
	}
}

func TotalPowerSmoothed(ctx context.Context, api *lotusapi.FullNodeStruct, tsk *types.TipSet) (builtin.FilterEstimate, error) {
	pact, err := api.StateGetActor(ctx, power.Address, tsk.Key())
	if err != nil {
		return builtin.FilterEstimate{}, err
	}

	pas, err := power.Load(adt.WrapStore(ctx, cbor.NewCborStore(NewTieredBlockstore(api))), pact)
	if err != nil {
		return builtin.FilterEstimate{}, err
	}

	smoothedPow, err := pas.TotalPowerSmoothed()
	if err != nil {
		return builtin.FilterEstimate{}, err
	}

	return builtin.FilterEstimate{
		PositionEstimate: smoothedPow.PositionEstimate,
		VelocityEstimate: smoothedPow.VelocityEstimate,
	}, nil
}

func ThisEpochRewardsSmoothed(ctx context.Context, api *lotusapi.FullNodeStruct, tsk *types.TipSet) (builtin.FilterEstimate, error) {
	ract, err := api.StateGetActor(ctx, reward.Address, tsk.Key())
	if err != nil {
		return builtin.FilterEstimate{}, err
	}

	ras, err := reward.Load(adt.WrapStore(ctx, cbor.NewCborStore(NewTieredBlockstore(api))), ract)
	if err != nil {
		return builtin.FilterEstimate{}, err
	}

	smoothedRew, err := ras.ThisEpochRewardSmoothed()
	if err != nil {
		return builtin.FilterEstimate{}, err
	}

	return builtin.FilterEstimate{
		PositionEstimate: smoothedRew.PositionEstimate,
		VelocityEstimate: smoothedRew.VelocityEstimate,
	}, nil
}
