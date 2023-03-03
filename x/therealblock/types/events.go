package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EmitEvent(ctx sdk.Context, eventType string, AttributeList ...string) {
	// Takes a list of attributes and emits an event
	// Example: EmitEvent(ctx, "project_created", "id", "1", "sponsor", "cosmos1...")
	var attributes []sdk.Attribute
	for i := 0; i < len(AttributeList); i += 2 {
		attributes = append(attributes, sdk.NewAttribute(AttributeList[i], AttributeList[i+1]))
	}
	ctx.EventManager().EmitEvent(sdk.NewEvent(eventType, attributes...))
}
