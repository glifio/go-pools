package deploy

import (
	"testing"

	"github.com/glifio/go-pools/constants"
)

func TestGetLPPlusIssuanceSeconds(t *testing.T) {
	testCases := []struct {
		name      string
		timestamp int
		want      int
	}{
		{
			name:      "before launch",
			timestamp: constants.LP_PLUS_MAINNET_WINDOW_START_TS - 1,
			want:      constants.LP_PLUS_ISSUANCE_SECONDS_PRE_LAUNCH,
		},
		{
			name:      "after launch",
			timestamp: constants.LP_PLUS_MAINNET_WINDOW_START_TS + 1,
			want:      constants.LP_PLUS_ISSUANCE_SECONDS,
		},
		{
			name:      "hardcoded launch timestamp before",
			timestamp: 1765823486,
			want:      constants.LP_PLUS_ISSUANCE_SECONDS_PRE_LAUNCH,
		},
		{
			name:      "hardcoded launch timestamp after",
			timestamp: 1765987200,
			want:      constants.LP_PLUS_ISSUANCE_SECONDS,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := GetLPPlusIssuanceSeconds(tc.timestamp)
			if got != tc.want {
				t.Errorf("GetLPPlusIssuanceSeconds(%d) = %d; want %d", tc.timestamp, got, tc.want)
			}
		})
	}
}
