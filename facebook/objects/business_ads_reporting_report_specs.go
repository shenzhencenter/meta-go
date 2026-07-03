package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BusinessAdsReportingReportSpecs struct {
	ActionReportTime                     *string                                `json:"action_report_time,omitempty"`
	AdAccountID                          *core.ID                               `json:"ad_account_id,omitempty"`
	AdAccountIds                         *[]core.ID                             `json:"ad_account_ids,omitempty"`
	AdAccounts                           *[]map[string]interface{}              `json:"ad_accounts,omitempty"`
	AttributionWindows                   *[]string                              `json:"attribution_windows,omitempty"`
	Business                             *Business                              `json:"business,omitempty"`
	BusinessAssetGroup                   *BusinessAssetGroup                    `json:"business_asset_group,omitempty"`
	ComparisonDateInterval               *map[string]interface{}                `json:"comparison_date_interval,omitempty"`
	CreationSource                       *string                                `json:"creation_source,omitempty"`
	CreationTime                         *core.Time                             `json:"creation_time,omitempty"`
	Currency                             *string                                `json:"currency,omitempty"`
	DatePreset                           *string                                `json:"date_preset,omitempty"`
	DefaultAttributionWindows            *[]string                              `json:"default_attribution_windows,omitempty"`
	Filtering                            *[]map[string]interface{}              `json:"filtering,omitempty"`
	Formatting                           *[]map[string][]map[string]interface{} `json:"formatting,omitempty"`
	ID                                   *core.ID                               `json:"id,omitempty"`
	LastAccessBy                         *Profile                               `json:"last_access_by,omitempty"`
	LastAccessTime                       *core.Time                             `json:"last_access_time,omitempty"`
	LastReportSnapshotID                 *core.ID                               `json:"last_report_snapshot_id,omitempty"`
	LastReportSnapshotTime               *core.Time                             `json:"last_report_snapshot_time,omitempty"`
	LastSharedReportExpiration           *core.Time                             `json:"last_shared_report_expiration,omitempty"`
	Limit                                *int                                   `json:"limit,omitempty"`
	LockedDimensions                     *int                                   `json:"locked_dimensions,omitempty"`
	ReportName                           *string                                `json:"report_name,omitempty"`
	ReportSnapshotAsyncPercentCompletion *int                                   `json:"report_snapshot_async_percent_completion,omitempty"`
	ReportSnapshotAsyncStatus            *string                                `json:"report_snapshot_async_status,omitempty"`
	ScheduleFrequency                    *string                                `json:"schedule_frequency,omitempty"`
	Scope                                *string                                `json:"scope,omitempty"`
	ShowDeprecateAwBanner                *bool                                  `json:"show_deprecate_aw_banner,omitempty"`
	Sorting                              *[]map[string]interface{}              `json:"sorting,omitempty"`
	StartDate                            *string                                `json:"start_date,omitempty"`
	Status                               *string                                `json:"status,omitempty"`
	Subscribers                          *[]string                              `json:"subscribers,omitempty"`
	UpdateBy                             *Profile                               `json:"update_by,omitempty"`
	UpdateTime                           *core.Time                             `json:"update_time,omitempty"`
	User                                 *Profile                               `json:"user,omitempty"`
	UserDimensions                       *[]string                              `json:"user_dimensions,omitempty"`
	UserMetrics                          *[]string                              `json:"user_metrics,omitempty"`
	ViewType                             *string                                `json:"view_type,omitempty"`
}

var BusinessAdsReportingReportSpecsFields = struct {
	ActionReportTime                     string
	AdAccountID                          string
	AdAccountIds                         string
	AdAccounts                           string
	AttributionWindows                   string
	Business                             string
	BusinessAssetGroup                   string
	ComparisonDateInterval               string
	CreationSource                       string
	CreationTime                         string
	Currency                             string
	DatePreset                           string
	DefaultAttributionWindows            string
	Filtering                            string
	Formatting                           string
	ID                                   string
	LastAccessBy                         string
	LastAccessTime                       string
	LastReportSnapshotID                 string
	LastReportSnapshotTime               string
	LastSharedReportExpiration           string
	Limit                                string
	LockedDimensions                     string
	ReportName                           string
	ReportSnapshotAsyncPercentCompletion string
	ReportSnapshotAsyncStatus            string
	ScheduleFrequency                    string
	Scope                                string
	ShowDeprecateAwBanner                string
	Sorting                              string
	StartDate                            string
	Status                               string
	Subscribers                          string
	UpdateBy                             string
	UpdateTime                           string
	User                                 string
	UserDimensions                       string
	UserMetrics                          string
	ViewType                             string
}{
	ActionReportTime:                     "action_report_time",
	AdAccountID:                          "ad_account_id",
	AdAccountIds:                         "ad_account_ids",
	AdAccounts:                           "ad_accounts",
	AttributionWindows:                   "attribution_windows",
	Business:                             "business",
	BusinessAssetGroup:                   "business_asset_group",
	ComparisonDateInterval:               "comparison_date_interval",
	CreationSource:                       "creation_source",
	CreationTime:                         "creation_time",
	Currency:                             "currency",
	DatePreset:                           "date_preset",
	DefaultAttributionWindows:            "default_attribution_windows",
	Filtering:                            "filtering",
	Formatting:                           "formatting",
	ID:                                   "id",
	LastAccessBy:                         "last_access_by",
	LastAccessTime:                       "last_access_time",
	LastReportSnapshotID:                 "last_report_snapshot_id",
	LastReportSnapshotTime:               "last_report_snapshot_time",
	LastSharedReportExpiration:           "last_shared_report_expiration",
	Limit:                                "limit",
	LockedDimensions:                     "locked_dimensions",
	ReportName:                           "report_name",
	ReportSnapshotAsyncPercentCompletion: "report_snapshot_async_percent_completion",
	ReportSnapshotAsyncStatus:            "report_snapshot_async_status",
	ScheduleFrequency:                    "schedule_frequency",
	Scope:                                "scope",
	ShowDeprecateAwBanner:                "show_deprecate_aw_banner",
	Sorting:                              "sorting",
	StartDate:                            "start_date",
	Status:                               "status",
	Subscribers:                          "subscribers",
	UpdateBy:                             "update_by",
	UpdateTime:                           "update_time",
	User:                                 "user",
	UserDimensions:                       "user_dimensions",
	UserMetrics:                          "user_metrics",
	ViewType:                             "view_type",
}
