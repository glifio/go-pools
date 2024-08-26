// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// SignedCredential is an auto generated low-level Go binding around an user-defined struct.
type SignedCredential_duplicate struct {
	Vc VerifiableCredential
	V  uint8
	R  [32]byte
	S  [32]byte
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_duplicate3 struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// AgentPoliceV2MetaData contains all meta data concerning the AgentPoliceV2 contract.
var AgentPoliceV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentStateRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCredential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationValueTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxMinersReached\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"faultySectors\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liveSectors\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTolerance\",\"type\":\"uint256\"}],\"name\":\"OverFaultySectorLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dtl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"OverLimitDTL\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"debt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"OverLimitQuota\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateChangeInProcess\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agentAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"}],\"name\":\"CredentialUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"Defaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agentID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"faultEpoch\",\"type\":\"uint256\"}],\"name\":\"FaultySectors\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"OnAdministration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_VERIFIABLE_CREDENTIAL_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"agentApproved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"agentLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowDTL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmAdministration\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmEquity\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"credentialUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"deriveStructHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distributeLiquidatedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"isValidCredential\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationDTL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMiners\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidator\",\"type\":\"uint64\"}],\"name\":\"prepareMinerForLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"registerCredentialUseBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sectorFaultyTolerancePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"setAgentDefaultDTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"agentIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"}],\"name\":\"setAgentLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowDTL\",\"type\":\"uint256\"}],\"name\":\"setBorrowDTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"_levels\",\"type\":\"uint256[10]\"}],\"name\":\"setLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidationDTL\",\"type\":\"uint256\"}],\"name\":\"setLiquidationDTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidationFee\",\"type\":\"uint256\"}],\"name\":\"setLiquidationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maxMiners\",\"type\":\"uint32\"}],\"name\":\"setMaxMiners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sectorFaultyTolerancePercent\",\"type\":\"uint256\"}],\"name\":\"setSectorFaultyTolerancePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"transferOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"validateCred\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x66038d7ea4c6800060085567016345785d8a0000600955600a805463ffffffff1916603217815561030060405269152d02c7e14af68000006101c09081526934f086f3b33b684000006101e0526969e10de76676d08000006102005269d3c21bcecceda1000000610220526a01a784379d99db42000000610240526a027b46536c66c8e3000000610260526a034f086f3b33b684000000610280526a0422ca8b0a00a4250000006102a0526a04f68ca6d8cd91c60000006102c0526a06342fd08f00f6378000006102e052620000d991600c9190620005f8565b50348015620000e6575f80fd5b5060405162004f3c38038062004f3c83398101604081905262000109916200073e565b82828187878582826200011d825f62000354565b610120526200012e81600162000354565b61014052815160208084019190912060e052815190820120610100524660a052620001bb60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c0526001600160a01b0390811661016052620001e39250831690506200038c565b90506001600160a01b0381166200020d57604051635435b28960e11b815260040160405180910390fd5b6200021881620003d6565b506200022d6001600160a01b0382166200038c565b90506001600160a01b0381166200025757604051635435b28960e11b815260040160405180910390fd5b62000262816200041d565b50506005805460ff60a01b19169055670a688906bd8b0000600655670bcbce7f1b15000060075560408051808201825260118152702927aaaa22a92faba324a62faa27a5a2a760791b60209091015251630d37324f60e11b815263aee0a13560e01b60048201526001600160a01b03821690631a6e649e90602401602060405180830381865afa158015620002f9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906200031f9190620007da565b6001600160a01b031661018052610160516200033b9062000464565b6001600160a01b03166101a05250620009a99350505050565b5f60208351101562000373576200036b8362000508565b905062000386565b8162000380848262000889565b5060ff90505b92915050565b5f80806200039a8462000553565b9150915081620003ac57509192915050565b5f80620003b98362000585565b9150915081620003cd575093949350505050565b95945050505050565b600380546001600160a01b0319169055620003fa6001600160a01b0382166200038c565b600280546001600160a01b0319166001600160a01b039290921691909117905550565b600580546001600160a01b0319169055620004416001600160a01b0382166200038c565b600480546001600160a01b0319166001600160a01b039290921691909117905550565b604080518082018252601581527f524f555445525f4d494e45525f5245474953545259000000000000000000000060209091015251630d37324f60e11b8152635324607160e01b60048201525f906001600160a01b03831690631a6e649e90602401602060405180830381865afa158015620004e2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190620003869190620007da565b5f80829050601f815111156200053e578260405163305a27a960e01b815260040162000535919062000951565b60405180910390fd5b80516200054b8262000985565b179392505050565b5f80600160401b600160a01b03831660ff60981b81036200057f57600192506001600160401b03841691505b50915091565b5f80825f526016600a60205f73fe000000000000000000000000000000000000025afa91505f51806001600160a01b031691508060a01c61ffff16905061040a8114620005d3575f92505f91505b50811580620005e357503d601614155b15620005f357505f928392509050565b915091565b82600a810192821562000634579160200282015b828111156200063457825182906001600160581b03169055916020019190600101906200060c565b506200064292915062000646565b5090565b5b8082111562000642575f815560010162000647565b634e487b7160e01b5f52604160045260245ffd5b5f5b838110156200068c57818101518382015260200162000672565b50505f910152565b5f82601f830112620006a4575f80fd5b81516001600160401b0380821115620006c157620006c16200065c565b604051601f8301601f19908116603f01168101908282118183101715620006ec57620006ec6200065c565b8160405283815286602085880101111562000705575f80fd5b6200071884602083016020890162000670565b9695505050505050565b80516001600160a01b038116811462000739575f80fd5b919050565b5f805f805f60a0868803121562000753575f80fd5b85516001600160401b03808211156200076a575f80fd5b6200077889838a0162000694565b965060208801519150808211156200078e575f80fd5b506200079d8882890162000694565b945050620007ae6040870162000722565b9250620007be6060870162000722565b9150620007ce6080870162000722565b90509295509295909350565b5f60208284031215620007eb575f80fd5b620007f68262000722565b9392505050565b600181811c908216806200081257607f821691505b6020821081036200083157634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000884575f81815260208120601f850160051c810160208610156200085f5750805b601f850160051c820191505b8181101562000880578281556001016200086b565b5050505b505050565b81516001600160401b03811115620008a557620008a56200065c565b620008bd81620008b68454620007fd565b8462000837565b602080601f831160018114620008f3575f8415620008db5750858301515b5f19600386901b1c1916600185901b17855562000880565b5f85815260208120601f198616915b82811015620009235788860151825594840194600190910190840162000902565b50858210156200094157878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b602081525f82518060208401526200097181604085016020870162000670565b601f01601f19169190910160400192915050565b8051602080830151919081101562000831575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051610180516101a0516144b362000a895f395f81816119fa01528181613050015261312e01525f8181611323015281816113fb01528181611519015281816115cf015281816116c8015281816118120152612fd701525f818161083501528181610d1d01528181610f550152818161160e015281816118510152818161213c015281816121ed015281816123bf015261270901525f612b6c01525f612b4001525f61377a01525f61375201525f6136ad01525f6136d701525f61370101526144b35ff3fe608060405234801561000f575f80fd5b50600436106102cd575f3560e01c806376048dc41161017c578063a2f3c210116100dd578063db6fccda11610093578063f2fde38b1161006e578063f2fde38b14610616578063fcad444814610629578063fef0bec814610631575f80fd5b8063db6fccda146105be578063e30c3978146105e3578063e5a8a2cb14610603575f80fd5b8063ac7e534e116100c3578063ac7e534e14610578578063b2596a6714610598578063b4e45883146105ab575f80fd5b8063a2f3c2101461055c578063a36a36301461056f575f80fd5b80638830d8c3116101325780639b83b8cd116101185780639b83b8cd146105175780639c18625f1461052a578063a0c9119a14610549575f80fd5b80638830d8c3146104e45780638da5cb5b146104f7575f80fd5b80637a018032116101625780637a018032146104ae5780638456cb59146104c157806384b0196e146104c9575f80fd5b806376048dc41461049357806379ba5097146104a6575f80fd5b8063472910d811610231578063570ca735116101e75780635e7981de116101c25780635e7981de14610464578063680af724146104775780636ec87bd01461048a575f80fd5b8063570ca7351461040e5780635c975abb1461042e5780635cf6862f14610451575f80fd5b8063516921491161021757806351692149146103c5578063529e2512146103e85780635654657a146103fb575f80fd5b8063472910d81461037a5780634ead15271461038d575f80fd5b80632b85012b116102865780633f4ba83a1161026c5780633f4ba83a1461034c578063403bb79d1461035457806340444c7c14610367575f80fd5b80632b85012b1461033b5780633727948814610343575f80fd5b8063260e63c6116102b6578063260e63c61461030c57806329605e771461031f57806329865d2a14610332575f80fd5b80630e927f74146102d15780631f2218de146102e6575b5f80fd5b6102e46102df36600461398d565b610644565b005b6102f96102f43660046139ed565b6109a2565b6040519081526020015b60405180910390f35b6102e461031a3660046139ed565b6109c7565b6102e461032d366004613a25565b6109ea565b6102f960065481565b6102f9610a39565b6102f960085481565b6102e4610a7b565b6102e4610362366004613cc4565b610a8d565b6102f9610375366004613d17565b610b7c565b6102e4610388366004613d49565b610ccd565b6103a061039b366004613d60565b610cda565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610303565b6103d86103d3366004613d49565b610d04565b6040519015158152602001610303565b6102e46103f6366004613d60565b610d18565b6102e46104093660046139ed565b610e9d565b6004546103a09073ffffffffffffffffffffffffffffffffffffffff1681565b60055474010000000000000000000000000000000000000000900460ff166103d8565b6102e461045f3660046139ed565b610f4f565b6102e4610472366004613d92565b6110de565b6102e4610485366004613d49565b611912565b6102f960075481565b6102e46104a1366004613dbc565b61191f565b6102e4611ab1565b6102e46104bc366004613e05565b611b0b565b6102e4611d64565b6104d1611d74565b6040516103039796959493929190613ebd565b6102e46104f2366004613f7a565b611dd2565b6002546103a09073ffffffffffffffffffffffffffffffffffffffff1681565b6102e4610525366004613f9d565b611e11565b6102f9610538366004613d49565b60166020525f908152604090205481565b6102e461055736600461401b565b611ea1565b6102f961056a366004613d17565b611f45565b6102f960095481565b6005546103a09073ffffffffffffffffffffffffffffffffffffffff1681565b6102f96105a6366004613d49565b611f57565b6102e46105b9366004613d49565b611f6d565b600a546105ce9063ffffffff1681565b60405163ffffffff9091168152602001610303565b6003546103a09073ffffffffffffffffffffffffffffffffffffffff1681565b6102e4610611366004614082565b611f7a565b6102e4610624366004613a25565b611f93565b6102e4612002565b6102e461063f366004613d49565b61205c565b61064c612069565b6106546120ba565b73ffffffffffffffffffffffffffffffffffffffff16630c7290f26040518163ffffffff1660e01b815260040160a060405180830381865afa15801561069c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106c091906140b9565b60800151156106fb576040517f6aefdaab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107988373ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610747573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061076b919061412b565b7fffffffff000000000000000000000000000000000000000000000000000000005f351661036285614142565b5f61085e6107b26107a9858061414d565b602001356121bf565b6107bc858061414d565b6107c46120ba565b73ffffffffffffffffffffffffffffffffffffffff1663679aefce6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561080c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610830919061412b565b6108597f0000000000000000000000000000000000000000000000000000000000000000612213565b6122f4565b5050905060065481101561089e576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff16639e4d74cc6108d98473ffffffffffffffffffffffffffffffffffffffff1661236d565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024015f604051808303815f87803b15801561093c575f80fd5b505af115801561094e573d5f803e3d5ffd5b505060405173ffffffffffffffffffffffffffffffffffffffff871681527fcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c29250602001905060405180910390a150505050565b5f600b816109b261056a85614189565b81526020019081526020015f20549050919050565b6109e733826109d984602001356121bf565b6109e16120ba565b5f6123b1565b50565b6109f2612069565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6040518060c00160405280609681526020016143e860969139604051602001610a629190614194565b6040516020818303038152906040528051906020012081565b610a83612069565b610a8b6125f1565b565b5f610a9782610cda565b82515190915073ffffffffffffffffffffffffffffffffffffffff8083169116141580610aca5750610ac88161266e565b155b80610ada57508151602001518414155b80610b0f5750815160a001517fffffffff00000000000000000000000000000000000000000000000000000000848116911614155b80610b3157508151604001514310801590610b2f57508151606001514311155b155b80610b3f5750815160200151155b15610b76576040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b5f6040518060c00160405280609681526020016143e860969139604051602001610ba69190614194565b60405160208183030381529060405280519060200120825f015183602001518460400151856060015186608001518760a001518860c001518960e00151604051602001610bf39190614194565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815282825280516020918201209083019a909a5273ffffffffffffffffffffffffffffffffffffffff909816978101979097526060870195909552608086019390935260a085019190915260c08401527fffffffff000000000000000000000000000000000000000000000000000000001660e083015267ffffffffffffffff1661010082015261012081019190915261014001604051602081830303815290604052805190602001209050919050565b610cd5612069565b600755565b5f610cfe610cea835f0151611f45565b836020015184604001518560600151612791565b92915050565b5f610d0e826121bf565b6060015192915050565b610d427f0000000000000000000000000000000000000000000000000000000000000000336127bd565b805f0151602001513373ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d93573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610db7919061412b565b14610dee576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b43600b5f610dfe845f0151611f45565b815260208082019290925260409081015f2092909255825160a081015181830151608083015160c09093015185513381529485019390935267ffffffffffffffff9092168385015292517fffffffff000000000000000000000000000000000000000000000000000000009093169290917f700450a0d828409c6031b95d040684b8a2eba879416b6960dfa8c818f245ddf4919081900360600190a350565b5f610eab82602001356121bf565b90505f610eb66120ba565b9050608083013515610ed857610ed33384848487608001356123b1565b505050565b5f610ee960e0850160c086016141a5565b67ffffffffffffffff161115610f1d57610ed333848484610f18610f1360e0850160c086016141a5565b61288c565b6123b1565b6040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610f797f0000000000000000000000000000000000000000000000000000000000000000612213565b90505f805f611009610f8e86602001356121bf565b86610f976120ba565b73ffffffffffffffffffffffffffffffffffffffff1663679aefce6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610fdf573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611003919061412b565b876122f4565b92509250925060065483111561106c576006546040517f0203b0e600000000000000000000000000000000000000000000000000000000815260048101849052602481018390526044810185905260648101919091526084015b60405180910390fd5b5f808061108161107b89614189565b88612908565b92509250925080156110d4576008546040517f4c1b369700000000000000000000000000000000000000000000000000000000815260048101859052602481018490526044810191909152606401611063565b5050505050505050565b6110e6612069565b6110ee6120ba565b73ffffffffffffffffffffffffffffffffffffffff16630c7290f26040518163ffffffff1660e01b815260040160a060405180830381865afa158015611136573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061115a91906140b9565b6080015115611195576040517f6aefdaab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff166369e25ec16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156111de573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061120291906141c0565b611238576040517f06dac5ba00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611282573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112a6919061412b565b90506112b181610d04565b156112e8576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018390527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906323b872dd906064016020604051808303815f875af115801561137e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113a291906141c0565b505f6113ac6120ba565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8083166004830152602482018690529192507f00000000000000000000000000000000000000000000000000000000000000009091169063095ea7b3906044016020604051808303815f875af1158015611443573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061146791906141c0565b506040517fc564f772000000000000000000000000000000000000000000000000000000008152600481018390526024810184905273ffffffffffffffffffffffffffffffffffffffff82169063c564f772906044015f604051808303815f87803b1580156114d4575f80fd5b505af11580156114e6573d5f803e3d5ffd5b50506040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201525f92507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1691506370a0823190602401602060405180830381865afa158015611574573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611598919061412b565b9050801561190b575f6115ab8286614206565b90505f6115c36009548361297a90919063ffffffff16565b905080831115611810577f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6116327f0000000000000000000000000000000000000000000000000000000000000000612995565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018490526044016020604051808303815f875af11580156116a1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116c591906141c0565b507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8873ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561174b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061176f9190614219565b6117798487614206565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303815f875af11580156117e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061180a91906141c0565b50611908565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb6118757f0000000000000000000000000000000000000000000000000000000000000000612995565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff9091166004820152602481018690526044016020604051808303815f875af11580156118e4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110d491906141c0565b50505b5050505050565b61191a612069565b600855565b61192d838361036284614142565b5f61193b6102f4838061414d565b1115611973576040517fcaa03ea500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b7fffffffff0000000000000000000000000000000000000000000000000000000082167f2bb9af4300000000000000000000000000000000000000000000000000000000148015611a7a5750600a546040517fe34413650000000000000000000000000000000000000000000000000000000081526004810185905263ffffffff909116907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063e344136590602401602060405180830381865afa158015611a54573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a78919061412b565b115b15610ed3576040517f189bdc8d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035473ffffffffffffffffffffffffffffffffffffffff163314611b02576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a8b33612a3b565b611b13612069565b611b1b6120ba565b73ffffffffffffffffffffffffffffffffffffffff16630c7290f26040518163ffffffff1660e01b815260040160a060405180830381865afa158015611b63573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611b8791906140b9565b6080015115611bc2576040517f6aefdaab00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611c5f8273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c0e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c32919061412b565b7fffffffff000000000000000000000000000000000000000000000000000000005f351661036284614142565b5f611c7a611c706107a9848061414d565b6107bc848061414d565b50509050600754811015611cba576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663615d92116040518163ffffffff1660e01b81526004015f604051808303815f87803b158015611cff575f80fd5b505af1158015611d11573d5f803e3d5ffd5b505060405173ffffffffffffffffffffffffffffffffffffffff861681527f0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b19250602001905060405180910390a1505050565b611d6c612069565b610a8b612aca565b5f6060805f805f6060611d85612b39565b611d8d612b65565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b611dda612b92565b600a80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000001663ffffffff92909216919091179055565b611e19612069565b6040517ff794a92300000000000000000000000000000000000000000000000000000000815267ffffffffffffffff80841660048301528216602482015273ffffffffffffffffffffffffffffffffffffffff84169063f794a923906044015f604051808303815f87803b158015611e8f575f80fd5b505af1158015611908573d5f803e3d5ffd5b611ea9612069565b828114611ee2576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f5b8381101561190b57828282818110611efe57611efe614234565b9050602002013560165f878785818110611f1a57611f1a614234565b9050602002013581526020019081526020015f20819055508080611f3d90614261565b915050611ee4565b5f610cfe611f5283610b7c565b612c09565b600c81600a8110611f66575f80fd5b0154905081565b611f75612069565b600655565b611f82612b92565b611f8f600c82600a6138f4565b5050565b611f9b612069565b611fba8173ffffffffffffffffffffffffffffffffffffffff1661236d565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60055473ffffffffffffffffffffffffffffffffffffffff163314612053576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a8b33612c50565b612064612069565b600955565b60025473ffffffffffffffffffffffffffffffffffffffff163314610a8b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b604080518082018252601481527f524f555445525f494e46494e4954595f504f4f4c000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f392104570000000000000000000000000000000000000000000000000000000060048201525f907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690631a6e649e90602401602060405180830381865afa158015612196573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121ba9190614219565b905090565b6121e860405180608001604052805f81526020015f81526020015f81526020015f151581525090565b610cfe7f0000000000000000000000000000000000000000000000000000000000000000835f612cdf565b604080518082018252601281527f524f555445525f435245445f5041525345520000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527fd26df3b70000000000000000000000000000000000000000000000000000000060048201525f9073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024015b602060405180830381865afa1580156122d0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cfe9190614219565b5f805f6123018786612da4565b91506123168461231088614189565b90612dc3565b9050815f03612327575f9250612363565b805f03612356577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9250612363565b6123608282612e59565b92505b9450945094915050565b5f805f61237984612e6d565b915091508161238a57509192915050565b5f8061239583612eb9565b91509150816123a8575093949350505050565b95945050505050565b6123b9612f36565b5f6123e37f0000000000000000000000000000000000000000000000000000000000000000612213565b60208501519091505f8190036123fa57505061190b565b5f805f612475888a8973ffffffffffffffffffffffffffffffffffffffff1663679aefce6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561244b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061246f919061412b565b886122f4565b92509250925061248b8a8a602001358389612f8b565b6006548311156124e3576006546040517f0203b0e60000000000000000000000000000000000000000000000000000000081526004810184905260248101839052604481018590526064810191909152608401611063565b5f80806124f86124f28d614189565b89612908565b925092509250801561254b576008546040517f4c1b369700000000000000000000000000000000000000000000000000000000815260048101859052602481018490526044810191909152606401611063565b600c60165f8e6020013581526020019081526020015f2054600a811061257357612573614234565b01548511156125e25784600c60165f8f6020013581526020019081526020015f2054600a81106125a5576125a5614234565b01546040517f9598576400000000000000000000000000000000000000000000000000000000815260048101929092526024820152604401611063565b50505050505050505050505050565b6125f9613217565b600580547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b604080518082018252601081527f524f555445525f56435f49535355455200000000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f1603a47a0000000000000000000000000000000000000000000000000000000060048201525f9073ffffffffffffffffffffffffffffffffffffffff838116917f000000000000000000000000000000000000000000000000000000000000000090911690631a6e649e90602401602060405180830381865afa158015612750573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906127749190614219565b73ffffffffffffffffffffffffffffffffffffffff161492915050565b5f805f806127a18888888861326b565b9250925092506127b1828261335a565b50909695505050505050565b6127c68261345d565b6040517f1ffbb06400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301529190911690631ffbb06490602401602060405180830381865afa158015612832573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061285691906141c0565b611f8f576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517fff0000000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffff00000000000000000000000000000000000000000000000060c083901b16602c8201525f906034016040516020818303038152906040526128fe90614298565b60601c3192915050565b5f80806129158585613503565b9250612921858561355a565b91508115801561292f575082155b1561293b57505f612973565b8115801561294857505f83115b1561295557506001612973565b6008546129628484612e59565b111561297057506001612973565b505f5b9250925092565b5f61298e8383670de0b6b3a76400006135b1565b9392505050565b604080518082018252600f81527f524f555445525f54524541535552590000000000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527fa1858d5f0000000000000000000000000000000000000000000000000000000060048201525f9073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016122b5565b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055612a8273ffffffffffffffffffffffffffffffffffffffff821661236d565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b612ad2612f36565b600580547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586126443390565b60606121ba7f00000000000000000000000000000000000000000000000000000000000000005f6135eb565b60606121ba7f000000000000000000000000000000000000000000000000000000000000000060016135eb565b60045473ffffffffffffffffffffffffffffffffffffffff163314801590612bd2575060025473ffffffffffffffffffffffffffffffffffffffff163314155b15610a8b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f610cfe612c15613694565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b600580547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055612c9773ffffffffffffffffffffffffffffffffffffffff821661236d565b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b612d0860405180608001604052805f81526020015f81526020015f81526020015f151581525090565b6040517f6361f6de000000000000000000000000000000000000000000000000000000008152600481018490526024810183905273ffffffffffffffffffffffffffffffffffffffff851690636361f6de90604401608060405180830381865afa158015612d78573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d9c91906142e7565b949350505050565b5f80612db084846137ca565b509050808460200151612d9c919061432f565b60e08201516040517f1b2b5fad0000000000000000000000000000000000000000000000000000000081525f9173ffffffffffffffffffffffffffffffffffffffff841691631b2b5fad91612e1a91600401614342565b602060405180830381865afa158015612e35573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061298e919061412b565b5f61298e83670de0b6b3a7640000846135b1565b5f8073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103612eb3576001925067ffffffffffffffff841691505b50915091565b5f80825f526016600a60205f73fe000000000000000000000000000000000000025afa91505f518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a8114612f13575f92505f91505b50811580612f2257503d601614155b15612f3157505f928392509050565b915091565b60055474010000000000000000000000000000000000000000900460ff1615610a8b576040517fd93c066500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff808616600483018190525f929031917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa15801561301c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613040919061412b565b61304a919061432f565b90505f807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e3441365876040518263ffffffff1660e01b81526004016130a991815260200190565b602060405180830381865afa1580156130c4573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130e8919061412b565b90505f5b818110156131ca576040517f3d874d9f00000000000000000000000000000000000000000000000000000000815260048101889052602481018290526131ac907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1690633d874d9f90604401602060405180830381865afa158015613188573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f139190614354565b6131b6908461432f565b9250806131c281614261565b9150506130ec565b506131d5828461432f565b6131df858761432f565b1115611908576040517fb750281700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60055474010000000000000000000000000000000000000000900460ff16610a8b576040517f8dfc202b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156132a457505f91506003905082612363565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156132f5573d5f803e3d5ffd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0015191505073ffffffffffffffffffffffffffffffffffffffff811661334b57505f925060019150829050612363565b975f9750879650945050505050565b5f82600381111561336d5761336d61436f565b03613376575050565b600182600381111561338a5761338a61436f565b036133c1576040517ff645eedf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60028260038111156133d5576133d561436f565b0361340f576040517ffce698f700000000000000000000000000000000000000000000000000000000815260048101829052602401611063565b60038260038111156134235761342361436f565b03611f8f576040517fd78bce0c00000000000000000000000000000000000000000000000000000000815260048101829052602401611063565b604080518082018252601481527f524f555445525f4147454e545f464143544f5259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f29f616da0000000000000000000000000000000000000000000000000000000060048201525f9073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016122b5565b60e08201516040517f07124c060000000000000000000000000000000000000000000000000000000081525f9173ffffffffffffffffffffffffffffffffffffffff8416916307124c0691612e1a91600401614342565b60e08201516040517f402ed8cb0000000000000000000000000000000000000000000000000000000081525f9173ffffffffffffffffffffffffffffffffffffffff84169163402ed8cb91612e1a91600401614342565b5f827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04841183021582026135e4575f80fd5b5091020490565b606060ff8314613605576135fe836137e3565b9050610cfe565b8180546136119061439c565b80601f016020809104026020016040519081016040528092919081815260200182805461363d9061439c565b80156136885780601f1061365f57610100808354040283529160200191613688565b820191905f5260205f20905b81548152906001019060200180831161366b57829003601f168201915b50505050509050610cfe565b5f3073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156136f957507f000000000000000000000000000000000000000000000000000000000000000046145b1561372357507f000000000000000000000000000000000000000000000000000000000000000090565b6121ba604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f806137d7848443613820565b915091505b9250929050565b60605f6137ef83613856565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f805f8560400151846138339190614206565b905061383f8686613896565b915061384b82826138a3565b925050935093915050565b5f60ff8216601f811115610cfe576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208201515f9061298e90835b5f61298e8383670de0b6b3a76400005f827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04841183021582026138e5575f80fd5b50910281810615159190040190565b82600a8101928215613922579160200282015b82811115613922578235825591602001919060010190613907565b5061392e929150613932565b5090565b5b8082111561392e575f8155600101613933565b73ffffffffffffffffffffffffffffffffffffffff811681146109e7575f80fd5b803561397281613946565b919050565b5f60808284031215613987575f80fd5b50919050565b5f805f6060848603121561399f575f80fd5b83356139aa81613946565b9250602084013567ffffffffffffffff8111156139c5575f80fd5b6139d186828701613977565b92505060408401356139e281613946565b809150509250925092565b5f602082840312156139fd575f80fd5b813567ffffffffffffffff811115613a13575f80fd5b8201610100818503121561298e575f80fd5b5f60208284031215613a35575f80fd5b813561298e81613946565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114613972575f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b604051610100810167ffffffffffffffff81118282101715613ac057613ac0613a6f565b60405290565b6040516080810167ffffffffffffffff81118282101715613ac057613ac0613a6f565b67ffffffffffffffff811681146109e7575f80fd5b803561397281613ae9565b5f82601f830112613b18575f80fd5b813567ffffffffffffffff80821115613b3357613b33613a6f565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715613b7957613b79613a6f565b81604052838152866020858801011115613b91575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f6101008284031215613bc1575f80fd5b613bc9613a9c565b9050613bd482613967565b815260208201356020820152604082013560408201526060820135606082015260808201356080820152613c0a60a08301613a40565b60a0820152613c1b60c08301613afe565b60c082015260e082013567ffffffffffffffff811115613c39575f80fd5b613c4584828501613b09565b60e08301525092915050565b5f60808284031215613c61575f80fd5b613c69613ac6565b9050813567ffffffffffffffff811115613c81575f80fd5b613c8d84828501613bb0565b825250602082013560ff81168114613ca3575f80fd5b80602083015250604082013560408201526060820135606082015292915050565b5f805f60608486031215613cd6575f80fd5b83359250613ce660208501613a40565b9150604084013567ffffffffffffffff811115613d01575f80fd5b613d0d86828701613c51565b9150509250925092565b5f60208284031215613d27575f80fd5b813567ffffffffffffffff811115613d3d575f80fd5b612d9c84828501613bb0565b5f60208284031215613d59575f80fd5b5035919050565b5f60208284031215613d70575f80fd5b813567ffffffffffffffff811115613d86575f80fd5b612d9c84828501613c51565b5f8060408385031215613da3575f80fd5b8235613dae81613946565b946020939093013593505050565b5f805f60608486031215613dce575f80fd5b83359250613dde60208501613a40565b9150604084013567ffffffffffffffff811115613df9575f80fd5b613d0d86828701613977565b5f8060408385031215613e16575f80fd5b8235613e2181613946565b9150602083013567ffffffffffffffff811115613e3c575f80fd5b613e4885828601613977565b9150509250929050565b5f5b83811015613e6c578181015183820152602001613e54565b50505f910152565b5f8151808452613e8b816020860160208601613e52565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b7fff00000000000000000000000000000000000000000000000000000000000000881681525f602060e081840152613ef860e084018a613e74565b8381036040850152613f0a818a613e74565b6060850189905273ffffffffffffffffffffffffffffffffffffffff8816608086015260a0850187905284810360c086015285518082528387019250908301905f5b81811015613f6857835183529284019291840191600101613f4c565b50909c9b505050505050505050505050565b5f60208284031215613f8a575f80fd5b813563ffffffff8116811461298e575f80fd5b5f805f60608486031215613faf575f80fd5b8335613fba81613946565b92506020840135613fca81613ae9565b915060408401356139e281613ae9565b5f8083601f840112613fea575f80fd5b50813567ffffffffffffffff811115614001575f80fd5b6020830191508360208260051b85010111156137dc575f80fd5b5f805f806040858703121561402e575f80fd5b843567ffffffffffffffff80821115614045575f80fd5b61405188838901613fda565b90965094506020870135915080821115614069575f80fd5b5061407687828801613fda565b95989497509550505050565b5f610140808385031215614094575f80fd5b8381840111156140a2575f80fd5b509092915050565b80518015158114613972575f80fd5b5f60a082840312156140c9575f80fd5b60405160a0810181811067ffffffffffffffff821117156140ec576140ec613a6f565b80604052508251815260208301516020820152604083015160408201526060830151606082015261411f608084016140aa565b60808201529392505050565b5f6020828403121561413b575f80fd5b5051919050565b5f610cfe3683613c51565b5f82357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0183360301811261417f575f80fd5b9190910192915050565b5f610cfe3683613bb0565b5f825161417f818460208701613e52565b5f602082840312156141b5575f80fd5b813561298e81613ae9565b5f602082840312156141d0575f80fd5b61298e826140aa565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b81810381811115610cfe57610cfe6141d9565b5f60208284031215614229575f80fd5b815161298e81613946565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614291576142916141d9565b5060010190565b5f815160208301517fffffffffffffffffffffffffffffffffffffffff000000000000000000000000808216935060148310156142df5780818460140360031b1b83161693505b505050919050565b5f608082840312156142f7575f80fd5b6142ff613ac6565b825181526020830151602082015260408301516040820152614323606084016140aa565b60608201529392505050565b80820180821115610cfe57610cfe6141d9565b602081525f61298e6020830184613e74565b5f60208284031215614364575f80fd5b815161298e81613ae9565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b600181811c908216806143b057607f821691505b602082108103613987577f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffdfe56657269666961626c6543726564656e7469616c2861646472657373206973737565722c75696e74323536207375626a6563742c75696e743235362065706f63684973737565642c75696e743235362065706f636856616c6964556e74696c2c75696e743235362076616c75652c62797465733420616374696f6e2c75696e743634207461726765742c627974657320636c61696d29a2646970667358221220d04c74663c3dd6783588683125aae3c800fd7bc5222886d0873f60f173789e4864736f6c63430008150033",
}

