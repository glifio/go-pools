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

// Account is an auto generated low-level Go binding around an user-defined struct.
type Account struct {
	StartEpoch *big.Int
	Principal  *big.Int
	EpochsPaid *big.Int
	Defaulted  bool
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouteDNE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRoute\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"id\",\"type\":\"bytes4\"}],\"name\":\"PushRoute\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"}],\"name\":\"createAccountKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"}],\"name\":\"getAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"defaulted\",\"type\":\"bool\"}],\"internalType\":\"structAccount\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"id\",\"type\":\"bytes4\"}],\"name\":\"getRoute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"getRoute\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newRoute\",\"type\":\"address\"}],\"name\":\"pushRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"id\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"newRoute\",\"type\":\"address\"}],\"name\":\"pushRoute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"ids\",\"type\":\"string[]\"},{\"internalType\":\"address[]\",\"name\":\"newRoutes\",\"type\":\"address[]\"}],\"name\":\"pushRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"ids\",\"type\":\"bytes4[]\"},{\"internalType\":\"address[]\",\"name\":\"newRoutes\",\"type\":\"address[]\"}],\"name\":\"pushRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"name\":\"route\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolID\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"defaulted\",\"type\":\"bool\"}],\"internalType\":\"structAccount\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"setAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// CreateAccountKey is a free data retrieval call binding the contract method 0xb761b6f3.
//
// Solidity: function createAccountKey(uint256 agentID, uint256 poolID) pure returns(bytes32)
func (_Router *RouterCaller) CreateAccountKey(opts *bind.CallOpts, agentID *big.Int, poolID *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "createAccountKey", agentID, poolID)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CreateAccountKey is a free data retrieval call binding the contract method 0xb761b6f3.
//
// Solidity: function createAccountKey(uint256 agentID, uint256 poolID) pure returns(bytes32)
func (_Router *RouterSession) CreateAccountKey(agentID *big.Int, poolID *big.Int) ([32]byte, error) {
	return _Router.Contract.CreateAccountKey(&_Router.CallOpts, agentID, poolID)
}

// CreateAccountKey is a free data retrieval call binding the contract method 0xb761b6f3.
//
// Solidity: function createAccountKey(uint256 agentID, uint256 poolID) pure returns(bytes32)
func (_Router *RouterCallerSession) CreateAccountKey(agentID *big.Int, poolID *big.Int) ([32]byte, error) {
	return _Router.Contract.CreateAccountKey(&_Router.CallOpts, agentID, poolID)
}

// GetAccount is a free data retrieval call binding the contract method 0x6361f6de.
//
// Solidity: function getAccount(uint256 agentID, uint256 poolID) view returns((uint256,uint256,uint256,bool))
func (_Router *RouterCaller) GetAccount(opts *bind.CallOpts, agentID *big.Int, poolID *big.Int) (Account, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getAccount", agentID, poolID)

	if err != nil {
		return *new(Account), err
	}

	out0 := *abi.ConvertType(out[0], new(Account)).(*Account)

	return out0, err

}

// GetAccount is a free data retrieval call binding the contract method 0x6361f6de.
//
// Solidity: function getAccount(uint256 agentID, uint256 poolID) view returns((uint256,uint256,uint256,bool))
func (_Router *RouterSession) GetAccount(agentID *big.Int, poolID *big.Int) (Account, error) {
	return _Router.Contract.GetAccount(&_Router.CallOpts, agentID, poolID)
}

// GetAccount is a free data retrieval call binding the contract method 0x6361f6de.
//
// Solidity: function getAccount(uint256 agentID, uint256 poolID) view returns((uint256,uint256,uint256,bool))
func (_Router *RouterCallerSession) GetAccount(agentID *big.Int, poolID *big.Int) (Account, error) {
	return _Router.Contract.GetAccount(&_Router.CallOpts, agentID, poolID)
}

// GetRoute is a free data retrieval call binding the contract method 0x1a6e649e.
//
// Solidity: function getRoute(bytes4 id) view returns(address)
func (_Router *RouterCaller) GetRoute(opts *bind.CallOpts, id [4]byte) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getRoute", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoute is a free data retrieval call binding the contract method 0x1a6e649e.
//
// Solidity: function getRoute(bytes4 id) view returns(address)
func (_Router *RouterSession) GetRoute(id [4]byte) (common.Address, error) {
	return _Router.Contract.GetRoute(&_Router.CallOpts, id)
}

// GetRoute is a free data retrieval call binding the contract method 0x1a6e649e.
//
// Solidity: function getRoute(bytes4 id) view returns(address)
func (_Router *RouterCallerSession) GetRoute(id [4]byte) (common.Address, error) {
	return _Router.Contract.GetRoute(&_Router.CallOpts, id)
}

