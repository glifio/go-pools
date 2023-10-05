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

// PreStakeMetaData contains all meta data concerning the PreStake contract.
var PreStakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"},{\"internalType\":\"contractIPoolToken\",\"name\":\"_poolToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApprovePool\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approvePoolToTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"convertFILtoWFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000d9338038062000d93833981016040819052620000349162000209565b826001600160a01b0381166200005d5760405163e6c4247b60e01b815260040160405180910390fd5b6200007c816001600160a01b0316620000f760201b620008e81760201c565b600080546001600160a01b0319166001600160a01b039290921691821781556040517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350600280546001600160a01b039384166001600160a01b03199182161790915560038054929093169116179055506200025d565b60008080620001068462000143565b91509150816200011857509192915050565b600080620001268362000176565b91509150816200013a575093949350505050565b95945050505050565b600080600160401b600160a01b03831660ff60981b81036200017057600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620001ca5760009250600091505b50811580620001da57503d601614155b15620001eb57506000928392509050565b915091565b6001600160a01b03811681146200020657600080fd5b50565b6000806000606084860312156200021f57600080fd5b83516200022c81620001f0565b60208501519093506200023f81620001f0565b60408501519092506200025281620001f0565b809150509250925092565b610b26806200026d6000396000f3fe60806040526004361061009a5760003560e01c80637ea4e0e511610069578063ec18154e1161004e578063ec18154e1461018b578063f2fde38b146101ae578063f340fa01146101ce576100ab565b80637ea4e0e5146101145780638da5cb5b14610134576100ab565b80631a3ece7f146100b55780632b968958146100ca57806347e7ef24146100df57806379ba5097146100ff576100ab565b366100ab576100a933346101e1565b005b6100a933346101e1565b3480156100c157600080fd5b506100a96102f4565b3480156100d657600080fd5b506100a96103ca565b3480156100eb57600080fd5b506100a96100fa366004610a29565b6104bf565b34801561010b57600080fd5b506100a961056f565b34801561012057600080fd5b506100a961012f366004610a29565b61063c565b34801561014057600080fd5b506000546101619073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561019757600080fd5b506101a0610773565b604051908152602001610182565b3480156101ba57600080fd5b506100a96101c9366004610a53565b610819565b6100a96101dc366004610a53565b6108db565b6102008273ffffffffffffffffffffffffffffffffffffffff166108e8565b6003546040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8084166004830152602482018590529294509116906340c10f19906044016020604051808303816000875af115801561027b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029f9190610a75565b508173ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c826040516102e891815260200190565b60405180910390a25050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610345576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d0e30db0476040518263ffffffff1660e01b81526004016000604051808303818588803b1580156103af57600080fd5b505af11580156103c3573d6000803e3d6000fd5b5050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461041b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff161561046b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560405133907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3565b6002546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810183905273ffffffffffffffffffffffffffffffffffffffff909116906323b872dd906064016020604051808303816000875af115801561053c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105609190610a75565b5061056b82826101e1565b5050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146105c0576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff16331461068d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152602482018490529091169063095ea7b3906044016020604051808303816000875af1158015610706573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072a9190610a75565b508173ffffffffffffffffffffffffffffffffffffffff167f349c23e02899ddce96ca5c007ea92efb73b9dc4db1d0ffb87de71f87ac54a953826040516102e891815260200190565b6002546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600091479173ffffffffffffffffffffffffffffffffffffffff909116906370a0823190602401602060405180830381865afa1580156107e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080a9190610a97565b6108149190610ab0565b905090565b60005473ffffffffffffffffffffffffffffffffffffffff16331461086a576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831690811790915560405133907f3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a90600090a350565b6108e581346101e1565b50565b60008060006108f68461092f565b915091508161090757509192915050565b6000806109138361097c565b9150915081610926575093949350505050565b95945050505050565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103610976576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a81146109dc5760009250600091505b508115806109eb57503d601614155b156109fb57506000928392509050565b915091565b803573ffffffffffffffffffffffffffffffffffffffff81168114610a2457600080fd5b919050565b60008060408385031215610a3c57600080fd5b610a4583610a00565b946020939093013593505050565b600060208284031215610a6557600080fd5b610a6e82610a00565b9392505050565b600060208284031215610a8757600080fd5b81518015158114610a6e57600080fd5b600060208284031215610aa957600080fd5b5051919050565b80820180821115610aea577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220b7b03a31a3ce0f286c31c7b77f1281262d6ddccdac979c372ef7bd436c98690364736f6c63430008110033",
}

// PreStakeABI is the input ABI used to generate the binding from.
// Deprecated: Use PreStakeMetaData.ABI instead.
var PreStakeABI = PreStakeMetaData.ABI

// PreStakeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PreStakeMetaData.Bin instead.
var PreStakeBin = PreStakeMetaData.Bin

// DeployPreStake deploys a new Ethereum contract, binding an instance of PreStake to it.
func DeployPreStake(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _wFIL common.Address, _poolToken common.Address) (common.Address, *types.Transaction, *PreStake, error) {
	parsed, err := PreStakeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PreStakeBin), backend, _owner, _wFIL, _poolToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PreStake{PreStakeCaller: PreStakeCaller{contract: contract}, PreStakeTransactor: PreStakeTransactor{contract: contract}, PreStakeFilterer: PreStakeFilterer{contract: contract}}, nil
}

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

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PreStake *PreStakeTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PreStake *PreStakeSession) RevokeOwnership() (*types.Transaction, error) {
	return _PreStake.Contract.RevokeOwnership(&_PreStake.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PreStake *PreStakeTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _PreStake.Contract.RevokeOwnership(&_PreStake.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PreStake *PreStakeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PreStake.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PreStake *PreStakeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PreStake.Contract.TransferOwnership(&_PreStake.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PreStake *PreStakeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PreStake.Contract.TransferOwnership(&_PreStake.TransactOpts, _newOwner)
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

// PreStakeOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the PreStake contract.
type PreStakeOwnershipPendingIterator struct {
	Event *PreStakeOwnershipPending // Event containing the contract specifics and raw log

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
func (it *PreStakeOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PreStakeOwnershipPending)
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
		it.Event = new(PreStakeOwnershipPending)
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
func (it *PreStakeOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PreStakeOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PreStakeOwnershipPending represents a OwnershipPending event raised by the PreStake contract.
type PreStakeOwnershipPending struct {
	CurrentOwner common.Address
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_PreStake *PreStakeFilterer) FilterOwnershipPending(opts *bind.FilterOpts, currentOwner []common.Address, pendingOwner []common.Address) (*PreStakeOwnershipPendingIterator, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PreStake.contract.FilterLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PreStakeOwnershipPendingIterator{contract: _PreStake.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_PreStake *PreStakeFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *PreStakeOwnershipPending, currentOwner []common.Address, pendingOwner []common.Address) (event.Subscription, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PreStake.contract.WatchLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PreStakeOwnershipPending)
				if err := _PreStake.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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

// ParseOwnershipPending is a log parse operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_PreStake *PreStakeFilterer) ParseOwnershipPending(log types.Log) (*PreStakeOwnershipPending, error) {
	event := new(PreStakeOwnershipPending)
	if err := _PreStake.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PreStakeOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PreStake.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PreStakeOwnershipTransferredIterator{contract: _PreStake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PreStakeOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PreStake.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PreStake *PreStakeFilterer) ParseOwnershipTransferred(log types.Log) (*PreStakeOwnershipTransferred, error) {
	event := new(PreStakeOwnershipTransferred)
	if err := _PreStake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
