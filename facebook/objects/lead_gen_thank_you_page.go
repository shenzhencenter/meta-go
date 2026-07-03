package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type LeadGenThankYouPage struct {
	Body                *string                       `json:"body,omitempty"`
	BusinessPhoneNumber *string                       `json:"business_phone_number,omitempty"`
	ButtonText          *string                       `json:"button_text,omitempty"`
	ButtonType          *string                       `json:"button_type,omitempty"`
	CountryCode         *string                       `json:"country_code,omitempty"`
	EnableMessenger     *bool                         `json:"enable_messenger,omitempty"`
	GatedFile           *LeadGenThankYouPageGatedFile `json:"gated_file,omitempty"`
	ID                  *core.ID                      `json:"id,omitempty"`
	LeadGenUseCase      *string                       `json:"lead_gen_use_case,omitempty"`
	Status              *string                       `json:"status,omitempty"`
	Title               *string                       `json:"title,omitempty"`
	WebsiteURL          *string                       `json:"website_url,omitempty"`
}

var LeadGenThankYouPageFields = struct {
	Body                string
	BusinessPhoneNumber string
	ButtonText          string
	ButtonType          string
	CountryCode         string
	EnableMessenger     string
	GatedFile           string
	ID                  string
	LeadGenUseCase      string
	Status              string
	Title               string
	WebsiteURL          string
}{
	Body:                "body",
	BusinessPhoneNumber: "business_phone_number",
	ButtonText:          "button_text",
	ButtonType:          "button_type",
	CountryCode:         "country_code",
	EnableMessenger:     "enable_messenger",
	GatedFile:           "gated_file",
	ID:                  "id",
	LeadGenUseCase:      "lead_gen_use_case",
	Status:              "status",
	Title:               "title",
	WebsiteURL:          "website_url",
}
