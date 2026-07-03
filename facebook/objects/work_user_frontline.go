package objects

type WorkUserFrontline struct {
	HasAccess   *bool `json:"has_access,omitempty"`
	IsFrontline *bool `json:"is_frontline,omitempty"`
}

var WorkUserFrontlineFields = struct {
	HasAccess   string
	IsFrontline string
}{
	HasAccess:   "has_access",
	IsFrontline: "is_frontline",
}
