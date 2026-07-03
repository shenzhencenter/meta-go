package objects

type LeadGenPostSubmissionCheckResult struct {
	APICallResult     *string `json:"api_call_result,omitempty"`
	APIErrorMessage   *string `json:"api_error_message,omitempty"`
	ShownThankYouPage *string `json:"shown_thank_you_page,omitempty"`
}
