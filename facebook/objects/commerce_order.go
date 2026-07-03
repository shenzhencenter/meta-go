package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type CommerceOrder struct {
	BuyerDetails            *map[string]interface{} `json:"buyer_details,omitempty"`
	Channel                 *string                 `json:"channel,omitempty"`
	ContainsBopisItems      *bool                   `json:"contains_bopis_items,omitempty"`
	Created                 *string                 `json:"created,omitempty"`
	EstimatedPaymentDetails *map[string]interface{} `json:"estimated_payment_details,omitempty"`
	ID                      *core.ID                `json:"id,omitempty"`
	IsGroupBuy              *bool                   `json:"is_group_buy,omitempty"`
	IsTestOrder             *bool                   `json:"is_test_order,omitempty"`
	LastUpdated             *string                 `json:"last_updated,omitempty"`
	MerchantOrderID         *core.ID                `json:"merchant_order_id,omitempty"`
	OrderStatus             *map[string]interface{} `json:"order_status,omitempty"`
	SelectedShippingOption  *map[string]interface{} `json:"selected_shipping_option,omitempty"`
	ShipByDate              *string                 `json:"ship_by_date,omitempty"`
	ShippingAddress         *map[string]interface{} `json:"shipping_address,omitempty"`
}
