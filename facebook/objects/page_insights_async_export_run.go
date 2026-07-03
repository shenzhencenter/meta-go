package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type PageInsightsAsyncExportRun struct {
	DataLevel       *string                   `json:"data_level,omitempty"`
	Filters         *[]map[string]interface{} `json:"filters,omitempty"`
	Format          *string                   `json:"format,omitempty"`
	GenReportDate   *int                      `json:"gen_report_date,omitempty"`
	ID              *core.ID                  `json:"id,omitempty"`
	ReportEndDate   *int                      `json:"report_end_date,omitempty"`
	ReportStartDate *int                      `json:"report_start_date,omitempty"`
	Sorters         *[]map[string]interface{} `json:"sorters,omitempty"`
	Status          *string                   `json:"status,omitempty"`
}

var PageInsightsAsyncExportRunFields = struct {
	DataLevel       string
	Filters         string
	Format          string
	GenReportDate   string
	ID              string
	ReportEndDate   string
	ReportStartDate string
	Sorters         string
	Status          string
}{
	DataLevel:       "data_level",
	Filters:         "filters",
	Format:          "format",
	GenReportDate:   "gen_report_date",
	ID:              "id",
	ReportEndDate:   "report_end_date",
	ReportStartDate: "report_start_date",
	Sorters:         "sorters",
	Status:          "status",
}
