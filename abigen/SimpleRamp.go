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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_poolID\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPoolID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"burnIFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iFIL\",\"outputs\":[{\"internalType\":\"contractIPoolToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pool\",\"outputs\":[{\"internalType\":\"contractIPool\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"redeemF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshExtern\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalExitDemand\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wFIL\",\"outputs\":[{\"internalType\":\"contractIWFIL\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdrawF\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60c060405260006003553480156200001657600080fd5b5060405162002b3f38038062002b3f833981016040819052620000399162000376565b6001600160a01b03821660805260a0819052620000556200005d565b5050620003e1565b6200008d62000079608051620001b560201b62001aea1760201c565b60a0516200026260201b62001bcd1760201c565b600080546001600160a01b0319166001600160a01b0392909216918217905560408051630e5f46fb60e11b81529051631cbe8df6916004808201926020929091908290030181865afa158015620000e8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200010e9190620003a7565b600180546001600160a01b0319166001600160a01b03928316179055600054604080516338d52e0f60e01b8152905191909216916338d52e0f9160048083019260209291908290030181865afa1580156200016d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620001939190620003a7565b600280546001600160a01b0319166001600160a01b0392909216919091179055565b604080518082018252601481527f524f555445525f504f4f4c5f524547495354525900000000000000000000000060209091015251630d37324f60e11b8152631c86174160e11b60048201526000906001600160a01b03831690631a6e649e90602401602060405180830381865afa15801562000236573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200025c9190620003a7565b92915050565b6000826001600160a01b031663efde4e646040518163ffffffff1660e01b8152600401602060405180830381865afa158015620002a3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620002c99190620003c7565b821115620002ea576040516311ebac6360e31b815260040160405180910390fd5b6040516341d1de9760e01b8152600481018390526001600160a01b038416906341d1de9790602401602060405180830381865afa15801562000330573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620003569190620003a7565b9392505050565b6001600160a01b03811681146200037357600080fd5b50565b600080604083850312156200038a57600080fd5b825162000397816200035d565b6020939093015192949293505050565b600060208284031215620003ba57600080fd5b815162000356816200035d565b600060208284031215620003da57600080fd5b5051919050565b60805160a0516127006200043f6000396000610498015260008181610472015281816107920152818161094f01528181610cc201528181610fca0152818161117e01528181611332015281816115e501526121ff01526127006000f3fe6080604052600436106100f75760003560e01c80639adf10501161008a578063ce96cb7711610059578063ce96cb771461035e578063d905777e1461037e578063e8c9c4b71461039e578063fb932108146103b35761014f565b80639adf1050146102dc5780639f40a7b3146102f1578063a318c1a414610311578063bef44256146103315761014f565b806316f0115b116100c657806316f0115b146102285780633eb569c61461027a5780634cdad5061461028f5780636be82b78146102af5761014f565b80630a28a477146101a05780630c462c5f146101d35780630e0a7023146101f35780630fede599146102135761014f565b3661014f5760025473ffffffffffffffffffffffffffffffffffffffff16331461014d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b005b60025473ffffffffffffffffffffffffffffffffffffffff16331461014d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b3480156101ac57600080fd5b506101c06101bb3660046125a6565b6103d3565b6040519081526020015b60405180910390f35b3480156101df57600080fd5b506101c06101ee3660046125e1565b6106f3565b3480156101ff57600080fd5b506101c061020e3660046125e1565b6108b0565b34801561021f57600080fd5b506003546101c0565b34801561023457600080fd5b506000546102559073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101ca565b34801561028657600080fd5b5061014d610a64565b34801561029b57600080fd5b506101c06102aa3660046125a6565b610c23565b3480156102bb57600080fd5b506001546102559073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102e857600080fd5b5061014d610f21565b3480156102fd57600080fd5b506101c061030c3660046125e1565b610f2b565b34801561031d57600080fd5b506101c061032c3660046125e1565b6110df565b34801561033d57600080fd5b506002546102559073ffffffffffffffffffffffffffffffffffffffff1681565b34801561036a57600080fd5b506101c0610379366004612629565b611293565b34801561038a57600080fd5b506101c0610399366004612629565b611546565b3480156103aa57600080fd5b5061014d6118b7565b3480156103bf57600080fd5b5061014d6103ce366004612646565b6119ea565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa158015610441573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104659190612672565b15610510576104bc6104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b7f0000000000000000000000000000000000000000000000000000000000000000611bcd565b60005473ffffffffffffffffffffffffffffffffffffffff908116911614610510576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561057b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061059f9190612694565b8211156105ae57506000919050565b600154604080517f18160ddd000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916318160ddd9160048083019260209291908290030181865afa15801561061e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106429190612694565b905080156106ea576106e58160008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166301e1d1146040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106dd9190612694565b859190611d06565b6106ec565b825b9392505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa158015610761573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107859190612672565b1561080a576107b66104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b60005473ffffffffffffffffffffffffffffffffffffffff90811691161461080a576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8261082a8173ffffffffffffffffffffffffffffffffffffffff16611d34565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461088e576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610897866103d3565b91506108a7848684896001611d7b565b50949350505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa15801561091e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109429190612672565b156109c7576109736104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b60005473ffffffffffffffffffffffffffffffffffffffff9081169116146109c7576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826109e78173ffffffffffffffffffffffffffffffffffffffff16611d34565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a4b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610a5486610c23565b91506108a7848688856001611d7b565b600254604080517fd0e30db00000000000000000000000000000000000000000000000000000000081529051479273ffffffffffffffffffffffffffffffffffffffff169163d0e30db091849160048082019260009290919082900301818588803b158015610ad257600080fd5b505af1158015610ae6573d6000803e3d6000fd5b50506002546000546040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff928316955063a9059cbb94509116915083906370a0823190602401602060405180830381865afa158015610b67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b8b9190612694565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af1158015610bfb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1f9190612672565b5050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c91573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cb59190612672565b15610d3a57610ce66104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b60005473ffffffffffffffffffffffffffffffffffffffff908116911614610d3a576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600154604080517f18160ddd000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff16916318160ddd9160048083019260209291908290030181865afa158015610daa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dce9190612694565b90508015610e7857600054604080517f01e1d1140000000000000000000000000000000000000000000000000000000081529051610e739273ffffffffffffffffffffffffffffffffffffffff16916301e1d1149160048083019260209291908290030181865afa158015610e47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e6b9190612694565b849083611d06565b610e7a565b825b915060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ee7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0b9190612694565b821115610f1b5750600092915050565b50919050565b610f296121f7565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fbd9190612672565b1561104257610fee6104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b60005473ffffffffffffffffffffffffffffffffffffffff908116911614611042576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826110628173ffffffffffffffffffffffffffffffffffffffff16611d34565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146110c6576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110cf86610c23565b91506108a7848688856000611d7b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa15801561114d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111719190612672565b156111f6576111a26104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b60005473ffffffffffffffffffffffffffffffffffffffff9081169116146111f6576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826112168173ffffffffffffffffffffffffffffffffffffffff16611d34565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461127a576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611283866103d3565b91506108a7848684896000611d7b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa158015611301573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113259190612672565b156113aa576113566104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b60005473ffffffffffffffffffffffffffffffffffffffff9081169116146113aa576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000546001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152611540938116926307a2d13a929116906370a0823190602401602060405180830381865afa158015611429573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144d9190612694565b6040518263ffffffff1660e01b815260040161146b91815260200190565b602060405180830381865afa158015611488573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ac9190612694565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611517573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153b9190612694565b6123e6565b92915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663b2bcb0026040518163ffffffff1660e01b8152600401602060405180830381865afa1580156115b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d89190612672565b1561165d576116096104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b60005473ffffffffffffffffffffffffffffffffffffffff90811691161461165d576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8481166004830152909116906370a0823190602401602060405180830381865afa1580156116cd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f19190612694565b600080546040517f07a2d13a00000000000000000000000000000000000000000000000000000000815260048101849052929350909173ffffffffffffffffffffffffffffffffffffffff909116906307a2d13a90602401602060405180830381865afa158015611766573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061178a9190612694565b905060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117f7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061181b9190612694565b811115610f1b57600054604080517f5d66b00a00000000000000000000000000000000000000000000000000000000815290516106ec9273ffffffffffffffffffffffffffffffffffffffff1691635d66b00a9160048083019260209291908290030181865afa158015611893573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102aa9190612694565b6001546040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820181905273ffffffffffffffffffffffffffffffffffffffff90921691639dc29fac9183906370a0823190602401602060405180830381865afa15801561192f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119539190612694565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af11580156119c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119e79190612672565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314611a3b576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6002546000546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152306024820152604481018490529116906323b872dd906064016020604051808303816000875af1158015611abc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ae09190612672565b5050600060035550565b604080518082018252601481527f524f555445525f504f4f4c5f5245474953545259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f390c2e8200000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e90602401602060405180830381865afa158015611ba9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061154091906126ad565b60008273ffffffffffffffffffffffffffffffffffffffff1663efde4e646040518163ffffffff1660e01b8152600401602060405180830381865afa158015611c1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c3e9190612694565b821115611c77576040517f8f5d631800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f41d1de970000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff8416906341d1de9790602401602060405180830381865afa158015611ce2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ec91906126ad565b828202811515841585830485141716611d1e57600080fd5b6001826001830304018115150290509392505050565b6000806000611d42846123fc565b9150915081611d5357509192915050565b600080611d5f83612449565b9150915081611d72575093949350505050565b95945050505050565b611d9a8473ffffffffffffffffffffffffffffffffffffffff16611d34565b9350611dbb8573ffffffffffffffffffffffffffffffffffffffff16611d34565b945060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635d66b00a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611e28573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e4c9190612694565b821115611e85576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001546040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff878116600483015230602483015260448201869052909116906323b872dd906064016020604051808303816000875af1158015611f04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f289190612672565b506001546040517f9dc29fac0000000000000000000000000000000000000000000000000000000081523060048201526024810185905273ffffffffffffffffffffffffffffffffffffffff90911690639dc29fac906044016020604051808303816000875af1158015611fa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc49190612672565b50600382905560008054604080517f38dff49c000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff909216926338dff49c9260048084019382900301818387803b15801561203157600080fd5b505af1158015612045573d6000803e3d6000fd5b5050505080156120f9576002546040517f2e1a7d4d0000000000000000000000000000000000000000000000000000000081526004810184905273ffffffffffffffffffffffffffffffffffffffff90911690632e1a7d4d90602401600060405180830381600087803b1580156120bb57600080fd5b505af11580156120cf573d6000803e3d6000fd5b506120f49250505073ffffffffffffffffffffffffffffffffffffffff8516836124cd565b612198565b6002546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018590529091169063a9059cbb906044016020604051808303816000875af1158015612172573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121969190612672565b505b604080518381526020810185905273ffffffffffffffffffffffffffffffffffffffff808816929087169183917ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db910160405180910390a45050505050565b6122236104967f0000000000000000000000000000000000000000000000000000000000000000611aea565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169182179055604080517f1cbe8df60000000000000000000000000000000000000000000000000000000081529051631cbe8df6916004808201926020929091908290030181865afa1580156122bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122df91906126ad565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff928316179055600054604080517f38d52e0f000000000000000000000000000000000000000000000000000000008152905191909216916338d52e0f9160048083019260209291908290030181865afa15801561237b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061239f91906126ad565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60008183106123f557816106ec565b5090919050565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103612443576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a81146124a95760009250600091505b508115806124b857503d601614155b156124c857506000928392509050565b915091565b80471015612507576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114612561576040519150601f19603f3d011682016040523d82523d6000602084013e612566565b606091505b50509050806125a1576040517f3204506f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050565b6000602082840312156125b857600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff811681146119e757600080fd5b600080600080608085870312156125f757600080fd5b843593506020850135612609816125bf565b92506040850135612619816125bf565b9396929550929360600135925050565b60006020828403121561263b57600080fd5b81356106ec816125bf565b6000806040838503121561265957600080fd5b8235612664816125bf565b946020939093013593505050565b60006020828403121561268457600080fd5b815180151581146106ec57600080fd5b6000602082840312156126a657600080fd5b5051919050565b6000602082840312156126bf57600080fd5b81516106ec816125bf56fea264697066735822122009d80a59db35db31d3e6295d94fce16f2051df5ccde9c0b5f503ab0fe2ece06f64736f6c63430008110033",
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

