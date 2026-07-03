package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdAccountOptimizationGoalsAEMv2Eligibility struct {
	IsDisabled       *bool                                                             `json:"is_disabled,omitempty"`
	OptimizationGoal *enums.Adaccountoptimizationgoalsaemv2eligibilityOptimizationGoal `json:"optimization_goal,omitempty"`
}

var AdAccountOptimizationGoalsAEMv2EligibilityFields = struct {
	IsDisabled       string
	OptimizationGoal string
}{
	IsDisabled:       "is_disabled",
	OptimizationGoal: "optimization_goal",
}
