// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/jimpick/go-ethereum"
	"github.com/jimpick/go-ethereum/accounts/abi"
	"github.com/jimpick/go-ethereum/accounts/abi/bind"
	"github.com/jimpick/go-ethereum/common"
	"github.com/jimpick/go-ethereum/core/types"
	"github.com/jimpick/go-ethereum/event"
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
type VerifiableCredential_AgentPolice struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// AgentPoliceMetaData contains all meta data concerning the AgentPolice contract.
var AgentPoliceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_defaultWindow\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"contractIPoolRegistry\",\"name\":\"_poolRegistry\",\"type\":\"address\"},{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AgentStateRejected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCredential\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"Defaulted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agentID\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"faultEpoch\",\"type\":\"uint256\"}],\"name\":\"FaultySectors\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"OnAdministration\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_VERIFIABLE_CREDENTIAL_TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"agentApproved\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"agentLiquidated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmAdministration\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"confirmRmEquity\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"credentialUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"deriveStructHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"digest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distributeLiquidatedFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"isValidCredential\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAgent[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"name\":\"markAsFaulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxConsecutiveFaultEpochs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEpochsOwedTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxPoolsPerAgent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidator\",\"type\":\"uint64\"}],\"name\":\"prepareMinerForLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"administration\",\"type\":\"address\"}],\"name\":\"putAgentOnAdministrationDueToFaultySectorDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"recover\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"registerCredentialUseBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sectorFaultyTolerancePercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setAgentDefaultDueToFaultySectorDays\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"setAgentDefaulted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_defaultWindow\",\"type\":\"uint256\"}],\"name\":\"setDefaultWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxConsecutiveFaultEpochs\",\"type\":\"uint256\"}],\"name\":\"setMaxConsecutiveFaultEpochs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDTE\",\"type\":\"uint256\"}],\"name\":\"setMaxDTE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxEpochsOwedTolerance\",\"type\":\"uint256\"}],\"name\":\"setMaxEpochsOwedTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLTV\",\"type\":\"uint256\"}],\"name\":\"setMaxLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxPoolsPerAgent\",\"type\":\"uint256\"}],\"name\":\"setMaxPoolsPerAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sectorFaultyTolerancePercent\",\"type\":\"uint256\"}],\"name\":\"setSectorFaultyTolerancePercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"transferOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"validateCred\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AgentPoliceABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentPoliceMetaData.ABI instead.
var AgentPoliceABI = AgentPoliceMetaData.ABI

// AgentPolice is an auto generated Go binding around an Ethereum contract.
type AgentPolice struct {
	AgentPoliceCaller     // Read-only binding to the contract
	AgentPoliceTransactor // Write-only binding to the contract
	AgentPoliceFilterer   // Log filterer for contract events
}

// AgentPoliceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentPoliceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentPoliceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentPoliceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentPoliceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentPoliceSession struct {
	Contract     *AgentPolice      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentPoliceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentPoliceCallerSession struct {
	Contract *AgentPoliceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AgentPoliceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentPoliceTransactorSession struct {
	Contract     *AgentPoliceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AgentPoliceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentPoliceRaw struct {
	Contract *AgentPolice // Generic contract binding to access the raw methods on
}

// AgentPoliceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentPoliceCallerRaw struct {
	Contract *AgentPoliceCaller // Generic read-only contract binding to access the raw methods on
}

// AgentPoliceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentPoliceTransactorRaw struct {
	Contract *AgentPoliceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentPolice creates a new instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPolice(address common.Address, backend bind.ContractBackend) (*AgentPolice, error) {
	contract, err := bindAgentPolice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentPolice{AgentPoliceCaller: AgentPoliceCaller{contract: contract}, AgentPoliceTransactor: AgentPoliceTransactor{contract: contract}, AgentPoliceFilterer: AgentPoliceFilterer{contract: contract}}, nil
}

// NewAgentPoliceCaller creates a new read-only instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPoliceCaller(address common.Address, caller bind.ContractCaller) (*AgentPoliceCaller, error) {
	contract, err := bindAgentPolice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceCaller{contract: contract}, nil
}

// NewAgentPoliceTransactor creates a new write-only instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPoliceTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentPoliceTransactor, error) {
	contract, err := bindAgentPolice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceTransactor{contract: contract}, nil
}

// NewAgentPoliceFilterer creates a new log filterer instance of AgentPolice, bound to a specific deployed contract.
func NewAgentPoliceFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentPoliceFilterer, error) {
	contract, err := bindAgentPolice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceFilterer{contract: contract}, nil
}

