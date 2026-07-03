package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCalibratorExistingRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCalibratorExistingRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCalibratorExistingRuleBatchCall(id string, params GetCalibratorExistingRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCalibratorExistingRuleBatchRequest(id string, params GetCalibratorExistingRuleParams, options ...core.BatchOption) *core.BatchRequest[objects.CalibratorExistingRule] {
	return core.NewBatchRequest[objects.CalibratorExistingRule](GetCalibratorExistingRuleBatchCall(id, params, options...))
}

func DecodeGetCalibratorExistingRuleBatchResponse(response *core.BatchResponse) (*objects.CalibratorExistingRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.CalibratorExistingRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCalibratorExistingRule(ctx context.Context, client *core.Client, id string, params GetCalibratorExistingRuleParams) (*objects.CalibratorExistingRule, error) {
	var out objects.CalibratorExistingRule
	call := GetCalibratorExistingRuleBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