// BurnIFIL is a paid mutator transaction binding the contract method 0xe8c9c4b7.
//
// Solidity: function burnIFIL() returns()
func (_SimpleRamp *SimpleRampTransactor) BurnIFIL(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "burnIFIL")
}

// BurnIFIL is a paid mutator transaction binding the contract method 0xe8c9c4b7.
//
// Solidity: function burnIFIL() returns()
func (_SimpleRamp *SimpleRampSession) BurnIFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.BurnIFIL(&_SimpleRamp.TransactOpts)
}

// BurnIFIL is a paid mutator transaction binding the contract method 0xe8c9c4b7.
//
// Solidity: function burnIFIL() returns()
func (_SimpleRamp *SimpleRampTransactorSession) BurnIFIL() (*types.Transaction, error) {
	return _SimpleRamp.Contract.BurnIFIL(&_SimpleRamp.TransactOpts)
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

// RedeemF is a paid mutator transaction binding the contract method 0x0e0a7023.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampTransactor) RedeemF(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "redeemF", shares, receiver, owner, arg3)
}

// RedeemF is a paid mutator transaction binding the contract method 0x0e0a7023.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampSession) RedeemF(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.RedeemF(&_SimpleRamp.TransactOpts, shares, receiver, owner, arg3)
}

