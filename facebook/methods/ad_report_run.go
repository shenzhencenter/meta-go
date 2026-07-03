package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetAdReportRunInsightsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdReportRunInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdReportRunInsights(ctx context.Context, client *core.Client, id string, params GetAdReportRunInsightsParams) (*core.Cursor[objects.AdsInsights], error) {
	var out core.Cursor[objects.AdsInsights]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type GetAdReportRunParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdReportRunParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdReportRun(ctx context.Context, client *core.Client, id string, params GetAdReportRunParams) (*objects.AdReportRun, error) {
	var out objects.AdReportRun
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
