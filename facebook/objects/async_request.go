package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AsyncRequest struct {
	ID     *core.ID `json:"id,omitempty"`
	Result *string  `json:"result,omitempty"`
	Status *int     `json:"status,omitempty"`
	Type   *int     `json:"type,omitempty"`
}
