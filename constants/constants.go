package constants

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/builtin"
)

const MainnetChainID = 314
const CalibnetChainID = 314159
const LocalnetChainID = 31415926
const AnvilChainID = 31337

const SecondsInMinute = 60
const EpochsInMinute = SecondsInMinute / builtin.EpochDurationSeconds
const EpochsInDay = builtin.EpochsInDay
const EpochsInWeek = EpochsInDay * 7
const EpochsInYear = EpochsInDay * 365

var DefaultWindow = EpochsInWeek * 3
var RepeatBorrowEpochTolerance = EpochsInDay

// a credential is valid for 5 minutes
const CredentialMinutesValid = 5
const CredentialEpochsValid = CredentialMinutesValid * EpochsInMinute

// the ChainHeadLookbackEpochs epochs is the protocol wide parameter that is subtracted from ChainHead height to look at real time on chain info
const ChainHeadLookbackEpochs = 3

type Method string

// these method names must match the names in the Agent contract in order to get the right function signature
var (
	MethodBorrow        Method = "borrow"
	MethodPay           Method = "pay"
	MethodAddMiner      Method = "addMiner"
	MethodRemoveMiner   Method = "removeMiner"
	MethodWithdraw      Method = "withdraw"
	MethodPushFunds     Method = "pushFunds"
	MethodPullFunds     Method = "pullFunds"
	MethodSetRecovered  Method = "setRecovered"
	MethodDefaultDTL    Method = "setAgentDefaultDTL"
	MethodPlusDowngrade Method = "downgrade"
)

type Route string

var (
	RouteAgentPolice   Route = "ROUTER_AGENT_POLICE"
	RouteAgentFactory  Route = "ROUTER_AGENT_FACTORY"
	RoutePoolRegistry  Route = "ROUTER_POOL_REGISTRY"
	RouteMinerRegistry Route = "ROUTER_MINER_REGISTRY"
	RouteWFIL          Route = "ROUTER_WFIL_TOKEN"
	RouteVCIssuer      Route = "ROUTER_VC_ISSUER"
	RouteCredParser    Route = "ROUTER_CRED_PARSER"
	RouteAgentDeployer Route = "ROUTER_AGENT_DEPLOYER"
	RouteInfinityPool  Route = "ROUTER_INFINITY_POOL"
)

var MAX_UINT256 = new(big.Int).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil), big.NewInt(1))
var WAD = big.NewInt(1e18)

// the infinity pool is the first pool created, and has a 0 id
// cache this here to not fetch it from the contracts every time we need it
var INFINITY_POOL_ID = big.NewInt(0)

var CHUNKSIZE = big.NewInt(EpochsInDay)

var FAULTY_SECTOR_TOLERANCE = big.NewFloat(0.001)

var MAX_BORROW_DTL = big.NewInt(75e16)
var LIQUIDATION_DTL = big.NewInt(85e16)

const (
	InactiveTier = iota
	BronzeTier
	SilverTier
	GoldTier
)

// Tier DTL mapping for dynamic borrowing limits
var TierDTL = map[uint8]*big.Int{
	InactiveTier: big.NewInt(75e16),  // 75%
	BronzeTier:   big.NewInt(80e16),  // 80%
	SilverTier:   big.NewInt(857e15), // 85.7%
	GoldTier:     big.NewInt(90e16),  // 90%
}

var PLUS_TIERS = uint8(4)
