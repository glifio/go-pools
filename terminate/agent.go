package terminate

import (
	"math/big"

	"github.com/glifio/go-pools/constants"
)

func (ats PreviewAgentTerminationSummary) TotalAssets() *big.Int {
	totalAssets := new(big.Int).Add(ats.MinersAvailableBal, ats.AgentAvailableBal)
	totalAssets.Add(totalAssets, ats.InitialPledge)
	totalAssets.Add(totalAssets, ats.VestingBalance)
	return totalAssets
}

func (ats PreviewAgentTerminationSummary) LiquidationValue() *big.Int {
	totalAssets := ats.TotalAssets()
	// if the total assets is less than or equal the termination penalty, we return 0
	if totalAssets.Cmp(ats.TerminationPenalty) <= 0 {
		return big.NewInt(0)
	}
	// liquidation value = total assets - termination penalty
	return new(big.Int).Sub(totalAssets, ats.TerminationPenalty)
}

// 1e18 = 100%
func (ats PreviewAgentTerminationSummary) RecoveryRate() *big.Int {
	if ats.InitialPledge.Cmp(ats.TerminationPenalty) == -1 || ats.InitialPledge.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	// recovery rate = (initial pledge - termination penalty) / initial pledge
	recoveryRate := new(big.Int).Sub(ats.InitialPledge, ats.TerminationPenalty)
	recoveryRate.Mul(recoveryRate, constants.WAD)
	recoveryRate.Quo(recoveryRate, ats.InitialPledge)

	return recoveryRate
}

func (ats PreviewAgentTerminationSummary) LTV(principal *big.Int) *big.Int {
	liquidationValue := ats.LiquidationValue()
	ltv := big.NewInt(0)
	if liquidationValue.Cmp(big.NewInt(0)) > 0 && principal.Cmp(big.NewInt(0)) > 0 {
		// add wad precision in prior to the div to keep atto denomination before we convert to percentage
		ltv.Mul(principal, constants.WAD)
		ltv.Div(ltv, liquidationValue)
	}

	return ltv
}

func (ats PreviewAgentTerminationSummary) Margin(principal *big.Int) (*AgentMargin, error) {
	margin, maxBorrowAndSeal, maxBorrowAndWithdraw := Margin(ats.TotalAssets(), ats.TerminationPenalty)

	baseMargin := BaseMargin{
		AvailableBalance:     ats.AgentAvailableBal,
		LockedRewards:        ats.VestingBalance,
		InitialPledge:        ats.InitialPledge,
		Margin:               margin,
		TerminationPenalty:   ats.TerminationPenalty,
		MaxBorrowAndSeal:     maxBorrowAndSeal,
		MaxBorrowAndWithdraw: maxBorrowAndWithdraw,
	}

	borrowLimit := big.NewInt(0).Sub(maxBorrowAndSeal, principal)

	leverageRatio, _ := new(big.Float).Quo(
		new(big.Float).SetInt(ats.LiquidationValue()),
		new(big.Float).SetInt(margin),
	).Float64()

	// calculate margin call as (principal * 1e18) / liquidationDTL
	marginCall := big.NewInt(0).Div(big.NewInt(0).Mul(principal, big.NewInt(1e18)), constants.LIQUIDATION_DTL)

	// calculate withdraw limit as liquidation value - (principal * 1e18 / 75e16)
	withdrawLimit := big.NewInt(0).Sub(
		ats.LiquidationValue(),
		big.NewInt(0).Div(
			big.NewInt(0).Mul(principal, big.NewInt(1e18)),
			constants.MAX_BORROW_DTL,
		),
	)

	res := AgentMargin{
		AgentBalance:  ats.AgentAvailableBal,
		Principal:     principal,
		LeverageRatio: leverageRatio,
		WithdrawLimit: withdrawLimit,
		BaseMargin:    baseMargin,
		BorrowLimit:   borrowLimit,
		MarginCall:    marginCall,
	}

	return &res, nil
}

func (cs *AgentCollateralStats) Summarize() PreviewAgentTerminationSummary {
	availBal := big.NewInt(0)
	initialPledge := big.NewInt(0)
	vestingBal := big.NewInt(0)
	for _, miner := range cs.MinersTerminationStats {
		availBal.Add(availBal, miner.Available)
		initialPledge.Add(initialPledge, miner.Pledged)
		vestingBal.Add(vestingBal, miner.Vesting)
	}

	return PreviewAgentTerminationSummary{
		TerminationPenalty: cs.TerminationPenalty,
		InitialPledge:      initialPledge,
		VestingBalance:     vestingBal,
		MinersAvailableBal: availBal,
		AgentAvailableBal:  cs.AvailableBalance,
	}
}
