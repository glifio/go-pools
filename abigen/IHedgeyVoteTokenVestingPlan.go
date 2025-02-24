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

// IHedgeyVoteTokenVestingPlanPlan is an auto generated low-level Go binding around an user-defined struct.
type IHedgeyVoteTokenVestingPlanPlan struct {
	Token            common.Address
	Amount           *big.Int
	Start            *big.Int
	Cliff            *big.Int
	Rate             *big.Int
	Period           *big.Int
	VestingAdmin     common.Address
	AdminTransferOBO bool
}

// IHedgeyVoteTokenVestingPlanMetaData contains all meta data concerning the IHedgeyVoteTokenVestingPlan contract.
var IHedgeyVoteTokenVestingPlanMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newPlanId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vestingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"adminTransferOBO\",\"type\":\"bool\"}],\"name\":\"PlanCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainder\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"latestUnlock\",\"type\":\"uint256\"}],\"name\":\"PlanRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainder\",\"type\":\"uint256\"}],\"name\":\"PlanRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"PlanTransferredByVestingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferrable\",\"type\":\"bool\"}],\"name\":\"PlanVestingAdminTransferToggle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newVestingAdmin\",\"type\":\"address\"}],\"name\":\"VestingPlanAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"VotingVaultCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"_safeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newVestingAdmin\",\"type\":\"address\"}],\"name\":\"changeVestingPlanAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vestingAdmin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"adminTransferOBO\",\"type\":\"bool\"}],\"name\":\"createPlan\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"newPlanId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegateAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"delegatees\",\"type\":\"address[]\"}],\"name\":\"delegatePlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"revokeTime\",\"type\":\"uint256\"}],\"name\":\"futureRevokePlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"lockedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"redemptionTime\",\"type\":\"uint256\"}],\"name\":\"partialRedeemPlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionTime\",\"type\":\"uint256\"}],\"name\":\"planBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestUnlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"}],\"name\":\"plans\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"vestingAdmin\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"adminTransferOBO\",\"type\":\"bool\"}],\"internalType\":\"structIHedgeyVoteTokenVestingPlan.Plan\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemAllPlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"}],\"name\":\"redeemPlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"}],\"name\":\"revokePlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"}],\"name\":\"setupVoting\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"votingVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"transferrable\",\"type\":\"bool\"}],\"name\":\"toggleAdminTransferOBO\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IHedgeyVoteTokenVestingPlanABI is the input ABI used to generate the binding from.
// Deprecated: Use IHedgeyVoteTokenVestingPlanMetaData.ABI instead.
var IHedgeyVoteTokenVestingPlanABI = IHedgeyVoteTokenVestingPlanMetaData.ABI

// IHedgeyVoteTokenVestingPlan is an auto generated Go binding around an Ethereum contract.
type IHedgeyVoteTokenVestingPlan struct {
	IHedgeyVoteTokenVestingPlanCaller     // Read-only binding to the contract
	IHedgeyVoteTokenVestingPlanTransactor // Write-only binding to the contract
	IHedgeyVoteTokenVestingPlanFilterer   // Log filterer for contract events
}

// IHedgeyVoteTokenVestingPlanCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenVestingPlanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyVoteTokenVestingPlanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenVestingPlanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyVoteTokenVestingPlanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHedgeyVoteTokenVestingPlanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyVoteTokenVestingPlanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHedgeyVoteTokenVestingPlanSession struct {
	Contract     *IHedgeyVoteTokenVestingPlan // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                // Call options to use throughout this session
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IHedgeyVoteTokenVestingPlanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHedgeyVoteTokenVestingPlanCallerSession struct {
	Contract *IHedgeyVoteTokenVestingPlanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                      // Call options to use throughout this session
}

// IHedgeyVoteTokenVestingPlanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHedgeyVoteTokenVestingPlanTransactorSession struct {
	Contract     *IHedgeyVoteTokenVestingPlanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                      // Transaction auth options to use throughout this session
}

// IHedgeyVoteTokenVestingPlanRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHedgeyVoteTokenVestingPlanRaw struct {
	Contract *IHedgeyVoteTokenVestingPlan // Generic contract binding to access the raw methods on
}

// IHedgeyVoteTokenVestingPlanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenVestingPlanCallerRaw struct {
	Contract *IHedgeyVoteTokenVestingPlanCaller // Generic read-only contract binding to access the raw methods on
}

// IHedgeyVoteTokenVestingPlanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenVestingPlanTransactorRaw struct {
	Contract *IHedgeyVoteTokenVestingPlanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHedgeyVoteTokenVestingPlan creates a new instance of IHedgeyVoteTokenVestingPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenVestingPlan(address common.Address, backend bind.ContractBackend) (*IHedgeyVoteTokenVestingPlan, error) {
	contract, err := bindIHedgeyVoteTokenVestingPlan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlan{IHedgeyVoteTokenVestingPlanCaller: IHedgeyVoteTokenVestingPlanCaller{contract: contract}, IHedgeyVoteTokenVestingPlanTransactor: IHedgeyVoteTokenVestingPlanTransactor{contract: contract}, IHedgeyVoteTokenVestingPlanFilterer: IHedgeyVoteTokenVestingPlanFilterer{contract: contract}}, nil
}

