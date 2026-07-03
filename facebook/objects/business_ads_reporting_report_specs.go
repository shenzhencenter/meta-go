package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
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
	CreationTime                         *time.Time                             `json:"creation_time,omitempty"`
	Currency                             *string                                `json:"currency,omitempty"`
	DatePreset                           *string                                `json:"date_preset,omitempty"`
	DefaultAttributionWindows            *[]string                              `json:"default_attribution_windows,omitempty"`
	Filtering                            *[]map[string]interface{}              `json:"filtering,omitempty"`
	Formatting                           *[]map[string][]map[string]interface{} `json:"formatting,omitempty"`
	ID                                   *core.ID                               `json:"id,omitempty"`
	LastAccessBy                         *Profile                               `json:"last_access_by,omitempty"`
	LastAccessTime                       *time.Time                             `json:"last_access_time,omitempty"`
	LastReportSnapshotID                 *core.ID                               `json:"last_report_snapshot_id,omitempty"`
	LastReportSnapshotTime               *time.Time                             `json:"last_report_snapshot_time,omitempty"`
	LastSharedReportExpiration           *time.Time                             `json:"last_shared_report_expiration,omitempty"`
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
	UpdateTime                           *time.Time                             `json:"update_time,omitempty"`
	User                                 *Profile                               `json:"user,omitempty"`
	UserDimensions                       *[]string                              `json:"user_dimensions,omitempty"`
	UserMetrics                          *[]string                              `json:"user_metrics,omitempty"`
	ViewType                             *string                                `json:"view_type,omitempty"`
}
