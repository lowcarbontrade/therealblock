package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAdminDelete = "admin_delete"

var _ sdk.Msg = &MsgAdminDelete{}

func NewMsgAdminDelete(creator string, address string) *MsgAdminDelete {
	return &MsgAdminDelete{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgAdminDelete) Route() string {
	return RouterKey
}

func (msg *MsgAdminDelete) Type() string {
	return TypeMsgAdminDelete
}

func (msg *MsgAdminDelete) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAdminDelete) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAdminDelete) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
