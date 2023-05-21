package util

import "testing"

func TestDeriveNativeFromSECP256K1Key(t *testing.T) {
	PRIVATE_KEY := "yk+/lP+dCaUMgXUJji6LV5HzKE8lFJAEML99xlBmd3I="

	addr, err := DeriveNativeFromSECP256K1KeyString(PRIVATE_KEY)
	if err != nil {
		t.Errorf("DeriveNativeFromSECP256K1KeyString(%s) failed: %s", PRIVATE_KEY, err)
	}

	EXP := "f1kymj6ucraqgbyvwbukylaskyuzx3ix2zvjfhaqi"

	if addr.String() != "f1kymj6ucraqgbyvwbukylaskyuzx3ix2zvjfhaqi" {
		t.Errorf("DeriveNativeFromSECP256K1KeyString(%s) expected %s, got %s", PRIVATE_KEY, EXP, addr.String())
	}
}
