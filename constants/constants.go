package constants

import (
	"math/big"

	"github.com/filecoin-project/go-state-types/builtin"
)

const SecondsInMinute = 60
const EpochsInMinute = SecondsInMinute / builtin.EpochDurationSeconds

// a credential is valid for 30 minutes
const CredentialMinutesValid = 30
const CredentialEpochsValid = CredentialMinutesValid * EpochsInMinute

var (
	MethodBorrow      = "borrow"
	MethodPay         = "pay"
	MethodAddMiner    = "addMiner"
	MethodRemoveMiner = "removeMiner"
	MethodWithdraw    = "withdraw"
	MethodPushFunds   = "pushFunds"
	MethodPullFunds   = "pullFunds"
)

var MAX_UINT256 = new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
var WAD = big.NewInt(1e18)

// the infinity pool is the first pool created, and has a 0 id
// cache this here to not fetch it from the contracts every time we need it
var INFINITY_POOL_ID = big.NewInt(0)
