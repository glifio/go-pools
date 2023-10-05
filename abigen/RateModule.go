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
type Account_duplicate2 struct {
	StartEpoch *big.Int
	Principal  *big.Int
	EpochsPaid *big.Int
	Defaulted  bool
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_duplicate2 struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// RateModuleMetaData contains all meta data concerning the RateModule contract.
var RateModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256[61]\",\"name\":\"_rateLookup\",\"type\":\"uint256[61]\"},{\"internalType\":\"uint256[10]\",\"name\":\"_levels\",\"type\":\"uint256[10]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountLevel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"agentTotalValue\",\"type\":\"uint256\"}],\"name\":\"computeDTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedDailyRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountPrincipal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrincipal\",\"type\":\"uint256\"}],\"name\":\"computeDTI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"}],\"name\":\"computeLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credParser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"defaulted\",\"type\":\"bool\"}],\"internalType\":\"structAccount\",\"name\":\"account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDTE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDTI\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxLTV\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minGCRED\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"penaltyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rateLookup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"agentIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"}],\"name\":\"setAgentLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseRate\",\"type\":\"uint256\"}],\"name\":\"setBaseRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[10]\",\"name\":\"_levels\",\"type\":\"uint256[10]\"}],\"name\":\"setLevels\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDTE\",\"type\":\"uint256\"}],\"name\":\"setMaxDTE\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDTI\",\"type\":\"uint256\"}],\"name\":\"setMaxDTI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxLTV\",\"type\":\"uint256\"}],\"name\":\"setMaxLTV\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minGCRED\",\"type\":\"uint256\"}],\"name\":\"setMinGCRED\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[61]\",\"name\":\"_rateLookup\",\"type\":\"uint256[61]\"}],\"name\":\"setRateLookup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateCredParser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040526703782dace9d90000600255670de0b6b3a7640000600355670de0b6b3a7640000600455602860055567027f7d0bdb920000604d553480156200004657600080fd5b5060405162001cb338038062001cb3833981016040819052620000699162000479565b8362000089816001600160a01b03166200013b60201b62000a3e1760201c565b90506001600160a01b038116620000b357604051635435b28960e11b815260040160405180910390fd5b620000be8162000187565b50826001600160a01b03166080816001600160a01b031681525050620000f1608051620001da60201b62000a851760201c565b604f80546001600160a01b0319166001600160a01b039290921691909117905562000120600683603d62000329565b5062000130604382600a6200036c565b505050505062000550565b600080806200014a846200027c565b91509150816200015c57509192915050565b6000806200016a83620002af565b91509150816200017e575093949350505050565b95945050505050565b600180546001600160a01b0319169055620001b76001600160a01b0382166200013b602090811b62000a3e17901c565b600080546001600160a01b0319166001600160a01b039290921691909117905550565b60408051808201825260128152712927aaaa22a92fa1a922a22fa820a929a2a960711b60209091015251630d37324f60e11b815263d26df3b760e01b60048201526000906001600160a01b03831690631a6e649e90602401602060405180830381865afa15801562000250573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200027691906200052b565b92915050565b600080600160401b600160a01b03831660ff60981b8103620002a957600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620003035760009250600091505b508115806200031357503d601614155b156200032457506000928392509050565b915091565b82603d81019282156200035a579160200282015b828111156200035a5782518255916020019190600101906200033d565b50620003689291506200039c565b5090565b82600a81019282156200035a57916020028201828111156200035a5782518255916020019190600101906200033d565b5b808211156200036857600081556001016200039d565b80516001600160a01b0381168114620003cb57600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b6040516107a081016001600160401b03811182821017156200040c576200040c620003d0565b60405290565b6040516000906101408082016001600160401b03811183821017156200043c576200043c620003d0565b60405290915081908301848111156200045457600080fd5b835b818110156200047057805183526020928301920162000456565b50505092915050565b60008060008061092085870312156200049157600080fd5b6200049c85620003b3565b93506020620004ad818701620003b3565b935086605f870112620004bf57600080fd5b620004c9620003e6565b806107e0880189811115620004dd57600080fd5b604089015b81811015620004fb5780518452928401928401620004e2565b50819550896107ff8a01126200051057600080fd5b6200051c8a8262000412565b94505050505092959194509250565b6000602082840312156200053e57600080fd5b6200054982620003b3565b9392505050565b6080516117406200057360003960008181610477015261054601526117406000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c80639fc2f56b116100f9578063e30c397811610097578063eb36e96311610071578063eb36e963146103c6578063edf2460b146103d9578063f2fde38b146103ec578063f384bd05146103ff57600080fd5b8063e30c397814610380578063e4649162146103a0578063e5a8a2cb146103b357600080fd5b8063b2596a67116100d3578063b2596a6714610353578063d6b7494f14610366578063db29c1701461036e578063e08fdb551461037757600080fd5b80639fc2f56b14610324578063a0c9119a1461032d578063a73292f01461034057600080fd5b8063773fcf64116101665780638da5cb5b116101405780638da5cb5b146102ae5780639741fbfa146102ce5780639915bcad146102f15780639c18625f1461030457600080fd5b8063773fcf641461025957806377c0f07d1461026157806379ba5097146102a657600080fd5b80631f68f20a116101a25780631f68f20a146102045780633705afb614610220578063492bff0e14610233578063509edfb81461024657600080fd5b806308a0c375146101c9578063117573c5146101de5780631d08837b146101f1575b600080fd5b6101dc6101d736600461110b565b610408565b005b6101dc6101ec36600461110b565b610415565b6101dc6101ff36600461110b565b610422565b61020d604d5481565b6040519081526020015b60405180910390f35b61020d61022e366004611124565b61042f565b6101dc610241366004611146565b610444565b6101dc61025436600461110b565b61045d565b6101dc61046a565b604f546102819073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610217565b6101dc6104e2565b6000546102819073ffffffffffffffffffffffffffffffffffffffff1681565b6102e16102dc36600461118a565b61053e565b6040519015158152602001610217565b6101dc6102ff36600461110b565b6107ba565b61020d61031236600461110b565b604e6020526000908152604090205481565b61020d60035481565b6101dc61033b36600461122c565b6107c7565b61020d61034e36600461110b565b610875565b61020d61036136600461110b565b61088c565b61020d61089c565b61020d60055481565b61020d60025481565b6001546102819073ffffffffffffffffffffffffffffffffffffffff1681565b61020d6103ae366004611124565b6108ae565b6101dc6103c1366004611298565b6108f2565b61020d6103d43660046112ac565b610907565b61020d6103e73660046112e1565b610939565b6101dc6103fa366004611348565b6109cf565b61020d60045481565b610410610b69565b600455565b61041d610b69565b600355565b61042a610b69565b604d55565b600061043b8383610bba565b90505b92915050565b61044c610b69565b610459600682603d61108b565b5050565b610465610b69565b600555565b610472610b69565b61049b7f0000000000000000000000000000000000000000000000000000000000000000610a85565b604f80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60015473ffffffffffffffffffffffffffffffffffffffff163314610533576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61053c33610bcf565b565b60004361056a7f0000000000000000000000000000000000000000000000000000000000000000610c5e565b73ffffffffffffffffffffffffffffffffffffffff1663ac76eb406040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105d8919061136c565b6105e69060408601356113b4565b10156105f45750600061043e565b6043604e60008460200135815260200190815260200160002054600a811061061e5761061e6113c7565b0154836020013511156106335750600061043e565b604f546000906106629073ffffffffffffffffffffffffffffffffffffffff1661065c85611542565b90610d05565b90508060000361067657600191505061043e565b604f546000906106a59073ffffffffffffffffffffffffffffffffffffffff1661069f86611542565b90610d9e565b90506005548110156106bc5760009250505061043e565b604f546000906106eb9073ffffffffffffffffffffffffffffffffffffffff166106e587611542565b90610df6565b905080600003610701576000935050505061043e565b60045461070e848361042f565b1115610720576000935050505061043e565b600354604f546107559085906103ae9073ffffffffffffffffffffffffffffffffffffffff1661074f8a611542565b90610e4e565b1115610767576000935050505061043e565b600254604f546107ae9061079a9073ffffffffffffffffffffffffffffffffffffffff1661079489611542565b90610ea6565b6107a385610efe565b896020013587610939565b11159695505050505050565b6107c2610b69565b600255565b6107cf610b69565b828114610808576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b8381101561086e57828282818110610825576108256113c7565b90506020020135604e6000878785818110610842576108426113c7565b905060200201358152602001908152602001600020819055508080610866906115e4565b91505061080b565b5050505050565b600681603d811061088557600080fd5b0154905081565b604381600a811061088557600080fd5b60006108a9600554610efe565b905090565b60008282116108de57507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61043e565b61043b6108eb848461161c565b8490610bba565b6108fa610b69565b610459604382600a6110c9565b604f5460009061043e906109349073ffffffffffffffffffffffffffffffffffffffff1661069f85611542565b610efe565b6000806109468484610bba565b905060006109548288610f2e565b905080600003610988577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff925050506109c7565b60006109b561099a601e610e1061162f565b6109a590601861166a565b6109af888a610f43565b90610f43565b90506109c18183610f58565b93505050505b949350505050565b6109d7610b69565b6109f68173ffffffffffffffffffffffffffffffffffffffff16610a3e565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b6000806000610a4c84610f6d565b9150915081610a5d57509192915050565b600080610a6983610fba565b9150915081610a7c575093949350505050565b95945050505050565b604080518082018252601281527f524f555445525f435245445f5041525345520000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527fd26df3b700000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024015b602060405180830381865afa158015610b45573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043e9190611681565b60005473ffffffffffffffffffffffffffffffffffffffff16331461053c576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061043b83670de0b6b3a76400008461103e565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055610c1673ffffffffffffffffffffffffffffffffffffffff8216610a3e565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b604080518082018252601381527f524f555445525f4147454e545f504f4c49434500000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f6ea276a300000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e90602401610b28565b60e08201516040517f17bbd29a00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff8416916317bbd29a91610d5d9160040161169e565b602060405180830381865afa158015610d7a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043b919061136c565b60e08201516040517f68fd63d300000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff8416916368fd63d391610d5d9160040161169e565b60e08201516040517f1b2b5fad00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff841691631b2b5fad91610d5d9160040161169e565b60e08201516040517f6bb0d0cc00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff841691636bb0d0cc91610d5d9160040161169e565b60e08201516040517fd9bc1b7300000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff84169163d9bc1b7391610d5d9160040161169e565b600061043e600660055484610f13919061161c565b603d8110610f2357610f236113c7565b0154604d5490610f43565b600061043b8383670de0b6b3a764000061103e565b600061043b8383670de0b6b3a764000061105d565b600061043b83670de0b6b3a76400008461105d565b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103610fb4576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a811461101a5760009250600091505b5081158061102957503d601614155b1561103957506000928392509050565b915091565b82820281151584158583048514171661105657600080fd5b0492915050565b82820281151584158583048514171661107557600080fd5b6001826001830304018115150290509392505050565b82603d81019282156110b9579160200282015b828111156110b957823582559160200191906001019061109e565b506110c59291506110f6565b5090565b82600a81019282156110b957916020028201828111156110b957823582559160200191906001019061109e565b5b808211156110c557600081556001016110f7565b60006020828403121561111d57600080fd5b5035919050565b6000806040838503121561113757600080fd5b50508035926020909101359150565b60006107a080838503121561115a57600080fd5b83818401111561116957600080fd5b509092915050565b6000610100828403121561118457600080fd5b50919050565b60008082840360a081121561119e57600080fd5b60808112156111ac57600080fd5b50829150608083013567ffffffffffffffff8111156111ca57600080fd5b6111d685828601611171565b9150509250929050565b60008083601f8401126111f257600080fd5b50813567ffffffffffffffff81111561120a57600080fd5b6020830191508360208260051b850101111561122557600080fd5b9250929050565b6000806000806040858703121561124257600080fd5b843567ffffffffffffffff8082111561125a57600080fd5b611266888389016111e0565b9096509450602087013591508082111561127f57600080fd5b5061128c878288016111e0565b95989497509550505050565b600061014080838503121561115a57600080fd5b6000602082840312156112be57600080fd5b813567ffffffffffffffff8111156112d557600080fd5b6109c784828501611171565b600080600080608085870312156112f757600080fd5b5050823594602084013594506040840135936060013592509050565b73ffffffffffffffffffffffffffffffffffffffff8116811461133557600080fd5b50565b803561134381611313565b919050565b60006020828403121561135a57600080fd5b813561136581611313565b9392505050565b60006020828403121561137e57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561043e5761043e611385565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715611449576114496113f6565b60405290565b80357fffffffff000000000000000000000000000000000000000000000000000000008116811461134357600080fd5b803567ffffffffffffffff8116811461134357600080fd5b600082601f8301126114a857600080fd5b813567ffffffffffffffff808211156114c3576114c36113f6565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611509576115096113f6565b8160405283815286602085880101111561152257600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000610100823603121561155557600080fd5b61155d611425565b61156683611338565b81526020830135602082015260408301356040820152606083013560608201526080830135608082015261159c60a0840161144f565b60a08201526115ad60c0840161147f565b60c082015260e083013567ffffffffffffffff8111156115cc57600080fd5b6115d836828601611497565b60e08301525092915050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361161557611615611385565b5060010190565b8181038181111561043e5761043e611385565b600082611665577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b808202811582820484141761043e5761043e611385565b60006020828403121561169357600080fd5b815161136581611313565b600060208083528351808285015260005b818110156116cb578581018301518582016040015282016116af565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f830116850101925050509291505056fea2646970667358221220288b3508b2e666236d0620e316d0c1b40f28c34239c9686d994f655ee648641764736f6c63430008110033",
}

// RateModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use RateModuleMetaData.ABI instead.
var RateModuleABI = RateModuleMetaData.ABI

// RateModuleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RateModuleMetaData.Bin instead.
var RateModuleBin = RateModuleMetaData.Bin

// DeployRateModule deploys a new Ethereum contract, binding an instance of RateModule to it.
func DeployRateModule(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _router common.Address, _rateLookup [61]*big.Int, _levels [10]*big.Int) (common.Address, *types.Transaction, *RateModule, error) {
	parsed, err := RateModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RateModuleBin), backend, _owner, _router, _rateLookup, _levels)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RateModule{RateModuleCaller: RateModuleCaller{contract: contract}, RateModuleTransactor: RateModuleTransactor{contract: contract}, RateModuleFilterer: RateModuleFilterer{contract: contract}}, nil
}

// RateModule is an auto generated Go binding around an Ethereum contract.
type RateModule struct {
	RateModuleCaller     // Read-only binding to the contract
	RateModuleTransactor // Write-only binding to the contract
	RateModuleFilterer   // Log filterer for contract events
}

// RateModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type RateModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RateModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RateModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RateModuleSession struct {
	Contract     *RateModule       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RateModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RateModuleCallerSession struct {
	Contract *RateModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RateModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RateModuleTransactorSession struct {
	Contract     *RateModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RateModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type RateModuleRaw struct {
	Contract *RateModule // Generic contract binding to access the raw methods on
}

// RateModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RateModuleCallerRaw struct {
	Contract *RateModuleCaller // Generic read-only contract binding to access the raw methods on
}

// RateModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RateModuleTransactorRaw struct {
	Contract *RateModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRateModule creates a new instance of RateModule, bound to a specific deployed contract.
func NewRateModule(address common.Address, backend bind.ContractBackend) (*RateModule, error) {
	contract, err := bindRateModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RateModule{RateModuleCaller: RateModuleCaller{contract: contract}, RateModuleTransactor: RateModuleTransactor{contract: contract}, RateModuleFilterer: RateModuleFilterer{contract: contract}}, nil
}

// NewRateModuleCaller creates a new read-only instance of RateModule, bound to a specific deployed contract.
func NewRateModuleCaller(address common.Address, caller bind.ContractCaller) (*RateModuleCaller, error) {
	contract, err := bindRateModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RateModuleCaller{contract: contract}, nil
}

// NewRateModuleTransactor creates a new write-only instance of RateModule, bound to a specific deployed contract.
func NewRateModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*RateModuleTransactor, error) {
	contract, err := bindRateModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RateModuleTransactor{contract: contract}, nil
}