// AgentPoliceV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentPoliceV2MetaData.ABI instead.
var AgentPoliceV2ABI = AgentPoliceV2MetaData.ABI

// AgentPoliceV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AgentPoliceV2MetaData.Bin instead.
var AgentPoliceV2Bin = AgentPoliceV2MetaData.Bin

// DeployAgentPoliceV2 deploys a new Ethereum contract, binding an instance of AgentPoliceV2 to it.
func DeployAgentPoliceV2(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _version string, _owner common.Address, _operator common.Address, _router common.Address) (common.Address, *types.Transaction, *AgentPoliceV2, error) {
	parsed, err := AgentPoliceV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AgentPoliceV2Bin), backend, _name, _version, _owner, _operator, _router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AgentPoliceV2{AgentPoliceV2Caller: AgentPoliceV2Caller{contract: contract}, AgentPoliceV2Transactor: AgentPoliceV2Transactor{contract: contract}, AgentPoliceV2Filterer: AgentPoliceV2Filterer{contract: contract}}, nil
}

// AgentPoliceV2 is an auto generated Go binding around an Ethereum contract.
type AgentPoliceV2 struct {
	AgentPoliceV2Caller     // Read-only binding to the contract
	AgentPoliceV2Transactor // Write-only binding to the contract
	AgentPoliceV2Filterer   // Log filterer for contract events
}

