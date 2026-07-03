package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"github.com/shenzhencenter/facebook-go-sdk/facebook/enums"
	"time"
)

type AdAsyncRequestSet struct {
	CanceledCount      *int                                         `json:"canceled_count,omitempty"`
	CreatedTime        *time.Time                                   `json:"created_time,omitempty"`
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
	UpdatedTime        *time.Time                                   `json:"updated_time,omitempty"`
}
