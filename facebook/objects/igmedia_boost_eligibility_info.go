package objects

type IGMediaBoostEligibilityInfo struct {
	BoostIneligibleReason *string `json:"boost_ineligible_reason,omitempty"`
	EligibleToBoost       *bool   `json:"eligible_to_boost,omitempty"`
}

var IGMediaBoostEligibilityInfoFields = struct {
	BoostIneligibleReason string
	EligibleToBoost       string
}{
	BoostIneligibleReason: "boost_ineligible_reason",
	EligibleToBoost:       "eligible_to_boost",
}
