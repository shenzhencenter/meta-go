package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetMeasurementReport(ctx context.Context, client *core.Client, id string, params GetMeasurementReportParams) (*objects.MeasurementReport, error) {
	var out objects.MeasurementReport
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
