package sdk

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/glifio/go-pools/types"
)

func Init(
	ctx context.Context,
	sdk *types.PoolsSDK,
	agentPolice common.Address,
	minerRegistry common.Address,
	router common.Address,
	poolRegistry common.Address,
	agentFactory common.Address,
	iFIL common.Address,
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

	extern := &fevmExtern{
		dialAddr:     dialAddr,
		token:        token,
		adoAddr:      adoAddr,
		adoNamespace: adoNamespace,
	}

	fevmQueries := &fevmQueries{
		router:        router,
		iFIL:          iFIL,
		infinityPool:  infinityPool,
		agentFactory:  agentFactory,
		poolRegistry:  poolRegistry,
		minerRegistry: minerRegistry,
		agentPolice:   agentPolice,
		chainID:       chainID,
		extern:        extern,
	}

	fevmActions := &fevmActions{extern: extern, queries: fevmQueries}

	*sdk = &fevmConnection{
		query:  fevmQueries,
		act:    fevmActions,
		extern: extern,
	}

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
