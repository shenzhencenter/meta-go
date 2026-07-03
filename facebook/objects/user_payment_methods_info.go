package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserPaymentMethodsInfo struct {
	AccountID                      *core.ID                  `json:"account_id,omitempty"`
	AvailableCardTypes             *[]string                 `json:"available_card_types,omitempty"`
	AvailablePaymentMethods        *[]string                 `json:"available_payment_methods,omitempty"`
	AvailablePaymentMethodsDetails *[]map[string]interface{} `json:"available_payment_methods_details,omitempty"`
	Country                        *string                   `json:"country,omitempty"`
	Currency                       *string                   `json:"currency,omitempty"`
	ExistingPaymentMethods         *[]map[string]interface{} `json:"existing_payment_methods,omitempty"`
}
