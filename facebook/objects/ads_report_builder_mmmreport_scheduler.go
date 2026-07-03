package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdsReportBuilderMMMReportScheduler struct {
	AdAccountIds      *[]core.ID                `json:"ad_account_ids,omitempty"`
	Filtering         *[]map[string]interface{} `json:"filtering,omitempty"`
	ID                *core.ID                  `json:"id,omitempty"`
	ReportName        *string                   `json:"report_name,omitempty"`
	ScheduleFrequency *string                   `json:"schedule_frequency,omitempty"`
}
