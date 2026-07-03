package objects

type VideoStatusError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}
