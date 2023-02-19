package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) CreateProject(goCtx context.Context, msg *types.MsgCreateProject) (*types.MsgCreateProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var project = types.Project{
		Stages:  msg.Stages,
		Sponsor: msg.Sponsor,
		Target:  msg.Target,
	}
	id := k.AppendProject(
		ctx,
		project,
	)
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ProjectCreatedEventType,
			sdk.NewAttribute(types.ProjectCreatedEventAttributeKey, strconv.FormatUint(id, 10)),
			sdk.NewAttribute(types.ProjectCreatedEventAttributeCreator, msg.Sponsor),
		),
	)

	//TODO find out how to get the signer address of the transaction
	return &types.MsgCreateProjectResponse{
		Id:      id,
		Address: msg.Sponsor,
	}, nil
}
