package util

import (
	"errors"
	"log"
	"math/big"
	"testing"
)

func TestToFIL(t *testing.T) {
	var tests = []struct {
		bal *big.Int
		fil *big.Float
	}{
		{mustNewBigInt("0"), big.NewFloat(0)},
		{mustNewBigInt("1000000000000000000"), big.NewFloat(1)},
		{mustNewBigInt("100000000000000000000"), big.NewFloat(100)},
	}

	for _, test := range tests {
		fil := ToFIL(test.bal)
		if fil.Cmp(test.fil) != 0 {
			t.Errorf("ToFIL(%s) expected %d, got %d", test.bal, test.fil, fil)
		}
	}
}

func TestToAttoFILFromString(t *testing.T) {
	var tests = []struct {
		bal     string
		attofil *big.Int
	}{
		{"10", mustNewBigInt("10000000000000000000")},
		{"2", mustNewBigInt("2000000000000000000")},
		{"1.1009", mustNewBigInt("1100900000000000000")},
	}

	for _, test := range tests {
		attoFIL, err := parseFILAmount(test.bal)
		if err != nil {
			t.Errorf("parseFILAmount(%s) failed: %s", test.bal, err)
		}

		if attoFIL.Cmp(test.attofil) != 0 {
			t.Errorf("ToFIL(%s) expected %d, got %d", test.bal, test.attofil, attoFIL)
		}
	}
}

func parseFILAmount(amount string) (*big.Int, error) {
	amt, ok := new(big.Float).SetString(amount)
	if !ok {
		return nil, errors.New("invalid amount")
	}

	return ToAtto(amt), nil
}

func mustNewBigInt(val string) *big.Int {
	res, success := big.NewInt(0).SetString(val, 10)
	if !success {
		log.Panic("big int failed")
	}
	return res
}
