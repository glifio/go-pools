package terminate

import (
	"context"
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	filbig "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/builtin/v13/miner"
	lotusapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"
	"github.com/glifio/go-pools/util"
)

func TermPenaltyOnPledge(
	ctx context.Context,
	api *lotusapi.FullNodeStruct,
	ts *types.TipSet,
	filToPledge *big.Int,
	sectorSize uint64,
	ratioVerified float64,
) (cost *big.Int, penalty *big.Int, err error) {
	smoothedPow, err := util.TotalPowerSmoothed(ctx, api, ts)
	if err != nil {
		return nil, nil, err
	}

	smoothedRew, err := util.ThisEpochRewardsSmoothed(ctx, api, ts)
	if err != nil {
		return nil, nil, err
	}

	height := ts.Height()
	activation := 4071894
	expiration := 5619449
	duration := abi.ChainEpoch(expiration - activation)
	powerBaseEpoch := activation
	replacedDayReward := big.NewInt(0)

	dealWeight := big.NewInt(0)

	scaledSectorSize := int64(float64(sectorSize) * ratioVerified)
	verifiedDealWeight := new(big.Int).Mul(
		big.NewInt(scaledSectorSize),
		big.NewInt(int64(duration)),
	)

	pwr := miner.QAPowerForWeight(
		abi.SectorSize(sectorSize),
		duration,
		filbig.NewFromGo(dealWeight),
		filbig.NewFromGo(verifiedDealWeight),
	)

	expectedDayReward := miner.ExpectedRewardForPower(
		util.ConvertSmoothing(smoothedRew),
		util.ConvertSmoothing(smoothedPow),
		pwr,
		builtin.EpochsInDay,
	)

	expectedStoragePledge := miner.ExpectedRewardForPower(
		util.ConvertSmoothing(smoothedRew),
		util.ConvertSmoothing(smoothedPow),
		pwr,
		miner.InitialPledgeProjectionPeriod,
	)

	s := miner.SectorOnChainInfo{
		Activation:            abi.ChainEpoch(activation),
		Expiration:            abi.ChainEpoch(expiration),
		DealWeight:            abi.DealWeight(filbig.NewFromGo(dealWeight)),
		VerifiedDealWeight:    abi.DealWeight(filbig.NewFromGo(verifiedDealWeight)),
		ExpectedDayReward:     expectedDayReward,
		ExpectedStoragePledge: expectedStoragePledge,
		PowerBaseEpoch:        abi.ChainEpoch(powerBaseEpoch),
		ReplacedDayReward:     filbig.NewFromGo(replacedDayReward),
	}
	sectorPower := miner8.QAPowerForSector(abi.SectorSize(sectorSize), util.ConvertSectorType(&s))

	termFee := miner8.PledgePenaltyForTermination(
		s.ExpectedDayReward,
		height-s.PowerBaseEpoch,
		s.ExpectedStoragePledge,
		smoothedPow,
		sectorPower,
		smoothedRew,
		s.ReplacedDayReward,
		s.PowerBaseEpoch-s.Activation,
	)

	// Pledge for full duration
	pledge := miner.ExpectedRewardForPower(
		util.ConvertSmoothing(smoothedRew),
		util.ConvertSmoothing(smoothedPow),
		pwr,
		duration,
	)

	// Calculate number of sectors
	sectors := new(big.Int).Quo(filToPledge, pledge.Int)

	cost = new(big.Int).Mul(pledge.Int, sectors)
	penalty = new(big.Int).Mul(termFee.Int, sectors)

	return cost, penalty, nil
}
