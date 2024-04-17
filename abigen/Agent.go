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
type SignedCredential struct {
	Vc VerifiableCredential
	V  uint8
	R  [32]byte
	S  [32]byte
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// AgentMetaData contains all meta data concerning the Agent contract.
var AgentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"agentRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"apiRequestKey\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BadAgentState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"addMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"administration\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adoRequestKey\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"worker\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"controlAddresses\",\"type\":\"uint64[]\"}],\"name\":\"changeMinerWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"confirmChangeMinerWorker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newAgent\",\"type\":\"address\"}],\"name\":\"decommissionAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaulted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"faultySectorStartEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"migrateMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"operator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOperator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"liquidator\",\"type\":\"uint64\"}],\"name\":\"prepareMinerForLiquidation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"pullFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"pushFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"newMinerOwner\",\"type\":\"uint64\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"removeMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_administration\",\"type\":\"address\"}],\"name\":\"setAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newKey\",\"type\":\"address\"}],\"name\":\"setAdoRequestKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setFaulty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setInDefault\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"setRecovered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"transferOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignedCredential\",\"name\":\"sc\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AgentABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentMetaData.ABI instead.
var AgentABI = AgentMetaData.ABI

// Agent is an auto generated Go binding around an Ethereum contract.
type Agent struct {
	AgentCaller     // Read-only binding to the contract
	AgentTransactor // Write-only binding to the contract
	AgentFilterer   // Log filterer for contract events
}

// AgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentSession struct {
	Contract     *Agent            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentCallerSession struct {
	Contract *AgentCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentTransactorSession struct {
	Contract     *AgentTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentRaw struct {
	Contract *Agent // Generic contract binding to access the raw methods on
}

// AgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentCallerRaw struct {
	Contract *AgentCaller // Generic read-only contract binding to access the raw methods on
}

// AgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentTransactorRaw struct {
	Contract *AgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgent creates a new instance of Agent, bound to a specific deployed contract.
func NewAgent(address common.Address, backend bind.ContractBackend) (*Agent, error) {
	contract, err := bindAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Agent{AgentCaller: AgentCaller{contract: contract}, AgentTransactor: AgentTransactor{contract: contract}, AgentFilterer: AgentFilterer{contract: contract}}, nil
}

// NewAgentCaller creates a new read-only instance of Agent, bound to a specific deployed contract.
func NewAgentCaller(address common.Address, caller bind.ContractCaller) (*AgentCaller, error) {
	contract, err := bindAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentCaller{contract: contract}, nil
}

// NewAgentTransactor creates a new write-only instance of Agent, bound to a specific deployed contract.
func NewAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentTransactor, error) {
	contract, err := bindAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentTransactor{contract: contract}, nil
}

// NewAgentFilterer creates a new log filterer instance of Agent, bound to a specific deployed contract.
func NewAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentFilterer, error) {
	contract, err := bindAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentFilterer{contract: contract}, nil
}

// bindAgent binds a generic wrapper to an already deployed contract.
func bindAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agent *AgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agent.Contract.AgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agent *AgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.Contract.AgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agent *AgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agent.Contract.AgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Agent *AgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Agent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Agent *AgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Agent *AgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Agent.Contract.contract.Transact(opts, method, params...)
}

// Administration is a free data retrieval call binding the contract method 0x3847cb59.
//
// Solidity: function administration() view returns(address)
func (_Agent *AgentCaller) Administration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "administration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Administration is a free data retrieval call binding the contract method 0x3847cb59.
//
// Solidity: function administration() view returns(address)
func (_Agent *AgentSession) Administration() (common.Address, error) {
	return _Agent.Contract.Administration(&_Agent.CallOpts)
}

// Administration is a free data retrieval call binding the contract method 0x3847cb59.
//
// Solidity: function administration() view returns(address)
func (_Agent *AgentCallerSession) Administration() (common.Address, error) {
	return _Agent.Contract.Administration(&_Agent.CallOpts)
}

