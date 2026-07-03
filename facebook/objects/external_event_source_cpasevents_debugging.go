package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
