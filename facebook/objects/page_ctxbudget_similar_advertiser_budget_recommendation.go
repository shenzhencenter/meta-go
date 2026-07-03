package objects

type PageCTXBudgetSimilarAdvertiserBudgetRecommendation struct {
	Budget                              *string `json:"budget,omitempty"`
	BudgetNewModel                      *string `json:"budget_new_model,omitempty"`
	BudgetWithoutThreshold              *string `json:"budget_without_threshold,omitempty"`
	ReportedConversion                  *string `json:"reported_conversion,omitempty"`
	ReportedConversionsNewModel         *string `json:"reported_conversions_new_model,omitempty"`
	ReportedConversionsWithoutThreshold *string `json:"reported_conversions_without_threshold,omitempty"`
}
