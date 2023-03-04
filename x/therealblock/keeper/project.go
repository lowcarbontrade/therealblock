package keeper

import (
	sdkmath "cosmossdk.io/math"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"strings"
)

func (k Keeper) checkStages(stages []*types.Stage, target sdk.Coin) error {
	var total = sdkmath.NewInt(0)
	for _, stage := range stages {
		if strings.Compare(target.Denom, stage.Allocation.Denom) != 0 {
			return types.ErrCoinDiffDenom
		}
		total = total.Add(stage.Allocation.Amount)
	}
	if !total.Equal(target.Amount) {
		return types.ErrCoinDiffAmount
	}
	return nil
}

func (k Keeper) saveProject(ctx sdk.Context, project *types.Project) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	store.Set(getProjectIDBytes(project.Id), k.cdc.MustMarshal(project))
}

func getProjectIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) getProjectId(ctx sdk.Context, id uint64) (val types.Project, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	b := store.Get(getProjectIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}
