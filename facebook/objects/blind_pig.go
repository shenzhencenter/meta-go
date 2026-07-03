package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type BlindPig struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}
