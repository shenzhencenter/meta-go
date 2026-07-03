package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InstantArticleInsightsQueryResult struct {
	Breakdowns *map[string]string `json:"breakdowns,omitempty"`
	Name       *string            `json:"name,omitempty"`
	Time       *core.Time         `json:"time,omitempty"`
	Value      *string            `json:"value,omitempty"`
}

var InstantArticleInsightsQueryResultFields = struct {
	Breakdowns string
	Name       string
	Time       string
	Value      string
}{
	Breakdowns: "breakdowns",
	Name:       "name",
	Time:       "time",
	Value:      "value",
}
