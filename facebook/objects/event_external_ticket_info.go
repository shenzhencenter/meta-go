package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type EventExternalTicketInfo struct {
	ID            *core.ID        `json:"id,omitempty"`
	MaxSalesPrice *CurrencyAmount `json:"max_sales_price,omitempty"`
	MinSalesPrice *CurrencyAmount `json:"min_sales_price,omitempty"`
	SalesStatus   *string         `json:"sales_status,omitempty"`
}
