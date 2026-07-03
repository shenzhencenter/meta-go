package objects

type ProductItemError struct {
	Description   *string `json:"description,omitempty"`
	ErrorPriority *string `json:"error_priority,omitempty"`
	ErrorType     *string `json:"error_type,omitempty"`
	Title         *string `json:"title,omitempty"`
}
