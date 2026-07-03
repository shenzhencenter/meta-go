package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type TargetingGeoLocationGeoEntities struct {
	Country  *string  `json:"country,omitempty"`
	Key      *string  `json:"key,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Region   *string  `json:"region,omitempty"`
	RegionID *core.ID `json:"region_id,omitempty"`
}
