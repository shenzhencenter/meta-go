package objects

type AdCampaignGroupAgencyFeeConfigForApi struct {
	AgencyFeePct        *float64 `json:"agency_fee_pct,omitempty"`
	IsAgencyFeeDisabled *bool    `json:"is_agency_fee_disabled,omitempty"`
	IsDefaultAgencyFee  *bool    `json:"is_default_agency_fee,omitempty"`
}

var AdCampaignGroupAgencyFeeConfigForApiFields = struct {
	AgencyFeePct        string
	IsAgencyFeeDisabled string
	IsDefaultAgencyFee  string
}{
	AgencyFeePct:        "agency_fee_pct",
	IsAgencyFeeDisabled: "is_agency_fee_disabled",
	IsDefaultAgencyFee:  "is_default_agency_fee",
}
