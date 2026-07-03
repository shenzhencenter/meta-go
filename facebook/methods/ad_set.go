package methods

import (
	"context"
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"github.com/shenzhencenter/meta-go/facebook/objects"
	"net/http"
	"time"
)

type GetAdSetActivitiesParams struct {
	After      *string                                      `facebook:"after"`
	BusinessID *core.ID                                     `facebook:"business_id"`
	Category   *enums.AdcampaignactivitiesCategoryEnumParam `facebook:"category"`
	Limit      *int                                         `facebook:"limit"`
	Since      *time.Time                                   `facebook:"since"`
	UID        *core.ID                                     `facebook:"uid"`
	Until      *time.Time                                   `facebook:"until"`
	Extra      core.Params                                  `facebook:"-"`
}

func (params GetAdSetActivitiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.After != nil {
		out["after"] = *params.After
	}
	if params.BusinessID != nil {
		out["business_id"] = *params.BusinessID
	}
	if params.Category != nil {
		out["category"] = *params.Category
	}
	if params.Limit != nil {
		out["limit"] = *params.Limit
	}
	if params.Since != nil {
		out["since"] = *params.Since
	}
	if params.UID != nil {
		out["uid"] = *params.UID
	}
	if params.Until != nil {
		out["until"] = *params.Until
	}
	return out
}

func GetAdSetActivities(ctx context.Context, client *core.Client, id string, params GetAdSetActivitiesParams) (*core.Cursor[objects.AdActivity], error) {
	var out core.Cursor[objects.AdActivity]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "activities"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetAdStudiesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSetAdStudiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSetAdStudies(ctx context.Context, client *core.Client, id string, params GetAdSetAdStudiesParams) (*core.Cursor[objects.AdStudy], error) {
	var out core.Cursor[objects.AdStudy]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ad_studies"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetAdcreativesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSetAdcreativesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSetAdcreatives(ctx context.Context, client *core.Client, id string, params GetAdSetAdcreativesParams) (*core.Cursor[objects.AdCreative], error) {
	var out core.Cursor[objects.AdCreative]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adcreatives"), params.ToParams(), &out)
	return &out, err
}

type DeleteAdSetAdlabelsParams struct {
	Adlabels         []map[string]interface{}                             `facebook:"adlabels"`
	ExecutionOptions *[]enums.AdcampaignadlabelsExecutionOptionsEnumParam `facebook:"execution_options"`
	Extra            core.Params                                          `facebook:"-"`
}

func (params DeleteAdSetAdlabelsParams) ToParams() core.Params {
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

func DeleteAdSetAdlabels(ctx context.Context, client *core.Client, id string, params DeleteAdSetAdlabelsParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id, "adlabels"), params.ToParams(), &out)
	return &out, err
}

type CreateAdSetAdlabelsParams struct {
	Adlabels         []map[string]interface{}                             `facebook:"adlabels"`
	ExecutionOptions *[]enums.AdcampaignadlabelsExecutionOptionsEnumParam `facebook:"execution_options"`
	Extra            core.Params                                          `facebook:"-"`
}

