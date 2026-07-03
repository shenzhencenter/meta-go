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

func GetThirdPartyMeasurementReportDataset(ctx context.Context, client *core.Client, id string, params GetThirdPartyMeasurementReportDatasetParams) (*objects.ThirdPartyMeasurementReportDataset, error) {
	var out objects.ThirdPartyMeasurementReportDataset
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
