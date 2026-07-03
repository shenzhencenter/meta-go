package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type AdAsyncRequestSet struct {
	CanceledCount      *int                                         `json:"canceled_count,omitempty"`
	CreatedTime        *core.Time                                   `json:"created_time,omitempty"`
	ErrorCount         *int                                         `json:"error_count,omitempty"`
	ID                 *core.ID                                     `json:"id,omitempty"`
	InProgressCount    *int                                         `json:"in_progress_count,omitempty"`
	InitialCount       *uint64                                      `json:"initial_count,omitempty"`
	IsCompleted        *bool                                        `json:"is_completed,omitempty"`
	Name               *string                                      `json:"name,omitempty"`
	NotificationMode   *enums.AdasyncrequestsetNotificationModeEnum `json:"notification_mode,omitempty"`
	NotificationResult *AdAsyncRequestSetNotificationResult         `json:"notification_result,omitempty"`
	NotificationStatus *string                                      `json:"notification_status,omitempty"`
	NotificationURI    *string                                      `json:"notification_uri,omitempty"`
	OwnerID            *core.ID                                     `json:"owner_id,omitempty"`
	SuccessCount       *int                                         `json:"success_count,omitempty"`
	TotalCount         *uint64                                      `json:"total_count,omitempty"`
	UpdatedTime        *core.Time                                   `json:"updated_time,omitempty"`
}

var AdAsyncRequestSetFields = struct {
	CanceledCount      string
	CreatedTime        string
	ErrorCount         string
	ID                 string
	InProgressCount    string
	InitialCount       string
	IsCompleted        string
	Name               string
	NotificationMode   string
	NotificationResult string
	NotificationStatus string
	NotificationURI    string
	OwnerID            string
	SuccessCount       string
	TotalCount         string
	UpdatedTime        string
}{
	CanceledCount:      "canceled_count",
	CreatedTime:        "created_time",
	ErrorCount:         "error_count",
	ID:                 "id",
	InProgressCount:    "in_progress_count",
	InitialCount:       "initial_count",
	IsCompleted:        "is_completed",
	Name:               "name",
	NotificationMode:   "notification_mode",
	NotificationResult: "notification_result",
	NotificationStatus: "notification_status",
	NotificationURI:    "notification_uri",
	OwnerID:            "owner_id",
	SuccessCount:       "success_count",
	TotalCount:         "total_count",
	UpdatedTime:        "updated_time",
}
