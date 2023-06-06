package sdk

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (q *fevmQueries) StateWaitTx(ctx context.Context, hash common.Hash, ch chan *types.Receipt) {
	client, err := q.extern.ConnectEthClient()
	if err != nil {
		return
	}
	defer client.Close()

	for {
		time.Sleep(time.Millisecond * 5000)

		tx, err := client.TransactionReceipt(ctx, hash)
		if err == nil && tx != nil {
			ch <- tx
			return
		}
	}
}

func (q *fevmQueries) StateWaitReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	eapi, err := q.extern.ConnectEthClient()
	if err != nil {
		return nil, err
	}
	defer eapi.Close()

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	// wait for 1 confirmations before getting the receipt
	err = q.StateWaitForConfirmations(ctx, 1)
	if err != nil {
		return nil, err
	}

	ch := make(chan *types.Receipt)
	go q.StateWaitTx(ctx, hash, ch)

	var receipt *types.Receipt
	select {
	case <-ctx.Done():
		s.Stop()
		return nil, fmt.Errorf("timed out waiting for transaction")
	case receipt = <-ch:
		msg := fmt.Sprintf(" Receipt received at block %v.\n", receipt.BlockNumber)
		s.Suffix = msg
	}

	if receipt == nil {
		return nil, fmt.Errorf("transaction receipt not found")
	}

	if receipt.Status != 1 {
		return nil, fmt.Errorf("transaction failed: %v", receipt.Status)
	}

	return receipt, nil
}

func (q *fevmQueries) StateWaitNextTick(ctx context.Context, currentHeight *big.Int) error {
	eapi, err := q.extern.ConnectEthClient()
	if err != nil {
		return err
	}
	defer eapi.Close()

	target := currentHeight.Uint64() + 1
	for {
		time.Sleep(time.Millisecond * 5000)

		b, err := eapi.BlockNumber(ctx)
		if err != nil {
			return err
		}

		if b >= target {
			return nil
		}
	}
}

func (q *fevmQueries) StateWaitForConfirmations(ctx context.Context, confirmations uint64) error {
	eapi, err := q.extern.ConnectEthClient()
	if err != nil {
		return err
	}
	defer eapi.Close()

	b, err := eapi.BlockNumber(ctx)
	if err != nil {
		return err
	}

	target := b + confirmations
	for {
		time.Sleep(time.Millisecond * 5000)

		b, err := eapi.BlockNumber(ctx)
		if err != nil {
			return err
		}

		if b >= target {
			return nil
		}
	}
}
