package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/giansalex/nameservice/x/nameservice/types"
)

// MinPrice
func (k Keeper) MinPrice(ctx sdk.Context) (res sdk.Coins) {
	k.paramstore.Get(ctx, types.KeyMinPrice, &res)
	return
}

// Get all parameteras as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MinPrice(ctx),
	)
}

// set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}
