package objects

import (
	"time"
)

type InstagramInsightsValue struct {
	EndTime *time.Time              `json:"end_time,omitempty"`
	Value   *map[string]interface{} `json:"value,omitempty"`
}
