package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Place struct {
	ID            *core.ID  `json:"id,omitempty"`
	Location      *Location `json:"location,omitempty"`
	Name          *string   `json:"name,omitempty"`
	OverallRating *float64  `json:"overall_rating,omitempty"`
}

var PlaceFields = struct {
	ID            string
	Location      string
	Name          string
	OverallRating string
}{
	ID:            "id",
	Location:      "location",
	Name:          "name",
	OverallRating: "overall_rating",
}
