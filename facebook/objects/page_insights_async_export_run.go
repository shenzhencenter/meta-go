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
