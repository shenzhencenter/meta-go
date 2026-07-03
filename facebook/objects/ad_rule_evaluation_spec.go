package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdRuleEvaluationSpec struct {
	EvaluationType *enums.AdruleevaluationspecEvaluationType `json:"evaluation_type,omitempty"`
	Filters        *[]AdRuleFilters                          `json:"filters,omitempty"`
	Trigger        *AdRuleTrigger                            `json:"trigger,omitempty"`
}
