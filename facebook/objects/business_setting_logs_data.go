package objects

type BusinessSettingLogsData struct {
	Actor       *map[string]interface{} `json:"actor,omitempty"`
	EventObject *map[string]interface{} `json:"event_object,omitempty"`
	EventTime   *string                 `json:"event_time,omitempty"`
	EventType   *string                 `json:"event_type,omitempty"`
	ExtraData   *map[string]interface{} `json:"extra_data,omitempty"`
}
