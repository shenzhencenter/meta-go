package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExtendedCreditApplication struct {
	BillingCountry      *string         `json:"billing_country,omitempty"`
	City                *string         `json:"city,omitempty"`
	Cnpj                *string         `json:"cnpj,omitempty"`
	Country             *string         `json:"country,omitempty"`
	DisplayCurrency     *string         `json:"display_currency,omitempty"`
	DunsNumber          *string         `json:"duns_number,omitempty"`
	ID                  *core.ID        `json:"id,omitempty"`
	InvoiceEmailAddress *string         `json:"invoice_email_address,omitempty"`
	IsUmi               *bool           `json:"is_umi,omitempty"`
	LegalEntityName     *string         `json:"legal_entity_name,omitempty"`
	OriginalOnlineLimit *CurrencyAmount `json:"original_online_limit,omitempty"`
	PhoneNumber         *string         `json:"phone_number,omitempty"`
	PostalCode          *string         `json:"postal_code,omitempty"`
	ProductTypes        *[]string       `json:"product_types,omitempty"`
	ProposedCreditLimit *CurrencyAmount `json:"proposed_credit_limit,omitempty"`
	RegistrationNumber  *string         `json:"registration_number,omitempty"`
	RunID               *core.ID        `json:"run_id,omitempty"`
	State               *string         `json:"state,omitempty"`
	Status              *string         `json:"status,omitempty"`
	Street1             *string         `json:"street1,omitempty"`
	Street2             *string         `json:"street2,omitempty"`
	Submitter           *User           `json:"submitter,omitempty"`
	TaxExemptStatus     *string         `json:"tax_exempt_status,omitempty"`
	TaxID               *core.ID        `json:"tax_id,omitempty"`
	Terms               *string         `json:"terms,omitempty"`
}

var ExtendedCreditApplicationFields = struct {
	BillingCountry      string
	City                string
	Cnpj                string
	Country             string
	DisplayCurrency     string
	DunsNumber          string
	ID                  string
	InvoiceEmailAddress string
	IsUmi               string
	LegalEntityName     string
	OriginalOnlineLimit string
	PhoneNumber         string
	PostalCode          string
	ProductTypes        string
	ProposedCreditLimit string
	RegistrationNumber  string
	RunID               string
	State               string
	Status              string
	Street1             string
	Street2             string
	Submitter           string
	TaxExemptStatus     string
	TaxID               string
	Terms               string
}{
	BillingCountry:      "billing_country",
	City:                "city",
	Cnpj:                "cnpj",
	Country:             "country",
	DisplayCurrency:     "display_currency",
	DunsNumber:          "duns_number",
	ID:                  "id",
	InvoiceEmailAddress: "invoice_email_address",
	IsUmi:               "is_umi",
	LegalEntityName:     "legal_entity_name",
	OriginalOnlineLimit: "original_online_limit",
	PhoneNumber:         "phone_number",
	PostalCode:          "postal_code",
	ProductTypes:        "product_types",
	ProposedCreditLimit: "proposed_credit_limit",
	RegistrationNumber:  "registration_number",
	RunID:               "run_id",
	State:               "state",
	Status:              "status",
	Street1:             "street1",
	Street2:             "street2",
	Submitter:           "submitter",
	TaxExemptStatus:     "tax_exempt_status",
	TaxID:               "tax_id",
	Terms:               "terms",
}
