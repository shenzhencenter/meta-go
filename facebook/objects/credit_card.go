package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type CreditCard struct {
	BillingAddress                 *map[string]interface{} `json:"billing_address,omitempty"`
	CardCobadging                  *string                 `json:"card_cobadging,omitempty"`
	CardHolderName                 *string                 `json:"card_holder_name,omitempty"`
	CardType                       *string                 `json:"card_type,omitempty"`
	CredentialID                   *core.ID                `json:"credential_id,omitempty"`
	DefaultReceivingMethodProducts *[]string               `json:"default_receiving_method_products,omitempty"`
	ExpiryMonth                    *string                 `json:"expiry_month,omitempty"`
	ExpiryYear                     *string                 `json:"expiry_year,omitempty"`
	ID                             *core.ID                `json:"id,omitempty"`
	IsCvvTrickyBin                 *bool                   `json:"is_cvv_tricky_bin,omitempty"`
	IsEnabled                      *bool                   `json:"is_enabled,omitempty"`
	IsLastUsed                     *bool                   `json:"is_last_used,omitempty"`
	IsNetworkTokenizedInIndia      *bool                   `json:"is_network_tokenized_in_india,omitempty"`
	IsSoftDisabled                 *bool                   `json:"is_soft_disabled,omitempty"`
	IsUserVerified                 *bool                   `json:"is_user_verified,omitempty"`
	IsZipVerified                  *bool                   `json:"is_zip_verified,omitempty"`
	Last4                          *string                 `json:"last4,omitempty"`
	ReadableCardType               *string                 `json:"readable_card_type,omitempty"`
	TimeCreated                    *core.Time              `json:"time_created,omitempty"`
	TimeCreatedTs                  *int                    `json:"time_created_ts,omitempty"`
	Type                           *string                 `json:"type,omitempty"`
}

var CreditCardFields = struct {
	BillingAddress                 string
	CardCobadging                  string
	CardHolderName                 string
	CardType                       string
	CredentialID                   string
	DefaultReceivingMethodProducts string
	ExpiryMonth                    string
	ExpiryYear                     string
	ID                             string
	IsCvvTrickyBin                 string
	IsEnabled                      string
	IsLastUsed                     string
	IsNetworkTokenizedInIndia      string
	IsSoftDisabled                 string
	IsUserVerified                 string
	IsZipVerified                  string
	Last4                          string
	ReadableCardType               string
	TimeCreated                    string
	TimeCreatedTs                  string
	Type                           string
}{
	BillingAddress:                 "billing_address",
	CardCobadging:                  "card_cobadging",
	CardHolderName:                 "card_holder_name",
	CardType:                       "card_type",
	CredentialID:                   "credential_id",
	DefaultReceivingMethodProducts: "default_receiving_method_products",
	ExpiryMonth:                    "expiry_month",
	ExpiryYear:                     "expiry_year",
	ID:                             "id",
	IsCvvTrickyBin:                 "is_cvv_tricky_bin",
	IsEnabled:                      "is_enabled",
	IsLastUsed:                     "is_last_used",
	IsNetworkTokenizedInIndia:      "is_network_tokenized_in_india",
	IsSoftDisabled:                 "is_soft_disabled",
	IsUserVerified:                 "is_user_verified",
	IsZipVerified:                  "is_zip_verified",
	Last4:                          "last4",
	ReadableCardType:               "readable_card_type",
	TimeCreated:                    "time_created",
	TimeCreatedTs:                  "time_created_ts",
	Type:                           "type",
}
