package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdCreativeLiveVideoReminder struct {
	EnrollStatus *string    `json:"enroll_status,omitempty"`
	StartTime    *core.Time `json:"start_time,omitempty"`
}

var AdCreativeLiveVideoReminderFields = struct {
	EnrollStatus string
	StartTime    string
}{
	EnrollStatus: "enroll_status",
	StartTime:    "start_time",
}
