package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/giansalex/nameservice/x/nameservice/keeper"
	"github.com/giansalex/nameservice/x/nameservice/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the whois
	for _, elem := range genState.WhoisList {
		k.SetWhois(ctx, *elem)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all whois
	whoisList := k.GetAllWhois(ctx)
	for _, elem := range whoisList {
		elem := elem
		genesis.WhoisList = append(genesis.WhoisList, &elem)
	}

	genesis.Params = k.GetParams(ctx)

	return genesis
}
