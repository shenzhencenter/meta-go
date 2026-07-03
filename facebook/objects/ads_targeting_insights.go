package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsTargetingInsights struct {
	AudienceSize   *int     `json:"audience_size,omitempty"`
	Clicks         *uint64  `json:"clicks,omitempty"`
	ConversionCost *float64 `json:"conversion_cost,omitempty"`
	Conversions    *uint64  `json:"conversions,omitempty"`
	Description    *string  `json:"description,omitempty"`
	ID             *core.ID `json:"id,omitempty"`
	Impressions    *string  `json:"impressions,omitempty"`
	Name           *string  `json:"name,omitempty"`
	Revenue        *float64 `json:"revenue,omitempty"`
	Spend          *float64 `json:"spend,omitempty"`
	Type           *string  `json:"type,omitempty"`
}

var AdsTargetingInsightsFields = struct {
	AudienceSize   string
	Clicks         string
	ConversionCost string
	Conversions    string
	Description    string
	ID             string
	Impressions    string
	Name           string
	Revenue        string
	Spend          string
	Type           string
}{
	AudienceSize:   "audience_size",
	Clicks:         "clicks",
	ConversionCost: "conversion_cost",
	Conversions:    "conversions",
	Description:    "description",
	ID:             "id",
	Impressions:    "impressions",
	Name:           "name",
	Revenue:        "revenue",
	Spend:          "spend",
	Type:           "type",
}
