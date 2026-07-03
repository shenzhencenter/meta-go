package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsReportBuilderMMMReport struct {
	AsyncStatus  *string    `json:"async_status,omitempty"`
	ExportFormat *string    `json:"export_format,omitempty"`
	ExportName   *string    `json:"export_name,omitempty"`
	ExportType   *string    `json:"export_type,omitempty"`
	HasSeen      *bool      `json:"has_seen,omitempty"`
	ID           *core.ID   `json:"id,omitempty"`
	MmmStatus    *string    `json:"mmm_status,omitempty"`
	TimeStart    *core.Time `json:"time_start,omitempty"`
}

var AdsReportBuilderMMMReportFields = struct {
	AsyncStatus  string
	ExportFormat string
	ExportName   string
	ExportType   string
	HasSeen      string
	ID           string
	MmmStatus    string
	TimeStart    string
}{
	AsyncStatus:  "async_status",
	ExportFormat: "export_format",
	ExportName:   "export_name",
	ExportType:   "export_type",
	HasSeen:      "has_seen",
	ID:           "id",
	MmmStatus:    "mmm_status",
	TimeStart:    "time_start",
}
