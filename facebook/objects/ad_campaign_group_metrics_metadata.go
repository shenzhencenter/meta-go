package objects

type AdCampaignGroupMetricsMetadata struct {
	BudgetOptimization  *[]string `json:"budget_optimization,omitempty"`
	DuplicationFlowTips *[]string `json:"duplication_flow_tips,omitempty"`
}

var AdCampaignGroupMetricsMetadataFields = struct {
	BudgetOptimization  string
	DuplicationFlowTips string
}{
	BudgetOptimization:  "budget_optimization",
	DuplicationFlowTips: "duplication_flow_tips",
}
