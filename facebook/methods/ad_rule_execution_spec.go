package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdRuleExecutionSpecParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdRuleExecutionSpecParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdRuleExecutionSpecBatchCall(id string, params GetAdRuleExecutionSpecParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdRuleExecutionSpecBatchRequest(id string, params GetAdRuleExecutionSpecParams, options ...core.BatchOption) *core.BatchRequest[objects.AdRuleExecutionSpec] {
	return core.NewBatchRequest[objects.AdRuleExecutionSpec](GetAdRuleExecutionSpecBatchCall(id, params, options...))
}

func DecodeGetAdRuleExecutionSpecBatchResponse(response *core.BatchResponse) (*objects.AdRuleExecutionSpec, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdRuleExecutionSpec
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdRuleExecutionSpec(ctx context.Context, client *core.Client, id string, params GetAdRuleExecutionSpecParams) (*objects.AdRuleExecutionSpec, error) {
	var out objects.AdRuleExecutionSpec
	call := GetAdRuleExecutionSpecBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
