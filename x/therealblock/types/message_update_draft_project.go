package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateDraftProject = "update_draft_project"

var _ sdk.Msg = &MsgUpdateDraftProject{}

func NewMsgUpdateDraftProject(creator string, projectId uint64, target sdk.Coin, stages []*Stage) *MsgUpdateDraftProject {
	return &MsgUpdateDraftProject{
		Creator:   creator,
		ProjectId: projectId,
		Target:    target,
		Stages:    stages,
	}
}

func (msg *MsgUpdateDraftProject) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDraftProject) Type() string {
	return TypeMsgUpdateDraftProject
}

func (msg *MsgUpdateDraftProject) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDraftProject) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDraftProject) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
