package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CPASSetupPageStructureProgress struct {
	ID     *core.ID                  `json:"id,omitempty"`
	Issues *[]map[string]interface{} `json:"issues,omitempty"`
	Name   *string                   `json:"name,omitempty"`
}
