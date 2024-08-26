package econ

import (
	"context"
	"math"
	"math/big"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/actors/adt"
	"github.com/filecoin-project/lotus/chain/actors/builtin/reward"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/glifio/go-pools/abigen"
	"github.com/glifio/go-pools/util"
	cbor "github.com/ipfs/go-ipld-cbor"
)

// @dev rates have 2 WADs worth of precision (1e36) to maintain per epoch rate precision
func InterestOwed(ctx context.Context, account abigen.Account, rate *big.Int, chainHeadHeight abi.ChainEpoch) *big.Int {
	currentEpoch := new(big.Int).SetUint64(uint64(chainHeadHeight))

	interestOwed := big.NewInt(0)
	if account.EpochsPaid.Cmp(currentEpoch) == -1 {
		// Compute the number of epochs that are owed to get current
		epochsToPay := new(big.Int).Sub(currentEpoch, account.EpochsPaid)

		// Multiply the rate by the principal to get the per epoch interest rate, now has an extra WAD of precision after multiplication
		interestPerEpoch := new(big.Int).Mul(account.Principal, rate)

		// div out the extra wad from the multiplication above
		interestPerEpoch.Div(interestPerEpoch, big.NewInt(1e18))

		// Compute the total interest owed by multiplying how many epochs to pay, by the per epoch interest payment
		interestOwed = new(big.Int).Mul(interestPerEpoch, epochsToPay)

		// div out the extra wad in interestOwed from the WAD embedded in `rate` for precision
		interestOwed.Div(interestOwed, big.NewInt(1e18))
	}

	return interestOwed
}

func ComputeEDR(ctx context.Context, minerAddr address.Address, ts *types.TipSet, api *api.FullNodeStruct) (*big.Int, error) {
	pow, err := api.StateMinerPower(ctx, minerAddr, ts.Key())
	if err != nil {
		return nil, err
	}

	if pow.HasMinPower {
		// get the block rewards for this epoch
		reward, err := getBlockReward(ctx, api, ts)
		if err != nil {
			return nil, err
		}

		winRatio := new(big.Rat).SetFrac(
			types.BigMul(pow.MinerPower.QualityAdjPower, types.NewInt(build.BlocksPerEpoch)).Int,
			pow.TotalPower.QualityAdjPower.Int,
		)

		if winRatioFloat, _ := winRatio.Float64(); winRatioFloat > 0 {

			// if the corresponding poisson distribution isn't infinitely small then
			// throw it into the mix as well, accounting for multi-wins
			winRationWithPoissonFloat := -math.Expm1(-winRatioFloat)
			winRationWithPoisson := new(big.Rat).SetFloat64(winRationWithPoissonFloat)
			if winRationWithPoisson != nil {
				winRatio = winRationWithPoisson
			}

			daily, _ := new(big.Rat).Mul(
				winRatio,
				new(big.Rat).SetInt64(builtin.EpochsInDay),
			).Float64()

			dailyFloat := new(big.Float).SetFloat64(daily)

			rewardFloat := new(big.Float).SetInt(reward)

			totalRewardsBigFloat := new(big.Float).Mul(dailyFloat, rewardFloat)

			edr, _ := totalRewardsBigFloat.Int(nil)

			return edr, nil
		}
	}

	return big.NewInt(0), nil
}

func getBlockReward(ctx context.Context, api *api.FullNodeStruct, ts *types.TipSet) (*big.Int, error) {
	ract, err := api.StateGetActor(ctx, reward.Address, ts.Key())
	if err != nil {
		return big.NewInt(0), err
	}

	tbsRew := util.NewTieredBlockstore(api)

	rst, err := reward.Load(adt.WrapStore(ctx, cbor.NewCborStore(tbsRew)), ract)
	if err != nil {
		return big.NewInt(0), err
	}

	epochReward, err := rst.ThisEpochReward()
	if err != nil {
		return big.NewInt(0), err
	}

	// Divide by expected leaders per epoch to get block reward
	blockReward := types.BigDiv(epochReward, types.NewInt(uint64(builtin.ExpectedLeadersPerEpoch)))

	return blockReward.Int, nil
}
