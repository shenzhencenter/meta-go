package objects

import (
	"time"
)

type AdsPixelStats struct {
	Count                          *int       `json:"count,omitempty"`
	DiagnosticsHourlyLastTimestamp *time.Time `json:"diagnostics_hourly_last_timestamp,omitempty"`
	Event                          *string    `json:"event,omitempty"`
	Value                          *string    `json:"value,omitempty"`
}
