package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAREffectParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAREffectParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAREffectBatchCall(id string, params GetAREffectParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAREffectBatchRequest(id string, params GetAREffectParams, options ...core.BatchOption) *core.BatchRequest[objects.AREffect] {
	return core.NewBatchRequest[objects.AREffect](GetAREffectBatchCall(id, params, options...))
}

func DecodeGetAREffectBatchResponse(response *core.BatchResponse) (*objects.AREffect, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AREffect
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAREffectWithResponse(ctx context.Context, client *core.Client, id string, params GetAREffectParams) (*objects.AREffect, *core.Response, error) {
	var out objects.AREffect
	call := GetAREffectBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAREffect(ctx context.Context, client *core.Client, id string, params GetAREffectParams) (*objects.AREffect, error) {
	out, _, err := GetAREffectWithResponse(ctx, client, id, params)
	return out, err
}
