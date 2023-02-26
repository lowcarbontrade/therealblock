package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func (k Keeper) MintRBS(ctx sdk.Context, amount sdk.Coin, addrTo string) (string, error) {
	if !k.bankKeeper.HasSupply(ctx, amount.Denom) {
		return "", types.ErrCoinNotSupply
	}
	addr, err := sdk.AccAddressFromBech32(addrTo)
	if err != nil {
		return "", err
	}
	//TODO check what happens between mint and send (if an error occurs is the tx reverted?)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return "", err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, addr, sdk.NewCoins(amount)); err != nil {
		return "", err
	}
	return addr.String(), nil
}

func (k Keeper) BurnRBS(ctx sdk.Context, amount sdk.Coin, addrFrom string) (string, error) {
	addr, err := sdk.AccAddressFromBech32(addrFrom)
	if err != nil {
		return "", err
	}
	if !k.bankKeeper.HasBalance(ctx, addr, amount) {
		return "", types.ErrNotEnoughBalance
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return "", err
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(amount)); err != nil {
		return "", err
	}
	return addr.String(), nil
}
