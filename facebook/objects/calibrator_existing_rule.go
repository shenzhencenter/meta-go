package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CalibratorExistingRule struct {
	X7dVolume      *int       `json:"7d_volume,omitempty"`
	CreationSource *string    `json:"creation_source,omitempty"`
	CreationTime   *core.Time `json:"creation_time,omitempty"`
	Creator        *string    `json:"creator,omitempty"`
	EventType      *string    `json:"event_type,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	Rule           *string    `json:"rule,omitempty"`
	RuleType       *string    `json:"rule_type,omitempty"`
	SampleUrls     *[]string  `json:"sample_urls,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Transforms     *[]string  `json:"transforms,omitempty"`
}

var CalibratorExistingRuleFields = struct {
	X7dVolume      string
	CreationSource string
	CreationTime   string
	Creator        string
	EventType      string
	ID             string
	Rule           string
	RuleType       string
	SampleUrls     string
	Status         string
	Transforms     string
}{
	X7dVolume:      "7d_volume",
	CreationSource: "creation_source",
	CreationTime:   "creation_time",
	Creator:        "creator",
	EventType:      "event_type",
	ID:             "id",
	Rule:           "rule",
	RuleType:       "rule_type",
	SampleUrls:     "sample_urls",
	Status:         "status",
	Transforms:     "transforms",
}
