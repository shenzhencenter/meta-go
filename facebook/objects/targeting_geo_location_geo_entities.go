package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type TargetingGeoLocationGeoEntities struct {
	Country  *string  `json:"country,omitempty"`
	Key      *string  `json:"key,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Region   *string  `json:"region,omitempty"`
	RegionID *core.ID `json:"region_id,omitempty"`
}

var TargetingGeoLocationGeoEntitiesFields = struct {
	Country  string
	Key      string
	Name     string
	Region   string
	RegionID string
}{
	Country:  "country",
	Key:      "key",
	Name:     "name",
	Region:   "region",
	RegionID: "region_id",
}
