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

// PoolRegistryMetaData contains all meta data concerning the PoolRegistry contract.
var PoolRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treasuryFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MAX_TREASURY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"}],\"name\":\"addPoolToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"attachPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"poolIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"}],\"name\":\"removePoolFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setTreasuryFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"upgradePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolRegistryMetaData.ABI instead.
var PoolRegistryABI = PoolRegistryMetaData.ABI

// PoolRegistry is an auto generated Go binding around an Ethereum contract.
type PoolRegistry struct {
	PoolRegistryCaller     // Read-only binding to the contract
	PoolRegistryTransactor // Write-only binding to the contract
	PoolRegistryFilterer   // Log filterer for contract events
}

// PoolRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolRegistrySession struct {
	Contract     *PoolRegistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolRegistryCallerSession struct {
	Contract *PoolRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PoolRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolRegistryTransactorSession struct {
	Contract     *PoolRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PoolRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRegistryRaw struct {
	Contract *PoolRegistry // Generic contract binding to access the raw methods on
}

// PoolRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolRegistryCallerRaw struct {
	Contract *PoolRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PoolRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolRegistryTransactorRaw struct {
	Contract *PoolRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolRegistry creates a new instance of PoolRegistry, bound to a specific deployed contract.
func NewPoolRegistry(address common.Address, backend bind.ContractBackend) (*PoolRegistry, error) {
	contract, err := bindPoolRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolRegistry{PoolRegistryCaller: PoolRegistryCaller{contract: contract}, PoolRegistryTransactor: PoolRegistryTransactor{contract: contract}, PoolRegistryFilterer: PoolRegistryFilterer{contract: contract}}, nil
}

// NewPoolRegistryCaller creates a new read-only instance of PoolRegistry, bound to a specific deployed contract.
func NewPoolRegistryCaller(address common.Address, caller bind.ContractCaller) (*PoolRegistryCaller, error) {
	contract, err := bindPoolRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolRegistryCaller{contract: contract}, nil
}

// NewPoolRegistryTransactor creates a new write-only instance of PoolRegistry, bound to a specific deployed contract.
func NewPoolRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolRegistryTransactor, error) {
	contract, err := bindPoolRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolRegistryTransactor{contract: contract}, nil
}

// NewPoolRegistryFilterer creates a new log filterer instance of PoolRegistry, bound to a specific deployed contract.
func NewPoolRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolRegistryFilterer, error) {
	contract, err := bindPoolRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolRegistryFilterer{contract: contract}, nil
}

// bindPoolRegistry binds a generic wrapper to an already deployed contract.
func bindPoolRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolRegistry *PoolRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolRegistry.Contract.PoolRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolRegistry *PoolRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolRegistry.Contract.PoolRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolRegistry *PoolRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolRegistry.Contract.PoolRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolRegistry *PoolRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolRegistry *PoolRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolRegistry *PoolRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolRegistry.Contract.contract.Transact(opts, method, params...)
}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_PoolRegistry *PoolRegistryCaller) MAXTREASURYFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolRegistry.contract.Call(opts, &out, "MAX_TREASURY_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_PoolRegistry *PoolRegistrySession) MAXTREASURYFEE() (*big.Int, error) {
	return _PoolRegistry.Contract.MAXTREASURYFEE(&_PoolRegistry.CallOpts)
}

