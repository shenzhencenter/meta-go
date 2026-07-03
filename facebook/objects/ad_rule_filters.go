package objects

import (
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdRuleFilters struct {
	Field    *string                      `json:"field,omitempty"`
	Operator *enums.AdrulefiltersOperator `json:"operator,omitempty"`
	Value    *map[string]interface{}      `json:"value,omitempty"`
}
