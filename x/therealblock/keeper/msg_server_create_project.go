package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
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
	id, err := k.AppendProject(
		ctx,
		project,
	)
	if err != nil {
		return &types.MsgCreateProjectResponse{}, err
	}
	types.EmitEvent(ctx, types.EventTypeProjectCreated, id, msg.Sponsor)
	return &types.MsgCreateProjectResponse{
		Id:      id,
		Address: msg.Sponsor,
	}, nil
}

func (k Keeper) AppendProject(ctx sdk.Context, project types.Project) (uint64, error) {
	count := k.getProjectCount(ctx)
	project.Id = count
	if !k.bankKeeper.HasSupply(ctx, project.Target.Denom) {
		return 0, types.ErrCoinNotSupply
	}
	if err := k.checkStages(project.Stages, project.Target); err != nil {
		return 0, err
	}
	project.Current = sdk.NewCoin(project.Target.Denom, sdk.ZeroInt())
	project.State = types.ProjectStateDraft
	project.Investors = make([]*types.Investor, 0)
	k.saveProject(ctx, &project)
	k.setProjectCount(ctx, count+1)
	return count, nil
}

func (k Keeper) getProjectCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectCountKey))
	byteKey := types.KeyPrefix(types.ProjectCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) setProjectCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectCountKey))
	store.Set(types.KeyPrefix(types.ProjectCountKey), sdk.Uint64ToBigEndian(count))
}