// NewRateModuleFilterer creates a new log filterer instance of RateModule, bound to a specific deployed contract.
func NewRateModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*RateModuleFilterer, error) {
	contract, err := bindRateModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RateModuleFilterer{contract: contract}, nil
}

// bindRateModule binds a generic wrapper to an already deployed contract.
func bindRateModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RateModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RateModule *RateModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RateModule.Contract.RateModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RateModule *RateModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.Contract.RateModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RateModule *RateModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RateModule.Contract.RateModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RateModule *RateModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RateModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RateModule *RateModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RateModule *RateModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RateModule.Contract.contract.Transact(opts, method, params...)
}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCaller) AccountLevel(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "accountLevel", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_RateModule *RateModuleSession) AccountLevel(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.AccountLevel(&_RateModule.CallOpts, arg0)
}

// AccountLevel is a free data retrieval call binding the contract method 0x9c18625f.
//
// Solidity: function accountLevel(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCallerSession) AccountLevel(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.AccountLevel(&_RateModule.CallOpts, arg0)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_RateModule *RateModuleCaller) BaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "baseRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_RateModule *RateModuleSession) BaseRate() (*big.Int, error) {
	return _RateModule.Contract.BaseRate(&_RateModule.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_RateModule *RateModuleCallerSession) BaseRate() (*big.Int, error) {
	return _RateModule.Contract.BaseRate(&_RateModule.CallOpts)
}

