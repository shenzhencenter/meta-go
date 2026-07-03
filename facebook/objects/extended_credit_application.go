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
