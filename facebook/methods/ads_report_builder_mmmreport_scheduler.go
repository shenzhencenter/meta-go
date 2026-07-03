package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdsReportBuilderMMMReportSchedulerParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdsReportBuilderMMMReportSchedulerParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdsReportBuilderMMMReportSchedulerBatchCall(id string, params GetAdsReportBuilderMMMReportSchedulerParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdsReportBuilderMMMReportSchedulerBatchRequest(id string, params GetAdsReportBuilderMMMReportSchedulerParams, options ...core.BatchOption) *core.BatchRequest[objects.AdsReportBuilderMMMReportScheduler] {
	return core.NewBatchRequest[objects.AdsReportBuilderMMMReportScheduler](GetAdsReportBuilderMMMReportSchedulerBatchCall(id, params, options...))
}

func DecodeGetAdsReportBuilderMMMReportSchedulerBatchResponse(response *core.BatchResponse) (*objects.AdsReportBuilderMMMReportScheduler, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdsReportBuilderMMMReportScheduler
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdsReportBuilderMMMReportSchedulerWithResponse(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderMMMReportSchedulerParams) (*objects.AdsReportBuilderMMMReportScheduler, *core.Response, error) {
	var out objects.AdsReportBuilderMMMReportScheduler
	call := GetAdsReportBuilderMMMReportSchedulerBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdsReportBuilderMMMReportScheduler(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderMMMReportSchedulerParams) (*objects.AdsReportBuilderMMMReportScheduler, error) {
	out, _, err := GetAdsReportBuilderMMMReportSchedulerWithResponse(ctx, client, id, params)
	return out, err
}
