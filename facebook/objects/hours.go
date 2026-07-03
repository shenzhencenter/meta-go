package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type Hours struct {
	ID              *core.ID `json:"id,omitempty"`
	PermanentStatus *string  `json:"permanent_status,omitempty"`
}

var HoursFields = struct {
	ID              string
	PermanentStatus string
}{
	ID:              "id",
	PermanentStatus: "permanent_status",
}
