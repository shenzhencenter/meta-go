package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type EventExternalTicketInfo struct {
	ID            *core.ID        `json:"id,omitempty"`
	MaxSalesPrice *CurrencyAmount `json:"max_sales_price,omitempty"`
	MinSalesPrice *CurrencyAmount `json:"min_sales_price,omitempty"`
	SalesStatus   *string         `json:"sales_status,omitempty"`
}

var EventExternalTicketInfoFields = struct {
	ID            string
	MaxSalesPrice string
	MinSalesPrice string
	SalesStatus   string
}{
	ID:            "id",
	MaxSalesPrice: "max_sales_price",
	MinSalesPrice: "min_sales_price",
	SalesStatus:   "sales_status",
}
