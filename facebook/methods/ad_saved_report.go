package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdSavedReportParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSavedReportParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSavedReport(ctx context.Context, client *core.Client, id string, params GetAdSavedReportParams) (*objects.AdSavedReport, error) {
	var out objects.AdSavedReport
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
