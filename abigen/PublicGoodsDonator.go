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

// PublicGoodsDonatorMetaData contains all meta data concerning the PublicGoodsDonator contract.
var PublicGoodsDonatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"contractIPreStake\",\"name\":\"_preStake\",\"type\":\"address\"},{\"internalType\":\"contractIWFIL\",\"name\":\"_wFIL\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"_liquidStakingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"donationAmount\",\"type\":\"uint256\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawFunds\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"donationPercent\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"donationPercent\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preStake\",\"outputs\":[{\"internalType\":\"contractIPreStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x610100604052670de0b6b3a76400006080523480156200001e57600080fd5b506040516200142738038062001427833981016040819052620000419162000202565b836001600160a01b0381166200006a5760405163e6c4247b60e01b815260040160405180910390fd5b62000089816001600160a01b0316620000f060201b62000d0e1760201c565b600080546001600160a01b0319166001600160a01b039290921691821781556040517f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3506001600160a01b0392831660a05290821660c0521660e052506200026a565b60008080620000ff846200013c565b91509150816200011157509192915050565b6000806200011f836200016f565b915091508162000133575093949350505050565b95945050505050565b600080600160401b600160a01b03831660ff60981b81036200016957600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620001c35760009250600091505b50811580620001d357503d601614155b15620001e457506000928392509050565b915091565b6001600160a01b0381168114620001ff57600080fd5b50565b600080600080608085870312156200021957600080fd5b84516200022681620001e9565b60208601519094506200023981620001e9565b60408601519093506200024c81620001e9565b60608601519092506200025f81620001e9565b939692955090935050565b60805160a05160c05160e05161112f620002f8600039600081816103e40152818161053d0152818161065e01528181610730015281816109a801528181610ae60152610de5015260008181610256015261034501526000818161015c015281816103160152818161049c0152610a590152600081816101a1015281816108e80152610d59015261112f6000f3fe60806040526004361061007b5760003560e01c806379ba50971161004e57806379ba5097146100df5780638da5cb5b146100f4578063f1d72e251461014a578063f2fde38b1461017e57600080fd5b80630efe6a8b1461008057806324600fc3146100a25780632b968958146100b757806347e7ef24146100cc575b600080fd5b34801561008c57600080fd5b506100a061009b366004610fa5565b61019e565b005b3480156100ae57600080fd5b506100a06105dc565b3480156100c357600080fd5b506100a06107f0565b6100a06100da366004610fd8565b6108e5565b3480156100eb57600080fd5b506100a0610b7f565b34801561010057600080fd5b506000546101219073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b34801561015657600080fd5b506101217f000000000000000000000000000000000000000000000000000000000000000081565b34801561018a57600080fd5b506100a0610199366004611002565b610c4c565b807f00000000000000000000000000000000000000000000000000000000000000008111156101f9576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102188473ffffffffffffffffffffffffffffffffffffffff16610d0e565b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018590529094507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906323b872dd906064016020604051808303816000875af11580156102b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102d89190611024565b506040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166004830152602482018590527f0000000000000000000000000000000000000000000000000000000000000000169063095ea7b3906044016020604051808303816000875af115801561038e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b29190611024565b506040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610440573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104649190611046565b6040517f47e7ef24000000000000000000000000000000000000000000000000000000008152306004820152602481018690529091507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906347e7ef2490604401600060405180830381600087803b1580156104f557600080fd5b505af1158015610509573d6000803e3d6000fd5b50506040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600092507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1691506370a0823190602401602060405180830381865afa15801561059a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105be9190611046565b90506105d4866105ce848461108e565b86610d55565b505050505050565b60005473ffffffffffffffffffffffffffffffffffffffff16331461062d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa1580156106ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106de9190611046565b6000546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152602481018390529192507f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af1158015610779573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079d9190611024565b5060005460405182815273ffffffffffffffffffffffffffffffffffffffff909116907f21901fa892c430ea8bd38b9390225ac8e67eac75ee10ffba16feefc539a288f99060200160405180910390a250565b60005473ffffffffffffffffffffffffffffffffffffffff163314610841576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60015473ffffffffffffffffffffffffffffffffffffffff1615610891576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600080547fffffffffffffffffffffffff000000000000000000000000000000000000000016815560405133907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3565b807f0000000000000000000000000000000000000000000000000000000000000000811115610940576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61095f8373ffffffffffffffffffffffffffffffffffffffff16610d0e565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290935060009073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa1580156109ef573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a139190611046565b6040517ff340fa0100000000000000000000000000000000000000000000000000000000815230600482015290915073ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063f340fa019034906024016000604051808303818588803b158015610a9e57600080fd5b505af1158015610ab2573d6000803e3d6000fd5b50506040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600093507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1692506370a082319150602401602060405180830381865afa158015610b44573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b689190611046565b9050610b78856105ce848461108e565b5050505050565b60015473ffffffffffffffffffffffffffffffffffffffff163314610bd0576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b60005473ffffffffffffffffffffffffffffffffffffffff163314610c9d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831690811790915560405133907f3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a90600090a350565b6000806000610d1c84610eab565b9150915081610d2d57509192915050565b600080610d3983610ef8565b9150915081610d4c575093949350505050565b95945050505050565b60007f0000000000000000000000000000000000000000000000000000000000000000610d8283856110a7565b610d8c91906110be565b610d96908461108e565b6040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018390529192507f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303816000875af1158015610e30573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e549190611024565b508373ffffffffffffffffffffffffffffffffffffffff167f0553260a2e46b0577270d8992db02d30856ca880144c72d6e9503760946aef1382604051610e9d91815260200190565b60405180910390a250505050565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103610ef2576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a8114610f585760009250600091505b50811580610f6757503d601614155b15610f7757506000928392509050565b915091565b803573ffffffffffffffffffffffffffffffffffffffff81168114610fa057600080fd5b919050565b600080600060608486031215610fba57600080fd5b610fc384610f7c565b95602085013595506040909401359392505050565b60008060408385031215610feb57600080fd5b610ff483610f7c565b946020939093013593505050565b60006020828403121561101457600080fd5b61101d82610f7c565b9392505050565b60006020828403121561103657600080fd5b8151801515811461101d57600080fd5b60006020828403121561105857600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156110a1576110a161105f565b92915050565b80820281158282048414176110a1576110a161105f565b6000826110f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b50049056fea2646970667358221220f9ed443cb277f8d4f7ac2074b9cc5fbab0c728e02b993cf27b078ebd29e9bad064736f6c63430008110033",
}

