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

func GetAdsReportBuilderExportCore(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderExportCoreParams) (*objects.AdsReportBuilderExportCore, error) {
	var out objects.AdsReportBuilderExportCore
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
