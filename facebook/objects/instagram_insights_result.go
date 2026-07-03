package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InstagramInsightsResult struct {
	Description *string                   `json:"description,omitempty"`
	ID          *core.ID                  `json:"id,omitempty"`
	Name        *string                   `json:"name,omitempty"`
	Period      *string                   `json:"period,omitempty"`
	Title       *string                   `json:"title,omitempty"`
	TotalValue  *map[string]interface{}   `json:"total_value,omitempty"`
	Values      *[]InstagramInsightsValue `json:"values,omitempty"`
}

var InstagramInsightsResultFields = struct {
	Description string
	ID          string
	Name        string
	Period      string
	Title       string
	TotalValue  string
	Values      string
}{
	Description: "description",
	ID:          "id",
	Name:        "name",
	Period:      "period",
	Title:       "title",
	TotalValue:  "total_value",
	Values:      "values",
}
