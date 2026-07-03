package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
)

type WhatsAppBusinessAccount struct {
	AccountReviewStatus                   *string                                                             `json:"account_review_status,omitempty"`
	Analytics                             *map[string]interface{}                                             `json:"analytics,omitempty"`
	AuthInternationalRateEligibility      *map[string]interface{}                                             `json:"auth_international_rate_eligibility,omitempty"`
	BusinessVerificationStatus            *enums.WhatsappbusinessaccountBusinessVerificationStatus            `json:"business_verification_status,omitempty"`
	Country                               *string                                                             `json:"country,omitempty"`
	CreationTime                          *int                                                                `json:"creation_time,omitempty"`
	Currency                              *string                                                             `json:"currency,omitempty"`
	DisableMarketingMessagesOnCloudAPI    *bool                                                               `json:"disable_marketing_messages_on_cloud_api,omitempty"`
	HealthStatus                          *WhatsAppBusinessHealthStatusForMessageSend                         `json:"health_status,omitempty"`
	ID                                    *core.ID                                                            `json:"id,omitempty"`
	IsEnabledForInsights                  *bool                                                               `json:"is_enabled_for_insights,omitempty"`
	IsSharedWithPartners                  *bool                                                               `json:"is_shared_with_partners,omitempty"`
	LinkedCommerceAccount                 *CommerceMerchantSettings                                           `json:"linked_commerce_account,omitempty"`
	MarketingMessagesAdAccount            *map[string]interface{}                                             `json:"marketing_messages_ad_account,omitempty"`
	MarketingMessagesLiteAPIStatus        *string                                                             `json:"marketing_messages_lite_api_status,omitempty"`
	MarketingMessagesOnboardingStatus     *string                                                             `json:"marketing_messages_onboarding_status,omitempty"`
	MessageTemplateNamespace              *string                                                             `json:"message_template_namespace,omitempty"`
	Name                                  *string                                                             `json:"name,omitempty"`
	OnBehalfOfBusinessInfo                *map[string]interface{}                                             `json:"on_behalf_of_business_info,omitempty"`
	OwnerBusiness                         *Business                                                           `json:"owner_business,omitempty"`
	OwnerBusinessInfo                     *map[string]interface{}                                             `json:"owner_business_info,omitempty"`
	OwnershipType                         *string                                                             `json:"ownership_type,omitempty"`
	PrimaryBusinessLocation               *string                                                             `json:"primary_business_location,omitempty"`
	PrimaryFundingID                      *core.ID                                                            `json:"primary_funding_id,omitempty"`
	PurchaseOrderNumber                   *string                                                             `json:"purchase_order_number,omitempty"`
	Status                                *string                                                             `json:"status,omitempty"`
	TemplateAutoArchivalEnabled           *bool                                                               `json:"template_auto_archival_enabled,omitempty"`
	TimezoneID                            *core.ID                                                            `json:"timezone_id,omitempty"`
	WhatsappBusinessManagerMessagingLimit *enums.WhatsappbusinessaccountWhatsappBusinessManagerMessagingLimit `json:"whatsapp_business_manager_messaging_limit,omitempty"`
}
