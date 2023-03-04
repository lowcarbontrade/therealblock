package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) ChangeState(goCtx context.Context, msg *types.MsgChangeState) (*types.MsgChangeStateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	projectId, err := k.changeProjectState(ctx, msg.NewState, msg.ProjectId)
	if err != nil {
		return &types.MsgChangeStateResponse{}, err
	}
	return &types.MsgChangeStateResponse{
		ProjectId: projectId,
	}, nil
}

func (k Keeper) changeProjectState(ctx sdk.Context, newState string, projectId uint64) (uint64, error) {
	if err := types.IsValidState(newState); err != nil {
		return 0, err
	}
	project, found := k.getProjectId(ctx, projectId)
	if !found {
		return 0, types.ErrProjectNotFound
	}
	if project.State == newState {
		return 0, types.ErrProjectStateNotChanged
	}
	project.State = newState
	k.saveProject(ctx, &project)
	return project.Id, nil
}
