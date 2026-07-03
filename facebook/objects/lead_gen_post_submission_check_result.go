package objects

type LeadGenPostSubmissionCheckResult struct {
	APICallResult     *string `json:"api_call_result,omitempty"`
	APIErrorMessage   *string `json:"api_error_message,omitempty"`
	ShownThankYouPage *string `json:"shown_thank_you_page,omitempty"`
}

var LeadGenPostSubmissionCheckResultFields = struct {
	APICallResult     string
	APIErrorMessage   string
	ShownThankYouPage string
}{
	APICallResult:     "api_call_result",
	APIErrorMessage:   "api_error_message",
	ShownThankYouPage: "shown_thank_you_page",
}
