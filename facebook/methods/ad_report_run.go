package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdReportRunInsightsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdReportRunInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdReportRunInsightsBatchCall(id string, params GetAdReportRunInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetAdReportRunInsightsBatchRequest(id string, params GetAdReportRunInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsInsights]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsInsights]](GetAdReportRunInsightsBatchCall(id, params, options...))
}

func DecodeGetAdReportRunInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsInsights], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsInsights]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdReportRunInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdReportRunInsightsParams) (*core.Cursor[objects.AdsInsights], *core.Response, error) {
	var out core.Cursor[objects.AdsInsights]
	call := GetAdReportRunInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdReportRunInsights(ctx context.Context, client *core.Client, id string, params GetAdReportRunInsightsParams) (*core.Cursor[objects.AdsInsights], error) {
	out, _, err := GetAdReportRunInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdReportRunParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdReportRunParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdReportRunBatchCall(id string, params GetAdReportRunParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdReportRunBatchRequest(id string, params GetAdReportRunParams, options ...core.BatchOption) *core.BatchRequest[objects.AdReportRun] {
	return core.NewBatchRequest[objects.AdReportRun](GetAdReportRunBatchCall(id, params, options...))
}

func DecodeGetAdReportRunBatchResponse(response *core.BatchResponse) (*objects.AdReportRun, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdReportRun
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdReportRunWithResponse(ctx context.Context, client *core.Client, id string, params GetAdReportRunParams) (*objects.AdReportRun, *core.Response, error) {
	var out objects.AdReportRun
	call := GetAdReportRunBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdReportRun(ctx context.Context, client *core.Client, id string, params GetAdReportRunParams) (*objects.AdReportRun, error) {
	out, _, err := GetAdReportRunWithResponse(ctx, client, id, params)
	return out, err
}
