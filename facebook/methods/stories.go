package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetStoriesInsightsParams struct {
	Metric *[]enums.StoriesinsightsMetricEnumParam `facebook:"metric"`
	Extra  core.Params                             `facebook:"-"`
}

func (params GetStoriesInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Metric != nil {
		out["metric"] = *params.Metric
	}
	return out
}

func GetStoriesInsightsBatchCall(id string, params GetStoriesInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetStoriesInsightsBatchRequest(id string, params GetStoriesInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.InsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.InsightsResult]](GetStoriesInsightsBatchCall(id, params, options...))
}

func DecodeGetStoriesInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.InsightsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.InsightsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetStoriesInsights(ctx context.Context, client *core.Client, id string, params GetStoriesInsightsParams) (*core.Cursor[objects.InsightsResult], error) {
	var out core.Cursor[objects.InsightsResult]
	call := GetStoriesInsightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}

type GetStoriesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetStoriesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetStoriesBatchCall(id string, params GetStoriesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetStoriesBatchRequest(id string, params GetStoriesParams, options ...core.BatchOption) *core.BatchRequest[objects.Stories] {
	return core.NewBatchRequest[objects.Stories](GetStoriesBatchCall(id, params, options...))
}

func DecodeGetStoriesBatchResponse(response *core.BatchResponse) (*objects.Stories, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Stories
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetStories(ctx context.Context, client *core.Client, id string, params GetStoriesParams) (*objects.Stories, error) {
	var out objects.Stories
	call := GetStoriesBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
