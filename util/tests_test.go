package util

import (
	"math/big"
	"testing"
)

func TestApproxEqualRel(t *testing.T) {
	testCases := []struct {
		name           string
		a, b, rangeWAD *big.Int
		want           bool
	}{
		{
			name:     "Equal values",
			a:        big.NewInt(1e18), // 1 WAD
			b:        big.NewInt(1e18), // 1 WAD
			rangeWAD: big.NewInt(1e17), // 10%
			want:     true,
		},
		{
			name:     "Within range",
			a:        big.NewInt(1e18),        // 1 WAD
			b:        big.NewInt(1e18 + 5e16), // 1.05 WAD
			rangeWAD: big.NewInt(1e17),        // 10%
			want:     true,
		},
		{
			name:     "Outside range",
			a:        big.NewInt(1e18),        // 1 WAD
			b:        big.NewInt(1e18 + 2e17), // 1.2 WAD
			rangeWAD: big.NewInt(1e17),        // 10%
			want:     false,
		},
		{
			name:     "A is zero",
			a:        big.NewInt(0),    // 0 WAD
			b:        big.NewInt(1e18), // 1 WAD
			rangeWAD: big.NewInt(1e17), // 10%
			want:     false,            // Cannot compare if 'a' is zero
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := AssertApproxEqAbs(tc.a, tc.b, tc.rangeWAD)
			if got != tc.want {
				t.Errorf("approximatelyEqual(%v, %v, %v) = %v; want %v", tc.a, tc.b, tc.rangeWAD, got, tc.want)
			}
		})
	}
}
