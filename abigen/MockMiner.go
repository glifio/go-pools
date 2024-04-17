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

// ActiveBeneficiary is an auto generated low-level Go binding around an user-defined struct.
type ActiveBeneficiary struct {
	Beneficiary uint64
	Term        BeneficiaryTerm
}

// BeneficiaryTerm is an auto generated low-level Go binding around an user-defined struct.
type BeneficiaryTerm struct {
	Quota      *big.Int
	UsedQuota  *big.Int
	Expiration uint64
}

// GetBeneficiaryReturn is an auto generated low-level Go binding around an user-defined struct.
type GetBeneficiaryReturn struct {
	Active   ActiveBeneficiary
	Proposed PendingBeneficiaryChange
}

// PendingBeneficiaryChange is an auto generated low-level Go binding around an user-defined struct.
type PendingBeneficiaryChange struct {
	NewBeneficiary        uint64
	NewQuota              *big.Int
	NewExpiration         uint64
	ApprovedByBeneficiary bool
	ApprovedByNominee     bool
}

// MockMinerMetaData contains all meta data concerning the MockMiner contract.
var MockMinerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"activeBeneficiary\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"beneficiary\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used_quota\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"}],\"internalType\":\"structBeneficiaryTerm\",\"name\":\"term\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"name\":\"changeWorkerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBeneficiary\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"beneficiary\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"quota\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"used_quota\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiration\",\"type\":\"uint64\"}],\"internalType\":\"structBeneficiaryTerm\",\"name\":\"term\",\"type\":\"tuple\"}],\"internalType\":\"structActiveBeneficiary\",\"name\":\"active\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"new_beneficiary\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"new_quota\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"new_expiration\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"approved_by_beneficiary\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"approved_by_nominee\",\"type\":\"bool\"}],\"internalType\":\"structPendingBeneficiaryChange\",\"name\":\"proposed\",\"type\":\"tuple\"}],\"internalType\":\"structGetBeneficiaryReturn\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposed\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedBeneficiary\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"new_beneficiary\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"new_quota\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"new_expiration\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"approved_by_beneficiary\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"approved_by_nominee\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_id\",\"type\":\"uint64\"}],\"name\":\"setID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MockMinerABI is the input ABI used to generate the binding from.
// Deprecated: Use MockMinerMetaData.ABI instead.
var MockMinerABI = MockMinerMetaData.ABI

// MockMiner is an auto generated Go binding around an Ethereum contract.
type MockMiner struct {
	MockMinerCaller     // Read-only binding to the contract
	MockMinerTransactor // Write-only binding to the contract
	MockMinerFilterer   // Log filterer for contract events
}

