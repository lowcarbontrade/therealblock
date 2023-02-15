package therealblock_test

import (
	"testing"

	keepertest "github.com/realblocknetwork/therealblock/testutil/keeper"
	"github.com/realblocknetwork/therealblock/testutil/nullify"
	"github.com/realblocknetwork/therealblock/x/therealblock"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TherealblockKeeper(t)
	therealblock.InitGenesis(ctx, *k, genesisState)
	got := therealblock.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
