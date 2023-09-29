package sdk

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/filecoin-project/go-address"
	ltypes "github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	token "github.com/glifio/go-pools/jws"
	"github.com/glifio/go-pools/rpc"
	"github.com/glifio/go-pools/util"
	"github.com/glifio/go-pools/vc"
	"github.com/glifio/go-wallet-utils/accounts"
)

func (a *fevmActions) AgentCreate(
	ctx context.Context,
	owner common.Address,
	operator common.Address,
	request common.Address,
	wallet accounts.Wallet,
	account accounts.Account,
	passphrase string,
	proposer address.Address,
	approver address.Address,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	// fromAddr := owner

	/*
		nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
		if err != nil {
			return common.Hash{}, nil, err
		}
	*/

	args := []interface{}{owner, operator, request}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		wallet,
		account,
		passphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nil,
		args,
		abigen.NewAgentFactoryTransactor,
		a.queries.AgentFactory(),
		"Create",
		"Agent Create",
	)
}

func (a *fevmActions) AgentBorrow(
	ctx context.Context,
	agentAddr common.Address,
	poolID *big.Int,
	amount *big.Int,
	ownerWallet accounts.Wallet,
	ownerAccount accounts.Account,
	ownerPassphrase string,
	proposer address.Address,
	approver address.Address,
	requesterKey *ecdsa.PrivateKey,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, address.Undef, amount, constants.MethodBorrow, requesterKey, a.queries)
	if err != nil {
		return common.Hash{}, nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return common.Hash{}, nil, err
	}

	/*
		nonce, err := a.queries.ChainGetNonce(ctx, ownerAccount.EthAccount.Address)
		if err != nil {
			return common.Hash{}, nil, err
		}
	*/

	args := []interface{}{poolID, sc}
	// TODO: this isn't great because we'd rather not get the credential if the amount is too high
	agentData, err := vc.AbiDecodeClaim(sc.Vc.Claim)
	if err != nil {
		return common.Hash{}, nil, err
	}

	maxBorrowNow, err := a.queries.InfPoolAgentMaxBorrow(ctx, agentAddr, agentData)
	if err != nil {
		return common.Hash{}, nil, err
	}

	fmt.Printf("Debug: max borrow now: %0.09f\n", util.ToFIL(maxBorrowNow))

	// if amount.Cmp(maxBorrowNow) > 0 {
	// 	return nil, errors.New("amount exceeds max borrow - run `glif agent preview borrow <amount>` to get more information.")
	// }

	return util.WriteTx(
		ctx,
		lapi,
		client,
		ownerWallet,
		ownerAccount,
		ownerPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nil,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"Borrow",
		"Agent Borrow",
	)
}

func (a *fevmActions) AgentPay(
	ctx context.Context,
	agentAddr common.Address,
	poolID *big.Int,
	amount *big.Int,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
	proposer address.Address,
	approver address.Address,
	requesterKey *ecdsa.PrivateKey,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer closer()

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	jws, err := token.SignJWS(ctx, agentAddr, address.Undef, amount, constants.MethodPay, requesterKey, a.queries)
	if err != nil {
		return common.Hash{}, nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{poolID, sc}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		senderWallet,
		senderAccount,
		senderPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"Pay",
		"Agent Pay",
	)
}

func (a *fevmActions) AgentAddMiner(
	ctx context.Context,
	agentAddr common.Address,
	minerAddr address.Address,
	ownerWallet accounts.Wallet,
	ownerAccount accounts.Account,
	ownerPassphrase string,
	proposer address.Address,
	approver address.Address,
	requesterKey *ecdsa.PrivateKey,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, common.Big0, constants.MethodAddMiner, requesterKey, a.queries)
	if err != nil {
		return common.Hash{}, nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return common.Hash{}, nil, err
	}

	fromAddr := ownerAccount.EthAccount.Address

	nonce, err := a.queries.ChainGetNonce(ctx, fromAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{sc}

	txHash, tx, err := util.WriteTx(
		ctx,
		lapi,
		client,
		ownerWallet,
		ownerAccount,
		ownerPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"AddMiner",
		"Agent Add Miner",
	)
	if err != nil {
		return common.Hash{}, nil, err
	}

	return txHash, tx, nil
}

