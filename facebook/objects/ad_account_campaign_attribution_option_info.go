package objects

type AdAccountCampaignAttributionOptionInfo struct {
	IsEligible *bool   `json:"is_eligible,omitempty"`
	Value      *string `json:"value,omitempty"`
}

var AdAccountCampaignAttributionOptionInfoFields = struct {
	IsEligible string
	Value      string
}{
	IsEligible: "is_eligible",
	Value:      "value",
}
