package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-project [target] [stages]",
		Short: "Broadcast message create-project",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTarget, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argStages := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateProject(
				clientCtx.GetFromAddress().String(),
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
