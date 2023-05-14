package fevm

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/glifio/go-pools/types"
)

func InitFEVMConnection(
	ctx context.Context,
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
) (types.FEVMConnection, error) {
	client, err := ethclient.Dial(dialAddr)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	chainID, err := client.ChainID(ctx)
	if err != nil {
		return nil, err
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

	return &fevmConnection{
		query:  fevmQueries,
		act:    fevmActions,
		extern: extern,
	}, nil
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
