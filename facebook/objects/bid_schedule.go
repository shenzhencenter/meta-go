package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type BidSchedule struct {
	AdObjectID        *core.ID   `json:"ad_object_id,omitempty"`
	BidRecurrenceType *string    `json:"bid_recurrence_type,omitempty"`
	BidTimezone       *string    `json:"bid_timezone,omitempty"`
	BidValue          *int       `json:"bid_value,omitempty"`
	ID                *core.ID   `json:"id,omitempty"`
	Status            *string    `json:"status,omitempty"`
	TimeEnd           *core.Time `json:"time_end,omitempty"`
	TimeStart         *core.Time `json:"time_start,omitempty"`
}

var BidScheduleFields = struct {
	AdObjectID        string
	BidRecurrenceType string
	BidTimezone       string
	BidValue          string
	ID                string
	Status            string
	TimeEnd           string
	TimeStart         string
}{
	AdObjectID:        "ad_object_id",
	BidRecurrenceType: "bid_recurrence_type",
	BidTimezone:       "bid_timezone",
	BidValue:          "bid_value",
	ID:                "id",
	Status:            "status",
	TimeEnd:           "time_end",
	TimeStart:         "time_start",
}
