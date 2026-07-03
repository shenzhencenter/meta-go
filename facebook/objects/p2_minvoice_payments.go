package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type P2MInvoicePayments struct {
	PageID   *core.ID                  `json:"page_id,omitempty"`
	Payments *[]map[string]interface{} `json:"payments,omitempty"`
}

var P2MInvoicePaymentsFields = struct {
	PageID   string
	Payments string
}{
	PageID:   "page_id",
	Payments: "payments",
}
