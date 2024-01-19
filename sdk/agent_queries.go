package sdk

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	filtypes "github.com/filecoin-project/lotus/chain/types"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	"github.com/glifio/go-pools/econ"
	"github.com/glifio/go-pools/terminate"
	"github.com/glifio/go-pools/util"
	"github.com/glifio/go-pools/vc"
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

	return agentCaller.Owner(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentAdministrator(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	return agentCaller.Administration(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentDefaulted(ctx context.Context, address common.Address) (bool, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return false, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return false, err
	}

	return agentCaller.Defaulted(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentOperator(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	return agentCaller.Operator(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentRequester(ctx context.Context, address common.Address) (common.Address, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return common.Address{}, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return common.Address{}, err
	}

	return agentCaller.AdoRequestKey(&bind.CallOpts{Context: ctx})
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
	blockNumber *big.Int,
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

	minerCount, err := minerRegCaller.MinersCount(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, id)
	if err != nil {
		return nil, err
	}

	miners := []address.Address{}

	for i := big.NewInt(0); i.Cmp(minerCount) == -1; i.Add(i, big.NewInt(1)) {
		miner, err := minerRegCaller.GetMiner(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, id, i)
		if err != nil {
			return nil, err
		}

		minerAddr, err := address.NewIDAddress(miner)

		miners = append(miners, minerAddr)
	}
	return miners, nil
}

func (q *fevmQueries) AgentLiquidAssets(ctx context.Context, address common.Address, blockNumber *big.Int) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(address, client)
	if err != nil {
		return nil, err
	}

	assets, err := agentCaller.LiquidAssets(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber})
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (q *fevmQueries) AgentPrincipal(ctx context.Context, agentAddr common.Address, blockNumber *big.Int) (*big.Int, error) {
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

	poolIDs, err := poolRegistryCaller.PoolIDs(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, agentID)
	if err != nil {
		return nil, err
	}

	routerCaller, err := abigen.NewRouterCaller(q.router, ethClient)
	if err != nil {
		return nil, err
	}

	principal := big.NewInt(0)
	for _, poolID := range poolIDs {
		account, err := routerCaller.GetAccount(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, agentID, poolID)
		if err != nil {
			return nil, err
		}
		principal.Add(principal, account.Principal)
	}

	return principal, nil
}

func (q *fevmQueries) AgentAccount(ctx context.Context, agentAddr common.Address, poolID *big.Int, blockNumber *big.Int) (abigen.Account, error) {
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

	return routerCaller.GetAccount(&bind.CallOpts{Context: ctx, BlockNumber: blockNumber}, agentID, poolID)
}

func (q *fevmQueries) AgentAddrIDFromRcpt(ctx context.Context, receipt *types.Receipt) (common.Address, *big.Int, error) {
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

func (q *fevmQueries) AgentInterestOwed(ctx context.Context, agentAddr common.Address, tsk *filtypes.TipSet) (*big.Int, error) {
	account, err := q.AgentAccount(ctx, agentAddr, constants.INFINITY_POOL_ID, nil)
	if err != nil {
		log.Printf("Error getting agent account: %v", err)
		return nil, err
	}

	nullishVC, err := vc.NullishVerifiableCredential(*vc.EmptyAgentData())
	if err != nil {
		log.Printf("Error getting nullish VC: %v", err)
		return nil, err
	}

	rate, err := q.InfPoolGetRate(ctx, *nullishVC)
	if err != nil {
		log.Printf("Error getting rate: %v", err)
		return nil, err
	}

	if tsk == nil {
		tsk, err = q.ChainHead(ctx)
		if err != nil {
			log.Printf("Error getting chain head: %v", err)
			return nil, err
		}
	}

	owed := econ.InterestOwed(ctx, account, rate, tsk.Height())

	return owed, nil
}

func (q *fevmQueries) AgentFaultyEpochStart(ctx context.Context, agentAddr common.Address) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(agentAddr, client)
	if err != nil {
		return nil, err
	}

	return agentCaller.FaultySectorStartEpoch(&bind.CallOpts{Context: ctx})
}

func (q *fevmQueries) AgentVersion(ctx context.Context, agentAddr common.Address) (uint8, uint8, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return 0, 0, err
	}
	defer client.Close()

	agentCaller, err := abigen.NewAgentCaller(agentAddr, client)
	if err != nil {
		return 0, 0, err
	}

	agentVersion, err := agentCaller.Version(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, 0, err
	}

	agDeployer, err := q.RouterGetRoute(ctx, constants.RouteAgentDeployer)
	if err != nil {
		return 0, 0, err
	}

	agDeployerCaller, err := abigen.NewAgentDeployerCaller(agDeployer, client)
	if err != nil {
		return 0, 0, err
	}

	deployerVersion, err := agDeployerCaller.Version(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, 0, err
	}

	return agentVersion, deployerVersion, nil
}

func (q *fevmQueries) PreviewAgentTerminationQuick(ctx context.Context, agentAddr common.Address) (terminate.PreviewAgentTerminationSummary, error) {
	agentID, err := q.AgentID(ctx, agentAddr)
	if err != nil {
		return terminate.PreviewAgentTerminationSummary{}, err
	}

	agentCollateralStats, err := terminate.FetchAgentCollateralStats(ctx, agentID)
	if err != nil {
		return terminate.PreviewAgentTerminationSummary{}, err
	}

	availBal := big.NewInt(0)
	initialPledge := big.NewInt(0)
	vestingBal := big.NewInt(0)
	for _, miner := range agentCollateralStats.MinersTerminationStats {
		availBal.Add(availBal, miner.Available)
		initialPledge.Add(initialPledge, miner.Pledged)
		vestingBal.Add(vestingBal, miner.Vesting)
	}

	agentLiquidFIL, err := q.AgentLiquidAssets(ctx, agentAddr, nil)
	if err != nil {
		return terminate.PreviewAgentTerminationSummary{}, err
	}

	return terminate.PreviewAgentTerminationSummary{
		TerminationPenalty: agentCollateralStats.TerminationPenalty,
		InitialPledge:      initialPledge,
		VestingBalance:     vestingBal,
		MinersAvailableBal: availBal,
		AgentAvailableBal:  agentLiquidFIL,
	}, nil
}

var LookbackEpochs abi.ChainEpoch = 3

// PreviewAgentTermination preview terminating all the
// sectors on all the miners for an agent (using sampling and "off-chain"
// calculation) and will return the liquidation value of the agent.
func (q *fevmQueries) PreviewAgentTermination(ctx context.Context, agentAddr common.Address, tipset *ltypes.TipSet) (terminate.PreviewAgentTerminationSummary, error) {
	lapi, closer, err := q.extern.ConnectLotusClient()
	if err != nil {
		return terminate.PreviewAgentTerminationSummary{}, err
	}
	defer closer()

	// if no tipset is found, we use 3 epochs behind chainhead (so we dont get epoch syncronization errors)
	if tipset == nil {
		ch, err := lapi.ChainHead(ctx)
		if err != nil {
			return terminate.PreviewAgentTerminationSummary{}, err
		}

		tipset, err = lapi.ChainGetTipSetByHeight(context.Background(), abi.ChainEpoch(ch.Height()-LookbackEpochs), ltypes.EmptyTSK)
		if err != nil {
			return terminate.PreviewAgentTerminationSummary{}, err
		}
	}

	bigHeight := big.NewInt(int64(tipset.Height()))

	miners, err := q.AgentMiners(ctx, agentAddr, bigHeight)
	if err != nil {
		return terminate.PreviewAgentTerminationSummary{}, err
	}

	minerCount := int64(len(miners))

	tasks := make([]util.TaskFunc, minerCount)
	for i := int64(0); i < minerCount; i++ {
		minerAddr := miners[i]
		tasks[i] = func() (interface{}, error) {
			return terminate.PreviewTerminateSectorsQuick(context.Background(), lapi, minerAddr, tipset)
		}
	}

	results, err := util.Multiread(tasks)
	if err != nil {
		return terminate.PreviewAgentTerminationSummary{}, err
	}

	terminationPenaltyAgg := big.NewInt(0)
	initialPledgeAgg := big.NewInt(0)
	vestingBalanceAgg := big.NewInt(0)
	availableBalanceAgg := big.NewInt(0)

	for _, terminationStats := range results {
		terminationStat := terminationStats.(*terminate.PreviewTerminateSectorsReturn)
		// add the miners termination penalty to the aggregate
		terminationPenaltyAgg = new(big.Int).Add(terminationPenaltyAgg, terminationStat.SectorStats.TerminationPenalty)
		// add the miners bals to their aggregate counterpart
		initialPledgeAgg = new(big.Int).Add(initialPledgeAgg, terminationStat.InitialPledge)
		vestingBalanceAgg = new(big.Int).Add(vestingBalanceAgg, terminationStat.VestingBalance)
		availableBalanceAgg = new(big.Int).Add(availableBalanceAgg, terminationStat.Actor.Balance.Int)
	}

	agentLiquidFIL, err := q.AgentLiquidAssets(ctx, agentAddr, bigHeight)
	if err != nil {
		return terminate.PreviewAgentTerminationSummary{}, err
	}

	return terminate.PreviewAgentTerminationSummary{
		TerminationPenalty: terminationPenaltyAgg,
		InitialPledge:      initialPledgeAgg,
		VestingBalance:     vestingBalanceAgg,
		MinersAvailableBal: availableBalanceAgg,
		AgentAvailableBal:  agentLiquidFIL,
	}, nil
}
