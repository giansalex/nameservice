package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/giansalex/nameservice/x/nameservice/types"
)

// MinPrice - Min price for name
func (k Keeper) MinPrice(ctx sdk.Context) (res uint64) {
	k.paramspace.Get(ctx, types.KeyMinPrice, &res)
	return
}

// BondDenom - Bondable coin denomination
func (k Keeper) BondDenom(ctx sdk.Context) (res string) {
	k.paramspace.Get(ctx, types.KeyBondDenom, &res)
	return
}

// GetParams returns the total set of nameservice parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	return types.NewParams(
		k.MinPrice(ctx),
		k.BondDenom(ctx),
	)
}

// SetParams sets the nameservice parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
