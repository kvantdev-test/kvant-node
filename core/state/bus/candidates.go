package bus

import (
	"github.com/kvant-node/core/types"
	"math/big"
)

type Candidates interface {
	GetStakes(types.Pubkey) []Stake
	Punish(uint64, types.TmAddress) *big.Int
	GetCandidate(types.Pubkey) *Candidate
}

type Stake struct {
	Owner    types.Address
	Value    *big.Int
	Coin     types.CoinSymbol
	BipValue *big.Int
}

type Candidate struct {
	PubKey        types.Pubkey
	RewardAddress types.Address
	OwnerAddress  types.Address
	Commission    uint
	Status        byte
}
