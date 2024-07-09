package terminate

import (
	"context"
	"fmt"
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

	/*
		pos, _ := new(big.Int).SetString("8875796959452500272287209423592351924294053446108449795919", 10)
		vel, _ := new(big.Int).SetString("-3571342841146352955296449432600993921460860512773981", 10)
		smoothedPow := smoothing.FilterEstimate{
			PositionEstimate: filbig.NewFromGo(pos),
			VelocityEstimate: filbig.NewFromGo(vel),
		}
	*/
	smoothedPow, err := util.TotalPowerSmoothed(ctx, api, ts)
	if err != nil {
		return nil, nil, err
	}
	// fmt.Printf("Jim smoothedPow @%v: %+v\n", ts.Height(), smoothedPow)

	smoothedRew, err := util.ThisEpochRewardsSmoothed(ctx, api, ts)
	if err != nil {
		return nil, nil, err
	}

	/*
		pos, _ := new(big.Int).SetString("12972374068558641479895275459629742900278373039448326331292", 10)
		vel, _ := new(big.Int).SetString("-6196791619122846434103846045826164594101002962883399", 10)
		smoothedRew := smoothing.FilterEstimate{
			PositionEstimate: filbig.NewFromGo(pos),
			VelocityEstimate: filbig.NewFromGo(vel),
		}
	*/
	// height := abi.ChainEpoch(4072417)
	height := ts.Height()
	activation := 4071894
	expiration := 5619449
	duration := abi.ChainEpoch(expiration - activation)
	powerBaseEpoch := activation
	replacedDayReward := big.NewInt(0)

	dealWeight := big.NewInt(0)

	// FIXME
	verifiedDealWeight, _ := new(big.Int).SetString("53173584910090240", 10)

	// expectedDayReward, _ := new(big.Int).SetString("1445692053274289", 10)

	pwr := miner.QAPowerForWeight(
		abi.SectorSize(sectorSize),
		duration,
		filbig.NewFromGo(dealWeight),
		filbig.NewFromGo(verifiedDealWeight),
	)
	fmt.Printf("Jim pwr: %v\n", pwr)

	expectedDayReward := miner.ExpectedRewardForPower(
		util.ConvertSmoothing(smoothedRew),
		util.ConvertSmoothing(smoothedPow),
		pwr,
		builtin.EpochsInDay,
	)
	fmt.Printf("Jim expectedDayReward: %v\n", expectedDayReward)

	expectedStoragePledge, _ := new(big.Int).SetString("28791341617364739", 10)

	s := miner.SectorOnChainInfo{
		Activation:            abi.ChainEpoch(activation),
		Expiration:            abi.ChainEpoch(expiration),
		DealWeight:            abi.DealWeight(filbig.NewFromGo(dealWeight)),
		VerifiedDealWeight:    abi.DealWeight(filbig.NewFromGo(verifiedDealWeight)),
		ExpectedDayReward:     expectedDayReward,
		ExpectedStoragePledge: filbig.NewFromGo(expectedStoragePledge),
		PowerBaseEpoch:        abi.ChainEpoch(powerBaseEpoch),
		ReplacedDayReward:     filbig.NewFromGo(replacedDayReward),
	}
	sectorPower := miner8.QAPowerForSector(abi.SectorSize(sectorSize), util.ConvertSectorType(&s))

	// the termination penalty calculation
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
	return nil, termFee.Int, nil
}