// bindAgentPolice binds a generic wrapper to an already deployed contract.
func bindAgentPolice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentPoliceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentPolice *AgentPoliceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentPolice.Contract.AgentPoliceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentPolice *AgentPoliceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.Contract.AgentPoliceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentPolice *AgentPoliceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentPolice.Contract.AgentPoliceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentPolice *AgentPoliceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentPolice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentPolice *AgentPoliceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentPolice *AgentPoliceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentPolice.Contract.contract.Transact(opts, method, params...)
}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPolice *AgentPoliceCaller) VERIFIABLECREDENTIALTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "_VERIFIABLE_CREDENTIAL_TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPolice *AgentPoliceSession) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _AgentPolice.Contract.VERIFIABLECREDENTIALTYPEHASH(&_AgentPolice.CallOpts)
}

// VERIFIABLECREDENTIALTYPEHASH is a free data retrieval call binding the contract method 0x2b85012b.
//
// Solidity: function _VERIFIABLE_CREDENTIAL_TYPE_HASH() view returns(bytes32)
func (_AgentPolice *AgentPoliceCallerSession) VERIFIABLECREDENTIALTYPEHASH() ([32]byte, error) {
	return _AgentPolice.Contract.VERIFIABLECREDENTIALTYPEHASH(&_AgentPolice.CallOpts)
}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCaller) AgentApproved(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "agentApproved", vc)

	if err != nil {
		return err
	}

	return err

}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceSession) AgentApproved(vc VerifiableCredential) error {
	return _AgentPolice.Contract.AgentApproved(&_AgentPolice.CallOpts, vc)
}

// AgentApproved is a free data retrieval call binding the contract method 0x260e63c6.
//
// Solidity: function agentApproved((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) AgentApproved(vc VerifiableCredential) error {
	return _AgentPolice.Contract.AgentApproved(&_AgentPolice.CallOpts, vc)
}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPolice *AgentPoliceCaller) AgentLiquidated(opts *bind.CallOpts, agentID *big.Int) (bool, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "agentLiquidated", agentID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPolice *AgentPoliceSession) AgentLiquidated(agentID *big.Int) (bool, error) {
	return _AgentPolice.Contract.AgentLiquidated(&_AgentPolice.CallOpts, agentID)
}

// AgentLiquidated is a free data retrieval call binding the contract method 0x51692149.
//
// Solidity: function agentLiquidated(uint256 agentID) view returns(bool)
func (_AgentPolice *AgentPoliceCallerSession) AgentLiquidated(agentID *big.Int) (bool, error) {
	return _AgentPolice.Contract.AgentLiquidated(&_AgentPolice.CallOpts, agentID)
}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCaller) ConfirmRmAdministration(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "confirmRmAdministration", vc)

	if err != nil {
		return err
	}

	return err

}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceSession) ConfirmRmAdministration(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmAdministration(&_AgentPolice.CallOpts, vc)
}

// ConfirmRmAdministration is a free data retrieval call binding the contract method 0x5cf6862f.
//
// Solidity: function confirmRmAdministration((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) ConfirmRmAdministration(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmAdministration(&_AgentPolice.CallOpts, vc)
}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCaller) ConfirmRmEquity(opts *bind.CallOpts, vc VerifiableCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "confirmRmEquity", vc)

	if err != nil {
		return err
	}

	return err

}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceSession) ConfirmRmEquity(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmEquity(&_AgentPolice.CallOpts, vc)
}

// ConfirmRmEquity is a free data retrieval call binding the contract method 0x5654657a.
//
// Solidity: function confirmRmEquity((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) ConfirmRmEquity(vc VerifiableCredential) error {
	return _AgentPolice.Contract.ConfirmRmEquity(&_AgentPolice.CallOpts, vc)
}

// CredentialUsed is a free data retrieval call binding the contract method 0xf37c8b85.
//
// Solidity: function credentialUsed(uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_AgentPolice *AgentPoliceCaller) CredentialUsed(opts *bind.CallOpts, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "credentialUsed", v, r, s)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CredentialUsed is a free data retrieval call binding the contract method 0xf37c8b85.
//
// Solidity: function credentialUsed(uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_AgentPolice *AgentPoliceSession) CredentialUsed(v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _AgentPolice.Contract.CredentialUsed(&_AgentPolice.CallOpts, v, r, s)
}

// CredentialUsed is a free data retrieval call binding the contract method 0xf37c8b85.
//
// Solidity: function credentialUsed(uint8 v, bytes32 r, bytes32 s) view returns(bool)
func (_AgentPolice *AgentPoliceCallerSession) CredentialUsed(v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _AgentPolice.Contract.CredentialUsed(&_AgentPolice.CallOpts, v, r, s)
}

// DefaultWindow is a free data retrieval call binding the contract method 0x78ee44fd.
//
// Solidity: function defaultWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) DefaultWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "defaultWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultWindow is a free data retrieval call binding the contract method 0x78ee44fd.
//
// Solidity: function defaultWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) DefaultWindow() (*big.Int, error) {
	return _AgentPolice.Contract.DefaultWindow(&_AgentPolice.CallOpts)
}

