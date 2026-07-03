package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdKeywordStats struct {
	Actions            *[]AdsActionStats `json:"actions,omitempty"`
	Clicks             *uint64           `json:"clicks,omitempty"`
	CostPerTotalAction *float64          `json:"cost_per_total_action,omitempty"`
	CostPerUniqueClick *float64          `json:"cost_per_unique_click,omitempty"`
	Cpc                *float64          `json:"cpc,omitempty"`
	Cpm                *float64          `json:"cpm,omitempty"`
	Cpp                *float64          `json:"cpp,omitempty"`
	Ctr                *float64          `json:"ctr,omitempty"`
	Frequency          *float64          `json:"frequency,omitempty"`
	ID                 *core.ID          `json:"id,omitempty"`
	Impressions        *uint64           `json:"impressions,omitempty"`
	Name               *string           `json:"name,omitempty"`
	Reach              *uint64           `json:"reach,omitempty"`
	Spend              *float64          `json:"spend,omitempty"`
	TotalActions       *uint64           `json:"total_actions,omitempty"`
	TotalUniqueActions *uint64           `json:"total_unique_actions,omitempty"`
	UniqueActions      *[]AdsActionStats `json:"unique_actions,omitempty"`
	UniqueClicks       *uint64           `json:"unique_clicks,omitempty"`
	UniqueCtr          *float64          `json:"unique_ctr,omitempty"`
	UniqueImpressions  *uint64           `json:"unique_impressions,omitempty"`
}
