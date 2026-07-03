package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type PageUpcomingChange struct {
	ChangeType    *string             `json:"change_type,omitempty"`
	EffectiveTime *time.Time          `json:"effective_time,omitempty"`
	ID            *core.ID            `json:"id,omitempty"`
	Page          *Page               `json:"page,omitempty"`
	Proposal      *PageChangeProposal `json:"proposal,omitempty"`
	TimerStatus   *string             `json:"timer_status,omitempty"`
}
