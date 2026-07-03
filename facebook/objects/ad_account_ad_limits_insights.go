package objects

type AdAccountAdLimitsInsights struct {
	DateStart *string `json:"date_start,omitempty"`
	DateStop  *string `json:"date_stop,omitempty"`
}

var AdAccountAdLimitsInsightsFields = struct {
	DateStart string
	DateStop  string
}{
	DateStart: "date_start",
	DateStop:  "date_stop",
}
