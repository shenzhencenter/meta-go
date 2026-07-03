package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductDeliveryPreferenceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductDeliveryPreferenceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductDeliveryPreferenceBatchCall(id string, params GetProductDeliveryPreferenceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductDeliveryPreferenceBatchRequest(id string, params GetProductDeliveryPreferenceParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductDeliveryPreference] {
	return core.NewBatchRequest[objects.ProductDeliveryPreference](GetProductDeliveryPreferenceBatchCall(id, params, options...))
}

func DecodeGetProductDeliveryPreferenceBatchResponse(response *core.BatchResponse) (*objects.ProductDeliveryPreference, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductDeliveryPreference
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductDeliveryPreference(ctx context.Context, client *core.Client, id string, params GetProductDeliveryPreferenceParams) (*objects.ProductDeliveryPreference, error) {
	var out objects.ProductDeliveryPreference
	call := GetProductDeliveryPreferenceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
