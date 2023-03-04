package therealblock

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/realblocknetwork/therealblock/testutil/sample"
	therealblocksimulation "github.com/realblocknetwork/therealblock/x/therealblock/simulation"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = therealblocksimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateProject = "op_weight_msg_create_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateProject int = 100

	opWeightMsgInvestorBuyIn = "op_weight_msg_investor_buy_in"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInvestorBuyIn int = 100

	opWeightMsgChangeState = "op_weight_msg_change_state"
	// TODO: Determine the simulation weight value
	defaultWeightMsgChangeState int = 100

	opWeightMsgMoneyIn = "op_weight_msg_money_in"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMoneyIn int = 100

	opWeightMsgMoneyOut = "op_weight_msg_money_out"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMoneyOut int = 100

	opWeightMsgSponsorCancel = "op_weight_msg_sponsor_cancel"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSponsorCancel int = 100

	opWeightMsgSponsorAccept = "op_weight_msg_sponsor_accept"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSponsorAccept int = 100

	opWeightMsgAdminAdd = "op_weight_msg_admin_add"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAdminAdd int = 100

	opWeightMsgAdminDelete = "op_weight_msg_admin_delete"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAdminDelete int = 100

	opWeightMsgNextStage = "op_weight_msg_next_stage"
	// TODO: Determine the simulation weight value
	defaultWeightMsgNextStage int = 100

	opWeightMsgShareProfit = "op_weight_msg_share_profit"
	// TODO: Determine the simulation weight value
	defaultWeightMsgShareProfit int = 100

	opWeightMsgUpdateDraftProject = "op_weight_msg_update_draft_project"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDraftProject int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	therealblockGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&therealblockGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateProject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = defaultWeightMsgCreateProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateProject,
		therealblocksimulation.SimulateMsgCreateProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgInvestorBuyIn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInvestorBuyIn, &weightMsgInvestorBuyIn, nil,
		func(_ *rand.Rand) {
			weightMsgInvestorBuyIn = defaultWeightMsgInvestorBuyIn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInvestorBuyIn,
		therealblocksimulation.SimulateMsgInvestorBuyIn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgChangeState int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgChangeState, &weightMsgChangeState, nil,
		func(_ *rand.Rand) {
			weightMsgChangeState = defaultWeightMsgChangeState
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgChangeState,
		therealblocksimulation.SimulateMsgChangeState(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMoneyIn int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMoneyIn, &weightMsgMoneyIn, nil,
		func(_ *rand.Rand) {
			weightMsgMoneyIn = defaultWeightMsgMoneyIn
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMoneyIn,
		therealblocksimulation.SimulateMsgMoneyIn(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgMoneyOut int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMoneyOut, &weightMsgMoneyOut, nil,
		func(_ *rand.Rand) {
			weightMsgMoneyOut = defaultWeightMsgMoneyOut
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMoneyOut,
		therealblocksimulation.SimulateMsgMoneyOut(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSponsorCancel int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSponsorCancel, &weightMsgSponsorCancel, nil,
		func(_ *rand.Rand) {
			weightMsgSponsorCancel = defaultWeightMsgSponsorCancel
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSponsorCancel,
		therealblocksimulation.SimulateMsgSponsorCancel(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSponsorAccept int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSponsorAccept, &weightMsgSponsorAccept, nil,
		func(_ *rand.Rand) {
			weightMsgSponsorAccept = defaultWeightMsgSponsorAccept
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSponsorAccept,
		therealblocksimulation.SimulateMsgSponsorAccept(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAdminAdd int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAdminAdd, &weightMsgAdminAdd, nil,
		func(_ *rand.Rand) {
			weightMsgAdminAdd = defaultWeightMsgAdminAdd
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAdminAdd,
		therealblocksimulation.SimulateMsgAdminAdd(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAdminDelete int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAdminDelete, &weightMsgAdminDelete, nil,
		func(_ *rand.Rand) {
			weightMsgAdminDelete = defaultWeightMsgAdminDelete
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAdminDelete,
		therealblocksimulation.SimulateMsgAdminDelete(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgNextStage int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgNextStage, &weightMsgNextStage, nil,
		func(_ *rand.Rand) {
			weightMsgNextStage = defaultWeightMsgNextStage
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgNextStage,
		therealblocksimulation.SimulateMsgNextStage(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgShareProfit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgShareProfit, &weightMsgShareProfit, nil,
		func(_ *rand.Rand) {
			weightMsgShareProfit = defaultWeightMsgShareProfit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgShareProfit,
		therealblocksimulation.SimulateMsgShareProfit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDraftProject int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDraftProject, &weightMsgUpdateDraftProject, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDraftProject = defaultWeightMsgUpdateDraftProject
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDraftProject,
		therealblocksimulation.SimulateMsgUpdateDraftProject(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
