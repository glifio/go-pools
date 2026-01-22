package token

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestCheckAirdropEligibility_DuplicateOwner(t *testing.T) {
	// Test Agent 139 - should get 5,903.43 FIL
	agent139 := common.HexToAddress("0x207201C13A0640c99152f9aF05C16f95f27d659F")
	expectedAmount139 := new(big.Float).SetFloat64(5903.43)

	amount139, claimer139, err := CheckAirdropEligibility(agent139, false)
	if err != nil {
		t.Fatalf("Agent 139 eligibility check failed: %v", err)
	}

	// Check amount is approximately 5,903.43 FIL (within 0.01 FIL tolerance)
	diff139 := new(big.Float).Sub(amount139, expectedAmount139)
	diff139.Abs(diff139)
	tolerance := new(big.Float).SetFloat64(0.01)

	if diff139.Cmp(tolerance) > 0 {
		t.Errorf("Agent 139 amount incorrect: got %s FIL, want ~5903.43 FIL", amount139.Text('f', 2))
	}

	// Check claimer is the expected owner
	expectedOwner := common.HexToAddress("0x3e39a95489a06aeb044c4438790f74cf27a8ca82")
	if claimer139.Hex() != expectedOwner.Hex() {
		t.Errorf("Agent 139 claimer incorrect: got %s, want %s", claimer139.Hex(), expectedOwner.Hex())
	}

	t.Logf("✓ Agent 139: %s FIL -> claimer %s", amount139.Text('f', 2), claimer139.Hex())

	// Test Agent 149 - should get 487,247.74 FIL
	agent149 := common.HexToAddress("0x992309c1116Bb9DFB52d6793Be258E3eA9F0D76a")
	expectedAmount149 := new(big.Float).SetFloat64(487247.74)

	amount149, claimer149, err := CheckAirdropEligibility(agent149, false)
	if err != nil {
		t.Fatalf("Agent 149 eligibility check failed: %v", err)
	}

	// Check amount is approximately 487,247.74 FIL (within 0.01 FIL tolerance)
	diff149 := new(big.Float).Sub(amount149, expectedAmount149)
	diff149.Abs(diff149)

	if diff149.Cmp(tolerance) > 0 {
		t.Errorf("Agent 149 amount incorrect: got %s FIL, want ~487247.74 FIL", amount149.Text('f', 2))
	}

	// Check claimer is the expected owner
	if claimer149.Hex() != expectedOwner.Hex() {
		t.Errorf("Agent 149 claimer incorrect: got %s, want %s", claimer149.Hex(), expectedOwner.Hex())
	}

	t.Logf("✓ Agent 149: %s FIL -> claimer %s", amount149.Text('f', 2), claimer149.Hex())

	// Ensure the amounts are different
	if amount139.Cmp(amount149) == 0 {
		t.Error("Agent 139 and Agent 149 should have different amounts, but got the same")
	}
}
