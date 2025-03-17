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

// IHedgeyVoteTokenLockupPlanPlan is an auto generated low-level Go binding around an user-defined struct.
type IHedgeyVoteTokenLockupPlanPlan struct {
	Token  common.Address
	Amount *big.Int
	Start  *big.Int
	Cliff  *big.Int
	Rate   *big.Int
	Period *big.Int
}

// IHedgeyVoteTokenLockupPlanMetaData contains all meta data concerning the IHedgeyVoteTokenLockupPlan contract.
var IHedgeyVoteTokenLockupPlanMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"PlanCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountRedeemed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"planRemainder\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resetDate\",\"type\":\"uint256\"}],\"name\":\"PlanRedeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"}],\"name\":\"VotingVaultCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"_safeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"createPlan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegateAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"delegatees\",\"type\":\"address[]\"}],\"name\":\"delegatePlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"lockedBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockedBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"redemptionTime\",\"type\":\"uint256\"}],\"name\":\"partialRedeemPlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timeStamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redemptionTime\",\"type\":\"uint256\"}],\"name\":\"planBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainder\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestUnlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"}],\"name\":\"plans\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cliff\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"internalType\":\"structIHedgeyVoteTokenLockupPlan.Plan\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeemAllPlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"planIds\",\"type\":\"uint256[]\"}],\"name\":\"redeemPlans\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"}],\"name\":\"setupVoting\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"votingVault\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"planId\",\"type\":\"uint256\"}],\"name\":\"votingVaults\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IHedgeyVoteTokenLockupPlanABI is the input ABI used to generate the binding from.
// Deprecated: Use IHedgeyVoteTokenLockupPlanMetaData.ABI instead.
var IHedgeyVoteTokenLockupPlanABI = IHedgeyVoteTokenLockupPlanMetaData.ABI

// IHedgeyVoteTokenLockupPlan is an auto generated Go binding around an Ethereum contract.
type IHedgeyVoteTokenLockupPlan struct {
	IHedgeyVoteTokenLockupPlanCaller     // Read-only binding to the contract
	IHedgeyVoteTokenLockupPlanTransactor // Write-only binding to the contract
	IHedgeyVoteTokenLockupPlanFilterer   // Log filterer for contract events
}

// IHedgeyVoteTokenLockupPlanCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenLockupPlanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyVoteTokenLockupPlanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenLockupPlanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyVoteTokenLockupPlanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHedgeyVoteTokenLockupPlanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHedgeyVoteTokenLockupPlanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHedgeyVoteTokenLockupPlanSession struct {
	Contract     *IHedgeyVoteTokenLockupPlan // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IHedgeyVoteTokenLockupPlanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHedgeyVoteTokenLockupPlanCallerSession struct {
	Contract *IHedgeyVoteTokenLockupPlanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// IHedgeyVoteTokenLockupPlanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHedgeyVoteTokenLockupPlanTransactorSession struct {
	Contract     *IHedgeyVoteTokenLockupPlanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// IHedgeyVoteTokenLockupPlanRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHedgeyVoteTokenLockupPlanRaw struct {
	Contract *IHedgeyVoteTokenLockupPlan // Generic contract binding to access the raw methods on
}

// IHedgeyVoteTokenLockupPlanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenLockupPlanCallerRaw struct {
	Contract *IHedgeyVoteTokenLockupPlanCaller // Generic read-only contract binding to access the raw methods on
}

// IHedgeyVoteTokenLockupPlanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHedgeyVoteTokenLockupPlanTransactorRaw struct {
	Contract *IHedgeyVoteTokenLockupPlanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHedgeyVoteTokenLockupPlan creates a new instance of IHedgeyVoteTokenLockupPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenLockupPlan(address common.Address, backend bind.ContractBackend) (*IHedgeyVoteTokenLockupPlan, error) {
	contract, err := bindIHedgeyVoteTokenLockupPlan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenLockupPlan{IHedgeyVoteTokenLockupPlanCaller: IHedgeyVoteTokenLockupPlanCaller{contract: contract}, IHedgeyVoteTokenLockupPlanTransactor: IHedgeyVoteTokenLockupPlanTransactor{contract: contract}, IHedgeyVoteTokenLockupPlanFilterer: IHedgeyVoteTokenLockupPlanFilterer{contract: contract}}, nil
}

