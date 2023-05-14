package fevm

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/rpc"
)

func (q *fevmQueries) AgentID(ctx context.Context, address common.Address) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return nil, err
	}

	agentID, err := agentCaller.Id(nil)
	if err != nil {
		return nil, err
	}

	return agentID, nil
}

func (q *fevmQueries) AgentOwner(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	owner, err := agentCaller.Owner(nil)
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}

func (q *fevmQueries) AgentIsValid(ctx context.Context, address common.Address) (bool, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return false, err
	}
	defer client.Close()

	agentFactoryCaller, err := abigen.NewAgentFactoryCaller(q.agentFactory, client)
	if err != nil {
		return false, err
	}

	isAgent, err := agentFactoryCaller.IsAgent(&bind.CallOpts{Context: ctx}, address)
	if err != nil {
		return false, err
	}

	return isAgent, nil
}

func (q *fevmQueries) AgentMiners(
	ctx context.Context,
	agentAddr common.Address,
) ([]address.Address, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	id, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	minerRegCaller, err := abigen.NewMinerRegistryCaller(q.minerRegistry, ethClient)
	if err != nil {
		return nil, err
	}

	minerCount, err := minerRegCaller.MinersCount(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}

	miners := []address.Address{}

	for i := big.NewInt(0); i.Cmp(minerCount) == -1; i.Add(i, big.NewInt(1)) {
		miner, err := minerRegCaller.GetMiner(&bind.CallOpts{Context: ctx}, id, i)
		if err != nil {
			return nil, err
		}

		minerAddr, err := address.NewIDAddress(miner)

		miners = append(miners, minerAddr)
	}
	return miners, nil
}

func (q *fevmQueries) AgentLiquidAssets(ctx context.Context, address common.Address) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return nil, err
	}

	assets, err := agentCaller.LiquidAssets(nil)
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (q *fevmQueries) AgentTotalPrincipal(ctx context.Context, agentAddr common.Address) (*big.Int, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer ethClient.Close()

	agentID, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return nil, err
	}

	poolRegistryCaller, err := abigen.NewPoolRegistryCaller(q.poolRegistry, ethClient)
	if err != nil {
		return nil, err
	}

	poolIDs, err := poolRegistryCaller.PoolIDs(&bind.CallOpts{Context: ctx}, agentID)
	if err != nil {
		return nil, err
	}

	routerCaller, err := abigen.NewRouterCaller(q.router, ethClient)
	if err != nil {
		return nil, err
	}

	principal := big.NewInt(0)
	for _, poolID := range poolIDs {
		account, err := routerCaller.GetAccount(&bind.CallOpts{Context: ctx}, agentID, poolID)
		if err != nil {
			return nil, err
		}
		principal.Add(principal, account.Principal)
	}

	return principal, nil
}

func (q *fevmQueries) AgentAccount(ctx context.Context, agentAddr common.Address, poolID *big.Int) (abigen.Account, error) {
	ethClient, err := q.extern.ConnectEthClient()
	if err != nil {
		return abigen.Account{}, err
	}
	defer ethClient.Close()

	agentID, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return abigen.Account{}, err
	}

	routerCaller, err := abigen.NewRouterCaller(q.router, ethClient)
	if err != nil {
		return abigen.Account{}, err
	}

	return routerCaller.GetAccount(&bind.CallOpts{Context: ctx}, agentID, poolID)
}

func (q *fevmQueries) AgentAddrID(ctx context.Context, receipt *types.Receipt) (common.Address, *big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, nil, err
	}
	defer client.Close()

	agentFactoryABI, err := abigen.AgentFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, err
	}

	agentFactoryFilterer, err := abigen.NewAgentFactoryFilterer(q.agentFactory, client)
	if err != nil {
		return common.Address{}, nil, err
	}

	var agentID *big.Int
	var agentAddr common.Address

	for _, l := range receipt.Logs {
		event, err := agentFactoryABI.EventByID(l.Topics[0])
		if err != nil {
			return common.Address{}, nil, err
		}
		if event.Name == "CreateAgent" {
			createAgentEvent, err := agentFactoryFilterer.ParseCreateAgent(*l)
			if err != nil {
				return common.Address{}, nil, err
			}
			agentAddr = createAgentEvent.Agent
			agentID = createAgentEvent.AgentID
		}
	}

	return agentAddr, agentID, nil
}

func (q *fevmQueries) AgentOwes(ctx context.Context, agentAddr common.Address) (*big.Int, *big.Int, error) {
	closer, err := q.extern.ConnectAdoClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	defer closer()

	agentOwed, err := rpc.ADOClient.AmountOwed(ctx, agentAddr, constants.INFINITY_POOL_ID)
	if err != nil {
		return nil, nil, err
	}

	return agentOwed.AmountOwed, agentOwed.Gcred, nil
}
