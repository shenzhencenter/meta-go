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
