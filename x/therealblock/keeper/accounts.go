package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k Keeper) GetAccounts(ctx sdk.Context) []types.Account {
	var accounts []types.Account
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
	for ; iterator.Valid(); iterator.Next() {
		var account types.Account
		k.cdc.MustUnmarshal(iterator.Value(), &account)
		accounts = append(accounts, account)
	}
	err := iterator.Close()
	if err != nil {
		return nil
	}
	return accounts
}

func (k Keeper) SetAccounts(ctx sdk.Context, accounts []types.Account) {
	store := ctx.KVStore(k.storeKey)
	for _, account := range accounts {
		store.Set(types.KeyPrefix(types.GenAccountKey), k.cdc.MustMarshal(&account))
	}
}

func (k Keeper) SetAccount(ctx sdk.Context, account types.Account) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefix(types.GenAccountKey), k.cdc.MustMarshal(&account))
}