// AgentPoliceV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type AgentPoliceV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentPoliceV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentPoliceV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentPoliceV2Session struct {
	Contract     *AgentPoliceV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentPoliceV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentPoliceV2CallerSession struct {
	Contract *AgentPoliceV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AgentPoliceV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentPoliceV2TransactorSession struct {
	Contract     *AgentPoliceV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AgentPoliceV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type AgentPoliceV2Raw struct {
	Contract *AgentPoliceV2 // Generic contract binding to access the raw methods on
}

// AgentPoliceV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentPoliceV2CallerRaw struct {
	Contract *AgentPoliceV2Caller // Generic read-only contract binding to access the raw methods on
}

// AgentPoliceV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentPoliceV2TransactorRaw struct {
	Contract *AgentPoliceV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentPoliceV2 creates a new instance of AgentPoliceV2, bound to a specific deployed contract.
func NewAgentPoliceV2(address common.Address, backend bind.ContractBackend) (*AgentPoliceV2, error) {
	contract, err := bindAgentPoliceV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2{AgentPoliceV2Caller: AgentPoliceV2Caller{contract: contract}, AgentPoliceV2Transactor: AgentPoliceV2Transactor{contract: contract}, AgentPoliceV2Filterer: AgentPoliceV2Filterer{contract: contract}}, nil
}

// NewAgentPoliceV2Caller creates a new read-only instance of AgentPoliceV2, bound to a specific deployed contract.
func NewAgentPoliceV2Caller(address common.Address, caller bind.ContractCaller) (*AgentPoliceV2Caller, error) {
	contract, err := bindAgentPoliceV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2Caller{contract: contract}, nil
}

// NewAgentPoliceV2Transactor creates a new write-only instance of AgentPoliceV2, bound to a specific deployed contract.
func NewAgentPoliceV2Transactor(address common.Address, transactor bind.ContractTransactor) (*AgentPoliceV2Transactor, error) {
	contract, err := bindAgentPoliceV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2Transactor{contract: contract}, nil
}

// NewAgentPoliceV2Filterer creates a new log filterer instance of AgentPoliceV2, bound to a specific deployed contract.
func NewAgentPoliceV2Filterer(address common.Address, filterer bind.ContractFilterer) (*AgentPoliceV2Filterer, error) {
	contract, err := bindAgentPoliceV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2Filterer{contract: contract}, nil
}

// bindAgentPoliceV2 binds a generic wrapper to an already deployed contract.
func bindAgentPoliceV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentPoliceV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentPoliceV2 *AgentPoliceV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentPoliceV2.Contract.AgentPoliceV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentPoliceV2 *AgentPoliceV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.AgentPoliceV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentPoliceV2 *AgentPoliceV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.AgentPoliceV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentPoliceV2 *AgentPoliceV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentPoliceV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentPoliceV2 *AgentPoliceV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentPoliceV2 *AgentPoliceV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.contract.Transact(opts, method, params...)
}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2Caller) VERIFIABLECREDENTIALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "_VERIFIABLE_CREDENTIAL_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2Session) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _AgentPoliceV2.Contract.VERIFIABLECREDENTIALTYPEHASH(&_AgentPoliceV2.CallOpts)
}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _AgentPoliceV2.Contract.VERIFIABLECREDENTIALTYPEHASH(&_AgentPoliceV2.CallOpts)
}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Caller) AccountLevel(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "accountLevel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Session) AccountLevel(arg0 *big.Int) (*big.Int, error) {
	return _AgentPoliceV2.Contract.AccountLevel(&_AgentPoliceV2.CallOpts, arg0)
}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) AccountLevel(arg0 *big.Int) (*big.Int, error) {
	return _AgentPoliceV2.Contract.AccountLevel(&_AgentPoliceV2.CallOpts, arg0)
}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Caller) AgentApproved(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "agentApproved", vc)

	if err != nil {
		return err
	}

	return err

}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) AgentApproved(vc VerifiableCredential) error {
	return _AgentPoliceV2.Contract.AgentApproved(&_AgentPoliceV2.CallOpts, vc)
}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) AgentApproved(vc VerifiableCredential) error {
	return _AgentPoliceV2.Contract.AgentApproved(&_AgentPoliceV2.CallOpts, vc)
}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPoliceV2 *AgentPoliceV2Caller) AgentLiquidated(opts *bind.CallOpts, agentID *big.Int) (bool, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "agentLiquidated", agentID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPoliceV2 *AgentPoliceV2Session) AgentLiquidated(agentID *big.Int) (bool, error) {
	return _AgentPoliceV2.Contract.AgentLiquidated(&_AgentPoliceV2.CallOpts, agentID)
}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) AgentLiquidated(agentID *big.Int) (bool, error) {
	return _AgentPoliceV2.Contract.AgentLiquidated(&_AgentPoliceV2.CallOpts, agentID)
}

