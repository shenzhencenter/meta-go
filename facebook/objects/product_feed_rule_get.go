package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductFeedRuleGet struct {
	Attribute *string                           `json:"attribute,omitempty"`
	ID        *core.ID                          `json:"id,omitempty"`
	Params    *[]map[string]interface{}         `json:"params,omitempty"`
	RuleType  *enums.ProductfeedrulegetRuleType `json:"rule_type,omitempty"`
}
