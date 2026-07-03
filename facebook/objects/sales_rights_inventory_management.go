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

var SalesRightsInventoryManagementFields = struct {
	AvailableImpressions   string
	BookedImpressions      string
	OverbookedImpressions  string
	SupportedCountries     string
	TotalImpressions       string
	UnavailableImpressions string
	WarningMessages        string
}{
	AvailableImpressions:   "available_impressions",
	BookedImpressions:      "booked_impressions",
	OverbookedImpressions:  "overbooked_impressions",
	SupportedCountries:     "supported_countries",
	TotalImpressions:       "total_impressions",
	UnavailableImpressions: "unavailable_impressions",
	WarningMessages:        "warning_messages",
}