// BorrowDTL is a free data retrieval call binding the contract method 0x29865d2a.
//
// Solidity: function borrowDTL() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Caller) BorrowDTL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "borrowDTL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowDTL is a free data retrieval call binding the contract method 0x29865d2a.
//
// Solidity: function borrowDTL() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Session) BorrowDTL() (*big.Int, error) {
	return _AgentPoliceV2.Contract.BorrowDTL(&_AgentPoliceV2.CallOpts)
}

// BorrowDTL is a free data retrieval call binding the contract method 0x29865d2a.
//
// Solidity: function borrowDTL() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) BorrowDTL() (*big.Int, error) {
	return _AgentPoliceV2.Contract.BorrowDTL(&_AgentPoliceV2.CallOpts)
}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Caller) ConfirmRmAdministration(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "confirmRmAdministration", vc)

	if err != nil {
		return err
	}

	return err

}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) ConfirmRmAdministration(vc VerifiableCredential) error {
	return _AgentPoliceV2.Contract.ConfirmRmAdministration(&_AgentPoliceV2.CallOpts, vc)
}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) ConfirmRmAdministration(vc VerifiableCredential) error {
	return _AgentPoliceV2.Contract.ConfirmRmAdministration(&_AgentPoliceV2.CallOpts, vc)
}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Caller) ConfirmRmEquity(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "confirmRmEquity", vc)

	if err != nil {
		return err
	}

	return err

}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) ConfirmRmEquity(vc VerifiableCredential) error {
	return _AgentPoliceV2.Contract.ConfirmRmEquity(&_AgentPoliceV2.CallOpts, vc)
}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) ConfirmRmEquity(vc VerifiableCredential) error {
	return _AgentPoliceV2.Contract.ConfirmRmEquity(&_AgentPoliceV2.CallOpts, vc)
}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Caller) CredentialUsed(opts *bind.CallOpts, vc VerifiableCredential) (*big.Int, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "credentialUsed", vc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Session) CredentialUsed(vc VerifiableCredential) (*big.Int, error) {
	return _AgentPoliceV2.Contract.CredentialUsed(&_AgentPoliceV2.CallOpts, vc)
}

