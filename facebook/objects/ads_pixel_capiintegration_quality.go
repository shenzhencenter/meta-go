package objects

type AdsPixelCAPIIntegrationQuality struct {
	Acr                          *map[string]interface{}   `json:"acr,omitempty"`
	DataFreshness                *map[string]interface{}   `json:"data_freshness,omitempty"`
	DedupeKeyFeedback            *[]map[string]interface{} `json:"dedupe_key_feedback,omitempty"`
	EventAdSets                  *[]map[string]interface{} `json:"event_ad_sets,omitempty"`
	EventCoverage                *map[string]interface{}   `json:"event_coverage,omitempty"`
	EventMatchQuality            *map[string]interface{}   `json:"event_match_quality,omitempty"`
	EventName                    *string                   `json:"event_name,omitempty"`
	EventPotentialAlyAcrIncrease *map[string]interface{}   `json:"event_potential_aly_acr_increase,omitempty"`
	EventSpend                   *map[string]interface{}   `json:"event_spend,omitempty"`
}
