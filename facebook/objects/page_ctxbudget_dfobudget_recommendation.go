package objects

type PageCTXBudgetDFOBudgetRecommendation struct {
	Budget                              *string `json:"budget,omitempty"`
	BudgetLeads                         *string `json:"budget_leads,omitempty"`
	BudgetNewModel                      *string `json:"budget_new_model,omitempty"`
	BudgetPurchases                     *string `json:"budget_purchases,omitempty"`
	BudgetValue                         *string `json:"budget_value,omitempty"`
	BudgetWithoutThreshold              *string `json:"budget_without_threshold,omitempty"`
	ReportedConversion                  *string `json:"reported_conversion,omitempty"`
	ReportedConversionLeads             *string `json:"reported_conversion_leads,omitempty"`
	ReportedConversionPurchases         *string `json:"reported_conversion_purchases,omitempty"`
	ReportedConversionValue             *string `json:"reported_conversion_value,omitempty"`
	ReportedConversionsNewModel         *string `json:"reported_conversions_new_model,omitempty"`
	ReportedConversionsWithoutThreshold *string `json:"reported_conversions_without_threshold,omitempty"`
	ZoBudget                            *string `json:"zo_budget,omitempty"`
	ZoBudgetLeads                       *string `json:"zo_budget_leads,omitempty"`
	ZoBudgetPurchases                   *string `json:"zo_budget_purchases,omitempty"`
	ZoBudgetValue                       *string `json:"zo_budget_value,omitempty"`
}
