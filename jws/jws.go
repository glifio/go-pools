package jws

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/types"
	"github.com/golang-jwt/jwt/v5"
)

type RequestClaims struct {
	jwt.MapClaims
	AgentAddr       common.Address
	RequesterPubKey []byte
	Target          address.Address
	Value           *big.Int
	Method          constants.Method
	EpochHeight     *big.Int
}

var (
	InvalidAgentAddrErr = errors.New("invalid agent address")
	InvalidRequesterErr = errors.New("invalid requester")
	InvalidTokenErr     = errors.New("invalid token")
	TokenExpiredErr     = errors.New("token expired")
)

func SignJWS(ctx context.Context, agentAddr common.Address, target address.Address, value *big.Int, method constants.Method, key *ecdsa.PrivateKey, query types.FEVMQueries) (string, error) {
	epochHeight, err := query.ChainHeight(ctx)
	if err != nil {
		return "", err
	}

	pubKeyBytes := crypto.FromECDSAPub(&key.PublicKey)

	token := jwt.NewWithClaims(jwt.SigningMethodES256, RequestClaims{
		AgentAddr:       agentAddr,
		RequesterPubKey: pubKeyBytes,
		Target:          target,
		Value:           value,
		Method:          method,
		EpochHeight:     epochHeight,
	})

	return token.SignedString(key)
}

type ecdsaSignature struct {
	R, S *big.Int
}

// signatures last approx 2.5 minutes
var EXPIRATION_EPOCH_BUFFER = big.NewInt(5)

func VerifyJWS(ctx context.Context, jws string, query types.FEVMQueries) (*RequestClaims, error) {
	var jwsIssuerPubkey *ecdsa.PublicKey
	claims := &RequestClaims{}

	token, err := jwt.ParseWithClaims(jws, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		pubkey, err := crypto.UnmarshalPubkey(claims.RequesterPubKey)
		if err != nil {
			return nil, err
		}
		jwsIssuerPubkey = pubkey
		return pubkey, nil
	})

	if err != nil {
		return &RequestClaims{}, err
	}

	if !token.Valid {
		return &RequestClaims{}, InvalidTokenErr
	}

	if claims.AgentAddr.String() == "" {
		return &RequestClaims{}, InvalidAgentAddrErr
	}

	requesterFromClaims := crypto.PubkeyToAddress(*jwsIssuerPubkey)

	requester, err := query.AgentRequester(ctx, claims.AgentAddr)
	if err != nil {
		return &RequestClaims{}, err
	}

	if requester != requesterFromClaims {
		return &RequestClaims{}, InvalidRequesterErr
	}

	chainHeight, err := query.ChainHeight(ctx)
	if err != nil {
		return &RequestClaims{}, err
	}

	validEpoch := new(big.Int).Add(claims.EpochHeight, EXPIRATION_EPOCH_BUFFER)

	// make sure that the EpochHeight in the claims with the expiration buffer is gte than the current chain height
	if chainHeight.Cmp(validEpoch) > -1 {
		return &RequestClaims{}, TokenExpiredErr
	}

	return claims, nil
}
