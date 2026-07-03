package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdRuleExecutionOptions struct {
	Field    *string                               `json:"field,omitempty"`
	Operator *enums.AdruleexecutionoptionsOperator `json:"operator,omitempty"`
	Value    *map[string]interface{}               `json:"value,omitempty"`
}