func (params CreateAdSetAdlabelsParams) ToParams() core.Params {
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

func CreateAdSetAdlabels(ctx context.Context, client *core.Client, id string, params CreateAdSetAdlabelsParams) (*objects.AdSet, error) {
	var out objects.AdSet
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "adlabels"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetAdrulesGovernedParams struct {
	PassEvaluation *bool       `facebook:"pass_evaluation"`
	Extra          core.Params `facebook:"-"`
}

func (params GetAdSetAdrulesGovernedParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.PassEvaluation != nil {
		out["pass_evaluation"] = *params.PassEvaluation
	}
	return out
}

func GetAdSetAdrulesGoverned(ctx context.Context, client *core.Client, id string, params GetAdSetAdrulesGovernedParams) (*core.Cursor[objects.AdRule], error) {
	var out core.Cursor[objects.AdRule]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "adrules_governed"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetAdsParams struct {
	DatePreset      *enums.AdcampaignadsDatePresetEnumParam `facebook:"date_preset"`
	EffectiveStatus *[]string                               `facebook:"effective_status"`
	TimeRange       *map[string]interface{}                 `facebook:"time_range"`
	UpdatedSince    *int                                    `facebook:"updated_since"`
	Extra           core.Params                             `facebook:"-"`
}

func (params GetAdSetAdsParams) ToParams() core.Params {
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

func GetAdSetAds(ctx context.Context, client *core.Client, id string, params GetAdSetAdsParams) (*core.Cursor[objects.Ad], error) {
	var out core.Cursor[objects.Ad]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "ads"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetAsyncadrequestsParams struct {
	Statuses *[]enums.AdcampaignasyncadrequestsStatusesEnumParam `facebook:"statuses"`
	Extra    core.Params                                         `facebook:"-"`
}

func (params GetAdSetAsyncadrequestsParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.Statuses != nil {
		out["statuses"] = *params.Statuses
	}
	return out
}

func GetAdSetAsyncadrequests(ctx context.Context, client *core.Client, id string, params GetAdSetAsyncadrequestsParams) (*core.Cursor[objects.AdAsyncRequest], error) {
	var out core.Cursor[objects.AdAsyncRequest]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "asyncadrequests"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetBudgetSchedulesParams struct {
	TimeStart *time.Time  `facebook:"time_start"`
	TimeStop  *time.Time  `facebook:"time_stop"`
	Extra     core.Params `facebook:"-"`
}

func (params GetAdSetBudgetSchedulesParams) ToParams() core.Params {
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

func GetAdSetBudgetSchedules(ctx context.Context, client *core.Client, id string, params GetAdSetBudgetSchedulesParams) (*core.Cursor[objects.HighDemandPeriod], error) {
	var out core.Cursor[objects.HighDemandPeriod]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "budget_schedules"), params.ToParams(), &out)
	return &out, err
}

type CreateAdSetBudgetSchedulesParams struct {
	BudgetValue     uint64                                                  `facebook:"budget_value"`
	BudgetValueType enums.AdcampaignbudgetSchedulesBudgetValueTypeEnumParam `facebook:"budget_value_type"`
	TimeEnd         uint64                                                  `facebook:"time_end"`
	TimeStart       uint64                                                  `facebook:"time_start"`
	Extra           core.Params                                             `facebook:"-"`
}

func (params CreateAdSetBudgetSchedulesParams) ToParams() core.Params {
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

func CreateAdSetBudgetSchedules(ctx context.Context, client *core.Client, id string, params CreateAdSetBudgetSchedulesParams) (*objects.HighDemandPeriod, error) {
	var out objects.HighDemandPeriod
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "budget_schedules"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetCopiesParams struct {
	DatePreset      *enums.AdcampaigncopiesDatePresetEnumParam        `facebook:"date_preset"`
	EffectiveStatus *[]enums.AdcampaigncopiesEffectiveStatusEnumParam `facebook:"effective_status"`
	IsCompleted     *bool                                             `facebook:"is_completed"`
	TimeRange       *map[string]interface{}                           `facebook:"time_range"`
	Extra           core.Params                                       `facebook:"-"`
}

func (params GetAdSetCopiesParams) ToParams() core.Params {
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

func GetAdSetCopies(ctx context.Context, client *core.Client, id string, params GetAdSetCopiesParams) (*core.Cursor[objects.AdSet], error) {
	var out core.Cursor[objects.AdSet]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "copies"), params.ToParams(), &out)
	return &out, err
}

type CreateAdSetCopiesParams struct {
	CampaignID     *core.ID                                     `facebook:"campaign_id"`
	CreateDcoAdset *bool                                        `facebook:"create_dco_adset"`
	DeepCopy       *bool                                        `facebook:"deep_copy"`
	EndTime        *time.Time                                   `facebook:"end_time"`
	RenameOptions  *map[string]interface{}                      `facebook:"rename_options"`
	StartTime      *time.Time                                   `facebook:"start_time"`
	StatusOption   *enums.AdcampaigncopiesStatusOptionEnumParam `facebook:"status_option"`
	Extra          core.Params                                  `facebook:"-"`
}

func (params CreateAdSetCopiesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.CampaignID != nil {
		out["campaign_id"] = *params.CampaignID
	}
	if params.CreateDcoAdset != nil {
		out["create_dco_adset"] = *params.CreateDcoAdset
	}
	if params.DeepCopy != nil {
		out["deep_copy"] = *params.DeepCopy
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
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

func CreateAdSetCopies(ctx context.Context, client *core.Client, id string, params CreateAdSetCopiesParams) (*objects.AdSet, error) {
	var out objects.AdSet
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "copies"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetDeliveryEstimateParams struct {
	OptimizationGoal *enums.AdcampaigndeliveryEstimateOptimizationGoalEnumParam `facebook:"optimization_goal"`
	PromotedObject   *map[string]interface{}                                    `facebook:"promoted_object"`
	TargetingSpec    *objects.Targeting                                         `facebook:"targeting_spec"`
	Extra            core.Params                                                `facebook:"-"`
}

func (params GetAdSetDeliveryEstimateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.OptimizationGoal != nil {
		out["optimization_goal"] = *params.OptimizationGoal
	}
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	if params.TargetingSpec != nil {
		out["targeting_spec"] = *params.TargetingSpec
	}
	return out
}

func GetAdSetDeliveryEstimate(ctx context.Context, client *core.Client, id string, params GetAdSetDeliveryEstimateParams) (*core.Cursor[objects.AdCampaignDeliveryEstimate], error) {
	var out core.Cursor[objects.AdCampaignDeliveryEstimate]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "delivery_estimate"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdcampaigninsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdcampaigninsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdcampaigninsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdcampaigninsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdcampaigninsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                        `facebook:"default_summary"`
	ExportColumns                *[]string                                                    `facebook:"export_columns"`
	ExportFormat                 *string                                                      `facebook:"export_format"`
	ExportName                   *string                                                      `facebook:"export_name"`
	Fields                       *[]string                                                    `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                    `facebook:"filtering"`
	GraphCache                   *bool                                                        `facebook:"graph_cache"`
	Level                        *enums.AdcampaigninsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                         `facebook:"limit"`
	ProductIDLimit               *int                                                         `facebook:"product_id_limit"`
	Sort                         *[]string                                                    `facebook:"sort"`
	Summary                      *[]string                                                    `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdcampaigninsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                      `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                      `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                    `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                        `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                        `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                                  `facebook:"-"`
}

func (params GetAdSetInsightsParams) ToParams() core.Params {
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

func GetAdSetInsights(ctx context.Context, client *core.Client, id string, params GetAdSetInsightsParams) (*core.Cursor[objects.AdsInsights], error) {
	var out core.Cursor[objects.AdsInsights]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type CreateAdSetInsightsParams struct {
	ActionAttributionWindows     *[]enums.AdcampaigninsightsActionAttributionWindowsEnumParam `facebook:"action_attribution_windows"`
	ActionBreakdowns             *[]enums.AdcampaigninsightsActionBreakdownsEnumParam         `facebook:"action_breakdowns"`
	ActionReportTime             *enums.AdcampaigninsightsActionReportTimeEnumParam           `facebook:"action_report_time"`
	Breakdowns                   *[]enums.AdcampaigninsightsBreakdownsEnumParam               `facebook:"breakdowns"`
	DatePreset                   *enums.AdcampaigninsightsDatePresetEnumParam                 `facebook:"date_preset"`
	DefaultSummary               *bool                                                        `facebook:"default_summary"`
	ExportColumns                *[]string                                                    `facebook:"export_columns"`
	ExportFormat                 *string                                                      `facebook:"export_format"`
	ExportName                   *string                                                      `facebook:"export_name"`
	Fields                       *[]string                                                    `facebook:"fields"`
	Filtering                    *[]map[string]interface{}                                    `facebook:"filtering"`
	GraphCache                   *bool                                                        `facebook:"graph_cache"`
	Level                        *enums.AdcampaigninsightsLevelEnumParam                      `facebook:"level"`
	Limit                        *int                                                         `facebook:"limit"`
	ProductIDLimit               *int                                                         `facebook:"product_id_limit"`
	Sort                         *[]string                                                    `facebook:"sort"`
	Summary                      *[]string                                                    `facebook:"summary"`
	SummaryActionBreakdowns      *[]enums.AdcampaigninsightsSummaryActionBreakdownsEnumParam  `facebook:"summary_action_breakdowns"`
	TimeIncrement                *string                                                      `facebook:"time_increment"`
	TimeRange                    *map[string]interface{}                                      `facebook:"time_range"`
	TimeRanges                   *[]map[string]interface{}                                    `facebook:"time_ranges"`
	UseAccountAttributionSetting *bool                                                        `facebook:"use_account_attribution_setting"`
	UseUnifiedAttributionSetting *bool                                                        `facebook:"use_unified_attribution_setting"`
	Extra                        core.Params                                                  `facebook:"-"`
}

func (params CreateAdSetInsightsParams) ToParams() core.Params {
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

func CreateAdSetInsights(ctx context.Context, client *core.Client, id string, params CreateAdSetInsightsParams) (*objects.AdReportRun, error) {
	var out objects.AdReportRun
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id, "insights"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetMessageDeliveryEstimateParams struct {
	BidAmount            *uint64                                                           `facebook:"bid_amount"`
	DailyBudget          *uint64                                                           `facebook:"daily_budget"`
	IsDirectSendCampaign *bool                                                             `facebook:"is_direct_send_campaign"`
	LifetimeBudget       *uint64                                                           `facebook:"lifetime_budget"`
	LifetimeInDays       *uint64                                                           `facebook:"lifetime_in_days"`
	OptimizationGoal     *enums.AdcampaignmessageDeliveryEstimateOptimizationGoalEnumParam `facebook:"optimization_goal"`
	PacingType           *enums.AdcampaignmessageDeliveryEstimatePacingTypeEnumParam       `facebook:"pacing_type"`
	PromotedObject       *map[string]interface{}                                           `facebook:"promoted_object"`
	TargetingSpec        *objects.Targeting                                                `facebook:"targeting_spec"`
	Extra                core.Params                                                       `facebook:"-"`
}

func (params GetAdSetMessageDeliveryEstimateParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	if params.DailyBudget != nil {
		out["daily_budget"] = *params.DailyBudget
	}
	if params.IsDirectSendCampaign != nil {
		out["is_direct_send_campaign"] = *params.IsDirectSendCampaign
	}
	if params.LifetimeBudget != nil {
		out["lifetime_budget"] = *params.LifetimeBudget
	}
	if params.LifetimeInDays != nil {
		out["lifetime_in_days"] = *params.LifetimeInDays
	}
	if params.OptimizationGoal != nil {
		out["optimization_goal"] = *params.OptimizationGoal
	}
	if params.PacingType != nil {
		out["pacing_type"] = *params.PacingType
	}
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	if params.TargetingSpec != nil {
		out["targeting_spec"] = *params.TargetingSpec
	}
	return out
}

func GetAdSetMessageDeliveryEstimate(ctx context.Context, client *core.Client, id string, params GetAdSetMessageDeliveryEstimateParams) (*core.Cursor[objects.MessageDeliveryEstimate], error) {
	var out core.Cursor[objects.MessageDeliveryEstimate]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "message_delivery_estimate"), params.ToParams(), &out)
	return &out, err
}

type GetAdSetTargetingsentencelinesParams struct {
	Extra core.Params `facebook:"-"`
}

func (params GetAdSetTargetingsentencelinesParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func GetAdSetTargetingsentencelines(ctx context.Context, client *core.Client, id string, params GetAdSetTargetingsentencelinesParams) (*core.Cursor[objects.TargetingSentenceLine], error) {
	var out core.Cursor[objects.TargetingSentenceLine]
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id, "targetingsentencelines"), params.ToParams(), &out)
	return &out, err
}

type DeleteAdSetParams struct {
	Extra core.Params `facebook:"-"`
}

func (params DeleteAdSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	return out
}

func DeleteAdSet(ctx context.Context, client *core.Client, id string, params DeleteAdSetParams) (*map[string]interface{}, error) {
	var out map[string]interface{}
	err := client.Request(ctx, http.MethodDelete, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type GetAdSetParams struct {
	AmCallTags  *map[string]interface{}     `facebook:"am_call_tags"`
	DatePreset  *enums.AdcampaignDatePreset `facebook:"date_preset"`
	FromAdtable *bool                       `facebook:"from_adtable"`
	TimeRange   *map[string]interface{}     `facebook:"time_range"`
	Extra       core.Params                 `facebook:"-"`
}

func (params GetAdSetParams) ToParams() core.Params {
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

func GetAdSet(ctx context.Context, client *core.Client, id string, params GetAdSetParams) (*objects.AdSet, error) {
	var out objects.AdSet
	err := client.Request(ctx, http.MethodGet, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}

type UpdateAdSetParams struct {
	AccountID                                    *core.ID                                           `facebook:"account_id"`
	Adlabels                                     *[]map[string]interface{}                          `facebook:"adlabels"`
	AdsetSchedule                                *[]map[string]interface{}                          `facebook:"adset_schedule"`
	AttributionSpec                              *[]map[string]interface{}                          `facebook:"attribution_spec"`
	AutomaticManualState                         *enums.AdcampaignAutomaticManualState              `facebook:"automatic_manual_state"`
	BidAdjustments                               *map[string]interface{}                            `facebook:"bid_adjustments"`
	BidAmount                                    *int                                               `facebook:"bid_amount"`
	BidConstraints                               *map[string]map[string]interface{}                 `facebook:"bid_constraints"`
	BidStrategy                                  *enums.AdcampaignBidStrategy                       `facebook:"bid_strategy"`
	BillingEvent                                 *enums.AdcampaignBillingEvent                      `facebook:"billing_event"`
	BudgetScheduleSpecs                          *[]map[string]interface{}                          `facebook:"budget_schedule_specs"`
	CampaignAttribution                          *map[string]interface{}                            `facebook:"campaign_attribution"`
	CampaignSpec                                 *map[string]interface{}                            `facebook:"campaign_spec"`
	CostBiddingMode                              *enums.AdcampaignCostBiddingMode                   `facebook:"cost_bidding_mode"`
	CreativeSequence                             *[]string                                          `facebook:"creative_sequence"`
	CreativeSequenceRepetitionPattern            *enums.AdcampaignCreativeSequenceRepetitionPattern `facebook:"creative_sequence_repetition_pattern"`
	DailyBudget                                  *uint64                                            `facebook:"daily_budget"`
	DailyImps                                    *uint64                                            `facebook:"daily_imps"`
	DailyMinSpendTarget                          *uint64                                            `facebook:"daily_min_spend_target"`
	DailySpendCap                                *uint64                                            `facebook:"daily_spend_cap"`
	DateFormat                                   *string                                            `facebook:"date_format"`
	DestinationType                              *enums.AdcampaignDestinationType                   `facebook:"destination_type"`
	DsaBeneficiary                               *string                                            `facebook:"dsa_beneficiary"`
	DsaPayor                                     *string                                            `facebook:"dsa_payor"`
	EndTime                                      *time.Time                                         `facebook:"end_time"`
	ExecutionOptions                             *[]enums.AdcampaignExecutionOptions                `facebook:"execution_options"`
	ExistingCustomerBudgetPercentage             *uint64                                            `facebook:"existing_customer_budget_percentage"`
	FrequencyControlSpecs                        *[]map[string]interface{}                          `facebook:"frequency_control_specs"`
	FullFunnelExplorationMode                    *enums.AdcampaignFullFunnelExplorationMode         `facebook:"full_funnel_exploration_mode"`
	IsBaSkipDelayedEligible                      *bool                                              `facebook:"is_ba_skip_delayed_eligible"`
	IsBudgetScheduleEnabled                      *bool                                              `facebook:"is_budget_schedule_enabled"`
	IsDcFollowOptimized                          *bool                                              `facebook:"is_dc_follow_optimized"`
	IsIncrementalAttributionEnabled              *bool                                              `facebook:"is_incremental_attribution_enabled"`
	IsSacCfcaTermsCertified                      *bool                                              `facebook:"is_sac_cfca_terms_certified"`
	LifetimeBudget                               *uint64                                            `facebook:"lifetime_budget"`
	LifetimeImps                                 *uint64                                            `facebook:"lifetime_imps"`
	LifetimeMinSpendTarget                       *uint64                                            `facebook:"lifetime_min_spend_target"`
	LifetimeSpendCap                             *uint64                                            `facebook:"lifetime_spend_cap"`
	LiveVideoAdCampaignConfig                    *map[string]interface{}                            `facebook:"live_video_ad_campaign_config"`
	MaxBudgetSpendPercentage                     *uint64                                            `facebook:"max_budget_spend_percentage"`
	MetaMomentMakerSpec                          *map[string]interface{}                            `facebook:"meta_moment_maker_spec"`
	MinBudgetSpendPercentage                     *uint64                                            `facebook:"min_budget_spend_percentage"`
	MultiEventConversionAttributionWindowSeconds *uint64                                            `facebook:"multi_event_conversion_attribution_window_seconds"`
	MultiOptimizationGoalWeight                  *enums.AdcampaignMultiOptimizationGoalWeight       `facebook:"multi_optimization_goal_weight"`
	Name                                         *string                                            `facebook:"name"`
	OptimizationGoal                             *enums.AdcampaignOptimizationGoal                  `facebook:"optimization_goal"`
	OptimizationSubEvent                         *enums.AdcampaignOptimizationSubEvent              `facebook:"optimization_sub_event"`
	PacingType                                   *[]string                                          `facebook:"pacing_type"`
	PlacementSoftOptOut                          *map[string]interface{}                            `facebook:"placement_soft_opt_out"`
	PromotedObject                               *map[string]interface{}                            `facebook:"promoted_object"`
	RbPredictionID                               *core.ID                                           `facebook:"rb_prediction_id"`
	RegionalRegulatedCategories                  *[]enums.AdcampaignRegionalRegulatedCategories     `facebook:"regional_regulated_categories"`
	RegionalRegulationIdentities                 *map[string]interface{}                            `facebook:"regional_regulation_identities"`
	RelativeValue                                *float64                                           `facebook:"relative_value"`
	RfPredictionID                               *core.ID                                           `facebook:"rf_prediction_id"`
	StartTime                                    *time.Time                                         `facebook:"start_time"`
	Status                                       *enums.AdcampaignStatus                            `facebook:"status"`
	Targeting                                    *objects.Targeting                                 `facebook:"targeting"`
	TimeBasedAdRotationIDBlocks                  *[][]uint64                                        `facebook:"time_based_ad_rotation_id_blocks"`
	TimeBasedAdRotationIntervals                 *[]uint64                                          `facebook:"time_based_ad_rotation_intervals"`
	TimeStart                                    *time.Time                                         `facebook:"time_start"`
	TimeStop                                     *time.Time                                         `facebook:"time_stop"`
	TrendingTopicsSpec                           *map[string]interface{}                            `facebook:"trending_topics_spec"`
	TuneForCategory                              *enums.AdcampaignTuneForCategory                   `facebook:"tune_for_category"`
	ValueRuleSetID                               *core.ID                                           `facebook:"value_rule_set_id"`
	ValueRulesApplied                            *bool                                              `facebook:"value_rules_applied"`
	Extra                                        core.Params                                        `facebook:"-"`
}

func (params UpdateAdSetParams) ToParams() core.Params {
	out := core.Params{}
	for key, value := range params.Extra {
		out[key] = value
	}
	if params.AccountID != nil {
		out["account_id"] = *params.AccountID
	}
	if params.Adlabels != nil {
		out["adlabels"] = *params.Adlabels
	}
	if params.AdsetSchedule != nil {
		out["adset_schedule"] = *params.AdsetSchedule
	}
	if params.AttributionSpec != nil {
		out["attribution_spec"] = *params.AttributionSpec
	}
	if params.AutomaticManualState != nil {
		out["automatic_manual_state"] = *params.AutomaticManualState
	}
	if params.BidAdjustments != nil {
		out["bid_adjustments"] = *params.BidAdjustments
	}
	if params.BidAmount != nil {
		out["bid_amount"] = *params.BidAmount
	}
	if params.BidConstraints != nil {
		out["bid_constraints"] = *params.BidConstraints
	}
	if params.BidStrategy != nil {
		out["bid_strategy"] = *params.BidStrategy
	}
	if params.BillingEvent != nil {
		out["billing_event"] = *params.BillingEvent
	}
	if params.BudgetScheduleSpecs != nil {
		out["budget_schedule_specs"] = *params.BudgetScheduleSpecs
	}
	if params.CampaignAttribution != nil {
		out["campaign_attribution"] = *params.CampaignAttribution
	}
	if params.CampaignSpec != nil {
		out["campaign_spec"] = *params.CampaignSpec
	}
	if params.CostBiddingMode != nil {
		out["cost_bidding_mode"] = *params.CostBiddingMode
	}
	if params.CreativeSequence != nil {
		out["creative_sequence"] = *params.CreativeSequence
	}
	if params.CreativeSequenceRepetitionPattern != nil {
		out["creative_sequence_repetition_pattern"] = *params.CreativeSequenceRepetitionPattern
	}
	if params.DailyBudget != nil {
		out["daily_budget"] = *params.DailyBudget
	}
	if params.DailyImps != nil {
		out["daily_imps"] = *params.DailyImps
	}
	if params.DailyMinSpendTarget != nil {
		out["daily_min_spend_target"] = *params.DailyMinSpendTarget
	}
	if params.DailySpendCap != nil {
		out["daily_spend_cap"] = *params.DailySpendCap
	}
	if params.DateFormat != nil {
		out["date_format"] = *params.DateFormat
	}
	if params.DestinationType != nil {
		out["destination_type"] = *params.DestinationType
	}
	if params.DsaBeneficiary != nil {
		out["dsa_beneficiary"] = *params.DsaBeneficiary
	}
	if params.DsaPayor != nil {
		out["dsa_payor"] = *params.DsaPayor
	}
	if params.EndTime != nil {
		out["end_time"] = *params.EndTime
	}
	if params.ExecutionOptions != nil {
		out["execution_options"] = *params.ExecutionOptions
	}
	if params.ExistingCustomerBudgetPercentage != nil {
		out["existing_customer_budget_percentage"] = *params.ExistingCustomerBudgetPercentage
	}
	if params.FrequencyControlSpecs != nil {
		out["frequency_control_specs"] = *params.FrequencyControlSpecs
	}
	if params.FullFunnelExplorationMode != nil {
		out["full_funnel_exploration_mode"] = *params.FullFunnelExplorationMode
	}
	if params.IsBaSkipDelayedEligible != nil {
		out["is_ba_skip_delayed_eligible"] = *params.IsBaSkipDelayedEligible
	}
	if params.IsBudgetScheduleEnabled != nil {
		out["is_budget_schedule_enabled"] = *params.IsBudgetScheduleEnabled
	}
	if params.IsDcFollowOptimized != nil {
		out["is_dc_follow_optimized"] = *params.IsDcFollowOptimized
	}
	if params.IsIncrementalAttributionEnabled != nil {
		out["is_incremental_attribution_enabled"] = *params.IsIncrementalAttributionEnabled
	}
	if params.IsSacCfcaTermsCertified != nil {
		out["is_sac_cfca_terms_certified"] = *params.IsSacCfcaTermsCertified
	}
	if params.LifetimeBudget != nil {
		out["lifetime_budget"] = *params.LifetimeBudget
	}
	if params.LifetimeImps != nil {
		out["lifetime_imps"] = *params.LifetimeImps
	}
	if params.LifetimeMinSpendTarget != nil {
		out["lifetime_min_spend_target"] = *params.LifetimeMinSpendTarget
	}
	if params.LifetimeSpendCap != nil {
		out["lifetime_spend_cap"] = *params.LifetimeSpendCap
	}
	if params.LiveVideoAdCampaignConfig != nil {
		out["live_video_ad_campaign_config"] = *params.LiveVideoAdCampaignConfig
	}
	if params.MaxBudgetSpendPercentage != nil {
		out["max_budget_spend_percentage"] = *params.MaxBudgetSpendPercentage
	}
	if params.MetaMomentMakerSpec != nil {
		out["meta_moment_maker_spec"] = *params.MetaMomentMakerSpec
	}
	if params.MinBudgetSpendPercentage != nil {
		out["min_budget_spend_percentage"] = *params.MinBudgetSpendPercentage
	}
	if params.MultiEventConversionAttributionWindowSeconds != nil {
		out["multi_event_conversion_attribution_window_seconds"] = *params.MultiEventConversionAttributionWindowSeconds
	}
	if params.MultiOptimizationGoalWeight != nil {
		out["multi_optimization_goal_weight"] = *params.MultiOptimizationGoalWeight
	}
	if params.Name != nil {
		out["name"] = *params.Name
	}
	if params.OptimizationGoal != nil {
		out["optimization_goal"] = *params.OptimizationGoal
	}
	if params.OptimizationSubEvent != nil {
		out["optimization_sub_event"] = *params.OptimizationSubEvent
	}
	if params.PacingType != nil {
		out["pacing_type"] = *params.PacingType
	}
	if params.PlacementSoftOptOut != nil {
		out["placement_soft_opt_out"] = *params.PlacementSoftOptOut
	}
	if params.PromotedObject != nil {
		out["promoted_object"] = *params.PromotedObject
	}
	if params.RbPredictionID != nil {
		out["rb_prediction_id"] = *params.RbPredictionID
	}
	if params.RegionalRegulatedCategories != nil {
		out["regional_regulated_categories"] = *params.RegionalRegulatedCategories
	}
	if params.RegionalRegulationIdentities != nil {
		out["regional_regulation_identities"] = *params.RegionalRegulationIdentities
	}
	if params.RelativeValue != nil {
		out["relative_value"] = *params.RelativeValue
	}
	if params.RfPredictionID != nil {
		out["rf_prediction_id"] = *params.RfPredictionID
	}
	if params.StartTime != nil {
		out["start_time"] = *params.StartTime
	}
	if params.Status != nil {
		out["status"] = *params.Status
	}
	if params.Targeting != nil {
		out["targeting"] = *params.Targeting
	}
	if params.TimeBasedAdRotationIDBlocks != nil {
		out["time_based_ad_rotation_id_blocks"] = *params.TimeBasedAdRotationIDBlocks
	}
	if params.TimeBasedAdRotationIntervals != nil {
		out["time_based_ad_rotation_intervals"] = *params.TimeBasedAdRotationIntervals
	}
	if params.TimeStart != nil {
		out["time_start"] = *params.TimeStart
	}
	if params.TimeStop != nil {
		out["time_stop"] = *params.TimeStop
	}
	if params.TrendingTopicsSpec != nil {
		out["trending_topics_spec"] = *params.TrendingTopicsSpec
	}
	if params.TuneForCategory != nil {
		out["tune_for_category"] = *params.TuneForCategory
	}
	if params.ValueRuleSetID != nil {
		out["value_rule_set_id"] = *params.ValueRuleSetID
	}
	if params.ValueRulesApplied != nil {
		out["value_rules_applied"] = *params.ValueRulesApplied
	}
	return out
}

func UpdateAdSet(ctx context.Context, client *core.Client, id string, params UpdateAdSetParams) (*objects.AdSet, error) {
	var out objects.AdSet
	err := client.Request(ctx, http.MethodPost, core.GraphPath(id), params.ToParams(), &out)
	return &out, err
}
