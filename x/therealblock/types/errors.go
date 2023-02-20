package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/therealblock module sentinel errors
var (
	ErrSample           = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrProjectNotFound  = sdkerrors.Register(ModuleName, 1101, "project not found")
	ErrProjectNotActive = sdkerrors.Register(ModuleName, 1102, "project not in Active state")
)
