package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
