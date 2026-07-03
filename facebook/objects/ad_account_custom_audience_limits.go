package objects

type AdAccountCustomAudienceLimits struct {
	AudienceUpdateQuotaInTotal      *int     `json:"audience_update_quota_in_total,omitempty"`
	AudienceUpdateQuotaLeft         *float64 `json:"audience_update_quota_left,omitempty"`
	HasHitAudienceUpdateLimit       *bool    `json:"has_hit_audience_update_limit,omitempty"`
	NextAudienceUpdateAvailableTime *string  `json:"next_audience_update_available_time,omitempty"`
	RateLimitResetTime              *string  `json:"rate_limit_reset_time,omitempty"`
}
