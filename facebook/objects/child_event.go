package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ChildEvent struct {
	EndTime   *string  `json:"end_time,omitempty"`
	ID        *core.ID `json:"id,omitempty"`
	StartTime *string  `json:"start_time,omitempty"`
	TicketURI *string  `json:"ticket_uri,omitempty"`
}

var ChildEventFields = struct {
	EndTime   string
	ID        string
	StartTime string
	TicketURI string
}{
	EndTime:   "end_time",
	ID:        "id",
	StartTime: "start_time",
	TicketURI: "ticket_uri",
}
