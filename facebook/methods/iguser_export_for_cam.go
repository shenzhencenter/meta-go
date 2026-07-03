package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
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

func GetIGUserExportForCAMBrandedContentMediaBatchCall(id string, params GetIGUserExportForCAMBrandedContentMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "branded_content_media"), params.ToParams(), options...)
}

func NewGetIGUserExportForCAMBrandedContentMediaBatchRequest(id string, params GetIGUserExportForCAMBrandedContentMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMediaExportForCAM]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMediaExportForCAM]](GetIGUserExportForCAMBrandedContentMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserExportForCAMBrandedContentMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMediaExportForCAM]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserExportForCAMBrandedContentMedia(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMBrandedContentMediaParams) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	var out core.Cursor[objects.IGMediaExportForCAM]
	call := GetIGUserExportForCAMBrandedContentMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetIGUserExportForCAMInsightsBatchCall(id string, params GetIGUserExportForCAMInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetIGUserExportForCAMInsightsBatchRequest(id string, params GetIGUserExportForCAMInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGUserExportForCAMInsightsResult]] {
	return core.NewBatchRequest[core.Cursor[objects.IGUserExportForCAMInsightsResult]](GetIGUserExportForCAMInsightsBatchCall(id, params, options...))
}

func DecodeGetIGUserExportForCAMInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGUserExportForCAMInsightsResult], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGUserExportForCAMInsightsResult]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserExportForCAMInsights(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMInsightsParams) (*core.Cursor[objects.IGUserExportForCAMInsightsResult], error) {
	var out core.Cursor[objects.IGUserExportForCAMInsightsResult]
	call := GetIGUserExportForCAMInsightsBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetIGUserExportForCAMPastPartnershipAdsMediaBatchCall(id string, params GetIGUserExportForCAMPastPartnershipAdsMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "past_partnership_ads_media"), params.ToParams(), options...)
}

func NewGetIGUserExportForCAMPastPartnershipAdsMediaBatchRequest(id string, params GetIGUserExportForCAMPastPartnershipAdsMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMediaExportForCAM]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMediaExportForCAM]](GetIGUserExportForCAMPastPartnershipAdsMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserExportForCAMPastPartnershipAdsMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMediaExportForCAM]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserExportForCAMPastPartnershipAdsMedia(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMPastPartnershipAdsMediaParams) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	var out core.Cursor[objects.IGMediaExportForCAM]
	call := GetIGUserExportForCAMPastPartnershipAdsMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
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

func GetIGUserExportForCAMRecentMediaBatchCall(id string, params GetIGUserExportForCAMRecentMediaParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "recent_media"), params.ToParams(), options...)
}

func NewGetIGUserExportForCAMRecentMediaBatchRequest(id string, params GetIGUserExportForCAMRecentMediaParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.IGMediaExportForCAM]] {
	return core.NewBatchRequest[core.Cursor[objects.IGMediaExportForCAM]](GetIGUserExportForCAMRecentMediaBatchCall(id, params, options...))
}

func DecodeGetIGUserExportForCAMRecentMediaBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.IGMediaExportForCAM]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetIGUserExportForCAMRecentMedia(ctx context.Context, client *core.Client, id string, params GetIGUserExportForCAMRecentMediaParams) (*core.Cursor[objects.IGMediaExportForCAM], error) {
	var out core.Cursor[objects.IGMediaExportForCAM]
	call := GetIGUserExportForCAMRecentMediaBatchCall(id, params)
	err := client.Request(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, err
}
