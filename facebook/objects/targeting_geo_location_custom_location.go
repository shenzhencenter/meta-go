package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type TargetingGeoLocationCustomLocation struct {
	AddressString *string  `json:"address_string,omitempty"`
	Country       *string  `json:"country,omitempty"`
	CountryGroup  *string  `json:"country_group,omitempty"`
	CustomType    *string  `json:"custom_type,omitempty"`
	DistanceUnit  *string  `json:"distance_unit,omitempty"`
	Key           *string  `json:"key,omitempty"`
	Latitude      *float64 `json:"latitude,omitempty"`
	Longitude     *float64 `json:"longitude,omitempty"`
	MaxPopulation *int     `json:"max_population,omitempty"`
	MinPopulation *int     `json:"min_population,omitempty"`
	Name          *string  `json:"name,omitempty"`
	PrimaryCityID *core.ID `json:"primary_city_id,omitempty"`
	Radius        *float64 `json:"radius,omitempty"`
	RegionID      *core.ID `json:"region_id,omitempty"`
}
