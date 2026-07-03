package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPageLeadsAccessConfigParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageLeadsAccessConfigParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageLeadsAccessConfigBatchCall(id string, params GetPageLeadsAccessConfigParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPageLeadsAccessConfigBatchRequest(id string, params GetPageLeadsAccessConfigParams, options ...core.BatchOption) *core.BatchRequest[objects.PageLeadsAccessConfig] {
	return core.NewBatchRequest[objects.PageLeadsAccessConfig](GetPageLeadsAccessConfigBatchCall(id, params, options...))
}

func DecodeGetPageLeadsAccessConfigBatchResponse(response *core.BatchResponse) (*objects.PageLeadsAccessConfig, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PageLeadsAccessConfig
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageLeadsAccessConfig(ctx context.Context, client *core.Client, id string, params GetPageLeadsAccessConfigParams) (*objects.PageLeadsAccessConfig, error) {
	var out objects.PageLeadsAccessConfig
	call := GetPageLeadsAccessConfigBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
