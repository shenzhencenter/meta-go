package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func GetCommerceOrderCancellationsBatchCall(id string, params GetCommerceOrderCancellationsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "cancellations"), params.ToParams(), options...)
}

func NewGetCommerceOrderCancellationsBatchRequest(id string, params GetCommerceOrderCancellationsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrderCancellation]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrderCancellation]](GetCommerceOrderCancellationsBatchCall(id, params, options...))
}

func DecodeGetCommerceOrderCancellationsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrderCancellation], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrderCancellation]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceOrderCancellations(ctx context.Context, client *core.Client, id string, params GetCommerceOrderCancellationsParams) (*core.Cursor[objects.CommerceOrderCancellation], error) {
	var out core.Cursor[objects.CommerceOrderCancellation]
	call := GetCommerceOrderCancellationsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceOrderItemsBatchCall(id string, params GetCommerceOrderItemsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "items"), params.ToParams(), options...)
}

func NewGetCommerceOrderItemsBatchRequest(id string, params GetCommerceOrderItemsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrderItem]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrderItem]](GetCommerceOrderItemsBatchCall(id, params, options...))
}

func DecodeGetCommerceOrderItemsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrderItem], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrderItem]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceOrderItems(ctx context.Context, client *core.Client, id string, params GetCommerceOrderItemsParams) (*core.Cursor[objects.CommerceOrderItem], error) {
	var out core.Cursor[objects.CommerceOrderItem]
	call := GetCommerceOrderItemsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceOrderPaymentsBatchCall(id string, params GetCommerceOrderPaymentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "payments"), params.ToParams(), options...)
}

func NewGetCommerceOrderPaymentsBatchRequest(id string, params GetCommerceOrderPaymentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrderPayment]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrderPayment]](GetCommerceOrderPaymentsBatchCall(id, params, options...))
}

func DecodeGetCommerceOrderPaymentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrderPayment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrderPayment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceOrderPayments(ctx context.Context, client *core.Client, id string, params GetCommerceOrderPaymentsParams) (*core.Cursor[objects.CommerceOrderPayment], error) {
	var out core.Cursor[objects.CommerceOrderPayment]
	call := GetCommerceOrderPaymentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceOrderRefundsBatchCall(id string, params GetCommerceOrderRefundsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "refunds"), params.ToParams(), options...)
}

func NewGetCommerceOrderRefundsBatchRequest(id string, params GetCommerceOrderRefundsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrderRefund]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrderRefund]](GetCommerceOrderRefundsBatchCall(id, params, options...))
}

func DecodeGetCommerceOrderRefundsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrderRefund], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrderRefund]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceOrderRefunds(ctx context.Context, client *core.Client, id string, params GetCommerceOrderRefundsParams) (*core.Cursor[objects.CommerceOrderRefund], error) {
	var out core.Cursor[objects.CommerceOrderRefund]
	call := GetCommerceOrderRefundsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceOrderShipmentsBatchCall(id string, params GetCommerceOrderShipmentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "shipments"), params.ToParams(), options...)
}

func NewGetCommerceOrderShipmentsBatchRequest(id string, params GetCommerceOrderShipmentsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.CommerceOrderShipment]] {
	return core.NewBatchRequest[core.Cursor[objects.CommerceOrderShipment]](GetCommerceOrderShipmentsBatchCall(id, params, options...))
}

func DecodeGetCommerceOrderShipmentsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.CommerceOrderShipment], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.CommerceOrderShipment]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceOrderShipments(ctx context.Context, client *core.Client, id string, params GetCommerceOrderShipmentsParams) (*core.Cursor[objects.CommerceOrderShipment], error) {
	var out core.Cursor[objects.CommerceOrderShipment]
	call := GetCommerceOrderShipmentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func CreateCommerceOrderShipmentsBatchCall(id string, params CreateCommerceOrderShipmentsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "shipments"), params.ToParams(), options...)
}

func NewCreateCommerceOrderShipmentsBatchRequest(id string, params CreateCommerceOrderShipmentsParams, options ...core.BatchOption) *core.BatchRequest[objects.CommerceOrder] {
	return core.NewBatchRequest[objects.CommerceOrder](CreateCommerceOrderShipmentsBatchCall(id, params, options...))
}

func DecodeCreateCommerceOrderShipmentsBatchResponse(response *core.BatchResponse) (*objects.CommerceOrder, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CommerceOrder
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCommerceOrderShipments(ctx context.Context, client *core.Client, id string, params CreateCommerceOrderShipmentsParams) (*objects.CommerceOrder, error) {
	var out objects.CommerceOrder
	call := CreateCommerceOrderShipmentsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetCommerceOrderBatchCall(id string, params GetCommerceOrderParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCommerceOrderBatchRequest(id string, params GetCommerceOrderParams, options ...core.BatchOption) *core.BatchRequest[objects.CommerceOrder] {
	return core.NewBatchRequest[objects.CommerceOrder](GetCommerceOrderBatchCall(id, params, options...))
}

func DecodeGetCommerceOrderBatchResponse(response *core.BatchResponse) (*objects.CommerceOrder, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CommerceOrder
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCommerceOrder(ctx context.Context, client *core.Client, id string, params GetCommerceOrderParams) (*objects.CommerceOrder, error) {
	var out objects.CommerceOrder
	call := GetCommerceOrderBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
