package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"strings"
)

func (k msgServer) InvestorBuyIn(goCtx context.Context, msg *types.MsgInvestorBuyIn) (*types.MsgInvestorBuyInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var investor = types.Investor{
		Address: msg.Investor,
		Equity:  msg.Amount,
	}
	appendedAddr, err := k.AppendInvestorBuyIn(ctx, msg.ProjectId, investor)
	if err != nil {
		return &types.MsgInvestorBuyInResponse{}, err
	}
	return &types.MsgInvestorBuyInResponse{
		InvestorAddr: appendedAddr,
	}, nil
}

func (k Keeper) AppendInvestorBuyIn(ctx sdk.Context, id uint64, investor types.Investor) (string, error) {
	if investor.Equity.Amount.Equal(sdk.ZeroInt()) {
		return "", types.ErrCoinZeroAmount
	}
	project, found := k.getProjectId(ctx, id)
	if !found {
		return "", types.ErrProjectNotFound
	}
	if project.State != types.ProjectStateActive {
		return "", types.ErrProjectNotActive
	}
	if strings.Compare(project.Target.Denom, investor.Equity.Denom) != 0 {
		return "", types.ErrCoinDiffDenom
	}
	if project.Target.Sub(project.Current).IsLT(investor.Equity) {
		return "", types.ErrOverFunded
	}
	addr, err := sdk.AccAddressFromBech32(investor.Address)
	if err != nil {
		return "", err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(investor.Equity)); err != nil {
		return "", err
	}
	project.Investors = appendInvestor(project.Investors, &investor)
	project.Current = project.Current.Add(investor.Equity)
	if project.Target.IsEqual(project.Current) {
		project.State = types.ProjectStatePending
		types.EmitEvent(ctx, types.EventTypeProjectPending, project.Id, investor.Address)
	} else {
		types.EmitEvent(ctx, types.EventTypeProjectInvested, project.Id, investor.Address)
	}
	k.saveProject(ctx, &project)
	return investor.Address, nil
}

func appendInvestor(investors []*types.Investor, newInvestor *types.Investor) []*types.Investor {
	var found = false
	for _, investor := range investors {
		if strings.Compare(investor.Address, newInvestor.Address) == 0 {
			found = true
			investor.Equity = investor.Equity.Add(newInvestor.Equity)
			investor.Profit -= newInvestor.Equity.Amount.Int64()
		}
	}
	if !found {
		newInvestor.Profit -= newInvestor.Equity.Amount.Int64()
		return append(investors, newInvestor)
	}
	return investors
}
