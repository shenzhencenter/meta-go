package objects

type VoipInfo struct {
	HasMobileApp      *bool   `json:"has_mobile_app,omitempty"`
	HasPermission     *bool   `json:"has_permission,omitempty"`
	IsCallable        *bool   `json:"is_callable,omitempty"`
	IsCallableWebrtc  *bool   `json:"is_callable_webrtc,omitempty"`
	IsPushable        *bool   `json:"is_pushable,omitempty"`
	ReasonCode        *uint64 `json:"reason_code,omitempty"`
	ReasonDescription *string `json:"reason_description,omitempty"`
}

var VoipInfoFields = struct {
	HasMobileApp      string
	HasPermission     string
	IsCallable        string
	IsCallableWebrtc  string
	IsPushable        string
	ReasonCode        string
	ReasonDescription string
}{
	HasMobileApp:      "has_mobile_app",
	HasPermission:     "has_permission",
	IsCallable:        "is_callable",
	IsCallableWebrtc:  "is_callable_webrtc",
	IsPushable:        "is_pushable",
	ReasonCode:        "reason_code",
	ReasonDescription: "reason_description",
}
