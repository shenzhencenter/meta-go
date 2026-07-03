package objects

type AdAsyncRequestSetNotificationResult struct {
	Response *string `json:"response,omitempty"`
	Status   *string `json:"status,omitempty"`
}

var AdAsyncRequestSetNotificationResultFields = struct {
	Response string
	Status   string
}{
	Response: "response",
	Status:   "status",
}