// DefaultWindow is a free data retrieval call binding the contract method 0x78ee44fd.
//
// Solidity: function defaultWindow() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) DefaultWindow() (*big.Int, error) {
	return _AgentPolice.Contract.DefaultWindow(&_AgentPolice.CallOpts)
}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPolice *AgentPoliceCaller) DeriveStructHash(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "deriveStructHash", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPolice *AgentPoliceSession) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.DeriveStructHash(&_AgentPolice.CallOpts, vc)
}

// DeriveStructHash is a free data retrieval call binding the contract method 0x40444c7c.
//
// Solidity: function deriveStructHash((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) pure returns(bytes32)
func (_AgentPolice *AgentPoliceCallerSession) DeriveStructHash(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.DeriveStructHash(&_AgentPolice.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPolice *AgentPoliceCaller) Digest(opts *bind.CallOpts, vc VerifiableCredential) ([32]byte, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "digest", vc)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPolice *AgentPoliceSession) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.Digest(&_AgentPolice.CallOpts, vc)
}

// Digest is a free data retrieval call binding the contract method 0xa2f3c210.
//
// Solidity: function digest((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bytes32)
func (_AgentPolice *AgentPoliceCallerSession) Digest(vc VerifiableCredential) ([32]byte, error) {
	return _AgentPolice.Contract.Digest(&_AgentPolice.CallOpts, vc)
}

// IsValidCredential is a free data retrieval call binding the contract method 0x76048dc4.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPolice *AgentPoliceCaller) IsValidCredential(opts *bind.CallOpts, agent *big.Int, action [4]byte, sc SignedCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "isValidCredential", agent, action, sc)

	if err != nil {
		return err
	}

	return err

}

// IsValidCredential is a free data retrieval call binding the contract method 0x76048dc4.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPolice *AgentPoliceSession) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.IsValidCredential(&_AgentPolice.CallOpts, agent, action, sc)
}

// IsValidCredential is a free data retrieval call binding the contract method 0x76048dc4.
//
// Solidity: function isValidCredential(uint256 agent, bytes4 action, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) IsValidCredential(agent *big.Int, action [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.IsValidCredential(&_AgentPolice.CallOpts, agent, action, sc)
}

// MaxConsecutiveFaultEpochs is a free data retrieval call binding the contract method 0x395bbfe9.
//
// Solidity: function maxConsecutiveFaultEpochs() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxConsecutiveFaultEpochs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxConsecutiveFaultEpochs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxConsecutiveFaultEpochs is a free data retrieval call binding the contract method 0x395bbfe9.
//
// Solidity: function maxConsecutiveFaultEpochs() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxConsecutiveFaultEpochs() (*big.Int, error) {
	return _AgentPolice.Contract.MaxConsecutiveFaultEpochs(&_AgentPolice.CallOpts)
}

