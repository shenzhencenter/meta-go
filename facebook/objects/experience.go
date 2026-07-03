package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type Experience struct {
	Description *string                 `json:"description,omitempty"`
	From        *map[string]interface{} `json:"from,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	With        *[]User                 `json:"with,omitempty"`
}
