package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdRuleTrigger struct {
	Field    *string                      `json:"field,omitempty"`
	Operator *enums.AdruletriggerOperator `json:"operator,omitempty"`
	Type     *enums.AdruletriggerType     `json:"type,omitempty"`
	Value    *map[string]interface{}      `json:"value,omitempty"`
}

var AdRuleTriggerFields = struct {
	Field    string
	Operator string
	Type     string
	Value    string
}{
	Field:    "field",
	Operator: "operator",
	Type:     "type",
	Value:    "value",
}
