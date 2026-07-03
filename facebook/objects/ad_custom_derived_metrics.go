package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCustomDerivedMetrics struct {
	AdAccountID                *core.ID   `json:"ad_account_id,omitempty"`
	Business                   *Business  `json:"business,omitempty"`
	CreationTime               *core.Time `json:"creation_time,omitempty"`
	Creator                    *Profile   `json:"creator,omitempty"`
	CustomDerivedMetricType    *string    `json:"custom_derived_metric_type,omitempty"`
	DeletionTime               *core.Time `json:"deletion_time,omitempty"`
	Deletor                    *Profile   `json:"deletor,omitempty"`
	Description                *string    `json:"description,omitempty"`
	FormatType                 *string    `json:"format_type,omitempty"`
	Formula                    *string    `json:"formula,omitempty"`
	HasAttributionWindows      *bool      `json:"has_attribution_windows,omitempty"`
	HasInlineAttributionWindow *bool      `json:"has_inline_attribution_window,omitempty"`
	ID                         *core.ID   `json:"id,omitempty"`
	Name                       *string    `json:"name,omitempty"`
	Permission                 *string    `json:"permission,omitempty"`
	SavedReportID              *core.ID   `json:"saved_report_id,omitempty"`
	Scope                      *string    `json:"scope,omitempty"`
}

var AdCustomDerivedMetricsFields = struct {
	AdAccountID                string
	Business                   string
	CreationTime               string
	Creator                    string
	CustomDerivedMetricType    string
	DeletionTime               string
	Deletor                    string
	Description                string
	FormatType                 string
	Formula                    string
	HasAttributionWindows      string
	HasInlineAttributionWindow string
	ID                         string
	Name                       string
	Permission                 string
	SavedReportID              string
	Scope                      string
}{
	AdAccountID:                "ad_account_id",
	Business:                   "business",
	CreationTime:               "creation_time",
	Creator:                    "creator",
	CustomDerivedMetricType:    "custom_derived_metric_type",
	DeletionTime:               "deletion_time",
	Deletor:                    "deletor",
	Description:                "description",
	FormatType:                 "format_type",
	Formula:                    "formula",
	HasAttributionWindows:      "has_attribution_windows",
	HasInlineAttributionWindow: "has_inline_attribution_window",
	ID:                         "id",
	Name:                       "name",
	Permission:                 "permission",
	SavedReportID:              "saved_report_id",
	Scope:                      "scope",
}
