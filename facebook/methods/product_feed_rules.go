package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductFeedRulesRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedRulesRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedRulesRulesBatchCall(id string, params GetProductFeedRulesRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "rules"), params.ToParams(), options...)
}

func NewGetProductFeedRulesRulesBatchRequest(id string, params GetProductFeedRulesRulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductFeedRulesGet]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductFeedRulesGet]](GetProductFeedRulesRulesBatchCall(id, params, options...))
}

func DecodeGetProductFeedRulesRulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductFeedRulesGet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductFeedRulesGet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedRulesRulesWithResponse(ctx context.Context, client *core.Client, id string, params GetProductFeedRulesRulesParams) (*core.Cursor[objects.ProductFeedRulesGet], *core.Response, error) {
	var out core.Cursor[objects.ProductFeedRulesGet]
	call := GetProductFeedRulesRulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductFeedRulesRules(ctx context.Context, client *core.Client, id string, params GetProductFeedRulesRulesParams) (*core.Cursor[objects.ProductFeedRulesGet], error) {
	out, _, err := GetProductFeedRulesRulesWithResponse(ctx, client, id, params)
	return out, err
}
