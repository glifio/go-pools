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

// SimpleRampMetaData contains all meta data concerning the SimpleRamp contract.
var SimpleRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_poolID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"burnIFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iFIL\",\"outputs\":[{\"internalType\":\"contractIPoolToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshExtern\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalExitDemand\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wFIL\",\"outputs\":[{\"internalType\":\"contractIWFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SimpleRampABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleRampMetaData.ABI instead.
var SimpleRampABI = SimpleRampMetaData.ABI

// SimpleRamp is an auto generated Go binding around an Ethereum contract.
type SimpleRamp struct {
	SimpleRampCaller     // Read-only binding to the contract
	SimpleRampTransactor // Write-only binding to the contract
	SimpleRampFilterer   // Log filterer for contract events
}

// SimpleRampCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleRampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleRampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleRampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleRampSession struct {
	Contract     *SimpleRamp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleRampCallerSession struct {
	Contract *SimpleRampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimpleRampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleRampTransactorSession struct {
	Contract     *SimpleRampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleRampRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRampRaw struct {
	Contract *SimpleRamp // Generic contract binding to access the raw methods on
}

// SimpleRampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleRampCallerRaw struct {
	Contract *SimpleRampCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleRampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleRampTransactorRaw struct {
	Contract *SimpleRampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleRamp creates a new instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRamp(address common.Address, backend bind.ContractBackend) (*SimpleRamp, error) {
	contract, err := bindSimpleRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleRamp{SimpleRampCaller: SimpleRampCaller{contract: contract}, SimpleRampTransactor: SimpleRampTransactor{contract: contract}, SimpleRampFilterer: SimpleRampFilterer{contract: contract}}, nil
}

// NewSimpleRampCaller creates a new read-only instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRampCaller(address common.Address, caller bind.ContractCaller) (*SimpleRampCaller, error) {
	contract, err := bindSimpleRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRampCaller{contract: contract}, nil
}

// NewSimpleRampTransactor creates a new write-only instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRampTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleRampTransactor, error) {
	contract, err := bindSimpleRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRampTransactor{contract: contract}, nil
}

// NewSimpleRampFilterer creates a new log filterer instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRampFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleRampFilterer, error) {
	contract, err := bindSimpleRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleRampFilterer{contract: contract}, nil
}

// bindSimpleRamp binds a generic wrapper to an already deployed contract.
func bindSimpleRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRamp *SimpleRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRamp.Contract.SimpleRampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRamp *SimpleRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.Contract.SimpleRampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRamp *SimpleRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRamp.Contract.SimpleRampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRamp *SimpleRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRamp *SimpleRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRamp *SimpleRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRamp.Contract.contract.Transact(opts, method, params...)
}

// IFIL is a free data retrieval call binding the contract method 0x6be82b78.
//
// Solidity: function iFIL() view returns(address)
func (_SimpleRamp *SimpleRampCaller) IFIL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "iFIL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IFIL is a free data retrieval call binding the contract method 0x6be82b78.
//
// Solidity: function iFIL() view returns(address)
func (_SimpleRamp *SimpleRampSession) IFIL() (common.Address, error) {
	return _SimpleRamp.Contract.IFIL(&_SimpleRamp.CallOpts)
}

// IFIL is a free data retrieval call binding the contract method 0x6be82b78.
//
// Solidity: function iFIL() view returns(address)
func (_SimpleRamp *SimpleRampCallerSession) IFIL() (common.Address, error) {
	return _SimpleRamp.Contract.IFIL(&_SimpleRamp.CallOpts)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address account) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCaller) MaxRedeem(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "maxRedeem", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address account) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampSession) MaxRedeem(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxRedeem(&_SimpleRamp.CallOpts, account)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address account) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCallerSession) MaxRedeem(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxRedeem(&_SimpleRamp.CallOpts, account)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address account) view returns(uint256)
