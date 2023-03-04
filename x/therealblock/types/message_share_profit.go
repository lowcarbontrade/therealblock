package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgShareProfit = "share_profit"

var _ sdk.Msg = &MsgShareProfit{}

func NewMsgShareProfit(creator string, projectId uint64, amount sdk.Coin) *MsgShareProfit {
	return &MsgShareProfit{
		Creator:   creator,
		ProjectId: projectId,
		Amount:    amount,
	}
}

func (msg *MsgShareProfit) Route() string {
	return RouterKey
}

func (msg *MsgShareProfit) Type() string {
	return TypeMsgShareProfit
}

func (msg *MsgShareProfit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgShareProfit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgShareProfit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
