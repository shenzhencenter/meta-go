package objects

type IGShoppingReviewStatus struct {
	OnsiteEligibility *IGShoppingReviewStatusOnsiteEligibility       `json:"onsite_eligibility,omitempty"`
	Reasons           *[]IGShoppingReviewStatusReasonWithHelpMessage `json:"reasons,omitempty"`
	Status            *string                                        `json:"status,omitempty"`
}

var IGShoppingReviewStatusFields = struct {
	OnsiteEligibility string
	Reasons           string
	Status            string
}{
	OnsiteEligibility: "onsite_eligibility",
	Reasons:           "reasons",
	Status:            "status",
}
