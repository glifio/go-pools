package econ

import (
	"fmt"
	"math/big"

	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/util"
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

func NewMinerFi(
	balance *big.Int,
	availableBalance *big.Int,
	lockedRewards *big.Int,
	initialPledge *big.Int,
	feeDebt *big.Int,
	terminationFee *big.Int,
	liveSectors *big.Int,
	faultySectors *big.Int,
) *MinerFi {
	return &MinerFi{
		BaseFi: *NewBaseFi(
			balance,
			availableBalance,
			lockedRewards,
			initialPledge,
			feeDebt,
			terminationFee,
			liveSectors,
			faultySectors,
		),
	}
}

func NewAgentFi(
	agentAvailableBalance *big.Int,
	liability Liability,
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
		Liability:        liability,
		SpendableBalance: agentAvailableBalance,
	}
}

func EmptyBaseFi() *BaseFi {
	return &BaseFi{
		Balance:          big.NewInt(0),
		AvailableBalance: big.NewInt(0),
		LockedRewards:    big.NewInt(0),
		InitialPledge:    big.NewInt(0),
		FeeDebt:          big.NewInt(0),
		TerminationFee:   big.NewInt(0),
		LiveSectors:      big.NewInt(0),
		FaultySectors:    big.NewInt(0),
	}
}

func EmptyMinerFi() *MinerFi {
	return &MinerFi{
		BaseFi: *EmptyBaseFi(),
	}
}

func EmptyAgentFi() *AgentFi {
	return &AgentFi{
		BaseFi: *EmptyBaseFi(),
		Liability: Liability{
			Principal: big.NewInt(0),
			Interest:  big.NewInt(0),
		},
		SpendableBalance: big.NewInt(0),
	}
}

func (mfi *MinerFi) MaxBorrowAndSeal() *big.Int {
	lv := mfi.LiquidationValue()

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

func (mfi *MinerFi) MaxBorrowAndWithdraw() *big.Int {
	lv := mfi.LiquidationValue()

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

func (bfi *BaseFi) RecoveryRate() *big.Float {
	if bfi.Balance.Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}

	return computePerc(bfi.LiquidationValue(), bfi.Balance)
}

// MaxBorrowAndSeal = margin / (1 - max borrow DTL) - margin
func (afi *AgentFi) MaxBorrowAndSeal() *big.Int {
	margin := afi.Margin()

	if margin.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	maxBorrowAndSeal := big.NewInt(0).Sub(
		big.NewInt(0).Div(
			big.NewInt(0).Mul(margin, big.NewInt(1e18)),
			big.NewInt(0).Sub(constants.WAD, constants.MAX_BORROW_DTL),
		),
		margin,
	)

	return maxBorrowAndSeal
}

// MaxBorrowAndWithdraw = margin - lv * (1 - max borrow DTL)
func (afi *AgentFi) MaxBorrowAndWithdraw() *big.Int {
	margin := afi.Margin()

	if margin.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	maxBorrowAndWithdraw := big.NewInt(0).Sub(
		margin,
		big.NewInt(0).Div(
			big.NewInt(0).Mul(
				afi.LiquidationValue(),
				big.NewInt(0).Sub(constants.WAD, constants.MAX_BORROW_DTL),
			),
			big.NewInt(1e18), // div wad
		),
	)

	return maxBorrowAndWithdraw
}

func (afi *AgentFi) BorrowLimit() *big.Int {
	return new(big.Int).Sub(afi.MaxBorrowAndSeal(), afi.Debt())
}

func (afi *AgentFi) WithdrawLimit() *big.Int {
	withdrawLimit := big.NewInt(0).Sub(
		afi.LiquidationValue(),
		big.NewInt(0).Div(
			big.NewInt(0).Mul(afi.Debt(), big.NewInt(1e18)),
			constants.MAX_BORROW_DTL,
		),
	)

	// make sure the agent has enough available balance to withdraw the withdraw limit
	if afi.AvailableBalance.Cmp(withdrawLimit) < 0 {
		return afi.AvailableBalance
	}
	return withdrawLimit
}

// margin = liquidation value - principal
func (afi *AgentFi) Margin() *big.Int {
	return big.NewInt(0).Sub(afi.LiquidationValue(), afi.Debt())
}

func (afi *AgentFi) MarginCall() *big.Int {
	return new(big.Int).Div(big.NewInt(0).Mul(afi.Debt(), big.NewInt(1e18)), constants.LIQUIDATION_DTL)
}

// leverage ratio = liquidation value / margin
func (afi *AgentFi) LeverageRatio() *big.Float {
	// if debt is zero then leverage ratio is 1
	if afi.Debt().Cmp(big.NewInt(0)) == 0 {
		return new(big.Float).SetInt(big.NewInt(1))
	}
	// if margin is zero and debt is not zero then leverage ratio is infinite
	if afi.Margin().Cmp(big.NewInt(0)) == 0 {
		// set float to infinite
		return new(big.Float).SetInf(true)
	}

	return computePerc(afi.LiquidationValue(), afi.Margin())
}

func (afi *AgentFi) DTL() *big.Float {
	fmt.Printf("Jim agentFi: %+v\n", afi)
	// if debt is zero then DTL is 0
	if afi.Debt().Cmp(big.NewInt(0)) == 0 {
		return big.NewFloat(0)
	}
	lv := afi.LiquidationValue()
	// if liquidation value is zero and debt is non zero then DTL is infinite
	if lv.Cmp(big.NewInt(0)) == 0 {
		// return positive infinite float
		return new(big.Float).SetInf(true)
	}
	fmt.Printf("Jim debt: %+v\n", afi.Debt())
	fmt.Printf("Jim liquidationValue: %+v\n", afi.LiquidationValue())

	return computePerc(afi.Debt(), afi.LiquidationValue())
}

func (l *Liability) Debt() *big.Int {
	return new(big.Int).Add(l.Principal, l.Interest)
}

func computePerc(numerator, denominator *big.Int) *big.Float {
	if denominator.Cmp(big.NewInt(0)) == 0 {
		return new(big.Float).SetInf(true)
	}

	return util.ToFIL(util.DivWad(numerator, denominator))
}
