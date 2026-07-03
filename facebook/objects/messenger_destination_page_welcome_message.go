package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MessengerDestinationPageWelcomeMessage struct {
	ID                     *core.ID   `json:"id,omitempty"`
	PageWelcomeMessageBody *string    `json:"page_welcome_message_body,omitempty"`
	PageWelcomeMessageType *string    `json:"page_welcome_message_type,omitempty"`
	TemplateName           *string    `json:"template_name,omitempty"`
	TimeCreated            *core.Time `json:"time_created,omitempty"`
	TimeLastUsed           *core.Time `json:"time_last_used,omitempty"`
}

var MessengerDestinationPageWelcomeMessageFields = struct {
	ID                     string
	PageWelcomeMessageBody string
	PageWelcomeMessageType string
	TemplateName           string
	TimeCreated            string
	TimeLastUsed           string
}{
	ID:                     "id",
	PageWelcomeMessageBody: "page_welcome_message_body",
	PageWelcomeMessageType: "page_welcome_message_type",
	TemplateName:           "template_name",
	TimeCreated:            "time_created",
	TimeLastUsed:           "time_last_used",
}
