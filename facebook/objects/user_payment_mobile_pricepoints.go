package objects

type UserPaymentMobilePricepoints struct {
	MobileCountry    *string                   `json:"mobile_country,omitempty"`
	PhoneNumberLast4 *string                   `json:"phone_number_last4,omitempty"`
	Pricepoints      *[]map[string]interface{} `json:"pricepoints,omitempty"`
	UserCurrency     *string                   `json:"user_currency,omitempty"`
}

var UserPaymentMobilePricepointsFields = struct {
	MobileCountry    string
	PhoneNumberLast4 string
	Pricepoints      string
	UserCurrency     string
}{
	MobileCountry:    "mobile_country",
	PhoneNumberLast4: "phone_number_last4",
	Pricepoints:      "pricepoints",
	UserCurrency:     "user_currency",
}
