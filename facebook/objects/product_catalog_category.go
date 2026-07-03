package objects

type ProductCatalogCategory struct {
	CriteriaValue  *string              `json:"criteria_value,omitempty"`
	Description    *string              `json:"description,omitempty"`
	DestinationURI *string              `json:"destination_uri,omitempty"`
	ImageURL       *string              `json:"image_url,omitempty"`
	Name           *string              `json:"name,omitempty"`
	NumItems       *int                 `json:"num_items,omitempty"`
	Tokens         *[]map[string]string `json:"tokens,omitempty"`
}

var ProductCatalogCategoryFields = struct {
	CriteriaValue  string
	Description    string
	DestinationURI string
	ImageURL       string
	Name           string
	NumItems       string
	Tokens         string
}{
	CriteriaValue:  "criteria_value",
	Description:    "description",
	DestinationURI: "destination_uri",
	ImageURL:       "image_url",
	Name:           "name",
	NumItems:       "num_items",
	Tokens:         "tokens",
}
