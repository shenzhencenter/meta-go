package objects

type AdsPixelMicrodataStats struct {
	AllowedDomains                  *[]string                 `json:"allowed_domains,omitempty"`
	ErrorsStatsForTimeRanges        *[]map[string]interface{} `json:"errors_stats_for_time_ranges,omitempty"`
	HasValidEvents                  *bool                     `json:"has_valid_events,omitempty"`
	SuggestedAllowedDomainsCountMax *int                      `json:"suggested_allowed_domains_count_max,omitempty"`
	SuggestedTrustedDomains         *[]string                 `json:"suggested_trusted_domains,omitempty"`
}
