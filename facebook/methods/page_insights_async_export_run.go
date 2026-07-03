package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetPageInsightsAsyncExportRunParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetPageInsightsAsyncExportRunParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetPageInsightsAsyncExportRun(ctx context.Context, client *core.Client, id string, params GetPageInsightsAsyncExportRunParams) (*objects.PageInsightsAsyncExportRun, error) {
	var out objects.PageInsightsAsyncExportRun
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
