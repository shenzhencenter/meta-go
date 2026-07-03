package objects

type ConversionHealthKPI struct {
	HealthIndicator                      *string  `json:"health_indicator,omitempty"`
	ImpactedBrowsersMatchRate            *float64 `json:"impacted_browsers_match_rate,omitempty"`
	ImpactedBrowsersMatchRateMomTrend    *float64 `json:"impacted_browsers_match_rate_mom_trend,omitempty"`
	ImpactedBrowsersTrafficShare         *float64 `json:"impacted_browsers_traffic_share,omitempty"`
	ImpactedBrowsersTrafficShareMomTrend *float64 `json:"impacted_browsers_traffic_share_mom_trend,omitempty"`
	MatchRate                            *float64 `json:"match_rate,omitempty"`
	MatchRateMomTrend                    *float64 `json:"match_rate_mom_trend,omitempty"`
	MatchRateVerticalBenchmark           *float64 `json:"match_rate_vertical_benchmark,omitempty"`
	MatchRateVsBenchmarkMomTrend         *float64 `json:"match_rate_vs_benchmark_mom_trend,omitempty"`
}
