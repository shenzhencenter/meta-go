package objects

type BusinessSettingLogsData struct {
	Actor       *map[string]interface{} `json:"actor,omitempty"`
	EventObject *map[string]interface{} `json:"event_object,omitempty"`
	EventTime   *string                 `json:"event_time,omitempty"`
	EventType   *string                 `json:"event_type,omitempty"`
	ExtraData   *map[string]interface{} `json:"extra_data,omitempty"`
}

var BusinessSettingLogsDataFields = struct {
	Actor       string
	EventObject string
	EventTime   string
	EventType   string
	ExtraData   string
}{
	Actor:       "actor",
	EventObject: "event_object",
	EventTime:   "event_time",
	EventType:   "event_type",
	ExtraData:   "extra_data",
}
