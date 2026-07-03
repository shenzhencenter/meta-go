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

var AdsPixelCAPIIntegrationQualityFields = struct {
	Acr                          string
	DataFreshness                string
	DedupeKeyFeedback            string
	EventAdSets                  string
	EventCoverage                string
	EventMatchQuality            string
	EventName                    string
	EventPotentialAlyAcrIncrease string
	EventSpend                   string
}{
	Acr:                          "acr",
	DataFreshness:                "data_freshness",
	DedupeKeyFeedback:            "dedupe_key_feedback",
	EventAdSets:                  "event_ad_sets",
	EventCoverage:                "event_coverage",
	EventMatchQuality:            "event_match_quality",
	EventName:                    "event_name",
	EventPotentialAlyAcrIncrease: "event_potential_aly_acr_increase",
	EventSpend:                   "event_spend",
}