// MaxConsecutiveFaultEpochs is a free data retrieval call binding the contract method 0x395bbfe9.
//
// Solidity: function maxConsecutiveFaultEpochs() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxConsecutiveFaultEpochs() (*big.Int, error) {
	return _AgentPolice.Contract.MaxConsecutiveFaultEpochs(&_AgentPolice.CallOpts)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxDTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxDTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxDTE() (*big.Int, error) {
	return _AgentPolice.Contract.MaxDTE(&_AgentPolice.CallOpts)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxDTE() (*big.Int, error) {
	return _AgentPolice.Contract.MaxDTE(&_AgentPolice.CallOpts)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxEpochsOwedTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxEpochsOwedTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _AgentPolice.Contract.MaxEpochsOwedTolerance(&_AgentPolice.CallOpts)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _AgentPolice.Contract.MaxEpochsOwedTolerance(&_AgentPolice.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxLTV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxLTV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxLTV() (*big.Int, error) {
	return _AgentPolice.Contract.MaxLTV(&_AgentPolice.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxLTV() (*big.Int, error) {
	return _AgentPolice.Contract.MaxLTV(&_AgentPolice.CallOpts)
}

// MaxPoolsPerAgent is a free data retrieval call binding the contract method 0x56146380.
//
// Solidity: function maxPoolsPerAgent() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) MaxPoolsPerAgent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "maxPoolsPerAgent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPoolsPerAgent is a free data retrieval call binding the contract method 0x56146380.
//
// Solidity: function maxPoolsPerAgent() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) MaxPoolsPerAgent() (*big.Int, error) {
	return _AgentPolice.Contract.MaxPoolsPerAgent(&_AgentPolice.CallOpts)
}

// MaxPoolsPerAgent is a free data retrieval call binding the contract method 0x56146380.
//
// Solidity: function maxPoolsPerAgent() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) MaxPoolsPerAgent() (*big.Int, error) {
	return _AgentPolice.Contract.MaxPoolsPerAgent(&_AgentPolice.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPolice *AgentPoliceCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPolice *AgentPoliceSession) Operator() (common.Address, error) {
	return _AgentPolice.Contract.Operator(&_AgentPolice.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) Operator() (common.Address, error) {
	return _AgentPolice.Contract.Operator(&_AgentPolice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPolice *AgentPoliceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPolice *AgentPoliceSession) Owner() (common.Address, error) {
	return _AgentPolice.Contract.Owner(&_AgentPolice.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) Owner() (common.Address, error) {
	return _AgentPolice.Contract.Owner(&_AgentPolice.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPolice *AgentPoliceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPolice *AgentPoliceSession) Paused() (bool, error) {
	return _AgentPolice.Contract.Paused(&_AgentPolice.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AgentPolice *AgentPoliceCallerSession) Paused() (bool, error) {
	return _AgentPolice.Contract.Paused(&_AgentPolice.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPolice *AgentPoliceCaller) PendingOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "pendingOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPolice *AgentPoliceSession) PendingOperator() (common.Address, error) {
	return _AgentPolice.Contract.PendingOperator(&_AgentPolice.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) PendingOperator() (common.Address, error) {
	return _AgentPolice.Contract.PendingOperator(&_AgentPolice.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPolice *AgentPoliceCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPolice *AgentPoliceSession) PendingOwner() (common.Address, error) {
	return _AgentPolice.Contract.PendingOwner(&_AgentPolice.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) PendingOwner() (common.Address, error) {
	return _AgentPolice.Contract.PendingOwner(&_AgentPolice.CallOpts)
}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_AgentPolice *AgentPoliceCaller) Recover(opts *bind.CallOpts, sc SignedCredential) (common.Address, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "recover", sc)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_AgentPolice *AgentPoliceSession) Recover(sc SignedCredential) (common.Address, error) {
	return _AgentPolice.Contract.Recover(&_AgentPolice.CallOpts, sc)
}

// Recover is a free data retrieval call binding the contract method 0x4ead1527.
//
// Solidity: function recover(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns(address)
func (_AgentPolice *AgentPoliceCallerSession) Recover(sc SignedCredential) (common.Address, error) {
	return _AgentPolice.Contract.Recover(&_AgentPolice.CallOpts, sc)
}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPolice *AgentPoliceCaller) SectorFaultyTolerancePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "sectorFaultyTolerancePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPolice *AgentPoliceSession) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _AgentPolice.Contract.SectorFaultyTolerancePercent(&_AgentPolice.CallOpts)
}

// SectorFaultyTolerancePercent is a free data retrieval call binding the contract method 0x37279488.
//
// Solidity: function sectorFaultyTolerancePercent() view returns(uint256)
func (_AgentPolice *AgentPoliceCallerSession) SectorFaultyTolerancePercent() (*big.Int, error) {
	return _AgentPolice.Contract.SectorFaultyTolerancePercent(&_AgentPolice.CallOpts)
}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPolice *AgentPoliceCaller) ValidateCred(opts *bind.CallOpts, agent *big.Int, selector [4]byte, sc SignedCredential) error {
	var out []interface{}
	err := _AgentPolice.contract.Call(opts, &out, "validateCred", agent, selector, sc)

	if err != nil {
		return err
	}

	return err

}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPolice *AgentPoliceSession) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.ValidateCred(&_AgentPolice.CallOpts, agent, selector, sc)
}

// ValidateCred is a free data retrieval call binding the contract method 0x403bb79d.
//
// Solidity: function validateCred(uint256 agent, bytes4 selector, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) view returns()
func (_AgentPolice *AgentPoliceCallerSession) ValidateCred(agent *big.Int, selector [4]byte, sc SignedCredential) error {
	return _AgentPolice.Contract.ValidateCred(&_AgentPolice.CallOpts, agent, selector, sc)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPolice *AgentPoliceTransactor) AcceptOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "acceptOperator")
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPolice *AgentPoliceSession) AcceptOperator() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOperator(&_AgentPolice.TransactOpts)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_AgentPolice *AgentPoliceTransactorSession) AcceptOperator() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOperator(&_AgentPolice.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPolice *AgentPoliceTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPolice *AgentPoliceSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOwnership(&_AgentPolice.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AgentPolice *AgentPoliceTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AgentPolice.Contract.AcceptOwnership(&_AgentPolice.TransactOpts)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPolice *AgentPoliceTransactor) DistributeLiquidatedFunds(opts *bind.TransactOpts, agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "distributeLiquidatedFunds", agent, amount)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPolice *AgentPoliceSession) DistributeLiquidatedFunds(agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.DistributeLiquidatedFunds(&_AgentPolice.TransactOpts, agent, amount)
}

// DistributeLiquidatedFunds is a paid mutator transaction binding the contract method 0x5e7981de.
//
// Solidity: function distributeLiquidatedFunds(address agent, uint256 amount) returns()
func (_AgentPolice *AgentPoliceTransactorSession) DistributeLiquidatedFunds(agent common.Address, amount *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.DistributeLiquidatedFunds(&_AgentPolice.TransactOpts, agent, amount)
}

// MarkAsFaulty is a paid mutator transaction binding the contract method 0x27ddb79d.
//
// Solidity: function markAsFaulty(address[] agents) returns()
func (_AgentPolice *AgentPoliceTransactor) MarkAsFaulty(opts *bind.TransactOpts, agents []common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "markAsFaulty", agents)
}

// MarkAsFaulty is a paid mutator transaction binding the contract method 0x27ddb79d.
//
// Solidity: function markAsFaulty(address[] agents) returns()
func (_AgentPolice *AgentPoliceSession) MarkAsFaulty(agents []common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.MarkAsFaulty(&_AgentPolice.TransactOpts, agents)
}

// MarkAsFaulty is a paid mutator transaction binding the contract method 0x27ddb79d.
//
// Solidity: function markAsFaulty(address[] agents) returns()
func (_AgentPolice *AgentPoliceTransactorSession) MarkAsFaulty(agents []common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.MarkAsFaulty(&_AgentPolice.TransactOpts, agents)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPolice *AgentPoliceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPolice *AgentPoliceSession) Pause() (*types.Transaction, error) {
	return _AgentPolice.Contract.Pause(&_AgentPolice.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AgentPolice *AgentPoliceTransactorSession) Pause() (*types.Transaction, error) {
	return _AgentPolice.Contract.Pause(&_AgentPolice.TransactOpts)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPolice *AgentPoliceTransactor) PrepareMinerForLiquidation(opts *bind.TransactOpts, agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "prepareMinerForLiquidation", agent, miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPolice *AgentPoliceSession) PrepareMinerForLiquidation(agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPolice.Contract.PrepareMinerForLiquidation(&_AgentPolice.TransactOpts, agent, miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0x9b83b8cd.
//
// Solidity: function prepareMinerForLiquidation(address agent, uint64 miner, uint64 liquidator) returns()
func (_AgentPolice *AgentPoliceTransactorSession) PrepareMinerForLiquidation(agent common.Address, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _AgentPolice.Contract.PrepareMinerForLiquidation(&_AgentPolice.TransactOpts, agent, miner, liquidator)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x2c6de7c9.
//
// Solidity: function putAgentOnAdministration(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactor) PutAgentOnAdministration(opts *bind.TransactOpts, agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "putAgentOnAdministration", agent, administration)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x2c6de7c9.
//
// Solidity: function putAgentOnAdministration(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceSession) PutAgentOnAdministration(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministration(&_AgentPolice.TransactOpts, agent, administration)
}

// PutAgentOnAdministration is a paid mutator transaction binding the contract method 0x2c6de7c9.
//
// Solidity: function putAgentOnAdministration(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactorSession) PutAgentOnAdministration(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministration(&_AgentPolice.TransactOpts, agent, administration)
}

// PutAgentOnAdministrationDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xcb70f563.
//
// Solidity: function putAgentOnAdministrationDueToFaultySectorDays(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactor) PutAgentOnAdministrationDueToFaultySectorDays(opts *bind.TransactOpts, agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "putAgentOnAdministrationDueToFaultySectorDays", agent, administration)
}

// PutAgentOnAdministrationDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xcb70f563.
//
// Solidity: function putAgentOnAdministrationDueToFaultySectorDays(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceSession) PutAgentOnAdministrationDueToFaultySectorDays(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministrationDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent, administration)
}

// PutAgentOnAdministrationDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xcb70f563.
//
// Solidity: function putAgentOnAdministrationDueToFaultySectorDays(address agent, address administration) returns()
func (_AgentPolice *AgentPoliceTransactorSession) PutAgentOnAdministrationDueToFaultySectorDays(agent common.Address, administration common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.PutAgentOnAdministrationDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent, administration)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_AgentPolice *AgentPoliceTransactor) RefreshRoutes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "refreshRoutes")
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_AgentPolice *AgentPoliceSession) RefreshRoutes() (*types.Transaction, error) {
	return _AgentPolice.Contract.RefreshRoutes(&_AgentPolice.TransactOpts)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_AgentPolice *AgentPoliceTransactorSession) RefreshRoutes() (*types.Transaction, error) {
	return _AgentPolice.Contract.RefreshRoutes(&_AgentPolice.TransactOpts)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0x529e2512.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPolice *AgentPoliceTransactor) RegisterCredentialUseBlock(opts *bind.TransactOpts, sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "registerCredentialUseBlock", sc)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0x529e2512.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPolice *AgentPoliceSession) RegisterCredentialUseBlock(sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.RegisterCredentialUseBlock(&_AgentPolice.TransactOpts, sc)
}

