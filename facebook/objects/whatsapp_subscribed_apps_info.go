package objects

type WhatsappSubscribedAppsInfo struct {
	PageWhatsappNumber *string                   `json:"page_whatsapp_number,omitempty"`
	SubscribedApps     *[]map[string]interface{} `json:"subscribed_apps,omitempty"`
}

var WhatsappSubscribedAppsInfoFields = struct {
	PageWhatsappNumber string
	SubscribedApps     string
}{
	PageWhatsappNumber: "page_whatsapp_number",
	SubscribedApps:     "subscribed_apps",
}
