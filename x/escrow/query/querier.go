package query

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/spheronFdn/akash-node/x/escrow/keeper"
)

func NewQuerier(keeper keeper.Keeper, cdc *codec.LegacyAmino) sdk.Querier {
	return nil
}
