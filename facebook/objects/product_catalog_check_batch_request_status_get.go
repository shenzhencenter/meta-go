package objects

type ProductCatalogCheckBatchRequestStatusGet struct {
	Data   *[]map[string]interface{} `json:"data,omitempty"`
	Paging *map[string]interface{}   `json:"paging,omitempty"`
}