// CredentialUsed is a free data retrieval call binding the contract method 0x1f2218de.
//
// Solidity: function credentialUsed((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) CredentialUsed(vc VerifiableCredential) (*big.Int, error) {
	return _AgentPoliceV2.Contract.CredentialUsed(&_AgentPoliceV2.CallOpts, vc)
}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2Caller) DeriveStructHash(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "deriveStructHash", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2Session) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPoliceV2.Contract.DeriveStructHash(&_AgentPoliceV2.CallOpts, vc)
}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPoliceV2.Contract.DeriveStructHash(&_AgentPoliceV2.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2Caller) Digest(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "digest", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2Session) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPoliceV2.Contract.Digest(&_AgentPoliceV2.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPoliceV2.Contract.Digest(&_AgentPoliceV2.CallOpts, vc)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentPoliceV2 *AgentPoliceV2Caller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentPoliceV2 *AgentPoliceV2Session) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentPoliceV2.Contract.Eip712Domain(&_AgentPoliceV2.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _AgentPoliceV2.Contract.Eip712Domain(&_AgentPoliceV2.CallOpts)
}

// IsValidCredential is a free data retrieval call binding the contract method 0x76048dc4.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Caller) IsValidCredential(opts *bind.CallOpts, agent *big.Int, action [4]byte, sc SignedCredential) error {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "isValidCredential", agent, action, sc)

	if err != nil {
		return err
	}

	return err

}

