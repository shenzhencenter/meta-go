package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdColumnSizesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdColumnSizesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdColumnSizesBatchCall(id string, params GetAdColumnSizesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdColumnSizesBatchRequest(id string, params GetAdColumnSizesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdColumnSizes] {
	return core.NewBatchRequest[objects.AdColumnSizes](GetAdColumnSizesBatchCall(id, params, options...))
}

func DecodeGetAdColumnSizesBatchResponse(response *core.BatchResponse) (*objects.AdColumnSizes, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdColumnSizes
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdColumnSizesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdColumnSizesParams) (*objects.AdColumnSizes, *core.Response, error) {
	var out objects.AdColumnSizes
	call := GetAdColumnSizesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdColumnSizes(ctx context.Context, client *core.Client, id string, params GetAdColumnSizesParams) (*objects.AdColumnSizes, error) {
	out, _, err := GetAdColumnSizesWithResponse(ctx, client, id, params)
	return out, err
}
