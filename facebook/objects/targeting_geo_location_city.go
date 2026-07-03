package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type TargetingGeoLocationCity struct {
	Country      *string  `json:"country,omitempty"`
	DistanceUnit *string  `json:"distance_unit,omitempty"`
	Key          *string  `json:"key,omitempty"`
	Name         *string  `json:"name,omitempty"`
	Radius       *uint64  `json:"radius,omitempty"`
	Region       *string  `json:"region,omitempty"`
	RegionID     *core.ID `json:"region_id,omitempty"`
}