// GetRoute0 is a free data retrieval call binding the contract method 0x1b5f03a6.
//
// Solidity: function getRoute(string id) view returns(address)
func (_Router *RouterCaller) GetRoute0(opts *bind.CallOpts, id string) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getRoute0", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoute0 is a free data retrieval call binding the contract method 0x1b5f03a6.
//
// Solidity: function getRoute(string id) view returns(address)
func (_Router *RouterSession) GetRoute0(id string) (common.Address, error) {
	return _Router.Contract.GetRoute0(&_Router.CallOpts, id)
}

// GetRoute0 is a free data retrieval call binding the contract method 0x1b5f03a6.
//
// Solidity: function getRoute(string id) view returns(address)
func (_Router *RouterCallerSession) GetRoute0(id string) (common.Address, error) {
	return _Router.Contract.GetRoute0(&_Router.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Router *RouterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Router *RouterSession) PendingOwner() (common.Address, error) {
	return _Router.Contract.PendingOwner(&_Router.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Router *RouterCallerSession) PendingOwner() (common.Address, error) {
	return _Router.Contract.PendingOwner(&_Router.CallOpts)
}

// Route is a free data retrieval call binding the contract method 0xc5ae7a11.
//
// Solidity: function route(bytes4 ) view returns(address)
func (_Router *RouterCaller) Route(opts *bind.CallOpts, arg0 [4]byte) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "route", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Route is a free data retrieval call binding the contract method 0xc5ae7a11.
//
// Solidity: function route(bytes4 ) view returns(address)
func (_Router *RouterSession) Route(arg0 [4]byte) (common.Address, error) {
	return _Router.Contract.Route(&_Router.CallOpts, arg0)
}

// Route is a free data retrieval call binding the contract method 0xc5ae7a11.
//
// Solidity: function route(bytes4 ) view returns(address)
func (_Router *RouterCallerSession) Route(arg0 [4]byte) (common.Address, error) {
	return _Router.Contract.Route(&_Router.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Router *RouterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Router *RouterSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Router *RouterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Router.Contract.AcceptOwnership(&_Router.TransactOpts)
}

// PushRoute is a paid mutator transaction binding the contract method 0x19ac0743.
//
// Solidity: function pushRoute(string id, address newRoute) returns()
func (_Router *RouterTransactor) PushRoute(opts *bind.TransactOpts, id string, newRoute common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "pushRoute", id, newRoute)
}

// PushRoute is a paid mutator transaction binding the contract method 0x19ac0743.
//
// Solidity: function pushRoute(string id, address newRoute) returns()
func (_Router *RouterSession) PushRoute(id string, newRoute common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoute(&_Router.TransactOpts, id, newRoute)
}

// PushRoute is a paid mutator transaction binding the contract method 0x19ac0743.
//
// Solidity: function pushRoute(string id, address newRoute) returns()
func (_Router *RouterTransactorSession) PushRoute(id string, newRoute common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoute(&_Router.TransactOpts, id, newRoute)
}

// PushRoute0 is a paid mutator transaction binding the contract method 0x334908b2.
//
// Solidity: function pushRoute(bytes4 id, address newRoute) returns()
func (_Router *RouterTransactor) PushRoute0(opts *bind.TransactOpts, id [4]byte, newRoute common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "pushRoute0", id, newRoute)
}

// PushRoute0 is a paid mutator transaction binding the contract method 0x334908b2.
//
// Solidity: function pushRoute(bytes4 id, address newRoute) returns()
func (_Router *RouterSession) PushRoute0(id [4]byte, newRoute common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoute0(&_Router.TransactOpts, id, newRoute)
}

// PushRoute0 is a paid mutator transaction binding the contract method 0x334908b2.
//
// Solidity: function pushRoute(bytes4 id, address newRoute) returns()
func (_Router *RouterTransactorSession) PushRoute0(id [4]byte, newRoute common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoute0(&_Router.TransactOpts, id, newRoute)
}

// PushRoutes is a paid mutator transaction binding the contract method 0x7df3be51.
//
// Solidity: function pushRoutes(string[] ids, address[] newRoutes) returns()
func (_Router *RouterTransactor) PushRoutes(opts *bind.TransactOpts, ids []string, newRoutes []common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "pushRoutes", ids, newRoutes)
}

// PushRoutes is a paid mutator transaction binding the contract method 0x7df3be51.
//
// Solidity: function pushRoutes(string[] ids, address[] newRoutes) returns()
func (_Router *RouterSession) PushRoutes(ids []string, newRoutes []common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoutes(&_Router.TransactOpts, ids, newRoutes)
}

// PushRoutes is a paid mutator transaction binding the contract method 0x7df3be51.
//
// Solidity: function pushRoutes(string[] ids, address[] newRoutes) returns()
func (_Router *RouterTransactorSession) PushRoutes(ids []string, newRoutes []common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoutes(&_Router.TransactOpts, ids, newRoutes)
}

