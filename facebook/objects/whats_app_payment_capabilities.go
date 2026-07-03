package objects

type WhatsAppPaymentCapabilities struct {
	IsEnabled                *bool                     `json:"is_enabled,omitempty"`
	PaymentCapabilityDetails *[]map[string]interface{} `json:"payment_capability_details,omitempty"`
}
