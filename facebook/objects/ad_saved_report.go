package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdSavedReport struct {
	AppOwner               *Application              `json:"app_owner,omitempty"`
	Breakdowns             *[]string                 `json:"breakdowns,omitempty"`
	BuiltinColumnSet       *string                   `json:"builtin_column_set,omitempty"`
	CreationSource         *string                   `json:"creation_source,omitempty"`
	DateInterval           *map[string]interface{}   `json:"date_interval,omitempty"`
	DatePreset             *string                   `json:"date_preset,omitempty"`
	FormatVersion          *int                      `json:"format_version,omitempty"`
	ID                     *core.ID                  `json:"id,omitempty"`
	InsightsSection        *map[string]interface{}   `json:"insights_section,omitempty"`
	IsSharedUnread         *bool                     `json:"is_shared_unread,omitempty"`
	Level                  *string                   `json:"level,omitempty"`
	Name                   *string                   `json:"name,omitempty"`
	NormalizedFilter       *[]interface{}            `json:"normalized_filter,omitempty"`
	Sort                   *[]map[string]interface{} `json:"sort,omitempty"`
	UserAttributionWindows *[]string                 `json:"user_attribution_windows,omitempty"`
	UserColumns            *[]string                 `json:"user_columns,omitempty"`
	UserFilter             *[]interface{}            `json:"user_filter,omitempty"`
	UserOwner              *User                     `json:"user_owner,omitempty"`
}
