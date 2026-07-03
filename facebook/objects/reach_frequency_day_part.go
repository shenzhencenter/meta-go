package objects

type ReachFrequencyDayPart struct {
	Days        *[]int `json:"days,omitempty"`
	EndMinute   *int   `json:"end_minute,omitempty"`
	StartMinute *int   `json:"start_minute,omitempty"`
}
