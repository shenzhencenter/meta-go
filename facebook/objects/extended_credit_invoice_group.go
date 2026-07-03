package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
