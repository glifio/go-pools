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

// UpgradeToV2MetaData contains all meta data concerning the UpgradeToV2 contract.
var UpgradeToV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oldPoolBalance\",\"type\":\"uint256\"}],\"name\":\"OldPoolNotShutDownProperly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"varName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"PoolVarMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"varName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"PoolVarMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"name\":\"refreshProtocolRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAuth\",\"name\":\"_auth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agentPolice\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newInterestAccrued\",\"type\":\"uint256\"}],\"name\":\"verifyTotalAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UpgradeToV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use UpgradeToV2MetaData.ABI instead.
var UpgradeToV2ABI = UpgradeToV2MetaData.ABI

// UpgradeToV2 is an auto generated Go binding around an Ethereum contract.
type UpgradeToV2 struct {
	UpgradeToV2Caller     // Read-only binding to the contract
	UpgradeToV2Transactor // Write-only binding to the contract
	UpgradeToV2Filterer   // Log filterer for contract events
}

// UpgradeToV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type UpgradeToV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeToV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type UpgradeToV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeToV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UpgradeToV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UpgradeToV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UpgradeToV2Session struct {
	Contract     *UpgradeToV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UpgradeToV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UpgradeToV2CallerSession struct {
	Contract *UpgradeToV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UpgradeToV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UpgradeToV2TransactorSession struct {
	Contract     *UpgradeToV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UpgradeToV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type UpgradeToV2Raw struct {
	Contract *UpgradeToV2 // Generic contract binding to access the raw methods on
}

// UpgradeToV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UpgradeToV2CallerRaw struct {
	Contract *UpgradeToV2Caller // Generic read-only contract binding to access the raw methods on
}

// UpgradeToV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UpgradeToV2TransactorRaw struct {
	Contract *UpgradeToV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewUpgradeToV2 creates a new instance of UpgradeToV2, bound to a specific deployed contract.
func NewUpgradeToV2(address common.Address, backend bind.ContractBackend) (*UpgradeToV2, error) {
	contract, err := bindUpgradeToV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UpgradeToV2{UpgradeToV2Caller: UpgradeToV2Caller{contract: contract}, UpgradeToV2Transactor: UpgradeToV2Transactor{contract: contract}, UpgradeToV2Filterer: UpgradeToV2Filterer{contract: contract}}, nil
}

// NewUpgradeToV2Caller creates a new read-only instance of UpgradeToV2, bound to a specific deployed contract.
func NewUpgradeToV2Caller(address common.Address, caller bind.ContractCaller) (*UpgradeToV2Caller, error) {
	contract, err := bindUpgradeToV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeToV2Caller{contract: contract}, nil
}

// NewUpgradeToV2Transactor creates a new write-only instance of UpgradeToV2, bound to a specific deployed contract.
func NewUpgradeToV2Transactor(address common.Address, transactor bind.ContractTransactor) (*UpgradeToV2Transactor, error) {
	contract, err := bindUpgradeToV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UpgradeToV2Transactor{contract: contract}, nil
}

// NewUpgradeToV2Filterer creates a new log filterer instance of UpgradeToV2, bound to a specific deployed contract.
func NewUpgradeToV2Filterer(address common.Address, filterer bind.ContractFilterer) (*UpgradeToV2Filterer, error) {
	contract, err := bindUpgradeToV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UpgradeToV2Filterer{contract: contract}, nil
}

// bindUpgradeToV2 binds a generic wrapper to an already deployed contract.
func bindUpgradeToV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UpgradeToV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeToV2 *UpgradeToV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeToV2.Contract.UpgradeToV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeToV2 *UpgradeToV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.UpgradeToV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeToV2 *UpgradeToV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.UpgradeToV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UpgradeToV2 *UpgradeToV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UpgradeToV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UpgradeToV2 *UpgradeToV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UpgradeToV2 *UpgradeToV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeToV2 *UpgradeToV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeToV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeToV2 *UpgradeToV2Session) Owner() (common.Address, error) {
	return _UpgradeToV2.Contract.Owner(&_UpgradeToV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UpgradeToV2 *UpgradeToV2CallerSession) Owner() (common.Address, error) {
	return _UpgradeToV2.Contract.Owner(&_UpgradeToV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UpgradeToV2 *UpgradeToV2Caller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UpgradeToV2.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UpgradeToV2 *UpgradeToV2Session) PendingOwner() (common.Address, error) {
	return _UpgradeToV2.Contract.PendingOwner(&_UpgradeToV2.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_UpgradeToV2 *UpgradeToV2CallerSession) PendingOwner() (common.Address, error) {
	return _UpgradeToV2.Contract.PendingOwner(&_UpgradeToV2.CallOpts)
}

// VerifyTotalAssets is a free data retrieval call binding the contract method 0x8287aa2d.
//
// Solidity: function verifyTotalAssets(uint256 newInterestAccrued) view returns(bool)
func (_UpgradeToV2 *UpgradeToV2Caller) VerifyTotalAssets(opts *bind.CallOpts, newInterestAccrued *big.Int) (bool, error) {
	var out []interface{}
	err := _UpgradeToV2.contract.Call(opts, &out, "verifyTotalAssets", newInterestAccrued)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyTotalAssets is a free data retrieval call binding the contract method 0x8287aa2d.
//
// Solidity: function verifyTotalAssets(uint256 newInterestAccrued) view returns(bool)
func (_UpgradeToV2 *UpgradeToV2Session) VerifyTotalAssets(newInterestAccrued *big.Int) (bool, error) {
	return _UpgradeToV2.Contract.VerifyTotalAssets(&_UpgradeToV2.CallOpts, newInterestAccrued)
}

// VerifyTotalAssets is a free data retrieval call binding the contract method 0x8287aa2d.
//
// Solidity: function verifyTotalAssets(uint256 newInterestAccrued) view returns(bool)
func (_UpgradeToV2 *UpgradeToV2CallerSession) VerifyTotalAssets(newInterestAccrued *big.Int) (bool, error) {
	return _UpgradeToV2.Contract.VerifyTotalAssets(&_UpgradeToV2.CallOpts, newInterestAccrued)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UpgradeToV2 *UpgradeToV2Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeToV2.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UpgradeToV2 *UpgradeToV2Session) AcceptOwnership() (*types.Transaction, error) {
	return _UpgradeToV2.Contract.AcceptOwnership(&_UpgradeToV2.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_UpgradeToV2 *UpgradeToV2TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _UpgradeToV2.Contract.AcceptOwnership(&_UpgradeToV2.TransactOpts)
}

// RefreshProtocolRoutes is a paid mutator transaction binding the contract method 0x7abe3913.
//
// Solidity: function refreshProtocolRoutes(address[] agents) returns()
func (_UpgradeToV2 *UpgradeToV2Transactor) RefreshProtocolRoutes(opts *bind.TransactOpts, agents []common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.contract.Transact(opts, "refreshProtocolRoutes", agents)
}

// RefreshProtocolRoutes is a paid mutator transaction binding the contract method 0x7abe3913.
//
// Solidity: function refreshProtocolRoutes(address[] agents) returns()
func (_UpgradeToV2 *UpgradeToV2Session) RefreshProtocolRoutes(agents []common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.RefreshProtocolRoutes(&_UpgradeToV2.TransactOpts, agents)
}

// RefreshProtocolRoutes is a paid mutator transaction binding the contract method 0x7abe3913.
//
// Solidity: function refreshProtocolRoutes(address[] agents) returns()
func (_UpgradeToV2 *UpgradeToV2TransactorSession) RefreshProtocolRoutes(agents []common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.RefreshProtocolRoutes(&_UpgradeToV2.TransactOpts, agents)
}

// SetOwner is a paid mutator transaction binding the contract method 0x299a7bcc.
//
// Solidity: function setOwner(address _auth, address _newOwner) returns()
func (_UpgradeToV2 *UpgradeToV2Transactor) SetOwner(opts *bind.TransactOpts, _auth common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.contract.Transact(opts, "setOwner", _auth, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x299a7bcc.
//
// Solidity: function setOwner(address _auth, address _newOwner) returns()
func (_UpgradeToV2 *UpgradeToV2Session) SetOwner(_auth common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.SetOwner(&_UpgradeToV2.TransactOpts, _auth, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x299a7bcc.
//
// Solidity: function setOwner(address _auth, address _newOwner) returns()
func (_UpgradeToV2 *UpgradeToV2TransactorSession) SetOwner(_auth common.Address, _newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.SetOwner(&_UpgradeToV2.TransactOpts, _auth, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeToV2 *UpgradeToV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeToV2 *UpgradeToV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.TransferOwnership(&_UpgradeToV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UpgradeToV2 *UpgradeToV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.TransferOwnership(&_UpgradeToV2.TransactOpts, newOwner)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address agentPolice, address pool) payable returns()
func (_UpgradeToV2 *UpgradeToV2Transactor) Upgrade(opts *bind.TransactOpts, agentPolice common.Address, pool common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.contract.Transact(opts, "upgrade", agentPolice, pool)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address agentPolice, address pool) payable returns()
func (_UpgradeToV2 *UpgradeToV2Session) Upgrade(agentPolice common.Address, pool common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.Upgrade(&_UpgradeToV2.TransactOpts, agentPolice, pool)
}

// Upgrade is a paid mutator transaction binding the contract method 0x99a88ec4.
//
// Solidity: function upgrade(address agentPolice, address pool) payable returns()
func (_UpgradeToV2 *UpgradeToV2TransactorSession) Upgrade(agentPolice common.Address, pool common.Address) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.Upgrade(&_UpgradeToV2.TransactOpts, agentPolice, pool)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeToV2 *UpgradeToV2Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _UpgradeToV2.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeToV2 *UpgradeToV2Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.Fallback(&_UpgradeToV2.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_UpgradeToV2 *UpgradeToV2TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _UpgradeToV2.Contract.Fallback(&_UpgradeToV2.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeToV2 *UpgradeToV2Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UpgradeToV2.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeToV2 *UpgradeToV2Session) Receive() (*types.Transaction, error) {
	return _UpgradeToV2.Contract.Receive(&_UpgradeToV2.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UpgradeToV2 *UpgradeToV2TransactorSession) Receive() (*types.Transaction, error) {
	return _UpgradeToV2.Contract.Receive(&_UpgradeToV2.TransactOpts)
}
