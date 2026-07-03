package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type McomInvoiceLists struct {
	InvoiceDetails *[]McomInvoiceDetails `json:"invoice_details,omitempty"`
	InvoiceIds     *[]core.ID            `json:"invoice_ids,omitempty"`
	PageID         *core.ID              `json:"page_id,omitempty"`
}

var McomInvoiceListsFields = struct {
	InvoiceDetails string
	InvoiceIds     string
	PageID         string
}{
	InvoiceDetails: "invoice_details",
	InvoiceIds:     "invoice_ids",
	PageID:         "page_id",
}
