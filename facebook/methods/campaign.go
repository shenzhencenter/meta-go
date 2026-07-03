package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
)

type GetCampaignAdStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetCampaignAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetCampaignAdStudiesBatchCall(id string, params GetCampaignAdStudiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ad_studies"), params.ToParams(), options...)
}

func NewGetCampaignAdStudiesBatchRequest(id string, params GetCampaignAdStudiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdStudy]] {
	return core.NewBatchRequest[core.Cursor[objects.AdStudy]](GetCampaignAdStudiesBatchCall(id, params, options...))
}

func DecodeGetCampaignAdStudiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdStudy], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdStudy]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCampaignAdStudiesWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignAdStudiesParams) (*core.Cursor[objects.AdStudy], *core.Response, error) {
	var out core.Cursor[objects.AdStudy]
	call := GetCampaignAdStudiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaignAdStudies(ctx context.Context, client *core.Client, id string, params GetCampaignAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	out, _, err := GetCampaignAdStudiesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateCampaignAdlabelsParams struct {
	Adlabels         []map[string]interface{}                                  `facebook:"adlabels"`
	ExecutionOptions *[]enums.AdcampaigngroupadlabelsExecutionOptionsEnumParam `facebook:"execution_options"`
	Extra            core.Params                                               `facebook:"-"`
}

func (params CreateCampaignAdlabelsParams) ToParams() core.Params {
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

func CreateCampaignAdlabelsBatchCall(id string, params CreateCampaignAdlabelsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "adlabels"), params.ToParams(), options...)
}

func NewCreateCampaignAdlabelsBatchRequest(id string, params CreateCampaignAdlabelsParams, options ...core.BatchOption) *core.BatchRequest[objects.Campaign] {
	return core.NewBatchRequest[objects.Campaign](CreateCampaignAdlabelsBatchCall(id, params, options...))
}

func DecodeCreateCampaignAdlabelsBatchResponse(response *core.BatchResponse) (*objects.Campaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Campaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCampaignAdlabelsWithResponse(ctx context.Context, client *core.Client, id string, params CreateCampaignAdlabelsParams) (*objects.Campaign, *core.Response, error) {
	var out objects.Campaign
	call := CreateCampaignAdlabelsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateCampaignAdlabels(ctx context.Context, client *core.Client, id string, params CreateCampaignAdlabelsParams) (*objects.Campaign, error) {
	out, _, err := CreateCampaignAdlabelsWithResponse(ctx, client, id, params)
	return out, err
}

type GetCampaignAdrulesGovernedParams struct {
	PassEvaluation *bool       `facebook:"pass_evaluation"`
	Extra          core.Params `facebook:"-"`
}

func (params GetCampaignAdrulesGovernedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.PassEvaluation != nil {
		out["pass_evaluation"] = *params.PassEvaluation
	}
	return out
}

func GetCampaignAdrulesGovernedBatchCall(id string, params GetCampaignAdrulesGovernedParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adrules_governed"), params.ToParams(), options...)
}

func NewGetCampaignAdrulesGovernedBatchRequest(id string, params GetCampaignAdrulesGovernedParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdRule]] {
	return core.NewBatchRequest[core.Cursor[objects.AdRule]](GetCampaignAdrulesGovernedBatchCall(id, params, options...))
}

func DecodeGetCampaignAdrulesGovernedBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdRule], error) {
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

func GetCampaignAdrulesGovernedWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignAdrulesGovernedParams) (*core.Cursor[objects.AdRule], *core.Response, error) {
	var out core.Cursor[objects.AdRule]
	call := GetCampaignAdrulesGovernedBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaignAdrulesGoverned(ctx context.Context, client *core.Client, id string, params GetCampaignAdrulesGovernedParams) (*core.Cursor[objects.AdRule], error) {
	out, _, err := GetCampaignAdrulesGovernedWithResponse(ctx, client, id, params)
	return out, err
}

type GetCampaignAdsParams struct {
	DatePreset      *enums.AdcampaigngroupadsDatePresetEnumParam `facebook:"date_preset"`
	EffectiveStatus *[]string                                    `facebook:"effective_status"`
	TimeRange       *map[string]interface{}                      `facebook:"time_range"`
	UpdatedSince    *int                                         `facebook:"updated_since"`
	Extra           core.Params                                  `facebook:"-"`
}

func (params GetCampaignAdsParams) ToParams() core.Params {
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

func GetCampaignAdsBatchCall(id string, params GetCampaignAdsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "ads"), params.ToParams(), options...)
}

