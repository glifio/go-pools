package sdk

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/glifio/go-pools/abigen"
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
