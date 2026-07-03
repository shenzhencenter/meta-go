package objects

type AttributionSpec struct {
	EventType  *string `json:"event_type,omitempty"`
	WindowDays *int    `json:"window_days,omitempty"`
}
