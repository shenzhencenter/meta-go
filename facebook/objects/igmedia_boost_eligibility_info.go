package objects

type IGMediaBoostEligibilityInfo struct {
	BoostIneligibleReason *string `json:"boost_ineligible_reason,omitempty"`
	EligibleToBoost       *bool   `json:"eligible_to_boost,omitempty"`
}
