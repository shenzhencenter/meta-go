package objects

type AdRuleScheduleSpec struct {
	Schedule     *[]AdRuleSchedule `json:"schedule,omitempty"`
	ScheduleType *string           `json:"schedule_type,omitempty"`
}

var AdRuleScheduleSpecFields = struct {
	Schedule     string
	ScheduleType string
}{
	Schedule:     "schedule",
	ScheduleType: "schedule_type",
}
