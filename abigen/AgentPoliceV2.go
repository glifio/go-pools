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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentStateRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCredential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LiquidationValueTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxMinersReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverFaultySectorLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverLimitDTE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverLimitDTL\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverLimitQuota\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateChangeInProcess\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"CredentialUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"Defaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agentID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"faultEpoch\",\"type\":\"uint256\"}],\"name\":\"FaultySectors\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"OnAdministration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_VERIFIABLE_CREDENTIAL_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"agentApproved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"agentLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"borrowDTL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmAdministration\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmEquity\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"credentialUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"deriveStructHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distributeLiquidatedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"isValidCredential\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationDTL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidationFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMiners\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidator\",\"type\":\"uint64\"}],\"name\":\"prepareMinerForLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"registerCredentialUseBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sectorFaultyTolerancePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"setAgentDefaultDTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"agentIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"}],\"name\":\"setAgentLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_borrowDTL\",\"type\":\"uint256\"}],\"name\":\"setBorrowDTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"_levels\",\"type\":\"uint256[10]\"}],\"name\":\"setLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidationDTL\",\"type\":\"uint256\"}],\"name\":\"setLiquidationDTL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_liquidationFee\",\"type\":\"uint256\"}],\"name\":\"setLiquidationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_maxMiners\",\"type\":\"uint32\"}],\"name\":\"setMaxMiners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sectorFaultyTolerancePercent\",\"type\":\"uint256\"}],\"name\":\"setSectorFaultyTolerancePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"transferOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"validateCred\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AgentPoliceV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentPoliceV2MetaData.ABI instead.
var AgentPoliceV2ABI = AgentPoliceV2MetaData.ABI

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
	AgentID *big.Int
	Vc      VerifiableCredential
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCredentialUsed is a free log retrieval operation binding the contract event 0xe413f1e1e4e2982acfdd8591946c041a64d3366d4c7cbd5c773755720ebe794a.
//
// Solidity: event CredentialUsed(uint256 indexed agentID, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) FilterCredentialUsed(opts *bind.FilterOpts, agentID []*big.Int) (*AgentPoliceV2CredentialUsedIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPoliceV2.contract.FilterLogs(opts, "CredentialUsed", agentIDRule)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceV2CredentialUsedIterator{contract: _AgentPoliceV2.contract, event: "CredentialUsed", logs: logs, sub: sub}, nil
}

// WatchCredentialUsed is a free log subscription operation binding the contract event 0xe413f1e1e4e2982acfdd8591946c041a64d3366d4c7cbd5c773755720ebe794a.
//
// Solidity: event CredentialUsed(uint256 indexed agentID, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc)
func (_AgentPoliceV2 *AgentPoliceV2Filterer) WatchCredentialUsed(opts *bind.WatchOpts, sink chan<- *AgentPoliceV2CredentialUsed, agentID []*big.Int) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPoliceV2.contract.WatchLogs(opts, "CredentialUsed", agentIDRule)
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

// ParseCredentialUsed is a log parse operation binding the contract event 0xe413f1e1e4e2982acfdd8591946c041a64d3366d4c7cbd5c773755720ebe794a.
//
// Solidity: event CredentialUsed(uint256 indexed agentID, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc)
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