// AdoRequestKey is a free data retrieval call binding the contract method 0xd502cc1f.
//
// Solidity: function adoRequestKey() view returns(address)
func (_Agent *AgentCaller) AdoRequestKey(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "adoRequestKey")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdoRequestKey is a free data retrieval call binding the contract method 0xd502cc1f.
//
// Solidity: function adoRequestKey() view returns(address)
func (_Agent *AgentSession) AdoRequestKey() (common.Address, error) {
	return _Agent.Contract.AdoRequestKey(&_Agent.CallOpts)
}

// AdoRequestKey is a free data retrieval call binding the contract method 0xd502cc1f.
//
// Solidity: function adoRequestKey() view returns(address)
func (_Agent *AgentCallerSession) AdoRequestKey() (common.Address, error) {
	return _Agent.Contract.AdoRequestKey(&_Agent.CallOpts)
}

// Defaulted is a free data retrieval call binding the contract method 0x69e25ec1.
//
// Solidity: function defaulted() view returns(bool)
func (_Agent *AgentCaller) Defaulted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "defaulted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Defaulted is a free data retrieval call binding the contract method 0x69e25ec1.
//
// Solidity: function defaulted() view returns(bool)
func (_Agent *AgentSession) Defaulted() (bool, error) {
	return _Agent.Contract.Defaulted(&_Agent.CallOpts)
}

// Defaulted is a free data retrieval call binding the contract method 0x69e25ec1.
//
// Solidity: function defaulted() view returns(bool)
func (_Agent *AgentCallerSession) Defaulted() (bool, error) {
	return _Agent.Contract.Defaulted(&_Agent.CallOpts)
}

// FaultySectorStartEpoch is a free data retrieval call binding the contract method 0x8903f7f0.
//
// Solidity: function faultySectorStartEpoch() view returns(uint256)
func (_Agent *AgentCaller) FaultySectorStartEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "faultySectorStartEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FaultySectorStartEpoch is a free data retrieval call binding the contract method 0x8903f7f0.
//
// Solidity: function faultySectorStartEpoch() view returns(uint256)
func (_Agent *AgentSession) FaultySectorStartEpoch() (*big.Int, error) {
	return _Agent.Contract.FaultySectorStartEpoch(&_Agent.CallOpts)
}

