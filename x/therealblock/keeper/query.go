package keeper

import (
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

var _ types.QueryServer = Keeper{}
