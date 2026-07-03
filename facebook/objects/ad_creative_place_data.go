package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativePlaceData struct {
	AddressString    *string  `json:"address_string,omitempty"`
	Label            *string  `json:"label,omitempty"`
	Latitude         *float64 `json:"latitude,omitempty"`
	LocationSourceID *core.ID `json:"location_source_id,omitempty"`
	Longitude        *float64 `json:"longitude,omitempty"`
	Type             *string  `json:"type,omitempty"`
}
