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

// PublicGoodsDonatorMetaData contains all meta data concerning the PublicGoodsDonator contract.
var PublicGoodsDonatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIPreStake\",\"name\":\"_preStake\",\"type\":\"address\"},{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_liquidStakingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"donationAmount\",\"type\":\"uint256\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFunds\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"donationPercent\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"donationPercent\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preStake\",\"outputs\":[{\"internalType\":\"contractIPreStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PublicGoodsDonatorABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicGoodsDonatorMetaData.ABI instead.
var PublicGoodsDonatorABI = PublicGoodsDonatorMetaData.ABI

// PublicGoodsDonator is an auto generated Go binding around an Ethereum contract.
type PublicGoodsDonator struct {
	PublicGoodsDonatorCaller     // Read-only binding to the contract
	PublicGoodsDonatorTransactor // Write-only binding to the contract
	PublicGoodsDonatorFilterer   // Log filterer for contract events
}

// PublicGoodsDonatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicGoodsDonatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicGoodsDonatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicGoodsDonatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicGoodsDonatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicGoodsDonatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicGoodsDonatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicGoodsDonatorSession struct {
	Contract     *PublicGoodsDonator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PublicGoodsDonatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicGoodsDonatorCallerSession struct {
	Contract *PublicGoodsDonatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PublicGoodsDonatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicGoodsDonatorTransactorSession struct {
	Contract     *PublicGoodsDonatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PublicGoodsDonatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicGoodsDonatorRaw struct {
	Contract *PublicGoodsDonator // Generic contract binding to access the raw methods on
}

// PublicGoodsDonatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicGoodsDonatorCallerRaw struct {
	Contract *PublicGoodsDonatorCaller // Generic read-only contract binding to access the raw methods on
}

// PublicGoodsDonatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicGoodsDonatorTransactorRaw struct {
	Contract *PublicGoodsDonatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicGoodsDonator creates a new instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonator(address common.Address, backend bind.ContractBackend) (*PublicGoodsDonator, error) {
	contract, err := bindPublicGoodsDonator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonator{PublicGoodsDonatorCaller: PublicGoodsDonatorCaller{contract: contract}, PublicGoodsDonatorTransactor: PublicGoodsDonatorTransactor{contract: contract}, PublicGoodsDonatorFilterer: PublicGoodsDonatorFilterer{contract: contract}}, nil
}

// NewPublicGoodsDonatorCaller creates a new read-only instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonatorCaller(address common.Address, caller bind.ContractCaller) (*PublicGoodsDonatorCaller, error) {
	contract, err := bindPublicGoodsDonator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorCaller{contract: contract}, nil
}

// NewPublicGoodsDonatorTransactor creates a new write-only instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonatorTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicGoodsDonatorTransactor, error) {
	contract, err := bindPublicGoodsDonator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorTransactor{contract: contract}, nil
}

// NewPublicGoodsDonatorFilterer creates a new log filterer instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonatorFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicGoodsDonatorFilterer, error) {
	contract, err := bindPublicGoodsDonator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorFilterer{contract: contract}, nil
}

