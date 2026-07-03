package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdRuleEvaluationSpecParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdRuleEvaluationSpecParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdRuleEvaluationSpecBatchCall(id string, params GetAdRuleEvaluationSpecParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdRuleEvaluationSpecBatchRequest(id string, params GetAdRuleEvaluationSpecParams, options ...core.BatchOption) *core.BatchRequest[objects.AdRuleEvaluationSpec] {
	return core.NewBatchRequest[objects.AdRuleEvaluationSpec](GetAdRuleEvaluationSpecBatchCall(id, params, options...))
}

func DecodeGetAdRuleEvaluationSpecBatchResponse(response *core.BatchResponse) (*objects.AdRuleEvaluationSpec, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdRuleEvaluationSpec
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdRuleEvaluationSpec(ctx context.Context, client *core.Client, id string, params GetAdRuleEvaluationSpecParams) (*objects.AdRuleEvaluationSpec, error) {
	var out objects.AdRuleEvaluationSpec
	call := GetAdRuleEvaluationSpecBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
