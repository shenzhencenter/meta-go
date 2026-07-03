package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPagePostExperimentVideoInsightsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostExperimentVideoInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostExperimentVideoInsightsBatchCall(id string, params GetPagePostExperimentVideoInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "video_insights"), params.ToParams(), options...)
}

func NewGetPagePostExperimentVideoInsightsBatchRequest(id string, params GetPagePostExperimentVideoInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult]] {
	return core.NewBatchRequest[core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult]](GetPagePostExperimentVideoInsightsBatchCall(id, params, options...))
}

func DecodeGetPagePostExperimentVideoInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPagePostExperimentVideoInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostExperimentVideoInsightsParams) (*core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult], *core.Response, error) {
	var out core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult]
	call := GetPagePostExperimentVideoInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostExperimentVideoInsights(ctx context.Context, client *core.Client, id string, params GetPagePostExperimentVideoInsightsParams) (*core.Cursor[objects.PagePostExperimentVideoInsightsQueryResult], error) {
	out, _, err := GetPagePostExperimentVideoInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type DeletePagePostExperimentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeletePagePostExperimentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeletePagePostExperimentBatchCall(id string, params DeletePagePostExperimentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeletePagePostExperimentBatchRequest(id string, params DeletePagePostExperimentParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeletePagePostExperimentBatchCall(id, params, options...))
}

func DecodeDeletePagePostExperimentBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeletePagePostExperimentWithResponse(ctx context.Context, client *core.Client, id string, params DeletePagePostExperimentParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeletePagePostExperimentBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeletePagePostExperiment(ctx context.Context, client *core.Client, id string, params DeletePagePostExperimentParams) (*map[string]interface{}, error) {
	out, _, err := DeletePagePostExperimentWithResponse(ctx, client, id, params)
	return out, err
}

type GetPagePostExperimentParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPagePostExperimentParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPagePostExperimentBatchCall(id string, params GetPagePostExperimentParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPagePostExperimentBatchRequest(id string, params GetPagePostExperimentParams, options ...core.BatchOption) *core.BatchRequest[objects.PagePostExperiment] {
	return core.NewBatchRequest[objects.PagePostExperiment](GetPagePostExperimentBatchCall(id, params, options...))
}

func DecodeGetPagePostExperimentBatchResponse(response *core.BatchResponse) (*objects.PagePostExperiment, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PagePostExperiment
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPagePostExperimentWithResponse(ctx context.Context, client *core.Client, id string, params GetPagePostExperimentParams) (*objects.PagePostExperiment, *core.Response, error) {
	var out objects.PagePostExperiment
	call := GetPagePostExperimentBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPagePostExperiment(ctx context.Context, client *core.Client, id string, params GetPagePostExperimentParams) (*objects.PagePostExperiment, error) {
	out, _, err := GetPagePostExperimentWithResponse(ctx, client, id, params)
	return out, err
}
