package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"strings"
)

func (k msgServer) MoneyIn(goCtx context.Context, msg *types.MsgMoneyIn) (*types.MsgMoneyInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	if strings.Compare(msg.Amount.Denom, "rbs") != 0 {
		return nil, types.ErrInvalidDenom
	}
	addr, err := k.Mint(ctx, msg.Amount, msg.AddrTo)
	if err != nil {
		return &types.MsgMoneyInResponse{}, err
	}
	//TODO need more customizable events
	types.EmitEvent(ctx, types.EventTypeMoneyIn, 0, msg.AddrTo)
	return &types.MsgMoneyInResponse{
		AddrTo: addr,
	}, nil
}