func (_SimpleRamp *SimpleRampCaller) MaxWithdraw(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "maxWithdraw", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address account) view returns(uint256)
func (_SimpleRamp *SimpleRampSession) MaxWithdraw(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxWithdraw(&_SimpleRamp.CallOpts, account)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address account) view returns(uint256)
func (_SimpleRamp *SimpleRampCallerSession) MaxWithdraw(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxWithdraw(&_SimpleRamp.CallOpts, account)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SimpleRamp *SimpleRampCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SimpleRamp *SimpleRampSession) Pool() (common.Address, error) {
	return _SimpleRamp.Contract.Pool(&_SimpleRamp.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SimpleRamp *SimpleRampCallerSession) Pool() (common.Address, error) {
	return _SimpleRamp.Contract.Pool(&_SimpleRamp.CallOpts)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_SimpleRamp *SimpleRampCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_SimpleRamp *SimpleRampSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewRedeem(&_SimpleRamp.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_SimpleRamp *SimpleRampCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewRedeem(&_SimpleRamp.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewWithdraw(&_SimpleRamp.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewWithdraw(&_SimpleRamp.CallOpts, assets)
}

// TotalExitDemand is a free data retrieval call binding the contract method 0x0fede599.
//
// Solidity: function totalExitDemand() view returns(uint256)
func (_SimpleRamp *SimpleRampCaller) TotalExitDemand(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "totalExitDemand")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalExitDemand is a free data retrieval call binding the contract method 0x0fede599.
//
// Solidity: function totalExitDemand() view returns(uint256)
func (_SimpleRamp *SimpleRampSession) TotalExitDemand() (*big.Int, error) {
	return _SimpleRamp.Contract.TotalExitDemand(&_SimpleRamp.CallOpts)
}

// TotalExitDemand is a free data retrieval call binding the contract method 0x0fede599.
//
// Solidity: function totalExitDemand() view returns(uint256)
func (_SimpleRamp *SimpleRampCallerSession) TotalExitDemand() (*big.Int, error) {
	return _SimpleRamp.Contract.TotalExitDemand(&_SimpleRamp.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0xbef44256.
//
// Solidity: function wFIL() view returns(address)
func (_SimpleRamp *SimpleRampCaller) WFIL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "wFIL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WFIL is a free data retrieval call binding the contract method 0xbef44256.
//
// Solidity: function wFIL() view returns(address)
func (_SimpleRamp *SimpleRampSession) WFIL() (common.Address, error) {
	return _SimpleRamp.Contract.WFIL(&_SimpleRamp.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0xbef44256.
//
// Solidity: function wFIL() view returns(address)
func (_SimpleRamp *SimpleRampCallerSession) WFIL() (common.Address, error) {
	return _SimpleRamp.Contract.WFIL(&_SimpleRamp.CallOpts)
}

// BurnIFIL is a paid mutator transaction binding the contract method 0xe8c9c4b7.
//
// Solidity: function burnIFIL() returns()
func (_SimpleRamp *SimpleRampTransactor) BurnIFIL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "burnIFIL")
}

// BurnIFIL is a paid mutator transaction binding the contract method 0xe8c9c4b7.
//
// Solidity: function burnIFIL() returns()
func (_SimpleRamp *SimpleRampSession) BurnIFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.BurnIFIL(&_SimpleRamp.TransactOpts)
}

// BurnIFIL is a paid mutator transaction binding the contract method 0xe8c9c4b7.
//
// Solidity: function burnIFIL() returns()
func (_SimpleRamp *SimpleRampTransactorSession) BurnIFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.BurnIFIL(&_SimpleRamp.TransactOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address , uint256 amount) returns()
func (_SimpleRamp *SimpleRampTransactor) Distribute(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "distribute", arg0, amount)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address , uint256 amount) returns()
func (_SimpleRamp *SimpleRampSession) Distribute(arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Distribute(&_SimpleRamp.TransactOpts, arg0, amount)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address , uint256 amount) returns()
func (_SimpleRamp *SimpleRampTransactorSession) Distribute(arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Distribute(&_SimpleRamp.TransactOpts, arg0, amount)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x3eb569c6.
//
// Solidity: function recoverFIL() returns()
func (_SimpleRamp *SimpleRampTransactor) RecoverFIL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "recoverFIL")
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x3eb569c6.
//
// Solidity: function recoverFIL() returns()
func (_SimpleRamp *SimpleRampSession) RecoverFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RecoverFIL(&_SimpleRamp.TransactOpts)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x3eb569c6.
//
// Solidity: function recoverFIL() returns()
func (_SimpleRamp *SimpleRampTransactorSession) RecoverFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RecoverFIL(&_SimpleRamp.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "redeem", shares, receiver, owner, arg3)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Redeem(&_SimpleRamp.TransactOpts, shares, receiver, owner, arg3)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Redeem(&_SimpleRamp.TransactOpts, shares, receiver, owner, arg3)
}

// RefreshExtern is a paid mutator transaction binding the contract method 0x9adf1050.
//
// Solidity: function refreshExtern() returns()
func (_SimpleRamp *SimpleRampTransactor) RefreshExtern(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "refreshExtern")
}

// RefreshExtern is a paid mutator transaction binding the contract method 0x9adf1050.
//
// Solidity: function refreshExtern() returns()
func (_SimpleRamp *SimpleRampSession) RefreshExtern() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RefreshExtern(&_SimpleRamp.TransactOpts)
}

// RefreshExtern is a paid mutator transaction binding the contract method 0x9adf1050.
//
// Solidity: function refreshExtern() returns()
func (_SimpleRamp *SimpleRampTransactorSession) RefreshExtern() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RefreshExtern(&_SimpleRamp.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "withdraw", assets, receiver, owner, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Withdraw(&_SimpleRamp.TransactOpts, assets, receiver, owner, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Withdraw(&_SimpleRamp.TransactOpts, assets, receiver, owner, arg3)
}

// SimpleRampWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SimpleRamp contract.
type SimpleRampWithdrawIterator struct {
	Event *SimpleRampWithdraw // Event containing the contract specifics and raw log

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
func (it *SimpleRampWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRampWithdraw)
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
		it.Event = new(SimpleRampWithdraw)
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
func (it *SimpleRampWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRampWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRampWithdraw represents a Withdraw event raised by the SimpleRamp contract.
type SimpleRampWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_SimpleRamp *SimpleRampFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*SimpleRampWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleRamp.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleRampWithdrawIterator{contract: _SimpleRamp.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_SimpleRamp *SimpleRampFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SimpleRampWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleRamp.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRampWithdraw)
				if err := _SimpleRamp.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_SimpleRamp *SimpleRampFilterer) ParseWithdraw(log types.Log) (*SimpleRampWithdraw, error) {
	event := new(SimpleRampWithdraw)
	if err := _SimpleRamp.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
