package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) AdminAdd(goCtx context.Context, msg *types.MsgAdminAdd) (*types.MsgAdminAddResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	if k.IsAdminAccount(ctx, msg.NewAddress) {
		return nil, types.ErrAdminAccountExists
	}
	k.setAdminAccount(ctx, types.Account{Address: msg.NewAddress})
	if !k.IsAdminAccount(ctx, msg.NewAddress) {
		return nil, types.ErrAdminAccountNotSet
	}
	return &types.MsgAdminAddResponse{
		Address: msg.NewAddress,
	}, nil
}

func (k Keeper) setAdminAccount(ctx sdk.Context, account types.Account) {
	if !k.IsAdminAccount(ctx, account.Address) {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
		store.Set(types.KeyPrefix(account.Address), k.cdc.MustMarshal(&account))
	}
}
