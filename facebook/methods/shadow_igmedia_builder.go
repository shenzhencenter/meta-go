package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetShadowIGMediaBuilderParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetShadowIGMediaBuilderParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetShadowIGMediaBuilderBatchCall(id string, params GetShadowIGMediaBuilderParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetShadowIGMediaBuilderBatchRequest(id string, params GetShadowIGMediaBuilderParams, options ...core.BatchOption) *core.BatchRequest[objects.ShadowIGMediaBuilder] {
	return core.NewBatchRequest[objects.ShadowIGMediaBuilder](GetShadowIGMediaBuilderBatchCall(id, params, options...))
}

func DecodeGetShadowIGMediaBuilderBatchResponse(response *core.BatchResponse) (*objects.ShadowIGMediaBuilder, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ShadowIGMediaBuilder
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetShadowIGMediaBuilder(ctx context.Context, client *core.Client, id string, params GetShadowIGMediaBuilderParams) (*objects.ShadowIGMediaBuilder, error) {
	var out objects.ShadowIGMediaBuilder
	call := GetShadowIGMediaBuilderBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
