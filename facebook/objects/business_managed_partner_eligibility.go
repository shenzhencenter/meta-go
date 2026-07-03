package objects

type BusinessManagedPartnerEligibility struct {
	IsEligible        *bool   `json:"is_eligible,omitempty"`
	ReasonCode        *string `json:"reason_code,omitempty"`
	ReasonDescription *string `json:"reason_description,omitempty"`
}
