package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageChangeProposal struct {
	AcceptanceStatus   *string             `json:"acceptance_status,omitempty"`
	Category           *string             `json:"category,omitempty"`
	ID                 *core.ID            `json:"id,omitempty"`
	UpcomingChangeInfo *PageUpcomingChange `json:"upcoming_change_info,omitempty"`
}
