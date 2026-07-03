package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WorkUserBadges struct {
	Category    *string  `json:"category,omitempty"`
	Description *string  `json:"description,omitempty"`
	Icon        *string  `json:"icon,omitempty"`
	ID          *core.ID `json:"id,omitempty"`
	Name        *string  `json:"name,omitempty"`
}
