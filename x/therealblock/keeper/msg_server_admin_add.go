package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) AdminAdd(goCtx context.Context, msg *types.MsgAdminAdd) (*types.MsgAdminAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	if k.IsAdminAccount(ctx, msg.NewAddress) {
		return nil, types.ErrAdminAccountExists
	}
	k.deleteDevAccount(ctx)
	k.SetAdminAccount(ctx, types.Account{Address: msg.NewAddress})
	if !k.IsAdminAccount(ctx, msg.NewAddress) {
		return nil, types.ErrAdminAccountNotSet
	}
	return &types.MsgAdminAddResponse{
		Address: msg.NewAddress,
	}, nil
}

func (k msgServer) deleteDevAccount(ctx sdk.Context) {
	k.DeleteAdminAccount(ctx, types.DevAddr)
	if k.IsAdminAccount(ctx, types.DevAddr) {
		//TODO research more about panic handling and why might be needed instead of errors
		panic("Failed to delete dev account")
	}
}
