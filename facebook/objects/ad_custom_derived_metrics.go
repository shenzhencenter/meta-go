package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdCustomDerivedMetrics struct {
	AdAccountID                *core.ID   `json:"ad_account_id,omitempty"`
	Business                   *Business  `json:"business,omitempty"`
	CreationTime               *time.Time `json:"creation_time,omitempty"`
	Creator                    *Profile   `json:"creator,omitempty"`
	CustomDerivedMetricType    *string    `json:"custom_derived_metric_type,omitempty"`
	DeletionTime               *time.Time `json:"deletion_time,omitempty"`
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
