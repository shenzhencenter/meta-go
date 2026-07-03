package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ContentBlockList struct {
	Business *Business `json:"business,omitempty"`
	ID       *core.ID  `json:"id,omitempty"`
	Name     *string   `json:"name,omitempty"`
}