// NewIHedgeyVoteTokenVestingPlanCaller creates a new read-only instance of IHedgeyVoteTokenVestingPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenVestingPlanCaller(address common.Address, caller bind.ContractCaller) (*IHedgeyVoteTokenVestingPlanCaller, error) {
	contract, err := bindIHedgeyVoteTokenVestingPlan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanCaller{contract: contract}, nil
}

// NewIHedgeyVoteTokenVestingPlanTransactor creates a new write-only instance of IHedgeyVoteTokenVestingPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenVestingPlanTransactor(address common.Address, transactor bind.ContractTransactor) (*IHedgeyVoteTokenVestingPlanTransactor, error) {
	contract, err := bindIHedgeyVoteTokenVestingPlan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanTransactor{contract: contract}, nil
}

// NewIHedgeyVoteTokenVestingPlanFilterer creates a new log filterer instance of IHedgeyVoteTokenVestingPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenVestingPlanFilterer(address common.Address, filterer bind.ContractFilterer) (*IHedgeyVoteTokenVestingPlanFilterer, error) {
	contract, err := bindIHedgeyVoteTokenVestingPlan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanFilterer{contract: contract}, nil
}

// bindIHedgeyVoteTokenVestingPlan binds a generic wrapper to an already deployed contract.
func bindIHedgeyVoteTokenVestingPlan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHedgeyVoteTokenVestingPlanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHedgeyVoteTokenVestingPlan.Contract.IHedgeyVoteTokenVestingPlanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.IHedgeyVoteTokenVestingPlanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.IHedgeyVoteTokenVestingPlanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHedgeyVoteTokenVestingPlan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenVestingPlan.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.BalanceOf(&_IHedgeyVoteTokenVestingPlan.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.BalanceOf(&_IHedgeyVoteTokenVestingPlan.CallOpts, owner)
}

// LockedBalances is a free data retrieval call binding the contract method 0x52e5b7e4.
//
// Solidity: function lockedBalances(address holder, address token) view returns(uint256 lockedBalance)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCaller) LockedBalances(opts *bind.CallOpts, holder common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenVestingPlan.contract.Call(opts, &out, "lockedBalances", holder, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBalances is a free data retrieval call binding the contract method 0x52e5b7e4.
//
// Solidity: function lockedBalances(address holder, address token) view returns(uint256 lockedBalance)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) LockedBalances(holder common.Address, token common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.LockedBalances(&_IHedgeyVoteTokenVestingPlan.CallOpts, holder, token)
}

// LockedBalances is a free data retrieval call binding the contract method 0x52e5b7e4.
//
// Solidity: function lockedBalances(address holder, address token) view returns(uint256 lockedBalance)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCallerSession) LockedBalances(holder common.Address, token common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.LockedBalances(&_IHedgeyVoteTokenVestingPlan.CallOpts, holder, token)
}

// PlanBalanceOf is a free data retrieval call binding the contract method 0xc7d74fa7.
//
// Solidity: function planBalanceOf(uint256 planId, uint256 timeStamp, uint256 redemptionTime) view returns(uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCaller) PlanBalanceOf(opts *bind.CallOpts, planId *big.Int, timeStamp *big.Int, redemptionTime *big.Int) (struct {
	Balance      *big.Int
	Remainder    *big.Int
	LatestUnlock *big.Int
}, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenVestingPlan.contract.Call(opts, &out, "planBalanceOf", planId, timeStamp, redemptionTime)

	outstruct := new(struct {
		Balance      *big.Int
		Remainder    *big.Int
		LatestUnlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Remainder = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LatestUnlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PlanBalanceOf is a free data retrieval call binding the contract method 0xc7d74fa7.
//
// Solidity: function planBalanceOf(uint256 planId, uint256 timeStamp, uint256 redemptionTime) view returns(uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) PlanBalanceOf(planId *big.Int, timeStamp *big.Int, redemptionTime *big.Int) (struct {
	Balance      *big.Int
	Remainder    *big.Int
	LatestUnlock *big.Int
}, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.PlanBalanceOf(&_IHedgeyVoteTokenVestingPlan.CallOpts, planId, timeStamp, redemptionTime)
}

