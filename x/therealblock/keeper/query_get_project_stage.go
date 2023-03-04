package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetProjectStage(goCtx context.Context, req *types.QueryGetProjectStageRequest) (*types.QueryGetProjectStageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	stage, err := k.getStageInfo(ctx, req.ProjectId)
	if err != nil {
		return nil, err
	}
	return &types.QueryGetProjectStageResponse{
		Stage: stage,
	}, nil
}

func (k Keeper) getStageInfo(ctx sdk.Context, projectId uint64) (types.Stage, error) {
	project, found := k.getProjectId(ctx, projectId)
	if !found {
		return types.Stage{}, types.ErrProjectNotFound
	}
	if strings.Compare(project.State, types.ProjectStateFunded) == 0 ||
		strings.Compare(project.State, types.ProjectStateCompleted) == 0 {
		var aux = project.Target
		for _, stage := range project.Stages {
			aux = aux.Sub(stage.Allocation)
			if aux.Equal(project.Current) {
				return types.Stage{
					Name:       stage.Name,
					Allocation: stage.Allocation,
				}, nil
			}
		}
	}
	return types.Stage{}, types.ErrProjectNotFunded
}
