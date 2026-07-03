package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetMeasurementReportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetMeasurementReportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetMeasurementReportBatchCall(id string, params GetMeasurementReportParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetMeasurementReportBatchRequest(id string, params GetMeasurementReportParams, options ...core.BatchOption) *core.BatchRequest[objects.MeasurementReport] {
	return core.NewBatchRequest[objects.MeasurementReport](GetMeasurementReportBatchCall(id, params, options...))
}

func DecodeGetMeasurementReportBatchResponse(response *core.BatchResponse) (*objects.MeasurementReport, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.MeasurementReport
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetMeasurementReportWithResponse(ctx context.Context, client *core.Client, id string, params GetMeasurementReportParams) (*objects.MeasurementReport, *core.Response, error) {
	var out objects.MeasurementReport
	call := GetMeasurementReportBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetMeasurementReport(ctx context.Context, client *core.Client, id string, params GetMeasurementReportParams) (*objects.MeasurementReport, error) {
	out, _, err := GetMeasurementReportWithResponse(ctx, client, id, params)
	return out, err
}
