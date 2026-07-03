package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ValueBasedEligibleSource struct {
	ID    *core.ID `json:"id,omitempty"`
	Title *string  `json:"title,omitempty"`
	Type  *string  `json:"type,omitempty"`
}
