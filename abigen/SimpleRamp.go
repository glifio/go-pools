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

// SimpleRampMetaData contains all meta data concerning the SimpleRamp contract.
var SimpleRampMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_poolID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iFIL\",\"outputs\":[{\"internalType\":\"contractIPoolToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshExtern\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalExitDemand\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wFIL\",\"outputs\":[{\"internalType\":\"contractIWFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60c060405260006003553480156200001657600080fd5b5060405162001c4f38038062001c4f833981016040819052620000399162000376565b6001600160a01b03821660805260a0819052620000556200005d565b5050620003e1565b6200008d62000079608051620001b560201b62000dab1760201c565b60a0516200026260201b62000e8e1760201c565b600080546001600160a01b0319166001600160a01b0392909216918217905560408051630e5f46fb60e11b81529051631cbe8df6916004808201926020929091908290030181865afa158015620000e8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010e9190620003a7565b600180546001600160a01b0319166001600160a01b03928316179055600054604080516338d52e0f60e01b8152905191909216916338d52e0f9160048083019260209291908290030181865afa1580156200016d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001939190620003a7565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082018252601481527f524f555445525f504f4f4c5f524547495354525900000000000000000000000060209091015251630d37324f60e11b8152631c86174160e11b60048201526000906001600160a01b03831690631a6e649e90602401602060405180830381865afa15801562000236573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200025c9190620003a7565b92915050565b6000826001600160a01b031663efde4e646040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002a3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002c99190620003c7565b821115620002ea576040516311ebac6360e31b815260040160405180910390fd5b6040516341d1de9760e01b8152600481018390526001600160a01b038416906341d1de9790602401602060405180830381865afa15801562000330573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003569190620003a7565b9392505050565b6001600160a01b03811681146200037357600080fd5b50565b600080604083850312156200038a57600080fd5b825162000397816200035d565b6020939093015192949293505050565b600060208284031215620003ba57600080fd5b815162000356816200035d565b600060208284031215620003da57600080fd5b5051919050565b60805160a051611848620004076000396000610ff501526000610fcf01526118486000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80639adf10501161008c578063bef4425611610066578063bef44256146101c2578063ce96cb77146101e2578063d905777e146101f5578063fb9321081461020857600080fd5b80639adf1050146101945780639f40a7b31461019c578063a318c1a4146101af57600080fd5b80633eb569c6116100bd5780633eb569c6146101575780634cdad506146101615780636be82b781461017457600080fd5b80630a28a477146100e45780630fede5991461010a57806316f0115b14610112575b600080fd5b6100f76100f23660046116b1565b61021b565b6040519081526020015b60405180910390f35b6003546100f7565b6000546101329073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610101565b61015f610355565b005b6100f761016f3660046116b1565b610514565b6001546101329073ffffffffffffffffffffffffffffffffffffffff1681565b61015f61064d565b6100f76101aa3660046116ef565b610657565b6100f76101bd3660046116ef565b610753565b6002546101329073ffffffffffffffffffffffffffffffffffffffff1681565b6100f76101f0366004611737565b610846565b6100f7610203366004611737565b6109dd565b61015f610216366004611754565b610cab565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610289573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102ad9190611780565b8211156102bc57506000919050565b6000546040517fc6e6f5920000000000000000000000000000000000000000000000000000000081526004810184905273ffffffffffffffffffffffffffffffffffffffff9091169063c6e6f59290602401602060405180830381865afa15801561032b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034f9190611780565b92915050565b600254604080517fd0e30db00000000000000000000000000000000000000000000000000000000081529051479273ffffffffffffffffffffffffffffffffffffffff169163d0e30db091849160048082019260009290919082900301818588803b1580156103c357600080fd5b505af11580156103d7573d6000803e3d6000fd5b50506002546000546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff928316955063a9059cbb94509116915083906370a0823190602401602060405180830381865afa158015610458573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061047c9190611780565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156104ec573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105109190611799565b5050565b600080546040517f07a2d13a0000000000000000000000000000000000000000000000000000000081526004810184905273ffffffffffffffffffffffffffffffffffffffff909116906307a2d13a90602401602060405180830381865afa158015610584573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105a89190611780565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610615573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106399190611780565b81111561064857506000919050565b919050565b610655610fc7565b565b6000823373ffffffffffffffffffffffffffffffffffffffff8216146106a9576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517f07a2d13a0000000000000000000000000000000000000000000000000000000081526004810188905273ffffffffffffffffffffffffffffffffffffffff909116906307a2d13a90602401602060405180830381865afa158015610718573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073c9190611780565b915061074a848688856111dc565b50949350505050565b6000823373ffffffffffffffffffffffffffffffffffffffff8216146107a5576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546040517fc6e6f5920000000000000000000000000000000000000000000000000000000081526004810188905273ffffffffffffffffffffffffffffffffffffffff9091169063c6e6f59290602401602060405180830381865afa158015610814573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108389190611780565b915061074a848684896111dc565b600080546001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015261034f938116926307a2d13a929116906370a0823190602401602060405180830381865afa1580156108c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108ea9190611780565b6040518263ffffffff1660e01b815260040161090891815260200190565b602060405180830381865afa158015610925573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109499190611780565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109d89190611780565b61169b565b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff838116600483015260009216906370a0823190602401602060405180830381865afa158015610a4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a729190611780565b600080546040517f07a2d13a00000000000000000000000000000000000000000000000000000000815260048101849052929350909173ffffffffffffffffffffffffffffffffffffffff909116906307a2d13a90602401602060405180830381865afa158015610ae7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b0b9190611780565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b9c9190611780565b811115610ca557600054604080517f5d66b00a000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff9092169163c6e6f592918391635d66b00a916004808201926020929091908290030181865afa158015610c1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3f9190611780565b6040518263ffffffff1660e01b8152600401610c5d91815260200190565b602060405180830381865afa158015610c7a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9e9190611780565b9392505050565b50919050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610cfc576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152306024820152604481018490529116906323b872dd906064016020604051808303816000875af1158015610d7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610da19190611799565b5050600060035550565b604080518082018252601481527f524f555445525f504f4f4c5f5245474953545259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f390c2e8200000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e90602401602060405180830381865afa158015610e6a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034f91906117bb565b60008273ffffffffffffffffffffffffffffffffffffffff1663efde4e646040518163ffffffff1660e01b8152600401602060405180830381865afa158015610edb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eff9190611780565b821115610f38576040517f8f5d631800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f41d1de970000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8416906341d1de9790602401602060405180830381865afa158015610fa3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c9e91906117bb565b611019610ff37f0000000000000000000000000000000000000000000000000000000000000000610dab565b7f0000000000000000000000000000000000000000000000000000000000000000610e8e565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169182179055604080517f1cbe8df60000000000000000000000000000000000000000000000000000000081529051631cbe8df6916004808201926020929091908290030181865afa1580156110b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d591906117bb565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff928316179055600054604080517f38d52e0f000000000000000000000000000000000000000000000000000000008152905191909216916338d52e0f9160048083019260209291908290030181865afa158015611171573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061119591906117bb565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611247573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061126b9190611780565b8111156112a4576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015260009173ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015611313573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113379190611780565b6001546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8881166004830152306024830152604482018790529293509116906323b872dd906064016020604051808303816000875af11580156113b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113dc9190611799565b506001546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820181905273ffffffffffffffffffffffffffffffffffffffff90921691639dc29fac91849084906370a0823190602401602060405180830381865afa158015611457573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147b9190611780565b61148591906117d8565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156114f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115199190611799565b50600382905560008054604080517f38dff49c000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff909216926338dff49c9260048084019382900301818387803b15801561158657600080fd5b505af115801561159a573d6000803e3d6000fd5b50506002546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff888116600483015260248201879052909116925063a9059cbb91506044016020604051808303816000875af1158015611617573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163b9190611799565b50604080518381526020810185905273ffffffffffffffffffffffffffffffffffffffff808816929087169183917ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db910160405180910390a45050505050565b60008183106116aa5781610c9e565b5090919050565b6000602082840312156116c357600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff811681146116ec57600080fd5b50565b6000806000806080858703121561170557600080fd5b843593506020850135611717816116ca565b92506040850135611727816116ca565b9396929550929360600135925050565b60006020828403121561174957600080fd5b8135610c9e816116ca565b6000806040838503121561176757600080fd5b8235611772816116ca565b946020939093013593505050565b60006020828403121561179257600080fd5b5051919050565b6000602082840312156117ab57600080fd5b81518015158114610c9e57600080fd5b6000602082840312156117cd57600080fd5b8151610c9e816116ca565b8181038181111561034f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220780872456e7672cc4340642ed305800eee419e27d6aee965deae5fcb21135ad664736f6c63430008110033",
}

// SimpleRampABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpleRampMetaData.ABI instead.
var SimpleRampABI = SimpleRampMetaData.ABI

// SimpleRampBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpleRampMetaData.Bin instead.
var SimpleRampBin = SimpleRampMetaData.Bin

// DeploySimpleRamp deploys a new Ethereum contract, binding an instance of SimpleRamp to it.
func DeploySimpleRamp(auth *bind.TransactOpts, backend bind.ContractBackend, _router common.Address, _poolID *big.Int) (common.Address, *types.Transaction, *SimpleRamp, error) {
	parsed, err := SimpleRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpleRampBin), backend, _router, _poolID)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SimpleRamp{SimpleRampCaller: SimpleRampCaller{contract: contract}, SimpleRampTransactor: SimpleRampTransactor{contract: contract}, SimpleRampFilterer: SimpleRampFilterer{contract: contract}}, nil
}

// SimpleRamp is an auto generated Go binding around an Ethereum contract.
type SimpleRamp struct {
	SimpleRampCaller     // Read-only binding to the contract
	SimpleRampTransactor // Write-only binding to the contract
	SimpleRampFilterer   // Log filterer for contract events
}

// SimpleRampCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpleRampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpleRampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpleRampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpleRampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpleRampSession struct {
	Contract     *SimpleRamp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpleRampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpleRampCallerSession struct {
	Contract *SimpleRampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SimpleRampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpleRampTransactorSession struct {
	Contract     *SimpleRampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SimpleRampRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpleRampRaw struct {
	Contract *SimpleRamp // Generic contract binding to access the raw methods on
}

// SimpleRampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpleRampCallerRaw struct {
	Contract *SimpleRampCaller // Generic read-only contract binding to access the raw methods on
}

// SimpleRampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpleRampTransactorRaw struct {
	Contract *SimpleRampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpleRamp creates a new instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRamp(address common.Address, backend bind.ContractBackend) (*SimpleRamp, error) {
	contract, err := bindSimpleRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SimpleRamp{SimpleRampCaller: SimpleRampCaller{contract: contract}, SimpleRampTransactor: SimpleRampTransactor{contract: contract}, SimpleRampFilterer: SimpleRampFilterer{contract: contract}}, nil
}

// NewSimpleRampCaller creates a new read-only instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRampCaller(address common.Address, caller bind.ContractCaller) (*SimpleRampCaller, error) {
	contract, err := bindSimpleRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRampCaller{contract: contract}, nil
}

