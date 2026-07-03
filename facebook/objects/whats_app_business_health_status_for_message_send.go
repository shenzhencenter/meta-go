package objects

type WhatsAppBusinessHealthStatusForMessageSend struct {
	CanSendMessage *string                         `json:"can_send_message,omitempty"`
	Entities       *[]WhatsAppBusinessHealthStatus `json:"entities,omitempty"`
}
