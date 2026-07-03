package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdAccount struct {
	AccountID                          *core.ID                            `json:"account_id,omitempty"`
	AccountStatus                      *uint64                             `json:"account_status,omitempty"`
	AdAccountPromotableObjects         *AdAccountPromotableObjects         `json:"ad_account_promotable_objects,omitempty"`
	Age                                *float64                            `json:"age,omitempty"`
	AgencyClientDeclaration            *AgencyClientDeclaration            `json:"agency_client_declaration,omitempty"`
	AllCapabilities                    *[]string                           `json:"all_capabilities,omitempty"`
	AmountSpent                        *string                             `json:"amount_spent,omitempty"`
	AttributionSpec                    *[]AttributionSpec                  `json:"attribution_spec,omitempty"`
	Balance                            *string                             `json:"balance,omitempty"`
	BrandSafetyContentFilterLevels     *[]string                           `json:"brand_safety_content_filter_levels,omitempty"`
	Business                           *Business                           `json:"business,omitempty"`
	BusinessCity                       *string                             `json:"business_city,omitempty"`
	BusinessCountryCode                *string                             `json:"business_country_code,omitempty"`
	BusinessName                       *string                             `json:"business_name,omitempty"`
	BusinessState                      *string                             `json:"business_state,omitempty"`
	BusinessStreet                     *string                             `json:"business_street,omitempty"`
	BusinessStreet2                    *string                             `json:"business_street2,omitempty"`
	BusinessZip                        *string                             `json:"business_zip,omitempty"`
	Capabilities                       *[]string                           `json:"capabilities,omitempty"`
	CreatedTime                        *time.Time                          `json:"created_time,omitempty"`
	Currency                           *string                             `json:"currency,omitempty"`
	CustomAudienceInfo                 *CustomAudienceGroup                `json:"custom_audience_info,omitempty"`
	DefaultDsaBeneficiary              *string                             `json:"default_dsa_beneficiary,omitempty"`
	DefaultDsaPayor                    *string                             `json:"default_dsa_payor,omitempty"`
	DisableReason                      *uint64                             `json:"disable_reason,omitempty"`
	EndAdvertiser                      *string                             `json:"end_advertiser,omitempty"`
	EndAdvertiserName                  *string                             `json:"end_advertiser_name,omitempty"`
	ExistingCustomers                  *[]string                           `json:"existing_customers,omitempty"`
	ExpiredFundingSourceDetails        *FundingSourceDetails               `json:"expired_funding_source_details,omitempty"`
	ExtendedCreditInvoiceGroup         *ExtendedCreditInvoiceGroup         `json:"extended_credit_invoice_group,omitempty"`
	FailedDeliveryChecks               *[]DeliveryCheck                    `json:"failed_delivery_checks,omitempty"`
	FbEntity                           *uint64                             `json:"fb_entity,omitempty"`
	FundingSource                      *string                             `json:"funding_source,omitempty"`
	FundingSourceDetails               *FundingSourceDetails               `json:"funding_source_details,omitempty"`
	HasMigratedPermissions             *bool                               `json:"has_migrated_permissions,omitempty"`
	HasPageAuthorizedAdaccount         *bool                               `json:"has_page_authorized_adaccount,omitempty"`
	ID                                 *core.ID                            `json:"id,omitempty"`
	IoNumber                           *string                             `json:"io_number,omitempty"`
	IsAttributionSpecSystemDefault     *bool                               `json:"is_attribution_spec_system_default,omitempty"`
	IsBaSkipDelayedEligible            *bool                               `json:"is_ba_skip_delayed_eligible,omitempty"`
	IsDirectDealsEnabled               *bool                               `json:"is_direct_deals_enabled,omitempty"`
	IsInX3dsAuthorizationEnabledMarket *bool                               `json:"is_in_3ds_authorization_enabled_market,omitempty"`
	IsNotificationsEnabled             *bool                               `json:"is_notifications_enabled,omitempty"`
	IsPersonal                         *uint64                             `json:"is_personal,omitempty"`
	IsPrepayAccount                    *bool                               `json:"is_prepay_account,omitempty"`
	IsTaxIDRequired                    *bool                               `json:"is_tax_id_required,omitempty"`
	LiableAddress                      *CRMAddress                         `json:"liable_address,omitempty"`
	LineNumbers                        *[]int                              `json:"line_numbers,omitempty"`
	MarketingMessagesSettings          *AdAccountMarketingMessagesSettings `json:"marketing_messages_settings,omitempty"`
	MediaAgency                        *string                             `json:"media_agency,omitempty"`
	MinCampaignGroupSpendCap           *string                             `json:"min_campaign_group_spend_cap,omitempty"`
	MinDailyBudget                     *uint64                             `json:"min_daily_budget,omitempty"`
	Name                               *string                             `json:"name,omitempty"`
	OffsiteCloSignalStatus             *int                                `json:"offsite_clo_signal_status,omitempty"`
	OffsitePixelsTosAccepted           *bool                               `json:"offsite_pixels_tos_accepted,omitempty"`
	OpportunityScore                   *float64                            `json:"opportunity_score,omitempty"`
	OpportunityScoreWeight             *int                                `json:"opportunity_score_weight,omitempty"`
	Owner                              *string                             `json:"owner,omitempty"`
	OwnerBusiness                      *Business                           `json:"owner_business,omitempty"`
	Partner                            *string                             `json:"partner,omitempty"`
	RfSpec                             *ReachFrequencySpec                 `json:"rf_spec,omitempty"`
	SendBillToAddress                  *CRMAddress                         `json:"send_bill_to_address,omitempty"`
	ShowCheckoutExperience             *bool                               `json:"show_checkout_experience,omitempty"`
	SoldToAddress                      *CRMAddress                         `json:"sold_to_address,omitempty"`
	SpendCap                           *string                             `json:"spend_cap,omitempty"`
	TaxID                              *core.ID                            `json:"tax_id,omitempty"`
	TaxIDStatus                        *uint64                             `json:"tax_id_status,omitempty"`
	TaxIDType                          *string                             `json:"tax_id_type,omitempty"`
	TimezoneID                         *core.ID                            `json:"timezone_id,omitempty"`
	TimezoneName                       *string                             `json:"timezone_name,omitempty"`
	TimezoneOffsetHoursUtc             *float64                            `json:"timezone_offset_hours_utc,omitempty"`
	TosAccepted                        *map[string]int                     `json:"tos_accepted,omitempty"`
	UserAccessExpireTime               *time.Time                          `json:"user_access_expire_time,omitempty"`
	UserTasks                          *[]string                           `json:"user_tasks,omitempty"`
	UserTosAccepted                    *map[string]int                     `json:"user_tos_accepted,omitempty"`
	ViewableBusiness                   *Business                           `json:"viewable_business,omitempty"`
}
