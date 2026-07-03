package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CTWAWhatsAppNumbersInfo struct {
	CanManageWaFlows        *bool    `json:"can_manage_wa_flows,omitempty"`
	FormattedWhatsappNumber *string  `json:"formatted_whatsapp_number,omitempty"`
	IsBusinessNumber        *bool    `json:"is_business_number,omitempty"`
	IsCallingEnabled        *bool    `json:"is_calling_enabled,omitempty"`
	NumberCountryPrefix     *string  `json:"number_country_prefix,omitempty"`
	PageWhatsappNumberID    *core.ID `json:"page_whatsapp_number_id,omitempty"`
	WabaID                  *core.ID `json:"waba_id,omitempty"`
	WhatsappNumber          *string  `json:"whatsapp_number,omitempty"`
	WhatsappSmbDevice       *string  `json:"whatsapp_smb_device,omitempty"`
}

var CTWAWhatsAppNumbersInfoFields = struct {
	CanManageWaFlows        string
	FormattedWhatsappNumber string
	IsBusinessNumber        string
	IsCallingEnabled        string
	NumberCountryPrefix     string
	PageWhatsappNumberID    string
	WabaID                  string
	WhatsappNumber          string
	WhatsappSmbDevice       string
}{
	CanManageWaFlows:        "can_manage_wa_flows",
	FormattedWhatsappNumber: "formatted_whatsapp_number",
	IsBusinessNumber:        "is_business_number",
	IsCallingEnabled:        "is_calling_enabled",
	NumberCountryPrefix:     "number_country_prefix",
	PageWhatsappNumberID:    "page_whatsapp_number_id",
	WabaID:                  "waba_id",
	WhatsappNumber:          "whatsapp_number",
	WhatsappSmbDevice:       "whatsapp_smb_device",
}
