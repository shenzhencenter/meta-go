package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsSegments struct {
	DailyAudienceSize     *int      `json:"daily_audience_size,omitempty"`
	DailyImpressions      *int      `json:"daily_impressions,omitempty"`
	Description           *string   `json:"description,omitempty"`
	ID                    *core.ID  `json:"id,omitempty"`
	Name                  *string   `json:"name,omitempty"`
	Path                  *[]string `json:"path,omitempty"`
	Popularity            *float64  `json:"popularity,omitempty"`
	ProjectedCpm          *int      `json:"projected_cpm,omitempty"`
	ProjectedDailyRevenue *int      `json:"projected_daily_revenue,omitempty"`
}
