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
