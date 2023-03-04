package keeper

import (
	"context"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) SponsorAccept(goCtx context.Context, msg *types.MsgSponsorAccept) (*types.MsgSponsorAcceptResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	projectId, err := k.sponsorAcceptProject(ctx, msg.ProjectId, msg.Creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgSponsorAcceptResponse{
		ProjectId: projectId,
	}, nil
}

func (k Keeper) sponsorAcceptProject(ctx sdk.Context, projectId uint64, sponsor string) (uint64, error) {
	project, found := k.getProjectId(ctx, projectId)
	if !found {
		return 0, types.ErrProjectNotFound
	}
	if strings.Compare(project.Sponsor, sponsor) != 0 {
		return 0, types.ErrNotProjectSponsor
	}
	if project.State != types.ProjectStatePending {
		return 0, types.ErrProjectNotPending
	}
	project.State = types.ProjectStateFunded
	if err := k.issueProjectTokens(ctx, &project); err != nil {
		return 0, err
	}
	k.saveProject(ctx, &project)
	return project.Id, nil
}

func (k Keeper) issueProjectTokens(ctx sdk.Context, project *types.Project) error {
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(getDenomFromProject(project.Id), project.Target.Amount))); err != nil {
		return err
	}
	for _, investor := range project.Investors {
		addr, err := sdk.AccAddressFromBech32(investor.Address)
		if err != nil {
			return err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(sdk.NewCoin(getDenomFromProject(project.Id), investor.Equity.Amount))); err != nil {
			return err
		}
	}
	addr, err := sdk.AccAddressFromBech32(project.Sponsor)
	if err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(project.Stages[0].Allocation)); err != nil {
		return err
	}
	project.Current = project.Current.Sub(project.Stages[0].Allocation)
	return nil
}

func getDenomFromProject(projectId uint64) string {
	return "wRBS-" + strconv.FormatUint(projectId, 10)
}
