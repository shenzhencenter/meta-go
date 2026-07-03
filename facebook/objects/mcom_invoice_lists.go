package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type McomInvoiceLists struct {
	InvoiceDetails *[]McomInvoiceDetails `json:"invoice_details,omitempty"`
	InvoiceIds     *[]core.ID            `json:"invoice_ids,omitempty"`
	PageID         *core.ID              `json:"page_id,omitempty"`
}
