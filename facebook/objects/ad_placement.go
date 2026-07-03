package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdPlacement struct {
	BundleID            *core.ID                `json:"bundle_id,omitempty"`
	DisplayFormat       *string                 `json:"display_format,omitempty"`
	ExternalPlacementID *core.ID                `json:"external_placement_id,omitempty"`
	GoogleDisplayFormat *string                 `json:"google_display_format,omitempty"`
	ID                  *core.ID                `json:"id,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	PlacementGroup      *map[string]interface{} `json:"placement_group,omitempty"`
	Platform            *string                 `json:"platform,omitempty"`
	Status              *string                 `json:"status,omitempty"`
}

var AdPlacementFields = struct {
	BundleID            string
	DisplayFormat       string
	ExternalPlacementID string
	GoogleDisplayFormat string
	ID                  string
	Name                string
	PlacementGroup      string
	Platform            string
	Status              string
}{
	BundleID:            "bundle_id",
	DisplayFormat:       "display_format",
	ExternalPlacementID: "external_placement_id",
	GoogleDisplayFormat: "google_display_format",
	ID:                  "id",
	Name:                "name",
	PlacementGroup:      "placement_group",
	Platform:            "platform",
	Status:              "status",
}