// RegisterCredentialUseBlock is a paid mutator transaction binding the contract method 0x529e2512.
//
// Solidity: function registerCredentialUseBlock(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_AgentPolice *AgentPoliceTransactorSession) RegisterCredentialUseBlock(sc SignedCredential) (*types.Transaction, error) {
	return _AgentPolice.Contract.RegisterCredentialUseBlock(&_AgentPolice.TransactOpts, sc)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AgentPolice *AgentPoliceTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AgentPolice *AgentPoliceSession) Resume() (*types.Transaction, error) {
	return _AgentPolice.Contract.Resume(&_AgentPolice.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_AgentPolice *AgentPoliceTransactorSession) Resume() (*types.Transaction, error) {
	return _AgentPolice.Contract.Resume(&_AgentPolice.TransactOpts)
}

// SetAgentDefaultDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xa86beff8.
//
// Solidity: function setAgentDefaultDueToFaultySectorDays(address agent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetAgentDefaultDueToFaultySectorDays(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setAgentDefaultDueToFaultySectorDays", agent)
}

// SetAgentDefaultDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xa86beff8.
//
// Solidity: function setAgentDefaultDueToFaultySectorDays(address agent) returns()
func (_AgentPolice *AgentPoliceSession) SetAgentDefaultDueToFaultySectorDays(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaultDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent)
}

