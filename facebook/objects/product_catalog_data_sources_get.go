package objects

type ProductCatalogDataSourcesGet struct {
	Data   *[]map[string]interface{} `json:"data,omitempty"`
	Paging *map[string]interface{}   `json:"paging,omitempty"`
}

var ProductCatalogDataSourcesGetFields = struct {
	Data   string
	Paging string
}{
	Data:   "data",
	Paging: "paging",
}
