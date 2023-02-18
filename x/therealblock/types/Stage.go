package types

import (
	"strconv"
	"strings"
)

type Stages []*Stage

// ParseStageNormalized will parse out a list of stages separated by commas
// Expected format: "{stage-name0}{stage-allocation0},...,{stage-nameN}{stage-allocationN}"
func ParseStageNormalized(stageStr string) (Stages, error) {
	//TODO assert argument format correctness
	splitStages := strings.Split(stageStr, ",")
	parsedStages := make([]*Stage, 0)

	for _, stage := range splitStages {
		splitStage := strings.Split(stage, ":")
		parsedString, err := strconv.ParseUint(splitStage[1], 10, 64)
		if err != nil {
			return parsedStages, err
		}
		parsedStages = append(parsedStages, &Stage{Name: splitStage[0], Allocation: parsedString})
	}
	return parsedStages, nil
}
