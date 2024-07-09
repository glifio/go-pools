package terminate

import (
	"fmt"
	"math/big"
	"testing"
)

func TestTermPenaltyOnPledge(t *testing.T) {
	filToPledge := big.NewInt(2 * 1e18)
	var sectorSize uint64 = 34359738368
	var ratioQAP float64 = 1.0
	cost, penalty, err := TermPenaltyOnPledge(filToPledge, sectorSize, ratioQAP)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Jim cost: %v\n", cost)
	fmt.Printf("Jim penalty: %v\n", penalty)
	t.Fatal("Jim2")
}
