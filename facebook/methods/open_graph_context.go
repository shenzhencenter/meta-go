package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetOpenGraphContextParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetOpenGraphContextParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetOpenGraphContextBatchCall(id string, params GetOpenGraphContextParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetOpenGraphContextBatchRequest(id string, params GetOpenGraphContextParams, options ...core.BatchOption) *core.BatchRequest[objects.OpenGraphContext] {
	return core.NewBatchRequest[objects.OpenGraphContext](GetOpenGraphContextBatchCall(id, params, options...))
}

func DecodeGetOpenGraphContextBatchResponse(response *core.BatchResponse) (*objects.OpenGraphContext, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.OpenGraphContext
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetOpenGraphContextWithResponse(ctx context.Context, client *core.Client, id string, params GetOpenGraphContextParams) (*objects.OpenGraphContext, *core.Response, error) {
	var out objects.OpenGraphContext
	call := GetOpenGraphContextBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetOpenGraphContext(ctx context.Context, client *core.Client, id string, params GetOpenGraphContextParams) (*objects.OpenGraphContext, error) {
	out, _, err := GetOpenGraphContextWithResponse(ctx, client, id, params)
	return out, err
}
