package objects

type MessengerCallSettings struct {
	AudioEnabled *bool                   `json:"audio_enabled,omitempty"`
	CallHours    *map[string]interface{} `json:"call_hours,omitempty"`
	CallRouting  *map[string]interface{} `json:"call_routing,omitempty"`
	IconEnabled  *bool                   `json:"icon_enabled,omitempty"`
	VideoEnabled *bool                   `json:"video_enabled,omitempty"`
}
