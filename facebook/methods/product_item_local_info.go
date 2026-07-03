package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductItemLocalInfoParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductItemLocalInfoParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductItemLocalInfoBatchCall(id string, params GetProductItemLocalInfoParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductItemLocalInfoBatchRequest(id string, params GetProductItemLocalInfoParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductItemLocalInfo] {
	return core.NewBatchRequest[objects.ProductItemLocalInfo](GetProductItemLocalInfoBatchCall(id, params, options...))
}

func DecodeGetProductItemLocalInfoBatchResponse(response *core.BatchResponse) (*objects.ProductItemLocalInfo, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductItemLocalInfo
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductItemLocalInfoWithResponse(ctx context.Context, client *core.Client, id string, params GetProductItemLocalInfoParams) (*objects.ProductItemLocalInfo, *core.Response, error) {
	var out objects.ProductItemLocalInfo
	call := GetProductItemLocalInfoBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductItemLocalInfo(ctx context.Context, client *core.Client, id string, params GetProductItemLocalInfoParams) (*objects.ProductItemLocalInfo, error) {
	out, _, err := GetProductItemLocalInfoWithResponse(ctx, client, id, params)
	return out, err
}