// NewSimpleRampTransactor creates a new write-only instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRampTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpleRampTransactor, error) {
	contract, err := bindSimpleRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpleRampTransactor{contract: contract}, nil
}

// NewSimpleRampFilterer creates a new log filterer instance of SimpleRamp, bound to a specific deployed contract.
func NewSimpleRampFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpleRampFilterer, error) {
	contract, err := bindSimpleRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpleRampFilterer{contract: contract}, nil
}

// bindSimpleRamp binds a generic wrapper to an already deployed contract.
func bindSimpleRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpleRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRamp *SimpleRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRamp.Contract.SimpleRampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRamp *SimpleRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.Contract.SimpleRampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRamp *SimpleRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRamp.Contract.SimpleRampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SimpleRamp *SimpleRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SimpleRamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SimpleRamp *SimpleRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SimpleRamp *SimpleRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SimpleRamp.Contract.contract.Transact(opts, method, params...)
}

// IFIL is a free data retrieval call binding the contract method 0x6be82b78.
//
// Solidity: function iFIL() view returns(address)
func (_SimpleRamp *SimpleRampCaller) IFIL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "iFIL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IFIL is a free data retrieval call binding the contract method 0x6be82b78.
//
// Solidity: function iFIL() view returns(address)
func (_SimpleRamp *SimpleRampSession) IFIL() (common.Address, error) {
	return _SimpleRamp.Contract.IFIL(&_SimpleRamp.CallOpts)
}

