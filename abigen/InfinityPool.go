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
type Account_duplicate struct {
	StartEpoch *big.Int
	Principal  *big.Int
	EpochsPaid *big.Int
	Defaulted  bool
}

// VerifiableCredential is an auto generated low-level Go binding around an user-defined struct.
type VerifiableCredential_duplicate struct {
	Issuer          common.Address
	Subject         *big.Int
	EpochIssued     *big.Int
	EpochValidUntil *big.Int
	Value           *big.Int
	Action          [4]byte
	Target          uint64
	Claim           []byte
}

// InfinityPoolMetaData contains all meta data concerning the InfinityPool contract.
var InfinityPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rateModule\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_liquidStakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_preStake\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountDNE\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyDefaulted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CallFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientFunds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidState\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PayUp\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolShuttingDown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lostFunds\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"}],\"name\":\"WriteOff\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"asset\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"convertToAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"convertToShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPool\",\"name\":\"newPool\",\"type\":\"address\"}],\"name\":\"decommissionPool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowedAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feesCollected\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAbsMinLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"}],\"name\":\"getAgentBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLiquidAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"getRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"harvestAmount\",\"type\":\"uint256\"}],\"name\":\"harvestFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvestToRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"id\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"defaulted\",\"type\":\"bool\"}],\"internalType\":\"structAccount\",\"name\":\"account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isShuttingDown\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accountPrincipal\",\"type\":\"uint256\"}],\"name\":\"jumpStartAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"jumpStartTotalBorrowed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"liquidStakingToken\",\"outputs\":[{\"internalType\":\"contractIPoolToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxEpochsOwedTolerance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"maxWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subject\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochIssued\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochValidUntil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"action\",\"type\":\"bytes4\"},{\"internalType\":\"uint64\",\"name\":\"target\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"claim\",\"type\":\"bytes\"}],\"internalType\":\"structVerifiableCredential\",\"name\":\"vc\",\"type\":\"tuple\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochsPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refund\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preStake\",\"outputs\":[{\"internalType\":\"contractIPreStake\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"previewRedeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"name\":\"previewWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ramp\",\"outputs\":[{\"internalType\":\"contractIOffRamp\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rateModule\",\"outputs\":[{\"internalType\":\"contractIRateModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"recoverFIL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshRoutes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxEpochsOwedTolerance\",\"type\":\"uint256\"}],\"name\":\"setMaxEpochsOwedTolerance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumLiquidity\",\"type\":\"uint256\"}],\"name\":\"setMinimumLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOffRamp\",\"name\":\"_ramp\",\"type\":\"address\"}],\"name\":\"setRamp\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRateModule\",\"name\":\"_rateModule\",\"type\":\"address\"}],\"name\":\"setRateModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutDown\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowableAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBorrowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFromPreStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assets\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"agentID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"recoveredFunds\",\"type\":\"uint256\"}],\"name\":\"writeOff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalOwed\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x610140604052600060075560006008556000600955601e610e1062000025919062000529565b620000329060186200054c565b6200003f9060016200054c565b600a55600b805460ff191690553480156200005957600080fd5b5060405162005155380380620051558339810160408190526200007c916200058f565b876200009c816001600160a01b03166200020060201b620036c21760201c565b90506001600160a01b038116620000c657604051635435b28960e11b815260040160405180910390fd5b620000d1816200024c565b506001600160a01b03878116608081905287821660e052600580546001600160a01b03191688841617905584821661012052600984905560c0839052908516610100526200012b906200029f602090811b6200370917901c565b600260006101000a8154816001600160a01b0302191690836001600160a01b03160217905550620001696080516200034d60201b620037ed1760201c565b600360006101000a8154816001600160a01b0302191690836001600160a01b03160217905550620001a7608051620003b660201b620038941760201c565b600460006101000a8154816001600160a01b0302191690836001600160a01b03160217905550620001e56080516200041f60201b6200393b1760201c565b6001600160a01b031660a05250620006499650505050505050565b600080806200020f846200047c565b91509150816200022157509192915050565b6000806200022f83620004af565b915091508162000243575093949350505050565b95945050505050565b600180546001600160a01b03191690556200027c6001600160a01b03821662000200602090811b620036c217901c565b600080546001600160a01b0319166001600160a01b039290921691909117905550565b604080518082018252601481527f524f555445525f504f4f4c5f524547495354525900000000000000000000000060209091015251630d37324f60e11b8152631c86174160e11b60048201526000906001600160a01b03831690631a6e649e906024015b602060405180830381865afa15801562000321573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000347919062000624565b92915050565b604080518082018252601381527f524f555445525f4147454e545f504f4c4943450000000000000000000000000060209091015251630d37324f60e11b8152636ea276a360e01b60048201526000906001600160a01b03831690631a6e649e9060240162000303565b604080518082018252601481527f524f555445525f4147454e545f464143544f525900000000000000000000000060209091015251630d37324f60e11b81526314fb0b6d60e11b60048201526000906001600160a01b03831690631a6e649e9060240162000303565b60408051808201825260118152702927aaaa22a92faba324a62faa27a5a2a760791b60209091015251630d37324f60e11b815263aee0a13560e01b60048201526000906001600160a01b03831690631a6e649e9060240162000303565b600080600160401b600160a01b03831660ff60981b8103620004a957600192506001600160401b03841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa9150600051806001600160a01b031691508060a01c61ffff16905061040a8114620005035760009250600091505b508115806200051357503d601614155b156200052457506000928392509050565b915091565b6000826200054757634e487b7160e01b600052601260045260246000fd5b500490565b80820281158282048414176200034757634e487b7160e01b600052601160045260246000fd5b80516001600160a01b03811681146200058a57600080fd5b919050565b600080600080600080600080610100898b031215620005ad57600080fd5b620005b88962000572565b9750620005c860208a0162000572565b9650620005d860408a0162000572565b9550620005e860608a0162000572565b9450620005f860808a0162000572565b93506200060860a08a0162000572565b60c08a015160e0909a0151989b979a5095989497939692505050565b6000602082840312156200063757600080fd5b620006428262000572565b9392505050565b60805160a05160c05160e0516101005161012051614996620007bf60003960008181610a1d01528181610cee015261143001526000818161049001528181610be201528181610e54015281816129e5015281816131cf015281816135cc0152613f3801526000818161056601528181610dac0152818161111b01528181611465015281816115c1015281816116a601528181611b0b01528181611ca101528181611db70152818161204e01528181612413015281816125f00152818161269f01528181612916015281816130250152613e690152600081816108410152818161100d0152818161194a01528181611a5801528181611d5b0152818161225f0152818161236401528181613158015281816134fc015281816135520152613c3d01526000610b1a0152600081816112a4015281816113120152818161138001528181611a3201528181611d3901528181611de10152818161233e01528181613136015281816134da0152613c1b01526149966000f3fe6080604052600436106103595760003560e01c80637d694730116101bb578063c564f772116100f7578063ef8b30f711610095578063f1d72e251161006f578063f1d72e2514610a0b578063f2fde38b14610a3f578063f340fa0114610a5f578063f4765e9e14610a7257610371565b8063ef8b30f7146109c0578063f071db5a146109e0578063f1ccc3df146109f657610371565b8063ce96cb77116100d1578063ce96cb7714610933578063d905777e14610953578063e30c397814610973578063eb36e963146109a057610371565b8063c564f772146108f3578063c63d75b6146105bd578063c6e6f5921461091357610371565b80639741fbfa11610164578063b3d7f6b91161013e578063b3d7f6b91461087d578063b460af941461089d578063b56cf011146108bd578063ba087652146108d357610371565b80639741fbfa146107ff578063af640d0f1461082f578063b2bcb0021461086357610371565b8063886f039a11610195578063886f039a146107925780638da5cb5b146107b257806394bf804d146107df57610371565b80637d6947301461071d5780637edc8fe71461075d57806386a2c9881461077d57610371565b806338d52e0f116102955780634cdad506116102335780635fc2faf71161020d5780635fc2faf7146106a85780636e3ac566146106c85780636e553f65146106e857806379ba50971461070857610371565b80634cdad5061461065357806359221e38146106735780635d66b00a1461069357610371565b8063402d267d1161026f578063402d267d146105bd5780634107e644146105fd578063415e819d1461061d5780634c19386c1461063d57610371565b806338d52e0f1461055457806338dff49c146105885780633d9c3c171461059d57610371565b806315d276e11161030257806331ff92fa116102dc57806331ff92fa146104d2578063336d391e146104f2578063338891eb1461051f578063358bded01461053457610371565b806315d276e11461042c5780631cbe8df61461047e578063282567b4146104b257610371565b806310b9e5831161033357806310b9e583146103e1578063124dfd66146103f65780631387b4ed1461041657610371565b806301e1d1141461037957806307a2d13a146103a15780630a28a477146103c157610371565b3661037157610366610a92565b61036f33610ad1565b005b610366610a92565b34801561038557600080fd5b5061038e610ce7565b6040519081526020015b60405180910390f35b3480156103ad57600080fd5b5061038e6103bc3660046143bd565b610e4f565b3480156103cd57600080fd5b5061038e6103dc3660046143bd565b610f0a565b3480156103ed57600080fd5b5061036f610fcc565b34801561040257600080fd5b5061038e6104113660046143f8565b611001565b34801561042257600080fd5b5061038e600a5481565b34801561043857600080fd5b506006546104599073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610398565b34801561048a57600080fd5b506104597f000000000000000000000000000000000000000000000000000000000000000081565b3480156104be57600080fd5b5061036f6104cd3660046143bd565b611243565b3480156104de57600080fd5b5061036f6104ed3660046143f8565b611250565b3480156104fe57600080fd5b506005546104599073ffffffffffffffffffffffffffffffffffffffff1681565b34801561052b57600080fd5b5061036f61129f565b34801561054057600080fd5b5061036f61054f3660046143bd565b6113eb565b34801561056057600080fd5b506104597f000000000000000000000000000000000000000000000000000000000000000081565b34801561059457600080fd5b5061036f6114d7565b3480156105a957600080fd5b5061036f6105b83660046143f8565b6117a3565b3480156105c957600080fd5b5061038e6105d83660046143f8565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90565b34801561060957600080fd5b5061036f6106183660046143bd565b6117f2565b34801561062957600080fd5b5061036f61063836600461442e565b611850565b34801561064957600080fd5b5061038e60085481565b34801561065f57600080fd5b5061038e61066e3660046143bd565b611b93565b34801561067f57600080fd5b5061036f61068e3660046143bd565b611c12565b34801561069f57600080fd5b5061038e611c59565b3480156106b457600080fd5b5061038e6106c33660046143bd565b611d32565b3480156106d457600080fd5b5061036f6106e33660046143bd565b611d89565b3480156106f457600080fd5b5061038e610703366004614463565b611e5d565b34801561071457600080fd5b5061036f611e71565b34801561072957600080fd5b5061073d61073836600461442e565b611ecb565b604080519485526020850193909352918301526060820152608001610398565b34801561076957600080fd5b5061036f6107783660046143f8565b612561565b34801561078957600080fd5b5061038e612592565b34801561079e57600080fd5b5061036f6107ad366004614493565b612695565b3480156107be57600080fd5b506000546104599073ffffffffffffffffffffffffffffffffffffffff1681565b3480156107eb57600080fd5b5061038e6107fa366004614463565b612884565b34801561080b57600080fd5b5061081f61081a3660046144c1565b612aad565b6040519015158152602001610398565b34801561083b57600080fd5b5061038e7f000000000000000000000000000000000000000000000000000000000000000081565b34801561086f57600080fd5b50600b5461081f9060ff1681565b34801561088957600080fd5b5061038e6108983660046143bd565b612b47565b3480156108a957600080fd5b5061038e6108b8366004614517565b612b52565b3480156108c957600080fd5b5061038e60095481565b3480156108df57600080fd5b5061038e6108ee366004614517565b612cb8565b3480156108ff57600080fd5b5061038e61090e366004614559565b612e0f565b34801561091f57600080fd5b5061038e61092e3660046143bd565b6131ca565b34801561093f57600080fd5b5061038e61094e3660046143f8565b613278565b34801561095f57600080fd5b5061038e61096e3660046143f8565b6132f8565b34801561097f57600080fd5b506001546104599073ffffffffffffffffffffffffffffffffffffffff1681565b3480156109ac57600080fd5b5061038e6109bb36600461442e565b613378565b3480156109cc57600080fd5b5061038e6109db3660046143bd565b6133cf565b3480156109ec57600080fd5b5061038e60075481565b348015610a0257600080fd5b5061038e6133da565b348015610a1757600080fd5b506104597f000000000000000000000000000000000000000000000000000000000000000081565b348015610a4b57600080fd5b5061036f610a5a3660046143f8565b6133ea565b61038e610a6d3660046143f8565b613459565b348015610a7e57600080fd5b5061036f610a8d36600461457b565b613471565b600b5460ff1615610acf576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b565b600034600003610b0d576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b16346133cf565b90507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663d0e30db0346040518263ffffffff1660e01b81526004016000604051808303818588803b158015610b8057600080fd5b505af1158015610b94573d6000803e3d6000fd5b50506040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152602482018690527f00000000000000000000000000000000000000000000000000000000000000001693506340c10f19925060440190506020604051808303816000875af1158015610c2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5391906145be565b50610c738273ffffffffffffffffffffffffffffffffffffffff166136c2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d73484604051610cda929190918252602082015260400190565b60405180910390a3919050565b60006007547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663ec18154e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7b91906145db565b6008546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610e08573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e2c91906145db565b610e369190614623565b610e409190614623565b610e4a9190614636565b905090565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ebd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ee191906145db565b90508015610f0157610efc610ef4610ce7565b8490836139e2565b610f03565b825b9392505050565b60065460009073ffffffffffffffffffffffffffffffffffffffff16610f3257506000919050565b6006546040517f0a28a4770000000000000000000000000000000000000000000000000000000081526004810184905273ffffffffffffffffffffffffffffffffffffffff90911690630a28a477906024015b602060405180830381865afa158015610fa2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc691906145db565b92915050565b610fd4613a01565b600b80547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b600061100b613a52565b7f00000000000000000000000000000000000000000000000000000000000000008273ffffffffffffffffffffffffffffffffffffffff1663af640d0f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611077573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061109b91906145db565b1415806110ab5750600b5460ff16155b156110e2576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6110ed600754611d89565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a9059cbb90849083906370a0823190602401602060405180830381865afa158015611181573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111a591906145db565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af1158015611215573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123991906145be565b5050600854919050565b61124b613a01565b600955565b611258613a01565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6112c87f0000000000000000000000000000000000000000000000000000000000000000613709565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556113367f00000000000000000000000000000000000000000000000000000000000000006137ed565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff929092169190911790556113a47f0000000000000000000000000000000000000000000000000000000000000000613894565b600480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6113f3613a01565b6040517f23b872dd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000081166004830152306024830152604482018390527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064015b6020604051808303816000875af11580156114af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114d391906145be565b5050565b6114df613aa3565b600654604080517f0fede599000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff1691630fede5999160048083019260209291908290030181865afa15801561154f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061157391906145db565b6006546040517f70a0823100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff91821660048201529192506000917f0000000000000000000000000000000000000000000000000000000000000000909116906370a0823190602401602060405180830381865afa15801561160a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061162e91906145db565b9050818110156114d35760006116546116478385614636565b61164f611c59565b613af2565b6006546040517f095ea7b300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152602481018390529192507f0000000000000000000000000000000000000000000000000000000000000000169063095ea7b3906044016020604051808303816000875af11580156116ef573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171391906145be565b506006546040517ffb9321080000000000000000000000000000000000000000000000000000000081523060048201526024810183905273ffffffffffffffffffffffffffffffffffffffff9091169063fb93210890604401600060405180830381600087803b15801561178657600080fd5b505af115801561179a573d6000803e3d6000fd5b50505050505050565b6117ab613a01565b600680547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6117fa613a01565b611807601e610e10614649565b611812906018614684565b81111561184b576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a55565b611858610a92565b8061186281613b08565b670de0b6b3a7640000826080013510156118a8576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81608001356118b5612592565b10156118ed576040517fbb55fd2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006118fc8360200135613be9565b905080604001516000036119c85743808252604080830182905260025490517f6834f901000000000000000000000000000000000000000000000000000000008152602086013560048201527f0000000000000000000000000000000000000000000000000000000000000000602482015273ffffffffffffffffffffffffffffffffffffffff90911690636834f90190604401600060405180830381600087803b1580156119aa57600080fd5b505af11580156119be573d6000803e3d6000fd5b5050505050611a13565b43600a5482604001516119db9190614623565b1015611a13576040517f3213caaf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b826080013581602001818151611a299190614623565b905250611a7c817f000000000000000000000000000000000000000000000000000000000000000060208601357f0000000000000000000000000000000000000000000000000000000000000000613c61565b826080013560086000828254611a929190614623565b909155505060405160808401358152602080850135917f044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b910160405180910390a26040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152608084013560248201527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a9059cbb906044016020604051808303816000875af1158015611b69573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b8d91906145be565b50505050565b60065460009073ffffffffffffffffffffffffffffffffffffffff16611bbb57506000919050565b6006546040517f4cdad5060000000000000000000000000000000000000000000000000000000081526004810184905273ffffffffffffffffffffffffffffffffffffffff90911690634cdad50690602401610f85565b611c1a613a52565b60085415611c54576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600855565b6040517f70a08231000000000000000000000000000000000000000000000000000000008152306004820152600090819073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015611ce8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0c91906145db565b90506007548111611d1f57600091505090565b600754611d2c9082614636565b91505090565b6000611d7f7f0000000000000000000000000000000000000000000000000000000000000000837f0000000000000000000000000000000000000000000000000000000000000000613c6d565b6020015192915050565b8060076000828254611d9b9190614636565b909155505073ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb611e057f0000000000000000000000000000000000000000000000000000000000000000613d38565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e084901b16815273ffffffffffffffffffffffffffffffffffffffff909116600482015260248101849052604401611490565b6000611e67610a92565b610f038383613ddf565b60015473ffffffffffffffffffffffffffffffffffffffff163314611ec2576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610acf3361402c565b60008060008084611edb81613b08565b6000611eea8760200135613be9565b9050611ef68151151590565b611f2c576040517f260d436000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611f3587613378565b9550600080600083606001511561211f57611fe58a60800135600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663666010326040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fbb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fdf91906145db565b906140bb565b60076000828254611ff69190614623565b90915550506040517f23b872dd00000000000000000000000000000000000000000000000000000000815233600482015230602482015260808b0135604482015273ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303816000875af1158015612097573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120bb91906145be565b50604080518a8152600060208281018290529282018190526060820152908b0135907fbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff69060800160405180910390a250505060400151935060009250829150612559565b43846040015110156121ae57600084604001514361213d9190614636565b602086015190915061214f908b6140bb565b925082612168670de0b6b3a764000060808e0135614684565b10156121a0576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6121aa83826140bb565b9350505b828a60800135116121ec5760006121c960808c0135846140d0565b905080856040018181516121dd9190614623565b90525050506080890135612338565b60006121fc8460808d0135614636565b9050839150846020015181106122fe578460200151975087600860008282546122259190614636565b90915550506002546040517f8152a6c600000000000000000000000000000000000000000000000000000000815260208d013560048201527f0000000000000000000000000000000000000000000000000000000000000000602482015273ffffffffffffffffffffffffffffffffffffffff90911690638152a6c690604401600060405180830381600087803b1580156122bf57600080fd5b505af11580156122d3573d6000803e3d6000fd5b50505060208601516122e6915082614636565b60008087526020870181905260408701529650612336565b80975080600860008282546123139190614636565b92505081905550808560200181815161232c9190614636565b9052504360408601525b505b612388847f000000000000000000000000000000000000000000000000000000000000000060208d01357f0000000000000000000000000000000000000000000000000000000000000000613c61565b6123f981600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663666010326040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fbb573d6000803e3d6000fd5b6007600082825461240a9190614623565b925050819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3330898e6080013561245f9190614636565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff938416600482015292909116602483015260448201526064016020604051808303816000875af11580156124d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124fc91906145be565b5060408085015181518b8152602081810192909252918201899052606082018890528b0135907fbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff69060800160405180910390a25050506040015193505b509193509193565b612569613a01565b471561258f5761258f73ffffffffffffffffffffffffffffffffffffffff8216476140e5565b50565b600b5460009060ff16156125a65750600090565b6007546040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000919073ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015612637573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061265b91906145db565b6126659190614636565b905060006126716133da565b9050808210156126845760009250505090565b61268e8183614636565b9250505090565b61269d613a01565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603612722576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb61275d8473ffffffffffffffffffffffffffffffffffffffff166136c2565b6040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff8516906370a0823190602401602060405180830381865afa1580156127c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127eb91906145db565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af115801561285b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061287f91906145be565b505050565b600061288e610a92565b61289783612b47565b90508015806128a4575082155b156128db576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906323b872dd906064016020604051808303816000875af1158015612974573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061299891906145be565b506040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018590527f000000000000000000000000000000000000000000000000000000000000000016906340c10f19906044016020604051808303816000875af1158015612a2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a5291906145be565b50604080518281526020810185905273ffffffffffffffffffffffffffffffffffffffff84169133917fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d791015b60405180910390a392915050565b6005546040517f9741fbfa00000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff1690639741fbfa90612b069086908690600401614858565b602060405180830381865afa158015612b23573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0391906145be565b6000610fc682610e4f565b6000612b5c613aa3565b60065473ffffffffffffffffffffffffffffffffffffffff1663a318c1a4858585612b85610ce7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b168152600481019490945273ffffffffffffffffffffffffffffffffffffffff92831660248501529116604483015260648201526084016020604051808303816000875af1158015612c04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c2891906145db565b90508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db8785604051612ca9929190918252602082015260400190565b60405180910390a49392505050565b6000612cc2613aa3565b60065473ffffffffffffffffffffffffffffffffffffffff16639f40a7b3858585612ceb610ce7565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b168152600481019490945273ffffffffffffffffffffffffffffffffffffffff92831660248501529116604483015260648201526084016020604051808303816000875af1158015612d6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612d8e91906145db565b90508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ffbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db8488604051612ca9929190918252602082015260400190565b6000612e196141b9565b6000612e2484613be9565b9050806060015115612e62576040517ff29cb5e200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001606082015260208101516000818511612e7e576000612e88565b612e888286614636565b905060008082118015612e9e5750438460400151105b15612ffb57600554604080517fd6b7494f000000000000000000000000000000000000000000000000000000008152905160009273ffffffffffffffffffffffffffffffffffffffff169163d6b7494f9160048083019260209291908290030181865afa158015612f13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f3791906145db565b90506000856040015143612f4b9190614636565b90506000612f5d82611fdf88866140bb565b9050848111612f6c5780612f6e565b845b9350612fe184600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663666010326040518163ffffffff1660e01b8152600401602060405180830381865afa158015611fbb573d6000803e3d6000fd5b60076000828254612ff29190614623565b90915550505050505b600086841161300b576000613015565b6130158785614636565b90506130218483614623565b95507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd33308a8a1161306e5789613070565b8a5b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff938416600482015292909116602483015260448201526064016020604051808303816000875af11580156130e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061310d91906145be565b508460200151600860008282546131249190614636565b90915550506020850181905261317c857f00000000000000000000000000000000000000000000000000000000000000008a7f0000000000000000000000000000000000000000000000000000000000000000613c61565b604080518881526020810183905290810183905288907f31d6d43b5ab97f9739e3e78ab72e97d7575b3233f23e8e929160920032436aa59060600160405180910390a2505050505092915050565b6000807f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166318160ddd6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613238573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061325c91906145db565b90508015610f0157610efc81613270610ce7565b8591906139e2565b60065460009073ffffffffffffffffffffffffffffffffffffffff166132a057506000919050565b6006546040517fce96cb7700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529091169063ce96cb7790602401610f85565b60065460009073ffffffffffffffffffffffffffffffffffffffff1661332057506000919050565b6006546040517fd905777e00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529091169063d905777e90602401610f85565b6005546040517feb36e96300000000000000000000000000000000000000000000000000000000815260009173ffffffffffffffffffffffffffffffffffffffff169063eb36e96390610f8590859060040161489d565b6000610fc6826131ca565b6000610e4a600954611fdf610ce7565b6133f2613a01565b6134118173ffffffffffffffffffffffffffffffffffffffff166136c2565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b6000613463610a92565b610fc682610ad1565b919050565b613479613a01565b600061348483613be9565b905080602001516000146134c4576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60208101829052438082526040820152613520817f0000000000000000000000000000000000000000000000000000000000000000857f0000000000000000000000000000000000000000000000000000000000000000613c61565b6002546040517f6834f901000000000000000000000000000000000000000000000000000000008152600481018590527f0000000000000000000000000000000000000000000000000000000000000000602482015273ffffffffffffffffffffffffffffffffffffffff90911690636834f90190604401600060405180830381600087803b1580156135b257600080fd5b505af11580156135c6573d6000803e3d6000fd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166340c10f1985613610856131ca565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff909216600483015260248201526044016020604051808303816000875af1158015613680573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136a491906145be565b5081600860008282546136b79190614623565b909155505050505050565b60008060006136d08461420a565b91509150816136e157509192915050565b6000806136ed83614257565b9150915081613700575093949350505050565b95945050505050565b604080518082018252601481527f524f555445525f504f4f4c5f5245474953545259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f390c2e8200000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024015b602060405180830381865afa1580156137c9573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc691906148b0565b604080518082018252601381527f524f555445525f4147454e545f504f4c49434500000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f6ea276a300000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016137ac565b604080518082018252601481527f524f555445525f4147454e545f464143544f5259000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527f29f616da00000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016137ac565b604080518082018252601181527f524f555445525f5746494c5f544f4b454e000000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527faee0a13500000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016137ac565b8282028115158415858304851417166139fa57600080fd5b0492915050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610acf576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60025473ffffffffffffffffffffffffffffffffffffffff163314610acf576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065473ffffffffffffffffffffffffffffffffffffffff16610acf576040517fbaf3f0f700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000818310613b015781610f03565b5090919050565b60208101351580613bb25750600480546040517ffd66091e000000000000000000000000000000000000000000000000000000008152339281019290925260208301359173ffffffffffffffffffffffffffffffffffffffff9091169063fd66091e90602401602060405180830381865afa158015613b8b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613baf91906145db565b14155b1561258f576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613c1660405180608001604052806000815260200160008152602001600081526020016000151581525090565b610fc67f0000000000000000000000000000000000000000000000000000000000000000837f0000000000000000000000000000000000000000000000000000000000000000613c6d565b611b8d838383876142db565b613c9a60405180608001604052806000815260200160008152602001600081526020016000151581525090565b6040517f6361f6de000000000000000000000000000000000000000000000000000000008152600481018490526024810183905273ffffffffffffffffffffffffffffffffffffffff851690636361f6de90604401608060405180830381865afa158015613d0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d3091906148cd565b949350505050565b604080518082018252600f81527f524f555445525f54524541535552590000000000000000000000000000000000602090910152517f1a6e649e0000000000000000000000000000000000000000000000000000000081527fa1858d5f00000000000000000000000000000000000000000000000000000000600482015260009073ffffffffffffffffffffffffffffffffffffffff831690631a6e649e906024016137ac565b6000613dea836133cf565b9050821580613df7575080155b15613e2e576040517fa86b651200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018490527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906323b872dd906064016020604051808303816000875af1158015613ec7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613eeb91906145be565b506040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152602482018390527f000000000000000000000000000000000000000000000000000000000000000016906340c10f19906044016020604051808303816000875af1158015613f81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613fa591906145be565b50613fc58273ffffffffffffffffffffffffffffffffffffffff166136c2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d78584604051612a9f929190918252602082015260400190565b600180547fffffffffffffffffffffffff000000000000000000000000000000000000000016905561407373ffffffffffffffffffffffffffffffffffffffff82166136c2565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905550565b6000610f038383670de0b6b3a764000061438f565b6000610f0383670de0b6b3a7640000846139e2565b8047101561411f576040517f356680b700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008273ffffffffffffffffffffffffffffffffffffffff168260405160006040518083038185875af1925050503d8060008114614179576040519150601f19603f3d011682016040523d82523d6000602084013e61417e565b606091505b505090508061287f576040517f3204506f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60035473ffffffffffffffffffffffffffffffffffffffff163314610acf576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000008103614251576001925067ffffffffffffffff841691505b50915091565b600080826000526016600a6020600073fe000000000000000000000000000000000000025afa91506000518073ffffffffffffffffffffffffffffffffffffffff1691508060a01c61ffff16905061040a81146142b75760009250600091505b508115806142c657503d601614155b156142d657506000928392509050565b915091565b604080517fc7afd04b000000000000000000000000000000000000000000000000000000008152600481018590526024810184905282516044820152602083015160648201529082015160848201526060820151151560a482015273ffffffffffffffffffffffffffffffffffffffff85169063c7afd04b9060c401600060405180830381600087803b15801561437157600080fd5b505af1158015614385573d6000803e3d6000fd5b5050505050505050565b8282028115158415858304851417166143a757600080fd5b6001826001830304018115150290509392505050565b6000602082840312156143cf57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff8116811461258f57600080fd5b60006020828403121561440a57600080fd5b8135610f03816143d6565b6000610100828403121561442857600080fd5b50919050565b60006020828403121561444057600080fd5b813567ffffffffffffffff81111561445757600080fd5b613d3084828501614415565b6000806040838503121561447657600080fd5b823591506020830135614488816143d6565b809150509250929050565b600080604083850312156144a657600080fd5b82356144b1816143d6565b91506020830135614488816143d6565b60008082840360a08112156144d557600080fd5b60808112156144e357600080fd5b50829150608083013567ffffffffffffffff81111561450157600080fd5b61450d85828601614415565b9150509250929050565b60008060006060848603121561452c57600080fd5b83359250602084013561453e816143d6565b9150604084013561454e816143d6565b809150509250925092565b6000806040838503121561456c57600080fd5b50508035926020909101359150565b60008060006060848603121561459057600080fd5b833561459b816143d6565b95602085013595506040909401359392505050565b801515811461258f57600080fd5b6000602082840312156145d057600080fd5b8151610f03816145b0565b6000602082840312156145ed57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b80820180821115610fc657610fc66145f4565b81810381811115610fc657610fc66145f4565b60008261467f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8082028115828204841417610fc657610fc66145f4565b803567ffffffffffffffff8116811461346c57600080fd5b60008083357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18436030181126146e857600080fd5b830160208101925035905067ffffffffffffffff81111561470857600080fd5b80360382131561471757600080fd5b9250929050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b60006101008235614777816143d6565b73ffffffffffffffffffffffffffffffffffffffff81168552506020830135602085015260408301356040850152606083013560608501526080830135608085015260a08301357fffffffff00000000000000000000000000000000000000000000000000000000811681146147ec57600080fd5b7fffffffff000000000000000000000000000000000000000000000000000000001660a085015261481f60c0840161469b565b67ffffffffffffffff1660c085015261483b60e08401846146b3565b8260e087015261484e838701828461471e565b9695505050505050565b82358152602083013560208201526040830135604082015260006060840135614880816145b0565b80151560608401525060a06080830152613d3060a0830184614767565b602081526000610f036020830184614767565b6000602082840312156148c257600080fd5b8151610f03816143d6565b6000608082840312156148df57600080fd5b6040516080810181811067ffffffffffffffff82111715614929577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b80604052508251815260208301516020820152604083015160408201526060830151614954816145b0565b6060820152939250505056fea2646970667358221220f50795dffc0aaf091a2fc2ad71b4d42e4e6f6c5ed91358402257a04f016a858064736f6c63430008110033",
}

// InfinityPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use InfinityPoolMetaData.ABI instead.
var InfinityPoolABI = InfinityPoolMetaData.ABI

// InfinityPoolBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use InfinityPoolMetaData.Bin instead.
var InfinityPoolBin = InfinityPoolMetaData.Bin

// DeployInfinityPool deploys a new Ethereum contract, binding an instance of InfinityPool to it.
func DeployInfinityPool(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _router common.Address, _asset common.Address, _rateModule common.Address, _liquidStakingToken common.Address, _preStake common.Address, _minimumLiquidity *big.Int, _id *big.Int) (common.Address, *types.Transaction, *InfinityPool, error) {
	parsed, err := InfinityPoolMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(InfinityPoolBin), backend, _owner, _router, _asset, _rateModule, _liquidStakingToken, _preStake, _minimumLiquidity, _id)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &InfinityPool{InfinityPoolCaller: InfinityPoolCaller{contract: contract}, InfinityPoolTransactor: InfinityPoolTransactor{contract: contract}, InfinityPoolFilterer: InfinityPoolFilterer{contract: contract}}, nil
}

// InfinityPool is an auto generated Go binding around an Ethereum contract.
type InfinityPool struct {
	InfinityPoolCaller     // Read-only binding to the contract
	InfinityPoolTransactor // Write-only binding to the contract
	InfinityPoolFilterer   // Log filterer for contract events
}

// InfinityPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type InfinityPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InfinityPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InfinityPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InfinityPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InfinityPoolSession struct {
	Contract     *InfinityPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InfinityPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InfinityPoolCallerSession struct {
	Contract *InfinityPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// InfinityPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InfinityPoolTransactorSession struct {
	Contract     *InfinityPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// InfinityPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type InfinityPoolRaw struct {
	Contract *InfinityPool // Generic contract binding to access the raw methods on
}

// InfinityPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InfinityPoolCallerRaw struct {
	Contract *InfinityPoolCaller // Generic read-only contract binding to access the raw methods on
}

// InfinityPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InfinityPoolTransactorRaw struct {
	Contract *InfinityPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInfinityPool creates a new instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPool(address common.Address, backend bind.ContractBackend) (*InfinityPool, error) {
	contract, err := bindInfinityPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InfinityPool{InfinityPoolCaller: InfinityPoolCaller{contract: contract}, InfinityPoolTransactor: InfinityPoolTransactor{contract: contract}, InfinityPoolFilterer: InfinityPoolFilterer{contract: contract}}, nil
}

// NewInfinityPoolCaller creates a new read-only instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPoolCaller(address common.Address, caller bind.ContractCaller) (*InfinityPoolCaller, error) {
	contract, err := bindInfinityPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolCaller{contract: contract}, nil
}

