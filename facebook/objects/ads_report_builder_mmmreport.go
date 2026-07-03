package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type AdsReportBuilderMMMReport struct {
	AsyncStatus  *string    `json:"async_status,omitempty"`
	ExportFormat *string    `json:"export_format,omitempty"`
	ExportName   *string    `json:"export_name,omitempty"`
	ExportType   *string    `json:"export_type,omitempty"`
	HasSeen      *bool      `json:"has_seen,omitempty"`
	ID           *core.ID   `json:"id,omitempty"`
	MmmStatus    *string    `json:"mmm_status,omitempty"`
	TimeStart    *time.Time `json:"time_start,omitempty"`
}
