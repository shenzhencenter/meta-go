package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type EventRegistrationSetting struct {
	ID            *core.ID   `json:"id,omitempty"`
	Questions     *string    `json:"questions,omitempty"`
	TargetType    *string    `json:"target_type,omitempty"`
	TicketTierIds *[]core.ID `json:"ticket_tier_ids,omitempty"`
}

var EventRegistrationSettingFields = struct {
	ID            string
	Questions     string
	TargetType    string
	TicketTierIds string
}{
	ID:            "id",
	Questions:     "questions",
	TargetType:    "target_type",
	TicketTierIds: "ticket_tier_ids",
}
