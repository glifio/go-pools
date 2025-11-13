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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyLpPlus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"name\":\"RWTFutureMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getPositionsBatched\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"internalType\":\"structPosition[]\",\"name\":\"positions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractILpPlus\",\"name\":\"_lpPlus\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpPlus\",\"outputs\":[{\"internalType\":\"contractILpPlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_strikePrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_allocatedFils\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_expirationDates\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILpPlus\",\"name\":\"_lpPlus\",\"type\":\"address\"}],\"name\":\"setLpPlus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f516020613ba15f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b604051613ada90816100c78239608051818181611aea0152611bcb0152f35b6001600160401b0319166001600160401b039081175f516020613ba15f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a714612b5d5750806306fdde0314612a5c578063081812fc146129d7578063095ea7b31461280a5780631d05159d1461213f57806323b872dd1461209b5780632fc928881461202b5780633c9774c514611f8b5780633f4ba83a14611e9257806342842e0e14611e615780634f1ef28614611b6257806352d1902d14611aa557806354fd4d5014611a6c5780635c975abb14611a0d5780636352211e146119b357806370a08231146118df578063715018a61461179f57806374acd2091461170d57806379ba50971461166c5780638456cb59146115995780638da5cb5b146115295780638f15b41414610d0c5780638f68679f14610cb257806395d89b4114610b5a578063a22cb46514610a30578063ad3cb1cc146109b3578063b80f55c9146106c2578063b88d4fde14610579578063c87633fd14610453578063c87b56dd146103d8578063e30c397814610367578063e985e9c5146102b05763f2fde38b14610189575f80fd5b346102ad5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ad5773ffffffffffffffffffffffffffffffffffffffff6101dd6101d8612c8d565b6131fa565b6101e5613262565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b80fd5b50346102ad5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ad576102e8612c8d565b73ffffffffffffffffffffffffffffffffffffffff61034c610308612cb0565b9273ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b91165f52602052602060ff60405f2054166040519015158152f35b50346102ad57807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102ad57602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5761041260043561313e565b505f604051610422602082612cef565b5261044b604051610434602082612cef565b5f8152604051918291602083526020830190612c4a565b0390f35b5f80fd5b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5773ffffffffffffffffffffffffffffffffffffffff61049f612c8d565b6104a7613262565b1680156105515773ffffffffffffffffffffffffffffffffffffffff6104cf61054f926131fa565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416177f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0155565b005b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461044f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f576105b0612c8d565b6105b8612cb0565b9060443560643567ffffffffffffffff811161044f576105eb6105e26105f1923690600401612e51565b936101d86131a7565b936131fa565b6105f96131a7565b610602846131fa565b9361060c826131fa565b73ffffffffffffffffffffffffffffffffffffffff8116156106965773ffffffffffffffffffffffffffffffffffffffff61064a819286339161342a565b9616951694808603610662575061054f94503361385b565b8386917f64283d7b000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b7f64a0ae92000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5760043567ffffffffffffffff811161044f57610711903690600401612d48565b6107196131a7565b73ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416330361098b578051905f5b82811061076557005b61076f8183612f20565b51805f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020525f60026040822082815582600182015501556107f1815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b73ffffffffffffffffffffffffffffffffffffffff811690825f83159384156108c3575b8282527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079302602052604082207fffffffffffffffffffffffff000000000000000000000000000000000000000081541690557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8280a450610898575060010161075c565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b610919835f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b6109608473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8154019055610815565b7f46892297000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5761044b6040516109f2604082612cef565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190612c4a565b3461044f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57610a67612c8d565b6024359081151580920361044f57610a9673ffffffffffffffffffffffffffffffffffffffff916101d86131a7565b16908115610b2e57335f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260409020825f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff83161790556040519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b507f5b08ba18000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154610bb781612e95565b8084529060018116908115610c705750600114610bf3575b61044b83610bdf81850382612cef565b604051918291602083526020830190612c4a565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f9081527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b808210610c5657509091508101602001610bdf610bcf565b919260018160209254838588010152019101909291610c3e565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b84019091019150610bdf9050610bcf565b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5760207f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054604051908152f35b3461044f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5760043567ffffffffffffffff811161044f57610d5b903690600401612e51565b60243567ffffffffffffffff811161044f57610d7b903690600401612e51565b9060443573ffffffffffffffffffffffffffffffffffffffff8116810361044f576064359073ffffffffffffffffffffffffffffffffffffffff821680920361044f577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549260ff8460401c16159367ffffffffffffffff811680159081611521575b6001149081611517575b15908161150e575b506114e6578460017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055611491575b50610e6c6137d5565b610e746137d5565b80519067ffffffffffffffff8211611292578190610eb27f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054612e95565b601f81116113c7575b50602090601f83116001146112ca575f926112bf575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b835167ffffffffffffffff811161129257610f5e7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154612e95565b601f8111611210575b506020601f821160011461112c578190610fde9495965f92611121575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b6101d86137d5565b90610fe76137d5565b610fef6137d5565b73ffffffffffffffffffffffffffffffffffffffff8216156110f5576104cf73ffffffffffffffffffffffffffffffffffffffff91611030611038946132a2565b610fd66137d5565b60017f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a005561106257005b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b7f1e4fbdf7000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b015190508680610f84565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216957f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f52815f20965f5b8181106111f8575091610fde959697918460019594106111c1575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930155610fd6565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055868080611194565b83830151895560019098019760209384019301611179565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f830160051c81019160208410611288575b601f0160051c01905b81811061127d5750610f67565b5f8155600101611270565b9091508190611267565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b015190508680610ed1565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81925f5b8181106113af5750908460019594939210611378575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055610f23565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c1916905586808061134b565b92936020600181928786015181550195019301611335565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f52601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410611469575b90601f8493920160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b81811061145b5750610ebb565b5f815584935060010161144e565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf819150611420565b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005585610e63565b7ff92ee8a9000000000000000000000000000000000000000000000000000000005f5260045ffd5b90501587610e10565b303b159150610e08565b869150610dfe565b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f576115cf613262565b6115d76131a7565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416036116e15761054f336132a2565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57611744612f6e565b506004355f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a02602052606061177c60405f20612f8c565b61179d60405180926040809180518452602081015160208501520151910152565bf35b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f576117d5613262565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00555f73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f576119196101d8612c8d565b73ffffffffffffffffffffffffffffffffffffffff8116156119875761197e60209173ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b54604051908152f35b7f89c62b64000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5760206119ef60043561313e565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57602060405160018152f35b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5773ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163003611b3a5760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57611b94612c8d565b60243567ffffffffffffffff811161044f57611bb4903690600401612e51565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016803014908115611e1f575b50611b3a57611c03613262565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181611deb575b50611c8357837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc859203611dc05750813b15611d9557807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115611d64575f8083602061054f95519101845af4611d5e61382c565b91613a0b565b505034611d6d57005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011611e17575b81611e0760209383612cef565b8101031261044f57519085611c52565b3d9150611dfa565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141583611bf6565b3461044f57611e6f36612da5565b60405191926105f1906105eb90611e87602086612cef565b5f85526101d86131a7565b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57611ec8613262565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff811615611f63577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461044f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57611fc8602435600435612fb4565b6040518091602082016020835281518091526020604084019201905f5b818110611ff3575050500390f35b91935091602060608261201d60019488516040809180518452602081015160208501520151910152565b019401910191849392611fe5565b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57602073ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416604051908152f35b3461044f576120bc6120c26120af36612da5565b92916101d89491946131a7565b926131fa565b73ffffffffffffffffffffffffffffffffffffffff8116156106965773ffffffffffffffffffffffffffffffffffffffff612100819284339161342a565b931692169180830361210e57005b7f64283d7b000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b3461044f5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57612176612c8d565b60243567ffffffffffffffff811161044f57612196903690600401612d48565b60443567ffffffffffffffff811161044f576121b6903690600401612d48565b60643567ffffffffffffffff811161044f576121d6903690600401612d48565b906121df6131a7565b73ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416330361098b57829192519061222b82612d30565b936122396040519586612cef565b82855261224583612d30565b937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020870195013686377f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054925f9273ffffffffffffffffffffffffffffffffffffffff891692831560209a8115965b8c8a821061232b578c8c6122ca8d8d612ee6565b7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a005560405192839281840190828552518091526040840192915f5b82811061231457505050500390f35b835185528695509381019392810192600101612305565b846123c0838f9e9f8a60028f968c7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026123878761238081612379612371829f8390612ee6565b9e8f9e612f20565b5198612f20565b5193612f20565b51926040519561239687612cd3565b865280860192835260408601938452895f525260405f209351845551600184015551910155612f20565b526040516123ce8f82612cef565b845f825261069657898f838b612422825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b927f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930273ffffffffffffffffffffffffffffffffffffffff8516918215159687612742575b6126ef575b845f525260405f20827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a4506126c357833b612533575b5090876001928f60608a7f03488890ace1d90f3afce239b92847940fe64099a8f090c57b2af9dd9f25ba32928b61251988612512818f8161250b91612f20565b5196612f20565b5192612f20565b51916040519384528301526040820152a3019a999a6122b6565b8e899f9e9d9c9b9a9961259282915f946040519586809481937f150b7a020000000000000000000000000000000000000000000000000000000083523360048401528360248401528a6044840152608060648401526084830190612c4a565b03925af19182915f9361266a575b50506125e7578e8e6125b061382c565b805191826125e457837f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b01fd5b7fffffffff000000000000000000000000000000000000000000000000000000009d9e98999a9b9c9d7f150b7a020000000000000000000000000000000000000000000000000000000091160361263e57876124cb565b877f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9080929350813d83116126bc575b6126828183612cef565b8101031261044f57517fffffffff000000000000000000000000000000000000000000000000000000008116810361044f57905f8f6125a0565b503d612678565b7f73c6ac6e000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b6127368b73ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b6001815401905561246b565b612798865f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b6127df8773ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8154019055612466565b3461044f5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57612841612c8d565b612850602435916101d86131a7565b9061285a8161313e565b331515806129b7575b80612945575b61291957819073ffffffffffffffffffffffffffffffffffffffff80851691167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9255f80a45f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f2091167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790555f80f35b7fa9fbf51f000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b5061298d8173ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b73ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f20541615612869565b503373ffffffffffffffffffffffffffffffffffffffff82161415612863565b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f5760206119ef600435612a178161313e565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b3461044f575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054612ab981612e95565b8084529060018116908115610c705750600114612ae05761044b83610bdf81850382612cef565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b808210612b4357509091508101602001610bdf610bcf565b919260018160209254838588010152019101909291612b2b565b3461044f5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261044f57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361044f57817f80ac58cd0000000000000000000000000000000000000000000000000000000060209314908115612c20575b8115612bf6575b5015158152f35b7f01ffc9a70000000000000000000000000000000000000000000000000000000091501483612bef565b7f5b5e139f0000000000000000000000000000000000000000000000000000000081149150612be8565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361044f57565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361044f57565b6060810190811067ffffffffffffffff82111761129257604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff82111761129257604052565b67ffffffffffffffff81116112925760051b60200190565b9080601f8301121561044f578135612d5f81612d30565b92612d6d6040519485612cef565b81845260208085019260051b82010192831161044f57602001905b828210612d955750505090565b8135815260209182019101612d88565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc606091011261044f5760043573ffffffffffffffffffffffffffffffffffffffff8116810361044f579060243573ffffffffffffffffffffffffffffffffffffffff8116810361044f579060443590565b67ffffffffffffffff811161129257601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561044f57602081359101612e6b82612e17565b92612e796040519485612cef565b8284528282011161044f57815f92602092838601378301015290565b90600182811c92168015612edc575b6020831014612eaf57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691612ea4565b91908201809211612ef357565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b8051821015612f345760209160051b010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b91908203918211612ef357565b60405190612f7b82612cd3565b5f6040838281528260208201520152565b90604051612f9981612cd3565b60406002829480548452600181015460208501520154910152565b91907f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201828111612ef3578411613102576130139084612ee6565b908082116130fa575b506130278382612f61565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe061306b61305586612d30565b956130636040519788612cef565b808752612d30565b015f5b8181106130e357505083815b8381106130875750505050565b806001915f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020526130dc60405f206130cb6130c58785612f61565b91612f8c565b6130d58287612f20565b5284612f20565b500161307a565b6020906130ee612f6e565b8282890101520161306e565b90505f61301c565b5090915050604051613115602082612cef565b5f81525f805b81811061312757505090565b602090613132612f6e565b8282860101520161311b565b613186815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff821615610898575090565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166131d257565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff000000000000000000000000000000000000001461324e575b1561324a5761323c906133b1565b90613245575090565b905090565b5090565b505067ffffffffffffffff8116600161322e565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300541633036116e157565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff160361341f575b82158015613414575b61340c57565b5f9250829150565b5060163d1415613406565b5f92508291506133fd565b73ffffffffffffffffffffffffffffffffffffffff613487835f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b931680151580613666575b505073ffffffffffffffffffffffffffffffffffffffff83168061359e575b73ffffffffffffffffffffffffffffffffffffffff8216918261354a575b50825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f20827fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790557fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a490565b6135919073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190555f6134cf565b6135f4835f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b61363b8473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190556134b1565b806136ee575b156136775780613492565b905073ffffffffffffffffffffffffffffffffffffffff83166136c057507f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7f177e802f000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b508073ffffffffffffffffffffffffffffffffffffffff851614801561377a575b8061366c57508073ffffffffffffffffffffffffffffffffffffffff613773855f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b161461366c565b506137c28473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b815f5260205260ff60405f20541661370f565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561380457565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b3d15613856573d9061383d82612e17565b9161384b6040519384612cef565b82523d5f602084013e565b606090565b93909293823b61386d575b5050505050565b6138da73ffffffffffffffffffffffffffffffffffffffff928360209516968460405197889687967f150b7a020000000000000000000000000000000000000000000000000000000088521660048701521660248501526044840152608060648401526084830190612c4a565b03815f865af15f91816139ae575b5061392f57506138f661382c565b8051908161392a57827f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b602001fd5b7fffffffff000000000000000000000000000000000000000000000000000000007f150b7a020000000000000000000000000000000000000000000000000000000091160361398357505f80808080613866565b7f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011613a03575b816139ca60209383612cef565b8101031261044f57517fffffffff000000000000000000000000000000000000000000000000000000008116810361044f57905f6138e8565b3d91506139bd565b90613a485750805115613a2057805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b81511580613a9b575b613a59575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b15613a5156fea26469706673582212208adbdb7985dae1f88236ce850d8a6ec4a18b697d73661572bffa8454d6df5c5164736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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
