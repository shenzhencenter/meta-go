package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
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

func GetAdsReportBuilderSavedReport(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderSavedReportParams) (*objects.AdsReportBuilderSavedReport, error) {
	var out objects.AdsReportBuilderSavedReport
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
