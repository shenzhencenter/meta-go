package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetAdAdcreativesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdAdcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdAdcreativesBatchCall(id string, params GetAdAdcreativesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adcreatives"), params.ToParams(), options...)
}

func NewGetAdAdcreativesBatchRequest(id string, params GetAdAdcreativesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdCreative]] {
	return core.NewBatchRequest[core.Cursor[objects.AdCreative]](GetAdAdcreativesBatchCall(id, params, options...))
}

func DecodeGetAdAdcreativesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdCreative], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdCreative]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAdcreativesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdAdcreativesParams) (*core.Cursor[objects.AdCreative], *core.Response, error) {
	var out core.Cursor[objects.AdCreative]
	call := GetAdAdcreativesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdAdcreatives(ctx context.Context, client *core.Client, id string, params GetAdAdcreativesParams) (*core.Cursor[objects.AdCreative], error) {
	out, _, err := GetAdAdcreativesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdAdlabelsParams struct {
	Adlabels         []map[string]interface{}                          `facebook:"adlabels"`
	ExecutionOptions *[]enums.AdgroupadlabelsExecutionOptionsEnumParam `facebook:"execution_options"`
	Extra            core.Params                                       `facebook:"-"`
}

func (params CreateAdAdlabelsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["adlabels"] = params.Adlabels
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	return out
}

func CreateAdAdlabelsBatchCall(id string, params CreateAdAdlabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adlabels"), params.ToParams(), options...)
}

func NewCreateAdAdlabelsBatchRequest(id string, params CreateAdAdlabelsParams, options ...core.BatchOption) *core.BatchRequest[objects.Ad] {
	return core.NewBatchRequest[objects.Ad](CreateAdAdlabelsBatchCall(id, params, options...))
}

func DecodeCreateAdAdlabelsBatchResponse(response *core.BatchResponse) (*objects.Ad, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Ad
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdAdlabelsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdAdlabelsParams) (*objects.Ad, *core.Response, error) {
	var out objects.Ad
	call := CreateAdAdlabelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdAdlabels(ctx context.Context, client *core.Client, id string, params CreateAdAdlabelsParams) (*objects.Ad, error) {
	out, _, err := CreateAdAdlabelsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdAdrulesGovernedParams struct {
	PassEvaluation *bool       `facebook:"pass_evaluation"`
	Extra          core.Params `facebook:"-"`
}

func (params GetAdAdrulesGovernedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.PassEvaluation != nil {
		out["pass_evaluation"] = *params.PassEvaluation
	}
	return out
}

func GetAdAdrulesGovernedBatchCall(id string, params GetAdAdrulesGovernedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adrules_governed"), params.ToParams(), options...)
}

func NewGetAdAdrulesGovernedBatchRequest(id string, params GetAdAdrulesGovernedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdRule]] {
	return core.NewBatchRequest[core.Cursor[objects.AdRule]](GetAdAdrulesGovernedBatchCall(id, params, options...))
}

func DecodeGetAdAdrulesGovernedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdRule], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdRule]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdAdrulesGovernedWithResponse(ctx context.Context, client *core.Client, id string, params GetAdAdrulesGovernedParams) (*core.Cursor[objects.AdRule], *core.Response, error) {
	var out core.Cursor[objects.AdRule]
	call := GetAdAdrulesGovernedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdAdrulesGoverned(ctx context.Context, client *core.Client, id string, params GetAdAdrulesGovernedParams) (*core.Cursor[objects.AdRule], error) {
	out, _, err := GetAdAdrulesGovernedWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdCopiesParams struct {
	DatePreset      *enums.AdgroupcopiesDatePresetEnumParam `facebook:"date_preset"`
	EffectiveStatus *[]string                               `facebook:"effective_status"`
	TimeRange       *map[string]interface{}                 `facebook:"time_range"`
	UpdatedSince    *int                                    `facebook:"updated_since"`
	Extra           core.Params                             `facebook:"-"`
}

func (params GetAdCopiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.EffectiveStatus != nil {
		out["effective_status"] = *params.EffectiveStatus
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	if params.UpdatedSince != nil {
		out["updated_since"] = *params.UpdatedSince
	}
	return out
}

func GetAdCopiesBatchCall(id string, params GetAdCopiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "copies"), params.ToParams(), options...)
}

func NewGetAdCopiesBatchRequest(id string, params GetAdCopiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Ad]] {
	return core.NewBatchRequest[core.Cursor[objects.Ad]](GetAdCopiesBatchCall(id, params, options...))
}

func DecodeGetAdCopiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Ad], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Ad]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdCopiesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdCopiesParams) (*core.Cursor[objects.Ad], *core.Response, error) {
	var out core.Cursor[objects.Ad]
	call := GetAdCopiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdCopies(ctx context.Context, client *core.Client, id string, params GetAdCopiesParams) (*core.Cursor[objects.Ad], error) {
	out, _, err := GetAdCopiesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdCopiesParams struct {
	AdsetID            *core.ID                                  `facebook:"adset_id"`
	CreativeParameters *objects.AdCreative                       `facebook:"creative_parameters"`
	RenameOptions      *map[string]interface{}                   `facebook:"rename_options"`
	StatusOption       *enums.AdgroupcopiesStatusOptionEnumParam `facebook:"status_option"`
	Extra              core.Params                               `facebook:"-"`
}

func (params CreateAdCopiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdsetID != nil {
		out["adset_id"] = *params.AdsetID
	}
	if params.CreativeParameters != nil {
		out["creative_parameters"] = *params.CreativeParameters
	}
	if params.RenameOptions != nil {
		out["rename_options"] = *params.RenameOptions
	}
	if params.StatusOption != nil {
		out["status_option"] = *params.StatusOption
	}
	return out
}

func CreateAdCopiesBatchCall(id string, params CreateAdCopiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "copies"), params.ToParams(), options...)
}

