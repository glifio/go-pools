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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyLpPlus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"name\":\"RWTFutureMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getPositionsBatched\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"internalType\":\"structPosition[]\",\"name\":\"positions\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_initialOwner\",\"type\":\"address\"},{\"internalType\":\"contractILpPlus\",\"name\":\"_lpPlus\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpPlus\",\"outputs\":[{\"internalType\":\"contractILpPlus\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_strikePrices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_allocatedFils\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_expirationDates\",\"type\":\"uint256[]\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILpPlus\",\"name\":\"_lpPlus\",\"type\":\"address\"}],\"name\":\"setLpPlus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenIdGenerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenIdToPosition\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allocatedFil\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationDate\",\"type\":\"uint256\"}],\"internalType\":\"structPosition\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x60a080604052346100c257306080525f516020614a415f395f51905f525460ff8160401c166100b3576002600160401b03196001600160401b03821601610060575b60405161497a90816100c7823960805181818161226401526123e90152f35b6001600160401b0319166001600160401b039081175f516020614a415f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80610041565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816301ffc9a7146137075750806306fdde0314613606578063081812fc14613581578063095ea7b3146133b457806318160ddd1461335a5780631d05159d14612bfe57806323b872dd14612b5a5780632f745c5914612a885780632fc9288814612a185780633c9774c5146129785780633f4ba83a1461287f57806342842e0e1461267f5780634f1ef286146123805780634f6ccce7146122dc57806352d1902d1461221f57806354fd4d50146121e65780635c975abb146121875780636352211e1461212d57806370a08231146120e6578063715018a614611fa657806374acd20914611f1457806379ba509714611e735780638456cb5914611da05780638da5cb5b14611d305780638f15b414146115385780638f68679f146114de57806395d89b4114611386578063a22cb4651461125c578063ac9650d814610fde578063ad3cb1cc14610f61578063b80f55c91461087e578063b88d4fde146105a5578063c87633fd1461047f578063c87b56dd14610404578063e30c397814610393578063e985e9c5146102dc5763f2fde38b146101b5575f80fd5b346102d95760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d95773ffffffffffffffffffffffffffffffffffffffff61020961020461386a565b613e00565b610211613e68565b16807fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416177f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e227008380a380f35b80fd5b50346102d95760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d95761031461386a565b73ffffffffffffffffffffffffffffffffffffffff61037861033461388d565b9273ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b91165f52602052602060ff60405f2054166040519015158152f35b50346102d957807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc3601126102d957602073ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005416604051908152f35b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5761043e600435613d44565b505f60405161044e6020826138cc565b526104776040516104606020826138cc565b5f8152604051918291602083526020830190613827565b0390f35b5f80fd5b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5773ffffffffffffffffffffffffffffffffffffffff6104cb61386a565b6104d3613e68565b16801561057d5773ffffffffffffffffffffffffffffffffffffffff6104fb61057b92613e00565b1673ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffff00000000000000000000000000000000000000007f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416177f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0155565b005b7fd92e233d000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461047b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576105dc61386a565b6105e461388d565b6044359060643567ffffffffffffffff811161047b5761061761060e61061d923690600401613a2e565b94610204613dad565b91613e00565b610625613dad565b61062e82613e00565b61063782613e00565b73ffffffffffffffffffffffffffffffffffffffff8116156108525773ffffffffffffffffffffffffffffffffffffffff61067581928733916140f2565b921691169080820361081f575050803b61068b57005b6020916106f773ffffffffffffffffffffffffffffffffffffffff8093169560405195869485947f150b7a020000000000000000000000000000000000000000000000000000000086523360048701521660248501526044840152608060648401526084830190613827565b03815f865af15f91816107c2575b5061074c575061071361404a565b8051908161074757827f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b602001fd5b7fffffffff000000000000000000000000000000000000000000000000000000007f150b7a020000000000000000000000000000000000000000000000000000000091160361079757005b7f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011610817575b816107de602093836138cc565b8101031261047b57517fffffffff000000000000000000000000000000000000000000000000000000008116810361047b579083610705565b3d91506107d1565b84907f64283d7b000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b7f64a0ae92000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760043567ffffffffffffffff811161047b576108cd903690600401613925565b6108d5613dad565b73ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0154163303610f39578051905f5b82811061092157005b61092b8183613ad0565b51805f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020525f60026040822082815582600182015501556109ad815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b73ffffffffffffffffffffffffffffffffffffffff811690811591821580610e71575b845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055845f837fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8280a48315610d6c5750507f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254835f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060405f205568010000000000000000811015610d3f5783610ae3826001610b1994017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255613cee565b9091907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff83549160031b92831b921b1916179055565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111610d1257835f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed03602052610b9a60405f205491613cee565b90549060031b1c610bae81610ae384613cee565b5f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205260405f2055825f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020525f60408120557f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02548015610ce5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610c8a610c5c82613cee565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82549160031b1b19169055565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025550610cba5750600101610918565b7f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d77575b50610b19565b610d88610d8383613e00565b613ea8565b845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f2054915f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f2091818103610e27575b50845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020525f60408120555f526020525f604081205586610d71565b815f528260205260405f2054815f52836020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f205588610de8565b610ec7855f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b610f0e8373ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81540190556109d0565b7f46892297000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57610477604051610fa06040826138cc565b600581527f352e302e300000000000000000000000000000000000000000000000000000006020820152604051918291602083526020830190613827565b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760043567ffffffffffffffff811161047b573660238201121561047b5780600401359067ffffffffffffffff821161047b573660248360051b8301011161047b579060209160405161105f84826138cc565b5f8152838101917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe085013684376110958461390d565b906110a360405192836138cc565b8482527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06110d08661390d565b01865f5b82811061124d575050505f907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbd81360301915b868110156111d25760248160051b830101358381121561047b5782019060248201359167ffffffffffffffff831161047b5760440191803603831361047b575f80600194896111a26111b6958f8e6040519483869484860198893784019083820190898252519283915e0101858152037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018352826138cc565b5190305af46111af61404a565b90306148ab565b6111c08287613ad0565b526111cb8186613ad0565b5001611107565b604080518981528551818b018190525f92600582901b8301810191888d01918d9085015b8287106112035785850386f35b90919293828061123d837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc08a600196030186528851613827565b96019201960195929190926111f6565b606085820183015281016110d4565b3461047b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5761129361386a565b6024359081151580920361047b576112c273ffffffffffffffffffffffffffffffffffffffff91610204613dad565b1690811561135a57335f9081527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260409020825f5260205260405f207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0081541660ff83161790556040519081527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c3160203392a3005b507f5b08ba18000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301546113e381613a72565b808452906001811690811561149c575060011461141f575b6104778361140b818503826138cc565b604051918291602083526020830190613827565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f9081527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e939250905b8082106114825750909150810160200161140b6113fb565b91926001816020925483858801015201910190929161146a565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660208086019190915291151560051b8401909101915061140b90506113fb565b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760207f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054604051908152f35b3461047b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760043567ffffffffffffffff811161047b57611587903690600401613a2e565b60243567ffffffffffffffff811161047b576115a7903690600401613a2e565b9060443573ffffffffffffffffffffffffffffffffffffffff8116810361047b576064359073ffffffffffffffffffffffffffffffffffffffff821680920361047b577ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00549260ff8460401c16159367ffffffffffffffff811680159081611d28575b6001149081611d1e575b159081611d15575b50611ced578460017fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008316177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0055611c98575b50611698614854565b6116a0614854565b6116a8614854565b80519067ffffffffffffffff8211610d3f5781906116e67f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930054613a72565b601f8111611bce575b50602090601f8311600114611ad1575f92611ac6575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079300555b835167ffffffffffffffff8111610d3f576117927f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930154613a72565b601f8111611a44575b506020601f82116001146119605781906118129495965f92611955575b50507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8260011b9260031b1c1916177f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab0079301555b610204614854565b9061181b614854565b611823614854565b73ffffffffffffffffffffffffffffffffffffffff821615611929576104fb73ffffffffffffffffffffffffffffffffffffffff9161186461186c94613f3b565b61180a614854565b60017f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a005561189657005b7fffffffffffffffffffffffffffffffffffffffffffffff00ffffffffffffffff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054167ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b7f1e4fbdf7000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b0151905086806117b8565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08216957f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f52815f20965f5b818110611a2c575091611812959697918460019594106119f5575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015561180a565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c191690558680806119c8565b838301518955600190980197602093840193016119ad565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793015f527ff4bad0a69248f59680a4f2b3000328cec71a413447c96781cfe5996daa8c456e601f830160051c81019160208410611abc575b601f0160051c01905b818110611ab1575061179b565b5f8155600101611aa4565b9091508190611a9b565b015190508680611705565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016917f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81925f5b818110611bb65750908460019594939210611b7f575b505050811b017f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930055611757565b01517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60f88460031b161c19169055868080611b52565b92936020600181928786015181550195019301611b3c565b9091507f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f52601f830160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81019060208410611c70575b90601f8493920160051c7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf8101905b818110611c6257506116ef565b5f8155849350600101611c55565b7f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf819150611c27565b7fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000001668010000000000000001177ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00558561168f565b7ff92ee8a9000000000000000000000000000000000000000000000000000000005f5260045ffd5b9050158761163c565b303b159150611634565b86915061162a565b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602073ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c1993005416604051908152f35b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57611dd6613e68565b611dde613dad565b60017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416177fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586020604051338152a1005b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b573373ffffffffffffffffffffffffffffffffffffffff7f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00541603611ee85761057b33613f3b565b7f118cdaa7000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57611f4b613b1e565b506004355f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a026020526060611f8360405f20613b3c565b611fa460405180926040809180518452602081015160208501520151910152565bf35b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57611fdc613e68565b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c00555f73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300547fffffffffffffffffffffffff000000000000000000000000000000000000000081167f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576020612125610d8361020461386a565b604051908152f35b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576020612169600435613d44565b73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602060ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f0330054166040519015158152f35b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602060405160018152f35b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001630036122b45760206040517f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8152f35b7fe07c8dba000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576004357f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025481101561235157612342602091613cee565b90549060031b1c604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f525f60045260245260445ffd5b60407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576123b261386a565b60243567ffffffffffffffff811161047b576123d2903690600401613a2e565b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001680301490811561263d575b506122b457612421613e68565b73ffffffffffffffffffffffffffffffffffffffff8216916040517f52d1902d000000000000000000000000000000000000000000000000000000008152602081600481875afa5f9181612609575b506124a157837f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8592036125de5750813b156125b357807fffffffffffffffffffffffff00000000000000000000000000000000000000007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416177f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc557fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b5f80a2815115612582575f8083602061057b95519101845af461257c61404a565b916148ab565b50503461258b57005b7fb398979f000000000000000000000000000000000000000000000000000000005f5260045ffd5b7f4c9c8ce3000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7faa1d49a4000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b9091506020813d602011612635575b81612625602093836138cc565b8101031261047b57519085612470565b3d9150612618565b905073ffffffffffffffffffffffffffffffffffffffff7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5416141583612414565b3461047b5761268d36613982565b6020916126af610617604051956126a486886138cc565b5f8752610204613dad565b6126b7613dad565b6126c082613e00565b6126c982613e00565b73ffffffffffffffffffffffffffffffffffffffff8116156108525773ffffffffffffffffffffffffffffffffffffffff61270781928733916140f2565b921691169080820361081f575050803b61271d57005b839161278873ffffffffffffffffffffffffffffffffffffffff8093169660405195869485947f150b7a020000000000000000000000000000000000000000000000000000000086523360048701521660248501526044840152608060648401526084830190613827565b03815f875af15f9181612827575b506127db57506127a461404a565b805191826127d857837f64a0ae92000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b01fd5b7f150b7a020000000000000000000000000000000000000000000000000000000091507fffffffff00000000000000000000000000000000000000000000000000000000160361079757005b9091508281813d8311612878575b61283f81836138cc565b8101031261047b57517fffffffff000000000000000000000000000000000000000000000000000000008116810361047b579084612796565b503d612835565b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576128b5613e68565b7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005460ff811615612950577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00167fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6020604051338152a1005b7f8dfc202b000000000000000000000000000000000000000000000000000000005f5260045ffd5b3461047b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576129b5602435600435613b64565b6040518091602082016020835281518091526020604084019201905f5b8181106129e0575050500390f35b919350916020606082612a0a60019488516040809180518452602081015160208501520151910152565b0194019101918493926129d2565b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57602073ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a015416604051908152f35b3461047b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57612abf61386a565b73ffffffffffffffffffffffffffffffffffffffff60243591612ae4610d8382613e00565b831015612b2b57165f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20905f52602052602060405f2054604051908152f35b7fa57d13dc000000000000000000000000000000000000000000000000000000005f521660045260245260445ffd5b3461047b57612b7b612b81612b6e36613982565b9291610204949194613dad565b92613e00565b73ffffffffffffffffffffffffffffffffffffffff8116156108525773ffffffffffffffffffffffffffffffffffffffff612bbf81928433916140f2565b9316921691808303612bcd57005b7f64283d7b000000000000000000000000000000000000000000000000000000005f5260045260245260445260645ffd5b3461047b5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57612c3561386a565b60243567ffffffffffffffff811161047b57612c55903690600401613925565b9060443567ffffffffffffffff811161047b57612c76903690600401613925565b60643567ffffffffffffffff811161047b57612c96903690600401613925565b612c9e613dad565b73ffffffffffffffffffffffffffffffffffffffff7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0154163303610f39578392935191612cea8361390d565b93612cf860405195866138cc565b838552612d048461390d565b957fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe06020870197013688377f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054935f9273ffffffffffffffffffffffffffffffffffffffff8316928315908115915b898710612dea578b8b612d868c8c613ac3565b7f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0055604051918291602083019060208452518091526040830191905f5b818110612dd1575050500390f35b8251845285945060209384019390920191600101612dc3565b869084612e82838e9f9e8c60028f612e2585612e1e8f93612e17612e0f848093613ac3565b9d8e9c613ad0565b5196613ad0565b5192613ad0565b519060405193612e34856138b0565b84526020840190815260408401918252875f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0260205260405f209351845551600184015551910155613ad0565b5261085257612ecf815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b968373ffffffffffffffffffffffffffffffffffffffff891698891580159283613290575b61323d575b845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f208a7fffffffffffffffffffffffff0000000000000000000000000000000000000000825416179055848a8c7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a415613141577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060405f205568010000000000000000811015610d3f5784610ae382600161300b94017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255613cee565b885f9a036130a3575b5061307757866001927f03488890ace1d90f3afce239b92847940fe64099a8f090c57b2af9dd9f25ba3260608c61304f86612e1e818d613ad0565b5161305a878d613ad0565b519060405192835260208301526040820152a301959a999a612d73565b7f73c6ac6e000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b6130af610d8386613e00565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111610d1257895f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20815f526020528460405f2055845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f2055613014565b89891461300b57613154610d8382613e00565b845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f2054908b5f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20918181036131f3575b50855f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020525f60408120555f526020525f604081205561300b565b815f528260205260405f2054815f52836020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f20555f6131b5565b6132848673ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b60018154019055612ef9565b506132e7855f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b8661332f8373ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8154019055612ef4565b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760207f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254604051908152f35b3461047b5760407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576133eb61386a565b6133fa60243591610204613dad565b9061340481613d44565b33151580613561575b806134ef575b6134c357819073ffffffffffffffffffffffffffffffffffffffff80851691167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9255f80a45f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f2091167fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790555f80f35b7fa9fbf51f000000000000000000000000000000000000000000000000000000005f523360045260245ffd5b506135378173ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b73ffffffffffffffffffffffffffffffffffffffff33165f5260205260ff60405f20541615613413565b503373ffffffffffffffffffffffffffffffffffffffff8216141561340d565b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b5760206121696004356135c181613d44565b505f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b3461047b575f7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b576040515f7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005461366381613a72565b808452906001811690811561149c575060011461368a576104778361140b818503826138cc565b7f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab00793005f9081527f37c58c799b6609234b945e882912ee9ad34948a1dfaa20a97485e1a7752bbf81939250905b8082106136ed5750909150810160200161140b6113fb565b9192600181602092548385880101520191019092916136d5565b3461047b5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261047b57600435907fffffffff00000000000000000000000000000000000000000000000000000000821680920361047b57817f780e9d630000000000000000000000000000000000000000000000000000000060209314908115613799575b5015158152f35b7f80ac58cd000000000000000000000000000000000000000000000000000000008114915081156137fd575b81156137d3575b5083613792565b7f01ffc9a700000000000000000000000000000000000000000000000000000000915014836137cc565b7f5b5e139f00000000000000000000000000000000000000000000000000000000811491506137c5565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f602080948051918291828752018686015e5f8582860101520116010190565b6004359073ffffffffffffffffffffffffffffffffffffffff8216820361047b57565b6024359073ffffffffffffffffffffffffffffffffffffffff8216820361047b57565b6060810190811067ffffffffffffffff821117610d3f57604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff821117610d3f57604052565b67ffffffffffffffff8111610d3f5760051b60200190565b9080601f8301121561047b57813561393c8161390d565b9261394a60405194856138cc565b81845260208085019260051b82010192831161047b57602001905b8282106139725750505090565b8135815260209182019101613965565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc606091011261047b5760043573ffffffffffffffffffffffffffffffffffffffff8116810361047b579060243573ffffffffffffffffffffffffffffffffffffffff8116810361047b579060443590565b67ffffffffffffffff8111610d3f57601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b81601f8201121561047b57602081359101613a48826139f4565b92613a5660405194856138cc565b8284528282011161047b57815f92602092838601378301015290565b90600182811c92168015613ab9575b6020831014613a8c57565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b91607f1691613a81565b91908201809211610d1257565b8051821015613ae45760209160051b010190565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b91908203918211610d1257565b60405190613b2b826138b0565b5f6040838281528260208201520152565b90604051613b49816138b0565b60406002829480548452600181015460208501520154910152565b91907f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a0054907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8201828111610d12578411613cb257613bc39084613ac3565b90808211613caa575b50613bd78382613b11565b927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0613c1b613c058661390d565b95613c1360405197886138cc565b80875261390d565b015f5b818110613c9357505083815b838110613c375750505050565b806001915f527f4767c911798e2a4728b741b24d4756cff1b3ea90e82ca42399659ef3e6f74a02602052613c8c60405f20613c7b613c758785613b11565b91613b3c565b613c858287613ad0565b5284613ad0565b5001613c2a565b602090613c9e613b1e565b82828901015201613c1e565b90505f613bcc565b5090915050604051613cc56020826138cc565b5f81525f805b818110613cd757505090565b602090613ce2613b1e565b82828601015201613ccb565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254811015613ae4577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025f5260205f2001905f90565b613d8c815f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9073ffffffffffffffffffffffffffffffffffffffff821615610cba575090565b60ff7fcd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f033005416613dd857565b7fd93c0665000000000000000000000000000000000000000000000000000000005f5260045ffd5b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000014613e54575b15613e5057613e4290614079565b90613e4b575090565b905090565b5090565b505067ffffffffffffffff81166001613e34565b73ffffffffffffffffffffffffffffffffffffffff7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054163303611ee857565b73ffffffffffffffffffffffffffffffffffffffff811615613f0f57613f0b9073ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b5490565b7f89c62b64000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b7fffffffffffffffffffffffff00000000000000000000000000000000000000007f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c0054167f237e158222e3e6968b72b9db0d8043aacf074ad9f650f0d1606b4d82ee432c005573ffffffffffffffffffffffffffffffffffffffff807f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930054921691827fffffffffffffffffffffffff00000000000000000000000000000000000000008216177f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930055167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b3d15614074573d9061405b826139f4565b9161406960405193846138cc565b82523d5f602084013e565b606090565b5f526016600a60205f73fe000000000000000000000000000000000000025afa905f519061040a8273ffffffffffffffffffffffffffffffffffffffff169260a01c61ffff16036140e7575b821580156140dc575b6140d457565b5f9250829150565b5060163d14156140ce565b5f92508291506140c5565b9073ffffffffffffffffffffffffffffffffffffffff614150825f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b9316801515806146e7575b505073ffffffffffffffffffffffffffffffffffffffff831691821592831561461f575b73ffffffffffffffffffffffffffffffffffffffff82169384159081156145cc575b845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930260205260405f20867fffffffffffffffffffffffff00000000000000000000000000000000000000008254161790558486847fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef5f80a4156144d0577f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0254845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020528060405f205568010000000000000000811015610d3f5784610ae38260016142b394017f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0255613cee565b1561442357505090507f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8101908111610d1257815f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205261433d60405f205491613cee565b90549060031b1c61435181610ae384613cee565b5f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0360205260405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed036020525f60408120557f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed02548015610ce5577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016143fe610c5c82613cee565b7f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed025590565b8303614430575b50505090565b610d8361443c91613e00565b917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8301928311610d12575f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f20825f526020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f20555f808061442a565b8185146142b3576144e3610d8387613e00565b845f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f205490835f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0060205260405f2091818103614582575b50855f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed016020525f60408120555f526020525f60408120556142b3565b815f528260205260405f2054815f52836020528060405f20555f527f645e039705490088daad89bae25049a34f4a9072d398537b1ab2425f24cbed0160205260405f20555f614544565b6146138473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b600181540190556141a1565b614675835f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205260405f207fffffffffffffffffffffffff00000000000000000000000000000000000000008154169055565b6146bc8573ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930360205260405f2090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff815401905561417f565b8061476d575b156146f8578061415b565b73ffffffffffffffffffffffffffffffffffffffff841661473f57507f7e273289000000000000000000000000000000000000000000000000000000005f5260045260245ffd5b7f177e802f000000000000000000000000000000000000000000000000000000005f5260045260245260445ffd5b508073ffffffffffffffffffffffffffffffffffffffff85161480156147f9575b806146ed57508073ffffffffffffffffffffffffffffffffffffffff6147f2845f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930460205273ffffffffffffffffffffffffffffffffffffffff60405f20541690565b16146146ed565b506148418473ffffffffffffffffffffffffffffffffffffffff165f527f80bb2b638cc20bc4d0a60d66940f3ab4a00c1d7b313497ca82fb0b4ab007930560205260405f2090565b815f5260205260ff60405f20541661478e565b60ff7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005460401c161561488357565b7fd7e6bcf8000000000000000000000000000000000000000000000000000000005f5260045ffd5b906148e857508051156148c057805190602001fd5b7fd6bda275000000000000000000000000000000000000000000000000000000005f5260045ffd5b8151158061493b575b6148f9575090565b73ffffffffffffffffffffffffffffffffffffffff907f9996b315000000000000000000000000000000000000000000000000000000005f521660045260245ffd5b50803b156148f156fea2646970667358221220d87bfabacbe634dfc8146498399032eedbcb6386feaec44018f789df537e86aa64736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
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