// IFIL is a free data retrieval call binding the contract method 0x6be82b78.
//
// Solidity: function iFIL() view returns(address)
func (_SimpleRamp *SimpleRampCallerSession) IFIL() (common.Address, error) {
	return _SimpleRamp.Contract.IFIL(&_SimpleRamp.CallOpts)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address account) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCaller) MaxRedeem(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "maxRedeem", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address account) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampSession) MaxRedeem(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxRedeem(&_SimpleRamp.CallOpts, account)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address account) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCallerSession) MaxRedeem(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxRedeem(&_SimpleRamp.CallOpts, account)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address account) view returns(uint256)
func (_SimpleRamp *SimpleRampCaller) MaxWithdraw(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "maxWithdraw", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address account) view returns(uint256)
func (_SimpleRamp *SimpleRampSession) MaxWithdraw(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxWithdraw(&_SimpleRamp.CallOpts, account)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address account) view returns(uint256)
func (_SimpleRamp *SimpleRampCallerSession) MaxWithdraw(account common.Address) (*big.Int, error) {
	return _SimpleRamp.Contract.MaxWithdraw(&_SimpleRamp.CallOpts, account)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SimpleRamp *SimpleRampCaller) Pool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "pool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SimpleRamp *SimpleRampSession) Pool() (common.Address, error) {
	return _SimpleRamp.Contract.Pool(&_SimpleRamp.CallOpts)
}

