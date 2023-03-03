package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) AdminDelete(goCtx context.Context, msg *types.MsgAdminDelete) (*types.MsgAdminDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	if err := k.DeleteAdminAccount(ctx, msg.Address); err != nil {
		return nil, err
	}
	if k.IsAdminAccount(ctx, msg.Address) {
		return nil, types.ErrAdminAccountNotDeleted
	}
	return &types.MsgAdminDeleteResponse{
		Address: msg.Address,
	}, nil
}
