package vc

import (
	"math/big"
	"reflect"
	"testing"
)

func TestEncodeDecodeClaims(t *testing.T) {
	// create new agent data
	agentData := &AgentData{
		AgentValue:                  big.NewInt(1),
		CollateralValue:             big.NewInt(2),
		ExpectedDailyFaultPenalties: big.NewInt(3),
		ExpectedDailyRewards:        big.NewInt(4),
		Gcred:                       big.NewInt(5),
		QaPower:                     big.NewInt(6),
		Principal:                   big.NewInt(7),
		FaultySectors:               big.NewInt(8),
		LiveSectors:                 big.NewInt(9),
		GreenScore:                  big.NewInt(10),
	}

	cl, err := AbiEncodeClaim(*agentData)
	if err != nil {
		t.Fatal(err)
	}

	agentData2, err := AbiDecodeClaim(cl)
	if err != nil {
		t.Fatal(err)
	}

	if !compareAgentData(agentData, agentData2) {
		t.Fatal("AgentData not equal")
	}
}

func compareAgentData(a, b *AgentData) bool {
	valA := reflect.ValueOf(a).Elem()
	valB := reflect.ValueOf(b).Elem()

	for i := 0; i < valA.NumField(); i++ {
		valueA := valA.Field(i).Interface().(*big.Int)
		valueB := valB.Field(i).Interface().(*big.Int)

		if valueA.Cmp(valueB) != 0 {
			return false
		}
	}

	return true
}
