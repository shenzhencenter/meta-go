package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type P2MInvoicePayments struct {
	PageID   *core.ID                  `json:"page_id,omitempty"`
	Payments *[]map[string]interface{} `json:"payments,omitempty"`
}