// NewInfinityPoolTransactor creates a new write-only instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*InfinityPoolTransactor, error) {
	contract, err := bindInfinityPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolTransactor{contract: contract}, nil
}

// NewInfinityPoolFilterer creates a new log filterer instance of InfinityPool, bound to a specific deployed contract.
func NewInfinityPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*InfinityPoolFilterer, error) {
	contract, err := bindInfinityPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolFilterer{contract: contract}, nil
}

// bindInfinityPool binds a generic wrapper to an already deployed contract.
func bindInfinityPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InfinityPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfinityPool *InfinityPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfinityPool.Contract.InfinityPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfinityPool *InfinityPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.Contract.InfinityPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfinityPool *InfinityPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfinityPool.Contract.InfinityPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InfinityPool *InfinityPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InfinityPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InfinityPool *InfinityPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InfinityPool *InfinityPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InfinityPool.Contract.contract.Transact(opts, method, params...)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPool *InfinityPoolCaller) Asset(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "asset")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPool *InfinityPoolSession) Asset() (common.Address, error) {
	return _InfinityPool.Contract.Asset(&_InfinityPool.CallOpts)
}

// Asset is a free data retrieval call binding the contract method 0x38d52e0f.
//
// Solidity: function asset() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) Asset() (common.Address, error) {
	return _InfinityPool.Contract.Asset(&_InfinityPool.CallOpts)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) ConvertToAssets(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "convertToAssets", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToAssets(&_InfinityPool.CallOpts, shares)
}

