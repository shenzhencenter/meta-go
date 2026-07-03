package objects

type CurrencyAmount struct {
	Amount             *string `json:"amount,omitempty"`
	AmountInHundredths *string `json:"amount_in_hundredths,omitempty"`
	Currency           *string `json:"currency,omitempty"`
	OffsettedAmount    *string `json:"offsetted_amount,omitempty"`
}

var CurrencyAmountFields = struct {
	Amount             string
	AmountInHundredths string
	Currency           string
	OffsettedAmount    string
}{
	Amount:             "amount",
	AmountInHundredths: "amount_in_hundredths",
	Currency:           "currency",
	OffsettedAmount:    "offsetted_amount",
}
