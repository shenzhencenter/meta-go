package objects

type LiveVideoAdBreakConfig struct {
	DefaultAdBreakDuration       *uint64 `json:"default_ad_break_duration,omitempty"`
	FailureReasonPollingInterval *uint64 `json:"failure_reason_polling_interval,omitempty"`
	FirstBreakEligibleSecs       *uint64 `json:"first_break_eligible_secs,omitempty"`
	GuideURL                     *string `json:"guide_url,omitempty"`
	IsEligibleToOnboard          *bool   `json:"is_eligible_to_onboard,omitempty"`
	IsEnabled                    *bool   `json:"is_enabled,omitempty"`
	OnboardingURL                *string `json:"onboarding_url,omitempty"`
	PreparingDuration            *uint64 `json:"preparing_duration,omitempty"`
	TimeBetweenAdBreaksSecs      *uint64 `json:"time_between_ad_breaks_secs,omitempty"`
	ViewerCountThreshold         *uint64 `json:"viewer_count_threshold,omitempty"`
}
