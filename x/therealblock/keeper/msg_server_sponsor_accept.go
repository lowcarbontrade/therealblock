package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) SponsorAccept(goCtx context.Context, msg *types.MsgSponsorAccept) (*types.MsgSponsorAcceptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	projectId, err := k.SponsorAcceptProject(ctx, msg.ProjectId, msg.Creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgSponsorAcceptResponse{
		ProjectId: projectId,
	}, nil
}
