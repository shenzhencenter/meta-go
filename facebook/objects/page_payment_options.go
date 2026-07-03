package objects

type PagePaymentOptions struct {
	Amex       *uint64 `json:"amex,omitempty"`
	CashOnly   *uint64 `json:"cash_only,omitempty"`
	Discover   *uint64 `json:"discover,omitempty"`
	Mastercard *uint64 `json:"mastercard,omitempty"`
	Visa       *uint64 `json:"visa,omitempty"`
}
