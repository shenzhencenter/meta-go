package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type DirectDebit struct {
	BankAccountLastX4              *string   `json:"bank_account_last_4,omitempty"`
	BankCodeLastX4                 *string   `json:"bank_code_last_4,omitempty"`
	BankName                       *string   `json:"bank_name,omitempty"`
	DefaultReceivingMethodProducts *[]string `json:"default_receiving_method_products,omitempty"`
	DisplayString                  *string   `json:"display_string,omitempty"`
	ID                             *core.ID  `json:"id,omitempty"`
	LastFourDigits                 *string   `json:"last_four_digits,omitempty"`
	OnboardingURL                  *string   `json:"onboarding_url,omitempty"`
	OwnerName                      *string   `json:"owner_name,omitempty"`
	Status                         *int      `json:"status,omitempty"`
}
