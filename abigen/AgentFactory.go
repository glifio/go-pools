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

// AgentFactoryMetaData contains all meta data concerning the AgentFactory contract.
var AgentFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"CreateAgent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"agents\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"adoRequestKey\",\"type\":\"address\"}],\"name\":\"create\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"isAgent\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"upgradeAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newAgent\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AgentFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentFactoryMetaData.ABI instead.
var AgentFactoryABI = AgentFactoryMetaData.ABI

// AgentFactory is an auto generated Go binding around an Ethereum contract.
type AgentFactory struct {
	AgentFactoryCaller     // Read-only binding to the contract
	AgentFactoryTransactor // Write-only binding to the contract
	AgentFactoryFilterer   // Log filterer for contract events
}

// AgentFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentFactorySession struct {
	Contract     *AgentFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AgentFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentFactoryCallerSession struct {
	Contract *AgentFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AgentFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentFactoryTransactorSession struct {
	Contract     *AgentFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AgentFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentFactoryRaw struct {
	Contract *AgentFactory // Generic contract binding to access the raw methods on
}

// AgentFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentFactoryCallerRaw struct {
	Contract *AgentFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AgentFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentFactoryTransactorRaw struct {
	Contract *AgentFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentFactory creates a new instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactory(address common.Address, backend bind.ContractBackend) (*AgentFactory, error) {
	contract, err := bindAgentFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentFactory{AgentFactoryCaller: AgentFactoryCaller{contract: contract}, AgentFactoryTransactor: AgentFactoryTransactor{contract: contract}, AgentFactoryFilterer: AgentFactoryFilterer{contract: contract}}, nil
}

// NewAgentFactoryCaller creates a new read-only instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactoryCaller(address common.Address, caller bind.ContractCaller) (*AgentFactoryCaller, error) {
	contract, err := bindAgentFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryCaller{contract: contract}, nil
}

// NewAgentFactoryTransactor creates a new write-only instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentFactoryTransactor, error) {
	contract, err := bindAgentFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryTransactor{contract: contract}, nil
}

// NewAgentFactoryFilterer creates a new log filterer instance of AgentFactory, bound to a specific deployed contract.
func NewAgentFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentFactoryFilterer, error) {
	contract, err := bindAgentFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryFilterer{contract: contract}, nil
}

// bindAgentFactory binds a generic wrapper to an already deployed contract.
func bindAgentFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AgentFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentFactory *AgentFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentFactory.Contract.AgentFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentFactory *AgentFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentFactory.Contract.AgentFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentFactory *AgentFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentFactory.Contract.AgentFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentFactory *AgentFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentFactory *AgentFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentFactory *AgentFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentFactory.Contract.contract.Transact(opts, method, params...)
}

