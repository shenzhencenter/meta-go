package objects

type ExternalEventSourceCPASEventsDebuggingInfo struct {
	Counts     *int    `json:"counts,omitempty"`
	Diagnostic *string `json:"diagnostic,omitempty"`
	EventName  *string `json:"event_name,omitempty"`
}

var ExternalEventSourceCPASEventsDebuggingInfoFields = struct {
	Counts     string
	Diagnostic string
	EventName  string
}{
	Counts:     "counts",
	Diagnostic: "diagnostic",
	EventName:  "event_name",
}
