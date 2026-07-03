package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type MeasurementReport struct {
	DownloadUrls *[]string               `json:"download_urls,omitempty"`
	ID           *core.ID                `json:"id,omitempty"`
	Metadata     *map[string]interface{} `json:"metadata,omitempty"`
	ReportType   *string                 `json:"report_type,omitempty"`
	Status       *string                 `json:"status,omitempty"`
}

var MeasurementReportFields = struct {
	DownloadUrls string
	ID           string
	Metadata     string
	ReportType   string
	Status       string
}{
	DownloadUrls: "download_urls",
	ID:           "id",
	Metadata:     "metadata",
	ReportType:   "report_type",
	Status:       "status",
}
