package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdvAInstanceParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdvAInstanceParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdvAInstanceBatchCall(id string, params GetAdvAInstanceParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdvAInstanceBatchRequest(id string, params GetAdvAInstanceParams, options ...core.BatchOption) *core.BatchRequest[objects.AdvAInstance] {
	return core.NewBatchRequest[objects.AdvAInstance](GetAdvAInstanceBatchCall(id, params, options...))
}

func DecodeGetAdvAInstanceBatchResponse(response *core.BatchResponse) (*objects.AdvAInstance, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdvAInstance
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdvAInstance(ctx context.Context, client *core.Client, id string, params GetAdvAInstanceParams) (*objects.AdvAInstance, error) {
	var out objects.AdvAInstance
	call := GetAdvAInstanceBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