// FaultySectorStartEpoch is a free data retrieval call binding the contract method 0x8903f7f0.
//
// Solidity: function faultySectorStartEpoch() view returns(uint256)
func (_Agent *AgentCallerSession) FaultySectorStartEpoch() (*big.Int, error) {
	return _Agent.Contract.FaultySectorStartEpoch(&_Agent.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Agent *AgentCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Agent *AgentSession) Id() (*big.Int, error) {
	return _Agent.Contract.Id(&_Agent.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_Agent *AgentCallerSession) Id() (*big.Int, error) {
	return _Agent.Contract.Id(&_Agent.CallOpts)
}

// LiquidAssets is a free data retrieval call binding the contract method 0xe492cdce.
//
// Solidity: function liquidAssets() view returns(uint256)
func (_Agent *AgentCaller) LiquidAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "liquidAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidAssets is a free data retrieval call binding the contract method 0xe492cdce.
//
// Solidity: function liquidAssets() view returns(uint256)
func (_Agent *AgentSession) LiquidAssets() (*big.Int, error) {
	return _Agent.Contract.LiquidAssets(&_Agent.CallOpts)
}

// LiquidAssets is a free data retrieval call binding the contract method 0xe492cdce.
//
// Solidity: function liquidAssets() view returns(uint256)
func (_Agent *AgentCallerSession) LiquidAssets() (*big.Int, error) {
	return _Agent.Contract.LiquidAssets(&_Agent.CallOpts)
}

// NewAgent is a free data retrieval call binding the contract method 0x56c01ad6.
//
// Solidity: function newAgent() view returns(address)
func (_Agent *AgentCaller) NewAgent(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "newAgent")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewAgent is a free data retrieval call binding the contract method 0x56c01ad6.
//
// Solidity: function newAgent() view returns(address)
func (_Agent *AgentSession) NewAgent() (common.Address, error) {
	return _Agent.Contract.NewAgent(&_Agent.CallOpts)
}

// NewAgent is a free data retrieval call binding the contract method 0x56c01ad6.
//
// Solidity: function newAgent() view returns(address)
func (_Agent *AgentCallerSession) NewAgent() (common.Address, error) {
	return _Agent.Contract.NewAgent(&_Agent.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Agent *AgentCaller) Operator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "operator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Agent *AgentSession) Operator() (common.Address, error) {
	return _Agent.Contract.Operator(&_Agent.CallOpts)
}

// Operator is a free data retrieval call binding the contract method 0x570ca735.
//
// Solidity: function operator() view returns(address)
func (_Agent *AgentCallerSession) Operator() (common.Address, error) {
	return _Agent.Contract.Operator(&_Agent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Agent *AgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Agent *AgentSession) Owner() (common.Address, error) {
	return _Agent.Contract.Owner(&_Agent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Agent *AgentCallerSession) Owner() (common.Address, error) {
	return _Agent.Contract.Owner(&_Agent.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_Agent *AgentCaller) PendingOperator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "pendingOperator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_Agent *AgentSession) PendingOperator() (common.Address, error) {
	return _Agent.Contract.PendingOperator(&_Agent.CallOpts)
}

// PendingOperator is a free data retrieval call binding the contract method 0xac7e534e.
//
// Solidity: function pendingOperator() view returns(address)
func (_Agent *AgentCallerSession) PendingOperator() (common.Address, error) {
	return _Agent.Contract.PendingOperator(&_Agent.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Agent *AgentCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Agent *AgentSession) PendingOwner() (common.Address, error) {
	return _Agent.Contract.PendingOwner(&_Agent.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Agent *AgentCallerSession) PendingOwner() (common.Address, error) {
	return _Agent.Contract.PendingOwner(&_Agent.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Agent *AgentCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Agent.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Agent *AgentSession) Version() (uint8, error) {
	return _Agent.Contract.Version(&_Agent.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_Agent *AgentCallerSession) Version() (uint8, error) {
	return _Agent.Contract.Version(&_Agent.CallOpts)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_Agent *AgentTransactor) AcceptOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "acceptOperator")
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_Agent *AgentSession) AcceptOperator() (*types.Transaction, error) {
	return _Agent.Contract.AcceptOperator(&_Agent.TransactOpts)
}

// AcceptOperator is a paid mutator transaction binding the contract method 0xfcad4448.
//
// Solidity: function acceptOperator() returns()
func (_Agent *AgentTransactorSession) AcceptOperator() (*types.Transaction, error) {
	return _Agent.Contract.AcceptOperator(&_Agent.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Agent *AgentTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Agent *AgentSession) AcceptOwnership() (*types.Transaction, error) {
	return _Agent.Contract.AcceptOwnership(&_Agent.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Agent *AgentTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Agent.Contract.AcceptOwnership(&_Agent.TransactOpts)
}

// AddMiner is a paid mutator transaction binding the contract method 0x2bb9af43.
//
// Solidity: function addMiner(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactor) AddMiner(opts *bind.TransactOpts, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "addMiner", sc)
}

// AddMiner is a paid mutator transaction binding the contract method 0x2bb9af43.
//
// Solidity: function addMiner(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentSession) AddMiner(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.AddMiner(&_Agent.TransactOpts, sc)
}

// AddMiner is a paid mutator transaction binding the contract method 0x2bb9af43.
//
// Solidity: function addMiner(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactorSession) AddMiner(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.AddMiner(&_Agent.TransactOpts, sc)
}

// Borrow is a paid mutator transaction binding the contract method 0x5859117d.
//
// Solidity: function borrow(uint256 poolID, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactor) Borrow(opts *bind.TransactOpts, poolID *big.Int, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "borrow", poolID, sc)
}

// Borrow is a paid mutator transaction binding the contract method 0x5859117d.
//
// Solidity: function borrow(uint256 poolID, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentSession) Borrow(poolID *big.Int, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.Borrow(&_Agent.TransactOpts, poolID, sc)
}

// Borrow is a paid mutator transaction binding the contract method 0x5859117d.
//
// Solidity: function borrow(uint256 poolID, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactorSession) Borrow(poolID *big.Int, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.Borrow(&_Agent.TransactOpts, poolID, sc)
}

// ChangeMinerWorker is a paid mutator transaction binding the contract method 0x8194360e.
//
// Solidity: function changeMinerWorker(uint64 miner, uint64 worker, uint64[] controlAddresses) returns()
func (_Agent *AgentTransactor) ChangeMinerWorker(opts *bind.TransactOpts, miner uint64, worker uint64, controlAddresses []uint64) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "changeMinerWorker", miner, worker, controlAddresses)
}