// SetAgentDefaultDueToFaultySectorDays is a paid mutator transaction binding the contract method 0xa86beff8.
//
// Solidity: function setAgentDefaultDueToFaultySectorDays(address agent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetAgentDefaultDueToFaultySectorDays(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaultDueToFaultySectorDays(&_AgentPolice.TransactOpts, agent)
}

// SetAgentDefaulted is a paid mutator transaction binding the contract method 0x8b21505a.
//
// Solidity: function setAgentDefaulted(address agent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetAgentDefaulted(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setAgentDefaulted", agent)
}

// SetAgentDefaulted is a paid mutator transaction binding the contract method 0x8b21505a.
//
// Solidity: function setAgentDefaulted(address agent) returns()
func (_AgentPolice *AgentPoliceSession) SetAgentDefaulted(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaulted(&_AgentPolice.TransactOpts, agent)
}

// SetAgentDefaulted is a paid mutator transaction binding the contract method 0x8b21505a.
//
// Solidity: function setAgentDefaulted(address agent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetAgentDefaulted(agent common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetAgentDefaulted(&_AgentPolice.TransactOpts, agent)
}

// SetDefaultWindow is a paid mutator transaction binding the contract method 0xacec9564.
//
// Solidity: function setDefaultWindow(uint256 _defaultWindow) returns()
func (_AgentPolice *AgentPoliceTransactor) SetDefaultWindow(opts *bind.TransactOpts, _defaultWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setDefaultWindow", _defaultWindow)
}

// SetDefaultWindow is a paid mutator transaction binding the contract method 0xacec9564.
//
// Solidity: function setDefaultWindow(uint256 _defaultWindow) returns()
func (_AgentPolice *AgentPoliceSession) SetDefaultWindow(_defaultWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetDefaultWindow(&_AgentPolice.TransactOpts, _defaultWindow)
}

// SetDefaultWindow is a paid mutator transaction binding the contract method 0xacec9564.
//
// Solidity: function setDefaultWindow(uint256 _defaultWindow) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetDefaultWindow(_defaultWindow *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetDefaultWindow(&_AgentPolice.TransactOpts, _defaultWindow)
}

// SetMaxConsecutiveFaultEpochs is a paid mutator transaction binding the contract method 0xaf10015e.
//
// Solidity: function setMaxConsecutiveFaultEpochs(uint256 _maxConsecutiveFaultEpochs) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxConsecutiveFaultEpochs(opts *bind.TransactOpts, _maxConsecutiveFaultEpochs *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxConsecutiveFaultEpochs", _maxConsecutiveFaultEpochs)
}

// SetMaxConsecutiveFaultEpochs is a paid mutator transaction binding the contract method 0xaf10015e.
//
// Solidity: function setMaxConsecutiveFaultEpochs(uint256 _maxConsecutiveFaultEpochs) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxConsecutiveFaultEpochs(_maxConsecutiveFaultEpochs *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxConsecutiveFaultEpochs(&_AgentPolice.TransactOpts, _maxConsecutiveFaultEpochs)
}

