package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	keeper.SetParams(ctx, data.Params)
	for _, record := range data.WhoisRecords {
		keeper.SetWhois(ctx, record.Value, record)
	}
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Whois
	iterator := k.GetNamesIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {

		name := string(iterator.Key())
		whois := k.GetWhois(ctx, name)
		records = append(records, whois)

	}
	return GenesisState{
		Params:       k.GetParams(ctx),
		WhoisRecords: records,
	}
}
