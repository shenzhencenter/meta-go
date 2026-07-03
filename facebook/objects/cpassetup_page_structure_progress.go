package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CPASSetupPageStructureProgress struct {
	ID     *core.ID                  `json:"id,omitempty"`
	Issues *[]map[string]interface{} `json:"issues,omitempty"`
	Name   *string                   `json:"name,omitempty"`
}

var CPASSetupPageStructureProgressFields = struct {
	ID     string
	Issues string
	Name   string
}{
	ID:     "id",
	Issues: "issues",
	Name:   "name",
}
