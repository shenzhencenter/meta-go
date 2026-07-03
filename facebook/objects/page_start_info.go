package objects

type PageStartInfo struct {
	Date *map[string]interface{} `json:"date,omitempty"`
	Type *string                 `json:"type,omitempty"`
}
