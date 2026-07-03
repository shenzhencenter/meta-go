package objects

type AdAccountPrepayDetails struct {
	DefaultFundingAmount         *CurrencyAmount `json:"default_funding_amount,omitempty"`
	MaxAcceptableAmount          *CurrencyAmount `json:"max_acceptable_amount,omitempty"`
	MinAcceptableAmount          *CurrencyAmount `json:"min_acceptable_amount,omitempty"`
	ShouldCollectBusinessDetails *bool           `json:"should_collect_business_details,omitempty"`
}

var AdAccountPrepayDetailsFields = struct {
	DefaultFundingAmount         string
	MaxAcceptableAmount          string
	MinAcceptableAmount          string
	ShouldCollectBusinessDetails string
}{
	DefaultFundingAmount:         "default_funding_amount",
	MaxAcceptableAmount:          "max_acceptable_amount",
	MinAcceptableAmount:          "min_acceptable_amount",
	ShouldCollectBusinessDetails: "should_collect_business_details",
}
