package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/glifio/go-pools/types"
)

func InitFEVMConnection(
	agentPolice common.Address,
	minerRegistry common.Address,
	router common.Address,
	poolRegistry common.Address,
	agentFactory common.Address,
	iFIL common.Address,
	wFIL common.Address,
	infinityPool common.Address,
	adoAddr string,
	adoNamespace string,
	dialAddr string,
	token string,
	chainID *big.Int,
) *fevmConnection {
	extern := &fevmExtern{
		dialAddr:     dialAddr,
		token:        token,
		adoAddr:      adoAddr,
		adoNamespace: adoNamespace,
	}

	fevmQueries := &fevmQueries{
		router:        router,
		iFIL:          iFIL,
		wFIL:          wFIL,
		infinityPool:  infinityPool,
		agentFactory:  agentFactory,
		poolRegistry:  poolRegistry,
		minerRegistry: minerRegistry,
		agentPolice:   agentPolice,
		chainID:       chainID,
		extern:        extern,
	}

	fevmActions := &fevmActions{extern: extern, queries: fevmQueries}

	return &fevmConnection{
		query:  fevmQueries,
		act:    fevmActions,
		extern: extern,
	}
}

func Init(
	ctx context.Context,
	sdk *types.PoolsSDK,
	agentPolice common.Address,
	minerRegistry common.Address,
	router common.Address,
	poolRegistry common.Address,
	agentFactory common.Address,
	iFIL common.Address,
	wFIL common.Address,
	infinityPool common.Address,
	adoAddr string,
	adoNamespace string,
	dialAddr string,
	token string,
) error {
	client, err := ethclient.Dial(dialAddr)
	if err != nil {
		return err
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return err
	}

	*sdk = InitFEVMConnection(
		agentPolice,
		minerRegistry,
		router,
		poolRegistry,
		agentFactory,
		iFIL,
		wFIL,
		infinityPool,
		adoAddr,
		adoNamespace,
		dialAddr,
		token,
		chainID,
	)

	return nil
}

// Implement the FEVMConnection interface
func (c *fevmConnection) Query() types.FEVMQueries {
	return c.query
}

func (c *fevmConnection) Act() types.FEVMActions {
	return c.act
}

func (c *fevmConnection) Extern() types.FEVMExtern {
	return c.extern
}
