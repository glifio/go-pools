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

// PingMetaData contains all meta data concerning the Ping contract.
var PingMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AmountWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Balance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"IsOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"}],\"name\":\"checkIsOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"}],\"name\":\"getBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsFEVM\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PingABI is the input ABI used to generate the binding from.
// Deprecated: Use PingMetaData.ABI instead.
var PingABI = PingMetaData.ABI

// Ping is an auto generated Go binding around an Ethereum contract.
type Ping struct {
	PingCaller     // Read-only binding to the contract
	PingTransactor // Write-only binding to the contract
	PingFilterer   // Log filterer for contract events
}

// PingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PingSession struct {
	Contract     *Ping             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PingCallerSession struct {
	Contract *PingCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PingTransactorSession struct {
	Contract     *PingTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PingRaw struct {
	Contract *Ping // Generic contract binding to access the raw methods on
}

// PingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PingCallerRaw struct {
	Contract *PingCaller // Generic read-only contract binding to access the raw methods on
}

// PingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PingTransactorRaw struct {
	Contract *PingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPing creates a new instance of Ping, bound to a specific deployed contract.
func NewPing(address common.Address, backend bind.ContractBackend) (*Ping, error) {
	contract, err := bindPing(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ping{PingCaller: PingCaller{contract: contract}, PingTransactor: PingTransactor{contract: contract}, PingFilterer: PingFilterer{contract: contract}}, nil
}

// NewPingCaller creates a new read-only instance of Ping, bound to a specific deployed contract.
func NewPingCaller(address common.Address, caller bind.ContractCaller) (*PingCaller, error) {
	contract, err := bindPing(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PingCaller{contract: contract}, nil
}

// NewPingTransactor creates a new write-only instance of Ping, bound to a specific deployed contract.
func NewPingTransactor(address common.Address, transactor bind.ContractTransactor) (*PingTransactor, error) {
	contract, err := bindPing(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PingTransactor{contract: contract}, nil
}

// NewPingFilterer creates a new log filterer instance of Ping, bound to a specific deployed contract.
func NewPingFilterer(address common.Address, filterer bind.ContractFilterer) (*PingFilterer, error) {
	contract, err := bindPing(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PingFilterer{contract: contract}, nil
}

// bindPing binds a generic wrapper to an already deployed contract.
func bindPing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ping *PingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ping.Contract.PingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ping *PingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ping.Contract.PingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ping *PingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ping.Contract.PingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ping *PingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ping.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ping *PingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ping.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ping *PingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ping.Contract.contract.Transact(opts, method, params...)
}

// GetIsFEVM is a free data retrieval call binding the contract method 0x4249538e.
//
// Solidity: function getIsFEVM() pure returns(bool)
func (_Ping *PingCaller) GetIsFEVM(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ping.contract.Call(opts, &out, "getIsFEVM")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsFEVM is a free data retrieval call binding the contract method 0x4249538e.
//
// Solidity: function getIsFEVM() pure returns(bool)
func (_Ping *PingSession) GetIsFEVM() (bool, error) {
	return _Ping.Contract.GetIsFEVM(&_Ping.CallOpts)
}

// GetIsFEVM is a free data retrieval call binding the contract method 0x4249538e.
//
// Solidity: function getIsFEVM() pure returns(bool)
func (_Ping *PingCallerSession) GetIsFEVM() (bool, error) {
	return _Ping.Contract.GetIsFEVM(&_Ping.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint256)
func (_Ping *PingCaller) GetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ping.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint256)
func (_Ping *PingSession) GetVersion() (*big.Int, error) {
	return _Ping.Contract.GetVersion(&_Ping.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(uint256)
func (_Ping *PingCallerSession) GetVersion() (*big.Int, error) {
	return _Ping.Contract.GetVersion(&_Ping.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf5914f5b.
//
// Solidity: function changeOwner(uint64 target) returns()
func (_Ping *PingTransactor) ChangeOwner(opts *bind.TransactOpts, target uint64) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "changeOwner", target)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf5914f5b.
//
// Solidity: function changeOwner(uint64 target) returns()
func (_Ping *PingSession) ChangeOwner(target uint64) (*types.Transaction, error) {
	return _Ping.Contract.ChangeOwner(&_Ping.TransactOpts, target)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf5914f5b.
//
// Solidity: function changeOwner(uint64 target) returns()
func (_Ping *PingTransactorSession) ChangeOwner(target uint64) (*types.Transaction, error) {
	return _Ping.Contract.ChangeOwner(&_Ping.TransactOpts, target)
}

// CheckIsOwner is a paid mutator transaction binding the contract method 0x7cb2e58e.
//
// Solidity: function checkIsOwner(uint64 target) returns()
func (_Ping *PingTransactor) CheckIsOwner(opts *bind.TransactOpts, target uint64) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "checkIsOwner", target)
}

// CheckIsOwner is a paid mutator transaction binding the contract method 0x7cb2e58e.
//
// Solidity: function checkIsOwner(uint64 target) returns()
func (_Ping *PingSession) CheckIsOwner(target uint64) (*types.Transaction, error) {
	return _Ping.Contract.CheckIsOwner(&_Ping.TransactOpts, target)
}

// CheckIsOwner is a paid mutator transaction binding the contract method 0x7cb2e58e.
//
// Solidity: function checkIsOwner(uint64 target) returns()
func (_Ping *PingTransactorSession) CheckIsOwner(target uint64) (*types.Transaction, error) {
	return _Ping.Contract.CheckIsOwner(&_Ping.TransactOpts, target)
}

// GetBalance is a paid mutator transaction binding the contract method 0x341b9e5a.
//
// Solidity: function getBalance(uint64 target) returns()
func (_Ping *PingTransactor) GetBalance(opts *bind.TransactOpts, target uint64) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "getBalance", target)
}

// GetBalance is a paid mutator transaction binding the contract method 0x341b9e5a.
//
// Solidity: function getBalance(uint64 target) returns()
func (_Ping *PingSession) GetBalance(target uint64) (*types.Transaction, error) {
	return _Ping.Contract.GetBalance(&_Ping.TransactOpts, target)
}

// GetBalance is a paid mutator transaction binding the contract method 0x341b9e5a.
//
// Solidity: function getBalance(uint64 target) returns()
func (_Ping *PingTransactorSession) GetBalance(target uint64) (*types.Transaction, error) {
	return _Ping.Contract.GetBalance(&_Ping.TransactOpts, target)
}

// Transfer is a paid mutator transaction binding the contract method 0x3823c189.
//
// Solidity: function transfer(uint64 target, uint256 amount) returns()
func (_Ping *PingTransactor) Transfer(opts *bind.TransactOpts, target uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "transfer", target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x3823c189.
//
// Solidity: function transfer(uint64 target, uint256 amount) returns()
func (_Ping *PingSession) Transfer(target uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.Transfer(&_Ping.TransactOpts, target, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x3823c189.
//
// Solidity: function transfer(uint64 target, uint256 amount) returns()
func (_Ping *PingTransactorSession) Transfer(target uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.Transfer(&_Ping.TransactOpts, target, amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x38792dd3.
//
// Solidity: function withdrawBalance(uint64 target, uint256 amount) returns()
func (_Ping *PingTransactor) WithdrawBalance(opts *bind.TransactOpts, target uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ping.contract.Transact(opts, "withdrawBalance", target, amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x38792dd3.
//
// Solidity: function withdrawBalance(uint64 target, uint256 amount) returns()
func (_Ping *PingSession) WithdrawBalance(target uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.WithdrawBalance(&_Ping.TransactOpts, target, amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x38792dd3.
//
// Solidity: function withdrawBalance(uint64 target, uint256 amount) returns()
func (_Ping *PingTransactorSession) WithdrawBalance(target uint64, amount *big.Int) (*types.Transaction, error) {
	return _Ping.Contract.WithdrawBalance(&_Ping.TransactOpts, target, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ping *PingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ping.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ping *PingSession) Receive() (*types.Transaction, error) {
	return _Ping.Contract.Receive(&_Ping.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ping *PingTransactorSession) Receive() (*types.Transaction, error) {
	return _Ping.Contract.Receive(&_Ping.TransactOpts)
}

// PingAmountWithdrawnIterator is returned from FilterAmountWithdrawn and is used to iterate over the raw logs and unpacked data for AmountWithdrawn events raised by the Ping contract.
type PingAmountWithdrawnIterator struct {
	Event *PingAmountWithdrawn // Event containing the contract specifics and raw log

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
func (it *PingAmountWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingAmountWithdrawn)
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
		it.Event = new(PingAmountWithdrawn)
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
func (it *PingAmountWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingAmountWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingAmountWithdrawn represents a AmountWithdrawn event raised by the Ping contract.
type PingAmountWithdrawn struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAmountWithdrawn is a free log retrieval operation binding the contract event 0x3e92aa4db3477ddd7c90d0ded7b985612c042e99cdeed9408343c650bdda73b9.
//
// Solidity: event AmountWithdrawn(uint256 arg0)
func (_Ping *PingFilterer) FilterAmountWithdrawn(opts *bind.FilterOpts) (*PingAmountWithdrawnIterator, error) {

	logs, sub, err := _Ping.contract.FilterLogs(opts, "AmountWithdrawn")
	if err != nil {
		return nil, err
	}
	return &PingAmountWithdrawnIterator{contract: _Ping.contract, event: "AmountWithdrawn", logs: logs, sub: sub}, nil
}

// WatchAmountWithdrawn is a free log subscription operation binding the contract event 0x3e92aa4db3477ddd7c90d0ded7b985612c042e99cdeed9408343c650bdda73b9.
//
// Solidity: event AmountWithdrawn(uint256 arg0)
func (_Ping *PingFilterer) WatchAmountWithdrawn(opts *bind.WatchOpts, sink chan<- *PingAmountWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Ping.contract.WatchLogs(opts, "AmountWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingAmountWithdrawn)
				if err := _Ping.contract.UnpackLog(event, "AmountWithdrawn", log); err != nil {
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

// ParseAmountWithdrawn is a log parse operation binding the contract event 0x3e92aa4db3477ddd7c90d0ded7b985612c042e99cdeed9408343c650bdda73b9.
//
// Solidity: event AmountWithdrawn(uint256 arg0)
func (_Ping *PingFilterer) ParseAmountWithdrawn(log types.Log) (*PingAmountWithdrawn, error) {
	event := new(PingAmountWithdrawn)
	if err := _Ping.contract.UnpackLog(event, "AmountWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingBalanceIterator is returned from FilterBalance and is used to iterate over the raw logs and unpacked data for Balance events raised by the Ping contract.
type PingBalanceIterator struct {
	Event *PingBalance // Event containing the contract specifics and raw log

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
func (it *PingBalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingBalance)
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
		it.Event = new(PingBalance)
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
func (it *PingBalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingBalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingBalance represents a Balance event raised by the Ping contract.
type PingBalance struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBalance is a free log retrieval operation binding the contract event 0xe8d947d7ebdd7b8b8fa2ad2022c7591418ac32d8c29c5f8a8fc9de52ffa54092.
//
// Solidity: event Balance(uint256 arg0)
func (_Ping *PingFilterer) FilterBalance(opts *bind.FilterOpts) (*PingBalanceIterator, error) {

	logs, sub, err := _Ping.contract.FilterLogs(opts, "Balance")
	if err != nil {
		return nil, err
	}
	return &PingBalanceIterator{contract: _Ping.contract, event: "Balance", logs: logs, sub: sub}, nil
}

// WatchBalance is a free log subscription operation binding the contract event 0xe8d947d7ebdd7b8b8fa2ad2022c7591418ac32d8c29c5f8a8fc9de52ffa54092.
//
// Solidity: event Balance(uint256 arg0)
func (_Ping *PingFilterer) WatchBalance(opts *bind.WatchOpts, sink chan<- *PingBalance) (event.Subscription, error) {

	logs, sub, err := _Ping.contract.WatchLogs(opts, "Balance")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingBalance)
				if err := _Ping.contract.UnpackLog(event, "Balance", log); err != nil {
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

// ParseBalance is a log parse operation binding the contract event 0xe8d947d7ebdd7b8b8fa2ad2022c7591418ac32d8c29c5f8a8fc9de52ffa54092.
//
// Solidity: event Balance(uint256 arg0)
func (_Ping *PingFilterer) ParseBalance(log types.Log) (*PingBalance, error) {
	event := new(PingBalance)
	if err := _Ping.contract.UnpackLog(event, "Balance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingIsOwnerIterator is returned from FilterIsOwner and is used to iterate over the raw logs and unpacked data for IsOwner events raised by the Ping contract.
type PingIsOwnerIterator struct {
	Event *PingIsOwner // Event containing the contract specifics and raw log

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
func (it *PingIsOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingIsOwner)
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
		it.Event = new(PingIsOwner)
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
func (it *PingIsOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingIsOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingIsOwner represents a IsOwner event raised by the Ping contract.
type PingIsOwner struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterIsOwner is a free log retrieval operation binding the contract event 0xf86afeaae77bccf494ba8205cc81f613ed84b3839d47a054b15a4ffd41ebd53e.
//
// Solidity: event IsOwner(bool arg0)
func (_Ping *PingFilterer) FilterIsOwner(opts *bind.FilterOpts) (*PingIsOwnerIterator, error) {

	logs, sub, err := _Ping.contract.FilterLogs(opts, "IsOwner")
	if err != nil {
		return nil, err
	}
	return &PingIsOwnerIterator{contract: _Ping.contract, event: "IsOwner", logs: logs, sub: sub}, nil
}

// WatchIsOwner is a free log subscription operation binding the contract event 0xf86afeaae77bccf494ba8205cc81f613ed84b3839d47a054b15a4ffd41ebd53e.
//
// Solidity: event IsOwner(bool arg0)
func (_Ping *PingFilterer) WatchIsOwner(opts *bind.WatchOpts, sink chan<- *PingIsOwner) (event.Subscription, error) {

	logs, sub, err := _Ping.contract.WatchLogs(opts, "IsOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingIsOwner)
				if err := _Ping.contract.UnpackLog(event, "IsOwner", log); err != nil {
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

// ParseIsOwner is a log parse operation binding the contract event 0xf86afeaae77bccf494ba8205cc81f613ed84b3839d47a054b15a4ffd41ebd53e.
//
// Solidity: event IsOwner(bool arg0)
func (_Ping *PingFilterer) ParseIsOwner(log types.Log) (*PingIsOwner, error) {
	event := new(PingIsOwner)
	if err := _Ping.contract.UnpackLog(event, "IsOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PingReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Ping contract.
type PingReceivedIterator struct {
	Event *PingReceived // Event containing the contract specifics and raw log

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
func (it *PingReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PingReceived)
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
		it.Event = new(PingReceived)
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
func (it *PingReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PingReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PingReceived represents a Received event raised by the Ping contract.
type PingReceived struct {
	Arg0 common.Address
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Ping *PingFilterer) FilterReceived(opts *bind.FilterOpts) (*PingReceivedIterator, error) {

	logs, sub, err := _Ping.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &PingReceivedIterator{contract: _Ping.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Ping *PingFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *PingReceived) (event.Subscription, error) {

	logs, sub, err := _Ping.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PingReceived)
				if err := _Ping.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address arg0, uint256 arg1)
func (_Ping *PingFilterer) ParseReceived(log types.Log) (*PingReceived, error) {
	event := new(PingReceived)
	if err := _Ping.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
