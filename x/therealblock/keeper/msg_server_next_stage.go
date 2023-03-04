package keeper

import (
	"context"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) NextStage(goCtx context.Context, msg *types.MsgNextStage) (*types.MsgNextStageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	projectId, err := k.nextProjectStage(ctx, msg.ProjectId)
	if err != nil {
		return nil, err
	}
	return &types.MsgNextStageResponse{
		ProjectId: projectId,
	}, nil
}

func (k Keeper) nextProjectStage(ctx sdk.Context, projectId uint64) (uint64, error) {
	project, found := k.getProjectId(ctx, projectId)
	if !found {
		return 0, types.ErrProjectNotFound
	}
	if project.State != types.ProjectStateFunded {
		return 0, types.ErrProjectNotFunded
	}
	allocation, hasNext := calculateNextStage(project)
	if hasNext {
		sponsor, err := sdk.AccAddressFromBech32(project.Sponsor)
		if err != nil {
			return 0, err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sponsor, sdk.NewCoins(allocation)); err != nil {
			return 0, err
		}
		project.Current = project.Current.Sub(allocation)
	} else {
		project.State = types.ProjectStateCompleted
	}
	k.saveProject(ctx, &project)
	return project.Id, nil
}

func calculateNextStage(project types.Project) (sdk.Coin, bool) {
	var aux = project.Target
	var found = false
	for _, stage := range project.Stages {
		if found {
			return stage.Allocation, true
		}
		aux = aux.Sub(stage.Allocation)
		if aux.Equal(project.Current) {
			found = true
		}
	}
	return sdk.NewCoin("end", sdkmath.NewInt(0)), false
}
