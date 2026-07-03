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
