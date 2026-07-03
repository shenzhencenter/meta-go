package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdNetworkAnalyticsAsyncQueryExport struct {
	ExportLink *string  `json:"export_link,omitempty"`
	QueryID    *core.ID `json:"query_id,omitempty"`
	Status     *string  `json:"status,omitempty"`
}

var AdNetworkAnalyticsAsyncQueryExportFields = struct {
	ExportLink string
	QueryID    string
	Status     string
}{
	ExportLink: "export_link",
	QueryID:    "query_id",
	Status:     "status",
}
