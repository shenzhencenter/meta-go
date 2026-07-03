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
