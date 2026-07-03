package objects

type AdAccountSubsidyAmountDetails struct {
	EnteredAmount *CurrencyAmount `json:"entered_amount,omitempty"`
	FeeAmount     *CurrencyAmount `json:"fee_amount,omitempty"`
	TotalAmount   *CurrencyAmount `json:"total_amount,omitempty"`
}
