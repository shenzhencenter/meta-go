package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdAccountOptimizationGoalsAEMv2Eligibility struct {
	IsDisabled       *bool                                                             `json:"is_disabled,omitempty"`
	OptimizationGoal *enums.Adaccountoptimizationgoalsaemv2eligibilityOptimizationGoal `json:"optimization_goal,omitempty"`
}
