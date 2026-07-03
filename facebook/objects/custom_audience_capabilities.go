package objects

type CustomAudienceCapabilities struct {
	Capabilities *map[string]interface{} `json:"capabilities,omitempty"`
}

var CustomAudienceCapabilitiesFields = struct {
	Capabilities string
}{
	Capabilities: "capabilities",
}
