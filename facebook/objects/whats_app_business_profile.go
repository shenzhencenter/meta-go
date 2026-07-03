package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type WhatsAppBusinessProfile struct {
	ID                      *core.ID                `json:"id,omitempty"`
	NameVerification        *map[string]interface{} `json:"name_verification,omitempty"`
	WhatsappBusinessAPIData *map[string]interface{} `json:"whatsapp_business_api_data,omitempty"`
}

var WhatsAppBusinessProfileFields = struct {
	ID                      string
	NameVerification        string
	WhatsappBusinessAPIData string
}{
	ID:                      "id",
	NameVerification:        "name_verification",
	WhatsappBusinessAPIData: "whatsapp_business_api_data",
}
