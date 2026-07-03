package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPivotRules struct {
	CreationTime *core.Time                `json:"creation_time,omitempty"`
	Creator      *Profile                  `json:"creator,omitempty"`
	Description  *string                   `json:"description,omitempty"`
	ID           *core.ID                  `json:"id,omitempty"`
	Name         *string                   `json:"name,omitempty"`
	Permission   *string                   `json:"permission,omitempty"`
	Rules        *[]map[string]interface{} `json:"rules,omitempty"`
	Scope        *string                   `json:"scope,omitempty"`
	UpdateBy     *Profile                  `json:"update_by,omitempty"`
	UpdateTime   *core.Time                `json:"update_time,omitempty"`
}

var AdsPivotRulesFields = struct {
	CreationTime string
	Creator      string
	Description  string
	ID           string
	Name         string
	Permission   string
	Rules        string
	Scope        string
	UpdateBy     string
	UpdateTime   string
}{
	CreationTime: "creation_time",
	Creator:      "creator",
	Description:  "description",
	ID:           "id",
	Name:         "name",
	Permission:   "permission",
	Rules:        "rules",
	Scope:        "scope",
	UpdateBy:     "update_by",
	UpdateTime:   "update_time",
}