func NewCreateAdCopiesBatchRequest(id string, params CreateAdCopiesParams, options ...core.BatchOption) *core.BatchRequest[objects.Ad] {
	return core.NewBatchRequest[objects.Ad](CreateAdCopiesBatchCall(id, params, options...))
}

func DecodeCreateAdCopiesBatchResponse(response *core.BatchResponse) (*objects.Ad, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Ad
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdCopiesWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdCopiesParams) (*objects.Ad, *core.Response, error) {
	var out objects.Ad
	call := CreateAdCopiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdCopies(ctx context.Context, client *core.Client, id string, params CreateAdCopiesParams) (*objects.Ad, error) {
	out, _, err := CreateAdCopiesWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdgroupinsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdgroupinsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdgroupinsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdgroupinsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdgroupinsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                     `facebook:"default_summary"`
	ExportColumns                *[]string                                                 `facebook:"export_columns"`
	ExportFormat                 *string                                                   `facebook:"export_format"`
	ExportName                   *string                                                   `facebook:"export_name"`
	Fields                       *[]string                                                 `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                 `facebook:"filtering"`
	GraphCache                   *bool                                                     `facebook:"graph_cache"`
	Level                        *enums.AdgroupinsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                      `facebook:"limit"`
	ProductIDLimit               *int                                                      `facebook:"product_id_limit"`
	Sort                         *[]string                                                 `facebook:"sort"`
	Summary                      *[]string                                                 `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdgroupinsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                   `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                   `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                 `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                     `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                     `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                               `facebook:"-"`
}

func (params GetAdInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActionAttributionWindows != nil {
		out["action_attribution_windows"] = *params.ActionAttributionWindows
	}
	if params.ActionBreakdowns != nil {
		out["action_breakdowns"] = *params.ActionBreakdowns
	}
	if params.ActionReportTime != nil {
		out["action_report_time"] = *params.ActionReportTime
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.DefaultSummary != nil {
		out["default_summary"] = *params.DefaultSummary
	}
	if params.ExportColumns != nil {
		out["export_columns"] = *params.ExportColumns
	}
	if params.ExportFormat != nil {
		out["export_format"] = *params.ExportFormat
	}
	if params.ExportName != nil {
		out["export_name"] = *params.ExportName
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	if params.GraphCache != nil {
		out["graph_cache"] = *params.GraphCache
	}
	if params.Level != nil {
		out["level"] = *params.Level
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.ProductIDLimit != nil {
		out["product_id_limit"] = *params.ProductIDLimit
	}
	if params.Sort != nil {
		out["sort"] = *params.Sort
	}
	if params.Summary != nil {
		out["summary"] = *params.Summary
	}
	if params.SummaryActionBreakdowns != nil {
		out["summary_action_breakdowns"] = *params.SummaryActionBreakdowns
	}
	if params.TimeIncrement != nil {
		out["time_increment"] = *params.TimeIncrement
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	if params.TimeRanges != nil {
		out["time_ranges"] = *params.TimeRanges
	}
	if params.UseAccountAttributionSetting != nil {
		out["use_account_attribution_setting"] = *params.UseAccountAttributionSetting
	}
	if params.UseUnifiedAttributionSetting != nil {
		out["use_unified_attribution_setting"] = *params.UseUnifiedAttributionSetting
	}
	return out
}

func GetAdInsightsBatchCall(id string, params GetAdInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetAdInsightsBatchRequest(id string, params GetAdInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsInsights]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsInsights]](GetAdInsightsBatchCall(id, params, options...))
}

func DecodeGetAdInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsInsights], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdsInsights]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdInsightsParams) (*core.Cursor[objects.AdsInsights], *core.Response, error) {
	var out core.Cursor[objects.AdsInsights]
	call := GetAdInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdInsights(ctx context.Context, client *core.Client, id string, params GetAdInsightsParams) (*core.Cursor[objects.AdsInsights], error) {
	out, _, err := GetAdInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateAdInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdgroupinsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdgroupinsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdgroupinsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdgroupinsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdgroupinsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                     `facebook:"default_summary"`
	ExportColumns                *[]string                                                 `facebook:"export_columns"`
	ExportFormat                 *string                                                   `facebook:"export_format"`
	ExportName                   *string                                                   `facebook:"export_name"`
	Fields                       *[]string                                                 `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                 `facebook:"filtering"`
	GraphCache                   *bool                                                     `facebook:"graph_cache"`
	Level                        *enums.AdgroupinsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                      `facebook:"limit"`
	ProductIDLimit               *int                                                      `facebook:"product_id_limit"`
	Sort                         *[]string                                                 `facebook:"sort"`
	Summary                      *[]string                                                 `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdgroupinsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                   `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                   `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                 `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                     `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                     `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                               `facebook:"-"`
}

func (params CreateAdInsightsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.ActionAttributionWindows != nil {
		out["action_attribution_windows"] = *params.ActionAttributionWindows
	}
	if params.ActionBreakdowns != nil {
		out["action_breakdowns"] = *params.ActionBreakdowns
	}
	if params.ActionReportTime != nil {
		out["action_report_time"] = *params.ActionReportTime
	}
	if params.Breakdowns != nil {
		out["breakdowns"] = *params.Breakdowns
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.DefaultSummary != nil {
		out["default_summary"] = *params.DefaultSummary
	}
	if params.ExportColumns != nil {
		out["export_columns"] = *params.ExportColumns
	}
	if params.ExportFormat != nil {
		out["export_format"] = *params.ExportFormat
	}
	if params.ExportName != nil {
		out["export_name"] = *params.ExportName
	}
	if params.Fields != nil {
		out["fields"] = *params.Fields
	}
	if params.Filtering != nil {
		out["filtering"] = *params.Filtering
	}
	if params.GraphCache != nil {
		out["graph_cache"] = *params.GraphCache
	}
	if params.Level != nil {
		out["level"] = *params.Level
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.ProductIDLimit != nil {
		out["product_id_limit"] = *params.ProductIDLimit
	}
	if params.Sort != nil {
		out["sort"] = *params.Sort
	}
	if params.Summary != nil {
		out["summary"] = *params.Summary
	}
	if params.SummaryActionBreakdowns != nil {
		out["summary_action_breakdowns"] = *params.SummaryActionBreakdowns
	}
	if params.TimeIncrement != nil {
		out["time_increment"] = *params.TimeIncrement
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	if params.TimeRanges != nil {
		out["time_ranges"] = *params.TimeRanges
	}
	if params.UseAccountAttributionSetting != nil {
		out["use_account_attribution_setting"] = *params.UseAccountAttributionSetting
	}
	if params.UseUnifiedAttributionSetting != nil {
		out["use_unified_attribution_setting"] = *params.UseUnifiedAttributionSetting
	}
	return out
}

func CreateAdInsightsBatchCall(id string, params CreateAdInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewCreateAdInsightsBatchRequest(id string, params CreateAdInsightsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdReportRun] {
	return core.NewBatchRequest[objects.AdReportRun](CreateAdInsightsBatchCall(id, params, options...))
}

func DecodeCreateAdInsightsBatchResponse(response *core.BatchResponse) (*objects.AdReportRun, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.AdReportRun
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateAdInsightsWithResponse(ctx context.Context, client *core.Client, id string, params CreateAdInsightsParams) (*objects.AdReportRun, *core.Response, error) {
	var out objects.AdReportRun
	call := CreateAdInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateAdInsights(ctx context.Context, client *core.Client, id string, params CreateAdInsightsParams) (*objects.AdReportRun, error) {
	out, _, err := CreateAdInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdLeadsParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdLeadsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdLeadsBatchCall(id string, params GetAdLeadsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "leads"), params.ToParams(), options...)
}

func NewGetAdLeadsBatchRequest(id string, params GetAdLeadsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Lead]] {
	return core.NewBatchRequest[core.Cursor[objects.Lead]](GetAdLeadsBatchCall(id, params, options...))
}

func DecodeGetAdLeadsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Lead], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Lead]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdLeadsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdLeadsParams) (*core.Cursor[objects.Lead], *core.Response, error) {
	var out core.Cursor[objects.Lead]
	call := GetAdLeadsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdLeads(ctx context.Context, client *core.Client, id string, params GetAdLeadsParams) (*core.Cursor[objects.Lead], error) {
	out, _, err := GetAdLeadsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdPreviewsParams struct {
	AdFormat             enums.AdgrouppreviewsAdFormatEnumParam         `facebook:"ad_format"`
	CreativeFeature      *enums.AdgrouppreviewsCreativeFeatureEnumParam `facebook:"creative_feature"`
	DynamicAssetLabel    *string                                        `facebook:"dynamic_asset_label"`
	DynamicCreativeSpec  *map[string]interface{}                        `facebook:"dynamic_creative_spec"`
	DynamicCustomization *map[string]interface{}                        `facebook:"dynamic_customization"`
	EndDate              *core.Time                                     `facebook:"end_date"`
	Height               *uint64                                        `facebook:"height"`
	Locale               *string                                        `facebook:"locale"`
	PlacePageID          *core.ID                                       `facebook:"place_page_id"`
	Post                 *map[string]interface{}                        `facebook:"post"`
	ProductItemIds       *[]core.ID                                     `facebook:"product_item_ids"`
	RenderType           *enums.AdgrouppreviewsRenderTypeEnumParam      `facebook:"render_type"`
	StartDate            *core.Time                                     `facebook:"start_date"`
	Width                *uint64                                        `facebook:"width"`
	Extra                core.Params                                    `facebook:"-"`
}

func (params GetAdPreviewsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["ad_format"] = params.AdFormat
	if params.CreativeFeature != nil {
		out["creative_feature"] = *params.CreativeFeature
	}
	if params.DynamicAssetLabel != nil {
		out["dynamic_asset_label"] = *params.DynamicAssetLabel
	}
	if params.DynamicCreativeSpec != nil {
		out["dynamic_creative_spec"] = *params.DynamicCreativeSpec
	}
	if params.DynamicCustomization != nil {
		out["dynamic_customization"] = *params.DynamicCustomization
	}
	if params.EndDate != nil {
		out["end_date"] = *params.EndDate
	}
	if params.Height != nil {
		out["height"] = *params.Height
	}
	if params.Locale != nil {
		out["locale"] = *params.Locale
	}
	if params.PlacePageID != nil {
		out["place_page_id"] = *params.PlacePageID
	}
	if params.Post != nil {
		out["post"] = *params.Post
	}
	if params.ProductItemIds != nil {
		out["product_item_ids"] = *params.ProductItemIds
	}
	if params.RenderType != nil {
		out["render_type"] = *params.RenderType
	}
	if params.StartDate != nil {
		out["start_date"] = *params.StartDate
	}
	if params.Width != nil {
		out["width"] = *params.Width
	}
	return out
}

func GetAdPreviewsBatchCall(id string, params GetAdPreviewsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "previews"), params.ToParams(), options...)
}

func NewGetAdPreviewsBatchRequest(id string, params GetAdPreviewsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdPreview]] {
	return core.NewBatchRequest[core.Cursor[objects.AdPreview]](GetAdPreviewsBatchCall(id, params, options...))
}

func DecodeGetAdPreviewsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdPreview], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdPreview]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdPreviewsWithResponse(ctx context.Context, client *core.Client, id string, params GetAdPreviewsParams) (*core.Cursor[objects.AdPreview], *core.Response, error) {
	var out core.Cursor[objects.AdPreview]
	call := GetAdPreviewsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdPreviews(ctx context.Context, client *core.Client, id string, params GetAdPreviewsParams) (*core.Cursor[objects.AdPreview], error) {
	out, _, err := GetAdPreviewsWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdTargetingsentencelinesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdTargetingsentencelinesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdTargetingsentencelinesBatchCall(id string, params GetAdTargetingsentencelinesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "targetingsentencelines"), params.ToParams(), options...)
}

