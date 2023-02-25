package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgInvestorBuyIn = "investor_buy_in"

var _ sdk.Msg = &MsgInvestorBuyIn{}

func NewMsgInvestorBuyIn(investor string, projectId uint64, amount sdk.Coin) *MsgInvestorBuyIn {
	return &MsgInvestorBuyIn{
		Investor:  investor,
		ProjectId: projectId,
		Amount:    amount,
	}
}

func (msg *MsgInvestorBuyIn) Route() string {
	return RouterKey
}

func (msg *MsgInvestorBuyIn) Type() string {
	return TypeMsgInvestorBuyIn
}

func (msg *MsgInvestorBuyIn) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Investor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInvestorBuyIn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInvestorBuyIn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Investor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
