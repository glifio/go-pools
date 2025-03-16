package token

import (
	"math/big"
)

// ComputePrice calculates a Q96.64 fixed-point price to a *big.Float denominated in base (not atto)
func ComputePrice(price *big.Int) *big.Float {
	if price == nil {
		return big.NewFloat(0)
	}

	// Convert to big.Float for decimal arithmetic
	priceFloat := new(big.Float).SetInt(price)
	priceFloat.SetPrec(256)

	// Calculate 2^96
	two96 := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(2), big.NewInt(96), nil))

	// Divide by 2^96
	sqrtPrice := new(big.Float).Quo(priceFloat, two96)

	// Square the result to get the actual price
	result := new(big.Float).Mul(sqrtPrice, sqrtPrice)

	// The result is now in terms of raw token units, so we need to adjust for decimals
	// Since both tokens have 18 decimals, we don't need additional adjustment

	return result
}

func GLFToFIL(sqrtPriceX96 *big.Int) *big.Float {
	// return 1/Price(Token1/Token0) to get the price of Token0 in Token1
	return new(big.Float).Quo(big.NewFloat(1), ComputePrice(sqrtPriceX96))
}

func FILToGLF(sqrtPriceX96 *big.Int) *big.Float {
	return ComputePrice(sqrtPriceX96)
}