// ChangeMinerWorker is a paid mutator transaction binding the contract method 0x8194360e.
//
// Solidity: function changeMinerWorker(uint64 miner, uint64 worker, uint64[] controlAddresses) returns()
func (_Agent *AgentSession) ChangeMinerWorker(miner uint64, worker uint64, controlAddresses []uint64) (*types.Transaction, error) {
	return _Agent.Contract.ChangeMinerWorker(&_Agent.TransactOpts, miner, worker, controlAddresses)
}

// ChangeMinerWorker is a paid mutator transaction binding the contract method 0x8194360e.
//
// Solidity: function changeMinerWorker(uint64 miner, uint64 worker, uint64[] controlAddresses) returns()
func (_Agent *AgentTransactorSession) ChangeMinerWorker(miner uint64, worker uint64, controlAddresses []uint64) (*types.Transaction, error) {
	return _Agent.Contract.ChangeMinerWorker(&_Agent.TransactOpts, miner, worker, controlAddresses)
}

// ConfirmChangeMinerWorker is a paid mutator transaction binding the contract method 0xd16ac933.
//
// Solidity: function confirmChangeMinerWorker(uint64 miner) returns()
func (_Agent *AgentTransactor) ConfirmChangeMinerWorker(opts *bind.TransactOpts, miner uint64) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "confirmChangeMinerWorker", miner)
}

// ConfirmChangeMinerWorker is a paid mutator transaction binding the contract method 0xd16ac933.
//
// Solidity: function confirmChangeMinerWorker(uint64 miner) returns()
func (_Agent *AgentSession) ConfirmChangeMinerWorker(miner uint64) (*types.Transaction, error) {
	return _Agent.Contract.ConfirmChangeMinerWorker(&_Agent.TransactOpts, miner)
}

// ConfirmChangeMinerWorker is a paid mutator transaction binding the contract method 0xd16ac933.
//
// Solidity: function confirmChangeMinerWorker(uint64 miner) returns()
func (_Agent *AgentTransactorSession) ConfirmChangeMinerWorker(miner uint64) (*types.Transaction, error) {
	return _Agent.Contract.ConfirmChangeMinerWorker(&_Agent.TransactOpts, miner)
}

// DecommissionAgent is a paid mutator transaction binding the contract method 0x11756607.
//
// Solidity: function decommissionAgent(address _newAgent) returns()
func (_Agent *AgentTransactor) DecommissionAgent(opts *bind.TransactOpts, _newAgent common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "decommissionAgent", _newAgent)
}

// DecommissionAgent is a paid mutator transaction binding the contract method 0x11756607.
//
// Solidity: function decommissionAgent(address _newAgent) returns()
func (_Agent *AgentSession) DecommissionAgent(_newAgent common.Address) (*types.Transaction, error) {
	return _Agent.Contract.DecommissionAgent(&_Agent.TransactOpts, _newAgent)
}

