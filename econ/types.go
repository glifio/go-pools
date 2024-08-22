package econ

import (
	"math/big"
)

type BaseFi struct {
	// total balance of the miner (and in the case of an Agent, includes Agent balance and locked FIL)
	Balance *big.Int `json:"balance"`
	// liquid FIL on the miner (and in the case of an Agent, includes Agent available balance, does not account for any fee debt)
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

// note that in feeDebt, AvailableBalance will be kept at 0, and feeDebt will be a positive number
// lotus actually returns a negative number for availableBalance, but we scrap that and return 0
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
