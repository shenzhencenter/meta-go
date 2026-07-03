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

var PageChangeProposalFields = struct {
	AcceptanceStatus   string
	Category           string
	ID                 string
	UpcomingChangeInfo string
}{
	AcceptanceStatus:   "acceptance_status",
	Category:           "category",
	ID:                 "id",
	UpcomingChangeInfo: "upcoming_change_info",
}
