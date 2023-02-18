package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateProject = "create_project"

var _ sdk.Msg = &MsgCreateProject{}

func NewMsgCreateProject(sponsor string, target sdk.Coin, stages []*Stage) *MsgCreateProject {
	return &MsgCreateProject{
		Sponsor: sponsor,
		Target:  target,
		Stages:  stages,
	}
}

func (msg *MsgCreateProject) Route() string {
	return RouterKey
}

func (msg *MsgCreateProject) Type() string {
	return TypeMsgCreateProject
}

func (msg *MsgCreateProject) GetSigners() []sdk.AccAddress {
	sponsor, err := sdk.AccAddressFromBech32(msg.Sponsor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sponsor}
}

func (msg *MsgCreateProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sponsor)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sponsor address (%s)", err)
	}
	return nil
}
