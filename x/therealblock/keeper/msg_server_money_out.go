package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) MoneyOut(goCtx context.Context, msg *types.MsgMoneyOut) (*types.MsgMoneyOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	burnerAddr, err := k.BurnRBS(ctx, msg.Amount, msg.Creator)
	if err != nil {
		return nil, err
	}

	types.EmitEvent(ctx, types.EventTypeMoneyOut, types.ProjectEventProjectCreator, msg.Creator)
	return &types.MsgMoneyOutResponse{
		BurnAddr: burnerAddr,
	}, nil
}
