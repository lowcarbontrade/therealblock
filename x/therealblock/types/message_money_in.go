package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgMoneyIn = "money_in"

var _ sdk.Msg = &MsgMoneyIn{}

func NewMsgMoneyIn(creator string, addrTo string, amount sdk.Coin) *MsgMoneyIn {
	return &MsgMoneyIn{
		Creator: creator,
		AddrTo:  addrTo,
		Amount:  amount,
	}
}

func (msg *MsgMoneyIn) Route() string {
	return RouterKey
}

func (msg *MsgMoneyIn) Type() string {
	return TypeMsgMoneyIn
}

func (msg *MsgMoneyIn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgMoneyIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMoneyIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
