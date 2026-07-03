package objects

type ReachFrequencyAdFormat struct {
	Details *map[string]interface{} `json:"details,omitempty"`
	Type    *string                 `json:"type,omitempty"`
}

var ReachFrequencyAdFormatFields = struct {
	Details string
	Type    string
}{
	Details: "details",
	Type:    "type",
}
