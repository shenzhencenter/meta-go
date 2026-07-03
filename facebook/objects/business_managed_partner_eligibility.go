package objects

type BusinessManagedPartnerEligibility struct {
	IsEligible        *bool   `json:"is_eligible,omitempty"`
	ReasonCode        *string `json:"reason_code,omitempty"`
	ReasonDescription *string `json:"reason_description,omitempty"`
}

var BusinessManagedPartnerEligibilityFields = struct {
	IsEligible        string
	ReasonCode        string
	ReasonDescription string
}{
	IsEligible:        "is_eligible",
	ReasonCode:        "reason_code",
	ReasonDescription: "reason_description",
}
