package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdNetworkAnalyticsAsyncQueryExport struct {
	ExportLink *string  `json:"export_link,omitempty"`
	QueryID    *core.ID `json:"query_id,omitempty"`
	Status     *string  `json:"status,omitempty"`
}
