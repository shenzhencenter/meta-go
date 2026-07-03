package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsReportBuilderSavedReportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsReportBuilderSavedReportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsReportBuilderSavedReportBatchCall(id string, params GetAdsReportBuilderSavedReportParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsReportBuilderSavedReportBatchRequest(id string, params GetAdsReportBuilderSavedReportParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsReportBuilderSavedReport] {
	return core.NewBatchRequest[objects.AdsReportBuilderSavedReport](GetAdsReportBuilderSavedReportBatchCall(id, params, options...))
}

func DecodeGetAdsReportBuilderSavedReportBatchResponse(response *core.BatchResponse) (*objects.AdsReportBuilderSavedReport, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsReportBuilderSavedReport
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsReportBuilderSavedReport(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderSavedReportParams) (*objects.AdsReportBuilderSavedReport, error) {
	var out objects.AdsReportBuilderSavedReport
	call := GetAdsReportBuilderSavedReportBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