// IsValidCredential is a free data retrieval call binding the contract method 0x76048dc4.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) error {
	return _AgentPoliceV2.Contract.IsValidCredential(&_AgentPoliceV2.CallOpts, agent, action, sc)
}

// IsValidCredential is a free data retrieval call binding the contract method 0x76048dc4.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) error {
	return _AgentPoliceV2.Contract.IsValidCredential(&_AgentPoliceV2.CallOpts, agent, action, sc)
}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Caller) Levels(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "levels", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Session) Levels(arg0 *big.Int) (*big.Int, error) {
	return _AgentPoliceV2.Contract.Levels(&_AgentPoliceV2.CallOpts, arg0)
}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) Levels(arg0 *big.Int) (*big.Int, error) {
	return _AgentPoliceV2.Contract.Levels(&_AgentPoliceV2.CallOpts, arg0)
}

// LiquidationDTL is a free data retrieval call binding the contract method 0x6ec87bd0.
//
// Solidity: function liquidationDTL() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Caller) LiquidationDTL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "liquidationDTL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationDTL is a free data retrieval call binding the contract method 0x6ec87bd0.
//
// Solidity: function liquidationDTL() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Session) LiquidationDTL() (*big.Int, error) {
	return _AgentPoliceV2.Contract.LiquidationDTL(&_AgentPoliceV2.CallOpts)
}

// LiquidationDTL is a free data retrieval call binding the contract method 0x6ec87bd0.
//
// Solidity: function liquidationDTL() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) LiquidationDTL() (*big.Int, error) {
	return _AgentPoliceV2.Contract.LiquidationDTL(&_AgentPoliceV2.CallOpts)
}

// LiquidationFee is a free data retrieval call binding the contract method 0xa36a3630.
//
// Solidity: function liquidationFee() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Caller) LiquidationFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "liquidationFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationFee is a free data retrieval call binding the contract method 0xa36a3630.
//
// Solidity: function liquidationFee() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Session) LiquidationFee() (*big.Int, error) {
	return _AgentPoliceV2.Contract.LiquidationFee(&_AgentPoliceV2.CallOpts)
}

// LiquidationFee is a free data retrieval call binding the contract method 0xa36a3630.
//
// Solidity: function liquidationFee() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) LiquidationFee() (*big.Int, error) {
	return _AgentPoliceV2.Contract.LiquidationFee(&_AgentPoliceV2.CallOpts)
}

// MaxMiners is a free data retrieval call binding the contract method 0xdb6fccda.
//
// Solidity: function maxMiners() view returns(uint32)
func (_AgentPoliceV2 *AgentPoliceV2Caller) MaxMiners(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "maxMiners")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxMiners is a free data retrieval call binding the contract method 0xdb6fccda.
//
// Solidity: function maxMiners() view returns(uint32)
func (_AgentPoliceV2 *AgentPoliceV2Session) MaxMiners() (uint32, error) {
	return _AgentPoliceV2.Contract.MaxMiners(&_AgentPoliceV2.CallOpts)
}

// MaxMiners is a free data retrieval call binding the contract method 0xdb6fccda.
//
// Solidity: function maxMiners() view returns(uint32)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) MaxMiners() (uint32, error) {
	return _AgentPoliceV2.Contract.MaxMiners(&_AgentPoliceV2.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Caller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Session) Operator() (common.Address, error) {
	return _AgentPoliceV2.Contract.Operator(&_AgentPoliceV2.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) Operator() (common.Address, error) {
	return _AgentPoliceV2.Contract.Operator(&_AgentPoliceV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Session) Owner() (common.Address, error) {
	return _AgentPoliceV2.Contract.Owner(&_AgentPoliceV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) Owner() (common.Address, error) {
	return _AgentPoliceV2.Contract.Owner(&_AgentPoliceV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPoliceV2 *AgentPoliceV2Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPoliceV2 *AgentPoliceV2Session) Paused() (bool, error) {
	return _AgentPoliceV2.Contract.Paused(&_AgentPoliceV2.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) Paused() (bool, error) {
	return _AgentPoliceV2.Contract.Paused(&_AgentPoliceV2.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Caller) PendingOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "pendingOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Session) PendingOperator() (common.Address, error) {
	return _AgentPoliceV2.Contract.PendingOperator(&_AgentPoliceV2.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) PendingOperator() (common.Address, error) {
	return _AgentPoliceV2.Contract.PendingOperator(&_AgentPoliceV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Session) PendingOwner() (common.Address, error) {
	return _AgentPoliceV2.Contract.PendingOwner(&_AgentPoliceV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) PendingOwner() (common.Address, error) {
	return _AgentPoliceV2.Contract.PendingOwner(&_AgentPoliceV2.CallOpts)
}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Caller) Recover(opts *bind.CallOpts, sc SignedCredential) (common.Address, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "recover", sc)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2Session) Recover(sc SignedCredential) (common.Address, error) {
	return _AgentPoliceV2.Contract.Recover(&_AgentPoliceV2.CallOpts, sc)
}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) Recover(sc SignedCredential) (common.Address, error) {
	return _AgentPoliceV2.Contract.Recover(&_AgentPoliceV2.CallOpts, sc)
}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Caller) SectorFaultyTolerancePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "sectorFaultyTolerancePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2Session) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _AgentPoliceV2.Contract.SectorFaultyTolerancePercent(&_AgentPoliceV2.CallOpts)
}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _AgentPoliceV2.Contract.SectorFaultyTolerancePercent(&_AgentPoliceV2.CallOpts)
}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Caller) ValidateCred(opts *bind.CallOpts, agent *big.Int, selector [4]byte, sc SignedCredential) error {
	var out []interface{}
	err := _AgentPoliceV2.contract.Call(opts, &out, "validateCred", agent, selector, sc)

	if err != nil {
		return err
	}

	return err

}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) error {
	return _AgentPoliceV2.Contract.ValidateCred(&_AgentPoliceV2.CallOpts, agent, selector, sc)
}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPoliceV2 *AgentPoliceV2CallerSession) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) error {
	return _AgentPoliceV2.Contract.ValidateCred(&_AgentPoliceV2.CallOpts, agent, selector, sc)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) AcceptOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "acceptOperator")
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) AcceptOperator() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.AcceptOperator(&_AgentPoliceV2.TransactOpts)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) AcceptOperator() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.AcceptOperator(&_AgentPoliceV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.AcceptOwnership(&_AgentPoliceV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.AcceptOwnership(&_AgentPoliceV2.TransactOpts)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) DistributeLiquidatedFunds(opts *bind.TransactOpts, agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "distributeLiquidatedFunds", agent, amount)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) DistributeLiquidatedFunds(agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.DistributeLiquidatedFunds(&_AgentPoliceV2.TransactOpts, agent, amount)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) DistributeLiquidatedFunds(agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.DistributeLiquidatedFunds(&_AgentPoliceV2.TransactOpts, agent, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) Pause() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.Pause(&_AgentPoliceV2.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) Pause() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.Pause(&_AgentPoliceV2.TransactOpts)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) PrepareMinerForLiquidation(opts *bind.TransactOpts, agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "prepareMinerForLiquidation", agent, miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) PrepareMinerForLiquidation(agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.PrepareMinerForLiquidation(&_AgentPoliceV2.TransactOpts, agent, miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) PrepareMinerForLiquidation(agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.PrepareMinerForLiquidation(&_AgentPoliceV2.TransactOpts, agent, miner, liquidator)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x0e927f74.
//
// Solidity: function putAgentOnAdministration(address agent, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc, address administration) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) PutAgentOnAdministration(opts *bind.TransactOpts, agent common.Address, sc SignedCredential, administration common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "putAgentOnAdministration", agent, sc, administration)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x0e927f74.
//
// Solidity: function putAgentOnAdministration(address agent, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc, address administration) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) PutAgentOnAdministration(agent common.Address, sc SignedCredential, administration common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.PutAgentOnAdministration(&_AgentPoliceV2.TransactOpts, agent, sc, administration)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x0e927f74.
//
// Solidity: function putAgentOnAdministration(address agent, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc, address administration) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) PutAgentOnAdministration(agent common.Address, sc SignedCredential, administration common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.PutAgentOnAdministration(&_AgentPoliceV2.TransactOpts, agent, sc, administration)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0x529e2512.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) RegisterCredentialUseBlock(opts *bind.TransactOpts, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "registerCredentialUseBlock", sc)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0x529e2512.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) RegisterCredentialUseBlock(sc SignedCredential) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.RegisterCredentialUseBlock(&_AgentPoliceV2.TransactOpts, sc)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0x529e2512.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) RegisterCredentialUseBlock(sc SignedCredential) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.RegisterCredentialUseBlock(&_AgentPoliceV2.TransactOpts, sc)
}

// SetAgentDefaultDTL is a paid mutator transaction binding the contract method 0x7a018032.
//
// Solidity: function setAgentDefaultDTL(address agent, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetAgentDefaultDTL(opts *bind.TransactOpts, agent common.Address, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setAgentDefaultDTL", agent, sc)
}

