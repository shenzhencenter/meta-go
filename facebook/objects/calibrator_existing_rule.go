package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type CalibratorExistingRule struct {
	X7dVolume      *int       `json:"7d_volume,omitempty"`
	CreationSource *string    `json:"creation_source,omitempty"`
	CreationTime   *time.Time `json:"creation_time,omitempty"`
	Creator        *string    `json:"creator,omitempty"`
	EventType      *string    `json:"event_type,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	Rule           *string    `json:"rule,omitempty"`
	RuleType       *string    `json:"rule_type,omitempty"`
	SampleUrls     *[]string  `json:"sample_urls,omitempty"`
	Status         *string    `json:"status,omitempty"`
	Transforms     *[]string  `json:"transforms,omitempty"`
}