// ConvertToAssets is a free data retrieval call binding the contract method 0x07a2d13a.
//
// Solidity: function convertToAssets(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) ConvertToAssets(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToAssets(&_InfinityPool.CallOpts, shares)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) ConvertToShares(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "convertToShares", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToShares(&_InfinityPool.CallOpts, assets)
}

// ConvertToShares is a free data retrieval call binding the contract method 0xc6e6f592.
//
// Solidity: function convertToShares(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) ConvertToShares(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.ConvertToShares(&_InfinityPool.CallOpts, assets)
}

// FeesCollected is a free data retrieval call binding the contract method 0xf071db5a.
//
// Solidity: function feesCollected() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) FeesCollected(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "feesCollected")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeesCollected is a free data retrieval call binding the contract method 0xf071db5a.
//
// Solidity: function feesCollected() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) FeesCollected() (*big.Int, error) {
	return _InfinityPool.Contract.FeesCollected(&_InfinityPool.CallOpts)
}

// FeesCollected is a free data retrieval call binding the contract method 0xf071db5a.
//
// Solidity: function feesCollected() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) FeesCollected() (*big.Int, error) {
	return _InfinityPool.Contract.FeesCollected(&_InfinityPool.CallOpts)
}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) GetAbsMinLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getAbsMinLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) GetAbsMinLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.GetAbsMinLiquidity(&_InfinityPool.CallOpts)
}

// GetAbsMinLiquidity is a free data retrieval call binding the contract method 0xf1ccc3df.
//
// Solidity: function getAbsMinLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) GetAbsMinLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.GetAbsMinLiquidity(&_InfinityPool.CallOpts)
}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) GetAgentBorrowed(opts *bind.CallOpts, agentID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getAgentBorrowed", agentID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) GetAgentBorrowed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.GetAgentBorrowed(&_InfinityPool.CallOpts, agentID)
}

// GetAgentBorrowed is a free data retrieval call binding the contract method 0x5fc2faf7.
//
// Solidity: function getAgentBorrowed(uint256 agentID) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) GetAgentBorrowed(agentID *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.GetAgentBorrowed(&_InfinityPool.CallOpts, agentID)
}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) GetLiquidAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getLiquidAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) GetLiquidAssets() (*big.Int, error) {
	return _InfinityPool.Contract.GetLiquidAssets(&_InfinityPool.CallOpts)
}

// GetLiquidAssets is a free data retrieval call binding the contract method 0x5d66b00a.
//
// Solidity: function getLiquidAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) GetLiquidAssets() (*big.Int, error) {
	return _InfinityPool.Contract.GetLiquidAssets(&_InfinityPool.CallOpts)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_InfinityPool *InfinityPoolCaller) GetRate(opts *bind.CallOpts, vc VerifiableCredential) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "getRate", vc)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_InfinityPool *InfinityPoolSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _InfinityPool.Contract.GetRate(&_InfinityPool.CallOpts, vc)
}

// GetRate is a free data retrieval call binding the contract method 0xeb36e963.
//
// Solidity: function getRate((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(uint256 rate)
func (_InfinityPool *InfinityPoolCallerSession) GetRate(vc VerifiableCredential) (*big.Int, error) {
	return _InfinityPool.Contract.GetRate(&_InfinityPool.CallOpts, vc)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) Id(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "id")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) Id() (*big.Int, error) {
	return _InfinityPool.Contract.Id(&_InfinityPool.CallOpts)
}

// Id is a free data retrieval call binding the contract method 0xaf640d0f.
//
// Solidity: function id() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) Id() (*big.Int, error) {
	return _InfinityPool.Contract.Id(&_InfinityPool.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool approved)
func (_InfinityPool *InfinityPoolCaller) IsApproved(opts *bind.CallOpts, account Account, vc VerifiableCredential) (bool, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "isApproved", account, vc)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool approved)
func (_InfinityPool *InfinityPoolSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _InfinityPool.Contract.IsApproved(&_InfinityPool.CallOpts, account, vc)
}

// IsApproved is a free data retrieval call binding the contract method 0x9741fbfa.
//
// Solidity: function isApproved((uint256,uint256,uint256,bool) account, (address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) view returns(bool approved)
func (_InfinityPool *InfinityPoolCallerSession) IsApproved(account Account, vc VerifiableCredential) (bool, error) {
	return _InfinityPool.Contract.IsApproved(&_InfinityPool.CallOpts, account, vc)
}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPool *InfinityPoolCaller) IsShuttingDown(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "isShuttingDown")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPool *InfinityPoolSession) IsShuttingDown() (bool, error) {
	return _InfinityPool.Contract.IsShuttingDown(&_InfinityPool.CallOpts)
}

// IsShuttingDown is a free data retrieval call binding the contract method 0xb2bcb002.
//
// Solidity: function isShuttingDown() view returns(bool)
func (_InfinityPool *InfinityPoolCallerSession) IsShuttingDown() (bool, error) {
	return _InfinityPool.Contract.IsShuttingDown(&_InfinityPool.CallOpts)
}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPool *InfinityPoolCaller) LiquidStakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "liquidStakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPool *InfinityPoolSession) LiquidStakingToken() (common.Address, error) {
	return _InfinityPool.Contract.LiquidStakingToken(&_InfinityPool.CallOpts)
}

// LiquidStakingToken is a free data retrieval call binding the contract method 0x1cbe8df6.
//
// Solidity: function liquidStakingToken() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) LiquidStakingToken() (common.Address, error) {
	return _InfinityPool.Contract.LiquidStakingToken(&_InfinityPool.CallOpts)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxDeposit(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxDeposit", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxDeposit(&_InfinityPool.CallOpts, arg0)
}

// MaxDeposit is a free data retrieval call binding the contract method 0x402d267d.
//
// Solidity: function maxDeposit(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxDeposit(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxDeposit(&_InfinityPool.CallOpts, arg0)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxEpochsOwedTolerance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxEpochsOwedTolerance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _InfinityPool.Contract.MaxEpochsOwedTolerance(&_InfinityPool.CallOpts)
}

// MaxEpochsOwedTolerance is a free data retrieval call binding the contract method 0x1387b4ed.
//
// Solidity: function maxEpochsOwedTolerance() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxEpochsOwedTolerance() (*big.Int, error) {
	return _InfinityPool.Contract.MaxEpochsOwedTolerance(&_InfinityPool.CallOpts)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxMint(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxMint", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxMint(&_InfinityPool.CallOpts, arg0)
}

// MaxMint is a free data retrieval call binding the contract method 0xc63d75b6.
//
// Solidity: function maxMint(address ) pure returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxMint(arg0 common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxMint(&_InfinityPool.CallOpts, arg0)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxRedeem(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxRedeem", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxRedeem(&_InfinityPool.CallOpts, owner)
}

// MaxRedeem is a free data retrieval call binding the contract method 0xd905777e.
//
// Solidity: function maxRedeem(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxRedeem(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxRedeem(&_InfinityPool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MaxWithdraw(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "maxWithdraw", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxWithdraw(&_InfinityPool.CallOpts, owner)
}

// MaxWithdraw is a free data retrieval call binding the contract method 0xce96cb77.
//
// Solidity: function maxWithdraw(address owner) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MaxWithdraw(owner common.Address) (*big.Int, error) {
	return _InfinityPool.Contract.MaxWithdraw(&_InfinityPool.CallOpts, owner)
}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) MinimumLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "minimumLiquidity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) MinimumLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.MinimumLiquidity(&_InfinityPool.CallOpts)
}

