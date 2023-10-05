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
	Bin: "0x60806040523480156200001157600080fd5b50604051620012a2380380620012a28339810160408190526200003491620001dd565b8062000054816001600160a01b03166200009160201b620007be1760201c565b90506001600160a01b0381166200007e57604051635435b28960e11b815260040160405180910390fd5b6200008981620000dd565b50506200020f565b60008080620000a08462000130565b9150915081620000b257509192915050565b600080620000c08362000163565b9150915081620000d4575093949350505050565b95945050505050565b600180546001600160a01b03191690556200010d6001600160a01b03821662000091602090811b620007be17901c565b600080546001600160a01b0319166001600160a01b039290921691909117905550565b600080600160401b600160a01b03831660ff60981b81036200015d57600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620001b75760009250600091505b50811580620001c757503d601614155b15620001d857506000928392509050565b915091565b600060208284031215620001f057600080fd5b81516001600160a01b03811681146200020857600080fd5b9392505050565b611083806200021f6000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063c7afd04b11610066578063c7afd04b14610268578063e30c39781461027b578063ef254abd1461029b578063f2fde38b146102ae57600080fd5b80638da5cb5b146101ca578063b761b6f3146101ea578063c5ae7a111461023257600080fd5b8063334908b2116100c8578063334908b2146101545780636361f6de1461016757806379ba5097146101af5780637df3be51146101b757600080fd5b806319ac0743146100ef5780631a6e649e146101045780631b5f03a614610141575b600080fd5b6101026100fd366004610c40565b6102c1565b005b610117610112366004610ccc565b6102ef565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b61011761014f366004610ce7565b610374565b610102610162366004610d29565b61039d565b61017a610175366004610d60565b610456565b604051610138919081518152602080830151908201526040808301519082015260609182015115159181019190915260800190565b6101026104ef565b6101026101c5366004610dc7565b61054b565b6000546101179073ffffffffffffffffffffffffffffffffffffffff1681565b6102246101f8366004610d60565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b604051908152602001610138565b610117610240366004610ccc565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b610102610276366004610e33565b6105fe565b6001546101179073ffffffffffffffffffffffffffffffffffffffff1681565b6101026102a9366004610dc7565b6106a0565b6101026102bc366004610e92565b61074f565b6102c9610805565b6102ea83836040516102dc929190610eaf565b60405180910390208261039d565b505050565b7fffffffff00000000000000000000000000000000000000000000000000000000811660009081526002602052604081205473ffffffffffffffffffffffffffffffffffffffff168061036e576040517fe6e42cb600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b60006103968383604051610389929190610eaf565b60405180910390206102ef565b9392505050565b6103a5610805565b7fffffffff00000000000000000000000000000000000000000000000000000000821660008181526002602090815260409182902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86169081179091558251908152908101929092527fd0bd3873536bdca53cdb3985d74c5be9fdb9172f717e7d305b3dab7828e1df87910160405180910390a15050565b61048360405180608001604052806000815260200160008152602001600081526020016000151581525090565b50604080516020808201949094528082019290925280518083038201815260608301808352815191850191909120600090815260039485905282902060e084019092528154815260018201546080840152600282015460a084015292015460ff16151560c09091015290565b60015473ffffffffffffffffffffffffffffffffffffffff163314610540576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61054933610856565b565b610553610805565b82811461058c576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b838110156105f7576105e58585838181106105ac576105ac610ebf565b90506020028101906105be9190610eee565b8585858181106105d0576105d0610ebf565b90506020020160208101906100fd9190610e92565b806105ef81610f53565b91505061058f565b5050505050565b3361061161060b306108e5565b846109c8565b73ffffffffffffffffffffffffffffffffffffffff161461065e576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60408051602080820186905281830185905282518083038401815260609092018352815191810191909120600090815260039091522081906105f78282610fb2565b6106a8610805565b8281146106e1576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b838110156105f75761073d85858381811061070157610701610ebf565b90506020020160208101906107169190610ccc565b84848481811061072857610728610ebf565b90506020020160208101906101629190610e92565b8061074781610f53565b9150506106e4565b610757610805565b6107768173ffffffffffffffffffffffffffffffffffffffff166107be565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b60008060006107cc84610b01565b91509150816107dd57509192915050565b6000806107e983610b4e565b91509150816107fc575093949350505050565b95945050505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610549576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905561089d73ffffffffffffffffffffffffffffffffffffffff82166107be565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b604080518082018252601481527f524f555445525f504f4f4c5f5245474953545259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f390c2e8200000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e90602401602060405180830381865afa1580156109a4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061036e9190611017565b60008273ffffffffffffffffffffffffffffffffffffffff1663efde4e646040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a399190611034565b821115610a72576040517f8f5d631800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f41d1de970000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8416906341d1de9790602401602060405180830381865afa158015610add573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103969190611017565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103610b48576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a8114610bae5760009250600091505b50811580610bbd57503d601614155b15610bcd57506000928392509050565b915091565b60008083601f840112610be457600080fd5b50813567ffffffffffffffff811115610bfc57600080fd5b602083019150836020828501011115610c1457600080fd5b9250929050565b73ffffffffffffffffffffffffffffffffffffffff81168114610c3d57600080fd5b50565b600080600060408486031215610c5557600080fd5b833567ffffffffffffffff811115610c6c57600080fd5b610c7886828701610bd2565b9094509250506020840135610c8c81610c1b565b809150509250925092565b80357fffffffff0000000000000000000000000000000000000000000000000000000081168114610cc757600080fd5b919050565b600060208284031215610cde57600080fd5b61039682610c97565b60008060208385031215610cfa57600080fd5b823567ffffffffffffffff811115610d1157600080fd5b610d1d85828601610bd2565b90969095509350505050565b60008060408385031215610d3c57600080fd5b610d4583610c97565b91506020830135610d5581610c1b565b809150509250929050565b60008060408385031215610d7357600080fd5b50508035926020909101359150565b60008083601f840112610d9457600080fd5b50813567ffffffffffffffff811115610dac57600080fd5b6020830191508360208260051b8501011115610c1457600080fd5b60008060008060408587031215610ddd57600080fd5b843567ffffffffffffffff80821115610df557600080fd5b610e0188838901610d82565b90965094506020870135915080821115610e1a57600080fd5b50610e2787828801610d82565b95989497509550505050565b600080600083850360c0811215610e4957600080fd5b843593506020850135925060807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc082011215610e8457600080fd5b506040840190509250925092565b600060208284031215610ea457600080fd5b813561039681610c1b565b8183823760009101908152919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe1843603018112610f2357600080fd5b83018035915067ffffffffffffffff821115610f3e57600080fd5b602001915036819003821315610c1457600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610fab577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b5060010190565b813581556020820135600182015560408201356002820155600381016060830135801515808214610fe257600080fd5b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00835416915060ff8116821783555050505050565b60006020828403121561102957600080fd5b815161039681610c1b565b60006020828403121561104657600080fd5b505191905056fea264697066735822122023b960f753777c9365fdcd82de3f99b2207165773cd62c9ecd10a31df02669e464736f6c63430008110033",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// RouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RouterMetaData.Bin instead.
var RouterBin = RouterMetaData.Bin

// DeployRouter deploys a new Ethereum contract, binding an instance of Router to it.
func DeployRouter(auth *bind.TransactOpts, backend bind.ContractBackend, owner common.Address) (common.Address, *types.Transaction, *Router, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RouterBin), backend, owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

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
