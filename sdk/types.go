package sdk

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/glifio/go-pools/types"
)

type fevmConnection struct {
	query  types.FEVMQueries
	act    types.FEVMActions
	extern types.FEVMExtern
	types.PoolsSDK
}

type fevmExtern struct {
	dialAddr     string
	token        string
	adoAddr      string
	adoNamespace string
	types.FEVMExtern
}

type fevmQueries struct {
	extern        types.FEVMExtern
	agentPolice   common.Address
	minerRegistry common.Address
	router        common.Address
	poolRegistry  common.Address
	agentFactory  common.Address
	iFIL          common.Address
	infinityPool  common.Address
	simpleRamp    common.Address
	wFIL          common.Address
	chainID       *big.Int
	types.FEVMQueries
}

type fevmActions struct {
	extern  types.FEVMExtern
	queries types.FEVMQueries
	types.FEVMActions
}
