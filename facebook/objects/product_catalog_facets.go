package objects

type ProductCatalogFacets struct {
	Facets    *[]map[string][]map[string]interface{} `json:"facets,omitempty"`
	ItemCount *int                                   `json:"item_count,omitempty"`
}
