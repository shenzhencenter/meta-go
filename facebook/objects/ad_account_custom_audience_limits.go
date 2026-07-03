package objects

type AdAccountCustomAudienceLimits struct {
	AudienceUpdateQuotaInTotal      *int     `json:"audience_update_quota_in_total,omitempty"`
	AudienceUpdateQuotaLeft         *float64 `json:"audience_update_quota_left,omitempty"`
	HasHitAudienceUpdateLimit       *bool    `json:"has_hit_audience_update_limit,omitempty"`
	NextAudienceUpdateAvailableTime *string  `json:"next_audience_update_available_time,omitempty"`
	RateLimitResetTime              *string  `json:"rate_limit_reset_time,omitempty"`
}

var AdAccountCustomAudienceLimitsFields = struct {
	AudienceUpdateQuotaInTotal      string
	AudienceUpdateQuotaLeft         string
	HasHitAudienceUpdateLimit       string
	NextAudienceUpdateAvailableTime string
	RateLimitResetTime              string
}{
	AudienceUpdateQuotaInTotal:      "audience_update_quota_in_total",
	AudienceUpdateQuotaLeft:         "audience_update_quota_left",
	HasHitAudienceUpdateLimit:       "has_hit_audience_update_limit",
	NextAudienceUpdateAvailableTime: "next_audience_update_available_time",
	RateLimitResetTime:              "rate_limit_reset_time",
}
