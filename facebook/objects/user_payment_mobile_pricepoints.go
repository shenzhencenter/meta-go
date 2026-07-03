package objects

type UserPaymentMobilePricepoints struct {
	MobileCountry    *string                   `json:"mobile_country,omitempty"`
	PhoneNumberLast4 *string                   `json:"phone_number_last4,omitempty"`
	Pricepoints      *[]map[string]interface{} `json:"pricepoints,omitempty"`
	UserCurrency     *string                   `json:"user_currency,omitempty"`
}
