package econ

import (
	"math/big"

	"github.com/glifio/go-pools/constants"
)

func NewBaseFi(
	balance *big.Int,
	availableBalance *big.Int,
	lockedRewards *big.Int,
	initialPledge *big.Int,
	feeDebt *big.Int,
	terminationFee *big.Int,
	liveSectors *big.Int,
	faultySectors *big.Int,
) *BaseFi {
	return &BaseFi{
		Balance:          balance,
		AvailableBalance: availableBalance,
		LockedRewards:    lockedRewards,
		InitialPledge:    initialPledge,
		FeeDebt:          feeDebt,
		TerminationFee:   terminationFee,

		LiveSectors:   liveSectors,
		FaultySectors: faultySectors,
	}
}

func (termRes *TerminateSectorResult) ToBaseFi() *BaseFi {
	return NewBaseFi(
		termRes.TotalBalance,
		termRes.AvailableBalance,
		termRes.VestingFunds,
		termRes.InitialPledge,
		termRes.FeeDebt,
		termRes.EstimatedTerminationFee,
		big.NewInt(int64(termRes.LiveSectors)),
		big.NewInt(int64(termRes.FaultySectors)),
	)
}

func NewAgentFi(
	agentAvailableBalance *big.Int,
	principal *big.Int,
	minerFis []*BaseFi,
) *AgentFi {
	// loop through all the BaseFi and create 1 consolidated BaseFi
	balance := big.NewInt(0)
	availableBalance := big.NewInt(0)
	lockedRewards := big.NewInt(0)
	initialPledge := big.NewInt(0)
	feeDebt := big.NewInt(0)
	terminationFee := big.NewInt(0)

	liveSectors := big.NewInt(0)
	faultySectors := big.NewInt(0)

	for _, minerFi := range minerFis {
		balance.Add(balance, minerFi.Balance)
		availableBalance.Add(availableBalance, minerFi.AvailableBalance)
		lockedRewards.Add(lockedRewards, minerFi.LockedRewards)
		initialPledge.Add(initialPledge, minerFi.InitialPledge)
		feeDebt.Add(feeDebt, minerFi.FeeDebt)
		terminationFee.Add(terminationFee, minerFi.TerminationFee)

		liveSectors.Add(liveSectors, minerFi.LiveSectors)
		faultySectors.Add(faultySectors, minerFi.FaultySectors)
	}

	// add the agent's available balance to the liquid balance across all miners (this does not account for any fee debt)
	availableBalance.Add(availableBalance, agentAvailableBalance)
	// add the agent's available balance and subtract the fee debt
	balance.Add(balance, agentAvailableBalance)
	balance.Sub(balance, feeDebt)
	// if balance is negative because of fee debt, we just set it to 0 (balances cannot go negative)
	if balance.Sign() < 0 {
		balance = big.NewInt(0)
	}

	return &AgentFi{
		BaseFi: BaseFi{
			Balance:          balance,
			AvailableBalance: availableBalance,
			LockedRewards:    lockedRewards,
			InitialPledge:    initialPledge,
			FeeDebt:          feeDebt,
			TerminationFee:   terminationFee,

			LiveSectors:   liveSectors,
			FaultySectors: faultySectors,
		},
		Principal: principal,
	}
}

func EmptyAgentFi() *AgentFi {
	return &AgentFi{
		BaseFi: BaseFi{
			Balance:          big.NewInt(0),
			AvailableBalance: big.NewInt(0),
			LockedRewards:    big.NewInt(0),
			InitialPledge:    big.NewInt(0),
			FeeDebt:          big.NewInt(0),
			TerminationFee:   big.NewInt(0),
			LiveSectors:      big.NewInt(0),
			FaultySectors:    big.NewInt(0),
		},
		Principal: big.NewInt(0),
	}
}

func (bfi *BaseFi) MaxBorrowAndSeal() *big.Int {
	lv := bfi.LiquidationValue()

	if lv.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

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

	if lv.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

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
	// make sure we dont have negative liquidation value
	if bfi.Balance.Cmp(bfi.TerminationFee) <= 0 {
		return big.NewInt(0)
	}

	return new(big.Int).Sub(bfi.Balance, bfi.TerminationFee)
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
	lv := afi.LiquidationValue()
	if lv.Cmp(big.NewInt(0)) == 0 {
		return 0
	}
	leverageRatio, _ := new(big.Float).Quo(
		new(big.Float).SetInt(afi.Principal),
		new(big.Float).SetInt(afi.LiquidationValue()),
	).Float64()

	return leverageRatio
}
