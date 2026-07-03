package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeWAMOWhatsAppIdentitySpec struct {
	WamoWhatsappIdentityID *core.ID `json:"wamo_whatsapp_identity_id,omitempty"`
	WhatsappPhoneNumber    *string  `json:"whatsapp_phone_number,omitempty"`
}

var AdCreativeWAMOWhatsAppIdentitySpecFields = struct {
	WamoWhatsappIdentityID string
	WhatsappPhoneNumber    string
}{
	WamoWhatsappIdentityID: "wamo_whatsapp_identity_id",
	WhatsappPhoneNumber:    "whatsapp_phone_number",
}
