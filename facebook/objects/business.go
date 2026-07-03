package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type Business struct {
	BlockOfflineAnalytics                      *bool                                                `json:"block_offline_analytics,omitempty"`
	CollaborativeAdsManagedPartnerBusinessInfo *ManagedPartnerBusiness                              `json:"collaborative_ads_managed_partner_business_info,omitempty"`
	CollaborativeAdsManagedPartnerEligibility  *BusinessManagedPartnerEligibility                   `json:"collaborative_ads_managed_partner_eligibility,omitempty"`
	CollaborativeAdsPartnerPremiumOptions      *BusinessPartnerPremiumOptions                       `json:"collaborative_ads_partner_premium_options,omitempty"`
	CreatedBy                                  *map[string]interface{}                              `json:"created_by,omitempty"`
	CreatedTime                                *time.Time                                           `json:"created_time,omitempty"`
	ExtendedUpdatedTime                        *time.Time                                           `json:"extended_updated_time,omitempty"`
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
	UpdatedTime                                *time.Time                                           `json:"updated_time,omitempty"`
	UserAccessExpireTime                       *time.Time                                           `json:"user_access_expire_time,omitempty"`
	VerificationStatus                         *enums.BusinessVerificationStatus                    `json:"verification_status,omitempty"`
	Vertical                                   *string                                              `json:"vertical,omitempty"`
	VerticalID                                 *core.ID                                             `json:"vertical_id,omitempty"`
	WhatsappBusinessManagerMessagingLimit      *enums.BusinessWhatsappBusinessManagerMessagingLimit `json:"whatsapp_business_manager_messaging_limit,omitempty"`
}
