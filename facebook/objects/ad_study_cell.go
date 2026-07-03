package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdStudyCell struct {
	AdEntitiesCount     *uint64    `json:"ad_entities_count,omitempty"`
	AdIds               *[]core.ID `json:"ad_ids,omitempty"`
	ControlPercentage   *float64   `json:"control_percentage,omitempty"`
	ID                  *core.ID   `json:"id,omitempty"`
	Name                *string    `json:"name,omitempty"`
	TreatmentPercentage *float64   `json:"treatment_percentage,omitempty"`
}

var AdStudyCellFields = struct {
	AdEntitiesCount     string
	AdIds               string
	ControlPercentage   string
	ID                  string
	Name                string
	TreatmentPercentage string
}{
	AdEntitiesCount:     "ad_entities_count",
	AdIds:               "ad_ids",
	ControlPercentage:   "control_percentage",
	ID:                  "id",
	Name:                "name",
	TreatmentPercentage: "treatment_percentage",
}
