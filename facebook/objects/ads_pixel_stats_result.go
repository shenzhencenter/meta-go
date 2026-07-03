package objects

import (
	"time"
)

type AdsPixelStatsResult struct {
	Aggregation *string          `json:"aggregation,omitempty"`
	Data        *[]AdsPixelStats `json:"data,omitempty"`
	StartTime   *time.Time       `json:"start_time,omitempty"`
}