// ComputeDTE is a free data retrieval call binding the contract method 0xe4649162.
//
// Solidity: function computeDTE(uint256 principal, uint256 agentTotalValue) pure returns(uint256)
func (_RateModule *RateModuleCaller) ComputeDTE(opts *bind.CallOpts, principal *big.Int, agentTotalValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "computeDTE", principal, agentTotalValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeDTE is a free data retrieval call binding the contract method 0xe4649162.
//
// Solidity: function computeDTE(uint256 principal, uint256 agentTotalValue) pure returns(uint256)
func (_RateModule *RateModuleSession) ComputeDTE(principal *big.Int, agentTotalValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTE(&_RateModule.CallOpts, principal, agentTotalValue)
}

// ComputeDTE is a free data retrieval call binding the contract method 0xe4649162.
//
// Solidity: function computeDTE(uint256 principal, uint256 agentTotalValue) pure returns(uint256)
func (_RateModule *RateModuleCallerSession) ComputeDTE(principal *big.Int, agentTotalValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTE(&_RateModule.CallOpts, principal, agentTotalValue)
}

// ComputeDTI is a free data retrieval call binding the contract method 0xedf2460b.
//
// Solidity: function computeDTI(uint256 expectedDailyRewards, uint256 rate, uint256 accountPrincipal, uint256 totalPrincipal) pure returns(uint256)
func (_RateModule *RateModuleCaller) ComputeDTI(opts *bind.CallOpts, expectedDailyRewards *big.Int, rate *big.Int, accountPrincipal *big.Int, totalPrincipal *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "computeDTI", expectedDailyRewards, rate, accountPrincipal, totalPrincipal)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeDTI is a free data retrieval call binding the contract method 0xedf2460b.
//
// Solidity: function computeDTI(uint256 expectedDailyRewards, uint256 rate, uint256 accountPrincipal, uint256 totalPrincipal) pure returns(uint256)
func (_RateModule *RateModuleSession) ComputeDTI(expectedDailyRewards *big.Int, rate *big.Int, accountPrincipal *big.Int, totalPrincipal *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTI(&_RateModule.CallOpts, expectedDailyRewards, rate, accountPrincipal, totalPrincipal)
}

// ComputeDTI is a free data retrieval call binding the contract method 0xedf2460b.
//
// Solidity: function computeDTI(uint256 expectedDailyRewards, uint256 rate, uint256 accountPrincipal, uint256 totalPrincipal) pure returns(uint256)
func (_RateModule *RateModuleCallerSession) ComputeDTI(expectedDailyRewards *big.Int, rate *big.Int, accountPrincipal *big.Int, totalPrincipal *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeDTI(&_RateModule.CallOpts, expectedDailyRewards, rate, accountPrincipal, totalPrincipal)
}

// ComputeLTV is a free data retrieval call binding the contract method 0x3705afb6.
//
// Solidity: function computeLTV(uint256 principal, uint256 collateralValue) pure returns(uint256)
func (_RateModule *RateModuleCaller) ComputeLTV(opts *bind.CallOpts, principal *big.Int, collateralValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "computeLTV", principal, collateralValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeLTV is a free data retrieval call binding the contract method 0x3705afb6.
//
// Solidity: function computeLTV(uint256 principal, uint256 collateralValue) pure returns(uint256)
func (_RateModule *RateModuleSession) ComputeLTV(principal *big.Int, collateralValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeLTV(&_RateModule.CallOpts, principal, collateralValue)
}

// ComputeLTV is a free data retrieval call binding the contract method 0x3705afb6.
//
// Solidity: function computeLTV(uint256 principal, uint256 collateralValue) pure returns(uint256)
func (_RateModule *RateModuleCallerSession) ComputeLTV(principal *big.Int, collateralValue *big.Int) (*big.Int, error) {
	return _RateModule.Contract.ComputeLTV(&_RateModule.CallOpts, principal, collateralValue)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_RateModule *RateModuleCaller) CredParser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "credParser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_RateModule *RateModuleSession) CredParser() (common.Address, error) {
	return _RateModule.Contract.CredParser(&_RateModule.CallOpts)
}

// CredParser is a free data retrieval call binding the contract method 0x77c0f07d.
//
// Solidity: function credParser() view returns(address)
func (_RateModule *RateModuleCallerSession) CredParser() (common.Address, error) {
	return _RateModule.Contract.CredParser(&_RateModule.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_RateModule *RateModuleCaller) GetRate(opts *bind.CallOpts, vc VerifiableCredential) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "getRate", vc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_RateModule *RateModuleSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _RateModule.Contract.GetRate(&_RateModule.CallOpts, vc)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_RateModule *RateModuleCallerSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _RateModule.Contract.GetRate(&_RateModule.CallOpts, vc)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_RateModule *RateModuleCaller) IsApproved(opts *bind.CallOpts, account Account, vc VerifiableCredential) (bool, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "isApproved", account, vc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_RateModule *RateModuleSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _RateModule.Contract.IsApproved(&_RateModule.CallOpts, account, vc)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool)
func (_RateModule *RateModuleCallerSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _RateModule.Contract.IsApproved(&_RateModule.CallOpts, account, vc)
}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCaller) Levels(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "levels", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_RateModule *RateModuleSession) Levels(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.Levels(&_RateModule.CallOpts, arg0)
}

