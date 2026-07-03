package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CTXDFOObjectiveDefaults struct {
	Objective        *string  `json:"objective,omitempty"`
	OptimizationGoal *string  `json:"optimization_goal,omitempty"`
	PageID           *core.ID `json:"page_id,omitempty"`
}

var CTXDFOObjectiveDefaultsFields = struct {
	Objective        string
	OptimizationGoal string
	PageID           string
}{
	Objective:        "objective",
	OptimizationGoal: "optimization_goal",
	PageID:           "page_id",
}
