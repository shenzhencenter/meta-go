package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsReportBuilderExportCoreParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsReportBuilderExportCoreParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsReportBuilderExportCoreBatchCall(id string, params GetAdsReportBuilderExportCoreParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsReportBuilderExportCoreBatchRequest(id string, params GetAdsReportBuilderExportCoreParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsReportBuilderExportCore] {
	return core.NewBatchRequest[objects.AdsReportBuilderExportCore](GetAdsReportBuilderExportCoreBatchCall(id, params, options...))
}

func DecodeGetAdsReportBuilderExportCoreBatchResponse(response *core.BatchResponse) (*objects.AdsReportBuilderExportCore, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsReportBuilderExportCore
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsReportBuilderExportCore(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderExportCoreParams) (*objects.AdsReportBuilderExportCore, error) {
	var out objects.AdsReportBuilderExportCore
	call := GetAdsReportBuilderExportCoreBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
