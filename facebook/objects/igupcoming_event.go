package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type IGUpcomingEvent struct {
	EndTime                *time.Time `json:"end_time,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	NotificationSubtypes   *[]string  `json:"notification_subtypes,omitempty"`
	NotificationTargetTime *string    `json:"notification_target_time,omitempty"`
	StartTime              *time.Time `json:"start_time,omitempty"`
	Title                  *string    `json:"title,omitempty"`
}
