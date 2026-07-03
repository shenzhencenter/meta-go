package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LifeEvent struct {
	Description *string    `json:"description,omitempty"`
	EndTime     *core.Time `json:"end_time,omitempty"`
	From        *Page      `json:"from,omitempty"`
	ID          *core.ID   `json:"id,omitempty"`
	IsHidden    *bool      `json:"is_hidden,omitempty"`
	StartTime   *core.Time `json:"start_time,omitempty"`
	Title       *string    `json:"title,omitempty"`
	UpdatedTime *core.Time `json:"updated_time,omitempty"`
}

var LifeEventFields = struct {
	Description string
	EndTime     string
	From        string
	ID          string
	IsHidden    string
	StartTime   string
	Title       string
	UpdatedTime string
}{
	Description: "description",
	EndTime:     "end_time",
	From:        "from",
	ID:          "id",
	IsHidden:    "is_hidden",
	StartTime:   "start_time",
	Title:       "title",
	UpdatedTime: "updated_time",
}
