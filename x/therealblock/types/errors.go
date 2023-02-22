package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/therealblock module sentinel errors
var (
	ErrSample                 = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrProjectNotFound        = sdkerrors.Register(ModuleName, 1101, "project not found")
	ErrProjectNotActive       = sdkerrors.Register(ModuleName, 1102, "project not in Active state")
	ErrInvalidState           = sdkerrors.Register(ModuleName, 1103, "not a valid project state, does not exist")
	ErrProjectStateNotChanged = sdkerrors.Register(ModuleName, 1104, "new state is the same as the current state")
)
