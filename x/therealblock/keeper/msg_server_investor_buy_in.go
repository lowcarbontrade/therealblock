package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
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
