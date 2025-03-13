package token

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func typedEIP712(
	chainId int64,
	verifyingContract common.Address,
	delegatee common.Address,
	nonce *big.Int,
	expiry *big.Int,
) apitypes.TypedData {
	return apitypes.TypedData{
		Types: apitypes.Types{
			"Delegation": []apitypes.Type{
				// the 0x address of the delegatee
				{Name: "delegatee", Type: "address"},
				// the nonce of the request
				{Name: "nonce", Type: "uint256"},
				// when this signature expires
				{Name: "expiry", Type: "uint256"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "Delegation",
		Domain: apitypes.TypedDataDomain{
			Name:              "glif.io",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainId),
			VerifyingContract: verifyingContract.String(),
		},
		Message: apitypes.TypedDataMessage{
			"delegatee": delegatee.String(),
			"nonce":     nonce.String(),
			"expiry":    expiry.String(),
		},
	}
}

func typedEIP712Digest(
	chainId int64,
	verifyingContract common.Address,
	delegatee common.Address,
	nonce *big.Int,
	expiry *big.Int,
) (common.Hash, error) {
	_, rawData, err := apitypes.TypedDataAndHash(typedEIP712(chainId, verifyingContract, delegatee, nonce, expiry))
	if err != nil {
		return common.Hash{}, err
	}

	hash := crypto.Keccak256Hash([]byte(rawData))
	hash.Bytes()

	return crypto.Keccak256Hash([]byte(rawData)), nil
}

// CreateDelegateSig generates the signature parameters for delegateBySig
// domainSeparator should be obtained from the contract's EIP712 domain
func CreateDelegateSig(chainId *big.Int, verifyingContract common.Address, wallet accounts.Wallet, account accounts.Account, passphrase string, delegatee common.Address, nonce *big.Int, expiry *big.Int) (uint8, [32]byte, [32]byte, error) {
	hashToSign, err := typedEIP712Digest(chainId.Int64(), verifyingContract, delegatee, nonce, expiry)
	if err != nil {
		return 0, [32]byte{}, [32]byte{}, err
	}

	// Sign the digest using the wallet with passphrase
	signature, err := wallet.SignDataWithPassphrase(account, passphrase, accounts.MimetypeDataWithValidator, hashToSign.Bytes())
	if err != nil {
		return 0, [32]byte{}, [32]byte{}, err
	}

	// Convert signature to v, r, s format
	var r, s [32]byte
	copy(r[:], signature[:32])
	copy(s[:], signature[32:64])
	v := signature[64] + 27 // Transform V from 0/1 to 27/28 for Ethereum

	return v, r, s, nil
}
