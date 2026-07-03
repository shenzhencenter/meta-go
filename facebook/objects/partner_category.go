package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PartnerCategory struct {
	ApproximateCount *int     `json:"approximate_count,omitempty"`
	Country          *string  `json:"country,omitempty"`
	Description      *string  `json:"description,omitempty"`
	Details          *string  `json:"details,omitempty"`
	ID               *core.ID `json:"id,omitempty"`
	IsPrivate        *bool    `json:"is_private,omitempty"`
	Name             *string  `json:"name,omitempty"`
	ParentCategory   *string  `json:"parent_category,omitempty"`
	Source           *string  `json:"source,omitempty"`
	Status           *string  `json:"status,omitempty"`
	TargetingType    *string  `json:"targeting_type,omitempty"`
}
