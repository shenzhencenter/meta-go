package objects

type WorkUserFrontline struct {
	HasAccess   *bool `json:"has_access,omitempty"`
	IsFrontline *bool `json:"is_frontline,omitempty"`
}
