package transaction

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/kvant-node/core/code"
	"github.com/kvant-node/core/commissions"
	"github.com/kvant-node/core/state"
	"github.com/kvant-node/core/types"
	"github.com/kvant-node/formula"
	"github.com/tendermint/tendermint/libs/kv"
	"math/big"
)

type CandidateTx interface {
	GetPubKey() types.Pubkey
}

type EditCandidateData struct {
	PubKey        types.Pubkey
	RewardAddress types.Address
	OwnerAddress  types.Address
}

func (data EditCandidateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		PubKey        string `json:"pub_key"`
		RewardAddress string `json:"reward_address"`
		OwnerAddress  string `json:"owner_address"`
	}{
		PubKey:        data.PubKey.String(),
		RewardAddress: data.RewardAddress.String(),
		OwnerAddress:  data.OwnerAddress.String(),
	})
}

func (data EditCandidateData) GetPubKey() types.Pubkey {
	return data.PubKey
}

func (data EditCandidateData) TotalSpend(tx *Transaction, context *state.State) (TotalSpends, []Conversion, *big.Int, *Response) {
	panic("implement me")
}

func (data EditCandidateData) BasicCheck(tx *Transaction, context *state.State) *Response {
	return checkCandidateOwnership(data, tx, context)
}

func (data EditCandidateData) String() string {
	return fmt.Sprintf("EDIT CANDIDATE pubkey: %x",
		data.PubKey)
}

func (data EditCandidateData) Gas() int64 {
	return commissions.EditCandidate
}

func (data EditCandidateData) Run(tx *Transaction, context *state.State, isCheck bool, rewardPool *big.Int, currentBlock uint64) Response {
	sender, _ := tx.Sender()

	response := data.BasicCheck(tx, context)
	if response != nil {
		return *response
	}

	commissionInBaseCoin := tx.CommissionInBaseCoin()
	commission := big.NewInt(0).Set(commissionInBaseCoin)

	if !tx.GasCoin.IsBaseCoin() {
		coin := context.Coins.GetCoin(tx.GasCoin)

		errResp := CheckReserveUnderflow(coin, commissionInBaseCoin)
		if errResp != nil {
			return *errResp
		}

		if coin.Reserve().Cmp(commissionInBaseCoin) < 0 {
			return Response{
				Code: code.CoinReserveNotSufficient,
				Log:  fmt.Sprintf("Coin reserve balance is not sufficient for transaction. Has: %s, required %s", coin.Reserve().String(), commissionInBaseCoin.String()),
				Info: EncodeError(map[string]string{
					"has_reserve": coin.Reserve().String(),
					"commission":  commissionInBaseCoin.String(),
					"gas_coin":    coin.CName,
				}),
			}
		}

		commission = formula.CalculateSaleAmount(coin.Volume(), coin.Reserve(), coin.Crr(), commissionInBaseCoin)
	}

	if context.Accounts.GetBalance(sender, tx.GasCoin).Cmp(commission) < 0 {
		return Response{
			Code: code.InsufficientFunds,
			Log:  fmt.Sprintf("Insufficient funds for sender account: %s. Wanted %s %s", sender.String(), commission.String(), tx.GasCoin),
			Info: EncodeError(map[string]string{
				"sender":       sender.String(),
				"needed_value": commission.String(),
				"gas_coin":     fmt.Sprintf("%s", tx.GasCoin),
			}),
		}
	}

	if !isCheck {
		rewardPool.Add(rewardPool, commissionInBaseCoin)

		context.Coins.SubReserve(tx.GasCoin, commissionInBaseCoin)
		context.Coins.SubVolume(tx.GasCoin, commission)

		context.Accounts.SubBalance(sender, tx.GasCoin, commission)
		context.Candidates.Edit(data.PubKey, data.RewardAddress, data.OwnerAddress)
		context.Accounts.SetNonce(sender, tx.Nonce)
	}

	tags := kv.Pairs{
		kv.Pair{Key: []byte("tx.type"), Value: []byte(hex.EncodeToString([]byte{byte(TypeEditCandidate)}))},
		kv.Pair{Key: []byte("tx.from"), Value: []byte(hex.EncodeToString(sender[:]))},
	}

	return Response{
		Code:      code.OK,
		GasUsed:   tx.Gas(),
		GasWanted: tx.Gas(),
		Tags:      tags,
	}
}

func checkCandidateOwnership(data CandidateTx, tx *Transaction, context *state.State) *Response {
	if !context.Candidates.Exists(data.GetPubKey()) {
		return &Response{
			Code: code.CandidateNotFound,
			Log:  fmt.Sprintf("Candidate with such public key (%s) not found", data.GetPubKey().String()),
			Info: EncodeError(map[string]string{
				"public_key": data.GetPubKey().String(),
			}),
		}
	}

	owner := context.Candidates.GetCandidateOwner(data.GetPubKey())
	sender, _ := tx.Sender()
	if owner != sender {
		return &Response{
			Code: code.IsNotOwnerOfCandidate,
			Log:  fmt.Sprintf("Sender is not an owner of a candidate")}
	}

	return nil
}
