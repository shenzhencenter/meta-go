package objects

type AdRuleScheduleSpec struct {
	Schedule     *[]AdRuleSchedule `json:"schedule,omitempty"`
	ScheduleType *string           `json:"schedule_type,omitempty"`
}