// PushRoutes0 is a paid mutator transaction binding the contract method 0xef254abd.
//
// Solidity: function pushRoutes(bytes4[] ids, address[] newRoutes) returns()
func (_Router *RouterTransactor) PushRoutes0(opts *bind.TransactOpts, ids [][4]byte, newRoutes []common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "pushRoutes0", ids, newRoutes)
}

// PushRoutes0 is a paid mutator transaction binding the contract method 0xef254abd.
//
// Solidity: function pushRoutes(bytes4[] ids, address[] newRoutes) returns()
func (_Router *RouterSession) PushRoutes0(ids [][4]byte, newRoutes []common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoutes0(&_Router.TransactOpts, ids, newRoutes)
}

// PushRoutes0 is a paid mutator transaction binding the contract method 0xef254abd.
//
// Solidity: function pushRoutes(bytes4[] ids, address[] newRoutes) returns()
func (_Router *RouterTransactorSession) PushRoutes0(ids [][4]byte, newRoutes []common.Address) (*types.Transaction, error) {
	return _Router.Contract.PushRoutes0(&_Router.TransactOpts, ids, newRoutes)
}

// SetAccount is a paid mutator transaction binding the contract method 0xc7afd04b.
//
// Solidity: function setAccount(uint256 agentID, uint256 poolID, (uint256,uint256,uint256,bool) account) returns()
func (_Router *RouterTransactor) SetAccount(opts *bind.TransactOpts, agentID *big.Int, poolID *big.Int, account Account) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setAccount", agentID, poolID, account)
}

// SetAccount is a paid mutator transaction binding the contract method 0xc7afd04b.
//
// Solidity: function setAccount(uint256 agentID, uint256 poolID, (uint256,uint256,uint256,bool) account) returns()
func (_Router *RouterSession) SetAccount(agentID *big.Int, poolID *big.Int, account Account) (*types.Transaction, error) {
	return _Router.Contract.SetAccount(&_Router.TransactOpts, agentID, poolID, account)
}

// SetAccount is a paid mutator transaction binding the contract method 0xc7afd04b.
//
// Solidity: function setAccount(uint256 agentID, uint256 poolID, (uint256,uint256,uint256,bool) account) returns()
func (_Router *RouterTransactorSession) SetAccount(agentID *big.Int, poolID *big.Int, account Account) (*types.Transaction, error) {
	return _Router.Contract.SetAccount(&_Router.TransactOpts, agentID, poolID, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// RouterPushRouteIterator is returned from FilterPushRoute and is used to iterate over the raw logs and unpacked data for PushRoute events raised by the Router contract.
type RouterPushRouteIterator struct {
	Event *RouterPushRoute // Event containing the contract specifics and raw log

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
func (it *RouterPushRouteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterPushRoute)
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
		it.Event = new(RouterPushRoute)
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
func (it *RouterPushRouteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterPushRouteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterPushRoute represents a PushRoute event raised by the Router contract.
type RouterPushRoute struct {
	NewRoute common.Address
	Id       [4]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPushRoute is a free log retrieval operation binding the contract event 0xd0bd3873536bdca53cdb3985d74c5be9fdb9172f717e7d305b3dab7828e1df87.
//
// Solidity: event PushRoute(address newRoute, bytes4 id)
func (_Router *RouterFilterer) FilterPushRoute(opts *bind.FilterOpts) (*RouterPushRouteIterator, error) {

	logs, sub, err := _Router.contract.FilterLogs(opts, "PushRoute")
	if err != nil {
		return nil, err
	}
	return &RouterPushRouteIterator{contract: _Router.contract, event: "PushRoute", logs: logs, sub: sub}, nil
}

// WatchPushRoute is a free log subscription operation binding the contract event 0xd0bd3873536bdca53cdb3985d74c5be9fdb9172f717e7d305b3dab7828e1df87.
//
// Solidity: event PushRoute(address newRoute, bytes4 id)
func (_Router *RouterFilterer) WatchPushRoute(opts *bind.WatchOpts, sink chan<- *RouterPushRoute) (event.Subscription, error) {

	logs, sub, err := _Router.contract.WatchLogs(opts, "PushRoute")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterPushRoute)
				if err := _Router.contract.UnpackLog(event, "PushRoute", log); err != nil {
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

// ParsePushRoute is a log parse operation binding the contract event 0xd0bd3873536bdca53cdb3985d74c5be9fdb9172f717e7d305b3dab7828e1df87.
//
// Solidity: event PushRoute(address newRoute, bytes4 id)
func (_Router *RouterFilterer) ParsePushRoute(log types.Log) (*RouterPushRoute, error) {
	event := new(RouterPushRoute)
	if err := _Router.contract.UnpackLog(event, "PushRoute", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
