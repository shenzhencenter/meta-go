package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
