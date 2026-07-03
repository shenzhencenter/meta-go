package objects

type MCExperienceConfigForApi struct {
	IsCampaignEnabled         *bool   `json:"is_campaign_enabled,omitempty"`
	IsTermsSigned             *bool   `json:"is_terms_signed,omitempty"`
	IsUserManuallyToggleMcOff *bool   `json:"is_user_manually_toggle_mc_off,omitempty"`
	MerchantType              *string `json:"merchant_type,omitempty"`
}

var MCExperienceConfigForApiFields = struct {
	IsCampaignEnabled         string
	IsTermsSigned             string
	IsUserManuallyToggleMcOff string
	MerchantType              string
}{
	IsCampaignEnabled:         "is_campaign_enabled",
	IsTermsSigned:             "is_terms_signed",
	IsUserManuallyToggleMcOff: "is_user_manually_toggle_mc_off",
	MerchantType:              "merchant_type",
}
