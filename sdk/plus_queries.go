package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/constants"
	poolstypes "github.com/glifio/go-pools/types"
)

func (q *fevmQueries) PlusTokenIDFromRcpt(ctx context.Context, receipt *types.Receipt) (*big.Int, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plusABI, err := abigen.PlusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}

	plusFilterer, err := abigen.NewPlusFilterer(q.plus, client)
	if err != nil {
		return nil, err
	}

	var tokenID *big.Int

	for _, l := range receipt.Logs {
		event, err := plusABI.EventByID(l.Topics[0])
		if err != nil {
			return nil, err
		}
		if event.Name == "CardMinted" {
			cardMintedEvent, err := plusFilterer.ParseCardMinted(*l)
			if err != nil {
				return nil, err
			}
			tokenID = cardMintedEvent.TokenId
			break
		}
	}

	return tokenID, nil
}

func (q *fevmQueries) PlusInfo(ctx context.Context, tokenID *big.Int, blockNumber *big.Int) (*poolstypes.PlusInfo, error) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer client.Close()

	plus, err := abigen.NewPlusCaller(q.plus, client)
	if err != nil {
		return nil, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	agentID, err := plus.TokenIdToAgentId(opts, tokenID)
	if err != nil {
		return nil, err
	}

	filCashbackEarned, err := plus.TokenIdToFilCashbackEarned(opts, tokenID)
	if err != nil {
		return nil, err
	}

	glfVaultBalance, err := plus.TokenIdToGlfVaultBalance(opts, tokenID)
	if err != nil {
		return nil, err
	}

	lastTierSwitchTimestamp, err := plus.TokenIdToLastTierSwitchTimestamp(opts, tokenID)
	if err != nil {
		return nil, err
	}

	personalCashBackPercent, err := plus.TokenIdToPersonalCashBackPercent(opts, tokenID)
	if err != nil {
		return nil, err
	}

	tier, err := plus.TokenIdToTier(opts, tokenID)
	if err != nil {
		return nil, err
	}

	return &poolstypes.PlusInfo{
		AgentID:                 agentID,
		FilCashbackEarned:       filCashbackEarned,
		GLFVaultBalance:         glfVaultBalance,
		LastTierSwitchTimestamp: lastTierSwitchTimestamp,
		PersonalCashBackPercent: personalCashBackPercent,
		Tier:                    tier,
	}, nil
}

func (q *fevmQueries) PlusTierInfo(ctx context.Context, blockNumber *big.Int) ([]abigen.TierInfo, error) {
	tiers := make([]abigen.TierInfo, 0, constants.PLUS_TIERS)

	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return tiers, err
	}
	defer client.Close()

	plus, err := abigen.NewPlusCaller(q.plus, client)
	if err != nil {
		return tiers, err
	}

	opts := &bind.CallOpts{Context: ctx, BlockNumber: blockNumber}

	for tier := uint8(0); tier < constants.PLUS_TIERS; tier++ {
		tierInfo, err := plus.TierToTierInfo(opts, tier)
		if err != nil {
			return nil, err
		}
		tiers = append(tiers, tierInfo)
	}

	return tiers, nil
}
