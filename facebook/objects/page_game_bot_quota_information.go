package objects

type PageGameBotQuotaInformation struct {
	Count      *int `json:"count,omitempty"`
	TimeWindow *int `json:"time_window,omitempty"`
}

var PageGameBotQuotaInformationFields = struct {
	Count      string
	TimeWindow string
}{
	Count:      "count",
	TimeWindow: "time_window",
}
