package objects

type Currency struct {
	CurrencyOffset     *uint64  `json:"currency_offset,omitempty"`
	UsdExchange        *float64 `json:"usd_exchange,omitempty"`
	UsdExchangeInverse *float64 `json:"usd_exchange_inverse,omitempty"`
	UserCurrency       *string  `json:"user_currency,omitempty"`
}
