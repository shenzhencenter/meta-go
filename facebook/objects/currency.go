package objects

type Currency struct {
	CurrencyOffset     *uint64  `json:"currency_offset,omitempty"`
	UsdExchange        *float64 `json:"usd_exchange,omitempty"`
	UsdExchangeInverse *float64 `json:"usd_exchange_inverse,omitempty"`
	UserCurrency       *string  `json:"user_currency,omitempty"`
}

var CurrencyFields = struct {
	CurrencyOffset     string
	UsdExchange        string
	UsdExchangeInverse string
	UserCurrency       string
}{
	CurrencyOffset:     "currency_offset",
	UsdExchange:        "usd_exchange",
	UsdExchangeInverse: "usd_exchange_inverse",
	UserCurrency:       "user_currency",
}
