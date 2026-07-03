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

func GetAdsReportBuilderMMMReportScheduler(ctx context.Context, client *core.Client, id string, params GetAdsReportBuilderMMMReportSchedulerParams) (*objects.AdsReportBuilderMMMReportScheduler, error) {
	var out objects.AdsReportBuilderMMMReportScheduler
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
