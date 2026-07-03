package objects

type ProductCatalogHotelRoomsBatch struct {
	Errors           *[]map[string]interface{} `json:"errors,omitempty"`
	ErrorsTotalCount *int                      `json:"errors_total_count,omitempty"`
	Handle           *string                   `json:"handle,omitempty"`
	Status           *string                   `json:"status,omitempty"`
}
