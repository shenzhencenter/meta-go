package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type ConnectionsTargeting struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}
