package objects

import (
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type AdRuleFilters struct {
	Field    *string                      `json:"field,omitempty"`
	Operator *enums.AdrulefiltersOperator `json:"operator,omitempty"`
	Value    *map[string]interface{}      `json:"value,omitempty"`
}
