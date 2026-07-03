package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type EventRegistrationSetting struct {
	ID            *core.ID   `json:"id,omitempty"`
	Questions     *string    `json:"questions,omitempty"`
	TargetType    *string    `json:"target_type,omitempty"`
	TicketTierIds *[]core.ID `json:"ticket_tier_ids,omitempty"`
}
