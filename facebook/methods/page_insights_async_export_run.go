package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPageInsightsAsyncExportRunParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageInsightsAsyncExportRunParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageInsightsAsyncExportRunBatchCall(id string, params GetPageInsightsAsyncExportRunParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetPageInsightsAsyncExportRunBatchRequest(id string, params GetPageInsightsAsyncExportRunParams, options ...core.BatchOption) *core.BatchRequest[objects.PageInsightsAsyncExportRun] {
	return core.NewBatchRequest[objects.PageInsightsAsyncExportRun](GetPageInsightsAsyncExportRunBatchCall(id, params, options...))
}

func DecodeGetPageInsightsAsyncExportRunBatchResponse(response *core.BatchResponse) (*objects.PageInsightsAsyncExportRun, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.PageInsightsAsyncExportRun
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetPageInsightsAsyncExportRunWithResponse(ctx context.Context, client *core.Client, id string, params GetPageInsightsAsyncExportRunParams) (*objects.PageInsightsAsyncExportRun, *core.Response, error) {
	var out objects.PageInsightsAsyncExportRun
	call := GetPageInsightsAsyncExportRunBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetPageInsightsAsyncExportRun(ctx context.Context, client *core.Client, id string, params GetPageInsightsAsyncExportRunParams) (*objects.PageInsightsAsyncExportRun, error) {
	out, _, err := GetPageInsightsAsyncExportRunWithResponse(ctx, client, id, params)
	return out, err
}
