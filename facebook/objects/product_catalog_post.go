package objects

type ProductCatalogPost struct {
	Firstoneof  *map[string]interface{} `json:"__firstOneOf,omitempty"`
	Secondoneof *map[string]interface{} `json:"__secondOneOf,omitempty"`
}
