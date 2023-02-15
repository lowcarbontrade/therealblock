package keeper_test

import (
	"testing"

	testkeeper "github.com/realblocknetwork/therealblock/testutil/keeper"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TherealblockKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