// PlanBalanceOf is a free data retrieval call binding the contract method 0xc7d74fa7.
//
// Solidity: function planBalanceOf(uint256 planId, uint256 timeStamp, uint256 redemptionTime) view returns(uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCallerSession) PlanBalanceOf(planId *big.Int, timeStamp *big.Int, redemptionTime *big.Int) (struct {
	Balance      *big.Int
	Remainder    *big.Int
	LatestUnlock *big.Int
}, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.PlanBalanceOf(&_IHedgeyVoteTokenVestingPlan.CallOpts, planId, timeStamp, redemptionTime)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 planId) view returns((address,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCaller) Plans(opts *bind.CallOpts, planId *big.Int) (IHedgeyVoteTokenVestingPlanPlan, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenVestingPlan.contract.Call(opts, &out, "plans", planId)

	if err != nil {
		return *new(IHedgeyVoteTokenVestingPlanPlan), err
	}

	out0 := *abi.ConvertType(out[0], new(IHedgeyVoteTokenVestingPlanPlan)).(*IHedgeyVoteTokenVestingPlanPlan)

	return out0, err

}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 planId) view returns((address,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) Plans(planId *big.Int) (IHedgeyVoteTokenVestingPlanPlan, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.Plans(&_IHedgeyVoteTokenVestingPlan.CallOpts, planId)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 planId) view returns((address,uint256,uint256,uint256,uint256,uint256,address,bool))
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCallerSession) Plans(planId *big.Int) (IHedgeyVoteTokenVestingPlanPlan, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.Plans(&_IHedgeyVoteTokenVestingPlan.CallOpts, planId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenVestingPlan.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.TokenOfOwnerByIndex(&_IHedgeyVoteTokenVestingPlan.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.TokenOfOwnerByIndex(&_IHedgeyVoteTokenVestingPlan.CallOpts, owner, index)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x24b6b8c0.
//
// Solidity: function _safeTransfer(address from, address to, uint256 tokenId, bytes data) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) SafeTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "_safeTransfer", from, to, tokenId, data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x24b6b8c0.
//
// Solidity: function _safeTransfer(address from, address to, uint256 tokenId, bytes data) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) SafeTransfer(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.SafeTransfer(&_IHedgeyVoteTokenVestingPlan.TransactOpts, from, to, tokenId, data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x24b6b8c0.
//
// Solidity: function _safeTransfer(address from, address to, uint256 tokenId, bytes data) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) SafeTransfer(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.SafeTransfer(&_IHedgeyVoteTokenVestingPlan.TransactOpts, from, to, tokenId, data)
}

// ChangeVestingPlanAdmin is a paid mutator transaction binding the contract method 0x23ed81fb.
//
// Solidity: function changeVestingPlanAdmin(uint256 planId, address newVestingAdmin) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) ChangeVestingPlanAdmin(opts *bind.TransactOpts, planId *big.Int, newVestingAdmin common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "changeVestingPlanAdmin", planId, newVestingAdmin)
}

// ChangeVestingPlanAdmin is a paid mutator transaction binding the contract method 0x23ed81fb.
//
// Solidity: function changeVestingPlanAdmin(uint256 planId, address newVestingAdmin) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) ChangeVestingPlanAdmin(planId *big.Int, newVestingAdmin common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.ChangeVestingPlanAdmin(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId, newVestingAdmin)
}

// ChangeVestingPlanAdmin is a paid mutator transaction binding the contract method 0x23ed81fb.
//
// Solidity: function changeVestingPlanAdmin(uint256 planId, address newVestingAdmin) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) ChangeVestingPlanAdmin(planId *big.Int, newVestingAdmin common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.ChangeVestingPlanAdmin(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId, newVestingAdmin)
}

// CreatePlan is a paid mutator transaction binding the contract method 0x050fbad7.
//
// Solidity: function createPlan(address recipient, address token, uint256 amount, uint256 start, uint256 cliff, uint256 rate, uint256 period, address vestingAdmin, bool adminTransferOBO) returns(uint256 newPlanId)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) CreatePlan(opts *bind.TransactOpts, recipient common.Address, token common.Address, amount *big.Int, start *big.Int, cliff *big.Int, rate *big.Int, period *big.Int, vestingAdmin common.Address, adminTransferOBO bool) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "createPlan", recipient, token, amount, start, cliff, rate, period, vestingAdmin, adminTransferOBO)
}

// CreatePlan is a paid mutator transaction binding the contract method 0x050fbad7.
//
// Solidity: function createPlan(address recipient, address token, uint256 amount, uint256 start, uint256 cliff, uint256 rate, uint256 period, address vestingAdmin, bool adminTransferOBO) returns(uint256 newPlanId)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) CreatePlan(recipient common.Address, token common.Address, amount *big.Int, start *big.Int, cliff *big.Int, rate *big.Int, period *big.Int, vestingAdmin common.Address, adminTransferOBO bool) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.CreatePlan(&_IHedgeyVoteTokenVestingPlan.TransactOpts, recipient, token, amount, start, cliff, rate, period, vestingAdmin, adminTransferOBO)
}

// CreatePlan is a paid mutator transaction binding the contract method 0x050fbad7.
//
// Solidity: function createPlan(address recipient, address token, uint256 amount, uint256 start, uint256 cliff, uint256 rate, uint256 period, address vestingAdmin, bool adminTransferOBO) returns(uint256 newPlanId)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) CreatePlan(recipient common.Address, token common.Address, amount *big.Int, start *big.Int, cliff *big.Int, rate *big.Int, period *big.Int, vestingAdmin common.Address, adminTransferOBO bool) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.CreatePlan(&_IHedgeyVoteTokenVestingPlan.TransactOpts, recipient, token, amount, start, cliff, rate, period, vestingAdmin, adminTransferOBO)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 planId, address delegatee) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) Delegate(opts *bind.TransactOpts, planId *big.Int, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "delegate", planId, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 planId, address delegatee) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) Delegate(planId *big.Int, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.Delegate(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 planId, address delegatee) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) Delegate(planId *big.Int, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.Delegate(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId, delegatee)
}

