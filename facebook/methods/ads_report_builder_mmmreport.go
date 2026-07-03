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

func GetAdsReportBuilderMMMReport(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderMMMReportParams) (*objects.AdsReportBuilderMMMReport, error) {
	var out objects.AdsReportBuilderMMMReport
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
