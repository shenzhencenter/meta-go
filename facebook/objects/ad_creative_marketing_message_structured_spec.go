package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeMarketingMessageStructuredSpec struct {
	AssetCustomization            *map[string]interface{}   `json:"asset_customization,omitempty"`
	Autoreply                     *map[string]interface{}   `json:"autoreply,omitempty"`
	Buttons                       *[]map[string]interface{} `json:"buttons,omitempty"`
	ConversationTemplateID        *core.ID                  `json:"conversation_template_id,omitempty"`
	Footer                        *string                   `json:"footer,omitempty"`
	Greeting                      *string                   `json:"greeting,omitempty"`
	IsOptimizedText               *bool                     `json:"is_optimized_text,omitempty"`
	Language                      *string                   `json:"language,omitempty"`
	Offer                         *map[string]interface{}   `json:"offer,omitempty"`
	ReferencedAdgroupID           *core.ID                  `json:"referenced_adgroup_id,omitempty"`
	WhatsAppBusinessPhoneNumberID *core.ID                  `json:"whats_app_business_phone_number_id,omitempty"`
}

var AdCreativeMarketingMessageStructuredSpecFields = struct {
	AssetCustomization            string
	Autoreply                     string
	Buttons                       string
	ConversationTemplateID        string
	Footer                        string
	Greeting                      string
	IsOptimizedText               string
	Language                      string
	Offer                         string
	ReferencedAdgroupID           string
	WhatsAppBusinessPhoneNumberID string
}{
	AssetCustomization:            "asset_customization",
	Autoreply:                     "autoreply",
	Buttons:                       "buttons",
	ConversationTemplateID:        "conversation_template_id",
	Footer:                        "footer",
	Greeting:                      "greeting",
	IsOptimizedText:               "is_optimized_text",
	Language:                      "language",
	Offer:                         "offer",
	ReferencedAdgroupID:           "referenced_adgroup_id",
	WhatsAppBusinessPhoneNumberID: "whats_app_business_phone_number_id",
}
