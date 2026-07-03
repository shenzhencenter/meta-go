package objects

type ExternalEventSourceCPASEventsDebuggingInfo struct {
	Counts     *int    `json:"counts,omitempty"`
	Diagnostic *string `json:"diagnostic,omitempty"`
	EventName  *string `json:"event_name,omitempty"`
}
