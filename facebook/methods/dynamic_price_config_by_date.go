package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetDynamicPriceConfigByDateParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetDynamicPriceConfigByDateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetDynamicPriceConfigByDateBatchCall(id string, params GetDynamicPriceConfigByDateParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetDynamicPriceConfigByDateBatchRequest(id string, params GetDynamicPriceConfigByDateParams, options ...core.BatchOption) *core.BatchRequest[objects.DynamicPriceConfigByDate] {
	return core.NewBatchRequest[objects.DynamicPriceConfigByDate](GetDynamicPriceConfigByDateBatchCall(id, params, options...))
}

func DecodeGetDynamicPriceConfigByDateBatchResponse(response *core.BatchResponse) (*objects.DynamicPriceConfigByDate, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.DynamicPriceConfigByDate
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetDynamicPriceConfigByDate(ctx context.Context, client *core.Client, id string, params GetDynamicPriceConfigByDateParams) (*objects.DynamicPriceConfigByDate, error) {
	var out objects.DynamicPriceConfigByDate
	call := GetDynamicPriceConfigByDateBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