// Pool is a free data retrieval call binding the contract method 0x16f0115b.
//
// Solidity: function pool() view returns(address)
func (_SimpleRamp *SimpleRampCallerSession) Pool() (common.Address, error) {
	return _SimpleRamp.Contract.Pool(&_SimpleRamp.CallOpts)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_SimpleRamp *SimpleRampCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_SimpleRamp *SimpleRampSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewRedeem(&_SimpleRamp.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256 assets)
func (_SimpleRamp *SimpleRampCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewRedeem(&_SimpleRamp.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewWithdraw(&_SimpleRamp.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256 shares)
func (_SimpleRamp *SimpleRampCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _SimpleRamp.Contract.PreviewWithdraw(&_SimpleRamp.CallOpts, assets)
}

// TotalExitDemand is a free data retrieval call binding the contract method 0x0fede599.
//
// Solidity: function totalExitDemand() view returns(uint256)
func (_SimpleRamp *SimpleRampCaller) TotalExitDemand(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "totalExitDemand")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalExitDemand is a free data retrieval call binding the contract method 0x0fede599.
//
// Solidity: function totalExitDemand() view returns(uint256)
func (_SimpleRamp *SimpleRampSession) TotalExitDemand() (*big.Int, error) {
	return _SimpleRamp.Contract.TotalExitDemand(&_SimpleRamp.CallOpts)
}

// TotalExitDemand is a free data retrieval call binding the contract method 0x0fede599.
//
// Solidity: function totalExitDemand() view returns(uint256)
func (_SimpleRamp *SimpleRampCallerSession) TotalExitDemand() (*big.Int, error) {
	return _SimpleRamp.Contract.TotalExitDemand(&_SimpleRamp.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0xbef44256.
//
// Solidity: function wFIL() view returns(address)
func (_SimpleRamp *SimpleRampCaller) WFIL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SimpleRamp.contract.Call(opts, &out, "wFIL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WFIL is a free data retrieval call binding the contract method 0xbef44256.
//
// Solidity: function wFIL() view returns(address)
func (_SimpleRamp *SimpleRampSession) WFIL() (common.Address, error) {
	return _SimpleRamp.Contract.WFIL(&_SimpleRamp.CallOpts)
}

// WFIL is a free data retrieval call binding the contract method 0xbef44256.
//
// Solidity: function wFIL() view returns(address)
func (_SimpleRamp *SimpleRampCallerSession) WFIL() (common.Address, error) {
	return _SimpleRamp.Contract.WFIL(&_SimpleRamp.CallOpts)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address , uint256 amount) returns()
func (_SimpleRamp *SimpleRampTransactor) Distribute(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "distribute", arg0, amount)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address , uint256 amount) returns()
func (_SimpleRamp *SimpleRampSession) Distribute(arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Distribute(&_SimpleRamp.TransactOpts, arg0, amount)
}

// Distribute is a paid mutator transaction binding the contract method 0xfb932108.
//
// Solidity: function distribute(address , uint256 amount) returns()
func (_SimpleRamp *SimpleRampTransactorSession) Distribute(arg0 common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Distribute(&_SimpleRamp.TransactOpts, arg0, amount)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x3eb569c6.
//
// Solidity: function recoverFIL() returns()
func (_SimpleRamp *SimpleRampTransactor) RecoverFIL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "recoverFIL")
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x3eb569c6.
//
// Solidity: function recoverFIL() returns()
func (_SimpleRamp *SimpleRampSession) RecoverFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RecoverFIL(&_SimpleRamp.TransactOpts)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x3eb569c6.
//
// Solidity: function recoverFIL() returns()
func (_SimpleRamp *SimpleRampTransactorSession) RecoverFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RecoverFIL(&_SimpleRamp.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "redeem", shares, receiver, owner, arg3)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Redeem(&_SimpleRamp.TransactOpts, shares, receiver, owner, arg3)
}

// Redeem is a paid mutator transaction binding the contract method 0x9f40a7b3.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Redeem(&_SimpleRamp.TransactOpts, shares, receiver, owner, arg3)
}

// RefreshExtern is a paid mutator transaction binding the contract method 0x9adf1050.
//
// Solidity: function refreshExtern() returns()
func (_SimpleRamp *SimpleRampTransactor) RefreshExtern(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "refreshExtern")
}

// RefreshExtern is a paid mutator transaction binding the contract method 0x9adf1050.
//
// Solidity: function refreshExtern() returns()
func (_SimpleRamp *SimpleRampSession) RefreshExtern() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RefreshExtern(&_SimpleRamp.TransactOpts)
}

// RefreshExtern is a paid mutator transaction binding the contract method 0x9adf1050.
//
// Solidity: function refreshExtern() returns()
func (_SimpleRamp *SimpleRampTransactorSession) RefreshExtern() (*types.Transaction, error) {
	return _SimpleRamp.Contract.RefreshExtern(&_SimpleRamp.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "withdraw", assets, receiver, owner, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Withdraw(&_SimpleRamp.TransactOpts, assets, receiver, owner, arg3)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa318c1a4.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Withdraw(&_SimpleRamp.TransactOpts, assets, receiver, owner, arg3)
}

// SimpleRampWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SimpleRamp contract.
type SimpleRampWithdrawIterator struct {
	Event *SimpleRampWithdraw // Event containing the contract specifics and raw log

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
func (it *SimpleRampWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpleRampWithdraw)
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
		it.Event = new(SimpleRampWithdraw)
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
func (it *SimpleRampWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpleRampWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpleRampWithdraw represents a Withdraw event raised by the SimpleRamp contract.
type SimpleRampWithdraw struct {
	Caller   common.Address
	Receiver common.Address
	Owner    common.Address
	Assets   *big.Int
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_SimpleRamp *SimpleRampFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*SimpleRampWithdrawIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleRamp.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &SimpleRampWithdrawIterator{contract: _SimpleRamp.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_SimpleRamp *SimpleRampFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SimpleRampWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _SimpleRamp.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpleRampWithdraw)
				if err := _SimpleRamp.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_SimpleRamp *SimpleRampFilterer) ParseWithdraw(log types.Log) (*SimpleRampWithdraw, error) {
	event := new(SimpleRampWithdraw)
	if err := _SimpleRamp.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