func NewGetCampaignAdsBatchRequest(id string, params GetCampaignAdsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Ad]] {
	return core.NewBatchRequest[core.Cursor[objects.Ad]](GetCampaignAdsBatchCall(id, params, options...))
}

func DecodeGetCampaignAdsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Ad], error) {
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

func GetCampaignAdsWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignAdsParams) (*core.Cursor[objects.Ad], *core.Response, error) {
	var out core.Cursor[objects.Ad]
	call := GetCampaignAdsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaignAds(ctx context.Context, client *core.Client, id string, params GetCampaignAdsParams) (*core.Cursor[objects.Ad], error) {
	out, _, err := GetCampaignAdsWithResponse(ctx, client, id, params)
	return out, err
}

type GetCampaignAdsetsParams struct {
	DatePreset      *enums.AdcampaigngroupadsetsDatePresetEnumParam        `facebook:"date_preset"`
	EffectiveStatus *[]enums.AdcampaigngroupadsetsEffectiveStatusEnumParam `facebook:"effective_status"`
	IsCompleted     *bool                                                  `facebook:"is_completed"`
	TimeRange       *map[string]interface{}                                `facebook:"time_range"`
	Extra           core.Params                                            `facebook:"-"`
}

func (params GetCampaignAdsetsParams) ToParams() core.Params {
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
	if params.IsCompleted != nil {
		out["is_completed"] = *params.IsCompleted
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	return out
}

func GetCampaignAdsetsBatchCall(id string, params GetCampaignAdsetsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "adsets"), params.ToParams(), options...)
}

func NewGetCampaignAdsetsBatchRequest(id string, params GetCampaignAdsetsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdSet]] {
	return core.NewBatchRequest[core.Cursor[objects.AdSet]](GetCampaignAdsetsBatchCall(id, params, options...))
}

func DecodeGetCampaignAdsetsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdSet], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.AdSet]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCampaignAdsetsWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignAdsetsParams) (*core.Cursor[objects.AdSet], *core.Response, error) {
	var out core.Cursor[objects.AdSet]
	call := GetCampaignAdsetsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaignAdsets(ctx context.Context, client *core.Client, id string, params GetCampaignAdsetsParams) (*core.Cursor[objects.AdSet], error) {
	out, _, err := GetCampaignAdsetsWithResponse(ctx, client, id, params)
	return out, err
}

type GetCampaignBudgetSchedulesParams struct {
	TimeStart *core.Time  `facebook:"time_start"`
	TimeStop  *core.Time  `facebook:"time_stop"`
	Extra     core.Params `facebook:"-"`
}

func (params GetCampaignBudgetSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.TimeStart != nil {
		out["time_start"] = *params.TimeStart
	}
	if params.TimeStop != nil {
		out["time_stop"] = *params.TimeStop
	}
	return out
}

func GetCampaignBudgetSchedulesBatchCall(id string, params GetCampaignBudgetSchedulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "budget_schedules"), params.ToParams(), options...)
}

func NewGetCampaignBudgetSchedulesBatchRequest(id string, params GetCampaignBudgetSchedulesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.HighDemandPeriod]] {
	return core.NewBatchRequest[core.Cursor[objects.HighDemandPeriod]](GetCampaignBudgetSchedulesBatchCall(id, params, options...))
}

func DecodeGetCampaignBudgetSchedulesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.HighDemandPeriod], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.HighDemandPeriod]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCampaignBudgetSchedulesWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignBudgetSchedulesParams) (*core.Cursor[objects.HighDemandPeriod], *core.Response, error) {
	var out core.Cursor[objects.HighDemandPeriod]
	call := GetCampaignBudgetSchedulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaignBudgetSchedules(ctx context.Context, client *core.Client, id string, params GetCampaignBudgetSchedulesParams) (*core.Cursor[objects.HighDemandPeriod], error) {
	out, _, err := GetCampaignBudgetSchedulesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateCampaignBudgetSchedulesParams struct {
	BudgetValue     uint64                                                       `facebook:"budget_value"`
	BudgetValueType enums.AdcampaigngroupbudgetSchedulesBudgetValueTypeEnumParam `facebook:"budget_value_type"`
	TimeEnd         uint64                                                       `facebook:"time_end"`
	TimeStart       uint64                                                       `facebook:"time_start"`
	Extra           core.Params                                                  `facebook:"-"`
}

func (params CreateCampaignBudgetSchedulesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	out["budget_value"] = params.BudgetValue
	out["budget_value_type"] = params.BudgetValueType
	out["time_end"] = params.TimeEnd
	out["time_start"] = params.TimeStart
	return out
}

func CreateCampaignBudgetSchedulesBatchCall(id string, params CreateCampaignBudgetSchedulesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "budget_schedules"), params.ToParams(), options...)
}

func NewCreateCampaignBudgetSchedulesBatchRequest(id string, params CreateCampaignBudgetSchedulesParams, options ...core.BatchOption) *core.BatchRequest[objects.HighDemandPeriod] {
	return core.NewBatchRequest[objects.HighDemandPeriod](CreateCampaignBudgetSchedulesBatchCall(id, params, options...))
}

func DecodeCreateCampaignBudgetSchedulesBatchResponse(response *core.BatchResponse) (*objects.HighDemandPeriod, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.HighDemandPeriod
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCampaignBudgetSchedulesWithResponse(ctx context.Context, client *core.Client, id string, params CreateCampaignBudgetSchedulesParams) (*objects.HighDemandPeriod, *core.Response, error) {
	var out objects.HighDemandPeriod
	call := CreateCampaignBudgetSchedulesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateCampaignBudgetSchedules(ctx context.Context, client *core.Client, id string, params CreateCampaignBudgetSchedulesParams) (*objects.HighDemandPeriod, error) {
	out, _, err := CreateCampaignBudgetSchedulesWithResponse(ctx, client, id, params)
	return out, err
}

type GetCampaignCopiesParams struct {
	DatePreset      *enums.AdcampaigngroupcopiesDatePresetEnumParam        `facebook:"date_preset"`
	EffectiveStatus *[]enums.AdcampaigngroupcopiesEffectiveStatusEnumParam `facebook:"effective_status"`
	IsCompleted     *bool                                                  `facebook:"is_completed"`
	TimeRange       *map[string]interface{}                                `facebook:"time_range"`
	Extra           core.Params                                            `facebook:"-"`
}

func (params GetCampaignCopiesParams) ToParams() core.Params {
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
	if params.IsCompleted != nil {
		out["is_completed"] = *params.IsCompleted
	}
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	return out
}

func GetCampaignCopiesBatchCall(id string, params GetCampaignCopiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "copies"), params.ToParams(), options...)
}

func NewGetCampaignCopiesBatchRequest(id string, params GetCampaignCopiesParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.Campaign]] {
	return core.NewBatchRequest[core.Cursor[objects.Campaign]](GetCampaignCopiesBatchCall(id, params, options...))
}

func DecodeGetCampaignCopiesBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.Campaign], error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out core.Cursor[objects.Campaign]
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCampaignCopiesWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignCopiesParams) (*core.Cursor[objects.Campaign], *core.Response, error) {
	var out core.Cursor[objects.Campaign]
	call := GetCampaignCopiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaignCopies(ctx context.Context, client *core.Client, id string, params GetCampaignCopiesParams) (*core.Cursor[objects.Campaign], error) {
	out, _, err := GetCampaignCopiesWithResponse(ctx, client, id, params)
	return out, err
}

type CreateCampaignCopiesParams struct {
	DeepCopy               *bool                                             `facebook:"deep_copy"`
	EndTime                *core.Time                                        `facebook:"end_time"`
	MigrateToAdvantagePlus *bool                                             `facebook:"migrate_to_advantage_plus"`
	ParameterOverrides     *map[string]interface{}                           `facebook:"parameter_overrides"`
	RenameOptions          *map[string]interface{}                           `facebook:"rename_options"`
	StartTime              *core.Time                                        `facebook:"start_time"`
	StatusOption           *enums.AdcampaigngroupcopiesStatusOptionEnumParam `facebook:"status_option"`
	Extra                  core.Params                                       `facebook:"-"`
}

func (params CreateCampaignCopiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.DeepCopy != nil {
		out["deep_copy"] = *params.DeepCopy
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.MigrateToAdvantagePlus != nil {
		out["migrate_to_advantage_plus"] = *params.MigrateToAdvantagePlus
	}
	if params.ParameterOverrides != nil {
		out["parameter_overrides"] = *params.ParameterOverrides
	}
	if params.RenameOptions != nil {
		out["rename_options"] = *params.RenameOptions
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.StatusOption != nil {
		out["status_option"] = *params.StatusOption
	}
	return out
}

func CreateCampaignCopiesBatchCall(id string, params CreateCampaignCopiesParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "copies"), params.ToParams(), options...)
}

func NewCreateCampaignCopiesBatchRequest(id string, params CreateCampaignCopiesParams, options ...core.BatchOption) *core.BatchRequest[objects.Campaign] {
	return core.NewBatchRequest[objects.Campaign](CreateCampaignCopiesBatchCall(id, params, options...))
}

func DecodeCreateCampaignCopiesBatchResponse(response *core.BatchResponse) (*objects.Campaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Campaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func CreateCampaignCopiesWithResponse(ctx context.Context, client *core.Client, id string, params CreateCampaignCopiesParams) (*objects.Campaign, *core.Response, error) {
	var out objects.Campaign
	call := CreateCampaignCopiesBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateCampaignCopies(ctx context.Context, client *core.Client, id string, params CreateCampaignCopiesParams) (*objects.Campaign, error) {
	out, _, err := CreateCampaignCopiesWithResponse(ctx, client, id, params)
	return out, err
}

type GetCampaignInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdcampaigngroupinsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdcampaigngroupinsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdcampaigngroupinsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdcampaigngroupinsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdcampaigngroupinsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                             `facebook:"default_summary"`
	ExportColumns                *[]string                                                         `facebook:"export_columns"`
	ExportFormat                 *string                                                           `facebook:"export_format"`
	ExportName                   *string                                                           `facebook:"export_name"`
	Fields                       *[]string                                                         `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                         `facebook:"filtering"`
	GraphCache                   *bool                                                             `facebook:"graph_cache"`
	Level                        *enums.AdcampaigngroupinsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                              `facebook:"limit"`
	ProductIDLimit               *int                                                              `facebook:"product_id_limit"`
	Sort                         *[]string                                                         `facebook:"sort"`
	Summary                      *[]string                                                         `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdcampaigngroupinsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                           `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                           `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                         `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                             `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                             `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                                       `facebook:"-"`
}

func (params GetCampaignInsightsParams) ToParams() core.Params {
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

func GetCampaignInsightsBatchCall(id string, params GetCampaignInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewGetCampaignInsightsBatchRequest(id string, params GetCampaignInsightsParams, options ...core.BatchOption) *core.BatchRequest[core.Cursor[objects.AdsInsights]] {
	return core.NewBatchRequest[core.Cursor[objects.AdsInsights]](GetCampaignInsightsBatchCall(id, params, options...))
}

func DecodeGetCampaignInsightsBatchResponse(response *core.BatchResponse) (*core.Cursor[objects.AdsInsights], error) {
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

func GetCampaignInsightsWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignInsightsParams) (*core.Cursor[objects.AdsInsights], *core.Response, error) {
	var out core.Cursor[objects.AdsInsights]
	call := GetCampaignInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaignInsights(ctx context.Context, client *core.Client, id string, params GetCampaignInsightsParams) (*core.Cursor[objects.AdsInsights], error) {
	out, _, err := GetCampaignInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type CreateCampaignInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdcampaigngroupinsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdcampaigngroupinsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdcampaigngroupinsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdcampaigngroupinsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdcampaigngroupinsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                             `facebook:"default_summary"`
	ExportColumns                *[]string                                                         `facebook:"export_columns"`
	ExportFormat                 *string                                                           `facebook:"export_format"`
	ExportName                   *string                                                           `facebook:"export_name"`
	Fields                       *[]string                                                         `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                         `facebook:"filtering"`
	GraphCache                   *bool                                                             `facebook:"graph_cache"`
	Level                        *enums.AdcampaigngroupinsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                              `facebook:"limit"`
	ProductIDLimit               *int                                                              `facebook:"product_id_limit"`
	Sort                         *[]string                                                         `facebook:"sort"`
	Summary                      *[]string                                                         `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdcampaigngroupinsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                           `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                           `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                         `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                             `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                             `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                                       `facebook:"-"`
}

func (params CreateCampaignInsightsParams) ToParams() core.Params {
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

func CreateCampaignInsightsBatchCall(id string, params CreateCampaignInsightsParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id, "insights"), params.ToParams(), options...)
}

func NewCreateCampaignInsightsBatchRequest(id string, params CreateCampaignInsightsParams, options ...core.BatchOption) *core.BatchRequest[objects.AdReportRun] {
	return core.NewBatchRequest[objects.AdReportRun](CreateCampaignInsightsBatchCall(id, params, options...))
}

func DecodeCreateCampaignInsightsBatchResponse(response *core.BatchResponse) (*objects.AdReportRun, error) {
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

func CreateCampaignInsightsWithResponse(ctx context.Context, client *core.Client, id string, params CreateCampaignInsightsParams) (*objects.AdReportRun, *core.Response, error) {
	var out objects.AdReportRun
	call := CreateCampaignInsightsBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func CreateCampaignInsights(ctx context.Context, client *core.Client, id string, params CreateCampaignInsightsParams) (*objects.AdReportRun, error) {
	out, _, err := CreateCampaignInsightsWithResponse(ctx, client, id, params)
	return out, err
}

type DeleteCampaignParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteCampaignParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteCampaignBatchCall(id string, params DeleteCampaignParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodDelete, core.GraphPath(id), params.ToParams(), options...)
}

func NewDeleteCampaignBatchRequest(id string, params DeleteCampaignParams, options ...core.BatchOption) *core.BatchRequest[map[string]interface{}] {
	return core.NewBatchRequest[map[string]interface{}](DeleteCampaignBatchCall(id, params, options...))
}

func DecodeDeleteCampaignBatchResponse(response *core.BatchResponse) (*map[string]interface{}, error) {
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

func DeleteCampaignWithResponse(ctx context.Context, client *core.Client, id string, params DeleteCampaignParams) (*map[string]interface{}, *core.Response, error) {
	var out map[string]interface{}
	call := DeleteCampaignBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func DeleteCampaign(ctx context.Context, client *core.Client, id string, params DeleteCampaignParams) (*map[string]interface{}, error) {
	out, _, err := DeleteCampaignWithResponse(ctx, client, id, params)
	return out, err
}

type GetCampaignParams struct {
	AmCallTags  *map[string]interface{}          `facebook:"am_call_tags"`
	DatePreset  *enums.AdcampaigngroupDatePreset `facebook:"date_preset"`
	FromAdtable *bool                            `facebook:"from_adtable"`
	TimeRange   *map[string]interface{}          `facebook:"time_range"`
	Extra       core.Params                      `facebook:"-"`
}

func (params GetCampaignParams) ToParams() core.Params {
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
	if params.TimeRange != nil {
		out["time_range"] = *params.TimeRange
	}
	return out
}

func GetCampaignBatchCall(id string, params GetCampaignParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodGet, core.GraphPath(id), params.ToParams(), options...)
}

func NewGetCampaignBatchRequest(id string, params GetCampaignParams, options ...core.BatchOption) *core.BatchRequest[objects.Campaign] {
	return core.NewBatchRequest[objects.Campaign](GetCampaignBatchCall(id, params, options...))
}

func DecodeGetCampaignBatchResponse(response *core.BatchResponse) (*objects.Campaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Campaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func GetCampaignWithResponse(ctx context.Context, client *core.Client, id string, params GetCampaignParams) (*objects.Campaign, *core.Response, error) {
	var out objects.Campaign
	call := GetCampaignBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func GetCampaign(ctx context.Context, client *core.Client, id string, params GetCampaignParams) (*objects.Campaign, error) {
	out, _, err := GetCampaignWithResponse(ctx, client, id, params)
	return out, err
}

type UpdateCampaignParams struct {
	Adlabels                    *[]map[string]interface{}                        `facebook:"adlabels"`
	AdsetBidAmounts             *map[string]interface{}                          `facebook:"adset_bid_amounts"`
	AdsetBudgets                *[]map[string]interface{}                        `facebook:"adset_budgets"`
	BidStrategy                 *enums.AdcampaigngroupBidStrategy                `facebook:"bid_strategy"`
	BudgetRebalanceFlag         *bool                                            `facebook:"budget_rebalance_flag"`
	BudgetScheduleSpecs         *[]map[string]interface{}                        `facebook:"budget_schedule_specs"`
	DailyBudget                 *uint64                                          `facebook:"daily_budget"`
	ExecutionOptions            *[]enums.AdcampaigngroupExecutionOptions         `facebook:"execution_options"`
	FrequencyControlSpecs       *[]map[string]interface{}                        `facebook:"frequency_control_specs"`
	IsAdsetBudgetSharingEnabled *bool                                            `facebook:"is_adset_budget_sharing_enabled"`
	IsBudgetScheduleEnabled     *bool                                            `facebook:"is_budget_schedule_enabled"`
	IsDirectSendCampaign        *bool                                            `facebook:"is_direct_send_campaign"`
	IsMessageCampaign           *bool                                            `facebook:"is_message_campaign"`
	IsMetaMomentMakerEnabled    *bool                                            `facebook:"is_meta_moment_maker_enabled"`
	IsReelsTrendingAdsEnabled   *bool                                            `facebook:"is_reels_trending_ads_enabled"`
	IsSkadnetworkAttribution    *bool                                            `facebook:"is_skadnetwork_attribution"`
	IterativeSplitTestConfigs   *[]map[string]interface{}                        `facebook:"iterative_split_test_configs"`
	LifetimeBudget              *uint64                                          `facebook:"lifetime_budget"`
	MigrateToAdvantagePlus      *bool                                            `facebook:"migrate_to_advantage_plus"`
	Name                        *string                                          `facebook:"name"`
	Objective                   *enums.AdcampaigngroupObjective                  `facebook:"objective"`
	PacingType                  *[]string                                        `facebook:"pacing_type"`
	PromotedObject              *map[string]interface{}                          `facebook:"promoted_object"`
	SmartPromotionType          *enums.AdcampaigngroupSmartPromotionType         `facebook:"smart_promotion_type"`
	SpecialAdCategories         *[]enums.AdcampaigngroupSpecialAdCategories      `facebook:"special_ad_categories"`
	SpecialAdCategory           *enums.AdcampaigngroupSpecialAdCategory          `facebook:"special_ad_category"`
	SpecialAdCategoryCountry    *[]enums.AdcampaigngroupSpecialAdCategoryCountry `facebook:"special_ad_category_country"`
	SpendCap                    *uint64                                          `facebook:"spend_cap"`
	StartTime                   *core.Time                                       `facebook:"start_time"`
	Status                      *enums.AdcampaigngroupStatus                     `facebook:"status"`
	StopTime                    *core.Time                                       `facebook:"stop_time"`
	Extra                       core.Params                                      `facebook:"-"`
}

func (params UpdateCampaignParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.AdsetBidAmounts != nil {
		out["adset_bid_amounts"] = *params.AdsetBidAmounts
	}
	if params.AdsetBudgets != nil {
		out["adset_budgets"] = *params.AdsetBudgets
	}
	if params.BidStrategy != nil {
		out["bid_strategy"] = *params.BidStrategy
	}
	if params.BudgetRebalanceFlag != nil {
		out["budget_rebalance_flag"] = *params.BudgetRebalanceFlag
	}
	if params.BudgetScheduleSpecs != nil {
		out["budget_schedule_specs"] = *params.BudgetScheduleSpecs
	}
	if params.DailyBudget != nil {
		out["daily_budget"] = *params.DailyBudget
	}
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	if params.FrequencyControlSpecs != nil {
		out["frequency_control_specs"] = *params.FrequencyControlSpecs
	}
	if params.IsAdsetBudgetSharingEnabled != nil {
		out["is_adset_budget_sharing_enabled"] = *params.IsAdsetBudgetSharingEnabled
	}
	if params.IsBudgetScheduleEnabled != nil {
		out["is_budget_schedule_enabled"] = *params.IsBudgetScheduleEnabled
	}
	if params.IsDirectSendCampaign != nil {
		out["is_direct_send_campaign"] = *params.IsDirectSendCampaign
	}
	if params.IsMessageCampaign != nil {
		out["is_message_campaign"] = *params.IsMessageCampaign
	}
	if params.IsMetaMomentMakerEnabled != nil {
		out["is_meta_moment_maker_enabled"] = *params.IsMetaMomentMakerEnabled
	}
	if params.IsReelsTrendingAdsEnabled != nil {
		out["is_reels_trending_ads_enabled"] = *params.IsReelsTrendingAdsEnabled
	}
	if params.IsSkadnetworkAttribution != nil {
		out["is_skadnetwork_attribution"] = *params.IsSkadnetworkAttribution
	}
	if params.IterativeSplitTestConfigs != nil {
		out["iterative_split_test_configs"] = *params.IterativeSplitTestConfigs
	}
	if params.LifetimeBudget != nil {
		out["lifetime_budget"] = *params.LifetimeBudget
	}
	if params.MigrateToAdvantagePlus != nil {
		out["migrate_to_advantage_plus"] = *params.MigrateToAdvantagePlus
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.Objective != nil {
		out["objective"] = *params.Objective
	}
	if params.PacingType != nil {
		out["pacing_type"] = *params.PacingType
	}
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	if params.SmartPromotionType != nil {
		out["smart_promotion_type"] = *params.SmartPromotionType
	}
	if params.SpecialAdCategories != nil {
		out["special_ad_categories"] = *params.SpecialAdCategories
	}
	if params.SpecialAdCategory != nil {
		out["special_ad_category"] = *params.SpecialAdCategory
	}
	if params.SpecialAdCategoryCountry != nil {
		out["special_ad_category_country"] = *params.SpecialAdCategoryCountry
	}
	if params.SpendCap != nil {
		out["spend_cap"] = *params.SpendCap
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.StopTime != nil {
		out["stop_time"] = *params.StopTime
	}
	return out
}

func UpdateCampaignBatchCall(id string, params UpdateCampaignParams, options ...core.BatchOption) core.BatchCall {
	return core.NewBatchCall(http.MethodPost, core.GraphPath(id), params.ToParams(), options...)
}

func NewUpdateCampaignBatchRequest(id string, params UpdateCampaignParams, options ...core.BatchOption) *core.BatchRequest[objects.Campaign] {
	return core.NewBatchRequest[objects.Campaign](UpdateCampaignBatchCall(id, params, options...))
}

func DecodeUpdateCampaignBatchResponse(response *core.BatchResponse) (*objects.Campaign, error) {
	if response == nil {
		return nil, nil
	}
	if err := response.Err(); err != nil {
		return nil, err
	}
	var out objects.Campaign
	if err := response.Decode(&out); err != nil {
		return nil, err
	}
	return &out, nil
}

func UpdateCampaignWithResponse(ctx context.Context, client *core.Client, id string, params UpdateCampaignParams) (*objects.Campaign, *core.Response, error) {
	var out objects.Campaign
	call := UpdateCampaignBatchCall(id, params)
	response, err := client.RequestWithResponse(ctx, call.Method, call.RelativeURL, call.Params, &out)
	return &out, response, err
}

func UpdateCampaign(ctx context.Context, client *core.Client, id string, params UpdateCampaignParams) (*objects.Campaign, error) {
	out, _, err := UpdateCampaignWithResponse(ctx, client, id, params)
	return out, err
}
