package objects

type AdAccountPaymentOptions struct {
	AvailableAltpayOptions  *[]map[string]interface{} `json:"available_altpay_options,omitempty"`
	AvailableCardTypes      *[]string                 `json:"available_card_types,omitempty"`
	AvailablePaymentOptions *[]string                 `json:"available_payment_options,omitempty"`
	ExistingPaymentMethods  *[]map[string]interface{} `json:"existing_payment_methods,omitempty"`
}
