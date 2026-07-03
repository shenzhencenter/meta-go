package objects

type EmailImport struct {
	Lists *[]map[string]interface{} `json:"lists,omitempty"`
	Total *uint64                   `json:"total,omitempty"`
}

var EmailImportFields = struct {
	Lists string
	Total string
}{
	Lists: "lists",
	Total: "total",
}
