package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPixelEventSuggestionRule struct {
	X7dVolume  *int       `json:"7d_volume,omitempty"`
	Dismissed  *bool      `json:"dismissed,omitempty"`
	EndTime    *core.Time `json:"end_time,omitempty"`
	EventType  *string    `json:"event_type,omitempty"`
	Rank       *int       `json:"rank,omitempty"`
	Rule       *string    `json:"rule,omitempty"`
	SampleUrls *[]string  `json:"sample_urls,omitempty"`
	StartTime  *core.Time `json:"start_time,omitempty"`
}

var AdsPixelEventSuggestionRuleFields = struct {
	X7dVolume  string
	Dismissed  string
	EndTime    string
	EventType  string
	Rank       string
	Rule       string
	SampleUrls string
	StartTime  string
}{
	X7dVolume:  "7d_volume",
	Dismissed:  "dismissed",
	EndTime:    "end_time",
	EventType:  "event_type",
	Rank:       "rank",
	Rule:       "rule",
	SampleUrls: "sample_urls",
	StartTime:  "start_time",
}
