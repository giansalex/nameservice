package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/giansalex/nameservice/x/nameservice/types"
)

func (k msgServer) SetName(goCtx context.Context, msg *types.MsgSetName) (*types.MsgSetNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	if !k.HasWhois(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Name %s doesn't exist", msg.Name))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Owner != k.GetWhoisOwner(ctx, msg.Name) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetNameValue(ctx, msg.Name, msg.Value)

	return &types.MsgSetNameResponse{}, nil
}

func (k msgServer) BuyName(goCtx context.Context, msg *types.MsgBuyName) (*types.MsgBuyNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var whois types.Whois
	if k.HasWhois(ctx, msg.Name) {
		whois = k.GetWhois(ctx, msg.Name)
		price, _ := sdk.ParseCoinsNormalized(whois.Price)
		bidPrice, _ := sdk.ParseCoinsNormalized(msg.Bid)

		if price.IsAllGT(bidPrice) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFunds, "Bid not high enough") // If not, throw an error
		}

		creator, err := sdk.AccAddressFromBech32(whois.Creator)
		if err != nil {
			panic(err)
		}
		buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
		if err != nil {
			return nil, err
		}

		err = k.bankKeeper.SendCoins(ctx, buyer, creator, bidPrice)
		if err != nil {
			return nil, err
		}

		whois.Creator = msg.Buyer
		whois.Price = msg.Bid

		k.SetWhois(ctx, whois)
	} else {
		bidPrice, _ := sdk.ParseCoinsNormalized(msg.Bid)
		buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
		if err != nil {
			return nil, err
		}

		err = k.bankKeeper.SubtractCoins(ctx, buyer, bidPrice)
		if err != nil {
			return nil, err
		}

		whois = types.Whois{
			Id:      msg.Name,
			Creator: msg.Buyer,
			Price:   msg.Bid,
		}
	}

	k.SetWhois(ctx, whois)

	return &types.MsgBuyNameResponse{}, nil
}

func (k msgServer) DeleteName(goCtx context.Context, msg *types.MsgDeleteName) (*types.MsgDeleteNameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasWhois(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetWhoisOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhois(ctx, msg.Id)

	return &types.MsgDeleteNameResponse{}, nil
}

func (k msgServer) CreateWhois(goCtx context.Context, msg *types.MsgCreateWhois) (*types.MsgCreateWhoisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	id := k.AppendWhois(
		ctx,
		msg.Creator,
		msg.Value,
		msg.Price,
	)

	return &types.MsgCreateWhoisResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateWhois(goCtx context.Context, msg *types.MsgUpdateWhois) (*types.MsgUpdateWhoisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var whois = types.Whois{
		Creator: msg.Creator,
		Id:      msg.Id,
		Value:   msg.Value,
		Price:   msg.Price,
	}

	// Checks that the element exists
	if !k.HasWhois(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetWhoisOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetWhois(ctx, whois)

	return &types.MsgUpdateWhoisResponse{}, nil
}

func (k msgServer) DeleteWhois(goCtx context.Context, msg *types.MsgDeleteWhois) (*types.MsgDeleteWhoisResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasWhois(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %s doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetWhoisOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveWhois(ctx, msg.Id)

	return &types.MsgDeleteWhoisResponse{}, nil
}
