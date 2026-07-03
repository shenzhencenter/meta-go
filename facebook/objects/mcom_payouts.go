package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
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

var McomPayoutsFields = struct {
	NumberOfOrders            string
	OrderIds                  string
	PayoutAmount              string
	PayoutProviderReferenceID string
	PayoutStatus              string
	PayoutTime                string
	Provider                  string
}{
	NumberOfOrders:            "number_of_orders",
	OrderIds:                  "order_ids",
	PayoutAmount:              "payout_amount",
	PayoutProviderReferenceID: "payout_provider_reference_id",
	PayoutStatus:              "payout_status",
	PayoutTime:                "payout_time",
	Provider:                  "provider",
}
