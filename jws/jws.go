package jws

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/types"
	"github.com/golang-jwt/jwt/v5"
)

type RequestClaims struct {
	AgentAddr   common.Address
	Target      address.Address
	Value       *big.Int
	EpochHeight *big.Int

	jwt.MapClaims
}

func SignJWS(ctx context.Context, agentAddr common.Address, target address.Address, value *big.Int, key *ecdsa.PrivateKey, poolsSDK types.PoolsSDK) (string, error) {
	epochHeight, err := poolsSDK.Query().ChainHeight(ctx)
	if err != nil {
		return "", err
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, RequestClaims{
		AgentAddr:   agentAddr,
		Target:      target,
		Value:       value,
		EpochHeight: epochHeight,
	})

	return t.SignedString(key)
}

// func VerifyJWS(jws string, key *ecdsa.PublicKey) (*RequestClaims, error) {}
