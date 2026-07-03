package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CRMAddress struct {
	City               *string  `json:"city,omitempty"`
	CnpjTaxID          *core.ID `json:"cnpj_tax_id,omitempty"`
	Country            *string  `json:"country,omitempty"`
	ID                 *core.ID `json:"id,omitempty"`
	PostalCode         *string  `json:"postal_code,omitempty"`
	RegistrationLabel  *string  `json:"registration_label,omitempty"`
	RegistrationNumber *string  `json:"registration_number,omitempty"`
	State              *string  `json:"state,omitempty"`
	Street1            *string  `json:"street1,omitempty"`
	Street2            *string  `json:"street2,omitempty"`
	Street3            *string  `json:"street3,omitempty"`
	Street4            *string  `json:"street4,omitempty"`
	ValidationStatus   *string  `json:"validation_status,omitempty"`
	VatTaxID           *core.ID `json:"vat_tax_id,omitempty"`
}