// RedeemF is a paid mutator transaction binding the contract method 0x0e0a7023.
//
// Solidity: function redeemF(uint256 shares, address receiver, address owner, uint256 ) returns(uint256 assets)
func (_SimpleRamp *SimpleRampTransactorSession) RedeemF(shares *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.RedeemF(&_SimpleRamp.TransactOpts, shares, receiver, owner, arg3)
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

// WithdrawF is a paid mutator transaction binding the contract method 0x0c462c5f.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampTransactor) WithdrawF(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.contract.Transact(opts, "withdrawF", assets, receiver, owner, arg3)
}

// WithdrawF is a paid mutator transaction binding the contract method 0x0c462c5f.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampSession) WithdrawF(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.WithdrawF(&_SimpleRamp.TransactOpts, assets, receiver, owner, arg3)
}

// WithdrawF is a paid mutator transaction binding the contract method 0x0c462c5f.
//
// Solidity: function withdrawF(uint256 assets, address receiver, address owner, uint256 ) returns(uint256 shares)
func (_SimpleRamp *SimpleRampTransactorSession) WithdrawF(assets *big.Int, receiver common.Address, owner common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _SimpleRamp.Contract.WithdrawF(&_SimpleRamp.TransactOpts, assets, receiver, owner, arg3)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRamp *SimpleRampTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SimpleRamp.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRamp *SimpleRampSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Fallback(&_SimpleRamp.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SimpleRamp *SimpleRampTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SimpleRamp.Contract.Fallback(&_SimpleRamp.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRamp *SimpleRampTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SimpleRamp.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRamp *SimpleRampSession) Receive() (*types.Transaction, error) {
	return _SimpleRamp.Contract.Receive(&_SimpleRamp.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SimpleRamp *SimpleRampTransactorSession) Receive() (*types.Transaction, error) {
	return _SimpleRamp.Contract.Receive(&_SimpleRamp.TransactOpts)
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
