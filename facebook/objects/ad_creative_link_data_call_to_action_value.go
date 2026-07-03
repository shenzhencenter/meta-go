package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeLinkDataCallToActionValue struct {
	AndroidURL            *string   `json:"android_url,omitempty"`
	AppDestination        *string   `json:"app_destination,omitempty"`
	AppLink               *string   `json:"app_link,omitempty"`
	Application           *string   `json:"application,omitempty"`
	EventID               *core.ID  `json:"event_id,omitempty"`
	IosURL                *string   `json:"ios_url,omitempty"`
	LandOnWhatsappCatalog *int      `json:"land_on_whatsapp_catalog,omitempty"`
	LandOnWhatsappProfile *int      `json:"land_on_whatsapp_profile,omitempty"`
	LeadGenFormID         *core.ID  `json:"lead_gen_form_id,omitempty"`
	Link                  *string   `json:"link,omitempty"`
	LinkCaption           *string   `json:"link_caption,omitempty"`
	LinkFormat            *string   `json:"link_format,omitempty"`
	ObjectStoreUrls       *[]string `json:"object_store_urls,omitempty"`
	Page                  *string   `json:"page,omitempty"`
	ProductLink           *string   `json:"product_link,omitempty"`
	WhatsappNumber        *string   `json:"whatsapp_number,omitempty"`
}

var AdCreativeLinkDataCallToActionValueFields = struct {
	AndroidURL            string
	AppDestination        string
	AppLink               string
	Application           string
	EventID               string
	IosURL                string
	LandOnWhatsappCatalog string
	LandOnWhatsappProfile string
	LeadGenFormID         string
	Link                  string
	LinkCaption           string
	LinkFormat            string
	ObjectStoreUrls       string
	Page                  string
	ProductLink           string
	WhatsappNumber        string
}{
	AndroidURL:            "android_url",
	AppDestination:        "app_destination",
	AppLink:               "app_link",
	Application:           "application",
	EventID:               "event_id",
	IosURL:                "ios_url",
	LandOnWhatsappCatalog: "land_on_whatsapp_catalog",
	LandOnWhatsappProfile: "land_on_whatsapp_profile",
	LeadGenFormID:         "lead_gen_form_id",
	Link:                  "link",
	LinkCaption:           "link_caption",
	LinkFormat:            "link_format",
	ObjectStoreUrls:       "object_store_urls",
	Page:                  "page",
	ProductLink:           "product_link",
	WhatsappNumber:        "whatsapp_number",
}
