package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CommerceMerchantSettings struct {
	CheckoutConfig                     *string                 `json:"checkout_config,omitempty"`
	ContactEmail                       *string                 `json:"contact_email,omitempty"`
	DisplayName                        *string                 `json:"display_name,omitempty"`
	FacebookChannel                    *map[string]interface{} `json:"facebook_channel,omitempty"`
	ID                                 *core.ID                `json:"id,omitempty"`
	InstagramChannel                   *map[string]interface{} `json:"instagram_channel,omitempty"`
	KoreaFtcListing                    *string                 `json:"korea_ftc_listing,omitempty"`
	MerchantPage                       *Profile                `json:"merchant_page,omitempty"`
	MerchantStatus                     *string                 `json:"merchant_status,omitempty"`
	OffsiteIabCheckoutEnabledCountries *[]string               `json:"offsite_iab_checkout_enabled_countries,omitempty"`
	PaymentProvider                    *string                 `json:"payment_provider,omitempty"`
	PrivacyPolicyLocalized             *string                 `json:"privacy_policy_localized,omitempty"`
	ShopsAdsSetup                      *map[string]interface{} `json:"shops_ads_setup,omitempty"`
	Terms                              *string                 `json:"terms,omitempty"`
}

var CommerceMerchantSettingsFields = struct {
	CheckoutConfig                     string
	ContactEmail                       string
	DisplayName                        string
	FacebookChannel                    string
	ID                                 string
	InstagramChannel                   string
	KoreaFtcListing                    string
	MerchantPage                       string
	MerchantStatus                     string
	OffsiteIabCheckoutEnabledCountries string
	PaymentProvider                    string
	PrivacyPolicyLocalized             string
	ShopsAdsSetup                      string
	Terms                              string
}{
	CheckoutConfig:                     "checkout_config",
	ContactEmail:                       "contact_email",
	DisplayName:                        "display_name",
	FacebookChannel:                    "facebook_channel",
	ID:                                 "id",
	InstagramChannel:                   "instagram_channel",
	KoreaFtcListing:                    "korea_ftc_listing",
	MerchantPage:                       "merchant_page",
	MerchantStatus:                     "merchant_status",
	OffsiteIabCheckoutEnabledCountries: "offsite_iab_checkout_enabled_countries",
	PaymentProvider:                    "payment_provider",
	PrivacyPolicyLocalized:             "privacy_policy_localized",
	ShopsAdsSetup:                      "shops_ads_setup",
	Terms:                              "terms",
}