// MinimumLiquidity is a free data retrieval call binding the contract method 0xb56cf011.
//
// Solidity: function minimumLiquidity() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) MinimumLiquidity() (*big.Int, error) {
	return _InfinityPool.Contract.MinimumLiquidity(&_InfinityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPool *InfinityPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPool *InfinityPoolSession) Owner() (common.Address, error) {
	return _InfinityPool.Contract.Owner(&_InfinityPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) Owner() (common.Address, error) {
	return _InfinityPool.Contract.Owner(&_InfinityPool.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPool *InfinityPoolCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPool *InfinityPoolSession) PendingOwner() (common.Address, error) {
	return _InfinityPool.Contract.PendingOwner(&_InfinityPool.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) PendingOwner() (common.Address, error) {
	return _InfinityPool.Contract.PendingOwner(&_InfinityPool.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_InfinityPool *InfinityPoolCaller) PreStake(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "preStake")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_InfinityPool *InfinityPoolSession) PreStake() (common.Address, error) {
	return _InfinityPool.Contract.PreStake(&_InfinityPool.CallOpts)
}

// PreStake is a free data retrieval call binding the contract method 0xf1d72e25.
//
// Solidity: function preStake() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) PreStake() (common.Address, error) {
	return _InfinityPool.Contract.PreStake(&_InfinityPool.CallOpts)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewDeposit(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewDeposit", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewDeposit(&_InfinityPool.CallOpts, assets)
}

// PreviewDeposit is a free data retrieval call binding the contract method 0xef8b30f7.
//
// Solidity: function previewDeposit(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewDeposit(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewDeposit(&_InfinityPool.CallOpts, assets)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewMint(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewMint", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewMint(&_InfinityPool.CallOpts, shares)
}

// PreviewMint is a free data retrieval call binding the contract method 0xb3d7f6b9.
//
// Solidity: function previewMint(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewMint(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewMint(&_InfinityPool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewRedeem(opts *bind.CallOpts, shares *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewRedeem", shares)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewRedeem(&_InfinityPool.CallOpts, shares)
}

// PreviewRedeem is a free data retrieval call binding the contract method 0x4cdad506.
//
// Solidity: function previewRedeem(uint256 shares) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewRedeem(shares *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewRedeem(&_InfinityPool.CallOpts, shares)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) PreviewWithdraw(opts *bind.CallOpts, assets *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "previewWithdraw", assets)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewWithdraw(&_InfinityPool.CallOpts, assets)
}

// PreviewWithdraw is a free data retrieval call binding the contract method 0x0a28a477.
//
// Solidity: function previewWithdraw(uint256 assets) view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) PreviewWithdraw(assets *big.Int) (*big.Int, error) {
	return _InfinityPool.Contract.PreviewWithdraw(&_InfinityPool.CallOpts, assets)
}

// Ramp is a free data retrieval call binding the contract method 0x15d276e1.
//
// Solidity: function ramp() view returns(address)
func (_InfinityPool *InfinityPoolCaller) Ramp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "ramp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ramp is a free data retrieval call binding the contract method 0x15d276e1.
//
// Solidity: function ramp() view returns(address)
func (_InfinityPool *InfinityPoolSession) Ramp() (common.Address, error) {
	return _InfinityPool.Contract.Ramp(&_InfinityPool.CallOpts)
}

// Ramp is a free data retrieval call binding the contract method 0x15d276e1.
//
// Solidity: function ramp() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) Ramp() (common.Address, error) {
	return _InfinityPool.Contract.Ramp(&_InfinityPool.CallOpts)
}

// RateModule is a free data retrieval call binding the contract method 0x336d391e.
//
// Solidity: function rateModule() view returns(address)
func (_InfinityPool *InfinityPoolCaller) RateModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "rateModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RateModule is a free data retrieval call binding the contract method 0x336d391e.
//
// Solidity: function rateModule() view returns(address)
func (_InfinityPool *InfinityPoolSession) RateModule() (common.Address, error) {
	return _InfinityPool.Contract.RateModule(&_InfinityPool.CallOpts)
}

// RateModule is a free data retrieval call binding the contract method 0x336d391e.
//
// Solidity: function rateModule() view returns(address)
func (_InfinityPool *InfinityPoolCallerSession) RateModule() (common.Address, error) {
	return _InfinityPool.Contract.RateModule(&_InfinityPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) TotalAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "totalAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) TotalAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalAssets(&_InfinityPool.CallOpts)
}

// TotalAssets is a free data retrieval call binding the contract method 0x01e1d114.
//
// Solidity: function totalAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) TotalAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalAssets(&_InfinityPool.CallOpts)
}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) TotalBorrowableAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "totalBorrowableAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) TotalBorrowableAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowableAssets(&_InfinityPool.CallOpts)
}

// TotalBorrowableAssets is a free data retrieval call binding the contract method 0x86a2c988.
//
// Solidity: function totalBorrowableAssets() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) TotalBorrowableAssets() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowableAssets(&_InfinityPool.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPool *InfinityPoolCaller) TotalBorrowed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InfinityPool.contract.Call(opts, &out, "totalBorrowed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPool *InfinityPoolSession) TotalBorrowed() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowed(&_InfinityPool.CallOpts)
}

// TotalBorrowed is a free data retrieval call binding the contract method 0x4c19386c.
//
// Solidity: function totalBorrowed() view returns(uint256)
func (_InfinityPool *InfinityPoolCallerSession) TotalBorrowed() (*big.Int, error) {
	return _InfinityPool.Contract.TotalBorrowed(&_InfinityPool.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPool *InfinityPoolTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPool *InfinityPoolSession) AcceptOwnership() (*types.Transaction, error) {
	return _InfinityPool.Contract.AcceptOwnership(&_InfinityPool.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_InfinityPool *InfinityPoolTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _InfinityPool.Contract.AcceptOwnership(&_InfinityPool.TransactOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPool *InfinityPoolTransactor) Borrow(opts *bind.TransactOpts, vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "borrow", vc)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPool *InfinityPoolSession) Borrow(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Borrow(&_InfinityPool.TransactOpts, vc)
}

// Borrow is a paid mutator transaction binding the contract method 0x415e819d.
//
// Solidity: function borrow((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns()
func (_InfinityPool *InfinityPoolTransactorSession) Borrow(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Borrow(&_InfinityPool.TransactOpts, vc)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPool *InfinityPoolTransactor) DecommissionPool(opts *bind.TransactOpts, newPool common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "decommissionPool", newPool)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPool *InfinityPoolSession) DecommissionPool(newPool common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.DecommissionPool(&_InfinityPool.TransactOpts, newPool)
}

// DecommissionPool is a paid mutator transaction binding the contract method 0x124dfd66.
//
// Solidity: function decommissionPool(address newPool) returns(uint256 borrowedAmount)
func (_InfinityPool *InfinityPoolTransactorSession) DecommissionPool(newPool common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.DecommissionPool(&_InfinityPool.TransactOpts, newPool)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactor) Deposit(opts *bind.TransactOpts, assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "deposit", assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPool *InfinityPoolSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit(&_InfinityPool.TransactOpts, assets, receiver)
}

// Deposit is a paid mutator transaction binding the contract method 0x6e553f65.
//
// Solidity: function deposit(uint256 assets, address receiver) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactorSession) Deposit(assets *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit(&_InfinityPool.TransactOpts, assets, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactor) Deposit0(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "deposit0", receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPool *InfinityPoolSession) Deposit0(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit0(&_InfinityPool.TransactOpts, receiver)
}

// Deposit0 is a paid mutator transaction binding the contract method 0xf340fa01.
//
// Solidity: function deposit(address receiver) payable returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactorSession) Deposit0(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Deposit0(&_InfinityPool.TransactOpts, receiver)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPool *InfinityPoolTransactor) HarvestFees(opts *bind.TransactOpts, harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "harvestFees", harvestAmount)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPool *InfinityPoolSession) HarvestFees(harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestFees(&_InfinityPool.TransactOpts, harvestAmount)
}

// HarvestFees is a paid mutator transaction binding the contract method 0x6e3ac566.
//
// Solidity: function harvestFees(uint256 harvestAmount) returns()
func (_InfinityPool *InfinityPoolTransactorSession) HarvestFees(harvestAmount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestFees(&_InfinityPool.TransactOpts, harvestAmount)
}

// HarvestToRamp is a paid mutator transaction binding the contract method 0x38dff49c.
//
// Solidity: function harvestToRamp() returns()
func (_InfinityPool *InfinityPoolTransactor) HarvestToRamp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "harvestToRamp")
}

// HarvestToRamp is a paid mutator transaction binding the contract method 0x38dff49c.
//
// Solidity: function harvestToRamp() returns()
func (_InfinityPool *InfinityPoolSession) HarvestToRamp() (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestToRamp(&_InfinityPool.TransactOpts)
}

// HarvestToRamp is a paid mutator transaction binding the contract method 0x38dff49c.
//
// Solidity: function harvestToRamp() returns()
func (_InfinityPool *InfinityPoolTransactorSession) HarvestToRamp() (*types.Transaction, error) {
	return _InfinityPool.Contract.HarvestToRamp(&_InfinityPool.TransactOpts)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPool *InfinityPoolTransactor) JumpStartAccount(opts *bind.TransactOpts, receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "jumpStartAccount", receiver, agentID, accountPrincipal)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPool *InfinityPoolSession) JumpStartAccount(receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartAccount(&_InfinityPool.TransactOpts, receiver, agentID, accountPrincipal)
}

// JumpStartAccount is a paid mutator transaction binding the contract method 0xf4765e9e.
//
// Solidity: function jumpStartAccount(address receiver, uint256 agentID, uint256 accountPrincipal) returns()
func (_InfinityPool *InfinityPoolTransactorSession) JumpStartAccount(receiver common.Address, agentID *big.Int, accountPrincipal *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartAccount(&_InfinityPool.TransactOpts, receiver, agentID, accountPrincipal)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPool *InfinityPoolTransactor) JumpStartTotalBorrowed(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "jumpStartTotalBorrowed", amount)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPool *InfinityPoolSession) JumpStartTotalBorrowed(amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartTotalBorrowed(&_InfinityPool.TransactOpts, amount)
}