// AgentCount is a free data retrieval call binding the contract method 0xb7dc1284.
//
// Solidity: function agentCount() view returns(uint256)
func (_AgentFactory *AgentFactoryCaller) AgentCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AgentFactory.contract.Call(opts, &out, "agentCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AgentCount is a free data retrieval call binding the contract method 0xb7dc1284.
//
// Solidity: function agentCount() view returns(uint256)
func (_AgentFactory *AgentFactorySession) AgentCount() (*big.Int, error) {
	return _AgentFactory.Contract.AgentCount(&_AgentFactory.CallOpts)
}

// AgentCount is a free data retrieval call binding the contract method 0xb7dc1284.
//
// Solidity: function agentCount() view returns(uint256)
func (_AgentFactory *AgentFactoryCallerSession) AgentCount() (*big.Int, error) {
	return _AgentFactory.Contract.AgentCount(&_AgentFactory.CallOpts)
}

// Agents is a free data retrieval call binding the contract method 0xfd66091e.
//
// Solidity: function agents(address ) view returns(uint256)
func (_AgentFactory *AgentFactoryCaller) Agents(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AgentFactory.contract.Call(opts, &out, "agents", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Agents is a free data retrieval call binding the contract method 0xfd66091e.
//
// Solidity: function agents(address ) view returns(uint256)
func (_AgentFactory *AgentFactorySession) Agents(arg0 common.Address) (*big.Int, error) {
	return _AgentFactory.Contract.Agents(&_AgentFactory.CallOpts, arg0)
}

// Agents is a free data retrieval call binding the contract method 0xfd66091e.
//
// Solidity: function agents(address ) view returns(uint256)
func (_AgentFactory *AgentFactoryCallerSession) Agents(arg0 common.Address) (*big.Int, error) {
	return _AgentFactory.Contract.Agents(&_AgentFactory.CallOpts, arg0)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address agent) view returns(bool)
func (_AgentFactory *AgentFactoryCaller) IsAgent(opts *bind.CallOpts, agent common.Address) (bool, error) {
	var out []interface{}
	err := _AgentFactory.contract.Call(opts, &out, "isAgent", agent)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address agent) view returns(bool)
func (_AgentFactory *AgentFactorySession) IsAgent(agent common.Address) (bool, error) {
	return _AgentFactory.Contract.IsAgent(&_AgentFactory.CallOpts, agent)
}

// IsAgent is a free data retrieval call binding the contract method 0x1ffbb064.
//
// Solidity: function isAgent(address agent) view returns(bool)
func (_AgentFactory *AgentFactoryCallerSession) IsAgent(agent common.Address) (bool, error) {
	return _AgentFactory.Contract.IsAgent(&_AgentFactory.CallOpts, agent)
}

// Create is a paid mutator transaction binding the contract method 0x9c041ebd.
//
// Solidity: function create(address owner, address operator, address adoRequestKey) returns(address agent)
func (_AgentFactory *AgentFactoryTransactor) Create(opts *bind.TransactOpts, owner common.Address, operator common.Address, adoRequestKey common.Address) (*types.Transaction, error) {
	return _AgentFactory.contract.Transact(opts, "create", owner, operator, adoRequestKey)
}

// Create is a paid mutator transaction binding the contract method 0x9c041ebd.
//
// Solidity: function create(address owner, address operator, address adoRequestKey) returns(address agent)
func (_AgentFactory *AgentFactorySession) Create(owner common.Address, operator common.Address, adoRequestKey common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.Create(&_AgentFactory.TransactOpts, owner, operator, adoRequestKey)
}

// Create is a paid mutator transaction binding the contract method 0x9c041ebd.
//
// Solidity: function create(address owner, address operator, address adoRequestKey) returns(address agent)
func (_AgentFactory *AgentFactoryTransactorSession) Create(owner common.Address, operator common.Address, adoRequestKey common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.Create(&_AgentFactory.TransactOpts, owner, operator, adoRequestKey)
}

// UpgradeAgent is a paid mutator transaction binding the contract method 0xe4513014.
//
// Solidity: function upgradeAgent(address agent) returns(address newAgent)
func (_AgentFactory *AgentFactoryTransactor) UpgradeAgent(opts *bind.TransactOpts, agent common.Address) (*types.Transaction, error) {
	return _AgentFactory.contract.Transact(opts, "upgradeAgent", agent)
}

// UpgradeAgent is a paid mutator transaction binding the contract method 0xe4513014.
//
// Solidity: function upgradeAgent(address agent) returns(address newAgent)
func (_AgentFactory *AgentFactorySession) UpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.UpgradeAgent(&_AgentFactory.TransactOpts, agent)
}

// UpgradeAgent is a paid mutator transaction binding the contract method 0xe4513014.
//
// Solidity: function upgradeAgent(address agent) returns(address newAgent)
func (_AgentFactory *AgentFactoryTransactorSession) UpgradeAgent(agent common.Address) (*types.Transaction, error) {
	return _AgentFactory.Contract.UpgradeAgent(&_AgentFactory.TransactOpts, agent)
}

// AgentFactoryCreateAgentIterator is returned from FilterCreateAgent and is used to iterate over the raw logs and unpacked data for CreateAgent events raised by the AgentFactory contract.
type AgentFactoryCreateAgentIterator struct {
	Event *AgentFactoryCreateAgent // Event containing the contract specifics and raw log

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
func (it *AgentFactoryCreateAgentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentFactoryCreateAgent)
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
		it.Event = new(AgentFactoryCreateAgent)
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
func (it *AgentFactoryCreateAgentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentFactoryCreateAgentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentFactoryCreateAgent represents a CreateAgent event raised by the AgentFactory contract.
type AgentFactoryCreateAgent struct {
	AgentID *big.Int
	Agent   common.Address
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateAgent is a free log retrieval operation binding the contract event 0x804d5b2bc79ae16fc8fefb3681f47fbc5dd07c51404c26ef178a24dc8c037c4b.
//
// Solidity: event CreateAgent(uint256 indexed agentID, address indexed agent, address indexed creator)
func (_AgentFactory *AgentFactoryFilterer) FilterCreateAgent(opts *bind.FilterOpts, agentID []*big.Int, agent []common.Address, creator []common.Address) (*AgentFactoryCreateAgentIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AgentFactory.contract.FilterLogs(opts, "CreateAgent", agentIDRule, agentRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &AgentFactoryCreateAgentIterator{contract: _AgentFactory.contract, event: "CreateAgent", logs: logs, sub: sub}, nil
}

// WatchCreateAgent is a free log subscription operation binding the contract event 0x804d5b2bc79ae16fc8fefb3681f47fbc5dd07c51404c26ef178a24dc8c037c4b.
//
// Solidity: event CreateAgent(uint256 indexed agentID, address indexed agent, address indexed creator)
func (_AgentFactory *AgentFactoryFilterer) WatchCreateAgent(opts *bind.WatchOpts, sink chan<- *AgentFactoryCreateAgent, agentID []*big.Int, agent []common.Address, creator []common.Address) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AgentFactory.contract.WatchLogs(opts, "CreateAgent", agentIDRule, agentRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentFactoryCreateAgent)
				if err := _AgentFactory.contract.UnpackLog(event, "CreateAgent", log); err != nil {
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

// ParseCreateAgent is a log parse operation binding the contract event 0x804d5b2bc79ae16fc8fefb3681f47fbc5dd07c51404c26ef178a24dc8c037c4b.
//
// Solidity: event CreateAgent(uint256 indexed agentID, address indexed agent, address indexed creator)
func (_AgentFactory *AgentFactoryFilterer) ParseCreateAgent(log types.Log) (*AgentFactoryCreateAgent, error) {
	event := new(AgentFactoryCreateAgent)
	if err := _AgentFactory.contract.UnpackLog(event, "CreateAgent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
