package objects

type AdCampaignGroupAdvantageState struct {
	AdvantageAudienceState  *string `json:"advantage_audience_state,omitempty"`
	AdvantageBudgetState    *string `json:"advantage_budget_state,omitempty"`
	AdvantagePlacementState *string `json:"advantage_placement_state,omitempty"`
	AdvantageState          *string `json:"advantage_state,omitempty"`
}

var AdCampaignGroupAdvantageStateFields = struct {
	AdvantageAudienceState  string
	AdvantageBudgetState    string
	AdvantagePlacementState string
	AdvantageState          string
}{
	AdvantageAudienceState:  "advantage_audience_state",
	AdvantageBudgetState:    "advantage_budget_state",
	AdvantagePlacementState: "advantage_placement_state",
	AdvantageState:          "advantage_state",
}
