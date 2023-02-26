package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) SponsorCancel(goCtx context.Context, msg *types.MsgSponsorCancel) (*types.MsgSponsorCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	projectId, err := k.SponsorCancelProject(ctx, msg.ProjectId, msg.Creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgSponsorCancelResponse{
		ProjectId: projectId,
	}, nil
}
