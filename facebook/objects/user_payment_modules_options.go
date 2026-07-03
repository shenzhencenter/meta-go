package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type UserPaymentModulesOptions struct {
	AccountID               *core.ID                  `json:"account_id,omitempty"`
	AvailablePaymentOptions *[]map[string]interface{} `json:"available_payment_options,omitempty"`
	Country                 *string                   `json:"country,omitempty"`
	Currency                *string                   `json:"currency,omitempty"`
}
