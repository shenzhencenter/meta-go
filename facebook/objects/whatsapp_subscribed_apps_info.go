package objects

type WhatsappSubscribedAppsInfo struct {
	PageWhatsappNumber *string                   `json:"page_whatsapp_number,omitempty"`
	SubscribedApps     *[]map[string]interface{} `json:"subscribed_apps,omitempty"`
}
