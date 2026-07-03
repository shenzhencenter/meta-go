package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type TargetingGeoLocationPlace struct {
	Country       *string  `json:"country,omitempty"`
	DistanceUnit  *string  `json:"distance_unit,omitempty"`
	Key           *string  `json:"key,omitempty"`
	Latitude      *float64 `json:"latitude,omitempty"`
	Longitude     *float64 `json:"longitude,omitempty"`
	Name          *string  `json:"name,omitempty"`
	PrimaryCityID *core.ID `json:"primary_city_id,omitempty"`
	Radius        *float64 `json:"radius,omitempty"`
	RegionID      *core.ID `json:"region_id,omitempty"`
}

var TargetingGeoLocationPlaceFields = struct {
	Country       string
	DistanceUnit  string
	Key           string
	Latitude      string
	Longitude     string
	Name          string
	PrimaryCityID string
	Radius        string
	RegionID      string
}{
	Country:       "country",
	DistanceUnit:  "distance_unit",
	Key:           "key",
	Latitude:      "latitude",
	Longitude:     "longitude",
	Name:          "name",
	PrimaryCityID: "primary_city_id",
	Radius:        "radius",
	RegionID:      "region_id",
}
