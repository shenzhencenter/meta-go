package objects

type IGShoppingReviewStatusReasonWithHelpMessage struct {
	Code    *string `json:"code,omitempty"`
	HelpURL *string `json:"help_url,omitempty"`
	Message *string `json:"message,omitempty"`
}

var IGShoppingReviewStatusReasonWithHelpMessageFields = struct {
	Code    string
	HelpURL string
	Message string
}{
	Code:    "code",
	HelpURL: "help_url",
	Message: "message",
}