// JumpStartTotalBorrowed is a paid mutator transaction binding the contract method 0x59221e38.
//
// Solidity: function jumpStartTotalBorrowed(uint256 amount) returns()
func (_InfinityPool *InfinityPoolTransactorSession) JumpStartTotalBorrowed(amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.JumpStartTotalBorrowed(&_InfinityPool.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactor) Mint(opts *bind.TransactOpts, shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "mint", shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPool *InfinityPoolSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Mint(&_InfinityPool.TransactOpts, shares, receiver)
}

// Mint is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 shares, address receiver) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactorSession) Mint(shares *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Mint(&_InfinityPool.TransactOpts, shares, receiver)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolTransactor) Pay(opts *bind.TransactOpts, vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "pay", vc)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolSession) Pay(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Pay(&_InfinityPool.TransactOpts, vc)
}

// Pay is a paid mutator transaction binding the contract method 0x7d694730.
//
// Solidity: function pay((address,uint256,uint256,uint256,uint256,bytes4,uint64,bytes) vc) returns(uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolTransactorSession) Pay(vc VerifiableCredential) (*types.Transaction, error) {
	return _InfinityPool.Contract.Pay(&_InfinityPool.TransactOpts, vc)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPool *InfinityPoolTransactor) RecoverERC20(opts *bind.TransactOpts, receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "recoverERC20", receiver, token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPool *InfinityPoolSession) RecoverERC20(receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverERC20(&_InfinityPool.TransactOpts, receiver, token)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x886f039a.
//
// Solidity: function recoverERC20(address receiver, address token) returns()
func (_InfinityPool *InfinityPoolTransactorSession) RecoverERC20(receiver common.Address, token common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverERC20(&_InfinityPool.TransactOpts, receiver, token)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPool *InfinityPoolTransactor) RecoverFIL(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "recoverFIL", receiver)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPool *InfinityPoolSession) RecoverFIL(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverFIL(&_InfinityPool.TransactOpts, receiver)
}

// RecoverFIL is a paid mutator transaction binding the contract method 0x7edc8fe7.
//
// Solidity: function recoverFIL(address receiver) returns()
func (_InfinityPool *InfinityPoolTransactorSession) RecoverFIL(receiver common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.RecoverFIL(&_InfinityPool.TransactOpts, receiver)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactor) Redeem(opts *bind.TransactOpts, shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "redeem", shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPool *InfinityPoolSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Redeem(&_InfinityPool.TransactOpts, shares, receiver, owner)
}

// Redeem is a paid mutator transaction binding the contract method 0xba087652.
//
// Solidity: function redeem(uint256 shares, address receiver, address owner) returns(uint256 assets)
func (_InfinityPool *InfinityPoolTransactorSession) Redeem(shares *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Redeem(&_InfinityPool.TransactOpts, shares, receiver, owner)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPool *InfinityPoolTransactor) RefreshRoutes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "refreshRoutes")
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPool *InfinityPoolSession) RefreshRoutes() (*types.Transaction, error) {
	return _InfinityPool.Contract.RefreshRoutes(&_InfinityPool.TransactOpts)
}

// RefreshRoutes is a paid mutator transaction binding the contract method 0x338891eb.
//
// Solidity: function refreshRoutes() returns()
func (_InfinityPool *InfinityPoolTransactorSession) RefreshRoutes() (*types.Transaction, error) {
	return _InfinityPool.Contract.RefreshRoutes(&_InfinityPool.TransactOpts)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_InfinityPool *InfinityPoolTransactor) SetMaxEpochsOwedTolerance(opts *bind.TransactOpts, _maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setMaxEpochsOwedTolerance", _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_InfinityPool *InfinityPoolSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMaxEpochsOwedTolerance(&_InfinityPool.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMaxEpochsOwedTolerance is a paid mutator transaction binding the contract method 0x4107e644.
//
// Solidity: function setMaxEpochsOwedTolerance(uint256 _maxEpochsOwedTolerance) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetMaxEpochsOwedTolerance(_maxEpochsOwedTolerance *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMaxEpochsOwedTolerance(&_InfinityPool.TransactOpts, _maxEpochsOwedTolerance)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPool *InfinityPoolTransactor) SetMinimumLiquidity(opts *bind.TransactOpts, _minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setMinimumLiquidity", _minimumLiquidity)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPool *InfinityPoolSession) SetMinimumLiquidity(_minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMinimumLiquidity(&_InfinityPool.TransactOpts, _minimumLiquidity)
}

// SetMinimumLiquidity is a paid mutator transaction binding the contract method 0x282567b4.
//
// Solidity: function setMinimumLiquidity(uint256 _minimumLiquidity) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetMinimumLiquidity(_minimumLiquidity *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetMinimumLiquidity(&_InfinityPool.TransactOpts, _minimumLiquidity)
}

// SetRamp is a paid mutator transaction binding the contract method 0x3d9c3c17.
//
// Solidity: function setRamp(address _ramp) returns()
func (_InfinityPool *InfinityPoolTransactor) SetRamp(opts *bind.TransactOpts, _ramp common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setRamp", _ramp)
}

// SetRamp is a paid mutator transaction binding the contract method 0x3d9c3c17.
//
// Solidity: function setRamp(address _ramp) returns()
func (_InfinityPool *InfinityPoolSession) SetRamp(_ramp common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRamp(&_InfinityPool.TransactOpts, _ramp)
}

// SetRamp is a paid mutator transaction binding the contract method 0x3d9c3c17.
//
// Solidity: function setRamp(address _ramp) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetRamp(_ramp common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRamp(&_InfinityPool.TransactOpts, _ramp)
}

// SetRateModule is a paid mutator transaction binding the contract method 0x31ff92fa.
//
// Solidity: function setRateModule(address _rateModule) returns()
func (_InfinityPool *InfinityPoolTransactor) SetRateModule(opts *bind.TransactOpts, _rateModule common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "setRateModule", _rateModule)
}

// SetRateModule is a paid mutator transaction binding the contract method 0x31ff92fa.
//
// Solidity: function setRateModule(address _rateModule) returns()
func (_InfinityPool *InfinityPoolSession) SetRateModule(_rateModule common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRateModule(&_InfinityPool.TransactOpts, _rateModule)
}

// SetRateModule is a paid mutator transaction binding the contract method 0x31ff92fa.
//
// Solidity: function setRateModule(address _rateModule) returns()
func (_InfinityPool *InfinityPoolTransactorSession) SetRateModule(_rateModule common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.SetRateModule(&_InfinityPool.TransactOpts, _rateModule)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPool *InfinityPoolTransactor) ShutDown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "shutDown")
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPool *InfinityPoolSession) ShutDown() (*types.Transaction, error) {
	return _InfinityPool.Contract.ShutDown(&_InfinityPool.TransactOpts)
}

// ShutDown is a paid mutator transaction binding the contract method 0x10b9e583.
//
// Solidity: function shutDown() returns()
func (_InfinityPool *InfinityPoolTransactorSession) ShutDown() (*types.Transaction, error) {
	return _InfinityPool.Contract.ShutDown(&_InfinityPool.TransactOpts)
}

// TransferFromPreStake is a paid mutator transaction binding the contract method 0x358bded0.
//
// Solidity: function transferFromPreStake(uint256 _amount) returns()
func (_InfinityPool *InfinityPoolTransactor) TransferFromPreStake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "transferFromPreStake", _amount)
}

// TransferFromPreStake is a paid mutator transaction binding the contract method 0x358bded0.
//
// Solidity: function transferFromPreStake(uint256 _amount) returns()
func (_InfinityPool *InfinityPoolSession) TransferFromPreStake(_amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferFromPreStake(&_InfinityPool.TransactOpts, _amount)
}

