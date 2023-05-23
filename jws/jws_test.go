package jws

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/sdk"
	"github.com/glifio/go-pools/types"
	"github.com/stretchr/testify/mock"
)

// MockPoolsSDK is a mock implementation of PoolsSDK
type MockPoolsSDK struct {
	mock.Mock
}

func (m *MockPoolsSDK) Query() types.FEVMQueries {
	return m.Called().Get(0).(types.FEVMQueries)
}

func (m *MockPoolsSDK) Act() types.FEVMActions {
	return m.Called().Get(0).(types.FEVMActions)
}

func (m *MockPoolsSDK) Extern() types.FEVMExtern {
	return m.Called().Get(0).(types.FEVMExtern)
}

// MockFEVMQueries is a mock implementation of FEVMQueries
type MockFEVMQueries struct {
	types.FEVMQueries
	mock.Mock
}

func (m *MockFEVMQueries) ChainHeight(ctx context.Context) (*big.Int, error) {
	args := m.Called(ctx)
	return args.Get(0).(*big.Int), args.Error(1)
}

func (m *MockFEVMQueries) AgentRequester(ctx context.Context, agentAddr common.Address) (common.Address, error) {
	args := m.Called(ctx, agentAddr)
	return args.Get(0).(common.Address), args.Error(1)
}

func TestSignVerifyJWS(t *testing.T) {
	ZERO_ADDR := common.Address{}
	ctx := context.Background()

	privateKey, _ := crypto.GenerateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	signerAddr := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("Signer address: ", signerAddr)

	fevmConnect := sdk.InitFEVMConnection(
		ZERO_ADDR,
		ZERO_ADDR,
		ZERO_ADDR,
		ZERO_ADDR,
		ZERO_ADDR,
		ZERO_ADDR,
		ZERO_ADDR,
		ZERO_ADDR,
		"",
		"",
		"",
		"",
		common.Big0,
	)
	var agentAddr = common.HexToAddress("0xE8de74929076468BC59b079BDA683bc5bb813a39")
	var target, _ = address.NewFromString("f01869494")
	var value = big.NewInt(1000)

	// create mock FEVMQueries and set the expected return value for ChainHeight
	mockQueries := &MockFEVMQueries{
		FEVMQueries: fevmConnect.Query(),
	}
	mockQueries.On("ChainHeight", ctx).Return(big.NewInt(100), nil)
	mockQueries.On("AgentRequester", ctx, agentAddr).Return(signerAddr, nil)

	// create a mock PoolsSDK and set the expected return value for Query to be the mock
	mockSDK := new(MockPoolsSDK)
	mockSDK.On("Query").Return(mockQueries)

	jws, err := SignJWS(ctx, agentAddr, target, value, constants.MethodBorrow, privateKey, mockSDK)
	if err != nil {
		t.Fatal(err)
	}

	if jws == "" {
		t.Fatal("jws is empty")
	}

	claims, err := VerifyJWS(ctx, jws, mockSDK)
	if err != nil {
		t.Fatal(err)
	}

	if claims.AgentAddr != agentAddr {
		t.Fatal("agent address does not match")
	}

	if claims.Target != target {
		t.Fatal("target address does not match")
	}

	if claims.Value.Cmp(value) != 0 {
		t.Fatal("value does not match")
	}
}
