package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type WhatsAppBusinessHealthStatus struct {
	AdditionalInfo    *[]string                 `json:"additional_info,omitempty"`
	CanReceiveCallSip *string                   `json:"can_receive_call_sip,omitempty"`
	CanSendMessage    *string                   `json:"can_send_message,omitempty"`
	EntityType        *string                   `json:"entity_type,omitempty"`
	Errors            *[]map[string]interface{} `json:"errors,omitempty"`
	ID                *core.ID                  `json:"id,omitempty"`
}
