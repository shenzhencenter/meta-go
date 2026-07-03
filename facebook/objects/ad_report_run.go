package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
)

type AdReportRun struct {
	AccountID              *core.ID  `json:"account_id,omitempty"`
	AsyncPercentCompletion *uint64   `json:"async_percent_completion,omitempty"`
	AsyncReportURL         *string   `json:"async_report_url,omitempty"`
	AsyncStatus            *string   `json:"async_status,omitempty"`
	DateStart              *string   `json:"date_start,omitempty"`
	DateStop               *string   `json:"date_stop,omitempty"`
	Emails                 *[]string `json:"emails,omitempty"`
	ErrorCode              *int      `json:"error_code,omitempty"`
	ErrorMessage           *string   `json:"error_message,omitempty"`
	ErrorSubcode           *uint64   `json:"error_subcode,omitempty"`
	ErrorUserMsg           *string   `json:"error_user_msg,omitempty"`
	ErrorUserTitle         *string   `json:"error_user_title,omitempty"`
	FriendlyName           *string   `json:"friendly_name,omitempty"`
	ID                     *core.ID  `json:"id,omitempty"`
	IsAsyncExport          *int      `json:"is_async_export,omitempty"`
	IsBookmarked           *bool     `json:"is_bookmarked,omitempty"`
	IsRunning              *bool     `json:"is_running,omitempty"`
	ScheduleID             *core.ID  `json:"schedule_id,omitempty"`
	TimeCompleted          *uint64   `json:"time_completed,omitempty"`
	TimeRef                *uint64   `json:"time_ref,omitempty"`
}
