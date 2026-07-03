package objects

type IGShoppingReviewStatusOnsiteEligibility struct {
	IsEligible *bool                                          `json:"is_eligible,omitempty"`
	Reasons    *[]IGShoppingReviewStatusReasonWithHelpMessage `json:"reasons,omitempty"`
}

var IGShoppingReviewStatusOnsiteEligibilityFields = struct {
	IsEligible string
	Reasons    string
}{
	IsEligible: "is_eligible",
	Reasons:    "reasons",
}
