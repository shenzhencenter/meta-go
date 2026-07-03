package objects

type PageSettings struct {
	Setting *string                 `json:"setting,omitempty"`
	Value   *map[string]interface{} `json:"value,omitempty"`
}
