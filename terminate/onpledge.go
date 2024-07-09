package terminate

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	filbig "github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/go-state-types/builtin/v13/miner"
	"github.com/filecoin-project/specs-actors/v6/actors/util/smoothing"
	miner8 "github.com/filecoin-project/specs-actors/v8/actors/builtin/miner"
	"github.com/glifio/go-pools/util"
)

func TermPenaltyOnPledge(filToPledge *big.Int, sectorSize uint64, ratioQAP float64) (cost *big.Int, penalty *big.Int, err error) {
	pos, _ := new(big.Int).SetString("8875796959452500272287209423592351924294053446108449795919", 10)
	vel, _ := new(big.Int).SetString("-3571342841146352955296449432600993921460860512773981", 10)
	smoothedPow := smoothing.FilterEstimate{
		PositionEstimate: filbig.NewFromGo(pos),
		VelocityEstimate: filbig.NewFromGo(vel),
	}
	pos, _ = new(big.Int).SetString("12972374068558641479895275459629742900278373039448326331292", 10)
	vel, _ = new(big.Int).SetString("-6196791619122846434103846045826164594101002962883399", 10)
	smoothedRew := smoothing.FilterEstimate{
		PositionEstimate: filbig.NewFromGo(pos),
		VelocityEstimate: filbig.NewFromGo(vel),
	}
	height := abi.ChainEpoch(4072417)
	activation := 4071894
	expiration := 5619449
	powerBaseEpoch := activation
	replacedDayReward := big.NewInt(0)

	dealWeight := big.NewInt(0)

	// FIXME
	verifiedDealWeight, _ := new(big.Int).SetString("53173584910090240", 10)
	expectedDayReward, _ := new(big.Int).SetString("1445692053274289", 10)
	expectedStoragePledge, _ := new(big.Int).SetString("28791341617364739", 10)

	s := miner.SectorOnChainInfo{
		Activation:            abi.ChainEpoch(activation),
		Expiration:            abi.ChainEpoch(expiration),
		DealWeight:            abi.DealWeight(filbig.NewFromGo(dealWeight)),
		VerifiedDealWeight:    abi.DealWeight(filbig.NewFromGo(verifiedDealWeight)),
		ExpectedDayReward:     filbig.NewFromGo(expectedDayReward),
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
