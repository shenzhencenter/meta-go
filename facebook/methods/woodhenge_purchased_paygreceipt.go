package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetWoodhengePurchasedPAYGReceiptParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetWoodhengePurchasedPAYGReceiptParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetWoodhengePurchasedPAYGReceiptBatchCall(id string, params GetWoodhengePurchasedPAYGReceiptParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetWoodhengePurchasedPAYGReceiptBatchRequest(id string, params GetWoodhengePurchasedPAYGReceiptParams, options ...core.BatchOption) *core.BatchRequest[objects.WoodhengePurchasedPAYGReceipt] {
	return core.NewBatchRequest[objects.WoodhengePurchasedPAYGReceipt](GetWoodhengePurchasedPAYGReceiptBatchCall(id, params, options...))
}

func DecodeGetWoodhengePurchasedPAYGReceiptBatchResponse(response *core.BatchResponse) (*objects.WoodhengePurchasedPAYGReceipt, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.WoodhengePurchasedPAYGReceipt
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetWoodhengePurchasedPAYGReceipt(ctx context.Context, client *core.Client, id string, params GetWoodhengePurchasedPAYGReceiptParams) (*objects.WoodhengePurchasedPAYGReceipt, error) {
	var out objects.WoodhengePurchasedPAYGReceipt
	call := GetWoodhengePurchasedPAYGReceiptBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
