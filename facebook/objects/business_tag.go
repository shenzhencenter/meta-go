package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessTag struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}
