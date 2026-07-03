package objects

type CustomAudienceMatchRateInsights struct {
	EmailQuality         *string  `json:"email_quality,omitempty"`
	EmailUploadVolumePct *float64 `json:"email_upload_volume_pct,omitempty"`
	IsEligible           *bool    `json:"is_eligible,omitempty"`
	MadidQuality         *string  `json:"madid_quality,omitempty"`
	MadidUploadVolumePct *float64 `json:"madid_upload_volume_pct,omitempty"`
	MatchRateScore       *float64 `json:"match_rate_score,omitempty"`
	PhoneQuality         *string  `json:"phone_quality,omitempty"`
	PhoneUploadVolumePct *float64 `json:"phone_upload_volume_pct,omitempty"`
}

var CustomAudienceMatchRateInsightsFields = struct {
	EmailQuality         string
	EmailUploadVolumePct string
	IsEligible           string
	MadidQuality         string
	MadidUploadVolumePct string
	MatchRateScore       string
	PhoneQuality         string
	PhoneUploadVolumePct string
}{
	EmailQuality:         "email_quality",
	EmailUploadVolumePct: "email_upload_volume_pct",
	IsEligible:           "is_eligible",
	MadidQuality:         "madid_quality",
	MadidUploadVolumePct: "madid_upload_volume_pct",
	MatchRateScore:       "match_rate_score",
	PhoneQuality:         "phone_quality",
	PhoneUploadVolumePct: "phone_upload_volume_pct",
}
