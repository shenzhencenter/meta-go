package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
