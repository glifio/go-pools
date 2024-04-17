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

// PreStakeMetaData contains all meta data concerning the PreStake contract.
var PreStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"},{\"internalType\":\"contractIPoolToken\",\"name\":\"_poolToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApprovePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approvePoolToTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertFILtoWFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PreStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use PreStakeMetaData.ABI instead.
var PreStakeABI = PreStakeMetaData.ABI

// PreStake is an auto generated Go binding around an Ethereum contract.
type PreStake struct {
	PreStakeCaller     // Read-only binding to the contract
	PreStakeTransactor // Write-only binding to the contract
	PreStakeFilterer   // Log filterer for contract events
}

// PreStakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreStakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreStakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreStakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreStakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreStakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreStakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreStakeSession struct {
	Contract     *PreStake         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreStakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreStakeCallerSession struct {
	Contract *PreStakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PreStakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreStakeTransactorSession struct {
	Contract     *PreStakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PreStakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreStakeRaw struct {
	Contract *PreStake // Generic contract binding to access the raw methods on
}

// PreStakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreStakeCallerRaw struct {
	Contract *PreStakeCaller // Generic read-only contract binding to access the raw methods on
}

// PreStakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreStakeTransactorRaw struct {
	Contract *PreStakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreStake creates a new instance of PreStake, bound to a specific deployed contract.
func NewPreStake(address common.Address, backend bind.ContractBackend) (*PreStake, error) {
	contract, err := bindPreStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreStake{PreStakeCaller: PreStakeCaller{contract: contract}, PreStakeTransactor: PreStakeTransactor{contract: contract}, PreStakeFilterer: PreStakeFilterer{contract: contract}}, nil
}

// NewPreStakeCaller creates a new read-only instance of PreStake, bound to a specific deployed contract.
func NewPreStakeCaller(address common.Address, caller bind.ContractCaller) (*PreStakeCaller, error) {
	contract, err := bindPreStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreStakeCaller{contract: contract}, nil
}

// NewPreStakeTransactor creates a new write-only instance of PreStake, bound to a specific deployed contract.
func NewPreStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*PreStakeTransactor, error) {
	contract, err := bindPreStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreStakeTransactor{contract: contract}, nil
}

// NewPreStakeFilterer creates a new log filterer instance of PreStake, bound to a specific deployed contract.
func NewPreStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*PreStakeFilterer, error) {
	contract, err := bindPreStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreStakeFilterer{contract: contract}, nil
}