// PublicGoodsDonatorABI is the input ABI used to generate the binding from.
// Deprecated: Use PublicGoodsDonatorMetaData.ABI instead.
var PublicGoodsDonatorABI = PublicGoodsDonatorMetaData.ABI

// PublicGoodsDonatorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PublicGoodsDonatorMetaData.Bin instead.
var PublicGoodsDonatorBin = PublicGoodsDonatorMetaData.Bin

// DeployPublicGoodsDonator deploys a new Ethereum contract, binding an instance of PublicGoodsDonator to it.
func DeployPublicGoodsDonator(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _preStake common.Address, _wFIL common.Address, _liquidStakingToken common.Address) (common.Address, *types.Transaction, *PublicGoodsDonator, error) {
	parsed, err := PublicGoodsDonatorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PublicGoodsDonatorBin), backend, _owner, _preStake, _wFIL, _liquidStakingToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PublicGoodsDonator{PublicGoodsDonatorCaller: PublicGoodsDonatorCaller{contract: contract}, PublicGoodsDonatorTransactor: PublicGoodsDonatorTransactor{contract: contract}, PublicGoodsDonatorFilterer: PublicGoodsDonatorFilterer{contract: contract}}, nil
}

// PublicGoodsDonator is an auto generated Go binding around an Ethereum contract.
type PublicGoodsDonator struct {
	PublicGoodsDonatorCaller     // Read-only binding to the contract
	PublicGoodsDonatorTransactor // Write-only binding to the contract
	PublicGoodsDonatorFilterer   // Log filterer for contract events
}

// PublicGoodsDonatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PublicGoodsDonatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicGoodsDonatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PublicGoodsDonatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicGoodsDonatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PublicGoodsDonatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PublicGoodsDonatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PublicGoodsDonatorSession struct {
	Contract     *PublicGoodsDonator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PublicGoodsDonatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PublicGoodsDonatorCallerSession struct {
	Contract *PublicGoodsDonatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// PublicGoodsDonatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PublicGoodsDonatorTransactorSession struct {
	Contract     *PublicGoodsDonatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// PublicGoodsDonatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PublicGoodsDonatorRaw struct {
	Contract *PublicGoodsDonator // Generic contract binding to access the raw methods on
}

// PublicGoodsDonatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PublicGoodsDonatorCallerRaw struct {
	Contract *PublicGoodsDonatorCaller // Generic read-only contract binding to access the raw methods on
}

// PublicGoodsDonatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PublicGoodsDonatorTransactorRaw struct {
	Contract *PublicGoodsDonatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPublicGoodsDonator creates a new instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonator(address common.Address, backend bind.ContractBackend) (*PublicGoodsDonator, error) {
	contract, err := bindPublicGoodsDonator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonator{PublicGoodsDonatorCaller: PublicGoodsDonatorCaller{contract: contract}, PublicGoodsDonatorTransactor: PublicGoodsDonatorTransactor{contract: contract}, PublicGoodsDonatorFilterer: PublicGoodsDonatorFilterer{contract: contract}}, nil
}

