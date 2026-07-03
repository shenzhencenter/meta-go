package objects

type StreamFilter struct {
	FilterKey *string `json:"filter_key,omitempty"`
	Name      *string `json:"name,omitempty"`
	Type      *string `json:"type,omitempty"`
}

var StreamFilterFields = struct {
	FilterKey string
	Name      string
	Type      string
}{
	FilterKey: "filter_key",
	Name:      "name",
	Type:      "type",
}
