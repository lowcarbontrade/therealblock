package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateDraftProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-draft-project [project-id] [target] [stages]",
		Short: "Broadcast message update-draft-project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProjectId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			argTarget, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			argStages, err := types.ParseStageNormalized(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDraftProject(
				clientCtx.GetFromAddress().String(),
				argProjectId,
				argTarget,
				argStages,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