// MAXTREASURYFEE is a free data retrieval call binding the contract method 0xf2b3c809.
//
// Solidity: function MAX_TREASURY_FEE() view returns(uint256)
func (_PoolRegistry *PoolRegistryCallerSession) MAXTREASURYFEE() (*big.Int, error) {
	return _PoolRegistry.Contract.MAXTREASURYFEE(&_PoolRegistry.CallOpts)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_PoolRegistry *PoolRegistryCaller) AllPools(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PoolRegistry.contract.Call(opts, &out, "allPools", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_PoolRegistry *PoolRegistrySession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _PoolRegistry.Contract.AllPools(&_PoolRegistry.CallOpts, arg0)
}

// AllPools is a free data retrieval call binding the contract method 0x41d1de97.
//
// Solidity: function allPools(uint256 ) view returns(address)
func (_PoolRegistry *PoolRegistryCallerSession) AllPools(arg0 *big.Int) (common.Address, error) {
	return _PoolRegistry.Contract.AllPools(&_PoolRegistry.CallOpts, arg0)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_PoolRegistry *PoolRegistryCaller) AllPoolsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolRegistry.contract.Call(opts, &out, "allPoolsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_PoolRegistry *PoolRegistrySession) AllPoolsLength() (*big.Int, error) {
	return _PoolRegistry.Contract.AllPoolsLength(&_PoolRegistry.CallOpts)
}

// AllPoolsLength is a free data retrieval call binding the contract method 0xefde4e64.
//
// Solidity: function allPoolsLength() view returns(uint256)
func (_PoolRegistry *PoolRegistryCallerSession) AllPoolsLength() (*big.Int, error) {
	return _PoolRegistry.Contract.AllPoolsLength(&_PoolRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolRegistry *PoolRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolRegistry *PoolRegistrySession) Owner() (common.Address, error) {
	return _PoolRegistry.Contract.Owner(&_PoolRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolRegistry *PoolRegistryCallerSession) Owner() (common.Address, error) {
	return _PoolRegistry.Contract.Owner(&_PoolRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PoolRegistry *PoolRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PoolRegistry *PoolRegistrySession) PendingOwner() (common.Address, error) {
	return _PoolRegistry.Contract.PendingOwner(&_PoolRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PoolRegistry *PoolRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _PoolRegistry.Contract.PendingOwner(&_PoolRegistry.CallOpts)
}

// PoolIDs is a free data retrieval call binding the contract method 0x9a014e24.
//
// Solidity: function poolIDs(uint256 agentID) view returns(uint256[])
func (_PoolRegistry *PoolRegistryCaller) PoolIDs(opts *bind.CallOpts, agentID *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _PoolRegistry.contract.Call(opts, &out, "poolIDs", agentID)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// PoolIDs is a free data retrieval call binding the contract method 0x9a014e24.
//
// Solidity: function poolIDs(uint256 agentID) view returns(uint256[])
func (_PoolRegistry *PoolRegistrySession) PoolIDs(agentID *big.Int) ([]*big.Int, error) {
	return _PoolRegistry.Contract.PoolIDs(&_PoolRegistry.CallOpts, agentID)
}

// PoolIDs is a free data retrieval call binding the contract method 0x9a014e24.
//
// Solidity: function poolIDs(uint256 agentID) view returns(uint256[])
func (_PoolRegistry *PoolRegistryCallerSession) PoolIDs(agentID *big.Int) ([]*big.Int, error) {
	return _PoolRegistry.Contract.PoolIDs(&_PoolRegistry.CallOpts, agentID)
}

// TreasuryFeeRate is a free data retrieval call binding the contract method 0x66601032.
//
// Solidity: function treasuryFeeRate() view returns(uint256)
func (_PoolRegistry *PoolRegistryCaller) TreasuryFeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolRegistry.contract.Call(opts, &out, "treasuryFeeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TreasuryFeeRate is a free data retrieval call binding the contract method 0x66601032.
//
// Solidity: function treasuryFeeRate() view returns(uint256)
func (_PoolRegistry *PoolRegistrySession) TreasuryFeeRate() (*big.Int, error) {
	return _PoolRegistry.Contract.TreasuryFeeRate(&_PoolRegistry.CallOpts)
}

// TreasuryFeeRate is a free data retrieval call binding the contract method 0x66601032.
//
// Solidity: function treasuryFeeRate() view returns(uint256)
func (_PoolRegistry *PoolRegistryCallerSession) TreasuryFeeRate() (*big.Int, error) {
	return _PoolRegistry.Contract.TreasuryFeeRate(&_PoolRegistry.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PoolRegistry *PoolRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PoolRegistry *PoolRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _PoolRegistry.Contract.AcceptOwnership(&_PoolRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PoolRegistry *PoolRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PoolRegistry.Contract.AcceptOwnership(&_PoolRegistry.TransactOpts)
}

// AddPoolToList is a paid mutator transaction binding the contract method 0x6834f901.
//
// Solidity: function addPoolToList(uint256 agentID, uint256 pool) returns()
func (_PoolRegistry *PoolRegistryTransactor) AddPoolToList(opts *bind.TransactOpts, agentID *big.Int, pool *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.contract.Transact(opts, "addPoolToList", agentID, pool)
}

// AddPoolToList is a paid mutator transaction binding the contract method 0x6834f901.
//
// Solidity: function addPoolToList(uint256 agentID, uint256 pool) returns()
func (_PoolRegistry *PoolRegistrySession) AddPoolToList(agentID *big.Int, pool *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.Contract.AddPoolToList(&_PoolRegistry.TransactOpts, agentID, pool)
}

// AddPoolToList is a paid mutator transaction binding the contract method 0x6834f901.
//
// Solidity: function addPoolToList(uint256 agentID, uint256 pool) returns()
func (_PoolRegistry *PoolRegistryTransactorSession) AddPoolToList(agentID *big.Int, pool *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.Contract.AddPoolToList(&_PoolRegistry.TransactOpts, agentID, pool)
}

// AttachPool is a paid mutator transaction binding the contract method 0x12b1ee6c.
//
// Solidity: function attachPool(address pool) returns()
func (_PoolRegistry *PoolRegistryTransactor) AttachPool(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _PoolRegistry.contract.Transact(opts, "attachPool", pool)
}

// AttachPool is a paid mutator transaction binding the contract method 0x12b1ee6c.
//
// Solidity: function attachPool(address pool) returns()
func (_PoolRegistry *PoolRegistrySession) AttachPool(pool common.Address) (*types.Transaction, error) {
	return _PoolRegistry.Contract.AttachPool(&_PoolRegistry.TransactOpts, pool)
}

// AttachPool is a paid mutator transaction binding the contract method 0x12b1ee6c.
//
// Solidity: function attachPool(address pool) returns()
func (_PoolRegistry *PoolRegistryTransactorSession) AttachPool(pool common.Address) (*types.Transaction, error) {
	return _PoolRegistry.Contract.AttachPool(&_PoolRegistry.TransactOpts, pool)
}

// RemovePoolFromList is a paid mutator transaction binding the contract method 0x8152a6c6.
//
// Solidity: function removePoolFromList(uint256 agentID, uint256 pool) returns()
func (_PoolRegistry *PoolRegistryTransactor) RemovePoolFromList(opts *bind.TransactOpts, agentID *big.Int, pool *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.contract.Transact(opts, "removePoolFromList", agentID, pool)
}

// RemovePoolFromList is a paid mutator transaction binding the contract method 0x8152a6c6.
//
// Solidity: function removePoolFromList(uint256 agentID, uint256 pool) returns()
func (_PoolRegistry *PoolRegistrySession) RemovePoolFromList(agentID *big.Int, pool *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.Contract.RemovePoolFromList(&_PoolRegistry.TransactOpts, agentID, pool)
}

// RemovePoolFromList is a paid mutator transaction binding the contract method 0x8152a6c6.
//
// Solidity: function removePoolFromList(uint256 agentID, uint256 pool) returns()
func (_PoolRegistry *PoolRegistryTransactorSession) RemovePoolFromList(agentID *big.Int, pool *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.Contract.RemovePoolFromList(&_PoolRegistry.TransactOpts, agentID, pool)
}

// SetTreasuryFeeRate is a paid mutator transaction binding the contract method 0xce43303c.
//
// Solidity: function setTreasuryFeeRate(uint256 newFeeRate) returns()
func (_PoolRegistry *PoolRegistryTransactor) SetTreasuryFeeRate(opts *bind.TransactOpts, newFeeRate *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.contract.Transact(opts, "setTreasuryFeeRate", newFeeRate)
}

// SetTreasuryFeeRate is a paid mutator transaction binding the contract method 0xce43303c.
//
// Solidity: function setTreasuryFeeRate(uint256 newFeeRate) returns()
func (_PoolRegistry *PoolRegistrySession) SetTreasuryFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.Contract.SetTreasuryFeeRate(&_PoolRegistry.TransactOpts, newFeeRate)
}

// SetTreasuryFeeRate is a paid mutator transaction binding the contract method 0xce43303c.
//
// Solidity: function setTreasuryFeeRate(uint256 newFeeRate) returns()
func (_PoolRegistry *PoolRegistryTransactorSession) SetTreasuryFeeRate(newFeeRate *big.Int) (*types.Transaction, error) {
	return _PoolRegistry.Contract.SetTreasuryFeeRate(&_PoolRegistry.TransactOpts, newFeeRate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoolRegistry *PoolRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PoolRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoolRegistry *PoolRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoolRegistry.Contract.TransferOwnership(&_PoolRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PoolRegistry *PoolRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PoolRegistry.Contract.TransferOwnership(&_PoolRegistry.TransactOpts, newOwner)
}

// UpgradePool is a paid mutator transaction binding the contract method 0xa2ff484f.
//
// Solidity: function upgradePool(address newPool) returns()
func (_PoolRegistry *PoolRegistryTransactor) UpgradePool(opts *bind.TransactOpts, newPool common.Address) (*types.Transaction, error) {
	return _PoolRegistry.contract.Transact(opts, "upgradePool", newPool)
}

// UpgradePool is a paid mutator transaction binding the contract method 0xa2ff484f.
//
// Solidity: function upgradePool(address newPool) returns()
func (_PoolRegistry *PoolRegistrySession) UpgradePool(newPool common.Address) (*types.Transaction, error) {
	return _PoolRegistry.Contract.UpgradePool(&_PoolRegistry.TransactOpts, newPool)
}

// UpgradePool is a paid mutator transaction binding the contract method 0xa2ff484f.
//
// Solidity: function upgradePool(address newPool) returns()
func (_PoolRegistry *PoolRegistryTransactorSession) UpgradePool(newPool common.Address) (*types.Transaction, error) {
	return _PoolRegistry.Contract.UpgradePool(&_PoolRegistry.TransactOpts, newPool)
}
