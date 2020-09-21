package nameservice

import (
	"github.com/giansalex/nameservice/x/nameservice/keeper"
	"github.com/giansalex/nameservice/x/nameservice/types"
)

const (
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
	QuerierRoute = types.QuerierRoute
)

var (
	DefaultParamspace   = types.DefaultParamspace
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewMsgBuyName       = types.NewMsgBuyName
	NewMsgSetName       = types.NewMsgSetName
	NewMsgDeleteName    = types.NewMsgDeleteName
	NewWhois            = types.NewWhois
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	MsgSetName      = types.MsgSetName
	MsgBuyName      = types.MsgBuyName
	MsgDeleteName   = types.MsgDeleteName
	QueryResResolve = types.QueryResResolve
	QueryResNames   = types.QueryResNames
	Whois           = types.Whois
)
