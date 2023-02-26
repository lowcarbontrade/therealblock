package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	//TODO implement error handling
	ctx := sdk.UnwrapSDKContext(goCtx)
	var project = types.Project{
		Stages:  msg.Stages,
		Sponsor: msg.Sponsor,
		Target:  msg.Target,
	}
	id, err := k.AppendProject(
		ctx,
		project,
	)
	if err != nil {
		return &types.MsgCreateProjectResponse{}, err
	}
	types.EmitEvent(ctx, types.EventTypeProjectCreated, id, msg.Sponsor)

	//TODO find out how to get the signer address of the transaction
	return &types.MsgCreateProjectResponse{
		Id:      id,
		Address: msg.Sponsor,
	}, nil
}
