package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/giansalex/nameservice/x/nameservice/types"
)

// SetNameValue set name for specific whois
func (k Keeper) SetNameValue(ctx sdk.Context, name, value string) {
	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, whois)
}

// SetWhois set a specific whois in the store
func (k Keeper) SetWhois(ctx sdk.Context, whois types.Whois) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	b := k.cdc.MustMarshalBinaryBare(&whois)
	store.Set(types.KeyPrefix(types.WhoisKey+whois.Id), b)
}

// GetWhois returns a whois from its id
func (k Keeper) GetWhois(ctx sdk.Context, key string) types.Whois {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	var whois types.Whois
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.WhoisKey+key)), &whois)
	return whois
}

// HasWhois checks if the whois exists in the store
func (k Keeper) HasWhois(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	return store.Has(types.KeyPrefix(types.WhoisKey + id))
}

// GetWhoisOwner returns the creator of the whois
func (k Keeper) GetWhoisOwner(ctx sdk.Context, key string) string {
	return k.GetWhois(ctx, key).Creator
}

// DeleteWhois removes a whois from the store
func (k Keeper) RemoveWhois(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	store.Delete(types.KeyPrefix(types.WhoisKey + key))
}

// GetAllWhois returns all whois
func (k Keeper) GetAllWhois(ctx sdk.Context) (list []types.Whois) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.WhoisKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.WhoisKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Whois
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
