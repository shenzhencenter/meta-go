package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type EventTicketSetting struct {
	ID                 *core.ID `json:"id,omitempty"`
	TicketDeliveryType *string  `json:"ticket_delivery_type,omitempty"`
}
