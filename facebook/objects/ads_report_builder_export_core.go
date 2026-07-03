package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type AdsReportBuilderExportCore struct {
	AsyncPercentCompletion    *int       `json:"async_percent_completion,omitempty"`
	AsyncReportURL            *string    `json:"async_report_url,omitempty"`
	AsyncStatus               *string    `json:"async_status,omitempty"`
	ClientCreationValue       *string    `json:"client_creation_value,omitempty"`
	ExpiryTime                *time.Time `json:"expiry_time,omitempty"`
	ExportDownloadTime        *time.Time `json:"export_download_time,omitempty"`
	ExportFormat              *string    `json:"export_format,omitempty"`
	ExportName                *string    `json:"export_name,omitempty"`
	ExportType                *string    `json:"export_type,omitempty"`
	HasSeen                   *bool      `json:"has_seen,omitempty"`
	ID                        *core.ID   `json:"id,omitempty"`
	IsSharing                 *bool      `json:"is_sharing,omitempty"`
	LinkSharingExpirationTime *time.Time `json:"link_sharing_expiration_time,omitempty"`
	LinkSharingURI            *string    `json:"link_sharing_uri,omitempty"`
	TimeCompleted             *time.Time `json:"time_completed,omitempty"`
	TimeStart                 *time.Time `json:"time_start,omitempty"`
}
