package econ

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/deploy"
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
			afi, err := GetAgentFiFromAPI(tc.input, deploy.MainnetEventsURL)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			lv := afi.LiquidationValue()
			maxBorrow := afi.MaxBorrowAndSeal(constants.MAX_BORROW_DTL)
			maxWithdraw := afi.MaxBorrowAndWithdraw(constants.MAX_BORROW_DTL)

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

func TestFetchBaseFis(t *testing.T) {
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
			miners, bfis, err := GetBaseFisFromAPI(tc.input, deploy.MainnetEventsURL)
			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, len(miners), len(bfis))

			// now we get the agentFi for each agent
			afi, err := GetAgentFiFromAPI(tc.input, deploy.MainnetEventsURL)
			if err != nil {
				t.Fatalf("error getting agentFi from API: %v", err)
			}

			// create a new afi based on the bfi values
			afi2 := NewAgentFi(
				afi.SpendableBalance,
				Liability{
					Principal: afi.Principal,
					Interest:  afi.Interest,
				},
				bfis,
			)

			assert.True(t, afi2.LiquidationValue().Cmp(afi.LiquidationValue()) == 0)
			assert.True(t, afi2.MaxBorrowAndSeal(constants.MAX_BORROW_DTL).Cmp(afi.MaxBorrowAndSeal(constants.MAX_BORROW_DTL)) == 0)
			assert.True(t, afi2.MaxBorrowAndWithdraw(constants.MAX_BORROW_DTL).Cmp(afi.MaxBorrowAndWithdraw(constants.MAX_BORROW_DTL)) == 0)
			assert.True(t, afi2.BorrowLimit(constants.MAX_BORROW_DTL).Cmp(afi.BorrowLimit(constants.MAX_BORROW_DTL)) == 0)
			assert.True(t, afi2.WithdrawLimit(constants.MAX_BORROW_DTL).Cmp(afi.WithdrawLimit(constants.MAX_BORROW_DTL)) == 0)

			lv := afi.LiquidationValue()
			maxBorrow := afi.MaxBorrowAndSeal(constants.MAX_BORROW_DTL)
			maxWithdraw := afi.MaxBorrowAndWithdraw(constants.MAX_BORROW_DTL)

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

func TestGetPoolMetricsFromAPI(t *testing.T) {
	metrics, err := GetPoolMetricsFromAPI(deploy.MainnetEventsURL)
	assert.NoError(t, err)
	assert.NotNil(t, metrics)
}
