package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountTargetingUnified struct {
	AudienceSizeLowerBound *uint64   `json:"audience_size_lower_bound,omitempty"`
	AudienceSizeUpperBound *uint64   `json:"audience_size_upper_bound,omitempty"`
	ConversionLift         *float64  `json:"conversion_lift,omitempty"`
	Description            *string   `json:"description,omitempty"`
	ID                     *core.ID  `json:"id,omitempty"`
	Img                    *string   `json:"img,omitempty"`
	Info                   *string   `json:"info,omitempty"`
	InfoTitle              *string   `json:"info_title,omitempty"`
	IsRecommendation       *bool     `json:"is_recommendation,omitempty"`
	IsYouthAdsAgeGated     *bool     `json:"is_youth_ads_age_gated,omitempty"`
	Key                    *string   `json:"key,omitempty"`
	Link                   *string   `json:"link,omitempty"`
	Name                   *string   `json:"name,omitempty"`
	Parent                 *string   `json:"parent,omitempty"`
	Partner                *string   `json:"partner,omitempty"`
	Path                   *[]string `json:"path,omitempty"`
	PerformanceRating      *uint64   `json:"performance_rating,omitempty"`
	RawName                *string   `json:"raw_name,omitempty"`
	RecommendationModel    *string   `json:"recommendation_model,omitempty"`
	SearchInterestID       *core.ID  `json:"search_interest_id,omitempty"`
	Source                 *string   `json:"source,omitempty"`
	Spend                  *float64  `json:"spend,omitempty"`
	Type                   *string   `json:"type,omitempty"`
	Valid                  *bool     `json:"valid,omitempty"`
}
