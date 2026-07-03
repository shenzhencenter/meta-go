package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsPivotRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsPivotRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsPivotRulesBatchCall(id string, params GetAdsPivotRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsPivotRulesBatchRequest(id string, params GetAdsPivotRulesParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsPivotRules] {
	return core.NewBatchRequest[objects.AdsPivotRules](GetAdsPivotRulesBatchCall(id, params, options...))
}

func DecodeGetAdsPivotRulesBatchResponse(response *core.BatchResponse) (*objects.AdsPivotRules, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsPivotRules
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsPivotRulesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsPivotRulesParams) (*objects.AdsPivotRules, *core.Response, error) {
	var out objects.AdsPivotRules
	call := GetAdsPivotRulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsPivotRules(ctx context.Context, client *core.Client, id string, params GetAdsPivotRulesParams) (*objects.AdsPivotRules, error) {
	out, _, err := GetAdsPivotRulesWithResponse(ctx, client, id, params)
	return out, err
}
