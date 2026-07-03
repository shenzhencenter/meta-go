package objects

type OverrideDetails struct {
	Key    *string                 `json:"key,omitempty"`
	Type   *string                 `json:"type,omitempty"`
	Values *map[string]interface{} `json:"values,omitempty"`
}
