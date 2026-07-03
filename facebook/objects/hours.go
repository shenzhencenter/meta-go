package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type Hours struct {
	ID              *core.ID `json:"id,omitempty"`
	PermanentStatus *string  `json:"permanent_status,omitempty"`
}
