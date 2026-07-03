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

var AudienceOverlapFields = struct {
	EstimatedReach string
	ID             string
	Name           string
	Overlap        string
}{
	EstimatedReach: "estimated_reach",
	ID:             "id",
	Name:           "name",
	Overlap:        "overlap",
}
