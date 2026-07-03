package methods

import (
	"context"
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/objects"
	"net/http"
)

type GetIGUserExportForCAMBrandedContentMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserExportForCAMBrandedContentMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserExportForCAMBrandedContentMedia(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMBrandedContentMediaParams) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	var out core.Cursor[objects.IGMediaExportForCAM]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "branded_content_media"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserExportForCAMInsightsParams struct {
	Breakdown *enums.IguserexportforcaminsightsBreakdownEnumParam `facebook:"breakdown"`
	Metrics   *[]enums.IguserexportforcaminsightsMetricsEnumParam `facebook:"metrics"`
	Period    *enums.IguserexportforcaminsightsPeriodEnumParam    `facebook:"period"`
	TimeRange *enums.IguserexportforcaminsightsTimeRangeEnumParam `facebook:"time_range"`
	Extra     core.Params                                         `facebook:"-"`
}

func (params GetIGUserExportForCAMInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Breakdown != nil {
		out["breakdown"] = *params.Breakdown
	}
	if params.Metrics != nil {
		out["metrics"] = *params.Metrics
	}
	if params.Period != nil {
		out["period"] = *params.Period
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	return out
}

func GetIGUserExportForCAMInsights(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMInsightsParams) (*core.Cursor[objects.IGUserExportForCAMInsightsResult], error) {
	var out core.Cursor[objects.IGUserExportForCAMInsightsResult]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserExportForCAMPastPartnershipAdsMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserExportForCAMPastPartnershipAdsMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserExportForCAMPastPartnershipAdsMedia(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMPastPartnershipAdsMediaParams) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	var out core.Cursor[objects.IGMediaExportForCAM]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "past_partnership_ads_media"), params.ToParams(), &out)
	return &out, err
}

type GetIGUserExportForCAMRecentMediaParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetIGUserExportForCAMRecentMediaParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetIGUserExportForCAMRecentMedia(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMRecentMediaParams) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	var out core.Cursor[objects.IGMediaExportForCAM]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "recent_media"), params.ToParams(), &out)
	return &out, err
}
