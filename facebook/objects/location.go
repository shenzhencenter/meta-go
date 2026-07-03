package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Location struct {
	City        *string  `json:"city,omitempty"`
	CityID      *core.ID `json:"city_id,omitempty"`
	Country     *string  `json:"country,omitempty"`
	CountryCode *string  `json:"country_code,omitempty"`
	Latitude    *float64 `json:"latitude,omitempty"`
	LocatedIn   *string  `json:"located_in,omitempty"`
	Longitude   *float64 `json:"longitude,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Region      *string  `json:"region,omitempty"`
	RegionID    *core.ID `json:"region_id,omitempty"`
	State       *string  `json:"state,omitempty"`
	Street      *string  `json:"street,omitempty"`
	Zip         *string  `json:"zip,omitempty"`
}
