package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type ProductFeedSchedule struct {
	DayOfMonth    *uint64                            `json:"day_of_month,omitempty"`
	DayOfWeek     *string                            `json:"day_of_week,omitempty"`
	Hour          *uint64                            `json:"hour,omitempty"`
	ID            *core.ID                           `json:"id,omitempty"`
	Interval      *enums.ProductfeedscheduleInterval `json:"interval,omitempty"`
	IntervalCount *uint64                            `json:"interval_count,omitempty"`
	Minute        *uint64                            `json:"minute,omitempty"`
	Timezone      *string                            `json:"timezone,omitempty"`
	URL           *string                            `json:"url,omitempty"`
	Username      *string                            `json:"username,omitempty"`
}

var ProductFeedScheduleFields = struct {
	DayOfMonth    string
	DayOfWeek     string
	Hour          string
	ID            string
	Interval      string
	IntervalCount string
	Minute        string
	Timezone      string
	URL           string
	Username      string
}{
	DayOfMonth:    "day_of_month",
	DayOfWeek:     "day_of_week",
	Hour:          "hour",
	ID:            "id",
	Interval:      "interval",
	IntervalCount: "interval_count",
	Minute:        "minute",
	Timezone:      "timezone",
	URL:           "url",
	Username:      "username",
}
