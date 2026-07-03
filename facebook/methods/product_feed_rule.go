package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type DeleteProductFeedRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteProductFeedRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteProductFeedRuleBatchCall(id string, params DeleteProductFeedRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteProductFeedRuleBatchRequest(id string, params DeleteProductFeedRuleParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteProductFeedRuleBatchCall(id, params, options...))
}

func DecodeDeleteProductFeedRuleBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteProductFeedRuleWithResponse(ctx context.Context, client *core.Client, id string, params DeleteProductFeedRuleParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteProductFeedRuleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteProductFeedRule(ctx context.Context, client *core.Client, id string, params DeleteProductFeedRuleParams) (*map[string]interface{}, error) {
	out, _, err := DeleteProductFeedRuleWithResponse(ctx, client, id, params)
	return out, err
}

type GetProductFeedRuleParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedRuleBatchCall(id string, params GetProductFeedRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductFeedRuleBatchRequest(id string, params GetProductFeedRuleParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedRule] {
	return core.NewBatchRequest[objects.ProductFeedRule](GetProductFeedRuleBatchCall(id, params, options...))
}

func DecodeGetProductFeedRuleBatchResponse(response *core.BatchResponse) (*objects.ProductFeedRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeedRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedRuleWithResponse(ctx context.Context, client *core.Client, id string, params GetProductFeedRuleParams) (*objects.ProductFeedRule, *core.Response, error) {
	var out objects.ProductFeedRule
	call := GetProductFeedRuleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductFeedRule(ctx context.Context, client *core.Client, id string, params GetProductFeedRuleParams) (*objects.ProductFeedRule, error) {
	out, _, err := GetProductFeedRuleWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateProductFeedRuleParams struct {
	Params map[string]interface{} `facebook:"params"`
	Extra  core.Params            `facebook:"-"`
}

func (params UpdateProductFeedRuleParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["params"] = params.Params
	return out
}

func UpdateProductFeedRuleBatchCall(id string, params UpdateProductFeedRuleParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateProductFeedRuleBatchRequest(id string, params UpdateProductFeedRuleParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedRule] {
	return core.NewBatchRequest[objects.ProductFeedRule](UpdateProductFeedRuleBatchCall(id, params, options...))
}

func DecodeUpdateProductFeedRuleBatchResponse(response *core.BatchResponse) (*objects.ProductFeedRule, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeedRule
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateProductFeedRuleWithResponse(ctx context.Context, client *core.Client, id string, params UpdateProductFeedRuleParams) (*objects.ProductFeedRule, *core.Response, error) {
	var out objects.ProductFeedRule
	call := UpdateProductFeedRuleBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateProductFeedRule(ctx context.Context, client *core.Client, id string, params UpdateProductFeedRuleParams) (*objects.ProductFeedRule, error) {
	out, _, err := UpdateProductFeedRuleWithResponse(ctx, client, id, params)
	return out, err
}
