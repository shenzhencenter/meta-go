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

var AdAccountTargetingUnifiedFields = struct {
	AudienceSizeLowerBound string
	AudienceSizeUpperBound string
	ConversionLift         string
	Description            string
	ID                     string
	Img                    string
	Info                   string
	InfoTitle              string
	IsRecommendation       string
	IsYouthAdsAgeGated     string
	Key                    string
	Link                   string
	Name                   string
	Parent                 string
	Partner                string
	Path                   string
	PerformanceRating      string
	RawName                string
	RecommendationModel    string
	SearchInterestID       string
	Source                 string
	Spend                  string
	Type                   string
	Valid                  string
}{
	AudienceSizeLowerBound: "audience_size_lower_bound",
	AudienceSizeUpperBound: "audience_size_upper_bound",
	ConversionLift:         "conversion_lift",
	Description:            "description",
	ID:                     "id",
	Img:                    "img",
	Info:                   "info",
	InfoTitle:              "info_title",
	IsRecommendation:       "is_recommendation",
	IsYouthAdsAgeGated:     "is_youth_ads_age_gated",
	Key:                    "key",
	Link:                   "link",
	Name:                   "name",
	Parent:                 "parent",
	Partner:                "partner",
	Path:                   "path",
	PerformanceRating:      "performance_rating",
	RawName:                "raw_name",
	RecommendationModel:    "recommendation_model",
	SearchInterestID:       "search_interest_id",
	Source:                 "source",
	Spend:                  "spend",
	Type:                   "type",
	Valid:                  "valid",
}
