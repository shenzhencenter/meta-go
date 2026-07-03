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
