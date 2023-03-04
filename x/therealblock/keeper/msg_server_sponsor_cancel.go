package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) SponsorCancel(goCtx context.Context, msg *types.MsgSponsorCancel) (*types.MsgSponsorCancelResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	projectId, err := k.sponsorCancelProject(ctx, msg.ProjectId, msg.Creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgSponsorCancelResponse{
		ProjectId: projectId,
	}, nil
}

func (k Keeper) sponsorCancelProject(ctx sdk.Context, projectId uint64, sponsor string) (uint64, error) {
	project, found := k.getProjectId(ctx, projectId)
	if !found {
		return 0, types.ErrProjectNotFound
	}
	if strings.Compare(project.Sponsor, sponsor) != 0 {
		return 0, types.ErrNotProjectSponsor
	}
	if project.State != types.ProjectStateActive && project.State != types.ProjectStatePending {

		return 0, types.ErrProjectNotCancelable
	}
	if err := k.returnFundsCancel(ctx, &project); err != nil {
		return 0, err
	}
	project.State = types.ProjectStateCancelled
	k.saveProject(ctx, &project)
	return project.Id, nil
}

func (k Keeper) returnFundsCancel(ctx sdk.Context, project *types.Project) error {
	for _, investor := range project.Investors {
		addr, err := sdk.AccAddressFromBech32(investor.Address)
		if err != nil {
			return err
		}
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(investor.Equity)); err != nil {
			return err
		}
		project.Current.Amount = project.Current.Amount.Sub(investor.Equity.Amount)
		investor.Equity.Amount = sdk.ZeroInt()
	}
	return nil
}
