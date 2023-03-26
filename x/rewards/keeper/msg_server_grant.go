package keeper

import (
	"context"

    "github.com/gitopia/gitopia/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


func (k msgServer) Grant(goCtx context.Context,  msg *types.MsgGrant) (*types.MsgGrantResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fromAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "bad address " + msg.Creator)
	}

	toAddr, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "bad address " + msg.To)
	}

    err = k.bankKeeper.SendCoins(ctx, fromAddr, toAddr, msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "error transfering reward")
	}

	return &types.MsgGrantResponse{}, nil
}