// TransferFromPreStake is a paid mutator transaction binding the contract method 0x358bded0.
//
// Solidity: function transferFromPreStake(uint256 _amount) returns()
func (_InfinityPool *InfinityPoolTransactorSession) TransferFromPreStake(_amount *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferFromPreStake(&_InfinityPool.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPool *InfinityPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPool *InfinityPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferOwnership(&_InfinityPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_InfinityPool *InfinityPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.TransferOwnership(&_InfinityPool.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactor) Withdraw(opts *bind.TransactOpts, assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "withdraw", assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPool *InfinityPoolSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Withdraw(&_InfinityPool.TransactOpts, assets, receiver, owner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xb460af94.
//
// Solidity: function withdraw(uint256 assets, address receiver, address owner) returns(uint256 shares)
func (_InfinityPool *InfinityPoolTransactorSession) Withdraw(assets *big.Int, receiver common.Address, owner common.Address) (*types.Transaction, error) {
	return _InfinityPool.Contract.Withdraw(&_InfinityPool.TransactOpts, assets, receiver, owner)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns(uint256 totalOwed)
func (_InfinityPool *InfinityPoolTransactor) WriteOff(opts *bind.TransactOpts, agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPool.contract.Transact(opts, "writeOff", agentID, recoveredFunds)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns(uint256 totalOwed)
func (_InfinityPool *InfinityPoolSession) WriteOff(agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.WriteOff(&_InfinityPool.TransactOpts, agentID, recoveredFunds)
}

// WriteOff is a paid mutator transaction binding the contract method 0xc564f772.
//
// Solidity: function writeOff(uint256 agentID, uint256 recoveredFunds) returns(uint256 totalOwed)
func (_InfinityPool *InfinityPoolTransactorSession) WriteOff(agentID *big.Int, recoveredFunds *big.Int) (*types.Transaction, error) {
	return _InfinityPool.Contract.WriteOff(&_InfinityPool.TransactOpts, agentID, recoveredFunds)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPool *InfinityPoolTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _InfinityPool.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPool *InfinityPoolSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InfinityPool.Contract.Fallback(&_InfinityPool.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_InfinityPool *InfinityPoolTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _InfinityPool.Contract.Fallback(&_InfinityPool.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPool *InfinityPoolTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InfinityPool.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPool *InfinityPoolSession) Receive() (*types.Transaction, error) {
	return _InfinityPool.Contract.Receive(&_InfinityPool.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_InfinityPool *InfinityPoolTransactorSession) Receive() (*types.Transaction, error) {
	return _InfinityPool.Contract.Receive(&_InfinityPool.TransactOpts)
}

// InfinityPoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the InfinityPool contract.
type InfinityPoolBorrowIterator struct {
	Event *InfinityPoolBorrow // Event containing the contract specifics and raw log

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
func (it *InfinityPoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolBorrow)
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
		it.Event = new(InfinityPoolBorrow)
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
func (it *InfinityPoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolBorrow represents a Borrow event raised by the InfinityPool contract.
type InfinityPoolBorrow struct {
	Agent  *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPool *InfinityPoolFilterer) FilterBorrow(opts *bind.FilterOpts, agent []*big.Int) (*InfinityPoolBorrowIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Borrow", agentRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolBorrowIterator{contract: _InfinityPool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPool *InfinityPoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *InfinityPoolBorrow, agent []*big.Int) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Borrow", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolBorrow)
				if err := _InfinityPool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x044ec4d36b0d76019c4b5aec2a216dbca5ad6c6d671940c8fbbcd23cfe4e804b.
//
// Solidity: event Borrow(uint256 indexed agent, uint256 amount)
func (_InfinityPool *InfinityPoolFilterer) ParseBorrow(log types.Log) (*InfinityPoolBorrow, error) {
	event := new(InfinityPoolBorrow)
	if err := _InfinityPool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the InfinityPool contract.
type InfinityPoolDepositIterator struct {
	Event *InfinityPoolDeposit // Event containing the contract specifics and raw log

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
func (it *InfinityPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolDeposit)
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
		it.Event = new(InfinityPoolDeposit)
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
func (it *InfinityPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolDeposit represents a Deposit event raised by the InfinityPool contract.
type InfinityPoolDeposit struct {
	Caller common.Address
	Owner  common.Address
	Assets *big.Int
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) FilterDeposit(opts *bind.FilterOpts, caller []common.Address, owner []common.Address) (*InfinityPoolDepositIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolDepositIterator{contract: _InfinityPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *InfinityPoolDeposit, caller []common.Address, owner []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Deposit", callerRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolDeposit)
				if err := _InfinityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed caller, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) ParseDeposit(log types.Log) (*InfinityPoolDeposit, error) {
	event := new(InfinityPoolDeposit)
	if err := _InfinityPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolPayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the InfinityPool contract.
type InfinityPoolPayIterator struct {
	Event *InfinityPoolPay // Event containing the contract specifics and raw log

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
func (it *InfinityPoolPayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolPay)
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
		it.Event = new(InfinityPoolPay)
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
func (it *InfinityPoolPayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolPayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolPay represents a Pay event raised by the InfinityPool contract.
type InfinityPoolPay struct {
	Agent         *big.Int
	Rate          *big.Int
	EpochsPaid    *big.Int
	PrincipalPaid *big.Int
	Refund        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolFilterer) FilterPay(opts *bind.FilterOpts, agent []*big.Int) (*InfinityPoolPayIterator, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Pay", agentRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolPayIterator{contract: _InfinityPool.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *InfinityPoolPay, agent []*big.Int) (event.Subscription, error) {

	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Pay", agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolPay)
				if err := _InfinityPool.contract.UnpackLog(event, "Pay", log); err != nil {
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

// ParsePay is a log parse operation binding the contract event 0xbf0d333142d8550a97625890e80f07dc447be1a140bf10f54fbc8da279668ff6.
//
// Solidity: event Pay(uint256 indexed agent, uint256 rate, uint256 epochsPaid, uint256 principalPaid, uint256 refund)
func (_InfinityPool *InfinityPoolFilterer) ParsePay(log types.Log) (*InfinityPoolPay, error) {
	event := new(InfinityPoolPay)
	if err := _InfinityPool.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the InfinityPool contract.
type InfinityPoolWithdrawIterator struct {
	Event *InfinityPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *InfinityPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolWithdraw)
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
		it.Event = new(InfinityPoolWithdraw)
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
func (it *InfinityPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolWithdraw represents a Withdraw event raised by the InfinityPool contract.
type InfinityPoolWithdraw struct {
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
func (_InfinityPool *InfinityPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, caller []common.Address, receiver []common.Address, owner []common.Address) (*InfinityPoolWithdrawIterator, error) {

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

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolWithdrawIterator{contract: _InfinityPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xfbde797d201c681b91056529119e0b02407c7bb96a4a2c75c01fc9667232c8db.
//
// Solidity: event Withdraw(address indexed caller, address indexed receiver, address indexed owner, uint256 assets, uint256 shares)
func (_InfinityPool *InfinityPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *InfinityPoolWithdraw, caller []common.Address, receiver []common.Address, owner []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "Withdraw", callerRule, receiverRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolWithdraw)
				if err := _InfinityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_InfinityPool *InfinityPoolFilterer) ParseWithdraw(log types.Log) (*InfinityPoolWithdraw, error) {
	event := new(InfinityPoolWithdraw)
	if err := _InfinityPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InfinityPoolWriteOffIterator is returned from FilterWriteOff and is used to iterate over the raw logs and unpacked data for WriteOff events raised by the InfinityPool contract.
type InfinityPoolWriteOffIterator struct {
	Event *InfinityPoolWriteOff // Event containing the contract specifics and raw log

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
func (it *InfinityPoolWriteOffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InfinityPoolWriteOff)
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
		it.Event = new(InfinityPoolWriteOff)
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
func (it *InfinityPoolWriteOffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InfinityPoolWriteOffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InfinityPoolWriteOff represents a WriteOff event raised by the InfinityPool contract.
type InfinityPoolWriteOff struct {
	AgentID        *big.Int
	RecoveredFunds *big.Int
	LostFunds      *big.Int
	InterestPaid   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWriteOff is a free log retrieval operation binding the contract event 0x31d6d43b5ab97f9739e3e78ab72e97d7575b3233f23e8e929160920032436aa5.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 interestPaid)
func (_InfinityPool *InfinityPoolFilterer) FilterWriteOff(opts *bind.FilterOpts, agentID []*big.Int) (*InfinityPoolWriteOffIterator, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _InfinityPool.contract.FilterLogs(opts, "WriteOff", agentIDRule)
	if err != nil {
		return nil, err
	}
	return &InfinityPoolWriteOffIterator{contract: _InfinityPool.contract, event: "WriteOff", logs: logs, sub: sub}, nil
}

// WatchWriteOff is a free log subscription operation binding the contract event 0x31d6d43b5ab97f9739e3e78ab72e97d7575b3233f23e8e929160920032436aa5.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 interestPaid)
func (_InfinityPool *InfinityPoolFilterer) WatchWriteOff(opts *bind.WatchOpts, sink chan<- *InfinityPoolWriteOff, agentID []*big.Int) (event.Subscription, error) {

	var agentIDRule []interface{}
	for _, agentIDItem := range agentID {
		agentIDRule = append(agentIDRule, agentIDItem)
	}

	logs, sub, err := _InfinityPool.contract.WatchLogs(opts, "WriteOff", agentIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InfinityPoolWriteOff)
				if err := _InfinityPool.contract.UnpackLog(event, "WriteOff", log); err != nil {
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

// ParseWriteOff is a log parse operation binding the contract event 0x31d6d43b5ab97f9739e3e78ab72e97d7575b3233f23e8e929160920032436aa5.
//
// Solidity: event WriteOff(uint256 indexed agentID, uint256 recoveredFunds, uint256 lostFunds, uint256 interestPaid)
func (_InfinityPool *InfinityPoolFilterer) ParseWriteOff(log types.Log) (*InfinityPoolWriteOff, error) {
	event := new(InfinityPoolWriteOff)
	if err := _InfinityPool.contract.UnpackLog(event, "WriteOff", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
