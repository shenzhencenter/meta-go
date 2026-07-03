package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type HighDemandPeriodGet struct {
	AdObjectID      *core.ID                  `json:"ad_object_id,omitempty"`
	BudgetValue     *int                      `json:"budget_value,omitempty"`
	BudgetValueType *string                   `json:"budget_value_type,omitempty"`
	ID              *core.ID                  `json:"id,omitempty"`
	RecurrenceType  *string                   `json:"recurrence_type,omitempty"`
	TimeEnd         *string                   `json:"time_end,omitempty"`
	TimeStart       *string                   `json:"time_start,omitempty"`
	WeeklySchedule  *[]map[string]interface{} `json:"weekly_schedule,omitempty"`
}

var HighDemandPeriodGetFields = struct {
	AdObjectID      string
	BudgetValue     string
	BudgetValueType string
	ID              string
	RecurrenceType  string
	TimeEnd         string
	TimeStart       string
	WeeklySchedule  string
}{
	AdObjectID:      "ad_object_id",
	BudgetValue:     "budget_value",
	BudgetValueType: "budget_value_type",
	ID:              "id",
	RecurrenceType:  "recurrence_type",
	TimeEnd:         "time_end",
	TimeStart:       "time_start",
	WeeklySchedule:  "weekly_schedule",
}
