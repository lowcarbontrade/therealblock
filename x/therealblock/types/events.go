package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

// TODO implement multiple attributes EmitEvent(context, ...NewAttribute)
func EmitEvent(ctx sdk.Context, eventType string, projectId uint64, investorAddr string) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute("event_type", eventType),
			sdk.NewAttribute(ProjectEventProjectKey, strconv.FormatUint(projectId, 10)),
			sdk.NewAttribute(ProjectEventProjectCreator, investorAddr),
		),
	)
}
