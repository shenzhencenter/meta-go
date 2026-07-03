package objects

import (
	"time"
)

type InstantArticleInsightsQueryResult struct {
	Breakdowns *map[string]string `json:"breakdowns,omitempty"`
	Name       *string            `json:"name,omitempty"`
	Time       *time.Time         `json:"time,omitempty"`
	Value      *string            `json:"value,omitempty"`
}