// DelegateAll is a paid mutator transaction binding the contract method 0x1f25ccb6.
//
// Solidity: function delegateAll(address token, address delegatee) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) DelegateAll(opts *bind.TransactOpts, token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "delegateAll", token, delegatee)
}

// DelegateAll is a paid mutator transaction binding the contract method 0x1f25ccb6.
//
// Solidity: function delegateAll(address token, address delegatee) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) DelegateAll(token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.DelegateAll(&_IHedgeyVoteTokenVestingPlan.TransactOpts, token, delegatee)
}

// DelegateAll is a paid mutator transaction binding the contract method 0x1f25ccb6.
//
// Solidity: function delegateAll(address token, address delegatee) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) DelegateAll(token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.DelegateAll(&_IHedgeyVoteTokenVestingPlan.TransactOpts, token, delegatee)
}

// DelegatePlans is a paid mutator transaction binding the contract method 0xa8973e2b.
//
// Solidity: function delegatePlans(uint256[] planIds, address[] delegatees) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) DelegatePlans(opts *bind.TransactOpts, planIds []*big.Int, delegatees []common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "delegatePlans", planIds, delegatees)
}

// DelegatePlans is a paid mutator transaction binding the contract method 0xa8973e2b.
//
// Solidity: function delegatePlans(uint256[] planIds, address[] delegatees) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) DelegatePlans(planIds []*big.Int, delegatees []common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.DelegatePlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds, delegatees)
}

// DelegatePlans is a paid mutator transaction binding the contract method 0xa8973e2b.
//
// Solidity: function delegatePlans(uint256[] planIds, address[] delegatees) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) DelegatePlans(planIds []*big.Int, delegatees []common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.DelegatePlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds, delegatees)
}

// FutureRevokePlans is a paid mutator transaction binding the contract method 0xd0c524b7.
//
// Solidity: function futureRevokePlans(uint256[] planIds, uint256 revokeTime) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) FutureRevokePlans(opts *bind.TransactOpts, planIds []*big.Int, revokeTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "futureRevokePlans", planIds, revokeTime)
}

// FutureRevokePlans is a paid mutator transaction binding the contract method 0xd0c524b7.
//
// Solidity: function futureRevokePlans(uint256[] planIds, uint256 revokeTime) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) FutureRevokePlans(planIds []*big.Int, revokeTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.FutureRevokePlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds, revokeTime)
}

// FutureRevokePlans is a paid mutator transaction binding the contract method 0xd0c524b7.
//
// Solidity: function futureRevokePlans(uint256[] planIds, uint256 revokeTime) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) FutureRevokePlans(planIds []*big.Int, revokeTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.FutureRevokePlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds, revokeTime)
}

// PartialRedeemPlans is a paid mutator transaction binding the contract method 0x6b040218.
//
// Solidity: function partialRedeemPlans(uint256[] planIds, uint256 redemptionTime) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) PartialRedeemPlans(opts *bind.TransactOpts, planIds []*big.Int, redemptionTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "partialRedeemPlans", planIds, redemptionTime)
}

// PartialRedeemPlans is a paid mutator transaction binding the contract method 0x6b040218.
//
// Solidity: function partialRedeemPlans(uint256[] planIds, uint256 redemptionTime) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) PartialRedeemPlans(planIds []*big.Int, redemptionTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.PartialRedeemPlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds, redemptionTime)
}

// PartialRedeemPlans is a paid mutator transaction binding the contract method 0x6b040218.
//
// Solidity: function partialRedeemPlans(uint256[] planIds, uint256 redemptionTime) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) PartialRedeemPlans(planIds []*big.Int, redemptionTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.PartialRedeemPlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds, redemptionTime)
}

// RedeemAllPlans is a paid mutator transaction binding the contract method 0xb9bdac2c.
//
// Solidity: function redeemAllPlans() returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) RedeemAllPlans(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "redeemAllPlans")
}

// RedeemAllPlans is a paid mutator transaction binding the contract method 0xb9bdac2c.
//
// Solidity: function redeemAllPlans() returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) RedeemAllPlans() (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.RedeemAllPlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts)
}

// RedeemAllPlans is a paid mutator transaction binding the contract method 0xb9bdac2c.
//
// Solidity: function redeemAllPlans() returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) RedeemAllPlans() (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.RedeemAllPlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts)
}

// RedeemPlans is a paid mutator transaction binding the contract method 0x968b3e59.
//
// Solidity: function redeemPlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) RedeemPlans(opts *bind.TransactOpts, planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "redeemPlans", planIds)
}

// RedeemPlans is a paid mutator transaction binding the contract method 0x968b3e59.
//
// Solidity: function redeemPlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) RedeemPlans(planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.RedeemPlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds)
}

