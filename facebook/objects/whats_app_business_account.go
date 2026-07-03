package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
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

var WhatsAppBusinessAccountFields = struct {
	AccountReviewStatus                   string
	Analytics                             string
	AuthInternationalRateEligibility      string
	BusinessVerificationStatus            string
	Country                               string
	CreationTime                          string
	Currency                              string
	DisableMarketingMessagesOnCloudAPI    string
	HealthStatus                          string
	ID                                    string
	IsEnabledForInsights                  string
	IsSharedWithPartners                  string
	LinkedCommerceAccount                 string
	MarketingMessagesAdAccount            string
	MarketingMessagesLiteAPIStatus        string
	MarketingMessagesOnboardingStatus     string
	MessageTemplateNamespace              string
	Name                                  string
	OnBehalfOfBusinessInfo                string
	OwnerBusiness                         string
	OwnerBusinessInfo                     string
	OwnershipType                         string
	PrimaryBusinessLocation               string
	PrimaryFundingID                      string
	PurchaseOrderNumber                   string
	Status                                string
	TemplateAutoArchivalEnabled           string
	TimezoneID                            string
	WhatsappBusinessManagerMessagingLimit string
}{
	AccountReviewStatus:                   "account_review_status",
	Analytics:                             "analytics",
	AuthInternationalRateEligibility:      "auth_international_rate_eligibility",
	BusinessVerificationStatus:            "business_verification_status",
	Country:                               "country",
	CreationTime:                          "creation_time",
	Currency:                              "currency",
	DisableMarketingMessagesOnCloudAPI:    "disable_marketing_messages_on_cloud_api",
	HealthStatus:                          "health_status",
	ID:                                    "id",
	IsEnabledForInsights:                  "is_enabled_for_insights",
	IsSharedWithPartners:                  "is_shared_with_partners",
	LinkedCommerceAccount:                 "linked_commerce_account",
	MarketingMessagesAdAccount:            "marketing_messages_ad_account",
	MarketingMessagesLiteAPIStatus:        "marketing_messages_lite_api_status",
	MarketingMessagesOnboardingStatus:     "marketing_messages_onboarding_status",
	MessageTemplateNamespace:              "message_template_namespace",
	Name:                                  "name",
	OnBehalfOfBusinessInfo:                "on_behalf_of_business_info",
	OwnerBusiness:                         "owner_business",
	OwnerBusinessInfo:                     "owner_business_info",
	OwnershipType:                         "ownership_type",
	PrimaryBusinessLocation:               "primary_business_location",
	PrimaryFundingID:                      "primary_funding_id",
	PurchaseOrderNumber:                   "purchase_order_number",
	Status:                                "status",
	TemplateAutoArchivalEnabled:           "template_auto_archival_enabled",
	TimezoneID:                            "timezone_id",
	WhatsappBusinessManagerMessagingLimit: "whatsapp_business_manager_messaging_limit",
}
