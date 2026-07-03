package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdRuleEvaluationSpec struct {
	EvaluationType *enums.AdruleevaluationspecEvaluationType `json:"evaluation_type,omitempty"`
	Filters        *[]AdRuleFilters                          `json:"filters,omitempty"`
	Trigger        *AdRuleTrigger                            `json:"trigger,omitempty"`
}
