package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdStudyObjective struct {
	ID                 *core.ID  `json:"id,omitempty"`
	IsPrimary          *bool     `json:"is_primary,omitempty"`
	LastUpdatedResults *string   `json:"last_updated_results,omitempty"`
	Name               *string   `json:"name,omitempty"`
	Results            *[]string `json:"results,omitempty"`
	Type               *string   `json:"type,omitempty"`
}
