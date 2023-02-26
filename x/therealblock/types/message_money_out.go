package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMoneyOut = "money_out"

var _ sdk.Msg = &MsgMoneyOut{}

func NewMsgMoneyOut(creator string, amount sdk.Coin) *MsgMoneyOut {
	return &MsgMoneyOut{
		Creator: creator,
		Amount:  amount,
	}
}

func (msg *MsgMoneyOut) Route() string {
	return RouterKey
}

func (msg *MsgMoneyOut) Type() string {
	return TypeMsgMoneyOut
}

func (msg *MsgMoneyOut) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMoneyOut) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMoneyOut) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
