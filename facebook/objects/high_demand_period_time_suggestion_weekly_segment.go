package objects

type HighDemandPeriodTimeSuggestionWeeklySegment struct {
	Days         *[]string `json:"days,omitempty"`
	EndMinute    *int      `json:"end_minute,omitempty"`
	StartMinute  *int      `json:"start_minute,omitempty"`
	TimezoneType *string   `json:"timezone_type,omitempty"`
}

var HighDemandPeriodTimeSuggestionWeeklySegmentFields = struct {
	Days         string
	EndMinute    string
	StartMinute  string
	TimezoneType string
}{
	Days:         "days",
	EndMinute:    "end_minute",
	StartMinute:  "start_minute",
	TimezoneType: "timezone_type",
}
