package objects

type IGShoppingReviewStatusOnsiteEligibility struct {
	IsEligible *bool                                          `json:"is_eligible,omitempty"`
	Reasons    *[]IGShoppingReviewStatusReasonWithHelpMessage `json:"reasons,omitempty"`
}
