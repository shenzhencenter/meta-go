package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetProductFeedUploadErrorSamplesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadErrorSamplesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadErrorSamplesBatchCall(id string, params GetProductFeedUploadErrorSamplesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "samples"), params.ToParams(), options...)
}

func NewGetProductFeedUploadErrorSamplesBatchRequest(id string, params GetProductFeedUploadErrorSamplesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductFeedUploadErrorSample]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductFeedUploadErrorSample]](GetProductFeedUploadErrorSamplesBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadErrorSamplesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductFeedUploadErrorSample], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductFeedUploadErrorSample]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedUploadErrorSamplesWithResponse(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSamplesParams) (*core.Cursor[objects.ProductFeedUploadErrorSample], *core.Response, error) {
	var out core.Cursor[objects.ProductFeedUploadErrorSample]
	call := GetProductFeedUploadErrorSamplesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductFeedUploadErrorSamples(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSamplesParams) (*core.Cursor[objects.ProductFeedUploadErrorSample], error) {
	out, _, err := GetProductFeedUploadErrorSamplesWithResponse(ctx, client, id, params)
	return out, err
}

type GetProductFeedUploadErrorSuggestedRulesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadErrorSuggestedRulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadErrorSuggestedRulesBatchCall(id string, params GetProductFeedUploadErrorSuggestedRulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "suggested_rules"), params.ToParams(), options...)
}

func NewGetProductFeedUploadErrorSuggestedRulesBatchRequest(id string, params GetProductFeedUploadErrorSuggestedRulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.ProductFeedRuleSuggestion]] {
	return core.NewBatchRequest[core.Cursor[objects.ProductFeedRuleSuggestion]](GetProductFeedUploadErrorSuggestedRulesBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadErrorSuggestedRulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.ProductFeedRuleSuggestion], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.ProductFeedRuleSuggestion]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedUploadErrorSuggestedRulesWithResponse(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSuggestedRulesParams) (*core.Cursor[objects.ProductFeedRuleSuggestion], *core.Response, error) {
	var out core.Cursor[objects.ProductFeedRuleSuggestion]
	call := GetProductFeedUploadErrorSuggestedRulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductFeedUploadErrorSuggestedRules(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorSuggestedRulesParams) (*core.Cursor[objects.ProductFeedRuleSuggestion], error) {
	out, _, err := GetProductFeedUploadErrorSuggestedRulesWithResponse(ctx, client, id, params)
	return out, err
}

type GetProductFeedUploadErrorParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetProductFeedUploadErrorParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetProductFeedUploadErrorBatchCall(id string, params GetProductFeedUploadErrorParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetProductFeedUploadErrorBatchRequest(id string, params GetProductFeedUploadErrorParams, options ...core.BatchOption) *core.BatchRequest[objects.ProductFeedUploadError] {
	return core.NewBatchRequest[objects.ProductFeedUploadError](GetProductFeedUploadErrorBatchCall(id, params, options...))
}

func DecodeGetProductFeedUploadErrorBatchResponse(response *core.BatchResponse) (*objects.ProductFeedUploadError, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ProductFeedUploadError
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetProductFeedUploadErrorWithResponse(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorParams) (*objects.ProductFeedUploadError, *core.Response, error) {
	var out objects.ProductFeedUploadError
	call := GetProductFeedUploadErrorBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetProductFeedUploadError(ctx context.Context, client *core.Client, id string, params GetProductFeedUploadErrorParams) (*objects.ProductFeedUploadError, error) {
	out, _, err := GetProductFeedUploadErrorWithResponse(ctx, client, id, params)
	return out, err
}