// SetAgentDefaultDTL is a paid mutator transaction binding the contract method 0x7a018032.
//
// Solidity: function setAgentDefaultDTL(address agent, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetAgentDefaultDTL(agent common.Address, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetAgentDefaultDTL(&_AgentPoliceV2.TransactOpts, agent, sc)
}

// SetAgentDefaultDTL is a paid mutator transaction binding the contract method 0x7a018032.
//
// Solidity: function setAgentDefaultDTL(address agent, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetAgentDefaultDTL(agent common.Address, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetAgentDefaultDTL(&_AgentPoliceV2.TransactOpts, agent, sc)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetAgentLevels(opts *bind.TransactOpts, agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setAgentLevels", agentIDs, level)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetAgentLevels(agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetAgentLevels(&_AgentPoliceV2.TransactOpts, agentIDs, level)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetAgentLevels(agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetAgentLevels(&_AgentPoliceV2.TransactOpts, agentIDs, level)
}

// SetBorrowDTL is a paid mutator transaction binding the contract method 0xb4e45883.
//
// Solidity: function setBorrowDTL(uint256 _borrowDTL) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetBorrowDTL(opts *bind.TransactOpts, _borrowDTL *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setBorrowDTL", _borrowDTL)
}

// SetBorrowDTL is a paid mutator transaction binding the contract method 0xb4e45883.
//
// Solidity: function setBorrowDTL(uint256 _borrowDTL) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetBorrowDTL(_borrowDTL *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetBorrowDTL(&_AgentPoliceV2.TransactOpts, _borrowDTL)
}

// SetBorrowDTL is a paid mutator transaction binding the contract method 0xb4e45883.
//
// Solidity: function setBorrowDTL(uint256 _borrowDTL) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetBorrowDTL(_borrowDTL *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetBorrowDTL(&_AgentPoliceV2.TransactOpts, _borrowDTL)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetLevels(opts *bind.TransactOpts, _levels [10]*big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setLevels", _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetLevels(_levels [10]*big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetLevels(&_AgentPoliceV2.TransactOpts, _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetLevels(_levels [10]*big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetLevels(&_AgentPoliceV2.TransactOpts, _levels)
}

// SetLiquidationDTL is a paid mutator transaction binding the contract method 0x472910d8.
//
// Solidity: function setLiquidationDTL(uint256 _liquidationDTL) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetLiquidationDTL(opts *bind.TransactOpts, _liquidationDTL *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setLiquidationDTL", _liquidationDTL)
}

// SetLiquidationDTL is a paid mutator transaction binding the contract method 0x472910d8.
//
// Solidity: function setLiquidationDTL(uint256 _liquidationDTL) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetLiquidationDTL(_liquidationDTL *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetLiquidationDTL(&_AgentPoliceV2.TransactOpts, _liquidationDTL)
}

// SetLiquidationDTL is a paid mutator transaction binding the contract method 0x472910d8.
//
// Solidity: function setLiquidationDTL(uint256 _liquidationDTL) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetLiquidationDTL(_liquidationDTL *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetLiquidationDTL(&_AgentPoliceV2.TransactOpts, _liquidationDTL)
}

// SetLiquidationFee is a paid mutator transaction binding the contract method 0xfef0bec8.
//
// Solidity: function setLiquidationFee(uint256 _liquidationFee) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetLiquidationFee(opts *bind.TransactOpts, _liquidationFee *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setLiquidationFee", _liquidationFee)
}

// SetLiquidationFee is a paid mutator transaction binding the contract method 0xfef0bec8.
//
// Solidity: function setLiquidationFee(uint256 _liquidationFee) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetLiquidationFee(_liquidationFee *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetLiquidationFee(&_AgentPoliceV2.TransactOpts, _liquidationFee)
}

// SetLiquidationFee is a paid mutator transaction binding the contract method 0xfef0bec8.
//
// Solidity: function setLiquidationFee(uint256 _liquidationFee) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetLiquidationFee(_liquidationFee *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetLiquidationFee(&_AgentPoliceV2.TransactOpts, _liquidationFee)
}

// SetMaxMiners is a paid mutator transaction binding the contract method 0x8830d8c3.
//
// Solidity: function setMaxMiners(uint32 _maxMiners) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetMaxMiners(opts *bind.TransactOpts, _maxMiners uint32) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setMaxMiners", _maxMiners)
}

// SetMaxMiners is a paid mutator transaction binding the contract method 0x8830d8c3.
//
// Solidity: function setMaxMiners(uint32 _maxMiners) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetMaxMiners(_maxMiners uint32) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetMaxMiners(&_AgentPoliceV2.TransactOpts, _maxMiners)
}

// SetMaxMiners is a paid mutator transaction binding the contract method 0x8830d8c3.
//
// Solidity: function setMaxMiners(uint32 _maxMiners) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetMaxMiners(_maxMiners uint32) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetMaxMiners(&_AgentPoliceV2.TransactOpts, _maxMiners)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) SetSectorFaultyTolerancePercent(opts *bind.TransactOpts, _sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "setSectorFaultyTolerancePercent", _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetSectorFaultyTolerancePercent(&_AgentPoliceV2.TransactOpts, _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.SetSectorFaultyTolerancePercent(&_AgentPoliceV2.TransactOpts, _sectorFaultyTolerancePercent)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) TransferOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "transferOperator", newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.TransferOperator(&_AgentPoliceV2.TransactOpts, newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.TransferOperator(&_AgentPoliceV2.TransactOpts, newOperator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.TransferOwnership(&_AgentPoliceV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.TransferOwnership(&_AgentPoliceV2.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AgentPoliceV2 *AgentPoliceV2Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPoliceV2.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AgentPoliceV2 *AgentPoliceV2Session) Unpause() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.Unpause(&_AgentPoliceV2.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AgentPoliceV2 *AgentPoliceV2TransactorSession) Unpause() (*types.Transaction, error) {
	return _AgentPoliceV2.Contract.Unpause(&_AgentPoliceV2.TransactOpts)
}

// AgentPoliceV2CredentialUsedIterator is returned from FilterCredentialUsed and is used to iterate over the raw logs and unpacked data for CredentialUsed events raised by the AgentPoliceV2 contract.
type AgentPoliceV2CredentialUsedIterator struct {
	Event *AgentPoliceV2CredentialUsed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceV2CredentialUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceV2CredentialUsed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceV2CredentialUsed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceV2CredentialUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceV2CredentialUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceV2CredentialUsed represents a CredentialUsed event raised by the AgentPoliceV2 contract.
type AgentPoliceV2CredentialUsed struct {
	AgentID   *big.Int
	Action    [4]byte
	AgentAddr common.Address
	Value     *big.Int
	Target    uint64
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCredentialUsed is a free log retrieval operation binding the contract event 0x700450a0d828409c6031b95d040684b8a2eba879416b6960dfa8c818f245ddf4.
//
// Solidity: event CredentialUsed(uint256 indexed agentID, bytes4 indexed action, address agentAddr, uint256 value, uint64 target)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterCredentialUsed(opts *bind.FilterOpts, agentID []*big.Int, action [][4]byte) (*AgentPoliceV2CredentialUsedIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}
	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "CredentialUsed", agentIDRule, actionRule)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2CredentialUsedIterator{contract: _AgentPoliceV2.contract, event: "CredentialUsed", logs: logs, sub: sub}, nil
}

