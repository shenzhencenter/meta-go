package objects

type AdCampaignMultiAds struct {
	EnrollStatus *string `json:"enroll_status,omitempty"`
	SourceType   *string `json:"source_type,omitempty"`
}

var AdCampaignMultiAdsFields = struct {
	EnrollStatus string
	SourceType   string
}{
	EnrollStatus: "enroll_status",
	SourceType:   "source_type",
}
