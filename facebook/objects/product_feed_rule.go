package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ProductFeedRule struct {
	Attribute *string              `json:"attribute,omitempty"`
	ID        *core.ID             `json:"id,omitempty"`
	Params    *[]map[string]string `json:"params,omitempty"`
	RuleType  *string              `json:"rule_type,omitempty"`
}

var ProductFeedRuleFields = struct {
	Attribute string
	ID        string
	Params    string
	RuleType  string
}{
	Attribute: "attribute",
	ID:        "id",
	Params:    "params",
	RuleType:  "rule_type",
}
