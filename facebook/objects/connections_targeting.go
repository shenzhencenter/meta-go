package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ConnectionsTargeting struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}
