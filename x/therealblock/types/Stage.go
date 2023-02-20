package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

type Stages []*Stage

// ParseStageNormalized will parse out a list of stages separated by commas
// Expected format: "{stage-name0}{stage-allocation0},...,{stage-nameN}{stage-allocationN}"
// stage-allocations are of sdk.Coin type
// Example "stage1:1000token,stage2:2000token"
func ParseStageNormalized(stageStr string) (Stages, error) {
	//TODO assert argument format correctness
	splitStages := strings.Split(stageStr, ",")
	parsedStages := make([]*Stage, 0)

	for _, stage := range splitStages {
		splitStage := strings.Split(stage, ":")
		allocationParsed, err := sdk.ParseCoinNormalized(splitStage[1])
		if err != nil {
			return parsedStages, err
		}
		parsedStages = append(parsedStages, &Stage{Name: splitStage[0], Allocation: allocationParsed})
	}
	return parsedStages, nil
}
