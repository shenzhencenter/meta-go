package objects

type AdCampaignGroupAgencyFeeConfigForApi struct {
	AgencyFeePct        *float64 `json:"agency_fee_pct,omitempty"`
	IsAgencyFeeDisabled *bool    `json:"is_agency_fee_disabled,omitempty"`
	IsDefaultAgencyFee  *bool    `json:"is_default_agency_fee,omitempty"`
}
