package jws

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/sdk"
	"github.com/glifio/go-pools/types"
	"github.com/glifio/go-pools/util"
	"github.com/golang-jwt/jwt/v5"
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
	ctx := context.Background()

	agentAddr, target, value, mockQueries, signerPrivateKey, signerAddr := setup()

	mockQueries.On("ChainHeight", ctx).Return(big.NewInt(100), nil)
	mockQueries.On("AgentRequester", ctx, agentAddr).Return(signerAddr, nil)

	// create a mock PoolsSDK and set the expected return value for Query to be the mock
	mockSDK := new(MockPoolsSDK)
	mockSDK.On("Query").Return(mockQueries)

	jws, err := SignJWS(ctx, agentAddr, target, value, constants.MethodBorrow, signerPrivateKey, mockSDK)
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

func TestBadJWSPubkey(t *testing.T) {
	ctx := context.Background()

	agentAddr, target, value, mockQueries, signerPrivateKey, signerAddr := setup()

	mockQueries.On("ChainHeight", ctx).Return(big.NewInt(100), nil)
	mockQueries.On("AgentRequester", ctx, agentAddr).Return(signerAddr, nil)

	// create a mock PoolsSDK and set the expected return value for Query to be the mock
	mockSDK := new(MockPoolsSDK)
	mockSDK.On("Query").Return(mockQueries)

	privateKey, _ := crypto.GenerateKey()

	actionSelector, _ := util.MethodStrToBytes(constants.MethodBorrow)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, RequestClaims{
		AgentAddr:       agentAddr,
		RequesterPubKey: crypto.FromECDSAPub(&privateKey.PublicKey),
		Target:          target,
		Value:           value,
		ActionSelector:  actionSelector,
		EpochHeight:     big.NewInt(100),
	})

	// sign with the signer's private key, different from the public key encoded in the claims
	tokenStr, err := token.SignedString(signerPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = VerifyJWS(ctx, tokenStr, mockSDK)
	if err == nil {
		t.Fatal("expected crypto/ecdsa verification error")
	}

	containsInvalidECDSAVerificationErr := strings.Contains(err.Error(), jwt.ErrECDSAVerification.Error())
	containsInvalidSigErr := strings.Contains(err.Error(), jwt.ErrSignatureInvalid.Error())

	if !containsInvalidECDSAVerificationErr && !containsInvalidSigErr {
		t.Fatal("expected crypto/ecdsa verification error")
	}
}

func TestStaleJWS(t *testing.T) {
	ctx := context.Background()

	chainHeight := big.NewInt(100)

	agentAddr, target, value, mockQueries, signerPrivateKey, signerAddr := setup()
	// set the chain height to be 100
	mockQueries.On("ChainHeight", ctx).Return(chainHeight, nil)
	mockQueries.On("AgentRequester", ctx, agentAddr).Return(signerAddr, nil)

	// create a mock PoolsSDK and set the expected return value for Query to be the mock
	mockSDK := new(MockPoolsSDK)
	mockSDK.On("Query").Return(mockQueries)

	actionSelector, _ := util.MethodStrToBytes(constants.MethodBorrow)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, RequestClaims{
		AgentAddr:       agentAddr,
		RequesterPubKey: crypto.FromECDSAPub(&signerPrivateKey.PublicKey),
		Target:          target,
		Value:           value,
		ActionSelector:  actionSelector,
		// encode the chain height to be
		EpochHeight: new(big.Int).Sub(chainHeight, EXPIRATION_EPOCH_BUFFER),
	})

	// sign with the signer's private key, different from the public key encoded in the claims
	tokenStr, err := token.SignedString(signerPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = VerifyJWS(ctx, tokenStr, mockSDK)
	if err != TokenExpiredErr {
		t.Fatal("expected stale JWS error")
	}
}

var ZERO_ADDR = common.Address{}

func setup() (agentAddr common.Address, target address.Address, value *big.Int, mockQueries *MockFEVMQueries, privateKey *ecdsa.PrivateKey, signerAddr common.Address) {
	privateKey, _ = crypto.GenerateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	signerAddr = crypto.PubkeyToAddress(*publicKeyECDSA)

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
	agentAddr = common.HexToAddress("0xE8de74929076468BC59b079BDA683bc5bb813a39")
	target, _ = address.NewFromString("f01869494")
	value = big.NewInt(1000)

	// create mock FEVMQueries and set the expected return value for ChainHeight
	mockQueries = &MockFEVMQueries{
		FEVMQueries: fevmConnect.Query(),
	}

	return agentAddr, target, value, mockQueries, privateKey, signerAddr
}
