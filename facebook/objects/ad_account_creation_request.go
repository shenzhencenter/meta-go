package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdAccountCreationRequest struct {
	AdAccountsCurrency             *string                   `json:"ad_accounts_currency,omitempty"`
	AdAccountsInfo                 *[]map[string]interface{} `json:"ad_accounts_info,omitempty"`
	AdditionalComment              *string                   `json:"additional_comment,omitempty"`
	AddressInChinese               *string                   `json:"address_in_chinese,omitempty"`
	AddressInEnglish               *map[string]interface{}   `json:"address_in_english,omitempty"`
	AddressInLocalLanguage         *string                   `json:"address_in_local_language,omitempty"`
	AdvertiserBusiness             *Business                 `json:"advertiser_business,omitempty"`
	AppealReason                   *map[string]interface{}   `json:"appeal_reason,omitempty"`
	Business                       *Business                 `json:"business,omitempty"`
	BusinessRegistrationID         *core.ID                  `json:"business_registration_id,omitempty"`
	ChineseLegalEntityName         *string                   `json:"chinese_legal_entity_name,omitempty"`
	Contact                        *map[string]interface{}   `json:"contact,omitempty"`
	Creator                        *User                     `json:"creator,omitempty"`
	CreditCardID                   *core.ID                  `json:"credit_card_id,omitempty"`
	DisapprovalReasons             *[]map[string]interface{} `json:"disapproval_reasons,omitempty"`
	EnglishLegalEntityName         *string                   `json:"english_legal_entity_name,omitempty"`
	ExtendedCreditID               *core.ID                  `json:"extended_credit_id,omitempty"`
	ID                             *core.ID                  `json:"id,omitempty"`
	IsSmb                          *bool                     `json:"is_smb,omitempty"`
	IsTest                         *bool                     `json:"is_test,omitempty"`
	LegalEntityNameInLocalLanguage *string                   `json:"legal_entity_name_in_local_language,omitempty"`
	OeRequestID                    *core.ID                  `json:"oe_request_id,omitempty"`
	OfficialWebsiteURL             *string                   `json:"official_website_url,omitempty"`
	PlanningAgencyBusiness         *Business                 `json:"planning_agency_business,omitempty"`
	PlanningAgencyBusinessID       *core.ID                  `json:"planning_agency_business_id,omitempty"`
	PromotableAppIds               *[]core.ID                `json:"promotable_app_ids,omitempty"`
	PromotablePageIds              *[]core.ID                `json:"promotable_page_ids,omitempty"`
	PromotableUrls                 *[]string                 `json:"promotable_urls,omitempty"`
	RequestChangeReasons           *[]map[string]interface{} `json:"request_change_reasons,omitempty"`
	Status                         *string                   `json:"status,omitempty"`
	Subvertical                    *string                   `json:"subvertical,omitempty"`
	SubverticalV2                  *string                   `json:"subvertical_v2,omitempty"`
	TimeCreated                    *core.Time                `json:"time_created,omitempty"`
	Vertical                       *string                   `json:"vertical,omitempty"`
	VerticalV2                     *string                   `json:"vertical_v2,omitempty"`
}

var AdAccountCreationRequestFields = struct {
	AdAccountsCurrency             string
	AdAccountsInfo                 string
	AdditionalComment              string
	AddressInChinese               string
	AddressInEnglish               string
	AddressInLocalLanguage         string
	AdvertiserBusiness             string
	AppealReason                   string
	Business                       string
	BusinessRegistrationID         string
	ChineseLegalEntityName         string
	Contact                        string
	Creator                        string
	CreditCardID                   string
	DisapprovalReasons             string
	EnglishLegalEntityName         string
	ExtendedCreditID               string
	ID                             string
	IsSmb                          string
	IsTest                         string
	LegalEntityNameInLocalLanguage string
	OeRequestID                    string
	OfficialWebsiteURL             string
	PlanningAgencyBusiness         string
	PlanningAgencyBusinessID       string
	PromotableAppIds               string
	PromotablePageIds              string
	PromotableUrls                 string
	RequestChangeReasons           string
	Status                         string
	Subvertical                    string
	SubverticalV2                  string
	TimeCreated                    string
	Vertical                       string
	VerticalV2                     string
}{
	AdAccountsCurrency:             "ad_accounts_currency",
	AdAccountsInfo:                 "ad_accounts_info",
	AdditionalComment:              "additional_comment",
	AddressInChinese:               "address_in_chinese",
	AddressInEnglish:               "address_in_english",
	AddressInLocalLanguage:         "address_in_local_language",
	AdvertiserBusiness:             "advertiser_business",
	AppealReason:                   "appeal_reason",
	Business:                       "business",
	BusinessRegistrationID:         "business_registration_id",
	ChineseLegalEntityName:         "chinese_legal_entity_name",
	Contact:                        "contact",
	Creator:                        "creator",
	CreditCardID:                   "credit_card_id",
	DisapprovalReasons:             "disapproval_reasons",
	EnglishLegalEntityName:         "english_legal_entity_name",
	ExtendedCreditID:               "extended_credit_id",
	ID:                             "id",
	IsSmb:                          "is_smb",
	IsTest:                         "is_test",
	LegalEntityNameInLocalLanguage: "legal_entity_name_in_local_language",
	OeRequestID:                    "oe_request_id",
	OfficialWebsiteURL:             "official_website_url",
	PlanningAgencyBusiness:         "planning_agency_business",
	PlanningAgencyBusinessID:       "planning_agency_business_id",
	PromotableAppIds:               "promotable_app_ids",
	PromotablePageIds:              "promotable_page_ids",
	PromotableUrls:                 "promotable_urls",
	RequestChangeReasons:           "request_change_reasons",
	Status:                         "status",
	Subvertical:                    "subvertical",
	SubverticalV2:                  "subvertical_v2",
	TimeCreated:                    "time_created",
	Vertical:                       "vertical",
	VerticalV2:                     "vertical_v2",
}
