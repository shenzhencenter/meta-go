package objects

import (
	"time"
)

type AdsPixelEventSuggestionRule struct {
	X7dVolume  *int       `json:"7d_volume,omitempty"`
	Dismissed  *bool      `json:"dismissed,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	EventType  *string    `json:"event_type,omitempty"`
	Rank       *int       `json:"rank,omitempty"`
	Rule       *string    `json:"rule,omitempty"`
	SampleUrls *[]string  `json:"sample_urls,omitempty"`
	StartTime  *time.Time `json:"start_time,omitempty"`
}
