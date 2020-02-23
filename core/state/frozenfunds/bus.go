package frozenfunds

import (
	"github.com/kvant-node/core/types"
	"math/big"
)

type Bus struct {
	frozenfunds *FrozenFunds
}

func (b *Bus) AddFrozenFund(height uint64, address types.Address, pubkey types.Pubkey, coin types.CoinSymbol, value *big.Int) {
	b.frozenfunds.AddFund(height, address, pubkey, coin, value)
}

func NewBus(frozenfunds *FrozenFunds) *Bus {
	return &Bus{frozenfunds: frozenfunds}
}
