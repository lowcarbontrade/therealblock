package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k Keeper) AppendProject(ctx sdk.Context, project types.Project) uint64 {
	count := k.GetProjectCount(ctx)
	project.Id = count
	//TODO assert that coin denom exits (bank.keeper.hasSupply())
	project.Current = sdk.NewCoin(project.Target.Denom, sdk.ZeroInt())
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	appendedValue := k.cdc.MustMarshal(&project)
	store.Set(GetProjectIDBytes(project.Id), appendedValue)
	k.SetProjectCount(ctx, count+1)
	return count
}

func (k Keeper) GetProjectCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectCountKey))
	byteKey := types.KeyPrefix(types.ProjectCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}

func (k Keeper) SetProjectCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectCountKey))
	store.Set(types.KeyPrefix(types.ProjectCountKey), sdk.Uint64ToBigEndian(count))
}

func GetProjectIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