// bindPublicGoodsDonator binds a generic wrapper to an already deployed contract.
func bindPublicGoodsDonator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicGoodsDonatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicGoodsDonator *PublicGoodsDonatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicGoodsDonator.Contract.PublicGoodsDonatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicGoodsDonator *PublicGoodsDonatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.PublicGoodsDonatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicGoodsDonator *PublicGoodsDonatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.PublicGoodsDonatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicGoodsDonator *PublicGoodsDonatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicGoodsDonator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicGoodsDonator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorSession) Owner() (common.Address, error) {
	return _PublicGoodsDonator.Contract.Owner(&_PublicGoodsDonator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCallerSession) Owner() (common.Address, error) {
	return _PublicGoodsDonator.Contract.Owner(&_PublicGoodsDonator.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicGoodsDonator.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorSession) PendingOwner() (common.Address, error) {
	return _PublicGoodsDonator.Contract.PendingOwner(&_PublicGoodsDonator.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCallerSession) PendingOwner() (common.Address, error) {
	return _PublicGoodsDonator.Contract.PendingOwner(&_PublicGoodsDonator.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCaller) PreStake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicGoodsDonator.contract.Call(opts, &out, "preStake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorSession) PreStake() (common.Address, error) {
	return _PublicGoodsDonator.Contract.PreStake(&_PublicGoodsDonator.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCallerSession) PreStake() (common.Address, error) {
	return _PublicGoodsDonator.Contract.PreStake(&_PublicGoodsDonator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.AcceptOwnership(&_PublicGoodsDonator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.AcceptOwnership(&_PublicGoodsDonator.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recipient, uint256 amount, uint256 donationPercent) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "deposit", recipient, amount, donationPercent)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recipient, uint256 amount, uint256 donationPercent) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) Deposit(recipient common.Address, amount *big.Int, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit(&_PublicGoodsDonator.TransactOpts, recipient, amount, donationPercent)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recipient, uint256 amount, uint256 donationPercent) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) Deposit(recipient common.Address, amount *big.Int, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit(&_PublicGoodsDonator.TransactOpts, recipient, amount, donationPercent)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 donationPercent) payable returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) Deposit0(opts *bind.TransactOpts, recipient common.Address, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "deposit0", recipient, donationPercent)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 donationPercent) payable returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) Deposit0(recipient common.Address, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit0(&_PublicGoodsDonator.TransactOpts, recipient, donationPercent)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 donationPercent) payable returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) Deposit0(recipient common.Address, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit0(&_PublicGoodsDonator.TransactOpts, recipient, donationPercent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.TransferOwnership(&_PublicGoodsDonator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.TransferOwnership(&_PublicGoodsDonator.TransactOpts, newOwner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) WithdrawFunds() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.WithdrawFunds(&_PublicGoodsDonator.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.WithdrawFunds(&_PublicGoodsDonator.TransactOpts)
}

// PublicGoodsDonatorDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorDonateIterator struct {
	Event *PublicGoodsDonatorDonate // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorDonate)
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
		it.Event = new(PublicGoodsDonatorDonate)
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
func (it *PublicGoodsDonatorDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorDonate represents a Donate event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorDonate struct {
	Account        common.Address
	DonationAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0x0553260a2e46b0577270d8992db02d30856ca880144c72d6e9503760946aef13.
//
// Solidity: event Donate(address indexed account, uint256 donationAmount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterDonate(opts *bind.FilterOpts, account []common.Address) (*PublicGoodsDonatorDonateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "Donate", accountRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorDonateIterator{contract: _PublicGoodsDonator.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0x0553260a2e46b0577270d8992db02d30856ca880144c72d6e9503760946aef13.
//
// Solidity: event Donate(address indexed account, uint256 donationAmount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorDonate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "Donate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorDonate)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "Donate", log); err != nil {
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

// ParseDonate is a log parse operation binding the contract event 0x0553260a2e46b0577270d8992db02d30856ca880144c72d6e9503760946aef13.
//
// Solidity: event Donate(address indexed account, uint256 donationAmount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseDonate(log types.Log) (*PublicGoodsDonatorDonate, error) {
	event := new(PublicGoodsDonatorDonate)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicGoodsDonatorOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipTransferStartedIterator struct {
	Event *PublicGoodsDonatorOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorOwnershipTransferStarted)
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
		it.Event = new(PublicGoodsDonatorOwnershipTransferStarted)
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
func (it *PublicGoodsDonatorOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PublicGoodsDonatorOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorOwnershipTransferStartedIterator{contract: _PublicGoodsDonator.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorOwnershipTransferStarted)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseOwnershipTransferStarted(log types.Log) (*PublicGoodsDonatorOwnershipTransferStarted, error) {
	event := new(PublicGoodsDonatorOwnershipTransferStarted)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicGoodsDonatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipTransferredIterator struct {
	Event *PublicGoodsDonatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorOwnershipTransferred)
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
		it.Event = new(PublicGoodsDonatorOwnershipTransferred)
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
func (it *PublicGoodsDonatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorOwnershipTransferred represents a OwnershipTransferred event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PublicGoodsDonatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorOwnershipTransferredIterator{contract: _PublicGoodsDonator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorOwnershipTransferred)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseOwnershipTransferred(log types.Log) (*PublicGoodsDonatorOwnershipTransferred, error) {
	event := new(PublicGoodsDonatorOwnershipTransferred)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicGoodsDonatorWithdrawFundsIterator is returned from FilterWithdrawFunds and is used to iterate over the raw logs and unpacked data for WithdrawFunds events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorWithdrawFundsIterator struct {
	Event *PublicGoodsDonatorWithdrawFunds // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorWithdrawFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorWithdrawFunds)
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
		it.Event = new(PublicGoodsDonatorWithdrawFunds)
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
func (it *PublicGoodsDonatorWithdrawFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorWithdrawFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorWithdrawFunds represents a WithdrawFunds event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorWithdrawFunds struct {
	Wallet common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFunds is a free log retrieval operation binding the contract event 0x21901fa892c430ea8bd38b9390225ac8e67eac75ee10ffba16feefc539a288f9.
//
// Solidity: event WithdrawFunds(address indexed wallet, uint256 amount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterWithdrawFunds(opts *bind.FilterOpts, wallet []common.Address) (*PublicGoodsDonatorWithdrawFundsIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "WithdrawFunds", walletRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorWithdrawFundsIterator{contract: _PublicGoodsDonator.contract, event: "WithdrawFunds", logs: logs, sub: sub}, nil
}

// WatchWithdrawFunds is a free log subscription operation binding the contract event 0x21901fa892c430ea8bd38b9390225ac8e67eac75ee10ffba16feefc539a288f9.
//
// Solidity: event WithdrawFunds(address indexed wallet, uint256 amount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchWithdrawFunds(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorWithdrawFunds, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "WithdrawFunds", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorWithdrawFunds)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "WithdrawFunds", log); err != nil {
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

// ParseWithdrawFunds is a log parse operation binding the contract event 0x21901fa892c430ea8bd38b9390225ac8e67eac75ee10ffba16feefc539a288f9.
//
// Solidity: event WithdrawFunds(address indexed wallet, uint256 amount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseWithdrawFunds(log types.Log) (*PublicGoodsDonatorWithdrawFunds, error) {
	event := new(PublicGoodsDonatorWithdrawFunds)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "WithdrawFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
