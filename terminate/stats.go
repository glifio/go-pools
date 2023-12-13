package terminate

import (
	"math"
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
)

type SectorStats struct {
	TerminationPenalty    *big.Int
	SectorFeePenalty      *big.Int
	Activation            *big.Int
	MinActivation         abi.ChainEpoch
	MaxActivation         abi.ChainEpoch
	Age                   *big.Int
	MinAge                abi.ChainEpoch
	MaxAge                abi.ChainEpoch
	Expiration            *big.Int
	MinExpiration         abi.ChainEpoch
	MaxExpiration         abi.ChainEpoch
	DealWeight            *big.Int
	VerifiedDealWeight    *big.Int
	InitialPledge         *big.Int
	ExpectedDayReward     *big.Int
	ExpectedStoragePledge *big.Int
	ReplacedSectorAge     *big.Int
	ReplacedDayReward     *big.Int
}

func NewSectorStats() *SectorStats {
	return &SectorStats{
		TerminationPenalty:    big.NewInt(0),
		SectorFeePenalty:      big.NewInt(0),
		Activation:            big.NewInt(0),
		MinActivation:         abi.ChainEpoch(math.MaxInt64),
		MaxActivation:         abi.ChainEpoch(0),
		Age:                   big.NewInt(0),
		MinAge:                abi.ChainEpoch(math.MaxInt64),
		MaxAge:                abi.ChainEpoch(0),
		Expiration:            big.NewInt(0),
		MinExpiration:         abi.ChainEpoch(math.MaxInt64),
		MaxExpiration:         abi.ChainEpoch(0),
		DealWeight:            big.NewInt(0),
		VerifiedDealWeight:    big.NewInt(0),
		InitialPledge:         big.NewInt(0),
		ExpectedDayReward:     big.NewInt(0),
		ExpectedStoragePledge: big.NewInt(0),
		ReplacedSectorAge:     big.NewInt(0),
		ReplacedDayReward:     big.NewInt(0),
	}
}

func (st *SectorStats) Accumulate(addStats *SectorStats) *SectorStats {
	return &SectorStats{
		TerminationPenalty: st.TerminationPenalty.Add(st.TerminationPenalty,
			addStats.TerminationPenalty),
		SectorFeePenalty:      big.NewInt(0),                 // FIXME
		Activation:            big.NewInt(0),                 // FIXME
		MinActivation:         abi.ChainEpoch(math.MaxInt64), // FIXME
		MaxActivation:         abi.ChainEpoch(0),             // FIXME
		Age:                   big.NewInt(0),                 // FIXME
		MinAge:                abi.ChainEpoch(math.MaxInt64), // FIXME
		MaxAge:                abi.ChainEpoch(0),             // FIXME
		Expiration:            big.NewInt(0),                 // FIXME
		MinExpiration:         abi.ChainEpoch(math.MaxInt64), // FIXME
		MaxExpiration:         abi.ChainEpoch(0),             // FIXME
		DealWeight:            big.NewInt(0),                 // FIXME
		VerifiedDealWeight:    big.NewInt(0),                 // FIXME
		InitialPledge:         big.NewInt(0),                 // FIXME
		ExpectedDayReward:     big.NewInt(0),                 // FIXME
		ExpectedStoragePledge: big.NewInt(0),                 // FIXME
		ReplacedSectorAge:     big.NewInt(0),                 // FIXME
		ReplacedDayReward:     big.NewInt(0),                 // FIXME
	}
}

func (st *SectorStats) ScaleUp(sectorsCount int64, sampledSectorsCount int64) *SectorStats {
	return &SectorStats{
		TerminationPenalty: new(big.Int).Div(
			new(big.Int).Mul(st.TerminationPenalty, big.NewInt(sectorsCount)),
			big.NewInt(sampledSectorsCount),
		),
		SectorFeePenalty:      big.NewInt(0),                 // FIXME
		Activation:            big.NewInt(0),                 // FIXME
		MinActivation:         abi.ChainEpoch(math.MaxInt64), // FIXME
		MaxActivation:         abi.ChainEpoch(0),             // FIXME
		Age:                   big.NewInt(0),                 // FIXME
		MinAge:                abi.ChainEpoch(math.MaxInt64), // FIXME
		MaxAge:                abi.ChainEpoch(0),             // FIXME
		Expiration:            big.NewInt(0),                 // FIXME
		MinExpiration:         abi.ChainEpoch(math.MaxInt64), // FIXME
		MaxExpiration:         abi.ChainEpoch(0),             // FIXME
		DealWeight:            big.NewInt(0),                 // FIXME
		VerifiedDealWeight:    big.NewInt(0),                 // FIXME
		InitialPledge:         big.NewInt(0),                 // FIXME
		ExpectedDayReward:     big.NewInt(0),                 // FIXME
		ExpectedStoragePledge: big.NewInt(0),                 // FIXME
		ReplacedSectorAge:     big.NewInt(0),                 // FIXME
		ReplacedDayReward:     big.NewInt(0),                 // FIXME
	}
}
