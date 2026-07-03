package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsQuickViews struct {
	AttributionWindows           *[]string                 `json:"attribution_windows,omitempty"`
	Breakdowns                   *[]string                 `json:"breakdowns,omitempty"`
	ColumnFields                 *[]string                 `json:"column_fields,omitempty"`
	Description                  *string                   `json:"description,omitempty"`
	Grouping                     *string                   `json:"grouping,omitempty"`
	ID                           *core.ID                  `json:"id,omitempty"`
	IsAttributionWindowsDisabled *bool                     `json:"is_attribution_windows_disabled,omitempty"`
	IsBreakdownsDisabled         *bool                     `json:"is_breakdowns_disabled,omitempty"`
	IsColumnsAndSortDisabled     *bool                     `json:"is_columns_and_sort_disabled,omitempty"`
	IsFiltersDisabled            *bool                     `json:"is_filters_disabled,omitempty"`
	IsGroupingDisabled           *bool                     `json:"is_grouping_disabled,omitempty"`
	Name                         *string                   `json:"name,omitempty"`
	Owner                        *Profile                  `json:"owner,omitempty"`
	Permission                   *string                   `json:"permission,omitempty"`
	QuickViewType                *string                   `json:"quick_view_type,omitempty"`
	Sort                         *[]map[string]interface{} `json:"sort,omitempty"`
	TimeStampLastUsedByOwner     *core.Time                `json:"time_stamp_last_used_by_owner,omitempty"`
}

var AdsQuickViewsFields = struct {
	AttributionWindows           string
	Breakdowns                   string
	ColumnFields                 string
	Description                  string
	Grouping                     string
	ID                           string
	IsAttributionWindowsDisabled string
	IsBreakdownsDisabled         string
	IsColumnsAndSortDisabled     string
	IsFiltersDisabled            string
	IsGroupingDisabled           string
	Name                         string
	Owner                        string
	Permission                   string
	QuickViewType                string
	Sort                         string
	TimeStampLastUsedByOwner     string
}{
	AttributionWindows:           "attribution_windows",
	Breakdowns:                   "breakdowns",
	ColumnFields:                 "column_fields",
	Description:                  "description",
	Grouping:                     "grouping",
	ID:                           "id",
	IsAttributionWindowsDisabled: "is_attribution_windows_disabled",
	IsBreakdownsDisabled:         "is_breakdowns_disabled",
	IsColumnsAndSortDisabled:     "is_columns_and_sort_disabled",
	IsFiltersDisabled:            "is_filters_disabled",
	IsGroupingDisabled:           "is_grouping_disabled",
	Name:                         "name",
	Owner:                        "owner",
	Permission:                   "permission",
	QuickViewType:                "quick_view_type",
	Sort:                         "sort",
	TimeStampLastUsedByOwner:     "time_stamp_last_used_by_owner",
}
