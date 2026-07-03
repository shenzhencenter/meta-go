package objects

type WhatsAppBusinessHealthStatusForMessageSend struct {
	CanSendMessage *string                         `json:"can_send_message,omitempty"`
	Entities       *[]WhatsAppBusinessHealthStatus `json:"entities,omitempty"`
}

var WhatsAppBusinessHealthStatusForMessageSendFields = struct {
	CanSendMessage string
	Entities       string
}{
	CanSendMessage: "can_send_message",
	Entities:       "entities",
}
