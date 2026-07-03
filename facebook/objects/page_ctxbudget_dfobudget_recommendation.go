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

var PageCTXBudgetDFOBudgetRecommendationFields = struct {
	Budget                              string
	BudgetLeads                         string
	BudgetNewModel                      string
	BudgetPurchases                     string
	BudgetValue                         string
	BudgetWithoutThreshold              string
	ReportedConversion                  string
	ReportedConversionLeads             string
	ReportedConversionPurchases         string
	ReportedConversionValue             string
	ReportedConversionsNewModel         string
	ReportedConversionsWithoutThreshold string
	ZoBudget                            string
	ZoBudgetLeads                       string
	ZoBudgetPurchases                   string
	ZoBudgetValue                       string
}{
	Budget:                              "budget",
	BudgetLeads:                         "budget_leads",
	BudgetNewModel:                      "budget_new_model",
	BudgetPurchases:                     "budget_purchases",
	BudgetValue:                         "budget_value",
	BudgetWithoutThreshold:              "budget_without_threshold",
	ReportedConversion:                  "reported_conversion",
	ReportedConversionLeads:             "reported_conversion_leads",
	ReportedConversionPurchases:         "reported_conversion_purchases",
	ReportedConversionValue:             "reported_conversion_value",
	ReportedConversionsNewModel:         "reported_conversions_new_model",
	ReportedConversionsWithoutThreshold: "reported_conversions_without_threshold",
	ZoBudget:                            "zo_budget",
	ZoBudgetLeads:                       "zo_budget_leads",
	ZoBudgetPurchases:                   "zo_budget_purchases",
	ZoBudgetValue:                       "zo_budget_value",
}