// NewIHedgeyVoteTokenLockupPlanCaller creates a new read-only instance of IHedgeyVoteTokenLockupPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenLockupPlanCaller(address common.Address, caller bind.ContractCaller) (*IHedgeyVoteTokenLockupPlanCaller, error) {
	contract, err := bindIHedgeyVoteTokenLockupPlan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenLockupPlanCaller{contract: contract}, nil
}

// NewIHedgeyVoteTokenLockupPlanTransactor creates a new write-only instance of IHedgeyVoteTokenLockupPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenLockupPlanTransactor(address common.Address, transactor bind.ContractTransactor) (*IHedgeyVoteTokenLockupPlanTransactor, error) {
	contract, err := bindIHedgeyVoteTokenLockupPlan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenLockupPlanTransactor{contract: contract}, nil
}

// NewIHedgeyVoteTokenLockupPlanFilterer creates a new log filterer instance of IHedgeyVoteTokenLockupPlan, bound to a specific deployed contract.
func NewIHedgeyVoteTokenLockupPlanFilterer(address common.Address, filterer bind.ContractFilterer) (*IHedgeyVoteTokenLockupPlanFilterer, error) {
	contract, err := bindIHedgeyVoteTokenLockupPlan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenLockupPlanFilterer{contract: contract}, nil
}

// bindIHedgeyVoteTokenLockupPlan binds a generic wrapper to an already deployed contract.
func bindIHedgeyVoteTokenLockupPlan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHedgeyVoteTokenLockupPlanMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHedgeyVoteTokenLockupPlan.Contract.IHedgeyVoteTokenLockupPlanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.IHedgeyVoteTokenLockupPlanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.IHedgeyVoteTokenLockupPlanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHedgeyVoteTokenLockupPlan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenLockupPlan.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.BalanceOf(&_IHedgeyVoteTokenLockupPlan.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.BalanceOf(&_IHedgeyVoteTokenLockupPlan.CallOpts, owner)
}

// LockedBalances is a free data retrieval call binding the contract method 0x52e5b7e4.
//
// Solidity: function lockedBalances(address holder, address token) view returns(uint256 lockedBalance)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCaller) LockedBalances(opts *bind.CallOpts, holder common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenLockupPlan.contract.Call(opts, &out, "lockedBalances", holder, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockedBalances is a free data retrieval call binding the contract method 0x52e5b7e4.
//
// Solidity: function lockedBalances(address holder, address token) view returns(uint256 lockedBalance)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) LockedBalances(holder common.Address, token common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.LockedBalances(&_IHedgeyVoteTokenLockupPlan.CallOpts, holder, token)
}

// LockedBalances is a free data retrieval call binding the contract method 0x52e5b7e4.
//
// Solidity: function lockedBalances(address holder, address token) view returns(uint256 lockedBalance)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCallerSession) LockedBalances(holder common.Address, token common.Address) (*big.Int, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.LockedBalances(&_IHedgeyVoteTokenLockupPlan.CallOpts, holder, token)
}

// PlanBalanceOf is a free data retrieval call binding the contract method 0xc7d74fa7.
//
// Solidity: function planBalanceOf(uint256 planId, uint256 timeStamp, uint256 redemptionTime) view returns(uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCaller) PlanBalanceOf(opts *bind.CallOpts, planId *big.Int, timeStamp *big.Int, redemptionTime *big.Int) (struct {
	Balance      *big.Int
	Remainder    *big.Int
	LatestUnlock *big.Int
}, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenLockupPlan.contract.Call(opts, &out, "planBalanceOf", planId, timeStamp, redemptionTime)

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
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) PlanBalanceOf(planId *big.Int, timeStamp *big.Int, redemptionTime *big.Int) (struct {
	Balance      *big.Int
	Remainder    *big.Int
	LatestUnlock *big.Int
}, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.PlanBalanceOf(&_IHedgeyVoteTokenLockupPlan.CallOpts, planId, timeStamp, redemptionTime)
}

