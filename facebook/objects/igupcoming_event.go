package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type IGUpcomingEvent struct {
	EndTime                *core.Time `json:"end_time,omitempty"`
	ID                     *core.ID   `json:"id,omitempty"`
	NotificationSubtypes   *[]string  `json:"notification_subtypes,omitempty"`
	NotificationTargetTime *string    `json:"notification_target_time,omitempty"`
	StartTime              *core.Time `json:"start_time,omitempty"`
	Title                  *string    `json:"title,omitempty"`
}

var IGUpcomingEventFields = struct {
	EndTime                string
	ID                     string
	NotificationSubtypes   string
	NotificationTargetTime string
	StartTime              string
	Title                  string
}{
	EndTime:                "end_time",
	ID:                     "id",
	NotificationSubtypes:   "notification_subtypes",
	NotificationTargetTime: "notification_target_time",
	StartTime:              "start_time",
	Title:                  "title",
}
