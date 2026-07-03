package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
