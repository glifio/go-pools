package util

import "math/big"

var _WAD = big.NewInt(1e18)

func MulWad(x *big.Int, y *big.Int) *big.Int {
	z := new(big.Int).Div(new(big.Int).Mul(x, y), _WAD)
	return z
}

func DivWad(x *big.Int, y *big.Int) *big.Int {
	z := new(big.Int).Div(new(big.Int).Mul(x, _WAD), y)
	return z
}

func Diff(x *big.Int, y *big.Int) *big.Float {
	// calc the difference between to big int numbers
	// Convert big.Int to big.Float for percentage calculation
	aFloat := new(big.Float).SetInt(x)
	bFloat := new(big.Float).SetInt(y)

	// Calculate the difference (a - b)
	diff := new(big.Float).Sub(aFloat, bFloat)

	// if diff is zero return
	if diff.Cmp(big.NewFloat(0)) == 0 {
		return big.NewFloat(0)
	}

	// make diff a positive number
	if diff.Sign() == -1 {
		diff.Neg(diff)
	}

	// Calculate the percentage difference (diff / a * 100)
	percentDiff := new(big.Float).Quo(diff, aFloat)
	percentDiff.Mul(percentDiff, big.NewFloat(100))

	return percentDiff
}
