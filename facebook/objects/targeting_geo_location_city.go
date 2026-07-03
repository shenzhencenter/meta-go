package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var TargetingGeoLocationCityFields = struct {
	Country      string
	DistanceUnit string
	Key          string
	Name         string
	Radius       string
	Region       string
	RegionID     string
}{
	Country:      "country",
	DistanceUnit: "distance_unit",
	Key:          "key",
	Name:         "name",
	Radius:       "radius",
	Region:       "region",
	RegionID:     "region_id",
}
