package bus

import (
	"github.com/kvant-node/core/types"
	"math/big"
)

type FrozenFunds interface {
	AddFrozenFund(uint64, types.Address, types.Pubkey, types.CoinSymbol, *big.Int)
}
