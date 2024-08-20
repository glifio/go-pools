package econ

import (
	"fmt"
	"math/big"

	"github.com/glifio/go-pools/constants"
)

type BaseFi struct {
	AvailableBalance *big.Int `json:"availableBalance"`
	LockedRewards    *big.Int `json:"lockedRewards"`
	InitialPledge    *big.Int `json:"initialPledge"`
	FeeDebt          *big.Int `json:"feeDebt"`
	TerminationFee   *big.Int `json:"terminationFee"`
}

type AgentFi struct {
	BaseFi

	Principal *big.Int `json:"principal"`
}

type BaseFiRet struct {
	AvailableBalance     string `json:"availableBalance"`
	LockedRewards        string `json:"lockedRewards"`
	InitialPledge        string `json:"initialPledge"`
	FeeDebt              string `json:"feeDebt"`
	TerminationFee       string `json:"terminationFee"`
	Balance              string `json:"balance"`
	LiquidationValue     string `json:"liquidationValue"`
	MaxBorrowAndSeal     string `json:"maxBorrowAndSeal"`
	MaxBorrowAndWithdraw string `json:"maxBorrowAndWithdraw"`
}

type AgentFiRet struct {
	BaseFiRet

	Principal     string `json:"principal"`
	BorrowLimit   string `json:"borrowLimit"`
	WithdrawLimit string `json:"withdrawLimit"`
	MarginCall    string `json:"marginCall"`
	LeverageRatio string `json:"leverageRatio"`
}

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

func (bfi *BaseFi) String() *BaseFiRet {
	return &BaseFiRet{
		AvailableBalance:     bfi.AvailableBalance.String(),
		LockedRewards:        bfi.LockedRewards.String(),
		InitialPledge:        bfi.InitialPledge.String(),
		FeeDebt:              bfi.FeeDebt.String(),
		TerminationFee:       bfi.TerminationFee.String(),
		MaxBorrowAndSeal:     bfi.MaxBorrowAndSeal().String(),
		MaxBorrowAndWithdraw: bfi.MaxBorrowAndWithdraw().String(),
		Balance:              bfi.Balance().String(),
		LiquidationValue:     bfi.LiquidationValue().String(),
	}
}

func (afi *AgentFi) String() *AgentFiRet {
	return &AgentFiRet{
		BaseFiRet:     *afi.BaseFi.String(),
		Principal:     afi.Principal.String(),
		BorrowLimit:   afi.BorrowLimit().String(),
		WithdrawLimit: afi.WithdrawLimit().String(),
		MarginCall:    afi.MarginCall().String(),
		LeverageRatio: fmt.Sprintf("%0.02f", afi.LeverageRatio()),
	}
}
