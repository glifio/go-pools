package econ

import (
	"math/big"

	"github.com/filecoin-project/go-address"
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
	// MaxDTL represents the maximum debt-to-liquidation ratio for this agent
	MaxDTL *big.Int `json:"maxDTL"`
}

// note that in feeDebt, AvailableBalance will be kept at 0, and feeDebt will be a positive number
// lotus actually returns a negative number for availableBalance, but we scrap that and return 0
type TerminateSectorResult struct {
	TotalBalance     *big.Int
	AvailableBalance *big.Int
	VestingFunds     *big.Int
	InitialPledge    *big.Int
	FeeDebt          *big.Int

	EstimatedTerminationFee *big.Int

	LiveSectors   uint64
	FaultySectors uint64
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
	Tier                   uint8   `json:"tier"`
	LeverageRatio          float64 `json:"leverageRatio"`
	BorrowLimit            string  `json:"borrowLimit"`
	BorrowAndWithdrawLimit string  `json:"borrowAndWithdrawLimit"`
	WithdrawLimit          string  `json:"withdrawLimit"`
	LiveSectors            string  `json:"liveSectors"`
	FaultySectors          string  `json:"faultySectors"`
}

type MinerDetailsJSON struct {
	Miner                  uint64          `json:"miner"`
	AgentId                uint64          `json:"agentId"`
	Actions                uint16          `json:"actions"`
	MinerAddr              address.Address `json:"minerAddr"`
	AvailableBalance       string          `json:"availableBalance"`
	InitialPledge          string          `json:"initialPledge"`
	LockedRewards          string          `json:"lockedRewards"`
	FeeDebt                string          `json:"feeDebt"`
	Balance                string          `json:"balance"`
	EstimatedWeeklyRewards string          `json:"estimatedWeeklyRewards"`
	QAP                    string          `json:"qap"`
	RBP                    string          `json:"rbp"`
	LiveSectors            string          `json:"liveSectors"`
	FaultySectors          string          `json:"faultySectors"`
	RecoveringSectors      string          `json:"recoveringSectors"`
	Ratio                  string          `json:"ratio"`
	TerminationFee         string          `json:"terminationFee"`
	LiquidationValue       string          `json:"liquidationValue"`
}

type PoolMetrics struct {
	Height                    uint64
	Timestamp                 uint64
	PoolTotalAssets           *big.Int
	PoolTotalBorrowed         *big.Int
	PoolTotalBorrowableAssets *big.Int
	PoolExitReserve           *big.Int
	TotalAgentCount           uint64
	TotalMinerCollaterals     *big.Int
	TotalMinersCount          uint64
	TotalValueLocked          *big.Int
	TotalMinersSectors        *big.Int
	TotalMinerQAP             *big.Int
	TotalMinerRBP             *big.Int
	TotalMinerEDR             *big.Int
}

type PoolsMetricsJSON struct {
	Height                    uint64 `json:"height"`
	Timestamp                 uint64 `json:"timestamp"`
	PoolTotalAssets           string `json:"poolTotalAssets"`
	PoolTotalBorrowed         string `json:"poolTotalBorrowed"`
	PoolTotalBorrowableAssets string `json:"poolTotalBorrowableAssets"`
	PoolExitReserve           string `json:"poolExitReserve"`
	TotalAgentCount           uint64 `json:"totalAgentCount"`
	TotalMinerCollaterals     string `json:"totalMinerCollaterals"`
	TotalMinersCount          uint64 `json:"totalMinersCount"`
	TotalValueLocked          string `json:"totalValueLocked"`
	TotalMinersSectors        string `json:"totalMinersSectors"`
	TotalMinerQAP             string `json:"totalMinerQAP"`
	TotalMinerRBP             string `json:"totalMinerRBP"`
	TotalMinerEDR             string `json:"totalMinerEDR"`
}