// NewPublicGoodsDonatorCaller creates a new read-only instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonatorCaller(address common.Address, caller bind.ContractCaller) (*PublicGoodsDonatorCaller, error) {
	contract, err := bindPublicGoodsDonator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorCaller{contract: contract}, nil
}

// NewPublicGoodsDonatorTransactor creates a new write-only instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonatorTransactor(address common.Address, transactor bind.ContractTransactor) (*PublicGoodsDonatorTransactor, error) {
	contract, err := bindPublicGoodsDonator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorTransactor{contract: contract}, nil
}

// NewPublicGoodsDonatorFilterer creates a new log filterer instance of PublicGoodsDonator, bound to a specific deployed contract.
func NewPublicGoodsDonatorFilterer(address common.Address, filterer bind.ContractFilterer) (*PublicGoodsDonatorFilterer, error) {
	contract, err := bindPublicGoodsDonator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorFilterer{contract: contract}, nil
}

// bindPublicGoodsDonator binds a generic wrapper to an already deployed contract.
func bindPublicGoodsDonator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PublicGoodsDonatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicGoodsDonator *PublicGoodsDonatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicGoodsDonator.Contract.PublicGoodsDonatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicGoodsDonator *PublicGoodsDonatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.PublicGoodsDonatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicGoodsDonator *PublicGoodsDonatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.PublicGoodsDonatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PublicGoodsDonator *PublicGoodsDonatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PublicGoodsDonator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicGoodsDonator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorSession) Owner() (common.Address, error) {
	return _PublicGoodsDonator.Contract.Owner(&_PublicGoodsDonator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCallerSession) Owner() (common.Address, error) {
	return _PublicGoodsDonator.Contract.Owner(&_PublicGoodsDonator.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCaller) PreStake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PublicGoodsDonator.contract.Call(opts, &out, "preStake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorSession) PreStake() (common.Address, error) {
	return _PublicGoodsDonator.Contract.PreStake(&_PublicGoodsDonator.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_PublicGoodsDonator *PublicGoodsDonatorCallerSession) PreStake() (common.Address, error) {
	return _PublicGoodsDonator.Contract.PreStake(&_PublicGoodsDonator.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.AcceptOwnership(&_PublicGoodsDonator.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.AcceptOwnership(&_PublicGoodsDonator.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recipient, uint256 amount, uint256 donationPercent) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "deposit", recipient, amount, donationPercent)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recipient, uint256 amount, uint256 donationPercent) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) Deposit(recipient common.Address, amount *big.Int, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit(&_PublicGoodsDonator.TransactOpts, recipient, amount, donationPercent)
}

// Deposit is a paid mutator transaction binding the contract method 0x0efe6a8b.
//
// Solidity: function deposit(address recipient, uint256 amount, uint256 donationPercent) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) Deposit(recipient common.Address, amount *big.Int, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit(&_PublicGoodsDonator.TransactOpts, recipient, amount, donationPercent)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 donationPercent) payable returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) Deposit0(opts *bind.TransactOpts, recipient common.Address, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "deposit0", recipient, donationPercent)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 donationPercent) payable returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) Deposit0(recipient common.Address, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit0(&_PublicGoodsDonator.TransactOpts, recipient, donationPercent)
}

// Deposit0 is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 donationPercent) payable returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) Deposit0(recipient common.Address, donationPercent *big.Int) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.Deposit0(&_PublicGoodsDonator.TransactOpts, recipient, donationPercent)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) RevokeOwnership() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.RevokeOwnership(&_PublicGoodsDonator.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.RevokeOwnership(&_PublicGoodsDonator.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.TransferOwnership(&_PublicGoodsDonator.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.TransferOwnership(&_PublicGoodsDonator.TransactOpts, _newOwner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PublicGoodsDonator.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorSession) WithdrawFunds() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.WithdrawFunds(&_PublicGoodsDonator.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_PublicGoodsDonator *PublicGoodsDonatorTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _PublicGoodsDonator.Contract.WithdrawFunds(&_PublicGoodsDonator.TransactOpts)
}

// PublicGoodsDonatorDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorDonateIterator struct {
	Event *PublicGoodsDonatorDonate // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorDonate)
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
		it.Event = new(PublicGoodsDonatorDonate)
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
func (it *PublicGoodsDonatorDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorDonate represents a Donate event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorDonate struct {
	Account        common.Address
	DonationAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0x0553260a2e46b0577270d8992db02d30856ca880144c72d6e9503760946aef13.
//
// Solidity: event Donate(address indexed account, uint256 donationAmount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterDonate(opts *bind.FilterOpts, account []common.Address) (*PublicGoodsDonatorDonateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "Donate", accountRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorDonateIterator{contract: _PublicGoodsDonator.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0x0553260a2e46b0577270d8992db02d30856ca880144c72d6e9503760946aef13.
//
// Solidity: event Donate(address indexed account, uint256 donationAmount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorDonate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "Donate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorDonate)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "Donate", log); err != nil {
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

// ParseDonate is a log parse operation binding the contract event 0x0553260a2e46b0577270d8992db02d30856ca880144c72d6e9503760946aef13.
//
// Solidity: event Donate(address indexed account, uint256 donationAmount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseDonate(log types.Log) (*PublicGoodsDonatorDonate, error) {
	event := new(PublicGoodsDonatorDonate)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicGoodsDonatorOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipPendingIterator struct {
	Event *PublicGoodsDonatorOwnershipPending // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorOwnershipPending)
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
		it.Event = new(PublicGoodsDonatorOwnershipPending)
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
func (it *PublicGoodsDonatorOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorOwnershipPending represents a OwnershipPending event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipPending struct {
	CurrentOwner common.Address
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterOwnershipPending(opts *bind.FilterOpts, currentOwner []common.Address, pendingOwner []common.Address) (*PublicGoodsDonatorOwnershipPendingIterator, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorOwnershipPendingIterator{contract: _PublicGoodsDonator.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorOwnershipPending, currentOwner []common.Address, pendingOwner []common.Address) (event.Subscription, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorOwnershipPending)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseOwnershipPending(log types.Log) (*PublicGoodsDonatorOwnershipPending, error) {
	event := new(PublicGoodsDonatorOwnershipPending)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicGoodsDonatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipTransferredIterator struct {
	Event *PublicGoodsDonatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorOwnershipTransferred)
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
		it.Event = new(PublicGoodsDonatorOwnershipTransferred)
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
func (it *PublicGoodsDonatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorOwnershipTransferred represents a OwnershipTransferred event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PublicGoodsDonatorOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorOwnershipTransferredIterator{contract: _PublicGoodsDonator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorOwnershipTransferred)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseOwnershipTransferred(log types.Log) (*PublicGoodsDonatorOwnershipTransferred, error) {
	event := new(PublicGoodsDonatorOwnershipTransferred)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PublicGoodsDonatorWithdrawFundsIterator is returned from FilterWithdrawFunds and is used to iterate over the raw logs and unpacked data for WithdrawFunds events raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorWithdrawFundsIterator struct {
	Event *PublicGoodsDonatorWithdrawFunds // Event containing the contract specifics and raw log

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
func (it *PublicGoodsDonatorWithdrawFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PublicGoodsDonatorWithdrawFunds)
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
		it.Event = new(PublicGoodsDonatorWithdrawFunds)
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
func (it *PublicGoodsDonatorWithdrawFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PublicGoodsDonatorWithdrawFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PublicGoodsDonatorWithdrawFunds represents a WithdrawFunds event raised by the PublicGoodsDonator contract.
type PublicGoodsDonatorWithdrawFunds struct {
	Wallet common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawFunds is a free log retrieval operation binding the contract event 0x21901fa892c430ea8bd38b9390225ac8e67eac75ee10ffba16feefc539a288f9.
//
// Solidity: event WithdrawFunds(address indexed wallet, uint256 amount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) FilterWithdrawFunds(opts *bind.FilterOpts, wallet []common.Address) (*PublicGoodsDonatorWithdrawFundsIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.FilterLogs(opts, "WithdrawFunds", walletRule)
	if err != nil {
		return nil, err
	}
	return &PublicGoodsDonatorWithdrawFundsIterator{contract: _PublicGoodsDonator.contract, event: "WithdrawFunds", logs: logs, sub: sub}, nil
}

// WatchWithdrawFunds is a free log subscription operation binding the contract event 0x21901fa892c430ea8bd38b9390225ac8e67eac75ee10ffba16feefc539a288f9.
//
// Solidity: event WithdrawFunds(address indexed wallet, uint256 amount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) WatchWithdrawFunds(opts *bind.WatchOpts, sink chan<- *PublicGoodsDonatorWithdrawFunds, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _PublicGoodsDonator.contract.WatchLogs(opts, "WithdrawFunds", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PublicGoodsDonatorWithdrawFunds)
				if err := _PublicGoodsDonator.contract.UnpackLog(event, "WithdrawFunds", log); err != nil {
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

// ParseWithdrawFunds is a log parse operation binding the contract event 0x21901fa892c430ea8bd38b9390225ac8e67eac75ee10ffba16feefc539a288f9.
//
// Solidity: event WithdrawFunds(address indexed wallet, uint256 amount)
func (_PublicGoodsDonator *PublicGoodsDonatorFilterer) ParseWithdrawFunds(log types.Log) (*PublicGoodsDonatorWithdrawFunds, error) {
	event := new(PublicGoodsDonatorWithdrawFunds)
	if err := _PublicGoodsDonator.contract.UnpackLog(event, "WithdrawFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
