package objects

type LeadGenCustomDisclaimer struct {
	Body       *LeadGenCustomDisclaimerBody   `json:"body,omitempty"`
	Checkboxes *[]LeadGenLegalContentCheckbox `json:"checkboxes,omitempty"`
	Title      *string                        `json:"title,omitempty"`
}

var LeadGenCustomDisclaimerFields = struct {
	Body       string
	Checkboxes string
	Title      string
}{
	Body:       "body",
	Checkboxes: "checkboxes",
	Title:      "title",
}