func (a *fevmActions) AgentRemoveMiner(
	ctx context.Context,
	agentAddr common.Address,
	minerAddr address.Address,
	newOwnerAddr address.Address,
	ownerWallet accounts.Wallet,
	ownerAccount accounts.Account,
	ownerPassphrase string,
	proposer address.Address,
	approver address.Address,
	requesterKey *ecdsa.PrivateKey,
) (common.Hash, *types.Transaction, error) {
	if newOwnerAddr.Protocol() != address.ID {
		lapi, closer, err := a.extern.ConnectLotusClient()
		if err != nil {
			return common.Hash{}, nil, err
		}
		defer closer()

		idAddr, err := lapi.StateLookupID(context.Background(), newOwnerAddr, ltypes.EmptyTSK)
		if err != nil {
			return common.Hash{}, nil, err
		}
		newOwnerAddr = idAddr
	}

	newOwner64, err := address.IDFromAddress(newOwnerAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer closer()

	nonce, err := a.queries.ChainGetNonce(ctx, ownerAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, common.Big0, constants.MethodRemoveMiner, requesterKey, a.queries)
	if err != nil {
		return common.Hash{}, nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{newOwner64, sc}

	txHash, tx, err := util.WriteTx(
		ctx,
		lapi,
		client,
		ownerWallet,
		ownerAccount,
		ownerPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"RemoveMiner",
		"Agent Remove Miner",
	)
	if err != nil {
		return common.Hash{}, nil, err
	}

	return txHash, tx, nil
}

func (a *fevmActions) AgentChangeMinerWorker(
	ctx context.Context,
	agentAddr common.Address,
	minerAddr address.Address,
	workerAddr address.Address,
	controlAddrs []address.Address,
	ownerWallet accounts.Wallet,
	ownerAccount accounts.Account,
	ownerPassphrase string,
	proposer address.Address,
	approver address.Address,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	// convert miner address to ID address
	minerID, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	// convert worker address to ID address
	workerID, err := address.IDFromAddress(workerAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	// convert control addresses to ID addresses
	var controlIDs []uint64
	for _, controlAddr := range controlAddrs {
		controlID, err := address.IDFromAddress(controlAddr)
		if err != nil {
			return common.Hash{}, nil, err
		}
		controlIDs = append(controlIDs, controlID)
	}

	nonce, err := a.queries.ChainGetNonce(ctx, ownerAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{minerID, workerID, controlIDs}

	txHash, tx, err := util.WriteTx(
		ctx,
		lapi,
		client,
		ownerWallet,
		ownerAccount,
		ownerPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"ChangeMinerWorker",
		"Agent Change Miner Worker",
	)
	if err != nil {
		return common.Hash{}, nil, err
	}

	return txHash, tx, nil
}

func (a *fevmActions) AgentConfirmMinerWorkerChange(
	ctx context.Context,
	agentAddr common.Address,
	minerAddr address.Address,
	ownerWallet accounts.Wallet,
	ownerAccount accounts.Account,
	ownerPassphrase string,
	proposer address.Address,
	approver address.Address,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	minerU64, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, ownerAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{minerU64}

	txHash, tx, err := util.WriteTx(
		ctx,
		lapi,
		client,
		ownerWallet,
		ownerAccount,
		ownerPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"ConfirmChangeMinerWorker",
		"Agent Confirm Miner Worker Change",
	)
	if err != nil {
		return common.Hash{}, nil, err
	}

	return txHash, tx, nil
}

// AgentPullFunds pulls funds from the agent to a miner
func (a *fevmActions) AgentPullFunds(
	ctx context.Context,
	agentAddr common.Address,
	amount *big.Int,
	minerAddr address.Address,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
	proposer address.Address,
	approver address.Address,
	requesterKey *ecdsa.PrivateKey,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	agentID, err := a.queries.AgentID(ctx, agentAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	minerRegistryCaller, err := abigen.NewMinerRegistryCaller(a.queries.MinerRegistry(), client)
	if err != nil {
		return common.Hash{}, nil, err
	}

	minerU64, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	registered, err := minerRegistryCaller.MinerRegistered(nil, agentID, minerU64)
	if err != nil {
		return common.Hash{}, nil, err
	}

	if !registered {
		return common.Hash{}, nil, errors.New("Miner not registered with agent. Be sure to call `agent add-miner` first before pulling funds.")
	}

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, amount, constants.MethodPullFunds, requesterKey, a.queries)
	if err != nil {
		return common.Hash{}, nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return common.Hash{}, nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{sc}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		senderWallet,
		senderAccount,
		senderPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"PullFunds",
		"Agent Pull Funds",
	)
}

// AgentPushFunds pushes funds from the agent to a miner
func (a *fevmActions) AgentPushFunds(
	ctx context.Context,
	agentAddr common.Address,
	amount *big.Int,
	minerAddr address.Address,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
	proposer address.Address,
	approver address.Address,
	requesterKey *ecdsa.PrivateKey,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	agentID, err := a.queries.AgentID(ctx, agentAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	minerRegistryCaller, err := abigen.NewMinerRegistryCaller(a.queries.MinerRegistry(), client)
	if err != nil {
		return common.Hash{}, nil, err
	}

	minerU64, err := address.IDFromAddress(minerAddr)
	if err != nil {
		return common.Hash{}, nil, err
	}

	registered, err := minerRegistryCaller.MinerRegistered(nil, agentID, minerU64)
	if err != nil {
		return common.Hash{}, nil, err
	}

	if !registered {
		return common.Hash{}, nil, errors.New("Miner not registered with agent. Be sure to call `agent add-miner` first before pushing funds.")
	}

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer closer()

	jws, err := token.SignJWS(ctx, agentAddr, minerAddr, amount, constants.MethodPushFunds, requesterKey, a.queries)
	if err != nil {
		return common.Hash{}, nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return common.Hash{}, nil, err
	}

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{sc}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		senderWallet,
		senderAccount,
		senderPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"PushFunds",
		"Agent Push Funds",
	)
}

func (a *fevmActions) AgentWithdraw(
	ctx context.Context,
	agentAddr common.Address,
	receiver common.Address,
	amount *big.Int,
	ownerWallet accounts.Wallet,
	ownerAccount accounts.Account,
	ownerPassphrase string,
	proposer address.Address,
	approver address.Address,
	requesterKey *ecdsa.PrivateKey,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	closer, err := a.extern.ConnectAdoClient(ctx)
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer closer()

	nonce, err := a.queries.ChainGetNonce(ctx, ownerAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	jws, err := token.SignJWS(ctx, agentAddr, address.Undef, amount, constants.MethodWithdraw, requesterKey, a.queries)
	if err != nil {
		return common.Hash{}, nil, err
	}

	sc, err := rpc.ADOClient.SignCredential(ctx, jws)
	if err != nil {
		return common.Hash{}, nil, err
	}

	args := []interface{}{receiver, sc}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		ownerWallet,
		ownerAccount,
		ownerPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		args,
		abigen.NewAgentTransactor,
		agentAddr,
		"Withdraw",
		"Agent Withdraw",
	)
}

func (a *fevmActions) AgentRefreshRoutes(
	ctx context.Context,
	agentAddr common.Address,
	senderWallet accounts.Wallet,
	senderAccount accounts.Account,
	senderPassphrase string,
	proposer address.Address,
	approver address.Address,
) (common.Hash, *types.Transaction, error) {
	lapi, lcloser, err := a.extern.ConnectLotusClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer lcloser()

	client, err := a.extern.ConnectEthClient()
	if err != nil {
		return common.Hash{}, nil, err
	}
	defer client.Close()

	nonce, err := a.queries.ChainGetNonce(ctx, senderAccount.EthAccount.Address)
	if err != nil {
		return common.Hash{}, nil, err
	}

	return util.WriteTx(
		ctx,
		lapi,
		client,
		senderWallet,
		senderAccount,
		senderPassphrase,
		proposer,
		approver,
		a.queries.ChainID(),
		common.Big0,
		nonce,
		[]interface{}{},
		abigen.NewAgentTransactor,
		agentAddr,
		"RefreshRoutes",
		"Agent RefreshRoutes",
	)
}