// bindPreStake binds a generic wrapper to an already deployed contract.
func bindPreStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PreStakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreStake *PreStakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreStake.Contract.PreStakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreStake *PreStakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreStake.Contract.PreStakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreStake *PreStakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreStake.Contract.PreStakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreStake *PreStakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PreStake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreStake *PreStakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreStake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreStake *PreStakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreStake.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PreStake *PreStakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PreStake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PreStake *PreStakeSession) Owner() (common.Address, error) {
	return _PreStake.Contract.Owner(&_PreStake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PreStake *PreStakeCallerSession) Owner() (common.Address, error) {
	return _PreStake.Contract.Owner(&_PreStake.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PreStake *PreStakeCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PreStake.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PreStake *PreStakeSession) PendingOwner() (common.Address, error) {
	return _PreStake.Contract.PendingOwner(&_PreStake.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PreStake *PreStakeCallerSession) PendingOwner() (common.Address, error) {
	return _PreStake.Contract.PendingOwner(&_PreStake.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_PreStake *PreStakeCaller) TotalValueLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PreStake.contract.Call(opts, &out, "totalValueLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_PreStake *PreStakeSession) TotalValueLocked() (*big.Int, error) {
	return _PreStake.Contract.TotalValueLocked(&_PreStake.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_PreStake *PreStakeCallerSession) TotalValueLocked() (*big.Int, error) {
	return _PreStake.Contract.TotalValueLocked(&_PreStake.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PreStake *PreStakeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PreStake *PreStakeSession) AcceptOwnership() (*types.Transaction, error) {
	return _PreStake.Contract.AcceptOwnership(&_PreStake.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PreStake *PreStakeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PreStake.Contract.AcceptOwnership(&_PreStake.TransactOpts)
}

// ApprovePoolToTransfer is a paid mutator transaction binding the contract method 0x7ea4e0e5.
//
// Solidity: function approvePoolToTransfer(address pool, uint256 amount) returns()
func (_PreStake *PreStakeTransactor) ApprovePoolToTransfer(opts *bind.TransactOpts, pool common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "approvePoolToTransfer", pool, amount)
}

// ApprovePoolToTransfer is a paid mutator transaction binding the contract method 0x7ea4e0e5.
//
// Solidity: function approvePoolToTransfer(address pool, uint256 amount) returns()
func (_PreStake *PreStakeSession) ApprovePoolToTransfer(pool common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PreStake.Contract.ApprovePoolToTransfer(&_PreStake.TransactOpts, pool, amount)
}

// ApprovePoolToTransfer is a paid mutator transaction binding the contract method 0x7ea4e0e5.
//
// Solidity: function approvePoolToTransfer(address pool, uint256 amount) returns()
func (_PreStake *PreStakeTransactorSession) ApprovePoolToTransfer(pool common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PreStake.Contract.ApprovePoolToTransfer(&_PreStake.TransactOpts, pool, amount)
}

// ConvertFILtoWFIL is a paid mutator transaction binding the contract method 0x1a3ece7f.
//
// Solidity: function convertFILtoWFIL() returns()
func (_PreStake *PreStakeTransactor) ConvertFILtoWFIL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "convertFILtoWFIL")
}

// ConvertFILtoWFIL is a paid mutator transaction binding the contract method 0x1a3ece7f.
//
// Solidity: function convertFILtoWFIL() returns()
func (_PreStake *PreStakeSession) ConvertFILtoWFIL() (*types.Transaction, error) {
	return _PreStake.Contract.ConvertFILtoWFIL(&_PreStake.TransactOpts)
}

// ConvertFILtoWFIL is a paid mutator transaction binding the contract method 0x1a3ece7f.
//
// Solidity: function convertFILtoWFIL() returns()
func (_PreStake *PreStakeTransactorSession) ConvertFILtoWFIL() (*types.Transaction, error) {
	return _PreStake.Contract.ConvertFILtoWFIL(&_PreStake.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns()
func (_PreStake *PreStakeTransactor) Deposit(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "deposit", account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns()
func (_PreStake *PreStakeSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PreStake.Contract.Deposit(&_PreStake.TransactOpts, account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns()
func (_PreStake *PreStakeTransactorSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PreStake.Contract.Deposit(&_PreStake.TransactOpts, account, amount)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address account) payable returns()
func (_PreStake *PreStakeTransactor) Deposit0(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "deposit0", account)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address account) payable returns()
func (_PreStake *PreStakeSession) Deposit0(account common.Address) (*types.Transaction, error) {
	return _PreStake.Contract.Deposit0(&_PreStake.TransactOpts, account)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address account) payable returns()
func (_PreStake *PreStakeTransactorSession) Deposit0(account common.Address) (*types.Transaction, error) {
	return _PreStake.Contract.Deposit0(&_PreStake.TransactOpts, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PreStake *PreStakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PreStake *PreStakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PreStake.Contract.TransferOwnership(&_PreStake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PreStake *PreStakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PreStake.Contract.TransferOwnership(&_PreStake.TransactOpts, newOwner)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PreStake *PreStakeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _PreStake.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PreStake *PreStakeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PreStake.Contract.Fallback(&_PreStake.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_PreStake *PreStakeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _PreStake.Contract.Fallback(&_PreStake.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PreStake *PreStakeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreStake.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PreStake *PreStakeSession) Receive() (*types.Transaction, error) {
	return _PreStake.Contract.Receive(&_PreStake.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_PreStake *PreStakeTransactorSession) Receive() (*types.Transaction, error) {
	return _PreStake.Contract.Receive(&_PreStake.TransactOpts)
}

// PreStakeApprovePoolIterator is returned from FilterApprovePool and is used to iterate over the raw logs and unpacked data for ApprovePool events raised by the PreStake contract.
type PreStakeApprovePoolIterator struct {
	Event *PreStakeApprovePool // Event containing the contract specifics and raw log

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
func (it *PreStakeApprovePoolIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreStakeApprovePool)
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
		it.Event = new(PreStakeApprovePool)
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
func (it *PreStakeApprovePoolIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreStakeApprovePoolIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreStakeApprovePool represents a ApprovePool event raised by the PreStake contract.
type PreStakeApprovePool struct {
	Pool   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApprovePool is a free log retrieval operation binding the contract event 0x349c23e02899ddce96ca5c007ea92efb73b9dc4db1d0ffb87de71f87ac54a953.
//
// Solidity: event ApprovePool(address indexed pool, uint256 amount)
func (_PreStake *PreStakeFilterer) FilterApprovePool(opts *bind.FilterOpts, pool []common.Address) (*PreStakeApprovePoolIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PreStake.contract.FilterLogs(opts, "ApprovePool", poolRule)
	if err != nil {
		return nil, err
	}
	return &PreStakeApprovePoolIterator{contract: _PreStake.contract, event: "ApprovePool", logs: logs, sub: sub}, nil
}

// WatchApprovePool is a free log subscription operation binding the contract event 0x349c23e02899ddce96ca5c007ea92efb73b9dc4db1d0ffb87de71f87ac54a953.
//
// Solidity: event ApprovePool(address indexed pool, uint256 amount)
func (_PreStake *PreStakeFilterer) WatchApprovePool(opts *bind.WatchOpts, sink chan<- *PreStakeApprovePool, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _PreStake.contract.WatchLogs(opts, "ApprovePool", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreStakeApprovePool)
				if err := _PreStake.contract.UnpackLog(event, "ApprovePool", log); err != nil {
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

// ParseApprovePool is a log parse operation binding the contract event 0x349c23e02899ddce96ca5c007ea92efb73b9dc4db1d0ffb87de71f87ac54a953.
//
// Solidity: event ApprovePool(address indexed pool, uint256 amount)
func (_PreStake *PreStakeFilterer) ParseApprovePool(log types.Log) (*PreStakeApprovePool, error) {
	event := new(PreStakeApprovePool)
	if err := _PreStake.contract.UnpackLog(event, "ApprovePool", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreStakeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the PreStake contract.
type PreStakeDepositIterator struct {
	Event *PreStakeDeposit // Event containing the contract specifics and raw log

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
func (it *PreStakeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreStakeDeposit)
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
		it.Event = new(PreStakeDeposit)
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
func (it *PreStakeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreStakeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreStakeDeposit represents a Deposit event raised by the PreStake contract.
type PreStakeDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_PreStake *PreStakeFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*PreStakeDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreStake.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &PreStakeDepositIterator{contract: _PreStake.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_PreStake *PreStakeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *PreStakeDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PreStake.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreStakeDeposit)
				if err := _PreStake.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_PreStake *PreStakeFilterer) ParseDeposit(log types.Log) (*PreStakeDeposit, error) {
	event := new(PreStakeDeposit)
	if err := _PreStake.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreStakeOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the PreStake contract.
type PreStakeOwnershipTransferStartedIterator struct {
	Event *PreStakeOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *PreStakeOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreStakeOwnershipTransferStarted)
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
		it.Event = new(PreStakeOwnershipTransferStarted)
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
func (it *PreStakeOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreStakeOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreStakeOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the PreStake contract.
type PreStakeOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PreStakeOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PreStake.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PreStakeOwnershipTransferStartedIterator{contract: _PreStake.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *PreStakeOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PreStake.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreStakeOwnershipTransferStarted)
				if err := _PreStake.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) ParseOwnershipTransferStarted(log types.Log) (*PreStakeOwnershipTransferStarted, error) {
	event := new(PreStakeOwnershipTransferStarted)
	if err := _PreStake.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PreStakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PreStake contract.
type PreStakeOwnershipTransferredIterator struct {
	Event *PreStakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PreStakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreStakeOwnershipTransferred)
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
		it.Event = new(PreStakeOwnershipTransferred)
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
func (it *PreStakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreStakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreStakeOwnershipTransferred represents a OwnershipTransferred event raised by the PreStake contract.
type PreStakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PreStakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PreStake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PreStakeOwnershipTransferredIterator{contract: _PreStake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PreStakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PreStake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreStakeOwnershipTransferred)
				if err := _PreStake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) ParseOwnershipTransferred(log types.Log) (*PreStakeOwnershipTransferred, error) {
	event := new(PreStakeOwnershipTransferred)
	if err := _PreStake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
