package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type TargetingGeoLocationZip struct {
	Country       *string  `json:"country,omitempty"`
	Key           *string  `json:"key,omitempty"`
	Name          *string  `json:"name,omitempty"`
	PrimaryCityID *core.ID `json:"primary_city_id,omitempty"`
	RegionID      *core.ID `json:"region_id,omitempty"`
}