// PlanBalanceOf is a free data retrieval call binding the contract method 0xc7d74fa7.
//
// Solidity: function planBalanceOf(uint256 planId, uint256 timeStamp, uint256 redemptionTime) view returns(uint256 balance, uint256 remainder, uint256 latestUnlock)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCallerSession) PlanBalanceOf(planId *big.Int, timeStamp *big.Int, redemptionTime *big.Int) (struct {
	Balance      *big.Int
	Remainder    *big.Int
	LatestUnlock *big.Int
}, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.PlanBalanceOf(&_IHedgeyVoteTokenLockupPlan.CallOpts, planId, timeStamp, redemptionTime)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 planId) view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCaller) Plans(opts *bind.CallOpts, planId *big.Int) (IHedgeyVoteTokenLockupPlanPlan, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenLockupPlan.contract.Call(opts, &out, "plans", planId)

	if err != nil {
		return *new(IHedgeyVoteTokenLockupPlanPlan), err
	}

	out0 := *abi.ConvertType(out[0], new(IHedgeyVoteTokenLockupPlanPlan)).(*IHedgeyVoteTokenLockupPlanPlan)

	return out0, err

}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 planId) view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) Plans(planId *big.Int) (IHedgeyVoteTokenLockupPlanPlan, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.Plans(&_IHedgeyVoteTokenLockupPlan.CallOpts, planId)
}

// Plans is a free data retrieval call binding the contract method 0xb1620616.
//
// Solidity: function plans(uint256 planId) view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCallerSession) Plans(planId *big.Int) (IHedgeyVoteTokenLockupPlanPlan, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.Plans(&_IHedgeyVoteTokenLockupPlan.CallOpts, planId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenLockupPlan.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.TokenOfOwnerByIndex(&_IHedgeyVoteTokenLockupPlan.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.TokenOfOwnerByIndex(&_IHedgeyVoteTokenLockupPlan.CallOpts, owner, index)
}

// VotingVaults is a free data retrieval call binding the contract method 0xa8713ec7.
//
// Solidity: function votingVaults(uint256 planId) view returns(address)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCaller) VotingVaults(opts *bind.CallOpts, planId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IHedgeyVoteTokenLockupPlan.contract.Call(opts, &out, "votingVaults", planId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VotingVaults is a free data retrieval call binding the contract method 0xa8713ec7.
//
// Solidity: function votingVaults(uint256 planId) view returns(address)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) VotingVaults(planId *big.Int) (common.Address, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.VotingVaults(&_IHedgeyVoteTokenLockupPlan.CallOpts, planId)
}

// VotingVaults is a free data retrieval call binding the contract method 0xa8713ec7.
//
// Solidity: function votingVaults(uint256 planId) view returns(address)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanCallerSession) VotingVaults(planId *big.Int) (common.Address, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.VotingVaults(&_IHedgeyVoteTokenLockupPlan.CallOpts, planId)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x24b6b8c0.
//
// Solidity: function _safeTransfer(address from, address to, uint256 tokenId, bytes data) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) SafeTransfer(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "_safeTransfer", from, to, tokenId, data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x24b6b8c0.
//
// Solidity: function _safeTransfer(address from, address to, uint256 tokenId, bytes data) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) SafeTransfer(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.SafeTransfer(&_IHedgeyVoteTokenLockupPlan.TransactOpts, from, to, tokenId, data)
}

// SafeTransfer is a paid mutator transaction binding the contract method 0x24b6b8c0.
//
// Solidity: function _safeTransfer(address from, address to, uint256 tokenId, bytes data) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) SafeTransfer(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.SafeTransfer(&_IHedgeyVoteTokenLockupPlan.TransactOpts, from, to, tokenId, data)
}

