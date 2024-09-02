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

type Liability struct {
	Principal *big.Int `json:"principal"`
	Interest  *big.Int `json:"interest"`
}

type MinerFi struct {
	BaseFi
}

type AgentFi struct {
	BaseFi
	Liability
	// represents the amount of FIL and WFIL that is held by the Agent's smart contract
	SpendableBalance *big.Int `json:"spendableBalance"`
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

type AgentMarginJSON struct {
	AgentId                uint64  `json:"agentId"`
	Balance                string  `json:"balance"`
	SpendableBalance       string  `json:"spendableBalance"`
	AvailableBalance       string  `json:"availableBalance"`
	LockedRewards          string  `json:"lockedRewards"`
	FeeDebt                string  `json:"feeDebt"`
	InitialPledge          string  `json:"initialPledge"`
	TerminationFee         string  `json:"terminationFee"`
	LiquidationValue       string  `json:"liquidationValue"`
	Margin                 string  `json:"margin"`
	MarginCall             string  `json:"marginCall"`
	Principal              string  `json:"principal"`
	Interest               string  `json:"interest"`
	DTL                    float64 `json:"dtl"`
	LeverageRatio          float64 `json:"leverageRatio"`
	BorrowLimit            string  `json:"borrowLimit"`
	BorrowAndWithdrawLimit string  `json:"borrowAndWithdrawLimit"`
	WithdrawLimit          string  `json:"withdrawLimit"`
	LiveSectors            string  `json:"liveSectors"`
	FaultySectors          string  `json:"faultySectors"`
}
