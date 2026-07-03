package objects

type CopyrightAttributionInsights struct {
	L7AttributionPageView       *int     `json:"l7_attribution_page_view,omitempty"`
	L7AttributionPageViewDelta  *float64 `json:"l7_attribution_page_view_delta,omitempty"`
	L7AttributionVideoView      *int     `json:"l7_attribution_video_view,omitempty"`
	L7AttributionVideoViewDelta *float64 `json:"l7_attribution_video_view_delta,omitempty"`
	MetricsEndingDate           *string  `json:"metrics_ending_date,omitempty"`
}

var CopyrightAttributionInsightsFields = struct {
	L7AttributionPageView       string
	L7AttributionPageViewDelta  string
	L7AttributionVideoView      string
	L7AttributionVideoViewDelta string
	MetricsEndingDate           string
}{
	L7AttributionPageView:       "l7_attribution_page_view",
	L7AttributionPageViewDelta:  "l7_attribution_page_view_delta",
	L7AttributionVideoView:      "l7_attribution_video_view",
	L7AttributionVideoViewDelta: "l7_attribution_video_view_delta",
	MetricsEndingDate:           "metrics_ending_date",
}
