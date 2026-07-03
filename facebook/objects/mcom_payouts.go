package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type McomPayouts struct {
	NumberOfOrders            *int                    `json:"number_of_orders,omitempty"`
	OrderIds                  *[]core.ID              `json:"order_ids,omitempty"`
	PayoutAmount              *map[string]interface{} `json:"payout_amount,omitempty"`
	PayoutProviderReferenceID *core.ID                `json:"payout_provider_reference_id,omitempty"`
	PayoutStatus              *string                 `json:"payout_status,omitempty"`
	PayoutTime                *int                    `json:"payout_time,omitempty"`
	Provider                  *string                 `json:"provider,omitempty"`
}
