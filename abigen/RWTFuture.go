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

// Position is an auto generated low-level Go binding around an user-defined struct.
type Position struct {
	StrikePrice    *big.Int
	AllocatedFil   *big.Int
	ExpirationDate *big.Int
}

// RWTFutureMetaData contains all meta data concerning the RWTFuture contract.
var RWTFutureMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrLpPlus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"RWTNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"name\":\"RWTFutureMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burnExpired\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getPositionsBatched\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"internalType\":\"structPosition[]\",\"name\":\"positions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractILpPlus\",\"name\":\"_lpPlus\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpPlus\",\"outputs\":[{\"internalType\":\"contractILpPlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_strikePrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_allocatedFils\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_expirationDates\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILpPlus\",\"name\":\"_lpPlus\",\"type\":\"address\"}],\"name\":\"setLpPlus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f516020614b935f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051614acc90816100c78239608051818181611d6f0152611ef40152f35b6001600160401b0319166001600160401b039081175f516020614b935f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a71461323f5750806306fdde031461313e578063081812fc146130b9578063095ea7b314612eec57806318160ddd14612e925780631d05159d1461270957806323b872dd146126655780632f745c59146125935780632fc92888146125235780633c9774c5146124835780633f4ba83a1461238a57806342842e0e1461218a5780634f1ef28614611e8b5780634f6ccce714611de757806352d1902d14611d2a57806354fd4d5014611cf15780635c975abb14611c925780636352211e14611c3857806370a0823114611bec578063715018a614611aac57806374acd20914611a1a57806379ba5097146119795780638456cb59146118a65780638da5cb5b146118365780638f15b414146110115780638f68679f14610fb757806395d89b4114610e5f578063a22cb46514610d35578063ac9650d814610ae0578063ad3cb1cc14610a63578063b5a0749e14610949578063b80f55c914610889578063b88d4fde146105b0578063c87633fd1461048a578063c87b56dd1461040f578063e30c39781461039e578063e985e9c5146102e75763f2fde38b146101c0575f80fd5b346102e45760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102e45773ffffffffffffffffffffffffffffffffffffffff61021461020f6133a2565b613994565b61021c613aac565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b80fd5b50346102e45760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102e45761031f6133a2565b73ffffffffffffffffffffffffffffffffffffffff61038361033f6133c5565b9273ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b91165f52602052602060ff60405f2054166040519015158152f35b50346102e457807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102e457602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576104496004356138ad565b505f604051610459602082613404565b5261048260405161046b602082613404565b5f815260405191829160208352602083019061335f565b0390f35b5f80fd5b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865773ffffffffffffffffffffffffffffffffffffffff6104d66133a2565b6104de613aac565b1680156105885773ffffffffffffffffffffffffffffffffffffffff61050661058692613994565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416177f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0155565b005b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b346104865760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576105e76133a2565b6105ef6133c5565b6044359060643567ffffffffffffffff811161048657610622610619610628923690600401613566565b9461020f613941565b91613994565b610630613941565b61063982613994565b61064282613994565b73ffffffffffffffffffffffffffffffffffffffff81161561085d5773ffffffffffffffffffffffffffffffffffffffff6106808192873391614244565b921691169080820361082a575050803b61069657005b60209161070273ffffffffffffffffffffffffffffffffffffffff8093169560405195869485947f150b7a02000000000000000000000000000000000000000000000000000000008652336004870152166024850152604484015260806064840152608483019061335f565b03815f865af15f91816107cd575b50610757575061071e613c8e565b8051908161075257827f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b602001fd5b7fffffffff000000000000000000000000000000000000000000000000000000007f150b7a02000000000000000000000000000000000000000000000000000000009116036107a257005b7f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011610822575b816107e960209383613404565b8101031261048657517fffffffff0000000000000000000000000000000000000000000000000000000081168103610486579083610710565b3d91506107dc565b84907f64283d7b000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b7f64a0ae92000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865760043567ffffffffffffffff8111610486576108d890369060040161345d565b6108e0613941565b6108e86139fc565b80515f5b8181106108f557005b8061094361090560019386613639565b51805f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020525f60026040822082815582878201550155613cbd565b016108ec565b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865760043567ffffffffffffffff8111610486576109989036906004016135aa565b906109a1613941565b5f5b82811015610586578060051b820135805f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020524260406109e7815f206136a5565b015111610a385790610a32826001935f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020525f60026040822082815582878201550155613cbd565b016109a3565b7ff7b20ee7000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657610482604051610aa2604082613404565b600581527f352e302e30000000000000000000000000000000000000000000000000000000602082015260405191829160208352602083019061335f565b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865760043567ffffffffffffffff811161048657610b2f9036906004016135aa565b602091604051610b3f8482613404565b5f8152838101917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08501368437610b7584613445565b90610b836040519283613404565b8482527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0610bb086613445565b01865f5b828110610d26575050505f907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe181360301915b86811015610cab578060051b820135838112156104865782019081359167ffffffffffffffff8311610486578901918036038313610486575f8060019489610c7b610c8f958f8e6040519483869484860198893784019083820190898252519283915e0101858152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08101835282613404565b5190305af4610c88613c8e565b90306149fd565b610c998287613639565b52610ca48186613639565b5001610be7565b604080518981528551818b018190525f92600582901b8301810191888d01918d9085015b828710610cdc5785850386f35b909192938280610d16837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a60019603018652885161335f565b9601920196019592919092610ccf565b60608582018301528101610bb4565b346104865760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657610d6c6133a2565b6024359081151580920361048657610d9b73ffffffffffffffffffffffffffffffffffffffff9161020f613941565b16908115610e3357335f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260409020825f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff83161790556040519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b507f5b08ba18000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154610ebc816135db565b8084529060018116908115610f755750600114610ef8575b61048283610ee481850382613404565b60405191829160208352602083019061335f565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f9081527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b808210610f5b57509091508101602001610ee4610ed4565b919260018160209254838588010152019101909291610f43565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b84019091019150610ee49050610ed4565b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865760207f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054604051908152f35b346104865760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865760043567ffffffffffffffff811161048657611060903690600401613566565b60243567ffffffffffffffff811161048657611080903690600401613566565b9060443573ffffffffffffffffffffffffffffffffffffffff81168103610486576064359073ffffffffffffffffffffffffffffffffffffffff8216809203610486577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549260ff8460401c16159367ffffffffffffffff81168015908161182e575b6001149081611824575b15908161181b575b506117f3578460017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005561179e575b506111716149a6565b6111796149a6565b6111816149a6565b80519067ffffffffffffffff821161159f5781906111bf7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300546135db565b601f81116116d4575b50602090601f83116001146115d7575f926115cc575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b835167ffffffffffffffff811161159f5761126b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301546135db565b601f811161151d575b506020601f82116001146114395781906112eb9495965f9261142e575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b61020f6149a6565b906112f46149a6565b6112fc6149a6565b73ffffffffffffffffffffffffffffffffffffffff8216156114025761050673ffffffffffffffffffffffffffffffffffffffff9161133d61134594613b7f565b6112e36149a6565b60017f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a005561136f57005b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b7f1e4fbdf7000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b015190508680611291565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216957f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f52815f20965f5b8181106115055750916112eb959697918460019594106114ce575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301556112e3565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558680806114a1565b83830151895560019098019760209384019301611486565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f830160051c81019160208410611595575b601f0160051c01905b81811061158a5750611274565b5f815560010161157d565b9091508190611574565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b0151905086806111de565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81925f5b8181106116bc5750908460019594939210611685575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055611230565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055868080611658565b92936020600181928786015181550195019301611642565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f52601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410611776575b90601f8493920160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b81811061176857506111c8565b5f815584935060010161175b565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81915061172d565b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005585611168565b7ff92ee8a9000000000000000000000000000000000000000000000000000000005f5260045ffd5b90501587611115565b303b15915061110d565b869150611103565b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576118dc613aac565b6118e4613941565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416036119ee5761058633613b7f565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657611a51613687565b506004355f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020526060611a8960405f206136a5565b611aaa60405180926040809180518452602081015160208501520151910152565bf35b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657611ae2613aac565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00555f73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576020611c30611c2b61020f6133a2565b613aec565b604051908152f35b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576020611c746004356138ad565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657602060405160018152f35b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003611dbf5760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576004357f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254811015611e5c57611e4d602091613857565b90549060031b1c604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f525f60045260245260445ffd5b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657611ebd6133a2565b60243567ffffffffffffffff811161048657611edd903690600401613566565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115612148575b50611dbf57611f2c613aac565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181612114575b50611fac57837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8592036120e95750813b156120be57807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a281511561208d575f8083602061058695519101845af4612087613c8e565b916149fd565b50503461209657005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011612140575b8161213060209383613404565b8101031261048657519085611f7b565b3d9150612123565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141583611f1f565b3461048657612198366134ba565b6020916121ba610622604051956121af8688613404565b5f875261020f613941565b6121c2613941565b6121cb82613994565b6121d482613994565b73ffffffffffffffffffffffffffffffffffffffff81161561085d5773ffffffffffffffffffffffffffffffffffffffff6122128192873391614244565b921691169080820361082a575050803b61222857005b839161229373ffffffffffffffffffffffffffffffffffffffff8093169660405195869485947f150b7a02000000000000000000000000000000000000000000000000000000008652336004870152166024850152604484015260806064840152608483019061335f565b03815f875af15f9181612332575b506122e657506122af613c8e565b805191826122e357837f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b01fd5b7f150b7a020000000000000000000000000000000000000000000000000000000091507fffffffff0000000000000000000000000000000000000000000000000000000016036107a257005b9091508281813d8311612383575b61234a8183613404565b8101031261048657517fffffffff00000000000000000000000000000000000000000000000000000000811681036104865790846122a1565b503d612340565b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576123c0613aac565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff81161561245b577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b346104865760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576124c06024356004356136cd565b6040518091602082016020835281518091526020604084019201905f5b8181106124eb575050500390f35b91935091602060608261251560019488516040809180518452602081015160208501520151910152565b0194019101918493926124dd565b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657602073ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416604051908152f35b346104865760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576125ca6133a2565b73ffffffffffffffffffffffffffffffffffffffff602435916125ef611c2b82613994565b83101561263657165f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20905f52602052602060405f2054604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f521660045260245260445ffd5b346104865761268661268c612679366134ba565b929161020f949194613941565b92613994565b73ffffffffffffffffffffffffffffffffffffffff81161561085d5773ffffffffffffffffffffffffffffffffffffffff6126ca8192843391614244565b93169216918083036126d857005b7f64283d7b000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b346104865760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576127406133a2565b60243567ffffffffffffffff81116104865761276090369060040161345d565b9060443567ffffffffffffffff81116104865761278190369060040161345d565b60643567ffffffffffffffff811161048657906127a38492369060040161345d565b6127ab613941565b6127b36139fc565b8251916127bf83613445565b936127cd6040519586613404565b8385526127d984613445565b957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020870197013688377f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054935f9273ffffffffffffffffffffffffffffffffffffffff8316928315908115915b8987106128bf578b8b61285b8c8c61362c565b7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0055604051918291602083019060208452518091526040830191905f5b8181106128a6575050500390f35b8251845285945060209384019390920191600101612898565b869084612957838e9f9e8c60028f6128fa856128f38f936128ec6128e484809361362c565b9d8e9c613639565b5196613639565b5192613639565b519060405193612909856133e8565b84526020840190815260408401918252875f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0260205260405f209351845551600184015551910155613639565b5261085d576129a4815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b968373ffffffffffffffffffffffffffffffffffffffff891698891580159283612dc8575b612d75575b845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f208a7fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055848a8c7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a415612c79577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060405f20556801000000000000000081101561159f5784612ae0826001612b1694017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255613857565b9091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83549160031b92831b921b1916179055565b885f9a03612bae575b50612b8257866001927f03488890ace1d90f3afce239b92847940fe64099a8f090c57b2af9dd9f25ba3260608c612b5a866128f3818d613639565b51612b65878d613639565b519060405192835260208301526040820152a301959a999a612848565b7f73c6ac6e000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b612bba611c2b86613994565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111612c4c57895f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20815f526020528460405f2055845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f2055612b1f565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b898914612b1657612c8c611c2b82613994565b845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f2054908b5f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f2091818103612d2b575b50855f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020525f60408120555f526020525f6040812055612b16565b815f528260205260405f2054815f52836020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f20555f612ced565b612dbc8673ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190556129ce565b50612e1f855f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b86612e678373ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190556129c9565b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126104865760207f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254604051908152f35b346104865760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657612f236133a2565b612f326024359161020f613941565b90612f3c816138ad565b33151580613099575b80613027575b612ffb57819073ffffffffffffffffffffffffffffffffffffffff80851691167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9255f80a45f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f2091167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790555f80f35b7fa9fbf51f000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b5061306f8173ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b73ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f20541615612f4b565b503373ffffffffffffffffffffffffffffffffffffffff82161415612f45565b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576020611c746004356130f9816138ad565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b34610486575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc360112610486576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005461319b816135db565b8084529060018116908115610f7557506001146131c25761048283610ee481850382613404565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b80821061322557509091508101602001610ee4610ed4565b91926001816020925483858801015201910190929161320d565b346104865760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261048657600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361048657817f780e9d6300000000000000000000000000000000000000000000000000000000602093149081156132d1575b5015158152f35b7f80ac58cd00000000000000000000000000000000000000000000000000000000811491508115613335575b811561330b575b50836132ca565b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483613304565b7f5b5e139f00000000000000000000000000000000000000000000000000000000811491506132fd565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361048657565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361048657565b6060810190811067ffffffffffffffff82111761159f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761159f57604052565b67ffffffffffffffff811161159f5760051b60200190565b9080601f8301121561048657813561347481613445565b926134826040519485613404565b81845260208085019260051b82010192831161048657602001905b8282106134aa5750505090565b813581526020918201910161349d565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60609101126104865760043573ffffffffffffffffffffffffffffffffffffffff81168103610486579060243573ffffffffffffffffffffffffffffffffffffffff81168103610486579060443590565b67ffffffffffffffff811161159f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f82011215610486576020813591016135808261352c565b9261358e6040519485613404565b8284528282011161048657815f92602092838601378301015290565b9181601f840112156104865782359167ffffffffffffffff8311610486576020808501948460051b01011161048657565b90600182811c92168015613622575b60208310146135f557565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f16916135ea565b91908201809211612c4c57565b805182101561364d5760209160051b010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b91908203918211612c4c57565b60405190613694826133e8565b5f6040838281528260208201520152565b906040516136b2816133e8565b60406002829480548452600181015460208501520154910152565b91907f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201828111612c4c57841161381b5761372c908461362c565b90808211613813575b50613740838261367a565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061378461376e86613445565b9561377c6040519788613404565b808752613445565b015f5b8181106137fc57505083815b8381106137a05750505050565b806001915f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020526137f560405f206137e46137de878561367a565b916136a5565b6137ee8287613639565b5284613639565b5001613793565b602090613807613687565b82828901015201613787565b90505f613735565b509091505060405161382e602082613404565b5f81525f805b81811061384057505090565b60209061384b613687565b82828601015201613834565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025481101561364d577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025f5260205f2001905f90565b6138f5815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff821615613916575090565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300541661396c57565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff00000000000000000000000000000000000000146139e8575b156139e4576139d6906141cb565b906139df575090565b905090565b5090565b505067ffffffffffffffff811660016139c8565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633148015613a6c575b15613a4457565b7fa696e8c1000000000000000000000000000000000000000000000000000000005f5260045ffd5b5073ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0154163314613a3d565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633036119ee57565b73ffffffffffffffffffffffffffffffffffffffff811615613b5357613b4f9073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b5490565b7f89c62b64000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b3d15613cb8573d90613c9f8261352c565b91613cad6040519384613404565b82523d5f602084013e565b606090565b613d05815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b73ffffffffffffffffffffffffffffffffffffffff811690811591821580614103575b845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055845f837fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8280a483156140035750507f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254835f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060405f20556801000000000000000081101561159f5783612ae0826001613e3b94017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255613857565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111612c4c57835f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed03602052613ebc60405f205491613857565b90549060031b1c613ed081612ae084613857565b5f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205260405f2055825f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020525f60408120557f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02548015613fd6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01613fac613f7e82613857565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255506139165750565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b61400e575b50613e3b565b61401a611c2b83613994565b845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f2054915f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20918181036140b9575b50845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020525f60408120555f526020525f60408120555f614008565b815f528260205260405f2054815f52836020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f20555f61407a565b614159855f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b6141a08373ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8154019055613d28565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff1603614239575b8215801561422e575b61422657565b5f9250829150565b5060163d1415614220565b5f9250829150614217565b9073ffffffffffffffffffffffffffffffffffffffff6142a2825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b931680151580614839575b505073ffffffffffffffffffffffffffffffffffffffff8316918215928315614771575b73ffffffffffffffffffffffffffffffffffffffff821693841590811561471e575b845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f20867fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790558486847fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a415614622577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060405f20556801000000000000000081101561159f5784612ae082600161440594017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255613857565b1561457557505090507f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111612c4c57815f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205261448f60405f205491613857565b90549060031b1c6144a381612ae084613857565b5f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205260405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020525f60408120557f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02548015613fd6577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01614550613f7e82613857565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025590565b8303614582575b50505090565b611c2b61458e91613994565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301928311612c4c575f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20825f526020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f20555f808061457c565b81851461440557614635611c2b87613994565b845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f205490835f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20918181036146d4575b50855f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020525f60408120555f526020525f6040812055614405565b815f528260205260405f2054815f52836020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f20555f614696565b6147658473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190556142f3565b6147c7835f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b61480e8573ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190556142d1565b806148bf575b1561484a57806142ad565b73ffffffffffffffffffffffffffffffffffffffff841661489157507f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7f177e802f000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b508073ffffffffffffffffffffffffffffffffffffffff851614801561494b575b8061483f57508073ffffffffffffffffffffffffffffffffffffffff614944845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b161461483f565b506149938473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b815f5260205260ff60405f2054166148e0565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c16156149d557565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b90614a3a5750805115614a1257805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580614a8d575b614a4b575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15614a4356fea26469706673582212209a137d10dbb15b946e008d57d8f93dd7cb6613a377bdfb6a1fc3253ae0a7d00764736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// RWTFutureABI is the input ABI used to generate the binding from.
// Deprecated: Use RWTFutureMetaData.ABI instead.
var RWTFutureABI = RWTFutureMetaData.ABI

// RWTFutureBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RWTFutureMetaData.Bin instead.
var RWTFutureBin = RWTFutureMetaData.Bin

// DeployRWTFuture deploys a new Ethereum contract, binding an instance of RWTFuture to it.
func DeployRWTFuture(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RWTFuture, error) {
	parsed, err := RWTFutureMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RWTFutureBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RWTFuture{RWTFutureCaller: RWTFutureCaller{contract: contract}, RWTFutureTransactor: RWTFutureTransactor{contract: contract}, RWTFutureFilterer: RWTFutureFilterer{contract: contract}}, nil
}

// RWTFuture is an auto generated Go binding around an Ethereum contract.
type RWTFuture struct {
	RWTFutureCaller     // Read-only binding to the contract
	RWTFutureTransactor // Write-only binding to the contract
	RWTFutureFilterer   // Log filterer for contract events
}

// RWTFutureCaller is an auto generated read-only Go binding around an Ethereum contract.
type RWTFutureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWTFutureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RWTFutureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWTFutureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RWTFutureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RWTFutureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RWTFutureSession struct {
	Contract     *RWTFuture        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RWTFutureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RWTFutureCallerSession struct {
	Contract *RWTFutureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RWTFutureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RWTFutureTransactorSession struct {
	Contract     *RWTFutureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RWTFutureRaw is an auto generated low-level Go binding around an Ethereum contract.
type RWTFutureRaw struct {
	Contract *RWTFuture // Generic contract binding to access the raw methods on
}

// RWTFutureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RWTFutureCallerRaw struct {
	Contract *RWTFutureCaller // Generic read-only contract binding to access the raw methods on
}

// RWTFutureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RWTFutureTransactorRaw struct {
	Contract *RWTFutureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRWTFuture creates a new instance of RWTFuture, bound to a specific deployed contract.
func NewRWTFuture(address common.Address, backend bind.ContractBackend) (*RWTFuture, error) {
	contract, err := bindRWTFuture(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RWTFuture{RWTFutureCaller: RWTFutureCaller{contract: contract}, RWTFutureTransactor: RWTFutureTransactor{contract: contract}, RWTFutureFilterer: RWTFutureFilterer{contract: contract}}, nil
}

// NewRWTFutureCaller creates a new read-only instance of RWTFuture, bound to a specific deployed contract.
func NewRWTFutureCaller(address common.Address, caller bind.ContractCaller) (*RWTFutureCaller, error) {
	contract, err := bindRWTFuture(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RWTFutureCaller{contract: contract}, nil
}

// NewRWTFutureTransactor creates a new write-only instance of RWTFuture, bound to a specific deployed contract.
func NewRWTFutureTransactor(address common.Address, transactor bind.ContractTransactor) (*RWTFutureTransactor, error) {
	contract, err := bindRWTFuture(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RWTFutureTransactor{contract: contract}, nil
}

// NewRWTFutureFilterer creates a new log filterer instance of RWTFuture, bound to a specific deployed contract.
func NewRWTFutureFilterer(address common.Address, filterer bind.ContractFilterer) (*RWTFutureFilterer, error) {
	contract, err := bindRWTFuture(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RWTFutureFilterer{contract: contract}, nil
}

// bindRWTFuture binds a generic wrapper to an already deployed contract.
func bindRWTFuture(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RWTFutureMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWTFuture *RWTFutureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWTFuture.Contract.RWTFutureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWTFuture *RWTFutureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWTFuture.Contract.RWTFutureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWTFuture *RWTFutureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWTFuture.Contract.RWTFutureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RWTFuture *RWTFutureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RWTFuture.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RWTFuture *RWTFutureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWTFuture.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RWTFuture *RWTFutureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RWTFuture.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RWTFuture *RWTFutureCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RWTFuture *RWTFutureSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RWTFuture.Contract.UPGRADEINTERFACEVERSION(&_RWTFuture.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_RWTFuture *RWTFutureCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _RWTFuture.Contract.UPGRADEINTERFACEVERSION(&_RWTFuture.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RWTFuture *RWTFutureCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RWTFuture *RWTFutureSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RWTFuture.Contract.BalanceOf(&_RWTFuture.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RWTFuture *RWTFutureCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RWTFuture.Contract.BalanceOf(&_RWTFuture.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RWTFuture *RWTFutureCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RWTFuture *RWTFutureSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RWTFuture.Contract.GetApproved(&_RWTFuture.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RWTFuture *RWTFutureCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RWTFuture.Contract.GetApproved(&_RWTFuture.CallOpts, tokenId)
}

// GetPositionsBatched is a free data retrieval call binding the contract method 0x3c9774c5.
//
// Solidity: function getPositionsBatched(uint256 _startTokenId, uint256 _limit) view returns((uint256,uint256,uint256)[] positions)
func (_RWTFuture *RWTFutureCaller) GetPositionsBatched(opts *bind.CallOpts, _startTokenId *big.Int, _limit *big.Int) ([]Position, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "getPositionsBatched", _startTokenId, _limit)

	if err != nil {
		return *new([]Position), err
	}

	out0 := *abi.ConvertType(out[0], new([]Position)).(*[]Position)

	return out0, err

}

// GetPositionsBatched is a free data retrieval call binding the contract method 0x3c9774c5.
//
// Solidity: function getPositionsBatched(uint256 _startTokenId, uint256 _limit) view returns((uint256,uint256,uint256)[] positions)
func (_RWTFuture *RWTFutureSession) GetPositionsBatched(_startTokenId *big.Int, _limit *big.Int) ([]Position, error) {
	return _RWTFuture.Contract.GetPositionsBatched(&_RWTFuture.CallOpts, _startTokenId, _limit)
}

// GetPositionsBatched is a free data retrieval call binding the contract method 0x3c9774c5.
//
// Solidity: function getPositionsBatched(uint256 _startTokenId, uint256 _limit) view returns((uint256,uint256,uint256)[] positions)
func (_RWTFuture *RWTFutureCallerSession) GetPositionsBatched(_startTokenId *big.Int, _limit *big.Int) ([]Position, error) {
	return _RWTFuture.Contract.GetPositionsBatched(&_RWTFuture.CallOpts, _startTokenId, _limit)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RWTFuture *RWTFutureCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RWTFuture *RWTFutureSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RWTFuture.Contract.IsApprovedForAll(&_RWTFuture.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RWTFuture *RWTFutureCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RWTFuture.Contract.IsApprovedForAll(&_RWTFuture.CallOpts, owner, operator)
}

// LpPlus is a free data retrieval call binding the contract method 0x2fc92888.
//
// Solidity: function lpPlus() view returns(address)
func (_RWTFuture *RWTFutureCaller) LpPlus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "lpPlus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LpPlus is a free data retrieval call binding the contract method 0x2fc92888.
//
// Solidity: function lpPlus() view returns(address)
func (_RWTFuture *RWTFutureSession) LpPlus() (common.Address, error) {
	return _RWTFuture.Contract.LpPlus(&_RWTFuture.CallOpts)
}

// LpPlus is a free data retrieval call binding the contract method 0x2fc92888.
//
// Solidity: function lpPlus() view returns(address)
func (_RWTFuture *RWTFutureCallerSession) LpPlus() (common.Address, error) {
	return _RWTFuture.Contract.LpPlus(&_RWTFuture.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RWTFuture *RWTFutureCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RWTFuture *RWTFutureSession) Name() (string, error) {
	return _RWTFuture.Contract.Name(&_RWTFuture.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RWTFuture *RWTFutureCallerSession) Name() (string, error) {
	return _RWTFuture.Contract.Name(&_RWTFuture.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWTFuture *RWTFutureCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWTFuture *RWTFutureSession) Owner() (common.Address, error) {
	return _RWTFuture.Contract.Owner(&_RWTFuture.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RWTFuture *RWTFutureCallerSession) Owner() (common.Address, error) {
	return _RWTFuture.Contract.Owner(&_RWTFuture.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RWTFuture *RWTFutureCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RWTFuture *RWTFutureSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RWTFuture.Contract.OwnerOf(&_RWTFuture.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RWTFuture *RWTFutureCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RWTFuture.Contract.OwnerOf(&_RWTFuture.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RWTFuture *RWTFutureCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RWTFuture *RWTFutureSession) Paused() (bool, error) {
	return _RWTFuture.Contract.Paused(&_RWTFuture.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RWTFuture *RWTFutureCallerSession) Paused() (bool, error) {
	return _RWTFuture.Contract.Paused(&_RWTFuture.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RWTFuture *RWTFutureCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RWTFuture *RWTFutureSession) PendingOwner() (common.Address, error) {
	return _RWTFuture.Contract.PendingOwner(&_RWTFuture.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_RWTFuture *RWTFutureCallerSession) PendingOwner() (common.Address, error) {
	return _RWTFuture.Contract.PendingOwner(&_RWTFuture.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RWTFuture *RWTFutureCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RWTFuture *RWTFutureSession) ProxiableUUID() ([32]byte, error) {
	return _RWTFuture.Contract.ProxiableUUID(&_RWTFuture.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RWTFuture *RWTFutureCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RWTFuture.Contract.ProxiableUUID(&_RWTFuture.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RWTFuture *RWTFutureCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RWTFuture *RWTFutureSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RWTFuture.Contract.SupportsInterface(&_RWTFuture.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RWTFuture *RWTFutureCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RWTFuture.Contract.SupportsInterface(&_RWTFuture.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RWTFuture *RWTFutureCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RWTFuture *RWTFutureSession) Symbol() (string, error) {
	return _RWTFuture.Contract.Symbol(&_RWTFuture.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RWTFuture *RWTFutureCallerSession) Symbol() (string, error) {
	return _RWTFuture.Contract.Symbol(&_RWTFuture.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RWTFuture *RWTFutureCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RWTFuture *RWTFutureSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RWTFuture.Contract.TokenByIndex(&_RWTFuture.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RWTFuture *RWTFutureCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RWTFuture.Contract.TokenByIndex(&_RWTFuture.CallOpts, index)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_RWTFuture *RWTFutureCaller) TokenIdGenerator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "tokenIdGenerator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_RWTFuture *RWTFutureSession) TokenIdGenerator() (*big.Int, error) {
	return _RWTFuture.Contract.TokenIdGenerator(&_RWTFuture.CallOpts)
}

// TokenIdGenerator is a free data retrieval call binding the contract method 0x8f68679f.
//
// Solidity: function tokenIdGenerator() view returns(uint256)
func (_RWTFuture *RWTFutureCallerSession) TokenIdGenerator() (*big.Int, error) {
	return _RWTFuture.Contract.TokenIdGenerator(&_RWTFuture.CallOpts)
}

// TokenIdToPosition is a free data retrieval call binding the contract method 0x74acd209.
//
// Solidity: function tokenIdToPosition(uint256 _tokenId) view returns((uint256,uint256,uint256))
func (_RWTFuture *RWTFutureCaller) TokenIdToPosition(opts *bind.CallOpts, _tokenId *big.Int) (Position, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "tokenIdToPosition", _tokenId)

	if err != nil {
		return *new(Position), err
	}

	out0 := *abi.ConvertType(out[0], new(Position)).(*Position)

	return out0, err

}

// TokenIdToPosition is a free data retrieval call binding the contract method 0x74acd209.
//
// Solidity: function tokenIdToPosition(uint256 _tokenId) view returns((uint256,uint256,uint256))
func (_RWTFuture *RWTFutureSession) TokenIdToPosition(_tokenId *big.Int) (Position, error) {
	return _RWTFuture.Contract.TokenIdToPosition(&_RWTFuture.CallOpts, _tokenId)
}

// TokenIdToPosition is a free data retrieval call binding the contract method 0x74acd209.
//
// Solidity: function tokenIdToPosition(uint256 _tokenId) view returns((uint256,uint256,uint256))
func (_RWTFuture *RWTFutureCallerSession) TokenIdToPosition(_tokenId *big.Int) (Position, error) {
	return _RWTFuture.Contract.TokenIdToPosition(&_RWTFuture.CallOpts, _tokenId)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RWTFuture *RWTFutureCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RWTFuture *RWTFutureSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RWTFuture.Contract.TokenOfOwnerByIndex(&_RWTFuture.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RWTFuture *RWTFutureCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RWTFuture.Contract.TokenOfOwnerByIndex(&_RWTFuture.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RWTFuture *RWTFutureCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RWTFuture *RWTFutureSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RWTFuture.Contract.TokenURI(&_RWTFuture.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RWTFuture *RWTFutureCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RWTFuture.Contract.TokenURI(&_RWTFuture.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RWTFuture *RWTFutureCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RWTFuture *RWTFutureSession) TotalSupply() (*big.Int, error) {
	return _RWTFuture.Contract.TotalSupply(&_RWTFuture.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RWTFuture *RWTFutureCallerSession) TotalSupply() (*big.Int, error) {
	return _RWTFuture.Contract.TotalSupply(&_RWTFuture.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_RWTFuture *RWTFutureCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RWTFuture.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_RWTFuture *RWTFutureSession) Version() (*big.Int, error) {
	return _RWTFuture.Contract.Version(&_RWTFuture.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_RWTFuture *RWTFutureCallerSession) Version() (*big.Int, error) {
	return _RWTFuture.Contract.Version(&_RWTFuture.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RWTFuture *RWTFutureTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RWTFuture *RWTFutureSession) AcceptOwnership() (*types.Transaction, error) {
	return _RWTFuture.Contract.AcceptOwnership(&_RWTFuture.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_RWTFuture *RWTFutureTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _RWTFuture.Contract.AcceptOwnership(&_RWTFuture.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_RWTFuture *RWTFutureTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_RWTFuture *RWTFutureSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.Approve(&_RWTFuture.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_RWTFuture *RWTFutureTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.Approve(&_RWTFuture.TransactOpts, _to, _tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xb80f55c9.
//
// Solidity: function burn(uint256[] _tokenIds) returns()
func (_RWTFuture *RWTFutureTransactor) Burn(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "burn", _tokenIds)
}

// Burn is a paid mutator transaction binding the contract method 0xb80f55c9.
//
// Solidity: function burn(uint256[] _tokenIds) returns()
func (_RWTFuture *RWTFutureSession) Burn(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.Burn(&_RWTFuture.TransactOpts, _tokenIds)
}

// Burn is a paid mutator transaction binding the contract method 0xb80f55c9.
//
// Solidity: function burn(uint256[] _tokenIds) returns()
func (_RWTFuture *RWTFutureTransactorSession) Burn(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.Burn(&_RWTFuture.TransactOpts, _tokenIds)
}

// BurnExpired is a paid mutator transaction binding the contract method 0xb5a0749e.
//
// Solidity: function burnExpired(uint256[] _tokenIds) returns()
func (_RWTFuture *RWTFutureTransactor) BurnExpired(opts *bind.TransactOpts, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "burnExpired", _tokenIds)
}

// BurnExpired is a paid mutator transaction binding the contract method 0xb5a0749e.
//
// Solidity: function burnExpired(uint256[] _tokenIds) returns()
func (_RWTFuture *RWTFutureSession) BurnExpired(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.BurnExpired(&_RWTFuture.TransactOpts, _tokenIds)
}

// BurnExpired is a paid mutator transaction binding the contract method 0xb5a0749e.
//
// Solidity: function burnExpired(uint256[] _tokenIds) returns()
func (_RWTFuture *RWTFutureTransactorSession) BurnExpired(_tokenIds []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.BurnExpired(&_RWTFuture.TransactOpts, _tokenIds)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, address _lpPlus) returns()
func (_RWTFuture *RWTFutureTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _initialOwner common.Address, _lpPlus common.Address) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "initialize", _name, _symbol, _initialOwner, _lpPlus)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, address _lpPlus) returns()
func (_RWTFuture *RWTFutureSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _lpPlus common.Address) (*types.Transaction, error) {
	return _RWTFuture.Contract.Initialize(&_RWTFuture.TransactOpts, _name, _symbol, _initialOwner, _lpPlus)
}

// Initialize is a paid mutator transaction binding the contract method 0x8f15b414.
//
// Solidity: function initialize(string _name, string _symbol, address _initialOwner, address _lpPlus) returns()
func (_RWTFuture *RWTFutureTransactorSession) Initialize(_name string, _symbol string, _initialOwner common.Address, _lpPlus common.Address) (*types.Transaction, error) {
	return _RWTFuture.Contract.Initialize(&_RWTFuture.TransactOpts, _name, _symbol, _initialOwner, _lpPlus)
}

// Mint is a paid mutator transaction binding the contract method 0x1d05159d.
//
// Solidity: function mint(address _to, uint256[] _strikePrices, uint256[] _allocatedFils, uint256[] _expirationDates) returns(uint256[] tokenIds)
func (_RWTFuture *RWTFutureTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _strikePrices []*big.Int, _allocatedFils []*big.Int, _expirationDates []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "mint", _to, _strikePrices, _allocatedFils, _expirationDates)
}

// Mint is a paid mutator transaction binding the contract method 0x1d05159d.
//
// Solidity: function mint(address _to, uint256[] _strikePrices, uint256[] _allocatedFils, uint256[] _expirationDates) returns(uint256[] tokenIds)
func (_RWTFuture *RWTFutureSession) Mint(_to common.Address, _strikePrices []*big.Int, _allocatedFils []*big.Int, _expirationDates []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.Mint(&_RWTFuture.TransactOpts, _to, _strikePrices, _allocatedFils, _expirationDates)
}

// Mint is a paid mutator transaction binding the contract method 0x1d05159d.
//
// Solidity: function mint(address _to, uint256[] _strikePrices, uint256[] _allocatedFils, uint256[] _expirationDates) returns(uint256[] tokenIds)
func (_RWTFuture *RWTFutureTransactorSession) Mint(_to common.Address, _strikePrices []*big.Int, _allocatedFils []*big.Int, _expirationDates []*big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.Mint(&_RWTFuture.TransactOpts, _to, _strikePrices, _allocatedFils, _expirationDates)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RWTFuture *RWTFutureTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RWTFuture *RWTFutureSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _RWTFuture.Contract.Multicall(&_RWTFuture.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_RWTFuture *RWTFutureTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _RWTFuture.Contract.Multicall(&_RWTFuture.TransactOpts, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RWTFuture *RWTFutureTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RWTFuture *RWTFutureSession) Pause() (*types.Transaction, error) {
	return _RWTFuture.Contract.Pause(&_RWTFuture.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RWTFuture *RWTFutureTransactorSession) Pause() (*types.Transaction, error) {
	return _RWTFuture.Contract.Pause(&_RWTFuture.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RWTFuture *RWTFutureTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RWTFuture *RWTFutureSession) RenounceOwnership() (*types.Transaction, error) {
	return _RWTFuture.Contract.RenounceOwnership(&_RWTFuture.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RWTFuture *RWTFutureTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RWTFuture.Contract.RenounceOwnership(&_RWTFuture.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RWTFuture *RWTFutureTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RWTFuture *RWTFutureSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.SafeTransferFrom(&_RWTFuture.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RWTFuture *RWTFutureTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.SafeTransferFrom(&_RWTFuture.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_RWTFuture *RWTFutureTransactor) SafeTransferFrom0(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "safeTransferFrom0", _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_RWTFuture *RWTFutureSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RWTFuture.Contract.SafeTransferFrom0(&_RWTFuture.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes _data) returns()
func (_RWTFuture *RWTFutureTransactorSession) SafeTransferFrom0(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RWTFuture.Contract.SafeTransferFrom0(&_RWTFuture.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_RWTFuture *RWTFutureTransactor) SetApprovalForAll(opts *bind.TransactOpts, _operator common.Address, _approved bool) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "setApprovalForAll", _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_RWTFuture *RWTFutureSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _RWTFuture.Contract.SetApprovalForAll(&_RWTFuture.TransactOpts, _operator, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address _operator, bool _approved) returns()
func (_RWTFuture *RWTFutureTransactorSession) SetApprovalForAll(_operator common.Address, _approved bool) (*types.Transaction, error) {
	return _RWTFuture.Contract.SetApprovalForAll(&_RWTFuture.TransactOpts, _operator, _approved)
}

// SetLpPlus is a paid mutator transaction binding the contract method 0xc87633fd.
//
// Solidity: function setLpPlus(address _lpPlus) returns()
func (_RWTFuture *RWTFutureTransactor) SetLpPlus(opts *bind.TransactOpts, _lpPlus common.Address) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "setLpPlus", _lpPlus)
}

// SetLpPlus is a paid mutator transaction binding the contract method 0xc87633fd.
//
// Solidity: function setLpPlus(address _lpPlus) returns()
func (_RWTFuture *RWTFutureSession) SetLpPlus(_lpPlus common.Address) (*types.Transaction, error) {
	return _RWTFuture.Contract.SetLpPlus(&_RWTFuture.TransactOpts, _lpPlus)
}

// SetLpPlus is a paid mutator transaction binding the contract method 0xc87633fd.
//
// Solidity: function setLpPlus(address _lpPlus) returns()
func (_RWTFuture *RWTFutureTransactorSession) SetLpPlus(_lpPlus common.Address) (*types.Transaction, error) {
	return _RWTFuture.Contract.SetLpPlus(&_RWTFuture.TransactOpts, _lpPlus)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_RWTFuture *RWTFutureTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_RWTFuture *RWTFutureSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.TransferFrom(&_RWTFuture.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_RWTFuture *RWTFutureTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _RWTFuture.Contract.TransferFrom(&_RWTFuture.TransactOpts, _from, _to, _tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RWTFuture *RWTFutureTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RWTFuture *RWTFutureSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RWTFuture.Contract.TransferOwnership(&_RWTFuture.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RWTFuture *RWTFutureTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RWTFuture.Contract.TransferOwnership(&_RWTFuture.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RWTFuture *RWTFutureTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RWTFuture *RWTFutureSession) Unpause() (*types.Transaction, error) {
	return _RWTFuture.Contract.Unpause(&_RWTFuture.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RWTFuture *RWTFutureTransactorSession) Unpause() (*types.Transaction, error) {
	return _RWTFuture.Contract.Unpause(&_RWTFuture.TransactOpts)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RWTFuture *RWTFutureTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RWTFuture.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RWTFuture *RWTFutureSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RWTFuture.Contract.UpgradeToAndCall(&_RWTFuture.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RWTFuture *RWTFutureTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RWTFuture.Contract.UpgradeToAndCall(&_RWTFuture.TransactOpts, newImplementation, data)
}

// RWTFutureApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RWTFuture contract.
type RWTFutureApprovalIterator struct {
	Event *RWTFutureApproval // Event containing the contract specifics and raw log

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
func (it *RWTFutureApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureApproval)
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
		it.Event = new(RWTFutureApproval)
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
func (it *RWTFutureApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureApproval represents a Approval event raised by the RWTFuture contract.
type RWTFutureApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RWTFuture *RWTFutureFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RWTFutureApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RWTFutureApprovalIterator{contract: _RWTFuture.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RWTFuture *RWTFutureFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RWTFutureApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureApproval)
				if err := _RWTFuture.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RWTFuture *RWTFutureFilterer) ParseApproval(log types.Log) (*RWTFutureApproval, error) {
	event := new(RWTFutureApproval)
	if err := _RWTFuture.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RWTFuture contract.
type RWTFutureApprovalForAllIterator struct {
	Event *RWTFutureApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RWTFutureApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureApprovalForAll)
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
		it.Event = new(RWTFutureApprovalForAll)
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
func (it *RWTFutureApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureApprovalForAll represents a ApprovalForAll event raised by the RWTFuture contract.
type RWTFutureApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RWTFuture *RWTFutureFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RWTFutureApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RWTFutureApprovalForAllIterator{contract: _RWTFuture.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RWTFuture *RWTFutureFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RWTFutureApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureApprovalForAll)
				if err := _RWTFuture.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RWTFuture *RWTFutureFilterer) ParseApprovalForAll(log types.Log) (*RWTFutureApprovalForAll, error) {
	event := new(RWTFutureApprovalForAll)
	if err := _RWTFuture.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RWTFuture contract.
type RWTFutureInitializedIterator struct {
	Event *RWTFutureInitialized // Event containing the contract specifics and raw log

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
func (it *RWTFutureInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureInitialized)
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
		it.Event = new(RWTFutureInitialized)
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
func (it *RWTFutureInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureInitialized represents a Initialized event raised by the RWTFuture contract.
type RWTFutureInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RWTFuture *RWTFutureFilterer) FilterInitialized(opts *bind.FilterOpts) (*RWTFutureInitializedIterator, error) {

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RWTFutureInitializedIterator{contract: _RWTFuture.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RWTFuture *RWTFutureFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RWTFutureInitialized) (event.Subscription, error) {

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureInitialized)
				if err := _RWTFuture.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_RWTFuture *RWTFutureFilterer) ParseInitialized(log types.Log) (*RWTFutureInitialized, error) {
	event := new(RWTFutureInitialized)
	if err := _RWTFuture.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the RWTFuture contract.
type RWTFutureOwnershipTransferStartedIterator struct {
	Event *RWTFutureOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *RWTFutureOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureOwnershipTransferStarted)
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
		it.Event = new(RWTFutureOwnershipTransferStarted)
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
func (it *RWTFutureOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the RWTFuture contract.
type RWTFutureOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_RWTFuture *RWTFutureFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RWTFutureOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RWTFutureOwnershipTransferStartedIterator{contract: _RWTFuture.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_RWTFuture *RWTFutureFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *RWTFutureOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureOwnershipTransferStarted)
				if err := _RWTFuture.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_RWTFuture *RWTFutureFilterer) ParseOwnershipTransferStarted(log types.Log) (*RWTFutureOwnershipTransferStarted, error) {
	event := new(RWTFutureOwnershipTransferStarted)
	if err := _RWTFuture.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RWTFuture contract.
type RWTFutureOwnershipTransferredIterator struct {
	Event *RWTFutureOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RWTFutureOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureOwnershipTransferred)
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
		it.Event = new(RWTFutureOwnershipTransferred)
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
func (it *RWTFutureOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureOwnershipTransferred represents a OwnershipTransferred event raised by the RWTFuture contract.
type RWTFutureOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RWTFuture *RWTFutureFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RWTFutureOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RWTFutureOwnershipTransferredIterator{contract: _RWTFuture.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RWTFuture *RWTFutureFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RWTFutureOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureOwnershipTransferred)
				if err := _RWTFuture.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RWTFuture *RWTFutureFilterer) ParseOwnershipTransferred(log types.Log) (*RWTFutureOwnershipTransferred, error) {
	event := new(RWTFutureOwnershipTransferred)
	if err := _RWTFuture.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFuturePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RWTFuture contract.
type RWTFuturePausedIterator struct {
	Event *RWTFuturePaused // Event containing the contract specifics and raw log

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
func (it *RWTFuturePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFuturePaused)
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
		it.Event = new(RWTFuturePaused)
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
func (it *RWTFuturePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFuturePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFuturePaused represents a Paused event raised by the RWTFuture contract.
type RWTFuturePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RWTFuture *RWTFutureFilterer) FilterPaused(opts *bind.FilterOpts) (*RWTFuturePausedIterator, error) {

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RWTFuturePausedIterator{contract: _RWTFuture.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RWTFuture *RWTFutureFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RWTFuturePaused) (event.Subscription, error) {

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFuturePaused)
				if err := _RWTFuture.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RWTFuture *RWTFutureFilterer) ParsePaused(log types.Log) (*RWTFuturePaused, error) {
	event := new(RWTFuturePaused)
	if err := _RWTFuture.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureRWTFutureMintedIterator is returned from FilterRWTFutureMinted and is used to iterate over the raw logs and unpacked data for RWTFutureMinted events raised by the RWTFuture contract.
type RWTFutureRWTFutureMintedIterator struct {
	Event *RWTFutureRWTFutureMinted // Event containing the contract specifics and raw log

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
func (it *RWTFutureRWTFutureMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureRWTFutureMinted)
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
		it.Event = new(RWTFutureRWTFutureMinted)
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
func (it *RWTFutureRWTFutureMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureRWTFutureMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureRWTFutureMinted represents a RWTFutureMinted event raised by the RWTFuture contract.
type RWTFutureRWTFutureMinted struct {
	TokenId        *big.Int
	Receiver       common.Address
	StrikePrice    *big.Int
	AllocatedFil   *big.Int
	ExpirationDate *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRWTFutureMinted is a free log retrieval operation binding the contract event 0x03488890ace1d90f3afce239b92847940fe64099a8f090c57b2af9dd9f25ba32.
//
// Solidity: event RWTFutureMinted(uint256 indexed tokenId, address indexed receiver, uint256 strikePrice, uint256 allocatedFil, uint256 expirationDate)
func (_RWTFuture *RWTFutureFilterer) FilterRWTFutureMinted(opts *bind.FilterOpts, tokenId []*big.Int, receiver []common.Address) (*RWTFutureRWTFutureMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "RWTFutureMinted", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &RWTFutureRWTFutureMintedIterator{contract: _RWTFuture.contract, event: "RWTFutureMinted", logs: logs, sub: sub}, nil
}

// WatchRWTFutureMinted is a free log subscription operation binding the contract event 0x03488890ace1d90f3afce239b92847940fe64099a8f090c57b2af9dd9f25ba32.
//
// Solidity: event RWTFutureMinted(uint256 indexed tokenId, address indexed receiver, uint256 strikePrice, uint256 allocatedFil, uint256 expirationDate)
func (_RWTFuture *RWTFutureFilterer) WatchRWTFutureMinted(opts *bind.WatchOpts, sink chan<- *RWTFutureRWTFutureMinted, tokenId []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "RWTFutureMinted", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureRWTFutureMinted)
				if err := _RWTFuture.contract.UnpackLog(event, "RWTFutureMinted", log); err != nil {
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

// ParseRWTFutureMinted is a log parse operation binding the contract event 0x03488890ace1d90f3afce239b92847940fe64099a8f090c57b2af9dd9f25ba32.
//
// Solidity: event RWTFutureMinted(uint256 indexed tokenId, address indexed receiver, uint256 strikePrice, uint256 allocatedFil, uint256 expirationDate)
func (_RWTFuture *RWTFutureFilterer) ParseRWTFutureMinted(log types.Log) (*RWTFutureRWTFutureMinted, error) {
	event := new(RWTFutureRWTFutureMinted)
	if err := _RWTFuture.contract.UnpackLog(event, "RWTFutureMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RWTFuture contract.
type RWTFutureTransferIterator struct {
	Event *RWTFutureTransfer // Event containing the contract specifics and raw log

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
func (it *RWTFutureTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureTransfer)
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
		it.Event = new(RWTFutureTransfer)
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
func (it *RWTFutureTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureTransfer represents a Transfer event raised by the RWTFuture contract.
type RWTFutureTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RWTFuture *RWTFutureFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RWTFutureTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RWTFutureTransferIterator{contract: _RWTFuture.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RWTFuture *RWTFutureFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RWTFutureTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureTransfer)
				if err := _RWTFuture.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RWTFuture *RWTFutureFilterer) ParseTransfer(log types.Log) (*RWTFutureTransfer, error) {
	event := new(RWTFutureTransfer)
	if err := _RWTFuture.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RWTFuture contract.
type RWTFutureUnpausedIterator struct {
	Event *RWTFutureUnpaused // Event containing the contract specifics and raw log

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
func (it *RWTFutureUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureUnpaused)
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
		it.Event = new(RWTFutureUnpaused)
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
func (it *RWTFutureUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureUnpaused represents a Unpaused event raised by the RWTFuture contract.
type RWTFutureUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RWTFuture *RWTFutureFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RWTFutureUnpausedIterator, error) {

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RWTFutureUnpausedIterator{contract: _RWTFuture.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RWTFuture *RWTFutureFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RWTFutureUnpaused) (event.Subscription, error) {

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureUnpaused)
				if err := _RWTFuture.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RWTFuture *RWTFutureFilterer) ParseUnpaused(log types.Log) (*RWTFutureUnpaused, error) {
	event := new(RWTFutureUnpaused)
	if err := _RWTFuture.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RWTFutureUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RWTFuture contract.
type RWTFutureUpgradedIterator struct {
	Event *RWTFutureUpgraded // Event containing the contract specifics and raw log

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
func (it *RWTFutureUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RWTFutureUpgraded)
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
		it.Event = new(RWTFutureUpgraded)
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
func (it *RWTFutureUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RWTFutureUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RWTFutureUpgraded represents a Upgraded event raised by the RWTFuture contract.
type RWTFutureUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RWTFuture *RWTFutureFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RWTFutureUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RWTFuture.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RWTFutureUpgradedIterator{contract: _RWTFuture.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RWTFuture *RWTFutureFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RWTFutureUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RWTFuture.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RWTFutureUpgraded)
				if err := _RWTFuture.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RWTFuture *RWTFutureFilterer) ParseUpgraded(log types.Log) (*RWTFutureUpgraded, error) {
	event := new(RWTFutureUpgraded)
	if err := _RWTFuture.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