// CreatePlan is a paid mutator transaction binding the contract method 0x4e897e16.
//
// Solidity: function createPlan(address recipient, address token, uint256 amount, uint256 start, uint256 cliff, uint256 rate, uint256 period) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) CreatePlan(opts *bind.TransactOpts, recipient common.Address, token common.Address, amount *big.Int, start *big.Int, cliff *big.Int, rate *big.Int, period *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "createPlan", recipient, token, amount, start, cliff, rate, period)
}

// CreatePlan is a paid mutator transaction binding the contract method 0x4e897e16.
//
// Solidity: function createPlan(address recipient, address token, uint256 amount, uint256 start, uint256 cliff, uint256 rate, uint256 period) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) CreatePlan(recipient common.Address, token common.Address, amount *big.Int, start *big.Int, cliff *big.Int, rate *big.Int, period *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.CreatePlan(&_IHedgeyVoteTokenLockupPlan.TransactOpts, recipient, token, amount, start, cliff, rate, period)
}

// CreatePlan is a paid mutator transaction binding the contract method 0x4e897e16.
//
// Solidity: function createPlan(address recipient, address token, uint256 amount, uint256 start, uint256 cliff, uint256 rate, uint256 period) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) CreatePlan(recipient common.Address, token common.Address, amount *big.Int, start *big.Int, cliff *big.Int, rate *big.Int, period *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.CreatePlan(&_IHedgeyVoteTokenLockupPlan.TransactOpts, recipient, token, amount, start, cliff, rate, period)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 planId, address delegatee) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) Delegate(opts *bind.TransactOpts, planId *big.Int, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "delegate", planId, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 planId, address delegatee) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) Delegate(planId *big.Int, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.Delegate(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planId, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x08bbb824.
//
// Solidity: function delegate(uint256 planId, address delegatee) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) Delegate(planId *big.Int, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.Delegate(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planId, delegatee)
}

// DelegateAll is a paid mutator transaction binding the contract method 0x1f25ccb6.
//
// Solidity: function delegateAll(address token, address delegatee) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) DelegateAll(opts *bind.TransactOpts, token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "delegateAll", token, delegatee)
}

// DelegateAll is a paid mutator transaction binding the contract method 0x1f25ccb6.
//
// Solidity: function delegateAll(address token, address delegatee) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) DelegateAll(token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.DelegateAll(&_IHedgeyVoteTokenLockupPlan.TransactOpts, token, delegatee)
}

// DelegateAll is a paid mutator transaction binding the contract method 0x1f25ccb6.
//
// Solidity: function delegateAll(address token, address delegatee) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) DelegateAll(token common.Address, delegatee common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.DelegateAll(&_IHedgeyVoteTokenLockupPlan.TransactOpts, token, delegatee)
}

// DelegatePlans is a paid mutator transaction binding the contract method 0xa8973e2b.
//
// Solidity: function delegatePlans(uint256[] planIds, address[] delegatees) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) DelegatePlans(opts *bind.TransactOpts, planIds []*big.Int, delegatees []common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "delegatePlans", planIds, delegatees)
}

// DelegatePlans is a paid mutator transaction binding the contract method 0xa8973e2b.
//
// Solidity: function delegatePlans(uint256[] planIds, address[] delegatees) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) DelegatePlans(planIds []*big.Int, delegatees []common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.DelegatePlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planIds, delegatees)
}

// DelegatePlans is a paid mutator transaction binding the contract method 0xa8973e2b.
//
// Solidity: function delegatePlans(uint256[] planIds, address[] delegatees) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) DelegatePlans(planIds []*big.Int, delegatees []common.Address) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.DelegatePlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planIds, delegatees)
}

// PartialRedeemPlans is a paid mutator transaction binding the contract method 0x6b040218.
//
// Solidity: function partialRedeemPlans(uint256[] planIds, uint256 redemptionTime) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) PartialRedeemPlans(opts *bind.TransactOpts, planIds []*big.Int, redemptionTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "partialRedeemPlans", planIds, redemptionTime)
}

// PartialRedeemPlans is a paid mutator transaction binding the contract method 0x6b040218.
//
// Solidity: function partialRedeemPlans(uint256[] planIds, uint256 redemptionTime) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) PartialRedeemPlans(planIds []*big.Int, redemptionTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.PartialRedeemPlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planIds, redemptionTime)
}

