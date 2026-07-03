package objects

type AdCampaignFrequencyControlSpecs struct {
	Event        *string `json:"event,omitempty"`
	IntervalDays *uint64 `json:"interval_days,omitempty"`
	MaxFrequency *uint64 `json:"max_frequency,omitempty"`
	Type         *string `json:"type,omitempty"`
}
