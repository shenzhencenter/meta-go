package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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
	TimeChangesRequested           *core.Time                `json:"time_changes_requested,omitempty"`
	TimeCreated                    *core.Time                `json:"time_created,omitempty"`
	TimeUpdated                    *core.Time                `json:"time_updated,omitempty"`
	TimeZone                       *string                   `json:"time_zone,omitempty"`
	UsedResellerLink               *bool                     `json:"used_reseller_link,omitempty"`
	UserID                         *core.ID                  `json:"user_id,omitempty"`
	UserName                       *string                   `json:"user_name,omitempty"`
	Vertical                       *string                   `json:"vertical,omitempty"`
	VerticalV2                     *string                   `json:"vertical_v2,omitempty"`
	ViewedByReseller               *bool                     `json:"viewed_by_reseller,omitempty"`
	ZipCode                        *string                   `json:"zip_code,omitempty"`
}

var ChinaBusinessOnboardingVettingRequestFields = struct {
	AdAccountCreationRequestStatus string
	AdAccountLimit                 string
	AdAccountNumber                string
	AdAccountsInfo                 string
	AdvertiserBusinessID           string
	AdvertiserBusinessName         string
	BusinessManagerID              string
	BusinessRegistration           string
	BusinessRegistrationID         string
	BusinessVerificationStatus     string
	ChineseAddress                 string
	ChineseLegalEntityName         string
	City                           string
	Contact                        string
	CouponCode                     string
	DisapproveReason               string
	EnglishBusinessName            string
	ID                             string
	OfficialWebsiteURL             string
	OrgAdAccountCount              string
	PaymentType                    string
	PlanningAgencyID               string
	PlanningAgencyName             string
	PromotableAppIds               string
	PromotablePageIds              string
	PromotablePages                string
	PromotableUrls                 string
	RequestChangesReason           string
	ReviewedUser                   string
	SpendLimit                     string
	Status                         string
	Subvertical                    string
	SubverticalV2                  string
	SupportingDocument             string
	TimeChangesRequested           string
	TimeCreated                    string
	TimeUpdated                    string
	TimeZone                       string
	UsedResellerLink               string
	UserID                         string
	UserName                       string
	Vertical                       string
	VerticalV2                     string
	ViewedByReseller               string
	ZipCode                        string
}{
	AdAccountCreationRequestStatus: "ad_account_creation_request_status",
	AdAccountLimit:                 "ad_account_limit",
	AdAccountNumber:                "ad_account_number",
	AdAccountsInfo:                 "ad_accounts_info",
	AdvertiserBusinessID:           "advertiser_business_id",
	AdvertiserBusinessName:         "advertiser_business_name",
	BusinessManagerID:              "business_manager_id",
	BusinessRegistration:           "business_registration",
	BusinessRegistrationID:         "business_registration_id",
	BusinessVerificationStatus:     "business_verification_status",
	ChineseAddress:                 "chinese_address",
	ChineseLegalEntityName:         "chinese_legal_entity_name",
	City:                           "city",
	Contact:                        "contact",
	CouponCode:                     "coupon_code",
	DisapproveReason:               "disapprove_reason",
	EnglishBusinessName:            "english_business_name",
	ID:                             "id",
	OfficialWebsiteURL:             "official_website_url",
	OrgAdAccountCount:              "org_ad_account_count",
	PaymentType:                    "payment_type",
	PlanningAgencyID:               "planning_agency_id",
	PlanningAgencyName:             "planning_agency_name",
	PromotableAppIds:               "promotable_app_ids",
	PromotablePageIds:              "promotable_page_ids",
	PromotablePages:                "promotable_pages",
	PromotableUrls:                 "promotable_urls",
	RequestChangesReason:           "request_changes_reason",
	ReviewedUser:                   "reviewed_user",
	SpendLimit:                     "spend_limit",
	Status:                         "status",
	Subvertical:                    "subvertical",
	SubverticalV2:                  "subvertical_v2",
	SupportingDocument:             "supporting_document",
	TimeChangesRequested:           "time_changes_requested",
	TimeCreated:                    "time_created",
	TimeUpdated:                    "time_updated",
	TimeZone:                       "time_zone",
	UsedResellerLink:               "used_reseller_link",
	UserID:                         "user_id",
	UserName:                       "user_name",
	Vertical:                       "vertical",
	VerticalV2:                     "vertical_v2",
	ViewedByReseller:               "viewed_by_reseller",
	ZipCode:                        "zip_code",
}
