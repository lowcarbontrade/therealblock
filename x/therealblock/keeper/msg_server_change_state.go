package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) ChangeState(goCtx context.Context, msg *types.MsgChangeState) (*types.MsgChangeStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	projectId, err := k.ChangeState(ctx, msg.NewState, msg.ProjectId)
	if err != nil {
		return &types.MsgChangeStateResponse{}, err
	}
	return &types.MsgChangeStateResponse{
		ProjectId: projectId,
	}, nil
}
