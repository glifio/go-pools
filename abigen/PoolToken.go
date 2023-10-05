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

// PoolTokenMetaData contains all meta data concerning the PoolToken contract.
var PoolTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"currentOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"OwnershipPending\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_burner\",\"type\":\"address\"}],\"name\":\"setBurner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"setMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200152a3803806200152a833981016040819052620000349162000253565b806040518060400160405280601881526020017f496e66696e69747920506f6f6c205374616b65642046494c0000000000000000815250604051806040016040528060048152602001631a51925360e21b815250601282600090816200009b91906200032a565b506001620000aa83826200032a565b506002805460ff191660ff9290921691909117905550506001600160a01b038116620000e95760405163e6c4247b60e01b815260040160405180910390fd5b62000108816001600160a01b03166200015a60201b62000c571760201c565b600680546001600160a01b0319166001600160a01b039290921691821790556040516000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35050620003f6565b600080806200016984620001a6565b91509150816200017b57509192915050565b6000806200018983620001d9565b91509150816200019d575093949350505050565b95945050505050565b600080600160401b600160a01b03831660ff60981b8103620001d357600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a81146200022d5760009250600091505b508115806200023d57503d601614155b156200024e57506000928392509050565b915091565b6000602082840312156200026657600080fd5b81516001600160a01b03811681146200027e57600080fd5b9392505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620002b057607f821691505b602082108103620002d157634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200032557600081815260208120601f850160051c81016020861015620003005750805b601f850160051c820191505b8181101562000321578281556001016200030c565b5050505b505050565b81516001600160401b0381111562000346576200034662000285565b6200035e816200035784546200029b565b84620002d7565b602080601f8311600181146200039657600084156200037d5750858301515b600019600386901b1c1916600185901b17855562000321565b600085815260208120601f198616915b82811015620003c757888601518255948401946001909101908401620003a6565b5085821015620003e65787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b61112480620004066000396000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c806370a08231116100cd578063a9059cbb11610081578063dd62ed3e11610066578063dd62ed3e146102de578063f2fde38b146102f1578063fca3b5aa1461030457600080fd5b8063a9059cbb146102b8578063a996d6ce146102cb57600080fd5b80638da5cb5b116100b25780638da5cb5b1461027d57806395d89b411461029d5780639dc29fac146102a557600080fd5b806370a082311461026257806379ba50971461027557600080fd5b806323b872dd116101245780632b968958116101095780632b96895814610226578063313ce5671461023057806340c10f191461024f57600080fd5b806323b872dd146101f357806327810b6e1461020657600080fd5b806306fdde03146101565780630754617214610174578063095ea7b3146101b957806318160ddd146101dc575b600080fd5b61015e610317565b60405161016b9190610ef6565b60405180910390f35b6008546101949073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161016b565b6101cc6101c7366004610f8b565b6103a5565b604051901515815260200161016b565b6101e560035481565b60405190815260200161016b565b6101cc610201366004610fb5565b61043b565b6009546101949073ffffffffffffffffffffffffffffffffffffffff1681565b61022e6105fc565b005b60025461023d9060ff1681565b60405160ff909116815260200161016b565b6101cc61025d366004610f8b565b6106f4565b6101e5610270366004610ff1565b61075b565b61022e6107ad565b6006546101949073ffffffffffffffffffffffffffffffffffffffff1681565b61015e61087e565b6101cc6102b3366004610f8b565b61088b565b6101cc6102c6366004610f8b565b6108e9565b61022e6102d9366004610ff1565b6109b7565b6101e56102ec366004611013565b610a4f565b61022e6102ff366004610ff1565b610afd565b61022e610312366004610ff1565b610bbf565b6000805461032490611046565b80601f016020809104026020016040519081016040528092919081815260200182805461035090611046565b801561039d5780601f106103725761010080835404028352916020019161039d565b820191906000526020600020905b81548152906001019060200180831161038057829003601f168201915b505050505081565b60006103c68373ffffffffffffffffffffffffffffffffffffffff16610c57565b33600081815260056020908152604080832073ffffffffffffffffffffffffffffffffffffffff8616808552908352928190208790555186815293965090927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a35060015b92915050565b600061045c8473ffffffffffffffffffffffffffffffffffffffff16610c57565b935061047d8373ffffffffffffffffffffffffffffffffffffffff16610c57565b73ffffffffffffffffffffffffffffffffffffffff851660009081526005602090815260408083203384529091529020549093507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610514576104e283826110c8565b73ffffffffffffffffffffffffffffffffffffffff861660009081526005602090815260408083203384529091529020555b73ffffffffffffffffffffffffffffffffffffffff8516600090815260046020526040812080548592906105499084906110c8565b909155505073ffffffffffffffffffffffffffffffffffffffff8416600090815260046020526040812080548592906105839084906110db565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516105e991815260200190565b60405180910390a3506001949350505050565b60065473ffffffffffffffffffffffffffffffffffffffff16331461064d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60075473ffffffffffffffffffffffffffffffffffffffff161561069d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600680547fffffffffffffffffffffffff000000000000000000000000000000000000000016905560405160009033907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3565b60085460009073ffffffffffffffffffffffffffffffffffffffff163314610748576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107528383610c9e565b50600192915050565b6000600460006107808473ffffffffffffffffffffffffffffffffffffffff16610c57565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040016000205492915050565b60075473ffffffffffffffffffffffffffffffffffffffff1633146107fe576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000008082163390811790935560078054909116905560405173ffffffffffffffffffffffffffffffffffffffff909116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a350565b6001805461032490611046565b60095460009073ffffffffffffffffffffffffffffffffffffffff1633146108df576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107528383610d63565b600061090a8373ffffffffffffffffffffffffffffffffffffffff16610c57565b3360009081526004602052604081208054929550849290919061092e9084906110c8565b909155505073ffffffffffffffffffffffffffffffffffffffff8316600090815260046020526040812080548492906109689084906110db565b909155505060405182815273ffffffffffffffffffffffffffffffffffffffff84169033907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610429565b60065473ffffffffffffffffffffffffffffffffffffffff163314610a08576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600980547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b600060056000610a748573ffffffffffffffffffffffffffffffffffffffff16610c57565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610acf8473ffffffffffffffffffffffffffffffffffffffff16610c57565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600020549392505050565b60065473ffffffffffffffffffffffffffffffffffffffff163314610b4e576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831690811790915560405133907f3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a90600090a350565b60065473ffffffffffffffffffffffffffffffffffffffff163314610c10576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000806000610c6584610e25565b9150915081610c7657509192915050565b600080610c8283610e72565b9150915081610c95575093949350505050565b95945050505050565b610cbd8273ffffffffffffffffffffffffffffffffffffffff16610c57565b91508060036000828254610cd191906110db565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526004602052604081208054839290610d0b9084906110db565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906020015b60405180910390a35050565b610d828273ffffffffffffffffffffffffffffffffffffffff16610c57565b73ffffffffffffffffffffffffffffffffffffffff8116600090815260046020526040812080549294508392909190610dbc9084906110c8565b925050819055508060036000828254610dd591906110c8565b909155505060405181815260009073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610d57565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103610e6c576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a8114610ed25760009250600091505b50811580610ee157503d601614155b15610ef157506000928392509050565b915091565b600060208083528351808285015260005b81811015610f2357858101830151858201604001528201610f07565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610f8657600080fd5b919050565b60008060408385031215610f9e57600080fd5b610fa783610f62565b946020939093013593505050565b600080600060608486031215610fca57600080fd5b610fd384610f62565b9250610fe160208501610f62565b9150604084013590509250925092565b60006020828403121561100357600080fd5b61100c82610f62565b9392505050565b6000806040838503121561102657600080fd5b61102f83610f62565b915061103d60208401610f62565b90509250929050565b600181811c9082168061105a57607f821691505b602082108103611093577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8181038181111561043557610435611099565b808201808211156104355761043561109956fea264697066735822122019a9d94764d591a976cbd3b72bf13410fa0be4d08cc85c46e9c1580634ec73fb64736f6c63430008110033",
}

// PoolTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolTokenMetaData.ABI instead.
var PoolTokenABI = PoolTokenMetaData.ABI

// PoolTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PoolTokenMetaData.Bin instead.
var PoolTokenBin = PoolTokenMetaData.Bin

// DeployPoolToken deploys a new Ethereum contract, binding an instance of PoolToken to it.
func DeployPoolToken(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *PoolToken, error) {
	parsed, err := PoolTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PoolTokenBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PoolToken{PoolTokenCaller: PoolTokenCaller{contract: contract}, PoolTokenTransactor: PoolTokenTransactor{contract: contract}, PoolTokenFilterer: PoolTokenFilterer{contract: contract}}, nil
}

// PoolToken is an auto generated Go binding around an Ethereum contract.
type PoolToken struct {
	PoolTokenCaller     // Read-only binding to the contract
	PoolTokenTransactor // Write-only binding to the contract
	PoolTokenFilterer   // Log filterer for contract events
}

// PoolTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolTokenSession struct {
	Contract     *PoolToken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolTokenCallerSession struct {
	Contract *PoolTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PoolTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTokenTransactorSession struct {
	Contract     *PoolTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PoolTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolTokenRaw struct {
	Contract *PoolToken // Generic contract binding to access the raw methods on
}

// PoolTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolTokenCallerRaw struct {
	Contract *PoolTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTokenTransactorRaw struct {
	Contract *PoolTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolToken creates a new instance of PoolToken, bound to a specific deployed contract.
func NewPoolToken(address common.Address, backend bind.ContractBackend) (*PoolToken, error) {
	contract, err := bindPoolToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolToken{PoolTokenCaller: PoolTokenCaller{contract: contract}, PoolTokenTransactor: PoolTokenTransactor{contract: contract}, PoolTokenFilterer: PoolTokenFilterer{contract: contract}}, nil
}

// NewPoolTokenCaller creates a new read-only instance of PoolToken, bound to a specific deployed contract.
func NewPoolTokenCaller(address common.Address, caller bind.ContractCaller) (*PoolTokenCaller, error) {
	contract, err := bindPoolToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTokenCaller{contract: contract}, nil
}

// NewPoolTokenTransactor creates a new write-only instance of PoolToken, bound to a specific deployed contract.
func NewPoolTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTokenTransactor, error) {
	contract, err := bindPoolToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTokenTransactor{contract: contract}, nil
}

// NewPoolTokenFilterer creates a new log filterer instance of PoolToken, bound to a specific deployed contract.
func NewPoolTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolTokenFilterer, error) {
	contract, err := bindPoolToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolTokenFilterer{contract: contract}, nil
}

// bindPoolToken binds a generic wrapper to an already deployed contract.
func bindPoolToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolToken *PoolTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolToken.Contract.PoolTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolToken *PoolTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolToken.Contract.PoolTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolToken *PoolTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolToken.Contract.PoolTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolToken *PoolTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolToken *PoolTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolToken *PoolTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_PoolToken *PoolTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_PoolToken *PoolTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PoolToken.Contract.Allowance(&_PoolToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_PoolToken *PoolTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PoolToken.Contract.Allowance(&_PoolToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_PoolToken *PoolTokenCaller) BalanceOf(opts *bind.CallOpts, _a common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "balanceOf", _a)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_PoolToken *PoolTokenSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _PoolToken.Contract.BalanceOf(&_PoolToken.CallOpts, _a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _a) view returns(uint256)
func (_PoolToken *PoolTokenCallerSession) BalanceOf(_a common.Address) (*big.Int, error) {
	return _PoolToken.Contract.BalanceOf(&_PoolToken.CallOpts, _a)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_PoolToken *PoolTokenCaller) Burner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "burner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_PoolToken *PoolTokenSession) Burner() (common.Address, error) {
	return _PoolToken.Contract.Burner(&_PoolToken.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_PoolToken *PoolTokenCallerSession) Burner() (common.Address, error) {
	return _PoolToken.Contract.Burner(&_PoolToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PoolToken *PoolTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PoolToken *PoolTokenSession) Decimals() (uint8, error) {
	return _PoolToken.Contract.Decimals(&_PoolToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PoolToken *PoolTokenCallerSession) Decimals() (uint8, error) {
	return _PoolToken.Contract.Decimals(&_PoolToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_PoolToken *PoolTokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_PoolToken *PoolTokenSession) Minter() (common.Address, error) {
	return _PoolToken.Contract.Minter(&_PoolToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_PoolToken *PoolTokenCallerSession) Minter() (common.Address, error) {
	return _PoolToken.Contract.Minter(&_PoolToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PoolToken *PoolTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PoolToken *PoolTokenSession) Name() (string, error) {
	return _PoolToken.Contract.Name(&_PoolToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PoolToken *PoolTokenCallerSession) Name() (string, error) {
	return _PoolToken.Contract.Name(&_PoolToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolToken *PoolTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolToken *PoolTokenSession) Owner() (common.Address, error) {
	return _PoolToken.Contract.Owner(&_PoolToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PoolToken *PoolTokenCallerSession) Owner() (common.Address, error) {
	return _PoolToken.Contract.Owner(&_PoolToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PoolToken *PoolTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PoolToken *PoolTokenSession) Symbol() (string, error) {
	return _PoolToken.Contract.Symbol(&_PoolToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PoolToken *PoolTokenCallerSession) Symbol() (string, error) {
	return _PoolToken.Contract.Symbol(&_PoolToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PoolToken *PoolTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PoolToken *PoolTokenSession) TotalSupply() (*big.Int, error) {
	return _PoolToken.Contract.TotalSupply(&_PoolToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PoolToken *PoolTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PoolToken.Contract.TotalSupply(&_PoolToken.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PoolToken *PoolTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PoolToken *PoolTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _PoolToken.Contract.AcceptOwnership(&_PoolToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PoolToken *PoolTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PoolToken.Contract.AcceptOwnership(&_PoolToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Approve(&_PoolToken.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Approve(&_PoolToken.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns(bool)
func (_PoolToken *PoolTokenTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns(bool)
func (_PoolToken *PoolTokenSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Burn(&_PoolToken.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns(bool)
func (_PoolToken *PoolTokenTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Burn(&_PoolToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_PoolToken *PoolTokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_PoolToken *PoolTokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Mint(&_PoolToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_PoolToken *PoolTokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Mint(&_PoolToken.TransactOpts, account, amount)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PoolToken *PoolTokenTransactor) RevokeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "revokeOwnership")
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PoolToken *PoolTokenSession) RevokeOwnership() (*types.Transaction, error) {
	return _PoolToken.Contract.RevokeOwnership(&_PoolToken.TransactOpts)
}

// RevokeOwnership is a paid mutator transaction binding the contract method 0x2b968958.
//
// Solidity: function revokeOwnership() returns()
func (_PoolToken *PoolTokenTransactorSession) RevokeOwnership() (*types.Transaction, error) {
	return _PoolToken.Contract.RevokeOwnership(&_PoolToken.TransactOpts)
}

// SetBurner is a paid mutator transaction binding the contract method 0xa996d6ce.
//
// Solidity: function setBurner(address _burner) returns()
func (_PoolToken *PoolTokenTransactor) SetBurner(opts *bind.TransactOpts, _burner common.Address) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "setBurner", _burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0xa996d6ce.
//
// Solidity: function setBurner(address _burner) returns()
func (_PoolToken *PoolTokenSession) SetBurner(_burner common.Address) (*types.Transaction, error) {
	return _PoolToken.Contract.SetBurner(&_PoolToken.TransactOpts, _burner)
}

// SetBurner is a paid mutator transaction binding the contract method 0xa996d6ce.
//
// Solidity: function setBurner(address _burner) returns()
func (_PoolToken *PoolTokenTransactorSession) SetBurner(_burner common.Address) (*types.Transaction, error) {
	return _PoolToken.Contract.SetBurner(&_PoolToken.TransactOpts, _burner)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_PoolToken *PoolTokenTransactor) SetMinter(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "setMinter", _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_PoolToken *PoolTokenSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _PoolToken.Contract.SetMinter(&_PoolToken.TransactOpts, _minter)
}

// SetMinter is a paid mutator transaction binding the contract method 0xfca3b5aa.
//
// Solidity: function setMinter(address _minter) returns()
func (_PoolToken *PoolTokenTransactorSession) SetMinter(_minter common.Address) (*types.Transaction, error) {
	return _PoolToken.Contract.SetMinter(&_PoolToken.TransactOpts, _minter)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Transfer(&_PoolToken.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.Transfer(&_PoolToken.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenTransactor) TransferFrom(opts *bind.TransactOpts, _owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "transferFrom", _owner, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenSession) TransferFrom(_owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.TransferFrom(&_PoolToken.TransactOpts, _owner, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _owner, address _to, uint256 _amount) returns(bool)
func (_PoolToken *PoolTokenTransactorSession) TransferFrom(_owner common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PoolToken.Contract.TransferFrom(&_PoolToken.TransactOpts, _owner, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PoolToken *PoolTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _PoolToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PoolToken *PoolTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PoolToken.Contract.TransferOwnership(&_PoolToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_PoolToken *PoolTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _PoolToken.Contract.TransferOwnership(&_PoolToken.TransactOpts, _newOwner)
}

// PoolTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PoolToken contract.
type PoolTokenApprovalIterator struct {
	Event *PoolTokenApproval // Event containing the contract specifics and raw log

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
func (it *PoolTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokenApproval)
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
		it.Event = new(PoolTokenApproval)
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
func (it *PoolTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokenApproval represents a Approval event raised by the PoolToken contract.
type PoolTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_PoolToken *PoolTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoolTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PoolToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokenApprovalIterator{contract: _PoolToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_PoolToken *PoolTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PoolToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokenApproval)
				if err := _PoolToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_PoolToken *PoolTokenFilterer) ParseApproval(log types.Log) (*PoolTokenApproval, error) {
	event := new(PoolTokenApproval)
	if err := _PoolToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTokenOwnershipPendingIterator is returned from FilterOwnershipPending and is used to iterate over the raw logs and unpacked data for OwnershipPending events raised by the PoolToken contract.
type PoolTokenOwnershipPendingIterator struct {
	Event *PoolTokenOwnershipPending // Event containing the contract specifics and raw log

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
func (it *PoolTokenOwnershipPendingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokenOwnershipPending)
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
		it.Event = new(PoolTokenOwnershipPending)
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
func (it *PoolTokenOwnershipPendingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokenOwnershipPendingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokenOwnershipPending represents a OwnershipPending event raised by the PoolToken contract.
type PoolTokenOwnershipPending struct {
	CurrentOwner common.Address
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipPending is a free log retrieval operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_PoolToken *PoolTokenFilterer) FilterOwnershipPending(opts *bind.FilterOpts, currentOwner []common.Address, pendingOwner []common.Address) (*PoolTokenOwnershipPendingIterator, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PoolToken.contract.FilterLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokenOwnershipPendingIterator{contract: _PoolToken.contract, event: "OwnershipPending", logs: logs, sub: sub}, nil
}

// WatchOwnershipPending is a free log subscription operation binding the contract event 0x3c672e6c16e239c29e969deaf4eae15d7002350bf5490175e7481155af04c83a.
//
// Solidity: event OwnershipPending(address indexed currentOwner, address indexed pendingOwner)
func (_PoolToken *PoolTokenFilterer) WatchOwnershipPending(opts *bind.WatchOpts, sink chan<- *PoolTokenOwnershipPending, currentOwner []common.Address, pendingOwner []common.Address) (event.Subscription, error) {

	var currentOwnerRule []interface{}
	for _, currentOwnerItem := range currentOwner {
		currentOwnerRule = append(currentOwnerRule, currentOwnerItem)
	}
	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _PoolToken.contract.WatchLogs(opts, "OwnershipPending", currentOwnerRule, pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokenOwnershipPending)
				if err := _PoolToken.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
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
func (_PoolToken *PoolTokenFilterer) ParseOwnershipPending(log types.Log) (*PoolTokenOwnershipPending, error) {
	event := new(PoolTokenOwnershipPending)
	if err := _PoolToken.contract.UnpackLog(event, "OwnershipPending", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PoolToken contract.
type PoolTokenOwnershipTransferredIterator struct {
	Event *PoolTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokenOwnershipTransferred)
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
		it.Event = new(PoolTokenOwnershipTransferred)
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
func (it *PoolTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokenOwnershipTransferred represents a OwnershipTransferred event raised by the PoolToken contract.
type PoolTokenOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PoolToken *PoolTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*PoolTokenOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoolToken.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokenOwnershipTransferredIterator{contract: _PoolToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_PoolToken *PoolTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolTokenOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PoolToken.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokenOwnershipTransferred)
				if err := _PoolToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PoolToken *PoolTokenFilterer) ParseOwnershipTransferred(log types.Log) (*PoolTokenOwnershipTransferred, error) {
	event := new(PoolTokenOwnershipTransferred)
	if err := _PoolToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PoolToken contract.
type PoolTokenTransferIterator struct {
	Event *PoolTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PoolTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTokenTransfer)
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
		it.Event = new(PoolTokenTransfer)
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
func (it *PoolTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTokenTransfer represents a Transfer event raised by the PoolToken contract.
type PoolTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_PoolToken *PoolTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoolTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoolTokenTransferIterator{contract: _PoolToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_PoolToken *PoolTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoolToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTokenTransfer)
				if err := _PoolToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_PoolToken *PoolTokenFilterer) ParseTransfer(log types.Log) (*PoolTokenTransfer, error) {
	event := new(PoolTokenTransfer)
	if err := _PoolToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
