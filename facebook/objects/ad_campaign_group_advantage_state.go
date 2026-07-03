package objects

type AdCampaignGroupAdvantageState struct {
	AdvantageAudienceState  *string `json:"advantage_audience_state,omitempty"`
	AdvantageBudgetState    *string `json:"advantage_budget_state,omitempty"`
	AdvantagePlacementState *string `json:"advantage_placement_state,omitempty"`
	AdvantageState          *string `json:"advantage_state,omitempty"`
}
