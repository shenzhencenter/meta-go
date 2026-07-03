package objects

type CurrencyAmount struct {
	Amount             *string `json:"amount,omitempty"`
	AmountInHundredths *string `json:"amount_in_hundredths,omitempty"`
	Currency           *string `json:"currency,omitempty"`
	OffsettedAmount    *string `json:"offsetted_amount,omitempty"`
}
