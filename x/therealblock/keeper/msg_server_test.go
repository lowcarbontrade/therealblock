package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/realblocknetwork/therealblock/testutil/keeper"
	"github.com/realblocknetwork/therealblock/x/therealblock/keeper"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TherealblockKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
