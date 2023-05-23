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
	"github.com/glifio/go-pools/util"
	"github.com/golang-jwt/jwt/v5"
)

type RequestClaims struct {
	jwt.MapClaims
	AgentAddr       common.Address
	RequesterPubKey []byte
	Target          address.Address
	Value           *big.Int
	ActionSelector  [4]byte
	EpochHeight     *big.Int
}

func SignJWS(ctx context.Context, agentAddr common.Address, target address.Address, value *big.Int, method constants.Method, key *ecdsa.PrivateKey, poolsSDK types.PoolsSDK) (string, error) {
	epochHeight, err := poolsSDK.Query().ChainHeight(ctx)
	if err != nil {
		return "", err
	}

	pubKeyBytes := crypto.FromECDSAPub(&key.PublicKey)

	actionSelector, err := util.MethodStrToBytes(method)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, RequestClaims{
		AgentAddr:       agentAddr,
		RequesterPubKey: pubKeyBytes,
		Target:          target,
		Value:           value,
		ActionSelector:  actionSelector,
		EpochHeight:     epochHeight,
	})

	return token.SignedString(key)
}

type ecdsaSignature struct {
	R, S *big.Int
}

// signatures last approx 60 seconds
var EXPIRATION_EPOCH_BUFFER = big.NewInt(2)

func VerifyJWS(ctx context.Context, jws string, poolsSDK types.PoolsSDK) (*RequestClaims, error) {
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
		return &RequestClaims{}, errors.New("invalid token")
	}

	if claims.AgentAddr.String() == "" {
		return &RequestClaims{}, errors.New("invalid agent address")
	}

	requesterFromClaims := crypto.PubkeyToAddress(*jwsIssuerPubkey)

	requester, err := poolsSDK.Query().AgentRequester(ctx, claims.AgentAddr)
	if err != nil {
		return &RequestClaims{}, err
	}

	if requester != requesterFromClaims {
		return &RequestClaims{}, errors.New("invalid requester")
	}

	chainHeight, err := poolsSDK.Query().ChainHeight(ctx)
	if err != nil {
		return &RequestClaims{}, err
	}

	validEpoch := new(big.Int).Add(claims.EpochHeight, EXPIRATION_EPOCH_BUFFER)

	// make sure that the EpochHeight in the claims with the expiration buffer is gte than the current chain height
	if chainHeight.Cmp(validEpoch) > -1 {
		return &RequestClaims{}, errors.New("invalid epoch height")
	}

	return claims, nil
}