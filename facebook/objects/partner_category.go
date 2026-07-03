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

var PartnerCategoryFields = struct {
	ApproximateCount string
	Country          string
	Description      string
	Details          string
	ID               string
	IsPrivate        string
	Name             string
	ParentCategory   string
	Source           string
	Status           string
	TargetingType    string
}{
	ApproximateCount: "approximate_count",
	Country:          "country",
	Description:      "description",
	Details:          "details",
	ID:               "id",
	IsPrivate:        "is_private",
	Name:             "name",
	ParentCategory:   "parent_category",
	Source:           "source",
	Status:           "status",
	TargetingType:    "targeting_type",
}
