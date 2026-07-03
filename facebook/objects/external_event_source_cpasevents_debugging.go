package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type ExternalEventSourceCPASEventsDebugging struct {
	ActualEventTime *int     `json:"actual_event_time,omitempty"`
	AppVersion      *string  `json:"app_version,omitempty"`
	ContentURL      *string  `json:"content_url,omitempty"`
	DeviceOs        *string  `json:"device_os,omitempty"`
	Diagnostic      *string  `json:"diagnostic,omitempty"`
	EventName       *string  `json:"event_name,omitempty"`
	EventTime       *int     `json:"event_time,omitempty"`
	MissingIds      *core.ID `json:"missing_ids,omitempty"`
	Severity        *string  `json:"severity,omitempty"`
}

var ExternalEventSourceCPASEventsDebuggingFields = struct {
	ActualEventTime string
	AppVersion      string
	ContentURL      string
	DeviceOs        string
	Diagnostic      string
	EventName       string
	EventTime       string
	MissingIds      string
	Severity        string
}{
	ActualEventTime: "actual_event_time",
	AppVersion:      "app_version",
	ContentURL:      "content_url",
	DeviceOs:        "device_os",
	Diagnostic:      "diagnostic",
	EventName:       "event_name",
	EventTime:       "event_time",
	MissingIds:      "missing_ids",
	Severity:        "severity",
}
