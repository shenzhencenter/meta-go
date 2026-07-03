package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type ResellerGuidance struct {
	AdAccountFirstSpendDate *string    `json:"ad_account_first_spend_date,omitempty"`
	AdAccountID             *core.ID   `json:"ad_account_id,omitempty"`
	AdoptedGuidanceL7d      *[]string  `json:"adopted_guidance_l7d,omitempty"`
	AdvertiserName          *string    `json:"advertiser_name,omitempty"`
	AttributedToResellerL7d *bool      `json:"attributed_to_reseller_l7d,omitempty"`
	AvailableGuidance       *[]string  `json:"available_guidance,omitempty"`
	GuidanceAdoptionRateL7d *float64   `json:"guidance_adoption_rate_l7d,omitempty"`
	NurturedByResellerL7d   *bool      `json:"nurtured_by_reseller_l7d,omitempty"`
	PlanningAgencyName      *string    `json:"planning_agency_name,omitempty"`
	RecommendationTime      *time.Time `json:"recommendation_time,omitempty"`
	ReportingDs             *string    `json:"reporting_ds,omitempty"`
	Reseller                *Business  `json:"reseller,omitempty"`
	RevenueL30d             *float64   `json:"revenue_l30d,omitempty"`
	UltimateAdvertiserName  *string    `json:"ultimate_advertiser_name,omitempty"`
}
