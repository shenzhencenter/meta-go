package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdToplineParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdToplineParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdToplineBatchCall(id string, params GetAdToplineParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdToplineBatchRequest(id string, params GetAdToplineParams, options ...core.BatchOption) *core.BatchRequest[objects.AdTopline] {
	return core.NewBatchRequest[objects.AdTopline](GetAdToplineBatchCall(id, params, options...))
}

func DecodeGetAdToplineBatchResponse(response *core.BatchResponse) (*objects.AdTopline, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdTopline
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdToplineWithResponse(ctx context.Context, client *core.Client, id string, params GetAdToplineParams) (*objects.AdTopline, *core.Response, error) {
	var out objects.AdTopline
	call := GetAdToplineBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdTopline(ctx context.Context, client *core.Client, id string, params GetAdToplineParams) (*objects.AdTopline, error) {
	out, _, err := GetAdToplineWithResponse(ctx, client, id, params)
	return out, err
}
