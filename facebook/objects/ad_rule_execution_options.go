package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdRuleExecutionOptions struct {
	Field    *string                               `json:"field,omitempty"`
	Operator *enums.AdruleexecutionoptionsOperator `json:"operator,omitempty"`
	Value    *map[string]interface{}               `json:"value,omitempty"`
}

var AdRuleExecutionOptionsFields = struct {
	Field    string
	Operator string
	Value    string
}{
	Field:    "field",
	Operator: "operator",
	Value:    "value",
}
