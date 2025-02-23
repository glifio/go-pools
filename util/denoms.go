package util

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/params"
)

func ToAtto(bal *big.Float) *big.Int {
	// Convert base to atto by multiplying by 10^18
	f := new(big.Float).Mul(bal, big.NewFloat(params.Ether))
	atto := new(big.Int)
	f.Int(atto) // Convert float to integer
	return atto
}

func ToFIL(atto *big.Int) *big.Float {
	// Convert atto to base by dividing by 10^18
	f := new(big.Float).SetInt(atto)
	baseValue := new(big.Float).Quo(f, big.NewFloat(params.Ether))
	return baseValue
}

func ParseFILStrToAtto(filStr string) (*big.Int, error) {
	// Split into integer and fractional parts
	parts := strings.Split(filStr, ".")
	if len(parts) > 2 {
		return nil, fmt.Errorf("invalid amount format: %s", filStr)
	}

	// Handle the integer part
	intPart, ok := new(big.Int).SetString(parts[0], 10)
	if !ok {
		return nil, fmt.Errorf("invalid integer part: %s", parts[0])
	}

	// Multiply integer part by 10^18
	result := new(big.Int).Mul(intPart, big.NewInt(1e18))

	// Handle fractional part if it exists
	if len(parts) == 2 {
		fracPart := parts[1]
		// Pad or truncate to 18 digits (attoFIL precision)
		if len(fracPart) > 18 {
			fracPart = fracPart[:18] // Truncate to 18 decimals
		}
		fracPart = fracPart + strings.Repeat("0", 18-len(fracPart)) // Pad with zeros

		fracInt, ok := new(big.Int).SetString(fracPart, 10)
		if !ok {
			return nil, fmt.Errorf("invalid fractional part: %s", fracPart)
		}
		result.Add(result, fracInt)
	}

	return result, nil
}

var WAD = big.NewInt(1000000000000000000)
