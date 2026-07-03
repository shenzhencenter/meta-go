package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdSavedReportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSavedReportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSavedReportBatchCall(id string, params GetAdSavedReportParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdSavedReportBatchRequest(id string, params GetAdSavedReportParams, options ...core.BatchOption) *core.BatchRequest[objects.AdSavedReport] {
	return core.NewBatchRequest[objects.AdSavedReport](GetAdSavedReportBatchCall(id, params, options...))
}

func DecodeGetAdSavedReportBatchResponse(response *core.BatchResponse) (*objects.AdSavedReport, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdSavedReport
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdSavedReport(ctx context.Context, client *core.Client, id string, params GetAdSavedReportParams) (*objects.AdSavedReport, error) {
	var out objects.AdSavedReport
	call := GetAdSavedReportBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
