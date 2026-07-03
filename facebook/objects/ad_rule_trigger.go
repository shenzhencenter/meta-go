package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdRuleTrigger struct {
	Field    *string                      `json:"field,omitempty"`
	Operator *enums.AdruletriggerOperator `json:"operator,omitempty"`
	Type     *enums.AdruletriggerType     `json:"type,omitempty"`
	Value    *map[string]interface{}      `json:"value,omitempty"`
}
