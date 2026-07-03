package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var CRMAddressFields = struct {
	City               string
	CnpjTaxID          string
	Country            string
	ID                 string
	PostalCode         string
	RegistrationLabel  string
	RegistrationNumber string
	State              string
	Street1            string
	Street2            string
	Street3            string
	Street4            string
	ValidationStatus   string
	VatTaxID           string
}{
	City:               "city",
	CnpjTaxID:          "cnpj_tax_id",
	Country:            "country",
	ID:                 "id",
	PostalCode:         "postal_code",
	RegistrationLabel:  "registration_label",
	RegistrationNumber: "registration_number",
	State:              "state",
	Street1:            "street1",
	Street2:            "street2",
	Street3:            "street3",
	Street4:            "street4",
	ValidationStatus:   "validation_status",
	VatTaxID:           "vat_tax_id",
}
