package objects

type StreamFilter struct {
	FilterKey *string `json:"filter_key,omitempty"`
	Name      *string `json:"name,omitempty"`
	Type      *string `json:"type,omitempty"`
}
