package objects

type OverrideDetails struct {
	Key    *string                 `json:"key,omitempty"`
	Type   *string                 `json:"type,omitempty"`
	Values *map[string]interface{} `json:"values,omitempty"`
}

var OverrideDetailsFields = struct {
	Key    string
	Type   string
	Values string
}{
	Key:    "key",
	Type:   "type",
	Values: "values",
}
