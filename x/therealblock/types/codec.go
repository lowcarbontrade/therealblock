package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateProject{}, "therealblock/CreateProject", nil)
	cdc.RegisterConcrete(&MsgInvestorBuyIn{}, "therealblock/InvestorBuyIn", nil)
	cdc.RegisterConcrete(&MsgChangeState{}, "therealblock/ChangeState", nil)
	cdc.RegisterConcrete(&MsgMoneyIn{}, "therealblock/MoneyIn", nil)
	cdc.RegisterConcrete(&MsgMoneyOut{}, "therealblock/MoneyOut", nil)
	cdc.RegisterConcrete(&MsgSponsorCancel{}, "therealblock/SponsorCancel", nil)
	cdc.RegisterConcrete(&MsgSponsorAccept{}, "therealblock/SponsorAccept", nil)
	cdc.RegisterConcrete(&MsgAdminAdd{}, "therealblock/AdminAdd", nil)
	cdc.RegisterConcrete(&MsgAdminDelete{}, "therealblock/AdminDelete", nil)
	cdc.RegisterConcrete(&MsgNextStage{}, "therealblock/NextStage", nil)
	cdc.RegisterConcrete(&MsgShareProfit{}, "therealblock/ShareProfit", nil)
	cdc.RegisterConcrete(&MsgUpdateDraftProject{}, "therealblock/UpdateDraftProject", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProject{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInvestorBuyIn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgChangeState{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMoneyIn{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMoneyOut{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSponsorCancel{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSponsorAccept{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAdminAdd{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAdminDelete{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNextStage{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgShareProfit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateDraftProject{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