// PartialRedeemPlans is a paid mutator transaction binding the contract method 0x6b040218.
//
// Solidity: function partialRedeemPlans(uint256[] planIds, uint256 redemptionTime) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) PartialRedeemPlans(planIds []*big.Int, redemptionTime *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.PartialRedeemPlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planIds, redemptionTime)
}

// RedeemAllPlans is a paid mutator transaction binding the contract method 0xb9bdac2c.
//
// Solidity: function redeemAllPlans() returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) RedeemAllPlans(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "redeemAllPlans")
}

// RedeemAllPlans is a paid mutator transaction binding the contract method 0xb9bdac2c.
//
// Solidity: function redeemAllPlans() returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) RedeemAllPlans() (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.RedeemAllPlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts)
}

// RedeemAllPlans is a paid mutator transaction binding the contract method 0xb9bdac2c.
//
// Solidity: function redeemAllPlans() returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) RedeemAllPlans() (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.RedeemAllPlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts)
}

// RedeemPlans is a paid mutator transaction binding the contract method 0x968b3e59.
//
// Solidity: function redeemPlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) RedeemPlans(opts *bind.TransactOpts, planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "redeemPlans", planIds)
}

// RedeemPlans is a paid mutator transaction binding the contract method 0x968b3e59.
//
// Solidity: function redeemPlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) RedeemPlans(planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.RedeemPlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planIds)
}

// RedeemPlans is a paid mutator transaction binding the contract method 0x968b3e59.
//
// Solidity: function redeemPlans(uint256[] planIds) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) RedeemPlans(planIds []*big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.RedeemPlans(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planIds)
}

// SetupVoting is a paid mutator transaction binding the contract method 0xde6923b8.
//
// Solidity: function setupVoting(uint256 planId) returns(address votingVault)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) SetupVoting(opts *bind.TransactOpts, planId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "setupVoting", planId)
}

// SetupVoting is a paid mutator transaction binding the contract method 0xde6923b8.
//
// Solidity: function setupVoting(uint256 planId) returns(address votingVault)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) SetupVoting(planId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.SetupVoting(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planId)
}

// SetupVoting is a paid mutator transaction binding the contract method 0xde6923b8.
//
// Solidity: function setupVoting(uint256 planId) returns(address votingVault)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) SetupVoting(planId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.SetupVoting(&_IHedgeyVoteTokenLockupPlan.TransactOpts, planId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.TransferFrom(&_IHedgeyVoteTokenLockupPlan.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHedgeyVoteTokenLockupPlan.Contract.TransferFrom(&_IHedgeyVoteTokenLockupPlan.TransactOpts, from, to, tokenId)
}

