package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsReportBuilderMMMReportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsReportBuilderMMMReportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsReportBuilderMMMReportBatchCall(id string, params GetAdsReportBuilderMMMReportParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsReportBuilderMMMReportBatchRequest(id string, params GetAdsReportBuilderMMMReportParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsReportBuilderMMMReport] {
	return core.NewBatchRequest[objects.AdsReportBuilderMMMReport](GetAdsReportBuilderMMMReportBatchCall(id, params, options...))
}

func DecodeGetAdsReportBuilderMMMReportBatchResponse(response *core.BatchResponse) (*objects.AdsReportBuilderMMMReport, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsReportBuilderMMMReport
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsReportBuilderMMMReport(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderMMMReportParams) (*objects.AdsReportBuilderMMMReport, error) {
	var out objects.AdsReportBuilderMMMReport
	call := GetAdsReportBuilderMMMReportBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
