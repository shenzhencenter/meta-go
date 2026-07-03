package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetThirdPartyMeasurementReportDatasetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetThirdPartyMeasurementReportDatasetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetThirdPartyMeasurementReportDatasetBatchCall(id string, params GetThirdPartyMeasurementReportDatasetParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetThirdPartyMeasurementReportDatasetBatchRequest(id string, params GetThirdPartyMeasurementReportDatasetParams, options ...core.BatchOption) *core.BatchRequest[objects.ThirdPartyMeasurementReportDataset] {
	return core.NewBatchRequest[objects.ThirdPartyMeasurementReportDataset](GetThirdPartyMeasurementReportDatasetBatchCall(id, params, options...))
}

func DecodeGetThirdPartyMeasurementReportDatasetBatchResponse(response *core.BatchResponse) (*objects.ThirdPartyMeasurementReportDataset, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.ThirdPartyMeasurementReportDataset
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetThirdPartyMeasurementReportDatasetWithResponse(ctx context.Context, client *core.Client, id string, params GetThirdPartyMeasurementReportDatasetParams) (*objects.ThirdPartyMeasurementReportDataset, *core.Response, error) {
	var out objects.ThirdPartyMeasurementReportDataset
	call := GetThirdPartyMeasurementReportDatasetBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetThirdPartyMeasurementReportDataset(ctx context.Context, client *core.Client, id string, params GetThirdPartyMeasurementReportDatasetParams) (*objects.ThirdPartyMeasurementReportDataset, error) {
	out, _, err := GetThirdPartyMeasurementReportDatasetWithResponse(ctx, client, id, params)
	return out, err
}
