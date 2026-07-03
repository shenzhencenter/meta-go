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

var CreativeMulticellTestConfigFields = struct {
	BudgetPercentage       string
	BudgetSource           string
	ConfiguredCellCount    string
	DailyBudget            string
	EntrySource            string
	LeadGenForms           string
	LifetimeBudget         string
	UseExistingDailyBudget string
}{
	BudgetPercentage:       "budget_percentage",
	BudgetSource:           "budget_source",
	ConfiguredCellCount:    "configured_cell_count",
	DailyBudget:            "daily_budget",
	EntrySource:            "entry_source",
	LeadGenForms:           "lead_gen_forms",
	LifetimeBudget:         "lifetime_budget",
	UseExistingDailyBudget: "use_existing_daily_budget",
}
