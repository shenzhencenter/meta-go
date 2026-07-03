package objects

type ProductItemError struct {
	Description   *string `json:"description,omitempty"`
	ErrorPriority *string `json:"error_priority,omitempty"`
	ErrorType     *string `json:"error_type,omitempty"`
	Title         *string `json:"title,omitempty"`
}

var ProductItemErrorFields = struct {
	Description   string
	ErrorPriority string
	ErrorType     string
	Title         string
}{
	Description:   "description",
	ErrorPriority: "error_priority",
	ErrorType:     "error_type",
	Title:         "title",
}
