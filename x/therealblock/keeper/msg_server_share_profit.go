package keeper

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k msgServer) ShareProfit(goCtx context.Context, msg *types.MsgShareProfit) (*types.MsgShareProfitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if strings.Compare(msg.Amount.Denom, "rbs") != 0 {
		return nil, types.ErrInvalidDenom
	}
	projectId, err := k.shareProfit(ctx, msg.ProjectId, msg.Amount, msg.Creator)
	if err != nil {
		return nil, err
	}
	return &types.MsgShareProfitResponse{
		ProjectId: projectId,
	}, nil
}

func (k Keeper) shareProfit(ctx sdk.Context, projectId uint64, profit sdk.Coin, signer string) (uint64, error) {
	project, found := k.getProjectId(ctx, projectId)
	if !found {
		return 0, types.ErrProjectNotFound
	}
	if strings.Compare(project.Sponsor, signer) != 0 {
		return 0, types.ErrNotProjectSponsor
	}
	if len(project.Investors) == 0 {
		return 0, types.ErrNoInvestors
	}
	sponsor, err := sdk.AccAddressFromBech32(project.Sponsor)
	if err != nil {
		return 0, err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sponsor, types.ModuleName, sdk.NewCoins(profit)); err != nil {
		return 0, err
	}
	for _, investor := range project.Investors {
		addr, err := sdk.AccAddressFromBech32(investor.Address)
		if err != nil {
			return 0, err
		}
		var equity = profit.Amount.Mul(
			investor.Equity.Amount.Mul(sdkmath.NewInt(100)).Quo(
				project.Target.Amount)).Quo(
			sdkmath.NewInt(100))
		if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(sdk.NewCoin("rbs", equity))); err != nil {
			return 0, err
		}
		investor.Profit += equity.Int64()
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ProjectKey))
	store.Set(getProjectIDBytes(project.Id), k.cdc.MustMarshal(&project))
	return project.Id, nil
}
