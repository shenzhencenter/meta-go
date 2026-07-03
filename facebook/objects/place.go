package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type Place struct {
	ID            *core.ID  `json:"id,omitempty"`
	Location      *Location `json:"location,omitempty"`
	Name          *string   `json:"name,omitempty"`
	OverallRating *float64  `json:"overall_rating,omitempty"`
}
