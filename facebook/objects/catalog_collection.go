package objects

type CatalogCollection struct {
	Description *string `json:"description,omitempty"`
	Title       *string `json:"title,omitempty"`
	URL         *string `json:"url,omitempty"`
}

var CatalogCollectionFields = struct {
	Description string
	Title       string
	URL         string
}{
	Description: "description",
	Title:       "title",
	URL:         "url",
}
