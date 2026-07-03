package objects

type MCExperienceConfigForApi struct {
	IsCampaignEnabled         *bool   `json:"is_campaign_enabled,omitempty"`
	IsTermsSigned             *bool   `json:"is_terms_signed,omitempty"`
	IsUserManuallyToggleMcOff *bool   `json:"is_user_manually_toggle_mc_off,omitempty"`
	MerchantType              *string `json:"merchant_type,omitempty"`
}
