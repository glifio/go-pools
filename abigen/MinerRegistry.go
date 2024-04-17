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

// MinerRegistryMetaData contains all meta data concerning the MinerRegistry contract.
var MinerRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"contractIAgentFactory\",\"name\":\"_agentFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"AddMiner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"RemoveMiner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"addMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getMiner\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"minerRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"minersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"miner\",\"type\":\"uint64\"}],\"name\":\"removeMiner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MinerRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MinerRegistryMetaData.ABI instead.
var MinerRegistryABI = MinerRegistryMetaData.ABI

// MinerRegistry is an auto generated Go binding around an Ethereum contract.
type MinerRegistry struct {
	MinerRegistryCaller     // Read-only binding to the contract
	MinerRegistryTransactor // Write-only binding to the contract
	MinerRegistryFilterer   // Log filterer for contract events
}

// MinerRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinerRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinerRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinerRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinerRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinerRegistrySession struct {
	Contract     *MinerRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinerRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinerRegistryCallerSession struct {
	Contract *MinerRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MinerRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinerRegistryTransactorSession struct {
	Contract     *MinerRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MinerRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinerRegistryRaw struct {
	Contract *MinerRegistry // Generic contract binding to access the raw methods on
}

// MinerRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinerRegistryCallerRaw struct {
	Contract *MinerRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MinerRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinerRegistryTransactorRaw struct {
	Contract *MinerRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinerRegistry creates a new instance of MinerRegistry, bound to a specific deployed contract.
func NewMinerRegistry(address common.Address, backend bind.ContractBackend) (*MinerRegistry, error) {
	contract, err := bindMinerRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinerRegistry{MinerRegistryCaller: MinerRegistryCaller{contract: contract}, MinerRegistryTransactor: MinerRegistryTransactor{contract: contract}, MinerRegistryFilterer: MinerRegistryFilterer{contract: contract}}, nil
}

// NewMinerRegistryCaller creates a new read-only instance of MinerRegistry, bound to a specific deployed contract.
func NewMinerRegistryCaller(address common.Address, caller bind.ContractCaller) (*MinerRegistryCaller, error) {
	contract, err := bindMinerRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinerRegistryCaller{contract: contract}, nil
}

// NewMinerRegistryTransactor creates a new write-only instance of MinerRegistry, bound to a specific deployed contract.
func NewMinerRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MinerRegistryTransactor, error) {
	contract, err := bindMinerRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinerRegistryTransactor{contract: contract}, nil
}

// NewMinerRegistryFilterer creates a new log filterer instance of MinerRegistry, bound to a specific deployed contract.
func NewMinerRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MinerRegistryFilterer, error) {
	contract, err := bindMinerRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinerRegistryFilterer{contract: contract}, nil
}

// bindMinerRegistry binds a generic wrapper to an already deployed contract.
func bindMinerRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MinerRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerRegistry *MinerRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinerRegistry.Contract.MinerRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerRegistry *MinerRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerRegistry.Contract.MinerRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerRegistry *MinerRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerRegistry.Contract.MinerRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinerRegistry *MinerRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MinerRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinerRegistry *MinerRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinerRegistry *MinerRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinerRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetMiner is a free data retrieval call binding the contract method 0x3d874d9f.
//
// Solidity: function getMiner(uint256 agentID, uint256 index) view returns(uint64)
func (_MinerRegistry *MinerRegistryCaller) GetMiner(opts *bind.CallOpts, agentID *big.Int, index *big.Int) (uint64, error) {
	var out []interface{}
	err := _MinerRegistry.contract.Call(opts, &out, "getMiner", agentID, index)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetMiner is a free data retrieval call binding the contract method 0x3d874d9f.
//
// Solidity: function getMiner(uint256 agentID, uint256 index) view returns(uint64)
func (_MinerRegistry *MinerRegistrySession) GetMiner(agentID *big.Int, index *big.Int) (uint64, error) {
	return _MinerRegistry.Contract.GetMiner(&_MinerRegistry.CallOpts, agentID, index)
}

// GetMiner is a free data retrieval call binding the contract method 0x3d874d9f.
//
// Solidity: function getMiner(uint256 agentID, uint256 index) view returns(uint64)
func (_MinerRegistry *MinerRegistryCallerSession) GetMiner(agentID *big.Int, index *big.Int) (uint64, error) {
	return _MinerRegistry.Contract.GetMiner(&_MinerRegistry.CallOpts, agentID, index)
}

// MinerRegistered is a free data retrieval call binding the contract method 0x37d3362a.
//
// Solidity: function minerRegistered(uint256 agentID, uint64 miner) view returns(bool)
func (_MinerRegistry *MinerRegistryCaller) MinerRegistered(opts *bind.CallOpts, agentID *big.Int, miner uint64) (bool, error) {
	var out []interface{}
	err := _MinerRegistry.contract.Call(opts, &out, "minerRegistered", agentID, miner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinerRegistered is a free data retrieval call binding the contract method 0x37d3362a.
//
// Solidity: function minerRegistered(uint256 agentID, uint64 miner) view returns(bool)
func (_MinerRegistry *MinerRegistrySession) MinerRegistered(agentID *big.Int, miner uint64) (bool, error) {
	return _MinerRegistry.Contract.MinerRegistered(&_MinerRegistry.CallOpts, agentID, miner)
}

// MinerRegistered is a free data retrieval call binding the contract method 0x37d3362a.
//
// Solidity: function minerRegistered(uint256 agentID, uint64 miner) view returns(bool)
func (_MinerRegistry *MinerRegistryCallerSession) MinerRegistered(agentID *big.Int, miner uint64) (bool, error) {
	return _MinerRegistry.Contract.MinerRegistered(&_MinerRegistry.CallOpts, agentID, miner)
}

// MinersCount is a free data retrieval call binding the contract method 0xe3441365.
//
// Solidity: function minersCount(uint256 agentID) view returns(uint256)
func (_MinerRegistry *MinerRegistryCaller) MinersCount(opts *bind.CallOpts, agentID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MinerRegistry.contract.Call(opts, &out, "minersCount", agentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinersCount is a free data retrieval call binding the contract method 0xe3441365.
//
// Solidity: function minersCount(uint256 agentID) view returns(uint256)
func (_MinerRegistry *MinerRegistrySession) MinersCount(agentID *big.Int) (*big.Int, error) {
	return _MinerRegistry.Contract.MinersCount(&_MinerRegistry.CallOpts, agentID)
}

// MinersCount is a free data retrieval call binding the contract method 0xe3441365.
//
// Solidity: function minersCount(uint256 agentID) view returns(uint256)
func (_MinerRegistry *MinerRegistryCallerSession) MinersCount(agentID *big.Int) (*big.Int, error) {
	return _MinerRegistry.Contract.MinersCount(&_MinerRegistry.CallOpts, agentID)
}

// AddMiner is a paid mutator transaction binding the contract method 0x60c9beaf.
//
// Solidity: function addMiner(uint256 agentID, uint64 miner) returns()
func (_MinerRegistry *MinerRegistryTransactor) AddMiner(opts *bind.TransactOpts, agentID *big.Int, miner uint64) (*types.Transaction, error) {
	return _MinerRegistry.contract.Transact(opts, "addMiner", agentID, miner)
}

// AddMiner is a paid mutator transaction binding the contract method 0x60c9beaf.
//
// Solidity: function addMiner(uint256 agentID, uint64 miner) returns()
func (_MinerRegistry *MinerRegistrySession) AddMiner(agentID *big.Int, miner uint64) (*types.Transaction, error) {
	return _MinerRegistry.Contract.AddMiner(&_MinerRegistry.TransactOpts, agentID, miner)
}

// AddMiner is a paid mutator transaction binding the contract method 0x60c9beaf.
//
// Solidity: function addMiner(uint256 agentID, uint64 miner) returns()
func (_MinerRegistry *MinerRegistryTransactorSession) AddMiner(agentID *big.Int, miner uint64) (*types.Transaction, error) {
	return _MinerRegistry.Contract.AddMiner(&_MinerRegistry.TransactOpts, agentID, miner)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_MinerRegistry *MinerRegistryTransactor) RefreshRoutes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinerRegistry.contract.Transact(opts, "refreshRoutes")
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_MinerRegistry *MinerRegistrySession) RefreshRoutes() (*types.Transaction, error) {
	return _MinerRegistry.Contract.RefreshRoutes(&_MinerRegistry.TransactOpts)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_MinerRegistry *MinerRegistryTransactorSession) RefreshRoutes() (*types.Transaction, error) {
	return _MinerRegistry.Contract.RefreshRoutes(&_MinerRegistry.TransactOpts)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x3cbe1a26.
//
// Solidity: function removeMiner(uint256 agentID, uint64 miner) returns()
func (_MinerRegistry *MinerRegistryTransactor) RemoveMiner(opts *bind.TransactOpts, agentID *big.Int, miner uint64) (*types.Transaction, error) {
	return _MinerRegistry.contract.Transact(opts, "removeMiner", agentID, miner)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x3cbe1a26.
//
// Solidity: function removeMiner(uint256 agentID, uint64 miner) returns()
func (_MinerRegistry *MinerRegistrySession) RemoveMiner(agentID *big.Int, miner uint64) (*types.Transaction, error) {
	return _MinerRegistry.Contract.RemoveMiner(&_MinerRegistry.TransactOpts, agentID, miner)
}

// RemoveMiner is a paid mutator transaction binding the contract method 0x3cbe1a26.
//
// Solidity: function removeMiner(uint256 agentID, uint64 miner) returns()
func (_MinerRegistry *MinerRegistryTransactorSession) RemoveMiner(agentID *big.Int, miner uint64) (*types.Transaction, error) {
	return _MinerRegistry.Contract.RemoveMiner(&_MinerRegistry.TransactOpts, agentID, miner)
}

// MinerRegistryAddMinerIterator is returned from FilterAddMiner and is used to iterate over the raw logs and unpacked data for AddMiner events raised by the MinerRegistry contract.
type MinerRegistryAddMinerIterator struct {
	Event *MinerRegistryAddMiner // Event containing the contract specifics and raw log

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
func (it *MinerRegistryAddMinerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinerRegistryAddMiner)
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
		it.Event = new(MinerRegistryAddMiner)
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
func (it *MinerRegistryAddMinerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinerRegistryAddMinerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinerRegistryAddMiner represents a AddMiner event raised by the MinerRegistry contract.
type MinerRegistryAddMiner struct {
	Agent common.Address
	Miner uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAddMiner is a free log retrieval operation binding the contract event 0x16dc2cbf4fac6621252a619403d8f179259201c83ee0cfb9c1210f36f49e8765.
//
// Solidity: event AddMiner(address indexed agent, uint64 indexed miner)
func (_MinerRegistry *MinerRegistryFilterer) FilterAddMiner(opts *bind.FilterOpts, agent []common.Address, miner []uint64) (*MinerRegistryAddMinerIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _MinerRegistry.contract.FilterLogs(opts, "AddMiner", agentRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &MinerRegistryAddMinerIterator{contract: _MinerRegistry.contract, event: "AddMiner", logs: logs, sub: sub}, nil
}

// WatchAddMiner is a free log subscription operation binding the contract event 0x16dc2cbf4fac6621252a619403d8f179259201c83ee0cfb9c1210f36f49e8765.
//
// Solidity: event AddMiner(address indexed agent, uint64 indexed miner)
func (_MinerRegistry *MinerRegistryFilterer) WatchAddMiner(opts *bind.WatchOpts, sink chan<- *MinerRegistryAddMiner, agent []common.Address, miner []uint64) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _MinerRegistry.contract.WatchLogs(opts, "AddMiner", agentRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinerRegistryAddMiner)
				if err := _MinerRegistry.contract.UnpackLog(event, "AddMiner", log); err != nil {
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

// ParseAddMiner is a log parse operation binding the contract event 0x16dc2cbf4fac6621252a619403d8f179259201c83ee0cfb9c1210f36f49e8765.
//
// Solidity: event AddMiner(address indexed agent, uint64 indexed miner)
func (_MinerRegistry *MinerRegistryFilterer) ParseAddMiner(log types.Log) (*MinerRegistryAddMiner, error) {
	event := new(MinerRegistryAddMiner)
	if err := _MinerRegistry.contract.UnpackLog(event, "AddMiner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MinerRegistryRemoveMinerIterator is returned from FilterRemoveMiner and is used to iterate over the raw logs and unpacked data for RemoveMiner events raised by the MinerRegistry contract.
type MinerRegistryRemoveMinerIterator struct {
	Event *MinerRegistryRemoveMiner // Event containing the contract specifics and raw log

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
func (it *MinerRegistryRemoveMinerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinerRegistryRemoveMiner)
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
		it.Event = new(MinerRegistryRemoveMiner)
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
func (it *MinerRegistryRemoveMinerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinerRegistryRemoveMinerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinerRegistryRemoveMiner represents a RemoveMiner event raised by the MinerRegistry contract.
type MinerRegistryRemoveMiner struct {
	Agent common.Address
	Miner uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRemoveMiner is a free log retrieval operation binding the contract event 0xb9d2bc5634927c75cf1ec9264544a54f7da51cba5fe5da9185b635e8a82f3c43.
//
// Solidity: event RemoveMiner(address indexed agent, uint64 indexed miner)
func (_MinerRegistry *MinerRegistryFilterer) FilterRemoveMiner(opts *bind.FilterOpts, agent []common.Address, miner []uint64) (*MinerRegistryRemoveMinerIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _MinerRegistry.contract.FilterLogs(opts, "RemoveMiner", agentRule, minerRule)
	if err != nil {
		return nil, err
	}
	return &MinerRegistryRemoveMinerIterator{contract: _MinerRegistry.contract, event: "RemoveMiner", logs: logs, sub: sub}, nil
}

// WatchRemoveMiner is a free log subscription operation binding the contract event 0xb9d2bc5634927c75cf1ec9264544a54f7da51cba5fe5da9185b635e8a82f3c43.
//
// Solidity: event RemoveMiner(address indexed agent, uint64 indexed miner)
func (_MinerRegistry *MinerRegistryFilterer) WatchRemoveMiner(opts *bind.WatchOpts, sink chan<- *MinerRegistryRemoveMiner, agent []common.Address, miner []uint64) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var minerRule []interface{}
	for _, minerItem := range miner {
		minerRule = append(minerRule, minerItem)
	}

	logs, sub, err := _MinerRegistry.contract.WatchLogs(opts, "RemoveMiner", agentRule, minerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinerRegistryRemoveMiner)
				if err := _MinerRegistry.contract.UnpackLog(event, "RemoveMiner", log); err != nil {
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

// ParseRemoveMiner is a log parse operation binding the contract event 0xb9d2bc5634927c75cf1ec9264544a54f7da51cba5fe5da9185b635e8a82f3c43.
//
// Solidity: event RemoveMiner(address indexed agent, uint64 indexed miner)
func (_MinerRegistry *MinerRegistryFilterer) ParseRemoveMiner(log types.Log) (*MinerRegistryRemoveMiner, error) {
	event := new(MinerRegistryRemoveMiner)
	if err := _MinerRegistry.contract.UnpackLog(event, "RemoveMiner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