// RedeemPlans is a paid mutator transaction binding the contract method 0x968b3e59.
//
// Solidity: function redeemPlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) RedeemPlans(planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.RedeemPlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds)
}

// RevokePlans is a paid mutator transaction binding the contract method 0x9e866c47.
//
// Solidity: function revokePlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) RevokePlans(opts *bind.TransactOpts, planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "revokePlans", planIds)
}

// RevokePlans is a paid mutator transaction binding the contract method 0x9e866c47.
//
// Solidity: function revokePlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) RevokePlans(planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.RevokePlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds)
}

// RevokePlans is a paid mutator transaction binding the contract method 0x9e866c47.
//
// Solidity: function revokePlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) RevokePlans(planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.RevokePlans(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planIds)
}

// SetupVoting is a paid mutator transaction binding the contract method 0xde6923b8.
//
// Solidity: function setupVoting(uint256 planId) returns(address votingVault)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) SetupVoting(opts *bind.TransactOpts, planId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "setupVoting", planId)
}

// SetupVoting is a paid mutator transaction binding the contract method 0xde6923b8.
//
// Solidity: function setupVoting(uint256 planId) returns(address votingVault)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) SetupVoting(planId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.SetupVoting(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId)
}

// SetupVoting is a paid mutator transaction binding the contract method 0xde6923b8.
//
// Solidity: function setupVoting(uint256 planId) returns(address votingVault)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) SetupVoting(planId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.SetupVoting(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId)
}

// ToggleAdminTransferOBO is a paid mutator transaction binding the contract method 0x7e5d8b4d.
//
// Solidity: function toggleAdminTransferOBO(uint256 planId, bool transferrable) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) ToggleAdminTransferOBO(opts *bind.TransactOpts, planId *big.Int, transferrable bool) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "toggleAdminTransferOBO", planId, transferrable)
}

// ToggleAdminTransferOBO is a paid mutator transaction binding the contract method 0x7e5d8b4d.
//
// Solidity: function toggleAdminTransferOBO(uint256 planId, bool transferrable) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) ToggleAdminTransferOBO(planId *big.Int, transferrable bool) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.ToggleAdminTransferOBO(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId, transferrable)
}

// ToggleAdminTransferOBO is a paid mutator transaction binding the contract method 0x7e5d8b4d.
//
// Solidity: function toggleAdminTransferOBO(uint256 planId, bool transferrable) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) ToggleAdminTransferOBO(planId *big.Int, transferrable bool) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.ToggleAdminTransferOBO(&_IHedgeyVoteTokenVestingPlan.TransactOpts, planId, transferrable)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.TransferFrom(&_IHedgeyVoteTokenVestingPlan.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenVestingPlan.Contract.TransferFrom(&_IHedgeyVoteTokenVestingPlan.TransactOpts, from, to, tokenId)
}

