package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdContract struct {
	AccountID             *core.ID `json:"account_id,omitempty"`
	AccountMgrFbid        *string  `json:"account_mgr_fbid,omitempty"`
	AccountMgrName        *string  `json:"account_mgr_name,omitempty"`
	AdopsPersonName       *string  `json:"adops_person_name,omitempty"`
	AdvertiserAddressFbid *string  `json:"advertiser_address_fbid,omitempty"`
	AdvertiserFbid        *string  `json:"advertiser_fbid,omitempty"`
	AdvertiserName        *string  `json:"advertiser_name,omitempty"`
	AgencyDiscount        *float64 `json:"agency_discount,omitempty"`
	AgencyName            *string  `json:"agency_name,omitempty"`
	BillToAddressFbid     *string  `json:"bill_to_address_fbid,omitempty"`
	BillToFbid            *string  `json:"bill_to_fbid,omitempty"`
	CampaignName          *string  `json:"campaign_name,omitempty"`
	CreatedBy             *string  `json:"created_by,omitempty"`
	CreatedDate           *uint64  `json:"created_date,omitempty"`
	CustomerIo            *string  `json:"customer_io,omitempty"`
	IoNumber              *uint64  `json:"io_number,omitempty"`
	IoTerms               *string  `json:"io_terms,omitempty"`
	IoType                *string  `json:"io_type,omitempty"`
	LastUpdatedBy         *string  `json:"last_updated_by,omitempty"`
	LastUpdatedDate       *uint64  `json:"last_updated_date,omitempty"`
	MaxEndDate            *uint64  `json:"max_end_date,omitempty"`
	MdcFbid               *string  `json:"mdc_fbid,omitempty"`
	MediaPlanNumber       *string  `json:"media_plan_number,omitempty"`
	MinStartDate          *uint64  `json:"min_start_date,omitempty"`
	MsaContract           *string  `json:"msa_contract,omitempty"`
	PaymentTerms          *string  `json:"payment_terms,omitempty"`
	RevHoldFlag           *bool    `json:"rev_hold_flag,omitempty"`
	RevHoldReleasedBy     *int     `json:"rev_hold_released_by,omitempty"`
	RevHoldReleasedOn     *uint64  `json:"rev_hold_released_on,omitempty"`
	SalesrepFbid          *string  `json:"salesrep_fbid,omitempty"`
	SalesrepName          *string  `json:"salesrep_name,omitempty"`
	SoldToAddressFbid     *string  `json:"sold_to_address_fbid,omitempty"`
	SoldToFbid            *string  `json:"sold_to_fbid,omitempty"`
	Status                *string  `json:"status,omitempty"`
	Subvertical           *string  `json:"subvertical,omitempty"`
	ThirdpartyBilled      *uint64  `json:"thirdparty_billed,omitempty"`
	ThirdpartyUID         *core.ID `json:"thirdparty_uid,omitempty"`
	ThirdpartyURL         *string  `json:"thirdparty_url,omitempty"`
	VatCountry            *string  `json:"vat_country,omitempty"`
	Version               *uint64  `json:"version,omitempty"`
	Vertical              *string  `json:"vertical,omitempty"`
}

var AdContractFields = struct {
	AccountID             string
	AccountMgrFbid        string
	AccountMgrName        string
	AdopsPersonName       string
	AdvertiserAddressFbid string
	AdvertiserFbid        string
	AdvertiserName        string
	AgencyDiscount        string
	AgencyName            string
	BillToAddressFbid     string
	BillToFbid            string
	CampaignName          string
	CreatedBy             string
	CreatedDate           string
	CustomerIo            string
	IoNumber              string
	IoTerms               string
	IoType                string
	LastUpdatedBy         string
	LastUpdatedDate       string
	MaxEndDate            string
	MdcFbid               string
	MediaPlanNumber       string
	MinStartDate          string
	MsaContract           string
	PaymentTerms          string
	RevHoldFlag           string
	RevHoldReleasedBy     string
	RevHoldReleasedOn     string
	SalesrepFbid          string
	SalesrepName          string
	SoldToAddressFbid     string
	SoldToFbid            string
	Status                string
	Subvertical           string
	ThirdpartyBilled      string
	ThirdpartyUID         string
	ThirdpartyURL         string
	VatCountry            string
	Version               string
	Vertical              string
}{
	AccountID:             "account_id",
	AccountMgrFbid:        "account_mgr_fbid",
	AccountMgrName:        "account_mgr_name",
	AdopsPersonName:       "adops_person_name",
	AdvertiserAddressFbid: "advertiser_address_fbid",
	AdvertiserFbid:        "advertiser_fbid",
	AdvertiserName:        "advertiser_name",
	AgencyDiscount:        "agency_discount",
	AgencyName:            "agency_name",
	BillToAddressFbid:     "bill_to_address_fbid",
	BillToFbid:            "bill_to_fbid",
	CampaignName:          "campaign_name",
	CreatedBy:             "created_by",
	CreatedDate:           "created_date",
	CustomerIo:            "customer_io",
	IoNumber:              "io_number",
	IoTerms:               "io_terms",
	IoType:                "io_type",
	LastUpdatedBy:         "last_updated_by",
	LastUpdatedDate:       "last_updated_date",
	MaxEndDate:            "max_end_date",
	MdcFbid:               "mdc_fbid",
	MediaPlanNumber:       "media_plan_number",
	MinStartDate:          "min_start_date",
	MsaContract:           "msa_contract",
	PaymentTerms:          "payment_terms",
	RevHoldFlag:           "rev_hold_flag",
	RevHoldReleasedBy:     "rev_hold_released_by",
	RevHoldReleasedOn:     "rev_hold_released_on",
	SalesrepFbid:          "salesrep_fbid",
	SalesrepName:          "salesrep_name",
	SoldToAddressFbid:     "sold_to_address_fbid",
	SoldToFbid:            "sold_to_fbid",
	Status:                "status",
	Subvertical:           "subvertical",
	ThirdpartyBilled:      "thirdparty_billed",
	ThirdpartyUID:         "thirdparty_uid",
	ThirdpartyURL:         "thirdparty_url",
	VatCountry:            "vat_country",
	Version:               "version",
	Vertical:              "vertical",
}
