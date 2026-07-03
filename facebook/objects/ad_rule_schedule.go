package objects

type AdRuleSchedule struct {
	Days        *[]int `json:"days,omitempty"`
	EndMinute   *int   `json:"end_minute,omitempty"`
	StartMinute *int   `json:"start_minute,omitempty"`
}

var AdRuleScheduleFields = struct {
	Days        string
	EndMinute   string
	StartMinute string
}{
	Days:        "days",
	EndMinute:   "end_minute",
	StartMinute: "start_minute",
}
