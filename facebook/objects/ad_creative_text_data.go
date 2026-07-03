package objects

type AdCreativeTextData struct {
	Message *string `json:"message,omitempty"`
}

var AdCreativeTextDataFields = struct {
	Message string
}{
	Message: "message",
}
