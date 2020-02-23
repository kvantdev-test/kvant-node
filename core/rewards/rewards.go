package rewards

import (
	"github.com/kvant-node/core/types"
	"github.com/kvant-node/helpers"
	"math/big"
)

const lastBlock = 100000000
const firstReward = 999
const lastReward = 10

var startHeight uint64 = 0
var BeforeGenesis = big.NewInt(0)

func GetRewardForBlock(blockHeight uint64) *big.Int {
	blockHeight += startHeight

	if blockHeight > lastBlock {
		return big.NewInt(0)
	}

	if blockHeight == lastBlock {
		return helpers.BipToPip(big.NewInt(lastReward))
	}

	reward := big.NewInt(firstReward)
	reward.Sub(reward, big.NewInt(int64(blockHeight/200000)))

	if reward.Cmp(types.Big0) < 1 {
		return helpers.BipToPip(big.NewInt(1))
	}

	return helpers.BipToPip(reward)
}

func SetStartHeight(sHeight uint64) {
	for i := uint64(1); i <= sHeight; i++ {
		BeforeGenesis.Add(BeforeGenesis, GetRewardForBlock(i))
	}

	startHeight = sHeight
}
