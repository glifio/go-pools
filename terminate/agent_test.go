package terminate

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/glifio/go-pools/util"
)

func TestAgentLiquidationValue(t *testing.T) {
	var testCases = []struct {
		terminationStats PreviewAgentTerminationSummary
		expRes           *big.Int
	}{
		// case 1 - 50% recovery rate
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(500)),
			InitialPledge:      util.ToAtto(big.NewFloat(1000)),
			AvailableBalance:   util.ToAtto(big.NewFloat(100)),
			VestingBalance:     util.ToAtto(big.NewFloat(100)),
		}, util.ToAtto(big.NewFloat(650))},
		// case 2 - 90% recovery rate, 0 available
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(100)),
			InitialPledge:      util.ToAtto(big.NewFloat(1000)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(100)),
		}, util.ToAtto(big.NewFloat(1000))},
		// case 3 termination bigger than initial pledge
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(1000)),
			InitialPledge:      util.ToAtto(big.NewFloat(100)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(0)),
		}, util.ToAtto(big.NewFloat(0))},
		// case 4 - 0's
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(0)),
			InitialPledge:      util.ToAtto(big.NewFloat(0)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(0)),
		}, util.ToAtto(big.NewFloat(0))},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			answer := tc.terminationStats.LiquidationValue()
			if answer.Cmp(tc.expRes) != 0 {
				t.Errorf("expected %s, got %s", tc.expRes, answer)
			}
		})
	}
}

func TestAgentRecoveryRate(t *testing.T) {
	var testCases = []struct {
		terminationStats PreviewAgentTerminationSummary
		expRes           *big.Float
	}{
		// case 1 - 50% recovery rate
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(500)),
			InitialPledge:      util.ToAtto(big.NewFloat(1000)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(0)),
		}, big.NewFloat(50.00)},
		// case 2 - 90% recovery rate, 0 available
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(100)),
			InitialPledge:      util.ToAtto(big.NewFloat(1000)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(100)),
		}, big.NewFloat(90.00)},
		// case 3 termination bigger than initial pledge
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(1000)),
			InitialPledge:      util.ToAtto(big.NewFloat(1000)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(0)),
		}, big.NewFloat(0)},
		// case 4 - 0's
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(0)),
			InitialPledge:      util.ToAtto(big.NewFloat(0)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(0)),
		}, big.NewFloat(0)},
		{PreviewAgentTerminationSummary{
			TerminationPenalty: util.ToAtto(big.NewFloat(0)),
			InitialPledge:      util.ToAtto(big.NewFloat(100)),
			AvailableBalance:   util.ToAtto(big.NewFloat(0)),
			VestingBalance:     util.ToAtto(big.NewFloat(0)),
		}, big.NewFloat(100)},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test case: %v", i), func(t *testing.T) {
			answer := tc.terminationStats.RecoveryRate()
			if answer.Cmp(tc.expRes) != 0 {
				t.Errorf("expected %0.09f, got %0.09f", tc.expRes, answer)
			}
		})
	}
}
