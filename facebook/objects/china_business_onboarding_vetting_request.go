package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type ChinaBusinessOnboardingVettingRequest struct {
	AdAccountCreationRequestStatus *string                   `json:"ad_account_creation_request_status,omitempty"`
	AdAccountLimit                 *int                      `json:"ad_account_limit,omitempty"`
	AdAccountNumber                *string                   `json:"ad_account_number,omitempty"`
	AdAccountsInfo                 *[]map[string]interface{} `json:"ad_accounts_info,omitempty"`
	AdvertiserBusinessID           *core.ID                  `json:"advertiser_business_id,omitempty"`
	AdvertiserBusinessName         *string                   `json:"advertiser_business_name,omitempty"`
	BusinessManagerID              *core.ID                  `json:"business_manager_id,omitempty"`
	BusinessRegistration           *string                   `json:"business_registration,omitempty"`
	BusinessRegistrationID         *core.ID                  `json:"business_registration_id,omitempty"`
	BusinessVerificationStatus     *string                   `json:"business_verification_status,omitempty"`
	ChineseAddress                 *string                   `json:"chinese_address,omitempty"`
	ChineseLegalEntityName         *string                   `json:"chinese_legal_entity_name,omitempty"`
	City                           *string                   `json:"city,omitempty"`
	Contact                        *string                   `json:"contact,omitempty"`
	CouponCode                     *string                   `json:"coupon_code,omitempty"`
	DisapproveReason               *string                   `json:"disapprove_reason,omitempty"`
	EnglishBusinessName            *string                   `json:"english_business_name,omitempty"`
	ID                             *core.ID                  `json:"id,omitempty"`
	OfficialWebsiteURL             *string                   `json:"official_website_url,omitempty"`
	OrgAdAccountCount              *int                      `json:"org_ad_account_count,omitempty"`
	PaymentType                    *string                   `json:"payment_type,omitempty"`
	PlanningAgencyID               *core.ID                  `json:"planning_agency_id,omitempty"`
	PlanningAgencyName             *string                   `json:"planning_agency_name,omitempty"`
	PromotableAppIds               *[]core.ID                `json:"promotable_app_ids,omitempty"`
	PromotablePageIds              *[]core.ID                `json:"promotable_page_ids,omitempty"`
	PromotablePages                *[]map[string]interface{} `json:"promotable_pages,omitempty"`
	PromotableUrls                 *[]string                 `json:"promotable_urls,omitempty"`
	RequestChangesReason           *string                   `json:"request_changes_reason,omitempty"`
	ReviewedUser                   *string                   `json:"reviewed_user,omitempty"`
	SpendLimit                     *int                      `json:"spend_limit,omitempty"`
	Status                         *string                   `json:"status,omitempty"`
	Subvertical                    *string                   `json:"subvertical,omitempty"`
	SubverticalV2                  *string                   `json:"subvertical_v2,omitempty"`
	SupportingDocument             *string                   `json:"supporting_document,omitempty"`
	TimeChangesRequested           *time.Time                `json:"time_changes_requested,omitempty"`
	TimeCreated                    *time.Time                `json:"time_created,omitempty"`
	TimeUpdated                    *time.Time                `json:"time_updated,omitempty"`
	TimeZone                       *string                   `json:"time_zone,omitempty"`
	UsedResellerLink               *bool                     `json:"used_reseller_link,omitempty"`
	UserID                         *core.ID                  `json:"user_id,omitempty"`
	UserName                       *string                   `json:"user_name,omitempty"`
	Vertical                       *string                   `json:"vertical,omitempty"`
	VerticalV2                     *string                   `json:"vertical_v2,omitempty"`
	ViewedByReseller               *bool                     `json:"viewed_by_reseller,omitempty"`
	ZipCode                        *string                   `json:"zip_code,omitempty"`
}
