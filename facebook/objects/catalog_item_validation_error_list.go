package objects

type CatalogItemValidationErrorList struct {
	Errors *[]map[string]interface{} `json:"errors,omitempty"`
}
