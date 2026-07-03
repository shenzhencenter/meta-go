package objects

type AdCreativeAppExperienceSpec struct {
	RevealDetails  *map[string]interface{} `json:"reveal_details,omitempty"`
	ShowSpotlights *map[string]interface{} `json:"show_spotlights,omitempty"`
	WebsiteSummary *map[string]interface{} `json:"website_summary,omitempty"`
}
