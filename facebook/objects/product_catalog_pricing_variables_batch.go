package objects

type ProductCatalogPricingVariablesBatch struct {
	Errors           *[]map[string]interface{} `json:"errors,omitempty"`
	ErrorsTotalCount *int                      `json:"errors_total_count,omitempty"`
	Handle           *string                   `json:"handle,omitempty"`
	Status           *string                   `json:"status,omitempty"`
}

var ProductCatalogPricingVariablesBatchFields = struct {
	Errors           string
	ErrorsTotalCount string
	Handle           string
	Status           string
}{
	Errors:           "errors",
	ErrorsTotalCount: "errors_total_count",
	Handle:           "handle",
	Status:           "status",
}