func NewGetAdTargetingsentencelinesBatchRequest(id string, params GetAdTargetingsentencelinesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.TargetingSentenceLine]] {
	return core.NewBatchRequest[core.Cursor[objects.TargetingSentenceLine]](GetAdTargetingsentencelinesBatchCall(id, params, options...))
}

func DecodeGetAdTargetingsentencelinesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.TargetingSentenceLine], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.TargetingSentenceLine]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdTargetingsentencelinesWithResponse(ctx context.Context, client *core.Client, id string, params GetAdTargetingsentencelinesParams) (*core.Cursor[objects.TargetingSentenceLine], *core.Response, error) {
	var out core.Cursor[objects.TargetingSentenceLine]
	call := GetAdTargetingsentencelinesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAdTargetingsentencelines(ctx context.Context, client *core.Client, id string, params GetAdTargetingsentencelinesParams) (*core.Cursor[objects.TargetingSentenceLine], error) {
	out, _, err := GetAdTargetingsentencelinesWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteAdParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdBatchCall(id string, params DeleteAdParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteAdBatchRequest(id string, params DeleteAdParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteAdBatchCall(id, params, options...))
}

func DecodeDeleteAdBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out map[string]interface{}
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func DeleteAdWithResponse(ctx context.Context, client *core.Client, id string, params DeleteAdParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteAdBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteAd(ctx context.Context, client *core.Client, id string, params DeleteAdParams) (*map[string]interface{}, error) {
	out, _, err := DeleteAdWithResponse(ctx, client, id, params)
	return out, err
}

type GetAdParams struct {
	AmCallTags              *map[string]interface{}  `facebook:"am_call_tags"`
	DatePreset              *enums.AdgroupDatePreset `facebook:"date_preset"`
	FromAdtable             *bool                    `facebook:"from_adtable"`
	ReviewFeedbackBreakdown *bool                    `facebook:"review_feedback_breakdown"`
	TimeRange               *map[string]interface{}  `facebook:"time_range"`
	Extra                   core.Params              `facebook:"-"`
}

func (params GetAdParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AmCallTags != nil {
		out["am_call_tags"] = *params.AmCallTags
	}
	if params.DatePreset != nil {
		out["date_preset"] = *params.DatePreset
	}
	if params.FromAdtable != nil {
		out["from_adtable"] = *params.FromAdtable
	}
	if params.ReviewFeedbackBreakdown != nil {
		out["review_feedback_breakdown"] = *params.ReviewFeedbackBreakdown
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	return out
}

func GetAdBatchCall(id string, params GetAdParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetAdBatchRequest(id string, params GetAdParams, options ...core.BatchOption) *core.BatchRequest[objects.Ad] {
	return core.NewBatchRequest[objects.Ad](GetAdBatchCall(id, params, options...))
}

func DecodeGetAdBatchResponse(response *core.BatchResponse) (*objects.Ad, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Ad
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetAdWithResponse(ctx context.Context, client *core.Client, id string, params GetAdParams) (*objects.Ad, *core.Response, error) {
	var out objects.Ad
	call := GetAdBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetAd(ctx context.Context, client *core.Client, id string, params GetAdParams) (*objects.Ad, error) {
	out, _, err := GetAdWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateAdParams struct {
	AdScheduleEndTime       *core.Time                       `facebook:"ad_schedule_end_time"`
	AdScheduleStartTime     *core.Time                       `facebook:"ad_schedule_start_time"`
	Adlabels                *[]map[string]interface{}        `facebook:"adlabels"`
	AdsetSpec               *objects.AdSet                   `facebook:"adset_spec"`
	AudienceID              *core.ID                         `facebook:"audience_id"`
	BidAmount               *int                             `facebook:"bid_amount"`
	ConversionDomain        *string                          `facebook:"conversion_domain"`
	Creative                *objects.AdCreative              `facebook:"creative"`
	CreativeAssetGroupsSpec *map[string]interface{}          `facebook:"creative_asset_groups_spec"`
	CreativeAutomationSpec  *map[string]interface{}          `facebook:"creative_automation_spec"`
	DisplaySequence         *uint64                          `facebook:"display_sequence"`
	DraftAdgroupID          *core.ID                         `facebook:"draft_adgroup_id"`
	EngagementAudience      *bool                            `facebook:"engagement_audience"`
	ExecutionOptions        *[]enums.AdgroupExecutionOptions `facebook:"execution_options"`
	IncludeDemolinkHashes   *bool                            `facebook:"include_demolink_hashes"`
	Name                    *string                          `facebook:"name"`
	Priority                *uint64                          `facebook:"priority"`
	Status                  *enums.AdgroupStatus             `facebook:"status"`
	TrackingSpecs           *map[string]interface{}          `facebook:"tracking_specs"`
	Extra                   core.Params                      `facebook:"-"`
}

func (params UpdateAdParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AdScheduleEndTime != nil {
		out["ad_schedule_end_time"] = *params.AdScheduleEndTime
	}
	if params.AdScheduleStartTime != nil {
		out["ad_schedule_start_time"] = *params.AdScheduleStartTime
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.AdsetSpec != nil {
		out["adset_spec"] = *params.AdsetSpec
	}
	if params.AudienceID != nil {
		out["audience_id"] = *params.AudienceID
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	if params.ConversionDomain != nil {
		out["conversion_domain"] = *params.ConversionDomain
	}
	if params.Creative != nil {
		out["creative"] = *params.Creative
	}
	if params.CreativeAssetGroupsSpec != nil {
		out["creative_asset_groups_spec"] = *params.CreativeAssetGroupsSpec
	}
	if params.CreativeAutomationSpec != nil {
		out["creative_automation_spec"] = *params.CreativeAutomationSpec
	}
	if params.DisplaySequence != nil {
		out["display_sequence"] = *params.DisplaySequence
	}
	if params.DraftAdgroupID != nil {
		out["draft_adgroup_id"] = *params.DraftAdgroupID
	}
	if params.EngagementAudience != nil {
		out["engagement_audience"] = *params.EngagementAudience
	}
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	if params.IncludeDemolinkHashes != nil {
		out["include_demolink_hashes"] = *params.IncludeDemolinkHashes
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Priority != nil {
		out["priority"] = *params.Priority
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.TrackingSpecs != nil {
		out["tracking_specs"] = *params.TrackingSpecs
	}
	return out
}

func UpdateAdBatchCall(id string, params UpdateAdParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateAdBatchRequest(id string, params UpdateAdParams, options ...core.BatchOption) *core.BatchRequest[objects.Ad] {
	return core.NewBatchRequest[objects.Ad](UpdateAdBatchCall(id, params, options...))
}

func DecodeUpdateAdBatchResponse(response *core.BatchResponse) (*objects.Ad, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Ad
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateAdWithResponse(ctx context.Context, client *core.Client, id string, params UpdateAdParams) (*objects.Ad, *core.Response, error) {
	var out objects.Ad
	call := UpdateAdBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateAd(ctx context.Context, client *core.Client, id string, params UpdateAdParams) (*objects.Ad, error) {
	out, _, err := UpdateAdWithResponse(ctx, client, id, params)
	return out, err
}
