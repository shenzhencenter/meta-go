package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Business struct {
	BlockOfflineAnalytics                      *bool                                                `json:"block_offline_analytics,omitempty"`
	CollaborativeAdsManagedPartnerBusinessInfo *ManagedPartnerBusiness                              `json:"collaborative_ads_managed_partner_business_info,omitempty"`
	CollaborativeAdsManagedPartnerEligibility  *BusinessManagedPartnerEligibility                   `json:"collaborative_ads_managed_partner_eligibility,omitempty"`
	CollaborativeAdsPartnerPremiumOptions      *BusinessPartnerPremiumOptions                       `json:"collaborative_ads_partner_premium_options,omitempty"`
	CreatedBy                                  *map[string]interface{}                              `json:"created_by,omitempty"`
	CreatedTime                                *core.Time                                           `json:"created_time,omitempty"`
	ExtendedUpdatedTime                        *core.Time                                           `json:"extended_updated_time,omitempty"`
	ID                                         *core.ID                                             `json:"id,omitempty"`
	IsHidden                                   *bool                                                `json:"is_hidden,omitempty"`
	Link                                       *string                                              `json:"link,omitempty"`
	MarketingMessagesOnboardingStatus          *MarketingMessagesOnboardingStatus                   `json:"marketing_messages_onboarding_status,omitempty"`
	Name                                       *string                                              `json:"name,omitempty"`
	PrimaryPage                                *Page                                                `json:"primary_page,omitempty"`
	ProfilePictureURI                          *string                                              `json:"profile_picture_uri,omitempty"`
	TimezoneID                                 *core.ID                                             `json:"timezone_id,omitempty"`
	TwoFactorType                              *string                                              `json:"two_factor_type,omitempty"`
	UpdatedBy                                  *map[string]interface{}                              `json:"updated_by,omitempty"`
	UpdatedTime                                *core.Time                                           `json:"updated_time,omitempty"`
	UserAccessExpireTime                       *core.Time                                           `json:"user_access_expire_time,omitempty"`
	VerificationStatus                         *enums.BusinessVerificationStatus                    `json:"verification_status,omitempty"`
	Vertical                                   *string                                              `json:"vertical,omitempty"`
	VerticalID                                 *core.ID                                             `json:"vertical_id,omitempty"`
	WhatsappBusinessManagerMessagingLimit      *enums.BusinessWhatsappBusinessManagerMessagingLimit `json:"whatsapp_business_manager_messaging_limit,omitempty"`
}

var BusinessFields = struct {
	BlockOfflineAnalytics                      string
	CollaborativeAdsManagedPartnerBusinessInfo string
	CollaborativeAdsManagedPartnerEligibility  string
	CollaborativeAdsPartnerPremiumOptions      string
	CreatedBy                                  string
	CreatedTime                                string
	ExtendedUpdatedTime                        string
	ID                                         string
	IsHidden                                   string
	Link                                       string
	MarketingMessagesOnboardingStatus          string
	Name                                       string
	PrimaryPage                                string
	ProfilePictureURI                          string
	TimezoneID                                 string
	TwoFactorType                              string
	UpdatedBy                                  string
	UpdatedTime                                string
	UserAccessExpireTime                       string
	VerificationStatus                         string
	Vertical                                   string
	VerticalID                                 string
	WhatsappBusinessManagerMessagingLimit      string
}{
	BlockOfflineAnalytics:                      "block_offline_analytics",
	CollaborativeAdsManagedPartnerBusinessInfo: "collaborative_ads_managed_partner_business_info",
	CollaborativeAdsManagedPartnerEligibility:  "collaborative_ads_managed_partner_eligibility",
	CollaborativeAdsPartnerPremiumOptions:      "collaborative_ads_partner_premium_options",
	CreatedBy:                                  "created_by",
	CreatedTime:                                "created_time",
	ExtendedUpdatedTime:                        "extended_updated_time",
	ID:                                         "id",
	IsHidden:                                   "is_hidden",
	Link:                                       "link",
	MarketingMessagesOnboardingStatus:          "marketing_messages_onboarding_status",
	Name:                                       "name",
	PrimaryPage:                                "primary_page",
	ProfilePictureURI:                          "profile_picture_uri",
	TimezoneID:                                 "timezone_id",
	TwoFactorType:                              "two_factor_type",
	UpdatedBy:                                  "updated_by",
	UpdatedTime:                                "updated_time",
	UserAccessExpireTime:                       "user_access_expire_time",
	VerificationStatus:                         "verification_status",
	Vertical:                                   "vertical",
	VerticalID:                                 "vertical_id",
	WhatsappBusinessManagerMessagingLimit:      "whatsapp_business_manager_messaging_limit",
}
