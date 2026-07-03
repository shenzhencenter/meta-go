package objects

type AdAccountPrepayDetails struct {
	DefaultFundingAmount         *CurrencyAmount `json:"default_funding_amount,omitempty"`
	MaxAcceptableAmount          *CurrencyAmount `json:"max_acceptable_amount,omitempty"`
	MinAcceptableAmount          *CurrencyAmount `json:"min_acceptable_amount,omitempty"`
	ShouldCollectBusinessDetails *bool           `json:"should_collect_business_details,omitempty"`
}
