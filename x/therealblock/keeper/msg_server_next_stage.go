package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) NextStage(goCtx context.Context, msg *types.MsgNextStage) (*types.MsgNextStageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	projectId, err := k.NextProjectStage(ctx, msg.ProjectId)
	if err != nil {
		return nil, err
	}
	return &types.MsgNextStageResponse{
		ProjectId: projectId,
	}, nil
}
