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
