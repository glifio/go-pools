package econ

import (
	"math/big"
)

type BaseFi struct {
	AvailableBalance *big.Int `json:"availableBalance"`
	LockedRewards    *big.Int `json:"lockedRewards"`
	InitialPledge    *big.Int `json:"initialPledge"`
	FeeDebt          *big.Int `json:"feeDebt"`
	TerminationFee   *big.Int `json:"terminationFee"`

	LiveSectors   *big.Int `json:"liveSectors"`
	FaultySectors *big.Int `json:"faultySectors"`
}

type AgentFi struct {
	BaseFi

	Principal *big.Int `json:"principal"`
}

type TerminateSectorResult struct {
	TotalBalance     *big.Int
	AvailableBalance *big.Int
	VestingFunds     *big.Int
	InitialPledge    *big.Int
	FeeDebt          *big.Int

	EstimatedInitialPledge    *big.Int
	InitialPledgeFromSample   *big.Int
	AvgInitialPledgePerSector *big.Int

	EstimatedTerminationFee    *big.Int
	TerminationFeeFromSample   *big.Int
	AvgTerminationFeePerPledge *big.Int

	SampledSectors uint64
	LiveSectors    uint64
	FaultySectors  uint64
}
