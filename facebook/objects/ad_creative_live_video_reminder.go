package objects

import (
	"time"
)

type AdCreativeLiveVideoReminder struct {
	EnrollStatus *string    `json:"enroll_status,omitempty"`
	StartTime    *time.Time `json:"start_time,omitempty"`
}