// SetMaxConsecutiveFaultEpochs is a paid mutator transaction binding the contract method 0xaf10015e.
//
// Solidity: function setMaxConsecutiveFaultEpochs(uint256 _maxConsecutiveFaultEpochs) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxConsecutiveFaultEpochs(_maxConsecutiveFaultEpochs *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxConsecutiveFaultEpochs(&_AgentPolice.TransactOpts, _maxConsecutiveFaultEpochs)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxDTE(opts *bind.TransactOpts, _maxDTE *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxDTE", _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxDTE(&_AgentPolice.TransactOpts, _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxDTE(&_AgentPolice.TransactOpts, _maxDTE)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxEpochsOwedTolerance(opts *bind.TransactOpts, _maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxEpochsOwedTolerance", _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxEpochsOwedTolerance(&_AgentPolice.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxEpochsOwedTolerance(&_AgentPolice.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxLTV(opts *bind.TransactOpts, _maxLTV *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxLTV", _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxLTV(&_AgentPolice.TransactOpts, _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxLTV(&_AgentPolice.TransactOpts, _maxLTV)
}

// SetMaxPoolsPerAgent is a paid mutator transaction binding the contract method 0x64af5690.
//
// Solidity: function setMaxPoolsPerAgent(uint256 _maxPoolsPerAgent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetMaxPoolsPerAgent(opts *bind.TransactOpts, _maxPoolsPerAgent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setMaxPoolsPerAgent", _maxPoolsPerAgent)
}

// SetMaxPoolsPerAgent is a paid mutator transaction binding the contract method 0x64af5690.
//
// Solidity: function setMaxPoolsPerAgent(uint256 _maxPoolsPerAgent) returns()
func (_AgentPolice *AgentPoliceSession) SetMaxPoolsPerAgent(_maxPoolsPerAgent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxPoolsPerAgent(&_AgentPolice.TransactOpts, _maxPoolsPerAgent)
}

// SetMaxPoolsPerAgent is a paid mutator transaction binding the contract method 0x64af5690.
//
// Solidity: function setMaxPoolsPerAgent(uint256 _maxPoolsPerAgent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetMaxPoolsPerAgent(_maxPoolsPerAgent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetMaxPoolsPerAgent(&_AgentPolice.TransactOpts, _maxPoolsPerAgent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPolice *AgentPoliceTransactor) SetSectorFaultyTolerancePercent(opts *bind.TransactOpts, _sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "setSectorFaultyTolerancePercent", _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPolice *AgentPoliceSession) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetSectorFaultyTolerancePercent(&_AgentPolice.TransactOpts, _sectorFaultyTolerancePercent)
}