// WatchCredentialUsed is a free log subscription operation binding the contract event 0x700450a0d828409c6031b95d040684b8a2eba879416b6960dfa8c818f245ddf4.
//
// Solidity: event CredentialUsed(uint256 indexed agentID, bytes4 indexed action, address agentAddr, uint256 value, uint64 target)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchCredentialUsed(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2CredentialUsed, agentID []*big.Int, action [][4]byte) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}
	var actionRule []interface{}
	for _, actionItem := range action {
		actionRule = append(actionRule, actionItem)
	}

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "CredentialUsed", agentIDRule, actionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceV2CredentialUsed)
				if err := _AgentPoliceV2.contract.UnpackLog(event, "CredentialUsed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCredentialUsed is a log parse operation binding the contract event 0x700450a0d828409c6031b95d040684b8a2eba879416b6960dfa8c818f245ddf4.
//
// Solidity: event CredentialUsed(uint256 indexed agentID, bytes4 indexed action, address agentAddr, uint256 value, uint64 target)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) ParseCredentialUsed(log types.Log) (*AgentPoliceV2CredentialUsed, error) {
	event := new(AgentPoliceV2CredentialUsed)
	if err := _AgentPoliceV2.contract.UnpackLog(event, "CredentialUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceV2DefaultedIterator is returned from FilterDefaulted and is used to iterate over the raw logs and unpacked data for Defaulted events raised by the AgentPoliceV2 contract.
type AgentPoliceV2DefaultedIterator struct {
	Event *AgentPoliceV2Defaulted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceV2DefaultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceV2Defaulted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceV2Defaulted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceV2DefaultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceV2DefaultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceV2Defaulted represents a Defaulted event raised by the AgentPoliceV2 contract.
type AgentPoliceV2Defaulted struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaulted is a free log retrieval operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterDefaulted(opts *bind.FilterOpts) (*AgentPoliceV2DefaultedIterator, error) {

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "Defaulted")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2DefaultedIterator{contract: _AgentPoliceV2.contract, event: "Defaulted", logs: logs, sub: sub}, nil
}

// WatchDefaulted is a free log subscription operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchDefaulted(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2Defaulted) (event.Subscription, error) {

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "Defaulted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceV2Defaulted)
				if err := _AgentPoliceV2.contract.UnpackLog(event, "Defaulted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaulted is a log parse operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) ParseDefaulted(log types.Log) (*AgentPoliceV2Defaulted, error) {
	event := new(AgentPoliceV2Defaulted)
	if err := _AgentPoliceV2.contract.UnpackLog(event, "Defaulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceV2EIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the AgentPoliceV2 contract.
type AgentPoliceV2EIP712DomainChangedIterator struct {
	Event *AgentPoliceV2EIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceV2EIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceV2EIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceV2EIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceV2EIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceV2EIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceV2EIP712DomainChanged represents a EIP712DomainChanged event raised by the AgentPoliceV2 contract.
type AgentPoliceV2EIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*AgentPoliceV2EIP712DomainChangedIterator, error) {

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2EIP712DomainChangedIterator{contract: _AgentPoliceV2.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2EIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceV2EIP712DomainChanged)
				if err := _AgentPoliceV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_AgentPoliceV2 *AgentPoliceV2Filterer) ParseEIP712DomainChanged(log types.Log) (*AgentPoliceV2EIP712DomainChanged, error) {
	event := new(AgentPoliceV2EIP712DomainChanged)
	if err := _AgentPoliceV2.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceV2FaultySectorsIterator is returned from FilterFaultySectors and is used to iterate over the raw logs and unpacked data for FaultySectors events raised by the AgentPoliceV2 contract.
type AgentPoliceV2FaultySectorsIterator struct {
	Event *AgentPoliceV2FaultySectors // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceV2FaultySectorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceV2FaultySectors)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceV2FaultySectors)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceV2FaultySectorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceV2FaultySectorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceV2FaultySectors represents a FaultySectors event raised by the AgentPoliceV2 contract.
type AgentPoliceV2FaultySectors struct {
	AgentID    common.Address
	FaultEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFaultySectors is a free log retrieval operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterFaultySectors(opts *bind.FilterOpts, agentID []common.Address) (*AgentPoliceV2FaultySectorsIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "FaultySectors", agentIDRule)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2FaultySectorsIterator{contract: _AgentPoliceV2.contract, event: "FaultySectors", logs: logs, sub: sub}, nil
}

// WatchFaultySectors is a free log subscription operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchFaultySectors(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2FaultySectors, agentID []common.Address) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "FaultySectors", agentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceV2FaultySectors)
				if err := _AgentPoliceV2.contract.UnpackLog(event, "FaultySectors", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFaultySectors is a log parse operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) ParseFaultySectors(log types.Log) (*AgentPoliceV2FaultySectors, error) {
	event := new(AgentPoliceV2FaultySectors)
	if err := _AgentPoliceV2.contract.UnpackLog(event, "FaultySectors", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceV2OnAdministrationIterator is returned from FilterOnAdministration and is used to iterate over the raw logs and unpacked data for OnAdministration events raised by the AgentPoliceV2 contract.
type AgentPoliceV2OnAdministrationIterator struct {
	Event *AgentPoliceV2OnAdministration // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceV2OnAdministrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceV2OnAdministration)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceV2OnAdministration)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceV2OnAdministrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceV2OnAdministrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceV2OnAdministration represents a OnAdministration event raised by the AgentPoliceV2 contract.
type AgentPoliceV2OnAdministration struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnAdministration is a free log retrieval operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterOnAdministration(opts *bind.FilterOpts) (*AgentPoliceV2OnAdministrationIterator, error) {

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "OnAdministration")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2OnAdministrationIterator{contract: _AgentPoliceV2.contract, event: "OnAdministration", logs: logs, sub: sub}, nil
}

// WatchOnAdministration is a free log subscription operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchOnAdministration(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2OnAdministration) (event.Subscription, error) {

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "OnAdministration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceV2OnAdministration)
				if err := _AgentPoliceV2.contract.UnpackLog(event, "OnAdministration", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOnAdministration is a log parse operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) ParseOnAdministration(log types.Log) (*AgentPoliceV2OnAdministration, error) {
	event := new(AgentPoliceV2OnAdministration)
	if err := _AgentPoliceV2.contract.UnpackLog(event, "OnAdministration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceV2PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AgentPoliceV2 contract.
type AgentPoliceV2PausedIterator struct {
	Event *AgentPoliceV2Paused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceV2PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceV2Paused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceV2Paused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceV2PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceV2PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceV2Paused represents a Paused event raised by the AgentPoliceV2 contract.
type AgentPoliceV2Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterPaused(opts *bind.FilterOpts) (*AgentPoliceV2PausedIterator, error) {

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2PausedIterator{contract: _AgentPoliceV2.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2Paused) (event.Subscription, error) {

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceV2Paused)
				if err := _AgentPoliceV2.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) ParsePaused(log types.Log) (*AgentPoliceV2Paused, error) {
	event := new(AgentPoliceV2Paused)
	if err := _AgentPoliceV2.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceV2UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AgentPoliceV2 contract.
type AgentPoliceV2UnpausedIterator struct {
	Event *AgentPoliceV2Unpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentPoliceV2UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceV2Unpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentPoliceV2Unpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentPoliceV2UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceV2UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceV2Unpaused represents a Unpaused event raised by the AgentPoliceV2 contract.
type AgentPoliceV2Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterUnpaused(opts *bind.FilterOpts) (*AgentPoliceV2UnpausedIterator, error) {

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2UnpausedIterator{contract: _AgentPoliceV2.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2Unpaused) (event.Subscription, error) {

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceV2Unpaused)
				if err := _AgentPoliceV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) ParseUnpaused(log types.Log) (*AgentPoliceV2Unpaused, error) {
	event := new(AgentPoliceV2Unpaused)
	if err := _AgentPoliceV2.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
