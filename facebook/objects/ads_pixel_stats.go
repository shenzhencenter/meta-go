package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type AdsPixelStats struct {
	Count                          *int       `json:"count,omitempty"`
	DiagnosticsHourlyLastTimestamp *core.Time `json:"diagnostics_hourly_last_timestamp,omitempty"`
	Event                          *string    `json:"event,omitempty"`
	Value                          *string    `json:"value,omitempty"`
}

var AdsPixelStatsFields = struct {
	Count                          string
	DiagnosticsHourlyLastTimestamp string
	Event                          string
	Value                          string
}{
	Count:                          "count",
	DiagnosticsHourlyLastTimestamp: "diagnostics_hourly_last_timestamp",
	Event:                          "event",
	Value:                          "value",
}
