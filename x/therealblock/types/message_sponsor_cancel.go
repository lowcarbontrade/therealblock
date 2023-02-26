package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSponsorCancel = "sponsor_cancel"

var _ sdk.Msg = &MsgSponsorCancel{}

func NewMsgSponsorCancel(creator string, projectId uint64) *MsgSponsorCancel {
	return &MsgSponsorCancel{
		Creator:   creator,
		ProjectId: projectId,
	}
}

func (msg *MsgSponsorCancel) Route() string {
	return RouterKey
}

func (msg *MsgSponsorCancel) Type() string {
	return TypeMsgSponsorCancel
}

func (msg *MsgSponsorCancel) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSponsorCancel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSponsorCancel) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