// Levels is a free data retrieval call binding the contract method 0xb2596a67.
//
// Solidity: function levels(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCallerSession) Levels(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.Levels(&_RateModule.CallOpts, arg0)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_RateModule *RateModuleCaller) MaxDTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "maxDTE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_RateModule *RateModuleSession) MaxDTE() (*big.Int, error) {
	return _RateModule.Contract.MaxDTE(&_RateModule.CallOpts)
}

// MaxDTE is a free data retrieval call binding the contract method 0x9fc2f56b.
//
// Solidity: function maxDTE() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MaxDTE() (*big.Int, error) {
	return _RateModule.Contract.MaxDTE(&_RateModule.CallOpts)
}

// MaxDTI is a free data retrieval call binding the contract method 0xe08fdb55.
//
// Solidity: function maxDTI() view returns(uint256)
func (_RateModule *RateModuleCaller) MaxDTI(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "maxDTI")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDTI is a free data retrieval call binding the contract method 0xe08fdb55.
//
// Solidity: function maxDTI() view returns(uint256)
func (_RateModule *RateModuleSession) MaxDTI() (*big.Int, error) {
	return _RateModule.Contract.MaxDTI(&_RateModule.CallOpts)
}

// MaxDTI is a free data retrieval call binding the contract method 0xe08fdb55.
//
// Solidity: function maxDTI() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MaxDTI() (*big.Int, error) {
	return _RateModule.Contract.MaxDTI(&_RateModule.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_RateModule *RateModuleCaller) MaxLTV(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "maxLTV")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_RateModule *RateModuleSession) MaxLTV() (*big.Int, error) {
	return _RateModule.Contract.MaxLTV(&_RateModule.CallOpts)
}

