package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	RecommendationTime      *core.Time `json:"recommendation_time,omitempty"`
	ReportingDs             *string    `json:"reporting_ds,omitempty"`
	Reseller                *Business  `json:"reseller,omitempty"`
	RevenueL30d             *float64   `json:"revenue_l30d,omitempty"`
	UltimateAdvertiserName  *string    `json:"ultimate_advertiser_name,omitempty"`
}

var ResellerGuidanceFields = struct {
	AdAccountFirstSpendDate string
	AdAccountID             string
	AdoptedGuidanceL7d      string
	AdvertiserName          string
	AttributedToResellerL7d string
	AvailableGuidance       string
	GuidanceAdoptionRateL7d string
	NurturedByResellerL7d   string
	PlanningAgencyName      string
	RecommendationTime      string
	ReportingDs             string
	Reseller                string
	RevenueL30d             string
	UltimateAdvertiserName  string
}{
	AdAccountFirstSpendDate: "ad_account_first_spend_date",
	AdAccountID:             "ad_account_id",
	AdoptedGuidanceL7d:      "adopted_guidance_l7d",
	AdvertiserName:          "advertiser_name",
	AttributedToResellerL7d: "attributed_to_reseller_l7d",
	AvailableGuidance:       "available_guidance",
	GuidanceAdoptionRateL7d: "guidance_adoption_rate_l7d",
	NurturedByResellerL7d:   "nurtured_by_reseller_l7d",
	PlanningAgencyName:      "planning_agency_name",
	RecommendationTime:      "recommendation_time",
	ReportingDs:             "reporting_ds",
	Reseller:                "reseller",
	RevenueL30d:             "revenue_l30d",
	UltimateAdvertiserName:  "ultimate_advertiser_name",
}
