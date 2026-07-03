package objects

type VideoStatusError struct {
	Code    *int    `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

var VideoStatusErrorFields = struct {
	Code    string
	Message string
}{
	Code:    "code",
	Message: "message",
}
