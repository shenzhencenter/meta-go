package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPixelStatsResult struct {
	Aggregation *string          `json:"aggregation,omitempty"`
	Data        *[]AdsPixelStats `json:"data,omitempty"`
	StartTime   *core.Time       `json:"start_time,omitempty"`
}

var AdsPixelStatsResultFields = struct {
	Aggregation string
	Data        string
	StartTime   string
}{
	Aggregation: "aggregation",
	Data:        "data",
	StartTime:   "start_time",
}
