package util

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/manifest"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/assert"
)

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

type MockFullNodeAPI struct {
	api.FullNode
}

var idStr = "f01234"
var maskedIDStr = "0xFF000000000000000000000000000000000004d2"

func (m *MockFullNodeAPI) StateLookupID(ctx context.Context, addr address.Address, tsk types.TipSetKey) (address.Address, error) {
	return address.NewFromString(idStr)
}

var evmActorID = "f05678"
var ethAccountID = "f09876"

func (m *MockFullNodeAPI) StateGetActor(ctx context.Context, actor address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	code := cid.Cid{}
	if actor.String() == evmActorID {
		code, _ = actors.GetActorCodeID(actorstypes.Version(actors.LatestVersion), manifest.EvmKey)
	} else if actor.String() == ethAccountID {
		code, _ = actors.GetActorCodeID(actorstypes.Version(actors.LatestVersion), manifest.EthAccountKey)
	}

	return &types.Actor{
		Code:    code,
		Head:    cid.Cid{},
		Nonce:   0,
		Balance: types.NewInt(0),
	}, nil
}

func TestParseAddress(t *testing.T) {
	// Create a mock FullNodeAPI
	mockAPI := &MockFullNodeAPI{}

	testCases := []struct {
		name        string
		input       string
		expected    common.Address
		expectError bool
	}{
		{
			name:        "Valid Ethereum Address",
			input:       "0x3972E844729522d367BFA1D64368346D7ccEEa59",
			expected:    common.HexToAddress("0x3972E844729522d367BFA1D64368346D7ccEEa59"),
			expectError: false,
		},
		{
			name:        "Valid Filecoin ID Address",
			input:       idStr,
			expected:    common.HexToAddress(maskedIDStr),
			expectError: false,
		},
		{
			name:        "Valid Filecoin Account f1 Address",
			input:       "f1ys5qqiciehcml3sp764ymbbytfn3qoar5fo3iwy",
			expected:    common.HexToAddress(maskedIDStr),
			expectError: false,
		},
		{
			name:        "Valid Filecoin Account f3 Address",
			input:       "f3vpyybzycb3wvhwkxcrodn3rqv66sd5hfho4lfq6p6igmrlgyb22v3ekdghp6km47ioki3gfo4zb4ezirhfaq",
			expected:    common.HexToAddress(maskedIDStr),
			expectError: false,
		},
		{
			name:        "Valid Filecoin Account f4 Address",
			input:       "f410fmdqxonrwz5peuit5tlbe6ih6zibu5ys223xctfi",
			expected:    common.HexToAddress("0x60E1773636CF5E4A227d9AC24F20fEca034ee25A"),
			expectError: false,
		},
		{
			name:        "Invalid Address",
			input:       "invalid_address",
			expectError: true,
		},
		{
			name:        "EVM Actor",
			input:       evmActorID,
			expectError: true,
		},
		{
			name:        "Eth Account",
			input:       ethAccountID,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := ParseAddress(context.Background(), tc.input, mockAPI, false)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				fmt.Println(result.String())
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, result)
			}
		})
	}
}