// DecommissionAgent is a paid mutator transaction binding the contract method 0x11756607.
//
// Solidity: function decommissionAgent(address _newAgent) returns()
func (_Agent *AgentTransactorSession) DecommissionAgent(_newAgent common.Address) (*types.Transaction, error) {
	return _Agent.Contract.DecommissionAgent(&_Agent.TransactOpts, _newAgent)
}

// MigrateMiner is a paid mutator transaction binding the contract method 0x81b0149a.
//
// Solidity: function migrateMiner(uint64 miner) returns()
func (_Agent *AgentTransactor) MigrateMiner(opts *bind.TransactOpts, miner uint64) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "migrateMiner", miner)
}

// MigrateMiner is a paid mutator transaction binding the contract method 0x81b0149a.
//
// Solidity: function migrateMiner(uint64 miner) returns()
func (_Agent *AgentSession) MigrateMiner(miner uint64) (*types.Transaction, error) {
	return _Agent.Contract.MigrateMiner(&_Agent.TransactOpts, miner)
}

// MigrateMiner is a paid mutator transaction binding the contract method 0x81b0149a.
//
// Solidity: function migrateMiner(uint64 miner) returns()
func (_Agent *AgentTransactorSession) MigrateMiner(miner uint64) (*types.Transaction, error) {
	return _Agent.Contract.MigrateMiner(&_Agent.TransactOpts, miner)
}

// Pay is a paid mutator transaction binding the contract method 0x9ba7e551.
//
// Solidity: function pay(uint256 poolID, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_Agent *AgentTransactor) Pay(opts *bind.TransactOpts, poolID *big.Int, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "pay", poolID, sc)
}

// Pay is a paid mutator transaction binding the contract method 0x9ba7e551.
//
// Solidity: function pay(uint256 poolID, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_Agent *AgentSession) Pay(poolID *big.Int, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.Pay(&_Agent.TransactOpts, poolID, sc)
}

// Pay is a paid mutator transaction binding the contract method 0x9ba7e551.
//
// Solidity: function pay(uint256 poolID, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_Agent *AgentTransactorSession) Pay(poolID *big.Int, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.Pay(&_Agent.TransactOpts, poolID, sc)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0xf794a923.
//
// Solidity: function prepareMinerForLiquidation(uint64 miner, uint64 liquidator) returns()
func (_Agent *AgentTransactor) PrepareMinerForLiquidation(opts *bind.TransactOpts, miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "prepareMinerForLiquidation", miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0xf794a923.
//
// Solidity: function prepareMinerForLiquidation(uint64 miner, uint64 liquidator) returns()
func (_Agent *AgentSession) PrepareMinerForLiquidation(miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _Agent.Contract.PrepareMinerForLiquidation(&_Agent.TransactOpts, miner, liquidator)
}

// PrepareMinerForLiquidation is a paid mutator transaction binding the contract method 0xf794a923.
//
// Solidity: function prepareMinerForLiquidation(uint64 miner, uint64 liquidator) returns()
func (_Agent *AgentTransactorSession) PrepareMinerForLiquidation(miner uint64, liquidator uint64) (*types.Transaction, error) {
	return _Agent.Contract.PrepareMinerForLiquidation(&_Agent.TransactOpts, miner, liquidator)
}

// PullFunds is a paid mutator transaction binding the contract method 0x7b8686ed.
//
// Solidity: function pullFunds(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactor) PullFunds(opts *bind.TransactOpts, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "pullFunds", sc)
}

// PullFunds is a paid mutator transaction binding the contract method 0x7b8686ed.
//
// Solidity: function pullFunds(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentSession) PullFunds(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.PullFunds(&_Agent.TransactOpts, sc)
}

// PullFunds is a paid mutator transaction binding the contract method 0x7b8686ed.
//
// Solidity: function pullFunds(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactorSession) PullFunds(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.PullFunds(&_Agent.TransactOpts, sc)
}

// PushFunds is a paid mutator transaction binding the contract method 0xf1991b54.
//
// Solidity: function pushFunds(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactor) PushFunds(opts *bind.TransactOpts, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "pushFunds", sc)
}

