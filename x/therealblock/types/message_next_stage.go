package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNextStage = "next_stage"

var _ sdk.Msg = &MsgNextStage{}

func NewMsgNextStage(creator string, projectId uint64) *MsgNextStage {
	return &MsgNextStage{
		Creator:   creator,
		ProjectId: projectId,
	}
}

func (msg *MsgNextStage) Route() string {
	return RouterKey
}

func (msg *MsgNextStage) Type() string {
	return TypeMsgNextStage
}

func (msg *MsgNextStage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNextStage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNextStage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
