package objects

type PageGetStartedNullstate struct {
	CtaTitle          *string `json:"cta_title,omitempty"`
	ProcessedGreeting *string `json:"processed_greeting,omitempty"`
	Responsiveness    *string `json:"responsiveness,omitempty"`
}

var PageGetStartedNullstateFields = struct {
	CtaTitle          string
	ProcessedGreeting string
	Responsiveness    string
}{
	CtaTitle:          "cta_title",
	ProcessedGreeting: "processed_greeting",
	Responsiveness:    "responsiveness",
}
