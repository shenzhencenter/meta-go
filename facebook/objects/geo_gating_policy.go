package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type GeoGatingPolicy struct {
	AfterSchedule  *string    `json:"after_schedule,omitempty"`
	ExcludeCountry *[]string  `json:"exclude_country,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	IncludeCountry *[]string  `json:"include_country,omitempty"`
	Name           *string    `json:"name,omitempty"`
	ValidFrom      *core.Time `json:"valid_from,omitempty"`
	ValidUntil     *core.Time `json:"valid_until,omitempty"`
}

var GeoGatingPolicyFields = struct {
	AfterSchedule  string
	ExcludeCountry string
	ID             string
	IncludeCountry string
	Name           string
	ValidFrom      string
	ValidUntil     string
}{
	AfterSchedule:  "after_schedule",
	ExcludeCountry: "exclude_country",
	ID:             "id",
	IncludeCountry: "include_country",
	Name:           "name",
	ValidFrom:      "valid_from",
	ValidUntil:     "valid_until",
}
