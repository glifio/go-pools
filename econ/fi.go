package econ

import (
	"math/big"

	"github.com/glifio/go-pools/constants"
)

func NewBaseFi(
	availableBalance *big.Int,
	lockedRewards *big.Int,
	initialPledge *big.Int,
	feeDebt *big.Int,
	terminationFee *big.Int,
) *BaseFi {
	return &BaseFi{
		AvailableBalance: availableBalance,
		LockedRewards:    lockedRewards,
		InitialPledge:    initialPledge,
		FeeDebt:          feeDebt,
		TerminationFee:   terminationFee,
	}
}

func NewAgentFi(
	agentAvailableBalance *big.Int,
	principal *big.Int,
	minerFis []*BaseFi,
) *AgentFi {
	// loop through all the BaseFi and create 1 consolidated BaseFi
	availableBalance := big.NewInt(0)
	lockedRewards := big.NewInt(0)
	initialPledge := big.NewInt(0)
	feeDebt := big.NewInt(0)
	terminationFee := big.NewInt(0)

	for _, minerFi := range minerFis {
		availableBalance.Add(availableBalance, minerFi.AvailableBalance)
		lockedRewards.Add(lockedRewards, minerFi.LockedRewards)
		initialPledge.Add(initialPledge, minerFi.InitialPledge)
		feeDebt.Add(feeDebt, minerFi.FeeDebt)
		terminationFee.Add(terminationFee, minerFi.TerminationFee)
	}

	// add the agent's available balance
	availableBalance.Add(availableBalance, agentAvailableBalance)

	return &AgentFi{
		BaseFi: BaseFi{
			AvailableBalance: availableBalance,
			LockedRewards:    lockedRewards,
			InitialPledge:    initialPledge,
			FeeDebt:          feeDebt,
			TerminationFee:   terminationFee,
		},
		Principal: principal,
	}
}

func (bfi *BaseFi) Balance() *big.Int {
	bal := new(big.Int).Add(bfi.AvailableBalance, bfi.LockedRewards)
	bal.Add(bal, bfi.InitialPledge)
	bal.Sub(bal, bfi.FeeDebt)
	return bal
}

func (bfi *BaseFi) MaxBorrowAndSeal() *big.Int {
	lv := bfi.LiquidationValue()
	maxBorrowAndSeal := big.NewInt(0).Sub(
		big.NewInt(0).Div(
			big.NewInt(0).Mul(lv, big.NewInt(1e18)),
			big.NewInt(0).Sub(constants.WAD, constants.MAX_BORROW_DTL),
		),
		lv,
	)

	return maxBorrowAndSeal
}

func (bfi *BaseFi) MaxBorrowAndWithdraw() *big.Int {
	lv := bfi.LiquidationValue()
	maxBorrowAndWithdraw := big.NewInt(0).Sub(
		lv,
		big.NewInt(0).Div(
			big.NewInt(0).Mul(
				lv,
				big.NewInt(0).Sub(constants.WAD, constants.MAX_BORROW_DTL),
			),
			big.NewInt(1e18),
		),
	)

	return maxBorrowAndWithdraw
}

func (bfi *BaseFi) LiquidationValue() *big.Int {
	bal := bfi.Balance()
	// make sure we dont have negative liquidation value
	if bal.Cmp(bfi.TerminationFee) <= 0 {
		return big.NewInt(0)
	}

	return new(big.Int).Sub(bfi.Balance(), bfi.TerminationFee)
}

func (afi *AgentFi) BorrowLimit() *big.Int {
	return new(big.Int).Sub(afi.MaxBorrowAndSeal(), afi.Principal)
}

func (afi *AgentFi) WithdrawLimit() *big.Int {
	return big.NewInt(0).Sub(
		afi.LiquidationValue(),
		big.NewInt(0).Div(
			big.NewInt(0).Mul(afi.Principal, big.NewInt(1e18)),
			constants.MAX_BORROW_DTL,
		),
	)
}

func (afi *AgentFi) MarginCall() *big.Int {
	return new(big.Int).Div(big.NewInt(0).Mul(afi.Principal, big.NewInt(1e18)), constants.LIQUIDATION_DTL)
}

func (afi *AgentFi) LeverageRatio() float64 {
	leverageRatio, _ := new(big.Float).Quo(
		new(big.Float).SetInt(afi.Principal),
		new(big.Float).SetInt(afi.LiquidationValue()),
	).Float64()

	return leverageRatio
}
