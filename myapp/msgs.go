package kvstore

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgCreateKVPair struct {
	Creator sdk.AccAddress `json:"creator"`
	Key     string         `json:"key"`
	Value   string         `json:"value"`
}
