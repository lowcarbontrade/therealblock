package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) MoneyOut(goCtx context.Context, msg *types.MsgMoneyOut) (*types.MsgMoneyOutResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if strings.Compare(msg.Amount.Denom, "rbs") != 0 {
		return nil, types.ErrInvalidDenom
	}
	burnerAddr, err := k.burn(ctx, msg.Amount, msg.Creator)
	if err != nil {
		return nil, err
	}
	//TODO customize emit event
	types.EmitEvent(ctx, types.EventTypeMoneyOut, 0, msg.Creator)
	return &types.MsgMoneyOutResponse{
		BurnAddr: burnerAddr,
	}, nil
}
