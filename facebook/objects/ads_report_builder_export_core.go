package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsReportBuilderExportCore struct {
	AsyncPercentCompletion    *int       `json:"async_percent_completion,omitempty"`
	AsyncReportURL            *string    `json:"async_report_url,omitempty"`
	AsyncStatus               *string    `json:"async_status,omitempty"`
	ClientCreationValue       *string    `json:"client_creation_value,omitempty"`
	ExpiryTime                *core.Time `json:"expiry_time,omitempty"`
	ExportDownloadTime        *core.Time `json:"export_download_time,omitempty"`
	ExportFormat              *string    `json:"export_format,omitempty"`
	ExportName                *string    `json:"export_name,omitempty"`
	ExportType                *string    `json:"export_type,omitempty"`
	HasSeen                   *bool      `json:"has_seen,omitempty"`
	ID                        *core.ID   `json:"id,omitempty"`
	IsSharing                 *bool      `json:"is_sharing,omitempty"`
	LinkSharingExpirationTime *core.Time `json:"link_sharing_expiration_time,omitempty"`
	LinkSharingURI            *string    `json:"link_sharing_uri,omitempty"`
	TimeCompleted             *core.Time `json:"time_completed,omitempty"`
	TimeStart                 *core.Time `json:"time_start,omitempty"`
}

var AdsReportBuilderExportCoreFields = struct {
	AsyncPercentCompletion    string
	AsyncReportURL            string
	AsyncStatus               string
	ClientCreationValue       string
	ExpiryTime                string
	ExportDownloadTime        string
	ExportFormat              string
	ExportName                string
	ExportType                string
	HasSeen                   string
	ID                        string
	IsSharing                 string
	LinkSharingExpirationTime string
	LinkSharingURI            string
	TimeCompleted             string
	TimeStart                 string
}{
	AsyncPercentCompletion:    "async_percent_completion",
	AsyncReportURL:            "async_report_url",
	AsyncStatus:               "async_status",
	ClientCreationValue:       "client_creation_value",
	ExpiryTime:                "expiry_time",
	ExportDownloadTime:        "export_download_time",
	ExportFormat:              "export_format",
	ExportName:                "export_name",
	ExportType:                "export_type",
	HasSeen:                   "has_seen",
	ID:                        "id",
	IsSharing:                 "is_sharing",
	LinkSharingExpirationTime: "link_sharing_expiration_time",
	LinkSharingURI:            "link_sharing_uri",
	TimeCompleted:             "time_completed",
	TimeStart:                 "time_start",
}
