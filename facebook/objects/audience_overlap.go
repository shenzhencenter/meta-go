package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AudienceOverlap struct {
	EstimatedReach *int     `json:"estimated_reach,omitempty"`
	ID             *core.ID `json:"id,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Overlap        *int     `json:"overlap,omitempty"`
}