// MaxLTV is a free data retrieval call binding the contract method 0xf384bd05.
//
// Solidity: function maxLTV() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MaxLTV() (*big.Int, error) {
	return _RateModule.Contract.MaxLTV(&_RateModule.CallOpts)
}

// MinGCRED is a free data retrieval call binding the contract method 0xdb29c170.
//
// Solidity: function minGCRED() view returns(uint256)
func (_RateModule *RateModuleCaller) MinGCRED(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "minGCRED")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinGCRED is a free data retrieval call binding the contract method 0xdb29c170.
//
// Solidity: function minGCRED() view returns(uint256)
func (_RateModule *RateModuleSession) MinGCRED() (*big.Int, error) {
	return _RateModule.Contract.MinGCRED(&_RateModule.CallOpts)
}

// MinGCRED is a free data retrieval call binding the contract method 0xdb29c170.
//
// Solidity: function minGCRED() view returns(uint256)
func (_RateModule *RateModuleCallerSession) MinGCRED() (*big.Int, error) {
	return _RateModule.Contract.MinGCRED(&_RateModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RateModule *RateModuleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RateModule *RateModuleSession) Owner() (common.Address, error) {
	return _RateModule.Contract.Owner(&_RateModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RateModule *RateModuleCallerSession) Owner() (common.Address, error) {
	return _RateModule.Contract.Owner(&_RateModule.CallOpts)
}

// PenaltyRate is a free data retrieval call binding the contract method 0xd6b7494f.
//
// Solidity: function penaltyRate() view returns(uint256 rate)
func (_RateModule *RateModuleCaller) PenaltyRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "penaltyRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PenaltyRate is a free data retrieval call binding the contract method 0xd6b7494f.
//
// Solidity: function penaltyRate() view returns(uint256 rate)
func (_RateModule *RateModuleSession) PenaltyRate() (*big.Int, error) {
	return _RateModule.Contract.PenaltyRate(&_RateModule.CallOpts)
}

// PenaltyRate is a free data retrieval call binding the contract method 0xd6b7494f.
//
// Solidity: function penaltyRate() view returns(uint256 rate)
func (_RateModule *RateModuleCallerSession) PenaltyRate() (*big.Int, error) {
	return _RateModule.Contract.PenaltyRate(&_RateModule.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RateModule *RateModuleCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RateModule *RateModuleSession) PendingOwner() (common.Address, error) {
	return _RateModule.Contract.PendingOwner(&_RateModule.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RateModule *RateModuleCallerSession) PendingOwner() (common.Address, error) {
	return _RateModule.Contract.PendingOwner(&_RateModule.CallOpts)
}

// RateLookup is a free data retrieval call binding the contract method 0xa73292f0.
//
// Solidity: function rateLookup(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCaller) RateLookup(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RateModule.contract.Call(opts, &out, "rateLookup", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RateLookup is a free data retrieval call binding the contract method 0xa73292f0.
//
// Solidity: function rateLookup(uint256 ) view returns(uint256)
func (_RateModule *RateModuleSession) RateLookup(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.RateLookup(&_RateModule.CallOpts, arg0)
}

// RateLookup is a free data retrieval call binding the contract method 0xa73292f0.
//
// Solidity: function rateLookup(uint256 ) view returns(uint256)
func (_RateModule *RateModuleCallerSession) RateLookup(arg0 *big.Int) (*big.Int, error) {
	return _RateModule.Contract.RateLookup(&_RateModule.CallOpts, arg0)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RateModule *RateModuleTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RateModule *RateModuleSession) AcceptOwnership() (*types.Transaction, error) {
	return _RateModule.Contract.AcceptOwnership(&_RateModule.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RateModule *RateModuleTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RateModule.Contract.AcceptOwnership(&_RateModule.TransactOpts)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_RateModule *RateModuleTransactor) SetAgentLevels(opts *bind.TransactOpts, agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setAgentLevels", agentIDs, level)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_RateModule *RateModuleSession) SetAgentLevels(agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetAgentLevels(&_RateModule.TransactOpts, agentIDs, level)
}

// SetAgentLevels is a paid mutator transaction binding the contract method 0xa0c9119a.
//
// Solidity: function setAgentLevels(uint256[] agentIDs, uint256[] level) returns()
func (_RateModule *RateModuleTransactorSession) SetAgentLevels(agentIDs []*big.Int, level []*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetAgentLevels(&_RateModule.TransactOpts, agentIDs, level)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 _baseRate) returns()
func (_RateModule *RateModuleTransactor) SetBaseRate(opts *bind.TransactOpts, _baseRate *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setBaseRate", _baseRate)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 _baseRate) returns()
func (_RateModule *RateModuleSession) SetBaseRate(_baseRate *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetBaseRate(&_RateModule.TransactOpts, _baseRate)
}

// SetBaseRate is a paid mutator transaction binding the contract method 0x1d08837b.
//
// Solidity: function setBaseRate(uint256 _baseRate) returns()
func (_RateModule *RateModuleTransactorSession) SetBaseRate(_baseRate *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetBaseRate(&_RateModule.TransactOpts, _baseRate)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_RateModule *RateModuleTransactor) SetLevels(opts *bind.TransactOpts, _levels [10]*big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setLevels", _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_RateModule *RateModuleSession) SetLevels(_levels [10]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetLevels(&_RateModule.TransactOpts, _levels)
}

// SetLevels is a paid mutator transaction binding the contract method 0xe5a8a2cb.
//
// Solidity: function setLevels(uint256[10] _levels) returns()
func (_RateModule *RateModuleTransactorSession) SetLevels(_levels [10]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetLevels(&_RateModule.TransactOpts, _levels)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_RateModule *RateModuleTransactor) SetMaxDTE(opts *bind.TransactOpts, _maxDTE *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMaxDTE", _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_RateModule *RateModuleSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTE(&_RateModule.TransactOpts, _maxDTE)
}

// SetMaxDTE is a paid mutator transaction binding the contract method 0x117573c5.
//
// Solidity: function setMaxDTE(uint256 _maxDTE) returns()
func (_RateModule *RateModuleTransactorSession) SetMaxDTE(_maxDTE *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTE(&_RateModule.TransactOpts, _maxDTE)
}

// SetMaxDTI is a paid mutator transaction binding the contract method 0x9915bcad.
//
// Solidity: function setMaxDTI(uint256 _maxDTI) returns()
func (_RateModule *RateModuleTransactor) SetMaxDTI(opts *bind.TransactOpts, _maxDTI *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMaxDTI", _maxDTI)
}

// SetMaxDTI is a paid mutator transaction binding the contract method 0x9915bcad.
//
// Solidity: function setMaxDTI(uint256 _maxDTI) returns()
func (_RateModule *RateModuleSession) SetMaxDTI(_maxDTI *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTI(&_RateModule.TransactOpts, _maxDTI)
}

// SetMaxDTI is a paid mutator transaction binding the contract method 0x9915bcad.
//
// Solidity: function setMaxDTI(uint256 _maxDTI) returns()
func (_RateModule *RateModuleTransactorSession) SetMaxDTI(_maxDTI *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxDTI(&_RateModule.TransactOpts, _maxDTI)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_RateModule *RateModuleTransactor) SetMaxLTV(opts *bind.TransactOpts, _maxLTV *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMaxLTV", _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_RateModule *RateModuleSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxLTV(&_RateModule.TransactOpts, _maxLTV)
}

// SetMaxLTV is a paid mutator transaction binding the contract method 0x08a0c375.
//
// Solidity: function setMaxLTV(uint256 _maxLTV) returns()
func (_RateModule *RateModuleTransactorSession) SetMaxLTV(_maxLTV *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMaxLTV(&_RateModule.TransactOpts, _maxLTV)
}

// SetMinGCRED is a paid mutator transaction binding the contract method 0x509edfb8.
//
// Solidity: function setMinGCRED(uint256 _minGCRED) returns()
func (_RateModule *RateModuleTransactor) SetMinGCRED(opts *bind.TransactOpts, _minGCRED *big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setMinGCRED", _minGCRED)
}

// SetMinGCRED is a paid mutator transaction binding the contract method 0x509edfb8.
//
// Solidity: function setMinGCRED(uint256 _minGCRED) returns()
func (_RateModule *RateModuleSession) SetMinGCRED(_minGCRED *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMinGCRED(&_RateModule.TransactOpts, _minGCRED)
}

// SetMinGCRED is a paid mutator transaction binding the contract method 0x509edfb8.
//
// Solidity: function setMinGCRED(uint256 _minGCRED) returns()
func (_RateModule *RateModuleTransactorSession) SetMinGCRED(_minGCRED *big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetMinGCRED(&_RateModule.TransactOpts, _minGCRED)
}

// SetRateLookup is a paid mutator transaction binding the contract method 0x492bff0e.
//
// Solidity: function setRateLookup(uint256[61] _rateLookup) returns()
func (_RateModule *RateModuleTransactor) SetRateLookup(opts *bind.TransactOpts, _rateLookup [61]*big.Int) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "setRateLookup", _rateLookup)
}

// SetRateLookup is a paid mutator transaction binding the contract method 0x492bff0e.
//
// Solidity: function setRateLookup(uint256[61] _rateLookup) returns()
func (_RateModule *RateModuleSession) SetRateLookup(_rateLookup [61]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetRateLookup(&_RateModule.TransactOpts, _rateLookup)
}

// SetRateLookup is a paid mutator transaction binding the contract method 0x492bff0e.
//
// Solidity: function setRateLookup(uint256[61] _rateLookup) returns()
func (_RateModule *RateModuleTransactorSession) SetRateLookup(_rateLookup [61]*big.Int) (*types.Transaction, error) {
	return _RateModule.Contract.SetRateLookup(&_RateModule.TransactOpts, _rateLookup)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RateModule *RateModuleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RateModule *RateModuleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RateModule.Contract.TransferOwnership(&_RateModule.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RateModule *RateModuleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RateModule.Contract.TransferOwnership(&_RateModule.TransactOpts, newOwner)
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_RateModule *RateModuleTransactor) UpdateCredParser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RateModule.contract.Transact(opts, "updateCredParser")
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_RateModule *RateModuleSession) UpdateCredParser() (*types.Transaction, error) {
	return _RateModule.Contract.UpdateCredParser(&_RateModule.TransactOpts)
}

// UpdateCredParser is a paid mutator transaction binding the contract method 0x773fcf64.
//
// Solidity: function updateCredParser() returns()
func (_RateModule *RateModuleTransactorSession) UpdateCredParser() (*types.Transaction, error) {
	return _RateModule.Contract.UpdateCredParser(&_RateModule.TransactOpts)
}
