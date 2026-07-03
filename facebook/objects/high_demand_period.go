package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"time"
)

type HighDemandPeriod struct {
	AdObjectID      *core.ID                                       `json:"ad_object_id,omitempty"`
	BudgetValue     *int                                           `json:"budget_value,omitempty"`
	BudgetValueType *string                                        `json:"budget_value_type,omitempty"`
	ID              *core.ID                                       `json:"id,omitempty"`
	RecurrenceType  *string                                        `json:"recurrence_type,omitempty"`
	TimeEnd         *time.Time                                     `json:"time_end,omitempty"`
	TimeStart       *time.Time                                     `json:"time_start,omitempty"`
	WeeklySchedule  *[]HighDemandPeriodTimeSuggestionWeeklySegment `json:"weekly_schedule,omitempty"`
}
