package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Status struct {
	Event       *Event                  `json:"event,omitempty"`
	From        *map[string]interface{} `json:"from,omitempty"`
	ID          *core.ID                `json:"id,omitempty"`
	Message     *string                 `json:"message,omitempty"`
	Place       *Place                  `json:"place,omitempty"`
	UpdatedTime *core.Time              `json:"updated_time,omitempty"`
}

var StatusFields = struct {
	Event       string
	From        string
	ID          string
	Message     string
	Place       string
	UpdatedTime string
}{
	Event:       "event",
	From:        "from",
	ID:          "id",
	Message:     "message",
	Place:       "place",
	UpdatedTime: "updated_time",
}
