package keeper

import (
	sdkmath "cosmossdk.io/math"
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"strings"
)

func (k Keeper) AppendProject(ctx sdk.Context, project types.Project) (uint64, error) {
	count := k.GetProjectCount(ctx)
	project.Id = count
	if !k.bankKeeper.HasSupply(ctx, project.Target.Denom) {
		return 0, types.ErrCoinNotSupply
	}
	if err := k.checkStages(ctx, project.Stages, project.Target); err != nil {
		return 0, err
	}
	project.Current = sdk.NewCoin(project.Target.Denom, sdk.ZeroInt())
	project.State = types.ProjectStateDraft
	project.Investors = make([]*types.Investor, 0)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	appendedValue := k.cdc.MustMarshal(&project)
	store.Set(GetProjectIDBytes(project.Id), appendedValue)
	k.SetProjectCount(ctx, count+1)
	return count, nil
}

func (k Keeper) checkStages(ctx sdk.Context, stages []*types.Stage, target sdk.Coin) error {
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

func (k Keeper) GetProjectId(ctx sdk.Context, id uint64) (val types.Project, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	b := store.Get(GetProjectIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) AppendInvestorBuyIn(ctx sdk.Context, id uint64, investor types.Investor) (string, error) {
	if investor.Equity.Amount.Equal(sdk.ZeroInt()) {
		return "", types.ErrCoinZeroAmount
	}
	project, found := k.GetProjectId(ctx, id)
	if !found {
		return "", types.ErrProjectNotFound
	}
	if project.State != types.ProjectStateActive {
		return "", types.ErrProjectNotActive
	}
	if strings.Compare(project.Target.Denom, investor.Equity.Denom) != 0 {
		return "", types.ErrCoinDiffDenom
	}
	if project.Target.Sub(project.Current).IsLT(investor.Equity) {
		return "", types.ErrOverFunded
	}
	addr, err := sdk.AccAddressFromBech32(investor.Address)
	if err != nil {
		return "", err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(investor.Equity)); err != nil {
		return "", err
	}
	project.Investors = appendInvestor(project.Investors, &investor)
	project.Current = project.Current.Add(investor.Equity)
	if project.Target.IsEqual(project.Current) {
		project.State = types.ProjectStateFunded
		types.EmitEvent(ctx, types.ProjectFundedEventType, project.Id, investor.Address)
	} else {
		types.EmitEvent(ctx, types.ProjectInvestedEventType, project.Id, investor.Address)
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	store.Set(GetProjectIDBytes(project.Id), k.cdc.MustMarshal(&project))
	return investor.Address, nil
}

func appendInvestor(investors []*types.Investor, newInvestor *types.Investor) []*types.Investor {
	var found = false
	for _, investor := range investors {
		if strings.Compare(investor.Address, newInvestor.Address) == 0 {
			found = true
			investor.Equity = investor.Equity.Add(newInvestor.Equity)
		}
	}
	if !found {
		return append(investors, newInvestor)
	}
	return investors
}

func (k Keeper) ChangeProjectState(ctx sdk.Context, newState string, projectId uint64) (uint64, error) {
	if err := types.IsValidState(newState); err != nil {
		return 0, err
	}
	project, found := k.GetProjectId(ctx, projectId)
	if !found {
		return 0, types.ErrProjectNotFound
	}
	if project.State == newState {
		return 0, types.ErrProjectStateNotChanged
	}
	project.State = newState
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	store.Set(GetProjectIDBytes(project.Id), k.cdc.MustMarshal(&project))
	return project.Id, nil
}
