package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WorkSkill struct {
	ID   *core.ID `json:"id,omitempty"`
	Name *string  `json:"name,omitempty"`
}
