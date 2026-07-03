package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type MessengerDestinationPageWelcomeMessage struct {
	ID                     *core.ID   `json:"id,omitempty"`
	PageWelcomeMessageBody *string    `json:"page_welcome_message_body,omitempty"`
	PageWelcomeMessageType *string    `json:"page_welcome_message_type,omitempty"`
	TemplateName           *string    `json:"template_name,omitempty"`
	TimeCreated            *time.Time `json:"time_created,omitempty"`
	TimeLastUsed           *time.Time `json:"time_last_used,omitempty"`
}
