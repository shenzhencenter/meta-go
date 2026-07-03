package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WhatsAppBusinessHealthStatus struct {
	AdditionalInfo    *[]string                 `json:"additional_info,omitempty"`
	CanReceiveCallSip *string                   `json:"can_receive_call_sip,omitempty"`
	CanSendMessage    *string                   `json:"can_send_message,omitempty"`
	EntityType        *string                   `json:"entity_type,omitempty"`
	Errors            *[]map[string]interface{} `json:"errors,omitempty"`
	ID                *core.ID                  `json:"id,omitempty"`
}

var WhatsAppBusinessHealthStatusFields = struct {
	AdditionalInfo    string
	CanReceiveCallSip string
	CanSendMessage    string
	EntityType        string
	Errors            string
	ID                string
}{
	AdditionalInfo:    "additional_info",
	CanReceiveCallSip: "can_receive_call_sip",
	CanSendMessage:    "can_send_message",
	EntityType:        "entity_type",
	Errors:            "errors",
	ID:                "id",
}
