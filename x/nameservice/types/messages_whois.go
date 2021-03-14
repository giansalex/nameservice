package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetName{}

func NewMsgSetName(name string, value string, owner string) *MsgSetName {
	return &MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

func (msg *MsgSetName) Route() string {
	return RouterKey
}

func (msg *MsgSetName) Type() string {
	return "SetName"
}

func (msg *MsgSetName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSetName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name and/or Value cannot be empty")
	}
	return nil
}

var _ sdk.Msg = &MsgBuyName{}

func NewMsgBuyName(name string, bid string, buyer string) *MsgBuyName {
	return &MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

func (msg *MsgBuyName) Route() string {
	return RouterKey
}

func (msg *MsgBuyName) Type() string {
	return "BuyName"
}

func (msg *MsgBuyName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgBuyName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid buyer address (%s)", err)
	}
	if len(msg.Name) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Name cannot be empty")
	}

	bidCoins, err := sdk.ParseCoinsNormalized(msg.Bid)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid bid coins (%s)", err)
	}

	if !bidCoins.IsAllPositive() {
		return sdkerrors.ErrInsufficientFunds
	}

	return nil
}

var _ sdk.Msg = &MsgDeleteName{}

func NewMsgDeleteName(id, creator string) *MsgDeleteName {
	return &MsgDeleteName{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteName) Route() string {
	return RouterKey
}

func (msg *MsgDeleteName) Type() string {
	return "DeleteName"
}

func (msg *MsgDeleteName) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteName) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteName) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateWhois{}

func NewMsgCreateWhois(creator string, value string, price string) *MsgCreateWhois {
	return &MsgCreateWhois{
		Creator: creator,
		Value:   value,
		Price:   price,
	}
}

func (msg *MsgCreateWhois) Route() string {
	return RouterKey
}

func (msg *MsgCreateWhois) Type() string {
	return "CreateWhois"
}

func (msg *MsgCreateWhois) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWhois) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWhois) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWhois{}

func NewMsgUpdateWhois(creator string, id string, value string, price string) *MsgUpdateWhois {
	return &MsgUpdateWhois{
		Id:      id,
		Creator: creator,
		Value:   value,
		Price:   price,
	}
}

func (msg *MsgUpdateWhois) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWhois) Type() string {
	return "UpdateWhois"
}

func (msg *MsgUpdateWhois) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWhois) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWhois) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateWhois{}

func NewMsgDeleteWhois(creator string, id string) *MsgDeleteWhois {
	return &MsgDeleteWhois{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteWhois) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWhois) Type() string {
	return "DeleteWhois"
}

func (msg *MsgDeleteWhois) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWhois) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWhois) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
