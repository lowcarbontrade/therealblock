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
	ErrCoinNotSupply          = sdkerrors.Register(ModuleName, 1105, "coin does not have supply")
	ErrNoStages               = sdkerrors.Register(ModuleName, 1106, "project has no stages defined")
	ErrCoinDiffDenom          = sdkerrors.Register(ModuleName, 1107, "coin denom does not match project target denom")
	ErrCoinDiffAmount         = sdkerrors.Register(ModuleName, 1108, "stages coin amount does not match project target amount")
	ErrInvalidStageFormat     = sdkerrors.Register(ModuleName, 1109, "invalid stage format")
	ErrOverFunded             = sdkerrors.Register(ModuleName, 1110, "project is overfunded")
	ErrCoinZeroAmount         = sdkerrors.Register(ModuleName, 1111, "coin amount is zero")
	ErrNotEnoughBalance       = sdkerrors.Register(ModuleName, 1112, "not enough balance to perform the operation")
	ErrNotProjectSponsor      = sdkerrors.Register(ModuleName, 1113, "signing address not the project sponsor")
)