// MockMinerCaller is an auto generated read-only Go binding around an Ethereum contract.
type MockMinerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMinerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MockMinerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMinerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MockMinerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MockMinerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MockMinerSession struct {
	Contract     *MockMiner        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MockMinerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MockMinerCallerSession struct {
	Contract *MockMinerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MockMinerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MockMinerTransactorSession struct {
	Contract     *MockMinerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MockMinerRaw is an auto generated low-level Go binding around an Ethereum contract.
type MockMinerRaw struct {
	Contract *MockMiner // Generic contract binding to access the raw methods on
}

// MockMinerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MockMinerCallerRaw struct {
	Contract *MockMinerCaller // Generic read-only contract binding to access the raw methods on
}

// MockMinerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MockMinerTransactorRaw struct {
	Contract *MockMinerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMockMiner creates a new instance of MockMiner, bound to a specific deployed contract.
func NewMockMiner(address common.Address, backend bind.ContractBackend) (*MockMiner, error) {
	contract, err := bindMockMiner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MockMiner{MockMinerCaller: MockMinerCaller{contract: contract}, MockMinerTransactor: MockMinerTransactor{contract: contract}, MockMinerFilterer: MockMinerFilterer{contract: contract}}, nil
}

// NewMockMinerCaller creates a new read-only instance of MockMiner, bound to a specific deployed contract.
func NewMockMinerCaller(address common.Address, caller bind.ContractCaller) (*MockMinerCaller, error) {
	contract, err := bindMockMiner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MockMinerCaller{contract: contract}, nil
}

// NewMockMinerTransactor creates a new write-only instance of MockMiner, bound to a specific deployed contract.
func NewMockMinerTransactor(address common.Address, transactor bind.ContractTransactor) (*MockMinerTransactor, error) {
	contract, err := bindMockMiner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MockMinerTransactor{contract: contract}, nil
}

// NewMockMinerFilterer creates a new log filterer instance of MockMiner, bound to a specific deployed contract.
func NewMockMinerFilterer(address common.Address, filterer bind.ContractFilterer) (*MockMinerFilterer, error) {
	contract, err := bindMockMiner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MockMinerFilterer{contract: contract}, nil
}

// bindMockMiner binds a generic wrapper to an already deployed contract.
func bindMockMiner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MockMinerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockMiner *MockMinerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMiner.Contract.MockMinerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockMiner *MockMinerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMiner.Contract.MockMinerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockMiner *MockMinerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMiner.Contract.MockMinerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MockMiner *MockMinerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MockMiner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MockMiner *MockMinerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMiner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MockMiner *MockMinerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MockMiner.Contract.contract.Transact(opts, method, params...)
}

// ActiveBeneficiary is a free data retrieval call binding the contract method 0xca5f9df3.
//
// Solidity: function activeBeneficiary() view returns(uint64 beneficiary, (uint256,uint256,uint64) term)
func (_MockMiner *MockMinerCaller) ActiveBeneficiary(opts *bind.CallOpts) (struct {
	Beneficiary uint64
	Term        BeneficiaryTerm
}, error) {
	var out []interface{}
	err := _MockMiner.contract.Call(opts, &out, "activeBeneficiary")

	outstruct := new(struct {
		Beneficiary uint64
		Term        BeneficiaryTerm
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Beneficiary = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Term = *abi.ConvertType(out[1], new(BeneficiaryTerm)).(*BeneficiaryTerm)

	return *outstruct, err

}

// ActiveBeneficiary is a free data retrieval call binding the contract method 0xca5f9df3.
//
// Solidity: function activeBeneficiary() view returns(uint64 beneficiary, (uint256,uint256,uint64) term)
func (_MockMiner *MockMinerSession) ActiveBeneficiary() (struct {
	Beneficiary uint64
	Term        BeneficiaryTerm
}, error) {
	return _MockMiner.Contract.ActiveBeneficiary(&_MockMiner.CallOpts)
}

// ActiveBeneficiary is a free data retrieval call binding the contract method 0xca5f9df3.
//
// Solidity: function activeBeneficiary() view returns(uint64 beneficiary, (uint256,uint256,uint64) term)
func (_MockMiner *MockMinerCallerSession) ActiveBeneficiary() (struct {
	Beneficiary uint64
	Term        BeneficiaryTerm
}, error) {
	return _MockMiner.Contract.ActiveBeneficiary(&_MockMiner.CallOpts)
}

// GetBeneficiary is a free data retrieval call binding the contract method 0x565a2e2c.
//
// Solidity: function getBeneficiary() view returns(((uint64,(uint256,uint256,uint64)),(uint64,uint256,uint64,bool,bool)))
func (_MockMiner *MockMinerCaller) GetBeneficiary(opts *bind.CallOpts) (GetBeneficiaryReturn, error) {
	var out []interface{}
	err := _MockMiner.contract.Call(opts, &out, "getBeneficiary")

	if err != nil {
		return *new(GetBeneficiaryReturn), err
	}

	out0 := *abi.ConvertType(out[0], new(GetBeneficiaryReturn)).(*GetBeneficiaryReturn)

	return out0, err

}

// GetBeneficiary is a free data retrieval call binding the contract method 0x565a2e2c.
//
// Solidity: function getBeneficiary() view returns(((uint64,(uint256,uint256,uint64)),(uint64,uint256,uint64,bool,bool)))
func (_MockMiner *MockMinerSession) GetBeneficiary() (GetBeneficiaryReturn, error) {
	return _MockMiner.Contract.GetBeneficiary(&_MockMiner.CallOpts)
}

// GetBeneficiary is a free data retrieval call binding the contract method 0x565a2e2c.
//
// Solidity: function getBeneficiary() view returns(((uint64,(uint256,uint256,uint64)),(uint64,uint256,uint64,bool,bool)))
func (_MockMiner *MockMinerCallerSession) GetBeneficiary() (GetBeneficiaryReturn, error) {
	return _MockMiner.Contract.GetBeneficiary(&_MockMiner.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint64)
func (_MockMiner *MockMinerCaller) Id(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MockMiner.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint64)
func (_MockMiner *MockMinerSession) Id() (uint64, error) {
	return _MockMiner.Contract.Id(&_MockMiner.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint64)
func (_MockMiner *MockMinerCallerSession) Id() (uint64, error) {
	return _MockMiner.Contract.Id(&_MockMiner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockMiner *MockMinerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockMiner.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockMiner *MockMinerSession) Owner() (common.Address, error) {
	return _MockMiner.Contract.Owner(&_MockMiner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MockMiner *MockMinerCallerSession) Owner() (common.Address, error) {
	return _MockMiner.Contract.Owner(&_MockMiner.CallOpts)
}

// Proposed is a free data retrieval call binding the contract method 0xd1851c92.
//
// Solidity: function proposed() view returns(address)
func (_MockMiner *MockMinerCaller) Proposed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MockMiner.contract.Call(opts, &out, "proposed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proposed is a free data retrieval call binding the contract method 0xd1851c92.
//
// Solidity: function proposed() view returns(address)
func (_MockMiner *MockMinerSession) Proposed() (common.Address, error) {
	return _MockMiner.Contract.Proposed(&_MockMiner.CallOpts)
}

// Proposed is a free data retrieval call binding the contract method 0xd1851c92.
//
// Solidity: function proposed() view returns(address)
func (_MockMiner *MockMinerCallerSession) Proposed() (common.Address, error) {
	return _MockMiner.Contract.Proposed(&_MockMiner.CallOpts)
}

// ProposedBeneficiary is a free data retrieval call binding the contract method 0xc7de66e5.
//
// Solidity: function proposedBeneficiary() view returns(uint64 new_beneficiary, uint256 new_quota, uint64 new_expiration, bool approved_by_beneficiary, bool approved_by_nominee)
func (_MockMiner *MockMinerCaller) ProposedBeneficiary(opts *bind.CallOpts) (struct {
	NewBeneficiary        uint64
	NewQuota              *big.Int
	NewExpiration         uint64
	ApprovedByBeneficiary bool
	ApprovedByNominee     bool
}, error) {
	var out []interface{}
	err := _MockMiner.contract.Call(opts, &out, "proposedBeneficiary")

	outstruct := new(struct {
		NewBeneficiary        uint64
		NewQuota              *big.Int
		NewExpiration         uint64
		ApprovedByBeneficiary bool
		ApprovedByNominee     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewBeneficiary = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.NewQuota = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NewExpiration = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.ApprovedByBeneficiary = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.ApprovedByNominee = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// ProposedBeneficiary is a free data retrieval call binding the contract method 0xc7de66e5.
//
// Solidity: function proposedBeneficiary() view returns(uint64 new_beneficiary, uint256 new_quota, uint64 new_expiration, bool approved_by_beneficiary, bool approved_by_nominee)
func (_MockMiner *MockMinerSession) ProposedBeneficiary() (struct {
	NewBeneficiary        uint64
	NewQuota              *big.Int
	NewExpiration         uint64
	ApprovedByBeneficiary bool
	ApprovedByNominee     bool
}, error) {
	return _MockMiner.Contract.ProposedBeneficiary(&_MockMiner.CallOpts)
}

// ProposedBeneficiary is a free data retrieval call binding the contract method 0xc7de66e5.
//
// Solidity: function proposedBeneficiary() view returns(uint64 new_beneficiary, uint256 new_quota, uint64 new_expiration, bool approved_by_beneficiary, bool approved_by_nominee)
func (_MockMiner *MockMinerCallerSession) ProposedBeneficiary() (struct {
	NewBeneficiary        uint64
	NewQuota              *big.Int
	NewExpiration         uint64
	ApprovedByBeneficiary bool
	ApprovedByNominee     bool
}, error) {
	return _MockMiner.Contract.ProposedBeneficiary(&_MockMiner.CallOpts)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_MockMiner *MockMinerTransactor) ChangeOwnerAddress(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MockMiner.contract.Transact(opts, "changeOwnerAddress", newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_MockMiner *MockMinerSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _MockMiner.Contract.ChangeOwnerAddress(&_MockMiner.TransactOpts, newOwner)
}

// ChangeOwnerAddress is a paid mutator transaction binding the contract method 0x85eac05f.
//
// Solidity: function changeOwnerAddress(address newOwner) returns()
func (_MockMiner *MockMinerTransactorSession) ChangeOwnerAddress(newOwner common.Address) (*types.Transaction, error) {
	return _MockMiner.Contract.ChangeOwnerAddress(&_MockMiner.TransactOpts, newOwner)
}

// ChangeWorkerAddress is a paid mutator transaction binding the contract method 0x9efe79f0.
//
// Solidity: function changeWorkerAddress(uint64 , uint64[] ) returns()
func (_MockMiner *MockMinerTransactor) ChangeWorkerAddress(opts *bind.TransactOpts, arg0 uint64, arg1 []uint64) (*types.Transaction, error) {
	return _MockMiner.contract.Transact(opts, "changeWorkerAddress", arg0, arg1)
}

// ChangeWorkerAddress is a paid mutator transaction binding the contract method 0x9efe79f0.
//
// Solidity: function changeWorkerAddress(uint64 , uint64[] ) returns()
func (_MockMiner *MockMinerSession) ChangeWorkerAddress(arg0 uint64, arg1 []uint64) (*types.Transaction, error) {
	return _MockMiner.Contract.ChangeWorkerAddress(&_MockMiner.TransactOpts, arg0, arg1)
}

// ChangeWorkerAddress is a paid mutator transaction binding the contract method 0x9efe79f0.
//
// Solidity: function changeWorkerAddress(uint64 , uint64[] ) returns()
func (_MockMiner *MockMinerTransactorSession) ChangeWorkerAddress(arg0 uint64, arg1 []uint64) (*types.Transaction, error) {
	return _MockMiner.Contract.ChangeWorkerAddress(&_MockMiner.TransactOpts, arg0, arg1)
}

// SetID is a paid mutator transaction binding the contract method 0xa467e366.
//
// Solidity: function setID(uint64 _id) returns()
func (_MockMiner *MockMinerTransactor) SetID(opts *bind.TransactOpts, _id uint64) (*types.Transaction, error) {
	return _MockMiner.contract.Transact(opts, "setID", _id)
}

// SetID is a paid mutator transaction binding the contract method 0xa467e366.
//
// Solidity: function setID(uint64 _id) returns()
func (_MockMiner *MockMinerSession) SetID(_id uint64) (*types.Transaction, error) {
	return _MockMiner.Contract.SetID(&_MockMiner.TransactOpts, _id)
}

// SetID is a paid mutator transaction binding the contract method 0xa467e366.
//
// Solidity: function setID(uint64 _id) returns()
func (_MockMiner *MockMinerTransactorSession) SetID(_id uint64) (*types.Transaction, error) {
	return _MockMiner.Contract.SetID(&_MockMiner.TransactOpts, _id)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns(uint256 amountWithdrawn)
func (_MockMiner *MockMinerTransactor) WithdrawBalance(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MockMiner.contract.Transact(opts, "withdrawBalance", amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns(uint256 amountWithdrawn)
func (_MockMiner *MockMinerSession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _MockMiner.Contract.WithdrawBalance(&_MockMiner.TransactOpts, amount)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0xda76d5cd.
//
// Solidity: function withdrawBalance(uint256 amount) returns(uint256 amountWithdrawn)
func (_MockMiner *MockMinerTransactorSession) WithdrawBalance(amount *big.Int) (*types.Transaction, error) {
	return _MockMiner.Contract.WithdrawBalance(&_MockMiner.TransactOpts, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MockMiner *MockMinerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MockMiner.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MockMiner *MockMinerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MockMiner.Contract.Fallback(&_MockMiner.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MockMiner *MockMinerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MockMiner.Contract.Fallback(&_MockMiner.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockMiner *MockMinerTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MockMiner.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockMiner *MockMinerSession) Receive() (*types.Transaction, error) {
	return _MockMiner.Contract.Receive(&_MockMiner.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_MockMiner *MockMinerTransactorSession) Receive() (*types.Transaction, error) {
	return _MockMiner.Contract.Receive(&_MockMiner.TransactOpts)
}
