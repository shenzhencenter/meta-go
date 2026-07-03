package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdAccountAccountControlsPost struct {
	ErrorCode    *int     `json:"error_code,omitempty"`
	ErrorMessage *string  `json:"error_message,omitempty"`
	ID           *core.ID `json:"id,omitempty"`
	Success      *bool    `json:"success,omitempty"`
}
