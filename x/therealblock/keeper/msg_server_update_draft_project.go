package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) UpdateDraftProject(goCtx context.Context, msg *types.MsgUpdateDraftProject) (*types.MsgUpdateDraftProjectResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	projectId, err := k.updateDraftProjectInfo(ctx, msg.ProjectId, msg.Target, msg.Stages, msg.Creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgUpdateDraftProjectResponse{
		ProjectId: projectId,
	}, nil
}

func (k Keeper) updateDraftProjectInfo(ctx sdk.Context, projectId uint64, target sdk.Coin, stages []*types.Stage, signer string) (uint64, error) {
	project, found := k.getProjectId(ctx, projectId)
	if !found {
		return 0, types.ErrProjectNotFound
	}
	if strings.Compare(signer, project.Sponsor) != 0 {
		return 0, types.ErrNotProjectSponsor
	}
	if strings.Compare(project.State, types.ProjectStateDraft) != 0 {
		return 0, types.ErrProjectNotDraft
	}
	if err := k.checkStages(stages, target); err != nil {
		return 0, err
	}
	project.Target = target
	project.Stages = stages
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	store.Set(getProjectIDBytes(project.Id), k.cdc.MustMarshal(&project))
	return project.Id, nil
}
