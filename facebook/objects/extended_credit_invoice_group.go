package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExtendedCreditInvoiceGroup struct {
	AutoEnroll       *bool                `json:"auto_enroll,omitempty"`
	BillToAddress    *CRMAddress          `json:"bill_to_address,omitempty"`
	CustomerPoNumber *string              `json:"customer_po_number,omitempty"`
	Email            *ExtendedCreditEmail `json:"email,omitempty"`
	Emails           *[]string            `json:"emails,omitempty"`
	ID               *core.ID             `json:"id,omitempty"`
	LiableAddress    *CRMAddress          `json:"liable_address,omitempty"`
	Name             *string              `json:"name,omitempty"`
	SoldToAddress    *CRMAddress          `json:"sold_to_address,omitempty"`
}

var ExtendedCreditInvoiceGroupFields = struct {
	AutoEnroll       string
	BillToAddress    string
	CustomerPoNumber string
	Email            string
	Emails           string
	ID               string
	LiableAddress    string
	Name             string
	SoldToAddress    string
}{
	AutoEnroll:       "auto_enroll",
	BillToAddress:    "bill_to_address",
	CustomerPoNumber: "customer_po_number",
	Email:            "email",
	Emails:           "emails",
	ID:               "id",
	LiableAddress:    "liable_address",
	Name:             "name",
	SoldToAddress:    "sold_to_address",
}
