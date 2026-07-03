package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type RightsManagerDataExport struct {
	DownloadURI    *string    `json:"download_uri,omitempty"`
	ExportScope    *string    `json:"export_scope,omitempty"`
	ID             *core.ID   `json:"id,omitempty"`
	Name           *string    `json:"name,omitempty"`
	RecordType     *string    `json:"record_type,omitempty"`
	TimeRangeEnd   *core.Time `json:"time_range_end,omitempty"`
	TimeRangeStart *core.Time `json:"time_range_start,omitempty"`
}

var RightsManagerDataExportFields = struct {
	DownloadURI    string
	ExportScope    string
	ID             string
	Name           string
	RecordType     string
	TimeRangeEnd   string
	TimeRangeStart string
}{
	DownloadURI:    "download_uri",
	ExportScope:    "export_scope",
	ID:             "id",
	Name:           "name",
	RecordType:     "record_type",
	TimeRangeEnd:   "time_range_end",
	TimeRangeStart: "time_range_start",
}
