package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetCommerceOrderCancellationsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceOrderCancellationsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceOrderCancellations(ctx context.Context, client *core.Client, id string, params GetCommerceOrderCancellationsParams) (*core.Cursor[objects.CommerceOrderCancellation], error) {
	var out core.Cursor[objects.CommerceOrderCancellation]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "cancellations"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceOrderItemsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceOrderItemsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceOrderItems(ctx context.Context, client *core.Client, id string, params GetCommerceOrderItemsParams) (*core.Cursor[objects.CommerceOrderItem], error) {
	var out core.Cursor[objects.CommerceOrderItem]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "items"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceOrderPaymentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceOrderPaymentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceOrderPayments(ctx context.Context, client *core.Client, id string, params GetCommerceOrderPaymentsParams) (*core.Cursor[objects.CommerceOrderPayment], error) {
	var out core.Cursor[objects.CommerceOrderPayment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "payments"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceOrderRefundsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceOrderRefundsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceOrderRefunds(ctx context.Context, client *core.Client, id string, params GetCommerceOrderRefundsParams) (*core.Cursor[objects.CommerceOrderRefund], error) {
	var out core.Cursor[objects.CommerceOrderRefund]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "refunds"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceOrderShipmentsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceOrderShipmentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceOrderShipments(ctx context.Context, client *core.Client, id string, params GetCommerceOrderShipmentsParams) (*core.Cursor[objects.CommerceOrderShipment], error) {
	var out core.Cursor[objects.CommerceOrderShipment]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "shipments"), params.ToParams(), &out)
	return &out, err
}

type CreateCommerceOrderShipmentsParams struct {
	ExternalRedemptionLink              *string                   `facebook:"external_redemption_link"`
	ExternalShipmentID                  *core.ID                  `facebook:"external_shipment_id"`
	Fulfillment                         *map[string]interface{}   `facebook:"fulfillment"`
	IdempotencyKey                      string                    `facebook:"idempotency_key"`
	Items                               *[]map[string]interface{} `facebook:"items"`
	MerchantOrderReference              *string                   `facebook:"merchant_order_reference"`
	ShipmentOriginPostalCode            *string                   `facebook:"shipment_origin_postal_code"`
	ShippingTaxDetails                  *map[string]interface{}   `facebook:"shipping_tax_details"`
	ShouldUseDefaultFulfillmentLocation *bool                     `facebook:"should_use_default_fulfillment_location"`
	TrackingInfo                        *map[string]interface{}   `facebook:"tracking_info"`
	Extra                               core.Params               `facebook:"-"`
}

func (params CreateCommerceOrderShipmentsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ExternalRedemptionLink != nil {
		out["external_redemption_link"] = *params.ExternalRedemptionLink
	}
	if params.ExternalShipmentID != nil {
		out["external_shipment_id"] = *params.ExternalShipmentID
	}
	if params.Fulfillment != nil {
		out["fulfillment"] = *params.Fulfillment
	}
	out["idempotency_key"] = params.IdempotencyKey
	if params.Items != nil {
		out["items"] = *params.Items
	}
	if params.MerchantOrderReference != nil {
		out["merchant_order_reference"] = *params.MerchantOrderReference
	}
	if params.ShipmentOriginPostalCode != nil {
		out["shipment_origin_postal_code"] = *params.ShipmentOriginPostalCode
	}
	if params.ShippingTaxDetails != nil {
		out["shipping_tax_details"] = *params.ShippingTaxDetails
	}
	if params.ShouldUseDefaultFulfillmentLocation != nil {
		out["should_use_default_fulfillment_location"] = *params.ShouldUseDefaultFulfillmentLocation
	}
	if params.TrackingInfo != nil {
		out["tracking_info"] = *params.TrackingInfo
	}
	return out
}

func CreateCommerceOrderShipments(ctx context.Context, client *core.Client, id string, params CreateCommerceOrderShipmentsParams) (*objects.CommerceOrder, error) {
	var out objects.CommerceOrder
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "shipments"), params.ToParams(), &out)
	return &out, err
}

type GetCommerceOrderParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCommerceOrderParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCommerceOrder(ctx context.Context, client *core.Client, id string, params GetCommerceOrderParams) (*objects.CommerceOrder, error) {
	var out objects.CommerceOrder
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
