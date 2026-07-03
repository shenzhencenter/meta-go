package objects

type AttributionSpec struct {
	EventType  *string `json:"event_type,omitempty"`
	WindowDays *int    `json:"window_days,omitempty"`
}

var AttributionSpecFields = struct {
	EventType  string
	WindowDays string
}{
	EventType:  "event_type",
	WindowDays: "window_days",
}
