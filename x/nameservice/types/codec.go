package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSetName{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(&MsgBuyName{}, "nameservice/BuyName", nil)
	cdc.RegisterConcrete(&MsgDeleteName{}, "nameservice/DeleteName", nil)
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateWhois{}, "nameservice/CreateWhois", nil)
	cdc.RegisterConcrete(&MsgUpdateWhois{}, "nameservice/UpdateWhois", nil)
	cdc.RegisterConcrete(&MsgDeleteWhois{}, "nameservice/DeleteWhois", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSetName{},
		&MsgBuyName{},
		&MsgDeleteName{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateWhois{},
		&MsgUpdateWhois{},
		&MsgDeleteWhois{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