// PushFunds is a paid mutator transaction binding the contract method 0xf1991b54.
//
// Solidity: function pushFunds(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentSession) PushFunds(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.PushFunds(&_Agent.TransactOpts, sc)
}

// PushFunds is a paid mutator transaction binding the contract method 0xf1991b54.
//
// Solidity: function pushFunds(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactorSession) PushFunds(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.PushFunds(&_Agent.TransactOpts, sc)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_Agent *AgentTransactor) RefreshRoutes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "refreshRoutes")
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_Agent *AgentSession) RefreshRoutes() (*types.Transaction, error) {
	return _Agent.Contract.RefreshRoutes(&_Agent.TransactOpts)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_Agent *AgentTransactorSession) RefreshRoutes() (*types.Transaction, error) {
	return _Agent.Contract.RefreshRoutes(&_Agent.TransactOpts)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x86423c9a.
//
// Solidity: function removeMiner(uint64 newMinerOwner, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactor) RemoveMiner(opts *bind.TransactOpts, newMinerOwner uint64, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "removeMiner", newMinerOwner, sc)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x86423c9a.
//
// Solidity: function removeMiner(uint64 newMinerOwner, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentSession) RemoveMiner(newMinerOwner uint64, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.RemoveMiner(&_Agent.TransactOpts, newMinerOwner, sc)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x86423c9a.
//
// Solidity: function removeMiner(uint64 newMinerOwner, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactorSession) RemoveMiner(newMinerOwner uint64, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.RemoveMiner(&_Agent.TransactOpts, newMinerOwner, sc)
}

// SetAdministration is a paid mutator transaction binding the contract method 0x9e4d74cc.
//
// Solidity: function setAdministration(address _administration) returns()
func (_Agent *AgentTransactor) SetAdministration(opts *bind.TransactOpts, _administration common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "setAdministration", _administration)
}

// SetAdministration is a paid mutator transaction binding the contract method 0x9e4d74cc.
//
// Solidity: function setAdministration(address _administration) returns()
func (_Agent *AgentSession) SetAdministration(_administration common.Address) (*types.Transaction, error) {
	return _Agent.Contract.SetAdministration(&_Agent.TransactOpts, _administration)
}

// SetAdministration is a paid mutator transaction binding the contract method 0x9e4d74cc.
//
// Solidity: function setAdministration(address _administration) returns()
func (_Agent *AgentTransactorSession) SetAdministration(_administration common.Address) (*types.Transaction, error) {
	return _Agent.Contract.SetAdministration(&_Agent.TransactOpts, _administration)
}

// SetAdoRequestKey is a paid mutator transaction binding the contract method 0x1aba0c5c.
//
// Solidity: function setAdoRequestKey(address _newKey) returns()
func (_Agent *AgentTransactor) SetAdoRequestKey(opts *bind.TransactOpts, _newKey common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "setAdoRequestKey", _newKey)
}

// SetAdoRequestKey is a paid mutator transaction binding the contract method 0x1aba0c5c.
//
// Solidity: function setAdoRequestKey(address _newKey) returns()
func (_Agent *AgentSession) SetAdoRequestKey(_newKey common.Address) (*types.Transaction, error) {
	return _Agent.Contract.SetAdoRequestKey(&_Agent.TransactOpts, _newKey)
}

// SetAdoRequestKey is a paid mutator transaction binding the contract method 0x1aba0c5c.
//
// Solidity: function setAdoRequestKey(address _newKey) returns()
func (_Agent *AgentTransactorSession) SetAdoRequestKey(_newKey common.Address) (*types.Transaction, error) {
	return _Agent.Contract.SetAdoRequestKey(&_Agent.TransactOpts, _newKey)
}

// SetFaulty is a paid mutator transaction binding the contract method 0xa7180449.
//
// Solidity: function setFaulty() returns()
func (_Agent *AgentTransactor) SetFaulty(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "setFaulty")
}

