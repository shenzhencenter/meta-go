package objects

type CreativeMulticellTestConfig struct {
	BudgetPercentage       *int    `json:"budget_percentage,omitempty"`
	BudgetSource           *string `json:"budget_source,omitempty"`
	ConfiguredCellCount    *int    `json:"configured_cell_count,omitempty"`
	DailyBudget            *int    `json:"daily_budget,omitempty"`
	EntrySource            *string `json:"entry_source,omitempty"`
	LeadGenForms           *string `json:"lead_gen_forms,omitempty"`
	LifetimeBudget         *int    `json:"lifetime_budget,omitempty"`
	UseExistingDailyBudget *bool   `json:"use_existing_daily_budget,omitempty"`
}
