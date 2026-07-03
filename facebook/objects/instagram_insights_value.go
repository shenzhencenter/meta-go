package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InstagramInsightsValue struct {
	EndTime *core.Time              `json:"end_time,omitempty"`
	Value   *map[string]interface{} `json:"value,omitempty"`
}

var InstagramInsightsValueFields = struct {
	EndTime string
	Value   string
}{
	EndTime: "end_time",
	Value:   "value",
}
