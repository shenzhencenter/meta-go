package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var DirectDebitFields = struct {
	BankAccountLastX4              string
	BankCodeLastX4                 string
	BankName                       string
	DefaultReceivingMethodProducts string
	DisplayString                  string
	ID                             string
	LastFourDigits                 string
	OnboardingURL                  string
	OwnerName                      string
	Status                         string
}{
	BankAccountLastX4:              "bank_account_last_4",
	BankCodeLastX4:                 "bank_code_last_4",
	BankName:                       "bank_name",
	DefaultReceivingMethodProducts: "default_receiving_method_products",
	DisplayString:                  "display_string",
	ID:                             "id",
	LastFourDigits:                 "last_four_digits",
	OnboardingURL:                  "onboarding_url",
	OwnerName:                      "owner_name",
	Status:                         "status",
}
