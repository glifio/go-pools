package econ

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestFetchAgentEcon(t *testing.T) {
	testCases := []struct {
		name             string
		input            common.Address
		expectError      bool
		expectPositiveLV bool
	}{
		{
			name:             "valid agent",
			input:            common.HexToAddress("0x64Ea1aD49A78B6A6878c4975566b8953b1Ef1e79"),
			expectError:      false,
			expectPositiveLV: true,
		},
		{
			name:             "empty agent",
			input:            common.HexToAddress("0xDBa96B0FDbc87C7eEb641Ee37EAFC55B355079E4"),
			expectError:      false,
			expectPositiveLV: false,
		},
		{
			name:             "invalid agent",
			input:            common.HexToAddress("0x00"),
			expectError:      false,
			expectPositiveLV: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			afi, err := GetAgentFiFromAPI(tc.input)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			lv := afi.LiquidationValue()
			maxBorrow := afi.MaxBorrowAndSeal()
			maxWithdraw := afi.MaxBorrowAndWithdraw()

			if tc.expectPositiveLV {
				assert.True(t, lv.Cmp(big.NewInt(0)) == 1)
				assert.True(t, maxBorrow.Cmp(big.NewInt(0)) == 1)
				assert.True(t, maxWithdraw.Cmp(big.NewInt(0)) == 1)
			} else {
				assert.True(t, lv.Cmp(big.NewInt(0)) == 0)
				assert.True(t, maxBorrow.Cmp(big.NewInt(0)) == 0)
				assert.True(t, maxWithdraw.Cmp(big.NewInt(0)) == 0)
			}
		})
	}
}
