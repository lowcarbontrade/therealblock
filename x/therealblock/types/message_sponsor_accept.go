package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSponsorAccept = "sponsor_accept"

var _ sdk.Msg = &MsgSponsorAccept{}

func NewMsgSponsorAccept(creator string, projectId uint64) *MsgSponsorAccept {
	return &MsgSponsorAccept{
		Creator:   creator,
		ProjectId: projectId,
	}
}

func (msg *MsgSponsorAccept) Route() string {
	return RouterKey
}

func (msg *MsgSponsorAccept) Type() string {
	return TypeMsgSponsorAccept
}

func (msg *MsgSponsorAccept) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSponsorAccept) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSponsorAccept) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