// SetFaulty is a paid mutator transaction binding the contract method 0xa7180449.
//
// Solidity: function setFaulty() returns()
func (_Agent *AgentSession) SetFaulty() (*types.Transaction, error) {
	return _Agent.Contract.SetFaulty(&_Agent.TransactOpts)
}

// SetFaulty is a paid mutator transaction binding the contract method 0xa7180449.
//
// Solidity: function setFaulty() returns()
func (_Agent *AgentTransactorSession) SetFaulty() (*types.Transaction, error) {
	return _Agent.Contract.SetFaulty(&_Agent.TransactOpts)
}

// SetInDefault is a paid mutator transaction binding the contract method 0x615d9211.
//
// Solidity: function setInDefault() returns()
func (_Agent *AgentTransactor) SetInDefault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "setInDefault")
}

// SetInDefault is a paid mutator transaction binding the contract method 0x615d9211.
//
// Solidity: function setInDefault() returns()
func (_Agent *AgentSession) SetInDefault() (*types.Transaction, error) {
	return _Agent.Contract.SetInDefault(&_Agent.TransactOpts)
}

// SetInDefault is a paid mutator transaction binding the contract method 0x615d9211.
//
// Solidity: function setInDefault() returns()
func (_Agent *AgentTransactorSession) SetInDefault() (*types.Transaction, error) {
	return _Agent.Contract.SetInDefault(&_Agent.TransactOpts)
}

// SetRecovered is a paid mutator transaction binding the contract method 0x99efd9e1.
//
// Solidity: function setRecovered(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactor) SetRecovered(opts *bind.TransactOpts, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "setRecovered", sc)
}

// SetRecovered is a paid mutator transaction binding the contract method 0x99efd9e1.
//
// Solidity: function setRecovered(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentSession) SetRecovered(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.SetRecovered(&_Agent.TransactOpts, sc)
}

// SetRecovered is a paid mutator transaction binding the contract method 0x99efd9e1.
//
// Solidity: function setRecovered(((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactorSession) SetRecovered(sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.SetRecovered(&_Agent.TransactOpts, sc)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_Agent *AgentTransactor) TransferOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "transferOperator", newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_Agent *AgentSession) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Agent.Contract.TransferOperator(&_Agent.TransactOpts, newOperator)
}

// TransferOperator is a paid mutator transaction binding the contract method 0x29605e77.
//
// Solidity: function transferOperator(address newOperator) returns()
func (_Agent *AgentTransactorSession) TransferOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Agent.Contract.TransferOperator(&_Agent.TransactOpts, newOperator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Agent *AgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Agent *AgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Agent.Contract.TransferOwnership(&_Agent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Agent *AgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Agent.Contract.TransferOwnership(&_Agent.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x67a01c22.
//
// Solidity: function withdraw(address receiver, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactor) Withdraw(opts *bind.TransactOpts, receiver common.Address, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.contract.Transact(opts, "withdraw", receiver, sc)
}

// Withdraw is a paid mutator transaction binding the contract method 0x67a01c22.
//
// Solidity: function withdraw(address receiver, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentSession) Withdraw(receiver common.Address, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.Withdraw(&_Agent.TransactOpts, receiver, sc)
}

// Withdraw is a paid mutator transaction binding the contract method 0x67a01c22.
//
// Solidity: function withdraw(address receiver, ((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes),uint8,bytes32,bytes32) sc) returns()
func (_Agent *AgentTransactorSession) Withdraw(receiver common.Address, sc SignedCredential) (*types.Transaction, error) {
	return _Agent.Contract.Withdraw(&_Agent.TransactOpts, receiver, sc)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Agent *AgentTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Agent.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Agent *AgentSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Agent.Contract.Fallback(&_Agent.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Agent *AgentTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Agent.Contract.Fallback(&_Agent.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Agent *AgentTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Agent.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Agent *AgentSession) Receive() (*types.Transaction, error) {
	return _Agent.Contract.Receive(&_Agent.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Agent *AgentTransactorSession) Receive() (*types.Transaction, error) {
	return _Agent.Contract.Receive(&_Agent.TransactOpts)
}
