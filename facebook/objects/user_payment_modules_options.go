package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserPaymentModulesOptions struct {
	AccountID               *core.ID                  `json:"account_id,omitempty"`
	AvailablePaymentOptions *[]map[string]interface{} `json:"available_payment_options,omitempty"`
	Country                 *string                   `json:"country,omitempty"`
	Currency                *string                   `json:"currency,omitempty"`
}

var UserPaymentModulesOptionsFields = struct {
	AccountID               string
	AvailablePaymentOptions string
	Country                 string
	Currency                string
}{
	AccountID:               "account_id",
	AvailablePaymentOptions: "available_payment_options",
	Country:                 "country",
	Currency:                "currency",
}
