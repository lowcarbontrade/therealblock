package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) AdminDelete(goCtx context.Context, msg *types.MsgAdminDelete) (*types.MsgAdminDeleteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.IsAdminAccount(ctx, msg.Creator) {
		return nil, types.ErrNotAdminAccount
	}
	if err := k.deleteAdminAccount(ctx, msg.Address); err != nil {
		return nil, err
	}
	if k.IsAdminAccount(ctx, msg.Address) {
		return nil, types.ErrAdminAccountNotDeleted
	}
	return &types.MsgAdminDeleteResponse{
		Address: msg.Address,
	}, nil
}

func (k Keeper) deleteAdminAccount(ctx sdk.Context, address string) error {
	adminCount, err := k.getCountAdminAccounts(ctx)
	if err != nil {
		return err
	}
	if adminCount == 1 {
		return types.ErrLastAdminAccount
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
	store.Delete(types.KeyPrefix(address))
	return nil
}

func (k Keeper) getCountAdminAccounts(ctx sdk.Context) (int64, error) {
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
	var count int64
	for ; iterator.Valid(); iterator.Next() {
		count++
	}
	err := iterator.Close()
	if err != nil {
		return 0, err
	}
	return count, nil
}
