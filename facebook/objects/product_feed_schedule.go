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
