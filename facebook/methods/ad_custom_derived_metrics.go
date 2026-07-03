package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdCustomDerivedMetricsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdCustomDerivedMetricsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdCustomDerivedMetricsBatchCall(id string, params GetAdCustomDerivedMetricsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdCustomDerivedMetricsBatchRequest(id string, params GetAdCustomDerivedMetricsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdCustomDerivedMetrics] {
	return core.NewBatchRequest[objects.AdCustomDerivedMetrics](GetAdCustomDerivedMetricsBatchCall(id, params, options...))
}

func DecodeGetAdCustomDerivedMetricsBatchResponse(response *core.BatchResponse) (*objects.AdCustomDerivedMetrics, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdCustomDerivedMetrics
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdCustomDerivedMetricsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdCustomDerivedMetricsParams) (*objects.AdCustomDerivedMetrics, *core.Response, error) {
	var out objects.AdCustomDerivedMetrics
	call := GetAdCustomDerivedMetricsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdCustomDerivedMetrics(ctx context.Context, client *core.Client, id string, params GetAdCustomDerivedMetricsParams) (*objects.AdCustomDerivedMetrics, error) {
	out, _, err := GetAdCustomDerivedMetricsWithResponse(ctx, client, id, params)
	return out, err
}
