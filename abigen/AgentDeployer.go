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

// AgentDeployerMetaData contains all meta data concerning the AgentDeployer contract.
var AgentDeployerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adoRequestKey\",\"type\":\"address\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AgentDeployerABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentDeployerMetaData.ABI instead.
var AgentDeployerABI = AgentDeployerMetaData.ABI

// AgentDeployer is an auto generated Go binding around an Ethereum contract.
type AgentDeployer struct {
	AgentDeployerCaller     // Read-only binding to the contract
	AgentDeployerTransactor // Write-only binding to the contract
	AgentDeployerFilterer   // Log filterer for contract events
}

// AgentDeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentDeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentDeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentDeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentDeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentDeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentDeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentDeployerSession struct {
	Contract     *AgentDeployer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentDeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentDeployerCallerSession struct {
	Contract *AgentDeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AgentDeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentDeployerTransactorSession struct {
	Contract     *AgentDeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AgentDeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentDeployerRaw struct {
	Contract *AgentDeployer // Generic contract binding to access the raw methods on
}

// AgentDeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentDeployerCallerRaw struct {
	Contract *AgentDeployerCaller // Generic read-only contract binding to access the raw methods on
}

// AgentDeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentDeployerTransactorRaw struct {
	Contract *AgentDeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentDeployer creates a new instance of AgentDeployer, bound to a specific deployed contract.
func NewAgentDeployer(address common.Address, backend bind.ContractBackend) (*AgentDeployer, error) {
	contract, err := bindAgentDeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentDeployer{AgentDeployerCaller: AgentDeployerCaller{contract: contract}, AgentDeployerTransactor: AgentDeployerTransactor{contract: contract}, AgentDeployerFilterer: AgentDeployerFilterer{contract: contract}}, nil
}

// NewAgentDeployerCaller creates a new read-only instance of AgentDeployer, bound to a specific deployed contract.
func NewAgentDeployerCaller(address common.Address, caller bind.ContractCaller) (*AgentDeployerCaller, error) {
	contract, err := bindAgentDeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentDeployerCaller{contract: contract}, nil
}

// NewAgentDeployerTransactor creates a new write-only instance of AgentDeployer, bound to a specific deployed contract.
func NewAgentDeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentDeployerTransactor, error) {
	contract, err := bindAgentDeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentDeployerTransactor{contract: contract}, nil
}

// NewAgentDeployerFilterer creates a new log filterer instance of AgentDeployer, bound to a specific deployed contract.
func NewAgentDeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentDeployerFilterer, error) {
	contract, err := bindAgentDeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentDeployerFilterer{contract: contract}, nil
}

// bindAgentDeployer binds a generic wrapper to an already deployed contract.
func bindAgentDeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentDeployerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentDeployer *AgentDeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentDeployer.Contract.AgentDeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentDeployer *AgentDeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentDeployer.Contract.AgentDeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentDeployer *AgentDeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentDeployer.Contract.AgentDeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentDeployer *AgentDeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentDeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentDeployer *AgentDeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentDeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentDeployer *AgentDeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentDeployer.Contract.contract.Transact(opts, method, params...)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_AgentDeployer *AgentDeployerCaller) Version(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _AgentDeployer.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_AgentDeployer *AgentDeployerSession) Version() (uint8, error) {
	return _AgentDeployer.Contract.Version(&_AgentDeployer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint8)
func (_AgentDeployer *AgentDeployerCallerSession) Version() (uint8, error) {
	return _AgentDeployer.Contract.Version(&_AgentDeployer.CallOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0x09b01d0f.
//
// Solidity: function deploy(address router, uint256 agentId, address owner, address operator, address adoRequestKey) returns(address agent)
func (_AgentDeployer *AgentDeployerTransactor) Deploy(opts *bind.TransactOpts, router common.Address, agentId *big.Int, owner common.Address, operator common.Address, adoRequestKey common.Address) (*types.Transaction, error) {
	return _AgentDeployer.contract.Transact(opts, "deploy", router, agentId, owner, operator, adoRequestKey)
}

// Deploy is a paid mutator transaction binding the contract method 0x09b01d0f.
//
// Solidity: function deploy(address router, uint256 agentId, address owner, address operator, address adoRequestKey) returns(address agent)
func (_AgentDeployer *AgentDeployerSession) Deploy(router common.Address, agentId *big.Int, owner common.Address, operator common.Address, adoRequestKey common.Address) (*types.Transaction, error) {
	return _AgentDeployer.Contract.Deploy(&_AgentDeployer.TransactOpts, router, agentId, owner, operator, adoRequestKey)
}

// Deploy is a paid mutator transaction binding the contract method 0x09b01d0f.
//
// Solidity: function deploy(address router, uint256 agentId, address owner, address operator, address adoRequestKey) returns(address agent)
func (_AgentDeployer *AgentDeployerTransactorSession) Deploy(router common.Address, agentId *big.Int, owner common.Address, operator common.Address, adoRequestKey common.Address) (*types.Transaction, error) {
	return _AgentDeployer.Contract.Deploy(&_AgentDeployer.TransactOpts, router, agentId, owner, operator, adoRequestKey)
}
