package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CTXDFOObjectiveDefaults struct {
	Objective        *string  `json:"objective,omitempty"`
	OptimizationGoal *string  `json:"optimization_goal,omitempty"`
	PageID           *core.ID `json:"page_id,omitempty"`
}
