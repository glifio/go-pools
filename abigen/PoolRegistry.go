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

// PoolRegistryMetaData contains all meta data concerning the PoolRegistry contract.
var PoolRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_treasuryFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MAX_TREASURY_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"}],\"name\":\"addPoolToList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPools\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPoolsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"attachPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"poolIDs\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pool\",\"type\":\"uint256\"}],\"name\":\"removePoolFromList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFeeRate\",\"type\":\"uint256\"}],\"name\":\"setTreasuryFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryFeeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"upgradePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162000fe838038062000fe883398101604081905262000034916200020c565b8162000054816001600160a01b0316620000a360201b620009731760201c565b90506001600160a01b0381166200007e57604051635435b28960e11b815260040160405180910390fd5b6200008981620000ef565b50600292909255506001600160a01b03166080526200024d565b60008080620000b28462000142565b9150915081620000c457509192915050565b600080620000d28362000175565b9150915081620000e6575093949350505050565b95945050505050565b600180546001600160a01b03191690556200011f6001600160a01b038216620000a3602090811b6200097317901c565b600080546001600160a01b0319166001600160a01b039290921691909117905550565b600080600160401b600160a01b03831660ff60981b81036200016f57600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620001c95760009250600091505b50811580620001d957503d601614155b15620001ea57506000928392509050565b915091565b80516001600160a01b03811681146200020757600080fd5b919050565b6000806000606084860312156200022257600080fd5b835192506200023460208501620001ef565b91506200024460408501620001ef565b90509250925092565b608051610d826200026660003960005050610d826000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80639a014e241161008c578063e30c397811610066578063e30c3978146101ec578063efde4e641461020c578063f2b3c80914610214578063f2fde38b1461022357600080fd5b80639a014e24146101a6578063a2ff484f146101c6578063ce43303c146101d957600080fd5b80636834f901116100c85780636834f9011461015857806379ba50971461016b5780638152a6c6146101735780638da5cb5b1461018657600080fd5b806312b1ee6c146100ef57806341d1de97146101045780636660103214610141575b600080fd5b6101026100fd366004610b90565b610236565b005b610117610112366004610bb4565b6102b5565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61014a60025481565b604051908152602001610138565b610102610166366004610bcd565b6102ec565b6101026103bb565b610102610181366004610bcd565b610417565b6000546101179073ffffffffffffffffffffffffffffffffffffffff1681565b6101b96101b4366004610bb4565b61058d565b6040516101389190610bef565b6101026101d4366004610b90565b6105ef565b6101026101e7366004610bb4565b6108b5565b6001546101179073ffffffffffffffffffffffffffffffffffffffff1681565b60035461014a565b61014a67016345785d8a000081565b610102610231366004610b90565b610904565b61023e6109ba565b600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600381815481106102c557600080fd5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff16905081565b600354819081111561032a576040517f8f5d631800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003818154811061033d5761033d610c33565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff163314610397576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50600091825260046020908152604083208054600181018255908452922090910155565b60015473ffffffffffffffffffffffffffffffffffffffff16331461040c576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61041533610a0b565b565b6003548190811115610455576040517f8f5d631800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6003818154811061046857610468610c33565b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1633146104c2576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000838152600460205260408120905b815481101561058657838282815481106104ee576104ee610c33565b906000526020600020015403610574578154829061050e90600190610c91565b8154811061051e5761051e610c33565b906000526020600020015482828154811061053b5761053b610c33565b90600052602060002001819055508180548061055957610559610caa565b60019003818190600052602060002001600090559055610586565b8061057e81610cd9565b9150506104d2565b5050505050565b6000818152600460209081526040918290208054835181840281018401909452808452606093928301828280156105e357602002820191906000526020600020905b8154815260200190600101908083116105cf575b50505050509050919050565b6105f76109ba565b60008173ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610644573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106689190610d11565b905060006003828154811061067f5761067f610c33565b60009182526020918290200154604080517fb2bcb002000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9092169350839263b2bcb002926004808401938290030181865afa1580156106f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071a9190610d2a565b610750576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826003838154811061076457610764610c33565b6000918252602082200180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9384161790556040517f124dfd660000000000000000000000000000000000000000000000000000000081528583166004820152909183169063124dfd66906024016020604051808303816000875af1158015610808573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061082c9190610d11565b6040517f59221e380000000000000000000000000000000000000000000000000000000081526004810182905290915073ffffffffffffffffffffffffffffffffffffffff8516906359221e3890602401600060405180830381600087803b15801561089757600080fd5b505af11580156108ab573d6000803e3d6000fd5b5050505050505050565b6108bd6109ba565b67016345785d8a00008111156108ff576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600255565b61090c6109ba565b61092b8173ffffffffffffffffffffffffffffffffffffffff16610973565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b600080600061098184610a9a565b915091508161099257509192915050565b60008061099e83610ae7565b91509150816109b1575093949350505050565b95945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610415576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055610a5273ffffffffffffffffffffffffffffffffffffffff8216610973565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103610ae1576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a8114610b475760009250600091505b50811580610b5657503d601614155b15610b6657506000928392509050565b915091565b73ffffffffffffffffffffffffffffffffffffffff81168114610b8d57600080fd5b50565b600060208284031215610ba257600080fd5b8135610bad81610b6b565b9392505050565b600060208284031215610bc657600080fd5b5035919050565b60008060408385031215610be057600080fd5b50508035926020909101359150565b6020808252825182820181905260009190848201906040850190845b81811015610c2757835183529284019291840191600101610c0b565b50909695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b81810381811115610ca457610ca4610c62565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610d0a57610d0a610c62565b5060010190565b600060208284031215610d2357600080fd5b5051919050565b600060208284031215610d3c57600080fd5b81518015158114610bad57600080fdfea2646970667358221220559e56939b01b754ee18e3886483f9f82a74270a5738807661e88f8f2d7e326064736f6c63430008110033",
}

// PoolRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolRegistryMetaData.ABI instead.
var PoolRegistryABI = PoolRegistryMetaData.ABI

// PoolRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolRegistryMetaData.Bin instead.
var PoolRegistryBin = PoolRegistryMetaData.Bin

// DeployPoolRegistry deploys a new Ethereum contract, binding an instance of PoolRegistry to it.
func DeployPoolRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, _treasuryFeeRate *big.Int, _owner common.Address, _router common.Address) (common.Address, *types.Transaction, *PoolRegistry, error) {
	parsed, err := PoolRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolRegistryBin), backend, _treasuryFeeRate, _owner, _router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolRegistry{PoolRegistryCaller: PoolRegistryCaller{contract: contract}, PoolRegistryTransactor: PoolRegistryTransactor{contract: contract}, PoolRegistryFilterer: PoolRegistryFilterer{contract: contract}}, nil
}

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
