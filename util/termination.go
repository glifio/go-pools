package util

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-state-types/builtin/v14/miner"
	smoothing14 "github.com/filecoin-project/go-state-types/builtin/v14/util/smoothing"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/power"
	"github.com/filecoin-project/lotus/chain/actors/builtin/reward"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/specs-actors/v6/actors/util/smoothing"
	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"
	cbor "github.com/ipfs/go-ipld-cbor"
)

// the SimpleQAPower bool is missing here on the conversion
func ConvertSectorType(sector *miner.SectorOnChainInfo) *miner8.SectorOnChainInfo {
	return &miner8.SectorOnChainInfo{
		SectorNumber:          sector.SectorNumber,
		SealProof:             sector.SealProof,
		SealedCID:             sector.SealedCID,
		DealIDs:               sector.DealIDs,
		Activation:            sector.Activation,
		Expiration:            sector.Expiration,
		DealWeight:            sector.DealWeight,
		VerifiedDealWeight:    sector.VerifiedDealWeight,
		InitialPledge:         sector.InitialPledge,
		ExpectedDayReward:     sector.ExpectedDayReward,
		ExpectedStoragePledge: sector.ExpectedStoragePledge,
		ReplacedDayReward:     sector.ReplacedDayReward,
		SectorKeyCID:          sector.SectorKeyCID,
	}
}

func ConvertSmoothing(fe smoothing.FilterEstimate) smoothing14.FilterEstimate {
	return smoothing14.FilterEstimate{
		PositionEstimate: fe.PositionEstimate,
		VelocityEstimate: fe.VelocityEstimate,
	}
}

func TotalPowerSmoothed(ctx context.Context, api *lotusapi.FullNodeStruct, tsk *types.TipSet) (smoothing.FilterEstimate, error) {
	pact, err := api.StateGetActor(ctx, power.Address, tsk.Key())
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	tbsPow := blockstore.NewTieredBstore(blockstore.NewAPIBlockstore(api), blockstore.NewMemory())

	pas, err := power.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsPow)), pact)
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	smoothedPow, err := pas.TotalPowerSmoothed()
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	return smoothing.NewEstimate(
		smoothedPow.PositionEstimate,
		smoothedPow.VelocityEstimate,
	), nil
}

func ThisEpochRewardsSmoothed(ctx context.Context, api *lotusapi.FullNodeStruct, tsk *types.TipSet) (smoothing.FilterEstimate, error) {
	ract, err := api.StateGetActor(ctx, reward.Address, tsk.Key())
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	tbsRew := blockstore.NewTieredBstore(blockstore.NewAPIBlockstore(api), blockstore.NewMemory())

	ras, err := reward.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsRew)), ract)
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	smoothedRew, err := ras.ThisEpochRewardSmoothed()
	if err != nil {
		return smoothing.FilterEstimate{}, err
	}

	return smoothing.NewEstimate(
		smoothedRew.PositionEstimate,
		smoothedRew.VelocityEstimate,
	), nil
}
