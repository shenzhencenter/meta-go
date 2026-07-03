package objects

type PagePaymentOptions struct {
	Amex       *uint64 `json:"amex,omitempty"`
	CashOnly   *uint64 `json:"cash_only,omitempty"`
	Discover   *uint64 `json:"discover,omitempty"`
	Mastercard *uint64 `json:"mastercard,omitempty"`
	Visa       *uint64 `json:"visa,omitempty"`
}

var PagePaymentOptionsFields = struct {
	Amex       string
	CashOnly   string
	Discover   string
	Mastercard string
	Visa       string
}{
	Amex:       "amex",
	CashOnly:   "cash_only",
	Discover:   "discover",
	Mastercard: "mastercard",
	Visa:       "visa",
}