// IHedgeyVoteTokenVestingPlanPlanCreatedIterator is returned from FilterPlanCreated and is used to iterate over the raw logs and unpacked data for PlanCreated events raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanCreatedIterator struct {
	Event *IHedgeyVoteTokenVestingPlanPlanCreated // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenVestingPlanPlanCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenVestingPlanPlanCreated)
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
		it.Event = new(IHedgeyVoteTokenVestingPlanPlanCreated)
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
func (it *IHedgeyVoteTokenVestingPlanPlanCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenVestingPlanPlanCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenVestingPlanPlanCreated represents a PlanCreated event raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanCreated struct {
	NewPlanId        *big.Int
	Recipient        common.Address
	Token            common.Address
	Amount           *big.Int
	Start            *big.Int
	Cliff            *big.Int
	End              *big.Int
	Rate             *big.Int
	Period           *big.Int
	VestingAdmin     common.Address
	AdminTransferOBO bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPlanCreated is a free log retrieval operation binding the contract event 0x6d5fb3665416b633057b4e53641a7dec63a802702a55e386df478871ea22af9b.
//
// Solidity: event PlanCreated(uint256 indexed newPlanId, address indexed recipient, address indexed token, uint256 amount, uint256 start, uint256 cliff, uint256 end, uint256 rate, uint256 period, address vestingAdmin, bool adminTransferOBO)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) FilterPlanCreated(opts *bind.FilterOpts, newPlanId []*big.Int, recipient []common.Address, token []common.Address) (*IHedgeyVoteTokenVestingPlanPlanCreatedIterator, error) {

	var newPlanIdRule []interface{}
	for _, newPlanIdItem := range newPlanId {
		newPlanIdRule = append(newPlanIdRule, newPlanIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.FilterLogs(opts, "PlanCreated", newPlanIdRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanPlanCreatedIterator{contract: _IHedgeyVoteTokenVestingPlan.contract, event: "PlanCreated", logs: logs, sub: sub}, nil
}

// WatchPlanCreated is a free log subscription operation binding the contract event 0x6d5fb3665416b633057b4e53641a7dec63a802702a55e386df478871ea22af9b.
//
// Solidity: event PlanCreated(uint256 indexed newPlanId, address indexed recipient, address indexed token, uint256 amount, uint256 start, uint256 cliff, uint256 end, uint256 rate, uint256 period, address vestingAdmin, bool adminTransferOBO)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) WatchPlanCreated(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenVestingPlanPlanCreated, newPlanId []*big.Int, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var newPlanIdRule []interface{}
	for _, newPlanIdItem := range newPlanId {
		newPlanIdRule = append(newPlanIdRule, newPlanIdItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.WatchLogs(opts, "PlanCreated", newPlanIdRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenVestingPlanPlanCreated)
				if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanCreated", log); err != nil {
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

// ParsePlanCreated is a log parse operation binding the contract event 0x6d5fb3665416b633057b4e53641a7dec63a802702a55e386df478871ea22af9b.
//
// Solidity: event PlanCreated(uint256 indexed newPlanId, address indexed recipient, address indexed token, uint256 amount, uint256 start, uint256 cliff, uint256 end, uint256 rate, uint256 period, address vestingAdmin, bool adminTransferOBO)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) ParsePlanCreated(log types.Log) (*IHedgeyVoteTokenVestingPlanPlanCreated, error) {
	event := new(IHedgeyVoteTokenVestingPlanPlanCreated)
	if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenVestingPlanPlanRedeemedIterator is returned from FilterPlanRedeemed and is used to iterate over the raw logs and unpacked data for PlanRedeemed events raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanRedeemedIterator struct {
	Event *IHedgeyVoteTokenVestingPlanPlanRedeemed // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenVestingPlanPlanRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenVestingPlanPlanRedeemed)
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
		it.Event = new(IHedgeyVoteTokenVestingPlanPlanRedeemed)
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
func (it *IHedgeyVoteTokenVestingPlanPlanRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenVestingPlanPlanRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenVestingPlanPlanRedeemed represents a PlanRedeemed event raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanRedeemed struct {
	PlanId       *big.Int
	Balance      *big.Int
	Remainder    *big.Int
	LatestUnlock *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPlanRedeemed is a free log retrieval operation binding the contract event 0xa6faee2246474597b6de7c76bf9a45d256737543cb0806e6e805b55b38c7663f.
//
// Solidity: event PlanRedeemed(uint256 planId, uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) FilterPlanRedeemed(opts *bind.FilterOpts) (*IHedgeyVoteTokenVestingPlanPlanRedeemedIterator, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.FilterLogs(opts, "PlanRedeemed")
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanPlanRedeemedIterator{contract: _IHedgeyVoteTokenVestingPlan.contract, event: "PlanRedeemed", logs: logs, sub: sub}, nil
}

// WatchPlanRedeemed is a free log subscription operation binding the contract event 0xa6faee2246474597b6de7c76bf9a45d256737543cb0806e6e805b55b38c7663f.
//
// Solidity: event PlanRedeemed(uint256 planId, uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) WatchPlanRedeemed(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenVestingPlanPlanRedeemed) (event.Subscription, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.WatchLogs(opts, "PlanRedeemed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenVestingPlanPlanRedeemed)
				if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanRedeemed", log); err != nil {
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

// ParsePlanRedeemed is a log parse operation binding the contract event 0xa6faee2246474597b6de7c76bf9a45d256737543cb0806e6e805b55b38c7663f.
//
// Solidity: event PlanRedeemed(uint256 planId, uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) ParsePlanRedeemed(log types.Log) (*IHedgeyVoteTokenVestingPlanPlanRedeemed, error) {
	event := new(IHedgeyVoteTokenVestingPlanPlanRedeemed)
	if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenVestingPlanPlanRevokedIterator is returned from FilterPlanRevoked and is used to iterate over the raw logs and unpacked data for PlanRevoked events raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanRevokedIterator struct {
	Event *IHedgeyVoteTokenVestingPlanPlanRevoked // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenVestingPlanPlanRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenVestingPlanPlanRevoked)
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
		it.Event = new(IHedgeyVoteTokenVestingPlanPlanRevoked)
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
func (it *IHedgeyVoteTokenVestingPlanPlanRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenVestingPlanPlanRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenVestingPlanPlanRevoked represents a PlanRevoked event raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanRevoked struct {
	PlanId    *big.Int
	Balance   *big.Int
	Remainder *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPlanRevoked is a free log retrieval operation binding the contract event 0xa4489f0d65c1250dc8e830211cb0442bfcbf6a32300cb9cfa67f2b2176bd18f3.
//
// Solidity: event PlanRevoked(uint256 planId, uint256 balance, uint256 remainder)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) FilterPlanRevoked(opts *bind.FilterOpts) (*IHedgeyVoteTokenVestingPlanPlanRevokedIterator, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.FilterLogs(opts, "PlanRevoked")
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanPlanRevokedIterator{contract: _IHedgeyVoteTokenVestingPlan.contract, event: "PlanRevoked", logs: logs, sub: sub}, nil
}

// WatchPlanRevoked is a free log subscription operation binding the contract event 0xa4489f0d65c1250dc8e830211cb0442bfcbf6a32300cb9cfa67f2b2176bd18f3.
//
// Solidity: event PlanRevoked(uint256 planId, uint256 balance, uint256 remainder)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) WatchPlanRevoked(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenVestingPlanPlanRevoked) (event.Subscription, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.WatchLogs(opts, "PlanRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenVestingPlanPlanRevoked)
				if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanRevoked", log); err != nil {
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

// ParsePlanRevoked is a log parse operation binding the contract event 0xa4489f0d65c1250dc8e830211cb0442bfcbf6a32300cb9cfa67f2b2176bd18f3.
//
// Solidity: event PlanRevoked(uint256 planId, uint256 balance, uint256 remainder)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) ParsePlanRevoked(log types.Log) (*IHedgeyVoteTokenVestingPlanPlanRevoked, error) {
	event := new(IHedgeyVoteTokenVestingPlanPlanRevoked)
	if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdminIterator is returned from FilterPlanTransferredByVestingAdmin and is used to iterate over the raw logs and unpacked data for PlanTransferredByVestingAdmin events raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdminIterator struct {
	Event *IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin)
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
		it.Event = new(IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin)
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
func (it *IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin represents a PlanTransferredByVestingAdmin event raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin struct {
	TokenId *big.Int
	From    common.Address
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPlanTransferredByVestingAdmin is a free log retrieval operation binding the contract event 0x52f56ec49519b0e1b434f05fcf39068afe732dec973010ca61efa6e0a087bdeb.
//
// Solidity: event PlanTransferredByVestingAdmin(uint256 tokenId, address from, address to)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) FilterPlanTransferredByVestingAdmin(opts *bind.FilterOpts) (*IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdminIterator, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.FilterLogs(opts, "PlanTransferredByVestingAdmin")
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdminIterator{contract: _IHedgeyVoteTokenVestingPlan.contract, event: "PlanTransferredByVestingAdmin", logs: logs, sub: sub}, nil
}

// WatchPlanTransferredByVestingAdmin is a free log subscription operation binding the contract event 0x52f56ec49519b0e1b434f05fcf39068afe732dec973010ca61efa6e0a087bdeb.
//
// Solidity: event PlanTransferredByVestingAdmin(uint256 tokenId, address from, address to)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) WatchPlanTransferredByVestingAdmin(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin) (event.Subscription, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.WatchLogs(opts, "PlanTransferredByVestingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin)
				if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanTransferredByVestingAdmin", log); err != nil {
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

// ParsePlanTransferredByVestingAdmin is a log parse operation binding the contract event 0x52f56ec49519b0e1b434f05fcf39068afe732dec973010ca61efa6e0a087bdeb.
//
// Solidity: event PlanTransferredByVestingAdmin(uint256 tokenId, address from, address to)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) ParsePlanTransferredByVestingAdmin(log types.Log) (*IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin, error) {
	event := new(IHedgeyVoteTokenVestingPlanPlanTransferredByVestingAdmin)
	if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanTransferredByVestingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggleIterator is returned from FilterPlanVestingAdminTransferToggle and is used to iterate over the raw logs and unpacked data for PlanVestingAdminTransferToggle events raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggleIterator struct {
	Event *IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle)
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
		it.Event = new(IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle)
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
func (it *IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle represents a PlanVestingAdminTransferToggle event raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle struct {
	PlanId        *big.Int
	Transferrable bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPlanVestingAdminTransferToggle is a free log retrieval operation binding the contract event 0x15e344fba3d60744c30f71eed830999cdb736cf0f1d09f414983adbaeb55d2a6.
//
// Solidity: event PlanVestingAdminTransferToggle(uint256 planId, bool transferrable)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) FilterPlanVestingAdminTransferToggle(opts *bind.FilterOpts) (*IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggleIterator, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.FilterLogs(opts, "PlanVestingAdminTransferToggle")
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggleIterator{contract: _IHedgeyVoteTokenVestingPlan.contract, event: "PlanVestingAdminTransferToggle", logs: logs, sub: sub}, nil
}

// WatchPlanVestingAdminTransferToggle is a free log subscription operation binding the contract event 0x15e344fba3d60744c30f71eed830999cdb736cf0f1d09f414983adbaeb55d2a6.
//
// Solidity: event PlanVestingAdminTransferToggle(uint256 planId, bool transferrable)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) WatchPlanVestingAdminTransferToggle(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle) (event.Subscription, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.WatchLogs(opts, "PlanVestingAdminTransferToggle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle)
				if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanVestingAdminTransferToggle", log); err != nil {
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

// ParsePlanVestingAdminTransferToggle is a log parse operation binding the contract event 0x15e344fba3d60744c30f71eed830999cdb736cf0f1d09f414983adbaeb55d2a6.
//
// Solidity: event PlanVestingAdminTransferToggle(uint256 planId, bool transferrable)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) ParsePlanVestingAdminTransferToggle(log types.Log) (*IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle, error) {
	event := new(IHedgeyVoteTokenVestingPlanPlanVestingAdminTransferToggle)
	if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "PlanVestingAdminTransferToggle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenVestingPlanVestingPlanAdminChangedIterator is returned from FilterVestingPlanAdminChanged and is used to iterate over the raw logs and unpacked data for VestingPlanAdminChanged events raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanVestingPlanAdminChangedIterator struct {
	Event *IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenVestingPlanVestingPlanAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged)
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
		it.Event = new(IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged)
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
func (it *IHedgeyVoteTokenVestingPlanVestingPlanAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenVestingPlanVestingPlanAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged represents a VestingPlanAdminChanged event raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged struct {
	PlanId          *big.Int
	NewVestingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVestingPlanAdminChanged is a free log retrieval operation binding the contract event 0x423d5f6f62e79d3e51b096bd7715b454f212603530cdca551d6a37554889725a.
//
// Solidity: event VestingPlanAdminChanged(uint256 planId, address newVestingAdmin)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) FilterVestingPlanAdminChanged(opts *bind.FilterOpts) (*IHedgeyVoteTokenVestingPlanVestingPlanAdminChangedIterator, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.FilterLogs(opts, "VestingPlanAdminChanged")
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanVestingPlanAdminChangedIterator{contract: _IHedgeyVoteTokenVestingPlan.contract, event: "VestingPlanAdminChanged", logs: logs, sub: sub}, nil
}

// WatchVestingPlanAdminChanged is a free log subscription operation binding the contract event 0x423d5f6f62e79d3e51b096bd7715b454f212603530cdca551d6a37554889725a.
//
// Solidity: event VestingPlanAdminChanged(uint256 planId, address newVestingAdmin)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) WatchVestingPlanAdminChanged(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged) (event.Subscription, error) {

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.WatchLogs(opts, "VestingPlanAdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged)
				if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "VestingPlanAdminChanged", log); err != nil {
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

// ParseVestingPlanAdminChanged is a log parse operation binding the contract event 0x423d5f6f62e79d3e51b096bd7715b454f212603530cdca551d6a37554889725a.
//
// Solidity: event VestingPlanAdminChanged(uint256 planId, address newVestingAdmin)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) ParseVestingPlanAdminChanged(log types.Log) (*IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged, error) {
	event := new(IHedgeyVoteTokenVestingPlanVestingPlanAdminChanged)
	if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "VestingPlanAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenVestingPlanVotingVaultCreatedIterator is returned from FilterVotingVaultCreated and is used to iterate over the raw logs and unpacked data for VotingVaultCreated events raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanVotingVaultCreatedIterator struct {
	Event *IHedgeyVoteTokenVestingPlanVotingVaultCreated // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenVestingPlanVotingVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenVestingPlanVotingVaultCreated)
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
		it.Event = new(IHedgeyVoteTokenVestingPlanVotingVaultCreated)
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
func (it *IHedgeyVoteTokenVestingPlanVotingVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenVestingPlanVotingVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenVestingPlanVotingVaultCreated represents a VotingVaultCreated event raised by the IHedgeyVoteTokenVestingPlan contract.
type IHedgeyVoteTokenVestingPlanVotingVaultCreated struct {
	Id           *big.Int
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVotingVaultCreated is a free log retrieval operation binding the contract event 0xa9649a60c9bf950652949a7d6e3dca992b30cca610efc7df4469b15a8f778ddd.
//
// Solidity: event VotingVaultCreated(uint256 indexed id, address vaultAddress)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) FilterVotingVaultCreated(opts *bind.FilterOpts, id []*big.Int) (*IHedgeyVoteTokenVestingPlanVotingVaultCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.FilterLogs(opts, "VotingVaultCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenVestingPlanVotingVaultCreatedIterator{contract: _IHedgeyVoteTokenVestingPlan.contract, event: "VotingVaultCreated", logs: logs, sub: sub}, nil
}

// WatchVotingVaultCreated is a free log subscription operation binding the contract event 0xa9649a60c9bf950652949a7d6e3dca992b30cca610efc7df4469b15a8f778ddd.
//
// Solidity: event VotingVaultCreated(uint256 indexed id, address vaultAddress)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) WatchVotingVaultCreated(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenVestingPlanVotingVaultCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyVoteTokenVestingPlan.contract.WatchLogs(opts, "VotingVaultCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenVestingPlanVotingVaultCreated)
				if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "VotingVaultCreated", log); err != nil {
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

// ParseVotingVaultCreated is a log parse operation binding the contract event 0xa9649a60c9bf950652949a7d6e3dca992b30cca610efc7df4469b15a8f778ddd.
//
// Solidity: event VotingVaultCreated(uint256 indexed id, address vaultAddress)
func (_IHedgeyVoteTokenVestingPlan *IHedgeyVoteTokenVestingPlanFilterer) ParseVotingVaultCreated(log types.Log) (*IHedgeyVoteTokenVestingPlanVotingVaultCreated, error) {
	event := new(IHedgeyVoteTokenVestingPlanVotingVaultCreated)
	if err := _IHedgeyVoteTokenVestingPlan.contract.UnpackLog(event, "VotingVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
