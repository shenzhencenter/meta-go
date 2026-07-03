package objects

type AdAccountMarketingMessagesSettings struct {
	WhatsappActivationStatus *string `json:"whatsapp_activation_status,omitempty"`
}

var AdAccountMarketingMessagesSettingsFields = struct {
	WhatsappActivationStatus string
}{
	WhatsappActivationStatus: "whatsapp_activation_status",
}
