package objects

type LeadGenCustomDisclaimer struct {
	Body       *LeadGenCustomDisclaimerBody   `json:"body,omitempty"`
	Checkboxes *[]LeadGenLegalContentCheckbox `json:"checkboxes,omitempty"`
	Title      *string                        `json:"title,omitempty"`
}
