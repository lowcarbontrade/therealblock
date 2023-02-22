package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgChangeState = "change_state"

var _ sdk.Msg = &MsgChangeState{}

func NewMsgChangeState(creator string, projectId uint64, newState string) *MsgChangeState {
	return &MsgChangeState{
		Creator:   creator,
		ProjectId: projectId,
		NewState:  newState,
	}
}

func (msg *MsgChangeState) Route() string {
	return RouterKey
}

func (msg *MsgChangeState) Type() string {
	return TypeMsgChangeState
}

func (msg *MsgChangeState) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgChangeState) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgChangeState) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
