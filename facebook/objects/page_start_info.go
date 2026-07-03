package objects

type PageStartInfo struct {
	Date *map[string]interface{} `json:"date,omitempty"`
	Type *string                 `json:"type,omitempty"`
}

var PageStartInfoFields = struct {
	Date string
	Type string
}{
	Date: "date",
	Type: "type",
}
