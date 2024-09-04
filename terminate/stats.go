package terminate

import (
	"math"
	"math/big"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v13/miner"
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
	PowerBaseEpoch        *big.Int
	MinPowerBaseEpoch     abi.ChainEpoch
	MaxPowerBaseEpoch     abi.ChainEpoch
	ReplacedDayReward     *big.Int
	SectorCount           int64
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
		PowerBaseEpoch:        big.NewInt(0),
		MinPowerBaseEpoch:     abi.ChainEpoch(math.MaxInt64),
		MaxPowerBaseEpoch:     abi.ChainEpoch(0),
		ReplacedDayReward:     big.NewInt(0),
		SectorCount:           0,
	}
}

func (st *SectorStats) AddSector(
	s *miner.SectorOnChainInfo,
	height abi.ChainEpoch,
	termFee *big.Int,
	sectorFeePenalty *big.Int,
) *SectorStats {
	age := height - s.Activation
	return &SectorStats{
		TerminationPenalty:    st.TerminationPenalty.Add(st.TerminationPenalty, termFee),
		SectorFeePenalty:      st.SectorFeePenalty.Add(st.SectorFeePenalty, sectorFeePenalty),
		Activation:            st.Activation.Add(st.Activation, big.NewInt(int64(s.Activation))),
		MinActivation:         min(st.MinActivation, s.Activation),
		MaxActivation:         max(st.MaxActivation, s.Activation),
		Age:                   st.Age.Add(st.Age, big.NewInt(int64(age))),
		MinAge:                min(st.MinAge, age),
		MaxAge:                max(st.MaxAge, age),
		Expiration:            st.Expiration.Add(st.Expiration, big.NewInt(int64(s.Expiration))),
		MinExpiration:         min(st.MinExpiration, s.Expiration),
		MaxExpiration:         max(st.MaxExpiration, s.Expiration),
		DealWeight:            st.DealWeight.Add(st.DealWeight, s.DealWeight.Int),
		VerifiedDealWeight:    st.VerifiedDealWeight.Add(st.VerifiedDealWeight, s.VerifiedDealWeight.Int),
		InitialPledge:         st.InitialPledge.Add(st.InitialPledge, s.InitialPledge.Int),
		ExpectedDayReward:     st.ExpectedDayReward.Add(st.ExpectedDayReward, s.ExpectedDayReward.Int),
		ExpectedStoragePledge: st.ExpectedStoragePledge.Add(st.ExpectedStoragePledge, s.ExpectedStoragePledge.Int),
		PowerBaseEpoch:        st.PowerBaseEpoch.Add(st.PowerBaseEpoch, big.NewInt(int64(s.PowerBaseEpoch))),
		MinPowerBaseEpoch:     min(st.MinPowerBaseEpoch, s.PowerBaseEpoch),
		MaxPowerBaseEpoch:     max(st.MaxPowerBaseEpoch, s.PowerBaseEpoch),
		ReplacedDayReward:     st.ReplacedDayReward.Add(st.ReplacedDayReward, s.ReplacedDayReward.Int),
		SectorCount:           st.SectorCount + 1,
	}
}

func (st *SectorStats) Accumulate(addStats *SectorStats) *SectorStats {
	return &SectorStats{
		TerminationPenalty:    st.TerminationPenalty.Add(st.TerminationPenalty, addStats.TerminationPenalty),
		SectorFeePenalty:      st.SectorFeePenalty.Add(st.SectorFeePenalty, addStats.SectorFeePenalty),
		Activation:            st.Activation.Add(st.Activation, addStats.Activation),
		MinActivation:         min(st.MinActivation, addStats.MinActivation),
		MaxActivation:         max(st.MaxActivation, addStats.MaxActivation),
		Age:                   st.Age.Add(st.Age, addStats.Age),
		MinAge:                min(st.MinAge, addStats.MinAge),
		MaxAge:                max(st.MaxAge, addStats.MaxAge),
		Expiration:            st.Expiration.Add(st.Expiration, addStats.Expiration),
		MinExpiration:         min(st.MinExpiration, addStats.MinExpiration),
		MaxExpiration:         max(st.MaxExpiration, addStats.MaxExpiration),
		DealWeight:            st.DealWeight.Add(st.DealWeight, addStats.DealWeight),
		VerifiedDealWeight:    st.VerifiedDealWeight.Add(st.VerifiedDealWeight, addStats.VerifiedDealWeight),
		InitialPledge:         st.InitialPledge.Add(st.InitialPledge, addStats.InitialPledge),
		ExpectedDayReward:     st.ExpectedDayReward.Add(st.ExpectedDayReward, addStats.ExpectedDayReward),
		ExpectedStoragePledge: st.ExpectedStoragePledge.Add(st.ExpectedStoragePledge, addStats.ExpectedStoragePledge),
		PowerBaseEpoch:        st.PowerBaseEpoch.Add(st.PowerBaseEpoch, addStats.PowerBaseEpoch),
		MinPowerBaseEpoch:     min(st.MinPowerBaseEpoch, addStats.MinPowerBaseEpoch),
		MaxPowerBaseEpoch:     max(st.MaxPowerBaseEpoch, addStats.MaxPowerBaseEpoch),
		ReplacedDayReward:     st.ReplacedDayReward.Add(st.ReplacedDayReward, addStats.ReplacedDayReward),
		SectorCount:           st.SectorCount + addStats.SectorCount,
	}
}

func (st *SectorStats) ScaleUp(sectorsCount int64, sampledSectorsCount int64) *SectorStats {
	scaleBigInt := func(val *big.Int) *big.Int {
		return new(big.Int).Div(
			new(big.Int).Mul(val, big.NewInt(sectorsCount)),
			big.NewInt(sampledSectorsCount),
		)
	}
	return &SectorStats{
		TerminationPenalty:    scaleBigInt(st.TerminationPenalty),
		SectorFeePenalty:      scaleBigInt(st.SectorFeePenalty),
		Activation:            scaleBigInt(st.Activation),
		MinActivation:         st.MinActivation,
		MaxActivation:         st.MaxActivation,
		Age:                   scaleBigInt(st.Age),
		MinAge:                st.MinAge,
		MaxAge:                st.MaxAge,
		Expiration:            scaleBigInt(st.Expiration),
		MinExpiration:         st.MinExpiration,
		MaxExpiration:         st.MaxExpiration,
		DealWeight:            scaleBigInt(st.DealWeight),
		VerifiedDealWeight:    scaleBigInt(st.VerifiedDealWeight),
		InitialPledge:         scaleBigInt(st.InitialPledge),
		ExpectedDayReward:     scaleBigInt(st.ExpectedDayReward),
		ExpectedStoragePledge: scaleBigInt(st.ExpectedStoragePledge),
		PowerBaseEpoch:        scaleBigInt(st.PowerBaseEpoch),
		MinPowerBaseEpoch:     st.MinPowerBaseEpoch,
		MaxPowerBaseEpoch:     st.MaxPowerBaseEpoch,
		ReplacedDayReward:     scaleBigInt(st.ReplacedDayReward),
		SectorCount:           int64(float64(st.SectorCount) * float64(sectorsCount) / float64(sampledSectorsCount)),
	}
}
