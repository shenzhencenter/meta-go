package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type BidSchedule struct {
	AdObjectID        *core.ID   `json:"ad_object_id,omitempty"`
	BidRecurrenceType *string    `json:"bid_recurrence_type,omitempty"`
	BidTimezone       *string    `json:"bid_timezone,omitempty"`
	BidValue          *int       `json:"bid_value,omitempty"`
	ID                *core.ID   `json:"id,omitempty"`
	Status            *string    `json:"status,omitempty"`
	TimeEnd           *time.Time `json:"time_end,omitempty"`
	TimeStart         *time.Time `json:"time_start,omitempty"`
}
