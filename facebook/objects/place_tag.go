package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PlaceTag struct {
	CreatedTime *core.Time `json:"created_time,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	Place       *Page      `json:"place,omitempty"`
}

var PlaceTagFields = struct {
	CreatedTime string
	ID          string
	Place       string
}{
	CreatedTime: "created_time",
	ID:          "id",
	Place:       "place",
}
