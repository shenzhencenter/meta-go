package objects

type SalesRightsInventoryManagement struct {
	AvailableImpressions   *int      `json:"available_impressions,omitempty"`
	BookedImpressions      *int      `json:"booked_impressions,omitempty"`
	OverbookedImpressions  *int      `json:"overbooked_impressions,omitempty"`
	SupportedCountries     *[]string `json:"supported_countries,omitempty"`
	TotalImpressions       *int      `json:"total_impressions,omitempty"`
	UnavailableImpressions *int      `json:"unavailable_impressions,omitempty"`
	WarningMessages        *[]string `json:"warning_messages,omitempty"`
}
