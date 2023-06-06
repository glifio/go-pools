package constants

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/builtin"
)

const MainnetChainID = 314
const CalibnetChainID = 314159
const LocalnetChainID = 31415926

const SecondsInMinute = 60
const EpochsInMinute = SecondsInMinute / builtin.EpochDurationSeconds
const EpochsInDay = builtin.EpochsInDay
const EpochsInWeek = EpochsInDay * 7
const EpochsInYear = EpochsInDay * 365

// a credential is valid for 30 minutes
const CredentialMinutesValid = 30
const CredentialEpochsValid = CredentialMinutesValid * EpochsInMinute

type Method string

// these method names must match the names in the Agent contract in order to get the right function signature
var (
	MethodBorrow      Method = "borrow"
	MethodPay         Method = "pay"
	MethodAddMiner    Method = "addMiner"
	MethodRemoveMiner Method = "removeMiner"
	MethodWithdraw    Method = "withdraw"
	MethodPushFunds   Method = "pushFunds"
	MethodPullFunds   Method = "pullFunds"
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
)

var dialAddr = "https://api.node.glif.io/rpc/v1"
var t_dialAddr = "https://api.calibration.node.glif.io/rpc/v1"

var MAX_UINT256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
var WAD = big.NewInt(1e18)

// the infinity pool is the first pool created, and has a 0 id
// cache this here to not fetch it from the contracts every time we need it
var INFINITY_POOL_ID = big.NewInt(0)
