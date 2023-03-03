package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k Keeper) GetAdminAccounts(ctx sdk.Context) ([]types.Account, error) {
	var accounts []types.Account
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
	for ; iterator.Valid(); iterator.Next() {
		var account types.Account
		k.cdc.MustUnmarshal(iterator.Value(), &account)
		accounts = append(accounts, account)
	}
	err := iterator.Close()
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (k Keeper) GetCountAdminAccounts(ctx sdk.Context) (int64, error) {
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

func (k Keeper) IsAdminAccount(ctx sdk.Context, address string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
	return store.Has(types.KeyPrefix(address))
}

func (k Keeper) SetAdminAccounts(ctx sdk.Context, accounts []types.Account) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
	for _, account := range accounts {
		if !store.Has(types.KeyPrefix(account.Address)) {
			store.Set(types.KeyPrefix(account.Address), k.cdc.MustMarshal(&account))
		}
	}
}

func (k Keeper) SetAdminAccount(ctx sdk.Context, account types.Account) {
	if !k.IsAdminAccount(ctx, account.Address) {
		store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
		store.Set(types.KeyPrefix(account.Address), k.cdc.MustMarshal(&account))
	}
}

func (k Keeper) DeleteAdminAccount(ctx sdk.Context, address string) error {
	adminCount, err := k.GetCountAdminAccounts(ctx)
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
