package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type GeoGatingPolicy struct {
	AfterSchedule  *string    `json:"after_schedule,omitempty"`
	ExcludeCountry *[]string  `json:"exclude_country,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	IncludeCountry *[]string  `json:"include_country,omitempty"`
	Name           *string    `json:"name,omitempty"`
	ValidFrom      *time.Time `json:"valid_from,omitempty"`
	ValidUntil     *time.Time `json:"valid_until,omitempty"`
}
