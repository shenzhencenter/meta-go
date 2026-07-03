package objects

type BusinessProductCatalogTOS struct {
	Accepted *bool   `json:"accepted,omitempty"`
	Content  *string `json:"content,omitempty"`
}

var BusinessProductCatalogTOSFields = struct {
	Accepted string
	Content  string
}{
	Accepted: "accepted",
	Content:  "content",
}
