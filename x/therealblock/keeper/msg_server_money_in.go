package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) MoneyIn(goCtx context.Context, msg *types.MsgMoneyIn) (*types.MsgMoneyInResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := k.MintRBS(ctx, msg.Amount, msg.AddrTo)
	if err != nil {
		return &types.MsgMoneyInResponse{}, err
	}
	//TODO need more customizable events
	types.EmitEvent(ctx, types.EventTypeMoneyIn, 0, msg.AddrTo)
	return &types.MsgMoneyInResponse{
		AddrTo: addr,
	}, nil
}