// IHedgeyVoteTokenLockupPlanPlanCreatedIterator is returned from FilterPlanCreated and is used to iterate over the raw logs and unpacked data for PlanCreated events raised by the IHedgeyVoteTokenLockupPlan contract.
type IHedgeyVoteTokenLockupPlanPlanCreatedIterator struct {
	Event *IHedgeyVoteTokenLockupPlanPlanCreated // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenLockupPlanPlanCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenLockupPlanPlanCreated)
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
		it.Event = new(IHedgeyVoteTokenLockupPlanPlanCreated)
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
func (it *IHedgeyVoteTokenLockupPlanPlanCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenLockupPlanPlanCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenLockupPlanPlanCreated represents a PlanCreated event raised by the IHedgeyVoteTokenLockupPlan contract.
type IHedgeyVoteTokenLockupPlanPlanCreated struct {
	Id        *big.Int
	Recipient common.Address
	Token     common.Address
	Amount    *big.Int
	Start     *big.Int
	Cliff     *big.Int
	End       *big.Int
	Rate      *big.Int
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPlanCreated is a free log retrieval operation binding the contract event 0xe7d9b7fd810a51c7f2f160d0c100b1bb756592fdeaf6b9b84425b44eca133e9b.
//
// Solidity: event PlanCreated(uint256 indexed id, address indexed recipient, address indexed token, uint256 amount, uint256 start, uint256 cliff, uint256 end, uint256 rate, uint256 period)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) FilterPlanCreated(opts *bind.FilterOpts, id []*big.Int, recipient []common.Address, token []common.Address) (*IHedgeyVoteTokenLockupPlanPlanCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IHedgeyVoteTokenLockupPlan.contract.FilterLogs(opts, "PlanCreated", idRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenLockupPlanPlanCreatedIterator{contract: _IHedgeyVoteTokenLockupPlan.contract, event: "PlanCreated", logs: logs, sub: sub}, nil
}

// WatchPlanCreated is a free log subscription operation binding the contract event 0xe7d9b7fd810a51c7f2f160d0c100b1bb756592fdeaf6b9b84425b44eca133e9b.
//
// Solidity: event PlanCreated(uint256 indexed id, address indexed recipient, address indexed token, uint256 amount, uint256 start, uint256 cliff, uint256 end, uint256 rate, uint256 period)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) WatchPlanCreated(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenLockupPlanPlanCreated, id []*big.Int, recipient []common.Address, token []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _IHedgeyVoteTokenLockupPlan.contract.WatchLogs(opts, "PlanCreated", idRule, recipientRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenLockupPlanPlanCreated)
				if err := _IHedgeyVoteTokenLockupPlan.contract.UnpackLog(event, "PlanCreated", log); err != nil {
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

// ParsePlanCreated is a log parse operation binding the contract event 0xe7d9b7fd810a51c7f2f160d0c100b1bb756592fdeaf6b9b84425b44eca133e9b.
//
// Solidity: event PlanCreated(uint256 indexed id, address indexed recipient, address indexed token, uint256 amount, uint256 start, uint256 cliff, uint256 end, uint256 rate, uint256 period)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) ParsePlanCreated(log types.Log) (*IHedgeyVoteTokenLockupPlanPlanCreated, error) {
	event := new(IHedgeyVoteTokenLockupPlanPlanCreated)
	if err := _IHedgeyVoteTokenLockupPlan.contract.UnpackLog(event, "PlanCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenLockupPlanPlanRedeemedIterator is returned from FilterPlanRedeemed and is used to iterate over the raw logs and unpacked data for PlanRedeemed events raised by the IHedgeyVoteTokenLockupPlan contract.
type IHedgeyVoteTokenLockupPlanPlanRedeemedIterator struct {
	Event *IHedgeyVoteTokenLockupPlanPlanRedeemed // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenLockupPlanPlanRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenLockupPlanPlanRedeemed)
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
		it.Event = new(IHedgeyVoteTokenLockupPlanPlanRedeemed)
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
func (it *IHedgeyVoteTokenLockupPlanPlanRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenLockupPlanPlanRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenLockupPlanPlanRedeemed represents a PlanRedeemed event raised by the IHedgeyVoteTokenLockupPlan contract.
type IHedgeyVoteTokenLockupPlanPlanRedeemed struct {
	Id             *big.Int
	AmountRedeemed *big.Int
	PlanRemainder  *big.Int
	ResetDate      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPlanRedeemed is a free log retrieval operation binding the contract event 0xa6faee2246474597b6de7c76bf9a45d256737543cb0806e6e805b55b38c7663f.
//
// Solidity: event PlanRedeemed(uint256 indexed id, uint256 amountRedeemed, uint256 planRemainder, uint256 resetDate)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) FilterPlanRedeemed(opts *bind.FilterOpts, id []*big.Int) (*IHedgeyVoteTokenLockupPlanPlanRedeemedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyVoteTokenLockupPlan.contract.FilterLogs(opts, "PlanRedeemed", idRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenLockupPlanPlanRedeemedIterator{contract: _IHedgeyVoteTokenLockupPlan.contract, event: "PlanRedeemed", logs: logs, sub: sub}, nil
}

// WatchPlanRedeemed is a free log subscription operation binding the contract event 0xa6faee2246474597b6de7c76bf9a45d256737543cb0806e6e805b55b38c7663f.
//
// Solidity: event PlanRedeemed(uint256 indexed id, uint256 amountRedeemed, uint256 planRemainder, uint256 resetDate)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) WatchPlanRedeemed(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenLockupPlanPlanRedeemed, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyVoteTokenLockupPlan.contract.WatchLogs(opts, "PlanRedeemed", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenLockupPlanPlanRedeemed)
				if err := _IHedgeyVoteTokenLockupPlan.contract.UnpackLog(event, "PlanRedeemed", log); err != nil {
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
// Solidity: event PlanRedeemed(uint256 indexed id, uint256 amountRedeemed, uint256 planRemainder, uint256 resetDate)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) ParsePlanRedeemed(log types.Log) (*IHedgeyVoteTokenLockupPlanPlanRedeemed, error) {
	event := new(IHedgeyVoteTokenLockupPlanPlanRedeemed)
	if err := _IHedgeyVoteTokenLockupPlan.contract.UnpackLog(event, "PlanRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IHedgeyVoteTokenLockupPlanVotingVaultCreatedIterator is returned from FilterVotingVaultCreated and is used to iterate over the raw logs and unpacked data for VotingVaultCreated events raised by the IHedgeyVoteTokenLockupPlan contract.
type IHedgeyVoteTokenLockupPlanVotingVaultCreatedIterator struct {
	Event *IHedgeyVoteTokenLockupPlanVotingVaultCreated // Event containing the contract specifics and raw log

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
func (it *IHedgeyVoteTokenLockupPlanVotingVaultCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IHedgeyVoteTokenLockupPlanVotingVaultCreated)
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
		it.Event = new(IHedgeyVoteTokenLockupPlanVotingVaultCreated)
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
func (it *IHedgeyVoteTokenLockupPlanVotingVaultCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IHedgeyVoteTokenLockupPlanVotingVaultCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IHedgeyVoteTokenLockupPlanVotingVaultCreated represents a VotingVaultCreated event raised by the IHedgeyVoteTokenLockupPlan contract.
type IHedgeyVoteTokenLockupPlanVotingVaultCreated struct {
	Id           *big.Int
	VaultAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVotingVaultCreated is a free log retrieval operation binding the contract event 0xa9649a60c9bf950652949a7d6e3dca992b30cca610efc7df4469b15a8f778ddd.
//
// Solidity: event VotingVaultCreated(uint256 indexed id, address vaultAddress)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) FilterVotingVaultCreated(opts *bind.FilterOpts, id []*big.Int) (*IHedgeyVoteTokenLockupPlanVotingVaultCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyVoteTokenLockupPlan.contract.FilterLogs(opts, "VotingVaultCreated", idRule)
	if err != nil {
		return nil, err
	}
	return &IHedgeyVoteTokenLockupPlanVotingVaultCreatedIterator{contract: _IHedgeyVoteTokenLockupPlan.contract, event: "VotingVaultCreated", logs: logs, sub: sub}, nil
}

// WatchVotingVaultCreated is a free log subscription operation binding the contract event 0xa9649a60c9bf950652949a7d6e3dca992b30cca610efc7df4469b15a8f778ddd.
//
// Solidity: event VotingVaultCreated(uint256 indexed id, address vaultAddress)
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) WatchVotingVaultCreated(opts *bind.WatchOpts, sink chan<- *IHedgeyVoteTokenLockupPlanVotingVaultCreated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IHedgeyVoteTokenLockupPlan.contract.WatchLogs(opts, "VotingVaultCreated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IHedgeyVoteTokenLockupPlanVotingVaultCreated)
				if err := _IHedgeyVoteTokenLockupPlan.contract.UnpackLog(event, "VotingVaultCreated", log); err != nil {
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
func (_IHedgeyVoteTokenLockupPlan *IHedgeyVoteTokenLockupPlanFilterer) ParseVotingVaultCreated(log types.Log) (*IHedgeyVoteTokenLockupPlanVotingVaultCreated, error) {
	event := new(IHedgeyVoteTokenLockupPlanVotingVaultCreated)
	if err := _IHedgeyVoteTokenLockupPlan.contract.UnpackLog(event, "VotingVaultCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
