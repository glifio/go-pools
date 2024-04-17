package jws

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/mock"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jimpick/go-ethereum/common"
	"github.com/jimpick/go-ethereum/crypto"
)

func TestSignVerifyJWS(t *testing.T) {
	ctx := context.Background()

	agentAddr, target, value, signerPrivateKey, signerAddr := setup()

	mockFEVMQueries := mock.NewFEVMQueries(t)
	mockFEVMQueries.On("ChainHeight", ctx).Return(big.NewInt(100), nil)
	mockFEVMQueries.On("AgentRequester", ctx, agentAddr).Return(signerAddr, nil)

	jws, err := SignJWS(ctx, agentAddr, target, value, constants.MethodBorrow, signerPrivateKey, mockFEVMQueries)
	if err != nil {
		t.Fatal(err)
	}

	if jws == "" {
		t.Fatal("jws is empty")
	}

	claims, err := VerifyJWS(ctx, jws, mockFEVMQueries, false)
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

	agentAddr, target, value, signerPrivateKey, _ := setup()

	mockQuery := mock.NewFEVMQueries(t)

	privateKey, _ := crypto.GenerateKey()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, RequestClaims{
		AgentAddr:       agentAddr,
		RequesterPubKey: crypto.FromECDSAPub(&privateKey.PublicKey),
		Target:          target,
		Value:           value,
		Method:          constants.MethodBorrow,
		EpochHeight:     big.NewInt(100),
	})

	// sign with the signer's private key, different from the public key encoded in the claims
	tokenStr, err := token.SignedString(signerPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = VerifyJWS(ctx, tokenStr, mockQuery, false)
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

	agentAddr, target, value, signerPrivateKey, signerAddr := setup()

	mockQueries := mock.NewFEVMQueries(t)
	// set the chain height to be 100
	mockQueries.On("ChainHeight", ctx).Return(chainHeight, nil)
	mockQueries.On("AgentRequester", ctx, agentAddr).Return(signerAddr, nil)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, RequestClaims{
		AgentAddr:       agentAddr,
		RequesterPubKey: crypto.FromECDSAPub(&signerPrivateKey.PublicKey),
		Target:          target,
		Value:           value,
		Method:          constants.MethodBorrow,
		// encode the chain height to be
		EpochHeight: new(big.Int).Sub(chainHeight, EXPIRATION_EPOCH_BUFFER),
	})

	// sign with the signer's private key, different from the public key encoded in the claims
	tokenStr, err := token.SignedString(signerPrivateKey)
	if err != nil {
		t.Fatal(err)
	}

	_, err = VerifyJWS(ctx, tokenStr, mockQueries, false)
	if err != TokenExpiredErr {
		t.Fatal("expected stale JWS error")
	}
}

func TestJWSFromWWW(t *testing.T) {
	ctx := context.Background()

	agentAddr, target, value, signerPrivateKey, signerAddr := setup()

	mockFEVMQueries := mock.NewFEVMQueries(t)
	mockFEVMQueries.On("ChainHeight", ctx).Return(big.NewInt(100), nil)
	mockFEVMQueries.On("AgentRequester", ctx, agentAddr).Return(signerAddr, nil)

	pubKeyBytes := crypto.FromECDSAPub(&signerPrivateKey.PublicKey)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, requestClaimsStrVal{
		AgentAddr:       agentAddr,
		RequesterPubKey: pubKeyBytes,
		Target:          target,
		Value:           value.String(),
		Method:          constants.MethodBorrow,
		EpochHeight:     big.NewInt(100),
	})

	jws, err := token.SignedString(signerPrivateKey)

	if err != nil {
		t.Fatal(err)
	}

	if jws == "" {
		t.Fatal("jws is empty")
	}

	// call should throw
	_, err = VerifyJWS(ctx, jws, mockFEVMQueries, false)
	if err == nil {
		t.Fatal("error expected - jws from www should fail big.Int parsing")
	}

	claims, err := VerifyJWS(ctx, jws, mockFEVMQueries, true)
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

var ZERO_ADDR = common.Address{}

func setup() (agentAddr common.Address, target address.Address, value *big.Int, privateKey *ecdsa.PrivateKey, signerAddr common.Address) {
	privateKey, _ = crypto.GenerateKey()
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	signerAddr = crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("Signer address: ", signerAddr)

	agentAddr = common.HexToAddress("0xE8de74929076468BC59b079BDA683bc5bb813a39")
	target, _ = address.NewFromString("f01869494")
	value = big.NewInt(1000)

	return agentAddr, target, value, privateKey, signerAddr
}
