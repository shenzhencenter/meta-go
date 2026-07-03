package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageUpcomingChange struct {
	ChangeType    *string             `json:"change_type,omitempty"`
	EffectiveTime *core.Time          `json:"effective_time,omitempty"`
	ID            *core.ID            `json:"id,omitempty"`
	Page          *Page               `json:"page,omitempty"`
	Proposal      *PageChangeProposal `json:"proposal,omitempty"`
	TimerStatus   *string             `json:"timer_status,omitempty"`
}

var PageUpcomingChangeFields = struct {
	ChangeType    string
	EffectiveTime string
	ID            string
	Page          string
	Proposal      string
	TimerStatus   string
}{
	ChangeType:    "change_type",
	EffectiveTime: "effective_time",
	ID:            "id",
	Page:          "page",
	Proposal:      "proposal",
	TimerStatus:   "timer_status",
}