// SetSectorFaultyTolerancePercent is a paid mutator transaction binding the contract method 0x680af724.
//
// Solidity: function setSectorFaultyTolerancePercent(uint256 _sectorFaultyTolerancePercent) returns()
func (_AgentPolice *AgentPoliceTransactorSession) SetSectorFaultyTolerancePercent(_sectorFaultyTolerancePercent *big.Int) (*types.Transaction, error) {
	return _AgentPolice.Contract.SetSectorFaultyTolerancePercent(&_AgentPolice.TransactOpts, _sectorFaultyTolerancePercent)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPolice *AgentPoliceTransactor) TransferOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "transferOperator", newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPolice *AgentPoliceSession) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOperator(&_AgentPolice.TransactOpts, newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_AgentPolice *AgentPoliceTransactorSession) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOperator(&_AgentPolice.TransactOpts, newOperator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPolice *AgentPoliceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentPolice.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPolice *AgentPoliceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOwnership(&_AgentPolice.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentPolice *AgentPoliceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentPolice.Contract.TransferOwnership(&_AgentPolice.TransactOpts, newOwner)
}

// AgentPoliceDefaultedIterator is returned from FilterDefaulted and is used to iterate over the raw logs and unpacked data for Defaulted events raised by the AgentPolice contract.
type AgentPoliceDefaultedIterator struct {
	Event *AgentPoliceDefaulted // Event containing the contract specifics and raw log

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
func (it *AgentPoliceDefaultedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceDefaulted)
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
		it.Event = new(AgentPoliceDefaulted)
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
func (it *AgentPoliceDefaultedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceDefaultedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceDefaulted represents a Defaulted event raised by the AgentPolice contract.
type AgentPoliceDefaulted struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDefaulted is a free log retrieval operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPolice *AgentPoliceFilterer) FilterDefaulted(opts *bind.FilterOpts) (*AgentPoliceDefaultedIterator, error) {

	logs, sub, err := _AgentPolice.contract.FilterLogs(opts, "Defaulted")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceDefaultedIterator{contract: _AgentPolice.contract, event: "Defaulted", logs: logs, sub: sub}, nil
}

// WatchDefaulted is a free log subscription operation binding the contract event 0x0be58193d1ccca9f578279acffb21de6a55ce255b7c4901d2d3f4fbcd15db8b1.
//
// Solidity: event Defaulted(address agent)
func (_AgentPolice *AgentPoliceFilterer) WatchDefaulted(opts *bind.WatchOpts, sink chan<- *AgentPoliceDefaulted) (event.Subscription, error) {

	logs, sub, err := _AgentPolice.contract.WatchLogs(opts, "Defaulted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceDefaulted)
				if err := _AgentPolice.contract.UnpackLog(event, "Defaulted", log); err != nil {
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
func (_AgentPolice *AgentPoliceFilterer) ParseDefaulted(log types.Log) (*AgentPoliceDefaulted, error) {
	event := new(AgentPoliceDefaulted)
	if err := _AgentPolice.contract.UnpackLog(event, "Defaulted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceFaultySectorsIterator is returned from FilterFaultySectors and is used to iterate over the raw logs and unpacked data for FaultySectors events raised by the AgentPolice contract.
type AgentPoliceFaultySectorsIterator struct {
	Event *AgentPoliceFaultySectors // Event containing the contract specifics and raw log

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
func (it *AgentPoliceFaultySectorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceFaultySectors)
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
		it.Event = new(AgentPoliceFaultySectors)
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
func (it *AgentPoliceFaultySectorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceFaultySectorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceFaultySectors represents a FaultySectors event raised by the AgentPolice contract.
type AgentPoliceFaultySectors struct {
	AgentID    common.Address
	FaultEpoch *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFaultySectors is a free log retrieval operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPolice *AgentPoliceFilterer) FilterFaultySectors(opts *bind.FilterOpts, agentID []common.Address) (*AgentPoliceFaultySectorsIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPolice.contract.FilterLogs(opts, "FaultySectors", agentIDRule)
	if err != nil {
		return nil, err
	}
	return &AgentPoliceFaultySectorsIterator{contract: _AgentPolice.contract, event: "FaultySectors", logs: logs, sub: sub}, nil
}

// WatchFaultySectors is a free log subscription operation binding the contract event 0xae3b2aa2246d19333b887bb4211187ae48506f786b1edae51bde48cef3f52bd6.
//
// Solidity: event FaultySectors(address indexed agentID, uint256 faultEpoch)
func (_AgentPolice *AgentPoliceFilterer) WatchFaultySectors(opts *bind.WatchOpts, sink chan<- *AgentPoliceFaultySectors, agentID []common.Address) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _AgentPolice.contract.WatchLogs(opts, "FaultySectors", agentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceFaultySectors)
				if err := _AgentPolice.contract.UnpackLog(event, "FaultySectors", log); err != nil {
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
func (_AgentPolice *AgentPoliceFilterer) ParseFaultySectors(log types.Log) (*AgentPoliceFaultySectors, error) {
	event := new(AgentPoliceFaultySectors)
	if err := _AgentPolice.contract.UnpackLog(event, "FaultySectors", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentPoliceOnAdministrationIterator is returned from FilterOnAdministration and is used to iterate over the raw logs and unpacked data for OnAdministration events raised by the AgentPolice contract.
type AgentPoliceOnAdministrationIterator struct {
	Event *AgentPoliceOnAdministration // Event containing the contract specifics and raw log

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
func (it *AgentPoliceOnAdministrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentPoliceOnAdministration)
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
		it.Event = new(AgentPoliceOnAdministration)
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
func (it *AgentPoliceOnAdministrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentPoliceOnAdministrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentPoliceOnAdministration represents a OnAdministration event raised by the AgentPolice contract.
type AgentPoliceOnAdministration struct {
	Agent common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOnAdministration is a free log retrieval operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPolice *AgentPoliceFilterer) FilterOnAdministration(opts *bind.FilterOpts) (*AgentPoliceOnAdministrationIterator, error) {

	logs, sub, err := _AgentPolice.contract.FilterLogs(opts, "OnAdministration")
	if err != nil {
		return nil, err
	}
	return &AgentPoliceOnAdministrationIterator{contract: _AgentPolice.contract, event: "OnAdministration", logs: logs, sub: sub}, nil
}

// WatchOnAdministration is a free log subscription operation binding the contract event 0xcd2997877d26346cc62bc950a5d0a8a22f65bed07abe918de90bd967f79718c2.
//
// Solidity: event OnAdministration(address agent)
func (_AgentPolice *AgentPoliceFilterer) WatchOnAdministration(opts *bind.WatchOpts, sink chan<- *AgentPoliceOnAdministration) (event.Subscription, error) {

	logs, sub, err := _AgentPolice.contract.WatchLogs(opts, "OnAdministration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentPoliceOnAdministration)
				if err := _AgentPolice.contract.UnpackLog(event, "OnAdministration", log); err != nil {
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
func (_AgentPolice *AgentPoliceFilterer) ParseOnAdministration(log types.Log) (*AgentPoliceOnAdministration, error) {
	event := new(AgentPoliceOnAdministration)
	if err := _AgentPolice.contract.UnpackLog(event, "OnAdministration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
